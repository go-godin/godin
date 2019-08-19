package godin

import (
	"fmt"
)

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
	for key := range moduleNameTypes {
		if source.IsSet(key) {
			mod := ModuleFactory(moduleNameTypes[key])
			modules = append(modules, mod)
		}
	}
	return modules, nil
}

func (res *ModuleResolver) Resolve(moduleName string) (Module, error) {
	for key, typ := range moduleNameTypes {
		if key == moduleName {
			return ModuleFactory(typ), nil
		}
	}
	return nil, fmt.Errorf("module '%s' cannot be resolved", moduleName)
}

func ModuleFactory(moduleType Type) Module {
	switch moduleType {
	case TransportGrpcServer:
		return NewGrpcServerModule()
	default:
		return nil
	}
}
