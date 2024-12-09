package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HemloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hemlo, world!")
}
func Ex4() {
	http.HandleFunc("/hello", HemloHandler)

	http.ListenAndServe(":3030", nil)
}

func TestUnitHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	// Type conversion, this converts a handler function to a handler type.
	handler := http.HandlerFunc(HemloHandler)
	// Activating the handler type we created from our function.
	handler.ServeHTTP(recorder, req)

	expectedResBody := "Hemlo, world!"
	if recorder.Body.String() != expectedResBody {
		t.Errorf("handler returned unexpected body: "+
			"got %v and wanted %v\n", recorder.Body.String(),
			expectedResBody)
	}
}

func TestIntegrationHemloHandler(t *testing.T) {
	// Test server
	ts := httptest.NewServer(http.HandlerFunc(HemloHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hemlo, world!"
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}
	actual := *readBody(t, res)
	if actual != expected {
		// %q escapes the characters of the strings, like raw string.
		t.Errorf("expected body %q, got %q", expected, actual)
	}
}

func readBody(t *testing.T, r *http.Response) *string {
	// I think that the builder optimises memory consumption while writing to it.
	// It's meant for io.Copy.
	buffer := new(strings.Builder)
	_, err := io.Copy(buffer, r.Body)
	if err != nil {
		t.Fatal(err)
	}
	returnVal := buffer.String()

	return &returnVal
}
