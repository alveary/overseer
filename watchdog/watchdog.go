package watchdog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alveary/overseer/service"
)

// Watchdog service checker
type Watchdog struct {
	service *service.Service
}

// NewWatchdog is the constructor function
func NewWatchdog(service *service.Service) Watchdog {
	return Watchdog{service: service}
}

func (watchdog Watchdog) requestAliveResource() error {
	resp, err := http.Head(watchdog.service.Alive)

	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("Watchdog Lookup Failed: (%s) Status: %s", watchdog.service.Alive, resp.StatusCode)
	}

	return nil
}

func (watchdog Watchdog) serviceTest(checkchan chan bool, errorchan chan error) {
	err := watchdog.requestAliveResource()

	if err != nil {
		errorchan <- err
		return
	}

	checkchan <- true
}

func (watchdog Watchdog) spawnWatchCycle(checkchan chan bool, errorchan chan error) {
	available := true

	defer func() {
		close(checkchan)
		close(errorchan)
	}()

	for available {

		go watchdog.serviceTest(checkchan, errorchan)

		select {
		case <-checkchan:
			fmt.Printf("Service \"%s\" is still alive\n", watchdog.service.Name)
		case <-errorchan:
			watchdog.service.AddFailure()
			fmt.Printf("Service \"%s\" is not available, increasing fail count to %d\n", watchdog.service.Name, watchdog.service.Fails)
		case <-time.After(time.Second * 3):
			watchdog.service.AddFailure()
			fmt.Printf("Service \"%s\" is not responding in time, increasing fail count to %d\n", watchdog.service.Name, watchdog.service.Fails)
		}

		time.Sleep(30 * time.Second)
	}
}

// Watch ...
func Watch(service *service.Service) {
	watchdog := NewWatchdog(service)
	go watchdog.spawnWatchCycle(make(chan bool), make(chan error))
}
