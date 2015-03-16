package registry

import (
	"fmt"
	"os"
	"time"

	"github.com/alveary/overseer/service"
	"github.com/alveary/overseer/watchdog"
	"github.com/jinzhu/gorm"

	// gorm dependency
	_ "github.com/lib/pq"
)

// Registry of microservices
type Registry struct {
	db gorm.DB
}

func (registry Registry) Services() (services []service.Service) {
	registry.db.Limit(100).Find(&services)
	return services
}

// Register registrates a new service for a service name
func (registry *Registry) Register(requestedService *service.Service) {
	registry.db.FirstOrCreate(&requestedService, service.Service{Name: requestedService.Name})
}

// UnleashWatchdogs sets up the check cycle for registered services
func (registry *Registry) UnleashWatchdogs() {
	go func() {
		for true {
			for _, registered := range registry.Services() {
				go func() {
					fmt.Println("NEW WATCH CYCLE")
					result, err := watchdog.Watch(&registered)
					if err != nil {
						registry.db.Save(result)
					}
				}()
			}

			time.Sleep(30 * time.Second)
		}
	}()
}

// NewRegistry is the constructor method
func NewRegistry() Registry {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open("postgres", fmt.Sprintf("user=%s port=%s dbname=%s host=%s password=%s", dbUser, dbPort, dbName, dbHost, dbPassword))
	if err != nil {
		fmt.Errorf("OMG, DATABASE FOO")
	}

	// Then you could invoke `*sql.DB`'s functions with it
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)
	db.DropTable(&service.Service{})
	db.AutoMigrate(&service.Service{})
	return Registry{db: db}
}
