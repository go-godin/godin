package godin

var (
	grpcServerTemplate = &BaseTemplate{
		Config: &TemplateConfiguration{
			Name:       "grpc-server",
			SourceFile: "transport/grpc/server.go.tmpl",
			TargetFile: "internal/transport/grpc/server.go",
			GoSource:   true,
			Skip:       false,
		},
	}
	grpcClientTemplate = &BaseTemplate{
		Config: &TemplateConfiguration{
			Name:       "grpc-client",
			SourceFile: "transport/grpc/client.go.tmpl",
			TargetFile: "internal/transport/grpc/client.go",
			GoSource:   true,
			Skip:       false,
		},
	}
)

// GrpcServerModule provides all templates for the gRPC server transport layer
type GrpcServerModule struct {
	*grpcServerConfig
	ServerTemplate Template
}

// grpcServerConfig defines the transport.grpc.server module configuration struct.
type grpcServerConfig struct {
	DefaultPort    int    `yaml:"defaultPort"`
	DefaultAddress string `yaml:"defaultAddress"`
}

// NewGrpcServerModule returns a new pre-initialized GrpcServerModule.
// The module returns reasonable defaults so that it could be used directly.
// Upon calling Initialize(), the configuration will be loaded which eventually overwrites the defaults.
func NewGrpcServerModule() Module {
	return &GrpcServerModule{
		grpcServerConfig: &grpcServerConfig{
			DefaultPort:    50051,
			DefaultAddress: "0.0.0.0",
		},
		ServerTemplate: grpcServerTemplate,
	}
}

func (mod *GrpcServerModule) Install() error {
	return nil
}

func (mod *GrpcServerModule) Templates() (tpl []Template) {
	tpl = append(tpl, grpcServerTemplate)

	return tpl
}

func (mod *GrpcServerModule) Configure(source ResolvableConfig) error {
	cfg := &grpcServerConfig{}
	if err := source.Unmarshal(mod.Identifier(), cfg); err != nil {
		return err
	}
	mod.grpcServerConfig = cfg

	return nil
}

// Generate will render the modules templates
func (mod *GrpcServerModule) Generate(projectContext, protobufContext interface{}, templateRootPath, outputRootPath string) error {
	if err := mod.ServerTemplate.Render(projectContext, protobufContext, mod.grpcServerConfig, templateRootPath, outputRootPath); err != nil {
		return err
	}
	return nil
}

// Configuration returns the module configuration as interface.
func (mod *GrpcServerModule) Configuration() interface{} {
	return mod.grpcServerConfig
}

// Identifier returns the GrpcServerModule key. It's used as unique identifier and - though only internally - also
// used as configuration key. Note that the Identifier() must not always correspond to the configuration key used by the
// module.
func (mod *GrpcServerModule) Identifier() string {
	return "transport.grpc.server"
}

func (mod *GrpcServerModule) OutputPaths() (paths []string) {
	paths = append(paths, grpcServerTemplate.Configuration().TargetFile)

	return paths
}

// GrpcServerModule provides all templates for the gRPC server transport layer
type GrpcClientModule struct {
	*grpcClientConfig
	ClientTemplate Template
	project        ProjectConfiguration
}

// grpcServerConfig defines the transport.grpc.server module configuration struct.
type grpcClientConfig struct {
	RetryCount int `yaml:"retryCount"`
}

// NewGrpcServerModule returns a new pre-initialized GrpcServerModule.
// The module returns reasonable defaults so that it could be used directly.
// Upon calling Initialize(), the configuration will be loaded which eventually overwrites the defaults.
func NewGrpcClientModule(configuration ProjectConfiguration) Module {
	return &GrpcClientModule{
		grpcClientConfig: &grpcClientConfig{
			RetryCount: 3,
		},
		ClientTemplate: grpcClientTemplate,
		project:        configuration,
	}
}

func (mod *GrpcClientModule) Install() error {
	return nil
}

func (mod *GrpcClientModule) Templates() (tpl []Template) {
	tpl = append(tpl, grpcClientTemplate)

	return tpl
}

func (mod *GrpcClientModule) Configure(source ResolvableConfig) error {
	cfg := &grpcClientConfig{}
	if err := source.Unmarshal(mod.Identifier(), cfg); err != nil {
		return err
	}
	mod.grpcClientConfig = cfg

	return nil
}

// Generate will render the modules templates
func (mod *GrpcClientModule) Generate(projectContext, protobufContext interface{}, templateRootPath, outputRootPath string) error {
	if err := mod.ClientTemplate.Render(projectContext, protobufContext, mod.grpcClientConfig, templateRootPath, outputRootPath); err != nil {
		return err
	}
	return nil
}

// Configuration returns the module configuration as interface.
func (mod *GrpcClientModule) Configuration() interface{} {
	return mod.grpcClientConfig
}

// Identifier returns the GrpcServerModule key. It's used as unique identifier and - though only internally - also
// used as configuration key. Note that the Identifier() must not always correspond to the configuration key used by the
// module.
func (mod *GrpcClientModule) Identifier() string {
	return "transport.grpc.client"
}

func (mod *GrpcClientModule) OutputPaths() (paths []string) {
	paths = append(paths, grpcServerTemplate.Configuration().TargetFile)

	return paths
}
