package registry

// Service ...
type Service struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	fails int
}

// Registry of microservices
type Registry struct {
	Services map[string][]Service
}

// Register registrates a new service for a service name
func (registry *Registry) Register(service Service) {
	registry.Services[service.Name] = append(registry.Services[service.Name], service)
}
