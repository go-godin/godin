package module

// GrpcServerModule provides all templates for the gRPC server transport layer
type GrpcServerModule struct {
	*BaseModule
	grpcServerConfig
	ServerTemplate Template
}

// grpcServerConfig defines the transport.grpc.server module configuration struct.
type grpcServerConfig struct {
	DefaultPort    int    `yaml:"defaultPort"`
	DefaultAddress string `yaml:"defaultAddress"`
}

// NewGrpcServerModule returns a new pre-initialized GrpcServerModule.
// The module returns reasonable defaults so that it could be used directly.
// Upon calling Load(), the configuration will be loaded which eventually overwrites the defaults.
func NewGrpcServerModule() Module {
	return &GrpcServerModule{
		BaseModule: &BaseModule{
			ModuleName: "grpc-server",
		},
		grpcServerConfig: grpcServerConfig{
			DefaultPort:    50051,
			DefaultAddress: "0.0.0.0",
		},
		ServerTemplate: grpcServerTemplate,
	}
}

func (mod *GrpcServerModule) Install() error {
	return nil
}

// Generate will render the modules templates
func (mod *GrpcServerModule) Generate(protobufContext interface{}, templateRootPath, outputRootPath string) error {
	if err := mod.ServerTemplate.Render(protobufContext, mod.grpcServerConfig, templateRootPath, outputRootPath); err != nil {
		return err
	}
	return nil
}

// Configuration returns the module configuration as interface.
func (mod *GrpcServerModule) Configuration() interface{} {
	return mod.grpcServerConfig
}

// ConfigurationKey returns the GrpcServerModule key. It's used as unique identifier and - though only internally - also
// used as configuration key. Note that the ConfigurationKey() must not always correspond to the configuration key used by the
// module.
func (cfg *grpcServerConfig) ConfigurationKey() string {
	return "transport.grpc.server"
}

var (
	grpcServerTemplate = &BaseTemplate{
		Config: &TemplateConfiguration{
			Name:       "grpc-server",
			SourceFile: "transport/grpc/server.go.tmpl",
			TargetFile: "transport/grpc/server.go",
			GoSource:   true,
			Skip:       false,
		},
	}
)
