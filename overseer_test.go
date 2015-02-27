package main

import (
	"testing"

	"github.com/go-martini/martini"
	"github.com/stretchr/testify/suite"
)

//
type OverseerTestSuite struct {
	suite.Suite
	app *martini.ClassicMartini
}

//
func (suite *OverseerTestSuite) SetupTest() {
	suite.app = AppEngine()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOverseermTestSuite(t *testing.T) {
	suite.Run(t, new(OverseerTestSuite))
}
