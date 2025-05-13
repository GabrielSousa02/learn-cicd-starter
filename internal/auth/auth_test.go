package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	httpHeader := http.Header{ }
	// TODO: Fail test to check if CI works
	// httpHeader.Add("Authorization", "ApiKey my-fake-api-key")
	got, gotAPIError := GetAPIKey(httpHeader)
	want  := "my-fake-api-key"

	if gotAPIError != nil {
		t.Fatalf("expected: no errors, got %v", gotAPIError)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got %v", want, got)
	}

}
