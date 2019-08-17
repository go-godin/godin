package server

import (
	"gitub.com/go-godin/godin"
)

type serverTemplate struct {
	godin.BaseTemplate
}

func NewServerTemplate() *serverTemplate {
	return &serverTemplate{
		BaseTemplate: godin.BaseTemplate{
			Config: &godin.TemplateConfiguration{
				Name:       "grpc-server",
				SourceFile: "transport/grpc/server.go.tmpl",
				TargetFile: "transport/grpc/server.go",
				GoSource:   true,
				Skip:       false,
			},
		},
	}
}
