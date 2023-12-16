package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"car-service/model"

	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(3 * time.Second)

	tests := []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"create customer", http.MethodPost, "customer", http.StatusCreated, toJSONString(model.Customer{
			ID:          1,
			Name:        "Chetan Kushwah",
			Email:       "chetankushwah929@gmail.com",
			Phone:       "1234567890",
			Address:     "123 Main St",
			City:        "Anytown",
			DateOfBirth: "1990-01-01",
			IsActive:    true,
		})},
		{"get all customers", http.MethodGet, "customer", http.StatusOK, nil},
		{"get customer by ID", http.MethodGet, "customer/1", http.StatusOK, nil},
		{"update customer", http.MethodPut, "customer", http.StatusOK, toJSONString(model.Customer{
			ID:          1,
			Name:        "Chetan Kushwah",
			Email:       "chetankushwah929@gmail.com",
			Phone:       "9876543210",
			Address:     "456 Oak St",
			City:        "Newtown",
			DateOfBirth: "1990-01-01",
			IsActive:    true,
		})},
		{"delete customer", http.MethodDelete, "customer/1", http.StatusNoContent, nil},
	}

	for i, tc := range tests {
		req, _ := request.NewMock(tc.method, "http://localhost:3000/"+tc.endpoint, bytes.NewBuffer(tc.body))

		c := http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()
	}
}

// toJSONString converts a struct to its JSON representation.
func toJSONString(data interface{}) []byte {
	result, _ := json.Marshal(data)
	return result
}
