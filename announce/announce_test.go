package announce

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AnnounceTestSuite struct {
	suite.Suite
}

// func (suite *RegistryTestSuite) SetupTest() {
// }
func (suite *AnnounceTestSuite) TestAnnouncement() {

	// assert.Equal(suite.T(), registry.Services["service-name"], &newService, "The new service should be registered")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRegistrymTestSuite(t *testing.T) {
	suite.Run(t, new(AnnounceTestSuite))
}
