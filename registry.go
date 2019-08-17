package godin

import (
	"fmt"
)

// ModuleRegistry defines the interface of how enabledModules are registered, stored and retrieved.
type ModuleRegistry interface {
	Register(module Module) error
	Modules() ModuleStore
}

type EnabledModuleRegistry interface {
	ModuleRegistry
	Get(key, identifier string) (Module, error)
}

type AvailableModuleRegistry interface {
	ModuleRegistry
	FindModule(key string) (Module, error)
}

type ModuleStore []Module

// EnabledRegistry implements the default ModuleRegistry.
type EnabledRegistry struct {
	modules map[string][]Module
}

func NewEnabledRegistry() EnabledModuleRegistry {
	return &EnabledRegistry{
		modules: make(map[string][]Module),
	}
}

// Register will register a new module. The module's ConfigurationKey() must be unique.
// The module is then added to the 'enabledModules' field.
func (reg *EnabledRegistry) Register(module Module) error {
	if module.ConfigurationKey() == "" {
		return fmt.Errorf("the module must provide a non-empty configuration key")
	}

	if reg.IsRegistered(module.ConfigurationKey(), module.ID()) {
		return fmt.Errorf("a module with the key '%s'  and ID '%s' is already registered", module.ConfigurationKey(), module.ID())
	}

	reg.modules[module.ConfigurationKey()] = append(reg.modules[module.ConfigurationKey()], module)

	return nil
}

// Get returns a module by key and identifier if it's registered.
func (reg *EnabledRegistry) Get(key, identifier string) (Module, error) {
	if group, ok := reg.modules[key]; ok {
		for _, module := range group {
			if module.ID() == identifier {
				return module, nil
			}
		}
	}
	return nil, fmt.Errorf("no module '%s' with ID '%s' found", key, identifier)
}

// IsRegistered returns 'true' if a module with the given key and identifier is registered. Otherwise 'false' is returned.
func (reg *EnabledRegistry) IsRegistered(key, identifier string) bool {
	if group, ok := reg.modules[key]; ok {
		for _, module := range group {
			if module.ID() == identifier {
				return true
			}
		}
	}
	return false
}


// Modules returns all registered enabledModules
func (reg *EnabledRegistry) Modules() ModuleStore {
	var modules []Module
	for _, group := range reg.modules {
		for _, module := range group {
			modules = append(modules, module)
		}
	}
	return modules
}

type AvailableRegistry struct {
	modules []Module
}

func (reg *AvailableRegistry) Register(module Module) error {
	if reg.isRegistered(module) {
		return fmt.Errorf("module '%s' is already registered", module.ConfigurationKey())
	}
	reg.modules = append(reg.modules, module)

	return nil
}

func (reg *AvailableRegistry) FindModule(key string) (Module, error) {
	for _, module := range reg.modules {
		if module.ConfigurationKey() == key {
			return module, nil
		}
	}
	return nil, fmt.Errorf("module '%s' is not registered", key)
}

// Modules returns all registered enabledModules
func (reg *AvailableRegistry) Modules() ModuleStore {
	return reg.modules
}

func (reg *AvailableRegistry) isRegistered(module Module) bool {
	for _, mod := range reg.modules {
		if mod.ConfigurationKey() == module.ConfigurationKey() {
			return true
		}
	}
	return false
}
