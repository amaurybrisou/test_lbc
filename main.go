package main

import (
	"amaurybrisou/test_lbc/fizz_buzz"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type fizzBuzzRequest struct {
	int1, int2, limit int64
	str1, str2        string
}

func (f fizzBuzzRequest) String() string {
	return fmt.Sprintf("int1: %d, int2: %d, limit: %d, str1: %s, str2: %s", f.int1, f.int2, f.limit, f.str1, f.str2)
}

var fizzBuzzRequests map[fizzBuzzRequest]int

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	requestParams := r.URL.Query()
	int1, int2, limit, str1, str2 := int64(3), int64(5), int64(100), "Fizz", "Buzz"

	if len(requestParams) > 0 {
		if len(requestParams) != 5 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "int1, int2, limit, str1, str2 parameters are required")
			return
		}

		var err error

		int1, err = strconv.ParseInt(requestParams.Get("int1"), 10, 0)
		if err != nil || int1 < 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "int1 should be a positive integer")
			return
		}

		int2, err = strconv.ParseInt(requestParams.Get("int2"), 10, 0)
		if err != nil || int2 < 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "int2 should be a positive integer")
			return
		}

		limit, err = strconv.ParseInt(requestParams.Get("limit"), 10, 0)
		if err != nil || limit < 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "limit should be a positive integer")
			return
		}

		str1 = requestParams.Get("str1")
		if str1 == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "str1 should be a string")
			return
		}

		str2 = requestParams.Get("str2")
		if str2 == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "str2 should be a string")
			return
		}

	}

	fizzBuzzRequests[fizzBuzzRequest{int1, int2, limit, str1, str2}]++
	fmt.Fprint(w, fizz_buzz.FizzBuzz(int(int1), int(int2), int(limit), str1, str2))
}

func frequenceHandler(w http.ResponseWriter, r *http.Request) {
	if len(fizzBuzzRequests) == 0 {
		fmt.Fprint(w, "No Requests so far")
		return
	}

	var (
		max = 0
		req fizzBuzzRequest
	)

	for k, m := range fizzBuzzRequests {
		if m > max {
			max = m
			req = k
		}
	}

	plural := ""
	if max > 1 {
		plural = "s"
	}

	fmt.Fprintf(w, "%v has been called : %d time%s", req, max, plural)
}

func main() {
	fizzBuzzRequests = make(map[fizzBuzzRequest]int)
	r := http.NewServeMux()

	setup(r)

	srv := http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func setup(r *http.ServeMux) {
	fizzBuzzRequests = make(map[fizzBuzzRequest]int)
	r.HandleFunc("/frequency", frequenceHandler)
	r.HandleFunc("/", fizzBuzzHandler)
}
