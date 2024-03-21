package controller_test

import (
	"io"
	"net/http"
	"testing"
	"todo/internal/initializer"

	"github.com/stretchr/testify/assert"
)

type RouteTest struct {
	description string // description of the test case
	route       string // route path to test

	// Expected output
	expectedError bool
	expectedCode  int
	expectedBody  string
}

// TODO: For now this file just tests whether the api is working at all, to setup idiomatic tests I will need to build out a test database, test fixtures and test environment variables.
func TestTodoRoutes(t *testing.T) {
	var tests = []RouteTest{
		// First test case
		{
			description:   "get HTTP status 200",
			route:         "/hello",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
		// Second test case
		{
			description:   "get HTTP status 404, when route doesnt exist",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	app := initializer.SetupApi()

	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
	}

}

// TODO: this currently isnt used, to use this properly, we would need to setup a test databse with fixtures, and a test environment.
func ReadBody(t *testing.T, res *http.Response, test RouteTest) {
	// Read the response body
	body, err := io.ReadAll(res.Body)

	// Reading the response body should work everytime, such that
	// the err variable should be nil
	assert.Nilf(t, err, test.description)

	// Verify, that the reponse body equals the expected body
	assert.Equalf(t, test.expectedBody, string(body), test.description)
}
