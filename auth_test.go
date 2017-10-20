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

	got := authR.GetURLValues()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Failed for getUrlValues expected %v got %v", expected, got)
	}
}

func TestGetHeader(t *testing.T) {
	var err error

	authResponse := AuthResponse{}

	_, err = authResponse.GetHeader()
	if err != errTokenType {
		t.Error("Failed for auth GetHeader expected errTokenType")
		return
	}

	authResponse.TokenType = "bearer"
	_, err = authResponse.GetHeader()
	if err != errAccesToken {
		t.Error("Failed for auth GetHeader expected errAccessToken")
		return
	}

	authResponse.AccessToken = "123456"

	expected := "Bearer 123456"
	got, _ := authResponse.GetHeader()
	if expected != got {
		t.Errorf("Failed for auth GetHeader expected %v got %v", expected, got)
	}
}
