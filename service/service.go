package service

// Service db schema
import "time"

type Service struct {
	ID        int    `sql:"AUTO_INCREMENT"`
	Name      string `sql:"size:255" json:"name"`
	Root      string `sql:"size:255" json:"root"`
	Alive     string `sql:"size:3000" json:"alive"`
	Fails     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (service *Service) AddFailure() {
	service.Fails = service.Fails + 1
}
