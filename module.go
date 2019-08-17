package godin

import (
	"crypto/rand"
	"encoding/hex"
)

// Module defines the interface of default godin enabledModules
type Module interface {
	Configurable
	//Addable
	//Removable

	New() Module

	// ID returns a unique ID for the current module instance. The purpose of the ID is to allow
	// enabledModules being installed multiple times and still be identified by their ID.
	// If a module hardcoded that identifier, it cannot be installed multiple times.
	ID() string

	// Configuration returns the module configuration as interface.
	// The configuration is returned as interface not only to allow for custom configuration structs.
	// The configuration is owned by the module and is of no concern by the other layers of godin.
	Configuration() interface{}

	// Install hook is called when 'godin add' is executed for that module. The hook enables the module to interfere
	// and prepare the module (e.g. prompt for values).
	// The installation is considered a success if error == nil.
	Install() error

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

// BaseModule defines the default behaviour of godin enabledModules.
type BaseModule struct {
	ModuleName string
	Identifier string
}

// ID returns a unique ID for the current module instance. The purpose of the ID is to allow
// enabledModules being installed multiple times and still be identified by their ID.
// If a module hardcoded that identifier, it cannot be installed multiple times.
func (mod *BaseModule) ID() string {
	if mod.Identifier == "" {
		b := make([]byte, 4) // 8 characters
		_, _ = rand.Read(b)
		mod.Identifier = hex.EncodeToString(b)
	}
	return mod.Identifier
}

// Name returns a printable name of the module.
func (mod *BaseModule) Name() string {
	return mod.ModuleName
}
