package registry

// Service ...
type Service struct {
	Name  string `json:"name"`
	Root  string `json:"url"`
	Alive string `json:"url"`
	fails int
}

func (service *Service) AddFailure() {
	service.fails = service.fails + 1
}
