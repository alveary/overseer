package watchdog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alveary/overseer/service"
)

// Watch ...
func Watch(service *service.Service) {

	go func() {
		available := true

		checkchan := make(chan bool)
		errorchan := make(chan error)
		defer func() {
			close(checkchan)
			close(errorchan)
		}()

		for available {

			go func() {
				resp, err := http.Head(service.Alive)
				if err != nil {
					errorchan <- err
					return
				}
				if resp.StatusCode > 299 {
					fmt.Println(fmt.Sprintf("Watchdog Lookup: (%s) Status: %s", service.Alive, resp.StatusCode))
					errorchan <- fmt.Errorf("Request Error: %s", resp.Status)
					return
				}

				checkchan <- true
			}()

			select {
			case <-checkchan:
				fmt.Printf("Service \"%s\" is still alive\n", service.Name)
			case <-errorchan:
				service.AddFailure()
				fmt.Printf("Service \"%s\" is not available, increasing fail count to %d\n", service.Name, service.Fails)
			case <-time.After(time.Second * 3):
				service.AddFailure()
				fmt.Printf("Service \"%s\" is not responding in time, increasing fail count to %d\n", service.Name, service.Fails)
			}

			time.Sleep(30 * time.Second)
		}

	}()

}
