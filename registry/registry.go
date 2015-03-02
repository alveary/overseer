package registry

import "github.com/alveary/overseer/service"

// Registry of microservices
type Registry struct {
	Services map[string]interface{}
}

// Register registrates a new service for a service name
func (registry *Registry) Register(service *service.Service) {
	registry.Services[service.Name] = service
}
