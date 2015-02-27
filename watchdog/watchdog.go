package watchdog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alveary/overseer/registry"
)

// Watch ...
func Watch(service *registry.Service) {

	go func() {
		available := true

		for available {

			checkchan := make(chan bool)
			errorchan := make(chan error)
			defer func() {
				close(checkchan)
				close(errorchan)
			}()

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
				// JUST GO ON
			case <-errorchan:
				service.AddFailure()
			case <-time.After(time.Second * 3):
				service.AddFailure()
			}

			time.Sleep(1 * time.Minute)

		}

	}()

}
