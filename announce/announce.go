package announce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// ServiceEntry ...
type ServiceEntry struct {
	ServiceName   string
	ServiceRoot   string
	AliveResource string
}

// Service provides a method to attach a new Service to the overseer stack
func Service(entry ServiceEntry) {
	json, _ := json.Marshal(entry)

	overseerRoot := os.Getenv("OVERSEER_ROOT")

	if overseerRoot == "" {
		fmt.Println("OVERSEER_ROOT is not set:")
	}

	http.Post(overseerRoot, "application/json", bytes.NewBuffer(json))
}
