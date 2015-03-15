package registry

import (
	"github.com/alveary/overseer/service"
	"github.com/alveary/overseer/watchdog"
)

// Registry of microservices
type Registry struct {
	Services map[string]interface{}
}

type registeredService struct {
	*service.Service
	watchdog watchdog.Watchdog
}

// Register registrates a new service for a service name
func (registry *Registry) Register(newService *service.Service) {
	regService := registeredService{
		newService,
		watchdog.Watch(newService),
	}

	if registry.Services[newService.Name] != nil {
		registry.Services[newService.Name].(registeredService).watchdog.Deregister()
	}

	registry.Services[newService.Name] = regService
}
