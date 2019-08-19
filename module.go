package godin

type Type int

const (
	TransportGrpcServer Type = iota
)

type ResolvableConfig interface {
	Resolvable
	Unmarshal(key string, target interface{}) error
}

// Module defines the interface of default godin enabledModules
// TODO: a module should provide it's cli flags
type Module interface {
	ConfigProvider

	Configure(source ResolvableConfig) error

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

// Configurable defines the interface of anything which can be configured must behave.
type Configurable interface {
	ConfigurationKey() string
}
