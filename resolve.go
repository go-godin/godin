package godin

type Resolvable interface {
	IsSet(key string) bool
	Get(key string) interface{}
}

type ModuleResolver struct {
}

var moduleNameTypes = map[string]Type{
	"transport.grpc.server": TransportGrpcServer,
}

func (res *ModuleResolver) ResolveAll(source Resolvable) (modules []Module, err error) {
	for key, _ := range moduleNameTypes {
		if source.IsSet(key) {
			mod := ModuleFactory(moduleNameTypes[key])
			modules = append(modules, mod)
		}
	}
	return modules, nil
}

func ModuleFactory(moduleType Type) Module {
	switch moduleType {
	case TransportGrpcServer:
		return NewGrpcServerModule()
	default:
		return nil
	}
}
