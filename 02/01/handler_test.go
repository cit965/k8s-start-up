package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com?param1=5&param2=3", nil)
	w := httptest.NewRecorder()

	Handle(w, req)

	// We should get a good status code
	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	// Make sure that the version was 1.3
	if want, got := "1.3", w.Result().Header.Get("API-VERSION"); want != got {
		t.Fatalf("expected API-VERSION to be %s, instead got: %s", want, got)
	}

	// Test the result
	var response Response
	_ = json.NewDecoder(w.Body).Decode(&response)

	if want, got := 8, response.Sum; want != got {
		t.Fatalf("expected sum to be %d, instead got: %d", want, got)
	}
}
