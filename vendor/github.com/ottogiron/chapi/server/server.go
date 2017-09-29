package server

import (
	"errors"
	"fmt"
	"net/http"
)

type Server interface {
	Run(addr string) error
	Register(plugin Plugin) error
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *route
}

type baseServer struct {
	plugins map[string]Plugin
}

func (baseServer *baseServer) Register(plugin Plugin) error {

	_, containsKey := baseServer.plugins[plugin.Name()]

	if containsKey {
		alreadyRegisteredMsg := "Plugin already registered:%s"
		return errors.New(fmt.Sprint(alreadyRegisteredMsg, plugin.Name()))
	}

	baseServer.plugins[plugin.Name()] = plugin

	return nil
}

func (baseServer *baseServer) registerPlugins(plugins map[string]Plugin, server Server, processedPlugins map[string]bool) error {

	for _, plugin := range plugins {
		dependenciesNames := plugin.Dependencies()
		depLen := len(dependenciesNames)
		if !processedPlugins[plugin.Name()] {
			if depLen == 0 {
				plugin.Register(server)
			} else {
				meet := pluginDependencyMeet(plugin, processedPlugins)
				if meet {
					plugin.Register(server)
				} else {

					dependencies := make(map[string]Plugin)
					for _, currentDependency := range dependenciesNames {

						if plugins[currentDependency] != nil {
							if currentDependency != plugin.Name() {
								dependencies[currentDependency] = plugins[currentDependency]
							} else {
								unmetDependenciesMsg := "Circular dependency %s"
								return fmt.Errorf(unmetDependenciesMsg, plugin.Name())
							}
						} else {
							unmetDependenciesMsg := "Dependencies for %s is unmet %s"
							return fmt.Errorf(unmetDependenciesMsg, plugin.Name(), currentDependency)
						}

						baseServer.registerPlugins(dependencies, server, processedPlugins)
						plugin.Register(server)
					}
				}
			}
			processedPlugins[plugin.Name()] = true
		}

	}
	return nil
}

func pluginDependencyMeet(plugin Plugin, processedPlugins map[string]bool) bool {

	for _, name := range plugin.Dependencies() {
		if !processedPlugins[name] {
			return false
		}
	}
	return true
}
