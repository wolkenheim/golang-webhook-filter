package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gofiber/fiber/v2"
	"dam-webhook/webhook"
)

type SingleTestCase struct {
	description string
	route string
	method string
	expectedError bool
	expectedCode  int
	checkBody bool
	expectedBody  string
	jsonValue string
}

func TestIndexRoute(t *testing.T) {
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []SingleTestCase{
		{
			description:   "hello world test",
			route:         "/",
			method: "GET",
			expectedError: false,
			expectedCode:  200,
			checkBody: true,
			expectedBody:  "hello world",
			jsonValue: ``,
		},
		{
			description:   "webhook post",
			route:         "/webhook",
			method: "POST",
			expectedError: false,
			expectedCode:  200,
			checkBody: false,
			expectedBody:  "",
			jsonValue: `{"assetId": "87c23cwqDD2111", "metadata": {"folderPath": "/Client/XXX Group Holding SE & Co. KGaA/my-image-name-jpg", "cf_approvalState_client1": "Approved", "cf_assetType": {"value": "Content Image"}}}`,
		},
	}

	// Setup the app as it is done in the main function
	app := Setup()

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case

		payload := bytes.NewBuffer([]byte(test.jsonValue))

		req, _ := http.NewRequest(
			test.method,
			test.route,
			payload,
		)

		req.Header.Set("Content-Type", "application/json")

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

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		if(test.checkBody){
			assert.Equalf(t, test.expectedBody, string(body), test.description)
		}

	}
}

func Setup() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
	app.Post("/webhook", webhook.CreateWebhook)
	app.Post("/mock-api", webhook.MockApi)


	readConfig("testing")

	return app;
}
