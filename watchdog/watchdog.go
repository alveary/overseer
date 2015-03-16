package watchdog

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/alveary/overseer/service"
)

// Watchdog service checker
type Watchdog struct {
	Service *service.Service
	Done    chan bool
	Check   chan bool
	Err     chan error
}

// NewWatchdog is the constructor function
func NewWatchdog(service *service.Service) Watchdog {
	done := make(chan bool)
	check := make(chan bool)
	err := make(chan error)

	return Watchdog{service, done, check, err}
}

func (watchdog Watchdog) requestAliveResource() error {
	resp, err := http.Head(watchdog.Service.Alive)

	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("Watchdog Lookup Failed: (%s) Status: %s", watchdog.Service.Alive, resp.StatusCode)
	}

	return nil
}

func (watchdog Watchdog) ServiceTest() {
	fmt.Println(watchdog.Service)
	err := watchdog.requestAliveResource()

	select {
	case <-watchdog.Done:
		return
	default:
		if err != nil {
			watchdog.Err <- err
			return
		}

		watchdog.Check <- true
	}
}

// Watch ...
func Watch(newService *service.Service) (*service.Service, error) {
	watchdog := NewWatchdog(newService)

	go watchdog.ServiceTest()

	select {
	case <-watchdog.Check:
		fmt.Printf("Service \"%s\" is still alive\n", watchdog.Service.Name)
		return watchdog.Service, nil
	case err := <-watchdog.Err:
		watchdog.Service.AddFailure()
		fmt.Printf("Service \"%s\" is not available, increasing fail count to %d\n", watchdog.Service.Name, watchdog.Service.Fails)
		return watchdog.Service, err
	case <-time.After(time.Second * 3):
		watchdog.Done <- true
		watchdog.Service.AddFailure()
		fmt.Printf("Service \"%s\" is not responding in time, increasing fail count to %d\n", watchdog.Service.Name, watchdog.Service.Fails)
		return watchdog.Service, errors.New("Timeout of registered Service")
	}
}
