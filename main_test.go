package main

import (
	"fmt"
	"testing"
	"net/http"
    "github.com/Antuan01/go-test/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/Antuan01/go-test/database"
	"encoding/json"
	"strings"
	"io/ioutil"
)

func TestApp(t *testing.T) {
	app := fiber.New()

    app.Use(cors.New())

    app.Use(logger.New())

    database.InitDB()

    routes.SetupRoutes(app)

	testValidation := []struct {
		testName string
		url string
		payload string
		expectedCode int
		expectedResponse string
		testCase int
		expectedReason string
	}{
		{
			"Reason must have more than 10 characters",
			"http://127.0.0.1:3000/api/report/post",
			`{"reason": "aoeu", "post_id": 1}`+"\n",
			422,
			`{"error":"reason must be at least 10 characters long"}`,
			0,
			"",
		},
		{
			"Only existing post shuld have reports",
			"http://127.0.0.1:3000/api/report/post",
			`{"reason": "long reason to test", "post_id": 345}`+"\n",
			422,
			`{"error":"post not found"}`,
			0,
			"",
		},
		{
			"Created record must have a reason text",
			"http://127.0.0.1:3000/api/report/post",
			`{"reason": "long reason to test", "post_id": 1}`+"\n",
			201,
			"",
			1,
			"long reason to test",
		},
		{
			"Reason must have more than 10 characters",
			"http://127.0.0.1:3000/api/report/comment",
			`{"reason": "aoeu", "comment_id": 1}`+"\n",
			422,
			`{"error":"reason must be at least 10 characters long"}`,
			0,
			"",
		},
		{
			"Only existing comments shuld have reports",
			"http://127.0.0.1:3000/api/report/comment",
			`{"reason": "long reason to test", "comment_id": 345}`+"\n",
			422,
			`{"error":"comment not found"}`,
			0,
			"",
		},
		{
			"Created record must have a reason text",
			"http://127.0.0.1:3000/api/report/comment",
			`{"reason": "long reason to test", "comment_id": 1}`+"\n",
			201,
			"",
			1,
			"long reason to test",
		},
	}

	for _, test := range(testValidation) {
		req, _ := http.NewRequest("POST", test.url, strings.NewReader(test.payload))
	
		req.Header.Set("Content-type", "application/json")
	
		res, err := app.Test(req, -1)
		
		if err != nil {
			fmt.Println("Error")
			fmt.Println(err)	
		}
	
		fmt.Println(res.StatusCode)
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
		plainResponse := string(data)
		
		if res.StatusCode != test.expectedCode {
			t.Errorf("Bad response client error should have code %d not %d", test.expectedCode, res.StatusCode)
		}

		if test.testCase == 0 {		
			if plainResponse != test.expectedResponse {
				t.Errorf(test.testName)
			}
		}

		if test.testCase == 1 {

			type Response struct {
				Reason string `json:"reason"`
			}

			var response Response

			err := json.Unmarshal([]byte(plainResponse), &response)

			if err != nil {
				fmt.Println("JSON decode error!")
				return
			}
		
			fmt.Println(response.Reason)

			if response.Reason != test.expectedReason {
				t.Errorf(test.testName)
			}
		}
	}
}