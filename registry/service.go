package registry

import "time"

type service map[string]string

func newService(name string, address string) service {
	return service{
		"name":       name,
		"address":    address,
		"fails":      "0",
		"created_at": time.Now().String(),
		"updated_at": time.Now().String(),
	}
}
