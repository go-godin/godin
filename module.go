package godin

import (
	"fmt"
)

// Module defines the interface of default godin modules
type Module interface {
	Configurable
	Addable
	Removable

	// Name should return a printable name of the module. Ideally this name is namespaced.
	Name() string

	// PrepareContext enables the module to modify the default Context in order to add the information it needs.
	PrepareContext(ctx *Context) interface{}

	// Generate is executed when 'godin generate' is called
	Generate(ctx *Context, templateRootPath, outputRootPath string) error
}

// Addable defines behaviour of anything which can be enabled and checked for
type Addable interface {
	Add()
	IsEnabled() bool
}

// Delete defines behaviour of anything which can be disabled
type Removable interface {
	Remove()
}

// ModuleRegistry defines the interface of how modules are registered, stored and retrieved.
type ModuleRegistry interface {
	Register(module Module) error
	Get(key string) (Module, error)
	GetEnabled() []Module
	Modules() []Module
	Keys() []string
}

// BaseModule defines the default behaviour of godin modules.
type BaseModule struct {
}


// DefaultRegistry implements the default ModuleRegistry.
type DefaultRegistry struct {
	modules []Module
}

// Register will register a new module. The module's Key() must be unique.
// The module is then added to the 'modules' field.
func (reg *DefaultRegistry) Register(module Module) error {
	if module.Key() == "" {
		return fmt.Errorf("the module must provide a non-empty key")
	}

	if reg.IsRegistered(module.Key()) {
		return fmt.Errorf("a module with the key '%s' is already registered", module.Key())
	}

	reg.modules = append(reg.modules, module)

	return nil
}

// Get returns a module by key if it's registered.
func (reg *DefaultRegistry) Get(key string) (Module, error) {
	for _, mod := range reg.modules {
		if mod.Key() == key {
			return mod, nil
		}
	}
	return nil, fmt.Errorf("no module '%s' found", key)
}

// IsRegistered returns 'true' if a module with the given key is registered. Otherwise 'false' is returned.
func (reg *DefaultRegistry) IsRegistered(key string) bool {
	for _, k := range reg.Keys() {
		if k == key {
			return true
		}
	}
	return false
}

// Keys returns a slice of keys of all registered modules
func (reg *DefaultRegistry) Keys() []string {
	var keys []string
	for _, mod := range reg.modules {
		keys = append(keys, mod.Key())
	}

	return keys
}

// GetEnabled iterates over all registered modules and calls their IsEnabled().
// All enabled modules are then returned.
func (reg *DefaultRegistry) GetEnabled() []Module {
	var enabled []Module
	for _, mod := range reg.modules {
		if mod.IsEnabled() {
			enabled = append(enabled, mod)
		}
	}

	return enabled
}

// Modules returns all registered modules
func (reg *DefaultRegistry) Modules() []Module {
	return reg.modules
}
