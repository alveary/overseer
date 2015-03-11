package registry

import (
	"testing"

	"github.com/alveary/overseer/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegistryTestSuite struct {
	suite.Suite
}

func (suite *RegistryTestSuite) SetupTest() {
}

func (suite *RegistryTestSuite) TestServiceRegistration() {
	registry := Registry{make(map[string]interface{})}
	newService := service.Service{"service-name", "root-url", "alive-url", 0}
	assert.Equal(suite.T(), registry.Services["service-name"], nil, "A new registry should be empty")

	registry.Register(&newService)

	assert.Equal(suite.T(), registry.Services["service-name"], &newService)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRegistryTestSuite(t *testing.T) {
	suite.Run(t, new(RegistryTestSuite))
}
