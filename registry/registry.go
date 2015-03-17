package registry

import (
	"time"

	"github.com/alveary/overseer/service"
	"github.com/alveary/overseer/watchdog"
)

// Registry of microservices
type Registry struct {
	Services map[string]interface{}
}

// Register registrates a new service for a service name
func (registry *Registry) Register(newService *service.Service) {
	registry.Services[newService.Name] = newService
}

func (registry *Registry) UnleashWatchdogs() {
	go func() {
		for true {
			for _, registered := range registry.Services {
				watchdog.Watch(registered.(*service.Service))
			}

			time.Sleep(30 * time.Second)
		}
	}()
}
