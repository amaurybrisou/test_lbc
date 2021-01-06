package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFrequencyHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := http.NewServeMux()
	setup(router)

	router.ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/frequency", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `int1: 3, int2: 5, limit: 100, str1: Fizz, str2: Buzz has been called : 1 time`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	req, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/frequency", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected = `int1: 3, int2: 5, limit: 100, str1: Fizz, str2: Buzz has been called : 2 times`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	req, err = http.NewRequest("GET", "/?int1=3&int2=5&limit=15&str1=Fizz&str2=Buzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/frequency", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected = `int1: 3, int2: 5, limit: 100, str1: Fizz, str2: Buzz has been called : 2 times`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	for i := 0; i < 15; i++ { //call another 15 times this handler
		req, err = http.NewRequest("GET", "/?int1=3&int2=5&limit=15&str1=Fizz&str2=Buzz", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr = httptest.NewRecorder()
		router.ServeHTTP(rr, req)
	}

	req, err = http.NewRequest("GET", "/frequency", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected = `int1: 3, int2: 5, limit: 15, str1: Fizz, str2: Buzz has been called : 16 times`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
