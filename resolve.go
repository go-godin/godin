package godin

import (
	"fmt"
)

type Resolvable interface {
	IsSet(key string) bool
	Get(key string) interface{}
}

type moduleResolver struct {
	project ProjectConfiguration
}

func NewModuleResolver(configuration *ProjectConfiguration) *moduleResolver {
	return &moduleResolver{project: *configuration}
}

var moduleNameTypes = map[string]Type{
	"service.endpoints":     Endpoints,
	"transport.grpc.server": TransportGrpcServer,
	"transport.grpc.client": TransportGrpcClient,
}

func (res *moduleResolver) ResolveAll(source Resolvable) (modules []Module, err error) {
	for key := range moduleNameTypes {
		if source.IsSet(key) {
			mod := res.Factory(moduleNameTypes[key])
			modules = append(modules, mod)
		}
	}
	return modules, nil
}

func (res *moduleResolver) Resolve(moduleName string) (Module, error) {
	for key, typ := range moduleNameTypes {
		if key == moduleName {
			return res.Factory(typ), nil
		}
	}
	return nil, fmt.Errorf("module '%s' cannot be resolved", moduleName)
}

func (res *moduleResolver) Factory(moduleType Type) Module {
	switch moduleType {
	case Endpoints:
		return NewEndpointsModule()
	case TransportGrpcServer:
		return NewGrpcServerModule()
	case TransportGrpcClient:
		return NewGrpcClientModule(res.project)
	default:
		return nil
	}
}
