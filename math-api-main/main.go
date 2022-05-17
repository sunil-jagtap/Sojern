package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MathRequest struct {
	Numbers []float64
	Q       int
}

type MathResponse struct {
	Description string
	Results     []float64
}

func parseParams(w http.ResponseWriter, r *http.Request, key string) (string, error) {
	params, ok := r.URL.Query()[key]
	if !ok {
		log.Printf("Query param %s is missing\n", key)
		return "", fmt.Errorf("query param %s is missing", key)
	}
	log.Printf("Query param '%s' has value %s\n", key, params)
	return params[0], nil
}

func parseNumbersFromQueryString(w http.ResponseWriter, r *http.Request) ([]float64, error) {
	numsAsString, err := parseParams(w, r, "numbers")
	if err != nil {
		return nil, err
	}
	var nums []float64
	for _, v := range strings.Split(numsAsString, ",") {
		// invalid floats are ignored
		if num, err := strconv.ParseFloat(v, 64); err == nil {
			nums = append(nums, num)
		}
	}
	return nums, nil
}

func parseQuantifierFromQueryString(w http.ResponseWriter, r *http.Request) (int, error) {
	q, err := parseParams(w, r, "q")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(q)
}

func parseRequest(w http.ResponseWriter, r *http.Request) MathRequest {
	nums, _ := parseNumbersFromQueryString(w, r)
	q, _ := parseQuantifierFromQueryString(w, r)
	return MathRequest{Numbers: nums, Q: q}
}

func FindMin(w http.ResponseWriter, r *http.Request, request MathRequest) []byte {
	c := Calculation{Numbers: request.Numbers}
	payload, _ := json.Marshal(MathResponse{fmt.Sprintf("Min of %v with quantifier %v", request.Numbers, request.Q), c.Min(request.Q)})
	return payload
}

func FindMax(w http.ResponseWriter, r *http.Request, request MathRequest) []byte {
	c := Calculation{Numbers: request.Numbers}
	payload, _ := json.Marshal(MathResponse{fmt.Sprintf("Max of %v with quantifier %v", request.Numbers, request.Q), c.Max(request.Q)})
	return payload
}

func FindAvg(w http.ResponseWriter, r *http.Request, request MathRequest) []byte {
	c := Calculation{Numbers: request.Numbers}
	payload, _ := json.Marshal(MathResponse{fmt.Sprintf("Average of %v", request.Numbers), []float64{c.Average()}})
	return payload
}

func FindMedian(w http.ResponseWriter, r *http.Request, request MathRequest) []byte {
	c := Calculation{Numbers: request.Numbers}
	payload, _ := json.Marshal(MathResponse{fmt.Sprintf("Median of %v", request.Numbers), []float64{c.Median()}})
	return payload
}

func FindPerentile(w http.ResponseWriter, r *http.Request, request MathRequest) []byte {
	c := Calculation{Numbers: request.Numbers}
	payload, _ := json.Marshal(MathResponse{fmt.Sprintf("%vth percentile of %v", request.Q, request.Numbers), []float64{c.Percentile(request.Q)}})
	return payload
}

func main() {
	http.HandleFunc("/min", Handler(FindMin))
	http.HandleFunc("/max", Handler(FindMax))
	http.HandleFunc("/avg", Handler(FindAvg))
	http.HandleFunc("/median", Handler(FindMedian))
	http.HandleFunc("/percentile", Handler(FindPerentile))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Wraps common logic between endpoints around the endpoint handler
func Handler(fn func(http.ResponseWriter, *http.Request, MathRequest) []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := parseRequest(w, r)
		payload := fn(w, r, request)
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
