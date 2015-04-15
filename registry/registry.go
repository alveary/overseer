package registry

import (
	"log"
	"os"

	"github.com/fzzy/radix/redis"
)

// Registry of services
type Registry struct {
	client *redis.Client
}

// New is the Registry Constructor
func New() (*Registry, error) {
	host := os.Getenv("OVERSEER_REDIS_HOST")
	port := os.Getenv("OVERSEER_REDIS_PORT")

	client, err := redis.Dial("tcp", host+":"+port)

	if err != nil {
		log.Printf("RedisConnectionFailed: %s", err)
		return &Registry{client: nil}, nil
	}

	return &Registry{client: client}, nil
}

// All retrieves all registered entries (keys)
func (registry *Registry) All() []string {
	if registry.client == nil {
		log.Printf("Service Not Available: %s", "redis")
		return []string{}
	}

	entries, err := registry.client.Cmd("KEYS", "*").List()

	if err != nil {
		log.Printf("RedisCMDFailed: (%s) - %s", "KEYS", err)
	}

	return entries
}
