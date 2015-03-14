package ask

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alveary/overseer/service"
)

// Ask lets you retrieve the service URI information
func Ask(serviceName string) (retrieved service.Service, err error) {
	overseerRoot := os.Getenv("OVERSEER_ROOT")
	responsechan := make(chan *http.Response)
	errorchan := make(chan error)

	go func() {
		resp, err := http.Get(overseerRoot + "/" + serviceName)
		if err != nil {
			errorchan <- err
			return
		}
		if resp.StatusCode > 299 {
			errorchan <- fmt.Errorf("Request Error: %s", resp.Status)
			return
		}

		responsechan <- resp
	}()

	select {
	case resp := <-responsechan:
		defer resp.Body.Close()

		dec := json.NewDecoder(resp.Body)
		dec.Decode(&retrieved)

		return retrieved, nil
	case err = <-errorchan:
		return retrieved, err
	case <-time.After(time.Second * 3):
		return retrieved, fmt.Errorf("Timeout of service registry call")
	}
}