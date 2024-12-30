package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyGoodAuthorizationHeader(t *testing.T) {

	headers := http.Header{}
	apiKey := "ApiKey"
	set_key := "Chappy"
	headers.Set("Authorization", apiKey+" "+set_key)
	api_key, err := GetAPIKey(headers)

	if err != nil {
		t.Errorf("%s", err)
	}

	if api_key != set_key {
		t.Errorf("Expected %v got %v", set_key, api_key)
	}
}

func TestGetAPIKeyBadAuthorizationHeader(t *testing.T) {
	headers := http.Header{}
	apiKey := "ApiKey"
	set_key := "Chappy"

	headers.Set("Authorization", apiKey+" ribals ApiKey ribald")

	api_key, err := GetAPIKey(headers)

	if err != nil {
		t.Errorf("%s", err)
	}

	if api_key == set_key {
		t.Errorf("Expected %v got %v", set_key, api_key)
	}
}
