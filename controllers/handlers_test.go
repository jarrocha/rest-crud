package controllers_test

import (
	"net/http/httptest"
	"testing"

	"gitlab.com/codelittinc/golang-interview-project-jaime/rest-crud/controllers"
)

var handler_test_string string = "%s(%s) -> got %v want %v\n"

func TestIndexHandler(t *testing.T) {
	tests := []struct {
		input1 string
		status int
	}{
		{"GET", 200},
		{"POST", 405},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.input1, "https://localhost:8081/", nil)
		w := httptest.NewRecorder()

		controllers.Index(w, req)
		resp := w.Result()
		got1 := resp.StatusCode
		if got1 != test.status {
			t.Errorf(handler_test_string, "TestIndexHandler", test.input1, got1, test.status)
		}
	}
}

func TestShowHandler(t *testing.T) {
	tests := []struct {
		input1 string
		status int
	}{
		{"GET", 500},
		{"POST", 405},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.input1, "https://localhost:8081/users/show", nil)
		w := httptest.NewRecorder()

		controllers.Show(w, req)
		resp := w.Result()
		got1 := resp.StatusCode
		if got1 != test.status {
			t.Errorf(handler_test_string, "TestShowHandler", test.input1, got1, test.status)
		}
	}
}

func TestCreateProcessHandler(t *testing.T) {
	tests := []struct {
		input1 string
		status int
	}{
		{"GET", 405},
		{"POST", 406},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.input1, "https://localhost:8081/process", nil)
		w := httptest.NewRecorder()

		controllers.CreateProcess(w, req)
		resp := w.Result()
		got1 := resp.StatusCode
		if got1 != test.status {
			t.Errorf(handler_test_string, "TestCreateProcessHandler",
				test.input1, got1, test.status)
		}
	}
}

func TestUpdateHandler(t *testing.T) {
	tests := []struct {
		input1 string
		status int
	}{
		{"GET", 500},
		{"POST", 405},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.input1, "https://localhost:8081/process", nil)
		w := httptest.NewRecorder()

		controllers.Update(w, req)
		resp := w.Result()
		got1 := resp.StatusCode
		if got1 != test.status {
			t.Errorf(handler_test_string, "TestUpdateHandler",
				test.input1, got1, test.status)
		}
	}
}
