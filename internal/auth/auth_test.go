package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func test_GetAPIKey(t *testing.T) {
	// test cases
	headers := make(http.Header)
	headers.Set("Authorization", "")
	got, _ := GetAPIKey(headers)
	want := ""
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("test_GetAPIKey failed")
	}
}
