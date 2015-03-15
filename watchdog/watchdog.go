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
	done    chan bool
}

// NewWatchdog is the constructor function
func NewWatchdog(service *service.Service, done chan bool) Watchdog {
	return Watchdog{service: service, done: done}
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

func (watchdog Watchdog) serviceTest(checkchan chan bool, errorchan chan error, donechan chan bool) {
	err := watchdog.requestAliveResource()

	select {
	case <-donechan:
		return
	default:
		if err != nil {
			errorchan <- err
			return
		}

		checkchan <- true
	}
}

func (watchdog Watchdog) Deregister() {
	defer func() {
		close(watchdog.done)
	}()

	fmt.Printf("Deregistering service \"%s\"\n", watchdog.service.Name)

	watchdog.done <- true
}

func (watchdog Watchdog) spawnWatchCycle(done chan bool) {
	checkchan := make(chan bool)
	errorchan := make(chan error)
	donechan := make(chan bool)

	available := true

	defer func() {
		close(checkchan)
		close(errorchan)
		close(donechan)
	}()

	for available {
		go watchdog.serviceTest(checkchan, errorchan, donechan)

		select {
		case <-done:
			available = false
			donechan <- true
			return
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
func Watch(service *service.Service) (watchdog Watchdog) {
	done := make(chan bool)
	watchdog = NewWatchdog(service, done)
	go watchdog.spawnWatchCycle(done)
	return watchdog
}
