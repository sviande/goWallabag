package goWallabag

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetUrlValues(t *testing.T) {

	expected := url.Values{}
	expected.Set("grant_type", "GrantType")
	expected.Set("client_id", "ClientID")
	expected.Set("client_secret", "ClientSecret")
	expected.Set("username", "Username")
	expected.Set("password", "Password")

	authR := AuthRequest{}
	authR.GrantType = "GrantType"
	authR.ClientID = "ClientID"
	authR.ClientSecret = "ClientSecret"
	authR.Username = "Username"
	authR.Password = "Password"

	got := authR.GetUrlValues()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Failed for getUrlValues expected %v got %v", expected, got)
	}
}

func TestGetHeader(t *testing.T) {
	authResponse := AuthResponse{
		TokenType:   "bearer",
		AccessToken: "123456",
	}

	expected := "Bearer 123456"
	got := authResponse.GetHeader()
	if expected != got {
		t.Errorf("Failed for auth GetHeader expected %v got %v", expected, got)
	}
}
