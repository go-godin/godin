package godin

import (
	"fmt"
)

// Registry defines the interface of how enabledModules are registered, stored and retrieved.
type Registry interface {
	Register(module Module) error
	Modules() Store
	Get(key, identifier string) (Module, error)
}

type Store []Module

// defaultRegistry implements the default Registry.
type defaultRegistry struct {
	modules map[string]Module
}

func NewRegistry() Registry {
	return &defaultRegistry{
		modules: make(map[string]Module),
	}
}

// Register will register a new module. The module's Identifier() must be unique.
// The module is then added to the 'enabledModules' field.
func (reg *defaultRegistry) Register(module Module) error {
	if module.Identifier() == "" {
		return fmt.Errorf("the module must provide a non-empty configuration key")
	}

	if reg.isRegistered(module.Identifier()) {
		return fmt.Errorf("a module with the identifier '%s' is already registered", module.Identifier())
	}

	reg.modules[module.Identifier()] = module

	return nil
}

// Get returns a module by key and identifier if it's registered.
func (reg *defaultRegistry) Get(key, identifier string) (Module, error) {
	if module, ok := reg.modules[key]; ok {
		return module, nil
	}
	return nil, fmt.Errorf("no module '%s' with ID '%s' found", key, identifier)
}

// isRegistered returns 'true' if a module with the given key and identifier is registered. Otherwise 'false' is returned.
func (reg *defaultRegistry) isRegistered(key string) bool {
	if _, ok := reg.modules[key]; ok {
		return true
	}
	return false
}

// enabledModules returns all registered enabledModules
func (reg *defaultRegistry) Modules() Store {
	var modules []Module
	for _, module := range reg.modules {
		modules = append(modules, module)
	}
	return modules
}
