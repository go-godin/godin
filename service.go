package godin

import (
	"strings"
)

var (
	serviceInterfaceTemplate = &BaseTemplate{
		Config: &TemplateConfiguration{
			Name:       "service-interface",
			SourceFile: "service/interface.go.tmpl",
			TargetFile: "internal/<ServiceName>/service.go",
			GoSource:   true,
			Skip:       false,
		},
	}
	modelsTemplate = &BaseTemplate{
		Config: &TemplateConfiguration{
			Name:       "service-models",
			SourceFile: "service/models.go.tmpl",
			TargetFile: "internal/<ServiceName>/models.go",
			GoSource:   true,
			Skip:       false,
		},
	}
)

type ServiceInterfaceModule struct {
	ServiceInterfaceTemplate Template
	ServiceModelsTemplate    Template
	*serviceInterfaceConfig
}

type serviceInterfaceConfig struct {
	Enabled bool
}

func NewServiceInterfaceModule(serviceName string) Module {
	tf := serviceInterfaceTemplate.Config.TargetFile
	serviceInterfaceTemplate.Config.TargetFile = strings.Replace(tf, "<ServiceName>", serviceName, 1)

	tf = modelsTemplate.Config.TargetFile
	modelsTemplate.Config.TargetFile = strings.Replace(tf, "<ServiceName>", serviceName, 1)

	return &ServiceInterfaceModule{
		ServiceInterfaceTemplate: serviceInterfaceTemplate,
		ServiceModelsTemplate:    modelsTemplate,
		serviceInterfaceConfig:   &serviceInterfaceConfig{true},
	}
}

func (e ServiceInterfaceModule) Identifier() string {
	return "service.interface"
}

func (e ServiceInterfaceModule) Configuration() interface{} {
	return e.serviceInterfaceConfig
}

func (e ServiceInterfaceModule) Configure(source ResolvableConfig) error {
	cfg := &serviceInterfaceConfig{}
	if err := source.Unmarshal(e.Identifier(), cfg); err != nil {
		return err
	}
	e.serviceInterfaceConfig = cfg
	return nil
}

func (e ServiceInterfaceModule) Templates() []Template {
	return []Template{
		serviceInterfaceTemplate,
		modelsTemplate,
	}
}

func (e ServiceInterfaceModule) OutputPaths() []string {
	return []string{
		e.ServiceInterfaceTemplate.Configuration().TargetFile,
		e.ServiceModelsTemplate.Configuration().TargetFile,
	}
}

func (e ServiceInterfaceModule) Install() error {
	return nil
}

func (e ServiceInterfaceModule) Generate(projectContext interface{}, protobufContext interface{}, templateRootPath, outputRootPath string) error {
	if err := e.ServiceInterfaceTemplate.Render(projectContext, protobufContext, e.serviceInterfaceConfig, templateRootPath, outputRootPath); err != nil {
		return err
	}

	if e.ServiceModelsTemplate.Configuration().TargetExists(outputRootPath) {
		e.ServiceModelsTemplate.Configuration().Skip = true
	}
	if err := e.ServiceModelsTemplate.Render(projectContext, protobufContext, e.serviceInterfaceConfig, templateRootPath, outputRootPath); err != nil {
		return err
	}
	return nil
}
