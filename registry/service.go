package registry

// Service ...
type Service struct {
	Name  string `json:"name"`
	Root  string `json:"root"`
	Alive string `json:"alive"`
	fails int
}

func (service *Service) AddFailure() {
	service.fails = service.fails + 1
}
