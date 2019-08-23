package godin

type Type int

const (
	TransportGrpcServer Type = iota
	TransportGrpcClient
)

type ResolvableConfig interface {
	Resolvable
	Unmarshal(key string, target interface{}) error
}

// Module defines the interface of default godin enabledModules
type Module interface {
	ConfigProvider

	// Configure is used to initialize a module based on some ResolveableConfig which can be
	// unmarshalled into the module's own configuration struct.
	Configure(source ResolvableConfig) error

	// Templates returns all template object which the module uses.
	Templates() []Template

	// OutputPaths returns a list of paths which must exist in order for the module to correctly generate the
	// templates.
	OutputPaths() []string

	// Install hook is called when 'godin add' is executed for that module. The hook enables the module to interfere
	// and prepare the module (e.g. prompt for values).
	// The installation is considered a success if error == nil.
	Install() error

	// Generate is executed when 'godin generate' is called
	Generate(protobufContext interface{}, templateRootPath, outputRootPath string) error
}
