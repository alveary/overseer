package watchdog

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alveary/overseer/service"
	"github.com/go-martini/martini"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//
type TestSuite struct {
	suite.Suite
	app *martini.ClassicMartini
}

//
func (suite *TestSuite) SetupTest() {

}

func responseHeaderServer(code int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
	}))
}

func (suite *TestSuite) TestWatchdogAliveLookupForAvailableServices() {
	server := responseHeaderServer(200)
	defer server.Close()

	service := service.Service{"service-name", server.URL, server.URL, 0}

	watchdog := NewWatchdog(&service)
	assert.Equal(suite.T(), watchdog.requestAliveResource(), nil, "a successful lookup should result in no error")
}

func (suite *TestSuite) TestWatchdogAliveLookupForUnavailableServices() {
	server := responseHeaderServer(400)
	defer server.Close()

	// mock service
	service := service.Service{"service-name", server.URL, server.URL, 0}

	watchdog := NewWatchdog(&service)

	assert.IsType(suite.T(), watchdog.requestAliveResource(), errors.New("Error String"), "a successful lookup should result in no error")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
