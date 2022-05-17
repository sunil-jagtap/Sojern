package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

const GREATER_THAN = 1
const LESS_THAN = -1
const EQUAL = 0

func CompareHadler(w http.ResponseWriter, r *http.Request) {
	subject := r.URL.Query().Get("s")
	comparator := r.URL.Query().Get("c")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Compare(subject, comparator))
}

// Compare 2 versions, the subject and the comparator or version to compare against.
// Returns 1 if subject > comparator, 0 if equal and -1 if less than.
// Reasoning for decisions made:
// - Using regex "[0-9]+" instead of "\\." - dispite the assumption that the version be
//   only made up of digits and dots (.) I felt this would avoid issues converting the
//   version parts to int
// - getOrInfer - this is to handle Compare("1.1", "1.1.0") where the subject will be
//   treated like "1.1.0"
func Compare(subject string, comparator string) (result int) {
	pattern, _ := regexp.Compile("[0-9]+")
	subject_parts := pattern.FindAllString(subject, -1)
	comparator_parts := pattern.FindAllString(comparator, -1)
	result = recursively_compare(subject_parts, comparator_parts, 0)
	return
}

func recursively_compare(subject_parts, comparator_parts []string, index int) int {
	subjectPart, _ := strconv.Atoi(getOrInfer(subject_parts, index))
	comparatorPart, _ := strconv.Atoi(getOrInfer(comparator_parts, index))
	if subjectPart > comparatorPart {
		return GREATER_THAN
	}
	if subjectPart < comparatorPart {
		return LESS_THAN
	}
	if index >= len(subject_parts) && index >= len(comparator_parts) {
		return EQUAL
	}
	return recursively_compare(subject_parts, comparator_parts, index+1)
}

func getOrInfer(a []string, index int) string {
	if index >= len(a) {
		return "0"
	}
	return a[index]
}

func main() {
	http.HandleFunc("/compare", CompareHadler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
