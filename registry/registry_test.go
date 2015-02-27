package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegistryTestSuite struct {
	suite.Suite
}

// func (suite *RegistryTestSuite) SetupTest() {
// }
func (suite *RegistryTestSuite) TestServiceRegistration() {
	registry := Registry{make(map[string][]Service)}
	newService := Service{Name: "fubar", URL: "test"}
	assert.Equal(suite.T(), len(registry.Services["fubar"]), 0, "A new registry should be empty")

	registry.Register(newService)

	assert.Equal(suite.T(), len(registry.Services["fubar"]), 1, "The new service should be registered")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRegistrymTestSuite(t *testing.T) {
	suite.Run(t, new(RegistryTestSuite))
}