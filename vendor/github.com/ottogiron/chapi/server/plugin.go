package server

import "github.com/satori/go.uuid"

type Plugin interface {
	Register(server Server)
	Dependencies() []string
	Name() string
}

type BasePlugin struct {
}

func (basePlugin *BasePlugin) Dependencies() []string {
	return []string{}
}

func (basePlugin *BasePlugin) Name() string {
	return uuid.NewV4().String()
}
