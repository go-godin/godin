package godin

var (
	serviceInterfaceTemplate = &BaseTemplate{
		Config: &TemplateConfiguration{
			Name:       "service-interface",
			SourceFile: "service/interface.go.tmpl",
			TargetFile: "internal/service/service.go",
			GoSource:   true,
			Skip:       false,
		},
	}
)

type ServiceInterfaceModule struct {
	ServiceInterfaceTemplate Template
	*serviceInterfaceConfig
}

type serviceInterfaceConfig struct {
	Enabled bool
}

func NewServiceInterfaceModule() Module {
	return &ServiceInterfaceModule{
		ServiceInterfaceTemplate: serviceInterfaceTemplate,
		serviceInterfaceConfig: &serviceInterfaceConfig{true},
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
	}
}

func (e ServiceInterfaceModule) OutputPaths() []string {
	return []string{
		e.ServiceInterfaceTemplate.Configuration().TargetFile,
	}
}

func (e ServiceInterfaceModule) Install() error {
	return nil
}

func (e ServiceInterfaceModule) Generate(projectContext interface{}, protobufContext interface{}, templateRootPath, outputRootPath string) error {
	if err := e.ServiceInterfaceTemplate.Render(projectContext, protobufContext, e.serviceInterfaceConfig, templateRootPath, outputRootPath); err != nil {
		return err
	}
	return nil
}
