package godin

var (
	endpointsTemplate = &BaseTemplate{
		Config: &TemplateConfiguration{
			Name:       "endpoints",
			SourceFile: "endpoint/endpoints.go.tmpl",
			TargetFile: "internal/endpoint/endpoints.go",
			GoSource:   true,
			Skip:       false,
		},
	}
	requestResponseTemplate = &BaseTemplate{
		Config: &TemplateConfiguration{
			Name:       "request_response",
			SourceFile: "endpoint/request_response.go.tmpl",
			TargetFile: "internal/endpoint/request_response.go",
			GoSource:   true,
			Skip:       false,
		},
	}
)

type EndpointsModule struct {
	EndpointsTemplate       Template
	RequestResponseTemplate Template
	*endpointsConfig
}

type endpointsConfig struct {
	ZipkinMiddleware bool `yaml:"zipkinMiddleware"`
}

func NewEndpointsModule() Module {
	return &EndpointsModule{
		EndpointsTemplate:       endpointsTemplate,
		RequestResponseTemplate: requestResponseTemplate,
		endpointsConfig:         &endpointsConfig{ZipkinMiddleware: false},
	}
}

func (e EndpointsModule) Identifier() string {
	return "service.endpoints"
}

func (e EndpointsModule) Configuration() interface{} {
	return e.endpointsConfig
}

func (e EndpointsModule) Configure(source ResolvableConfig) error {
	cfg := &endpointsConfig{}
	if err := source.Unmarshal(e.Identifier(), cfg); err != nil {
		return err
	}
	e.endpointsConfig = cfg
	return nil
}

func (e EndpointsModule) Templates() []Template {
	return []Template{endpointsTemplate}
}

func (e EndpointsModule) OutputPaths() []string {
	return []string{e.EndpointsTemplate.Configuration().TargetFile}
}

func (e EndpointsModule) Install() error {
	return nil
}

func (e EndpointsModule) Generate(projectContext interface{}, protobufContext interface{}, templateRootPath, outputRootPath string) error {
	if err := e.EndpointsTemplate.Render(projectContext, protobufContext, e.endpointsConfig, templateRootPath, outputRootPath); err != nil {
		return err
	}
	if err := e.RequestResponseTemplate.Render(projectContext, protobufContext, e.endpointsConfig, templateRootPath, outputRootPath); err != nil {
		return err
	}
	return nil
}
