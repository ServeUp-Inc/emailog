package server

import (
	"io/ioutil"
	"net/http"
  "bytes"
  "encoding/json"
	"testing"

  "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

  "github.com/ServeUp-Inc/emailog/models"
)

func checkResponse(
  t                  *testing.T,
  server             *fiber.App,
  req                *http.Request,
  description        string,
  expectedError      bool,
  expectedStatusCode int,
  expectedBody       string,
) {
  // The -1 disables request latency.
  res, err := server.Test(req, -1)

  assert.Equalf(
    t,
    expectedError, err != nil,
    "%s: invalid error", description)

  // As expected errors lead to broken responses,
  // the next test case needs to be processed
  if expectedError {
    return
  }

  assert.Equalf(
    t,
    expectedStatusCode, res.StatusCode,
    "%s: wrong status code found", description)

  body, err := ioutil.ReadAll(res.Body)

  assert.Nilf(
    t, err,
    "%s: error encountered while reading body", description)

  assert.Equalf(
    t,
    expectedBody, string(body),
    "%s: unexpected data in body", description)
}

func TestV1GetRoutes(t *testing.T) {
  getRouteTests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "get index route",
			route:         "v1/",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Not Found",
		},
		{
			description:   "get non existing route",
			route:         "/non-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Not Found",
		},
	}

	server := Create()

	for _, test := range getRouteTests {
		req, _ := http.NewRequest(
      http.MethodGet,
			test.route,
			nil,
		)

    checkResponse(
      t, server, req,
      test.description,
      test.expectedError,
      test.expectedCode,
      test.expectedBody,
    )
	}
}

func TestV1PutRoute(t *testing.T) {
  putRouteTests := []struct {
    description string

    // Test input
    route string
    lead  models.Lead

    // Expected output
    expectedError bool
    expectedCode  int
    expectedBody  string
  } {
    {
      description:   "add new entry",
      route:         "/v1/",
      lead:          models.Lead {
        Email: "johndoe@example.com",
        Msg:   "john is alive.",
      },
      expectedError: false,
      expectedCode:  200,
      expectedBody:  "OK",
    },
    {
      description:   "add duplicate entry",
      route:         "/v1/",
      lead:          models.Lead {
        Email: "johndoe@example.com",
        Msg:   "john is alive.",
      },
      expectedError: false,
      expectedCode:  200,
      expectedBody:  "OK",
    },
    {
      description:   "try adding invalid email",
      route:         "/v1/",
      lead:          models.Lead {
        Email: "johndoeexample.com",
        Msg:   "bad email",
      },
      expectedError: false,
      expectedCode:  400,
      expectedBody:  "Bad Request",
    },
  }

	server := Create()

	for _, test := range putRouteTests {
    json, err := json.Marshal(test.lead)
    if err != nil {
        panic(err)
    }
    
		req, _ := http.NewRequest(
      http.MethodPut,
			test.route,
			bytes.NewBuffer(json),
		)
    req.Header.Set("Content-Type", "application/json; charset=utf-8")

    checkResponse(
      t, server, req,
      test.description,
      test.expectedError,
      test.expectedCode,
      test.expectedBody,
    )
	}
}
