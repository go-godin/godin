package module

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/pkg/errors"
	"gitub.com/go-godin/godin/file"
)

type Template interface {
	Configuration() *TemplateConfiguration
	Render(protobufContext interface{}, moduleConfig interface{}, templateRootPath, outputRootPath string) error
}

// TemplateConfiguration specifies the base configuration for each template.
type TemplateConfiguration struct {
	Name       string
	SourceFile string
	TargetFile string
	GoSource   bool
	Skip       bool
}

// SourceExists checks - given the path to the templates - if the template source file exists
func (cfg *TemplateConfiguration) SourceExists(templateDir string) bool {
	if _, err := os.Stat(filepath.Join(templateDir, cfg.SourceFile)); os.IsNotExist(err) {
		return false
	}
	return true
}

// EnsureTargetPath ensures that the template's target path exists and is writeable by the current user.
// It will create missing folders.
func (cfg *TemplateConfiguration) EnsureTargetPath(outputDir string) error {

	// Check if the target folder exists
	targetPath := filepath.Join(outputDir, cfg.TargetFile)
	targetFolder := filepath.Dir(targetPath)
	if _, err := os.Stat(targetFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(targetFolder, 0755); err != nil {
			return err
		}
	}

	// ensure permission for existing
	testFile := filepath.Join(targetFolder, "tester")
	if _, err := os.Create(testFile); os.IsPermission(err) {
		if err := os.Chmod(targetFolder, 0755); err != nil {
			err = errors.Wrap(err, "target path not writeable and chmod failed")
			return err
		}
	}
	if err := os.Remove(testFile); err != nil {
		return errors.Wrap(err, "failed to remove permission tester")
	}

	return nil
}

// BaseTemplate defines some useful default behaviour for module templates
type BaseTemplate struct {
	Config *TemplateConfiguration
}

// Configuration returns the TemplateConfiguration
func (tpl *BaseTemplate) Configuration() *TemplateConfiguration {
	return tpl.Config
}

func (tpl *BaseTemplate) Render(protobufContext interface{}, moduleConfig interface{}, templateRootPath, outputRootPath string) error {
	if tpl.Config.Skip {
		fmt.Printf("[-] template disabled: %s\n", tpl.Config.SourceFile)
		return nil
	}
	fmt.Printf("[+] template enabled: %s\n", tpl.Config.SourceFile)

	if !tpl.Config.SourceExists(templateRootPath) {
		return fmt.Errorf("source template not found %s: %s", tpl.Config.SourceFile)
	}
	fmt.Printf("    -> template found \n")

	render := NewTemplateRenderer(*tpl.Config, templateRootPath)
	output, err := render.Render(tpl.prepareContext(protobufContext, moduleConfig))
	if err != nil {
		fmt.Println(err)
	}

	// write targetFile
	targetPath := path.Join(outputRootPath, tpl.Config.TargetFile)
	writer := file.NewFileWriter(targetPath, output)
	if err := writer.Write(true); err != nil {
		return fmt.Errorf("failed to write template '%s': %s", tpl.Config.SourceFile, err)
	}
	fmt.Printf("    -> target file written: %s\n", tpl.Config.TargetFile)

	/*
		if err := tpl.Config.EnsureTargetPath(app.OutputPath()); err != nil {
			fmt.Println(err)
		}

	*/
	return nil
}

// prepareContext aggregates the protobuf context (global context) with the module and template configuration.
func (tpl *BaseTemplate) prepareContext(protobufContext interface{}, moduleConfig interface{}) interface{} {
	return struct {
		CTX interface{}
		TPL *TemplateConfiguration
		MOD interface{}
	}{
		CTX: protobufContext,
		TPL: tpl.Config,
		MOD: moduleConfig,
	}
}

type TemplateRenderer struct {
	templateRootPath string
	template         TemplateConfiguration
}

func NewTemplateRenderer(config TemplateConfiguration, templateSource string) *TemplateRenderer {
	return &TemplateRenderer{
		template:         config,
		templateRootPath: templateSource,
	}
}

// Render the template given the template configuration and return the rendered buffer.
// If a template is configured to be a Go-source file, the rendered output will be formatted using go/format before returning.
//
// The given templateContext will be passed on template execution and is thus available from within the template.
// In addition to the templateContext, the 'github.com/Masterminds/sprig' template functions are injected to further increase
// template productivity.
func (r *TemplateRenderer) Render(templateContext interface{}) (rendered []byte, err error) {
	templatePath := filepath.Join(r.templateRootPath, r.template.SourceFile)

	buf, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}

	tpl := template.New(r.template.Name).Funcs(sprig.TxtFuncMap())
	tpl, err = tpl.Parse(string(buf))
	if err != nil {
		return nil, err
	}

	out := bytes.Buffer{}
	if err := tpl.Execute(&out, templateContext); err != nil {
		return nil, err
	}

	if r.template.GoSource {
		formatted, err := format.Source(out.Bytes())
		if err != nil {
			return nil, err
		}

		out = *bytes.NewBuffer(formatted)
	}

	return out.Bytes(), nil
}
