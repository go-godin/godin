package godin

type Configurator interface {
	FileExists() bool
	CreateFile() error
	MustSave()
}

const ConfigFileName = "godin"
const ConfigFileType = "yaml"

type godinConfiguration struct {
}