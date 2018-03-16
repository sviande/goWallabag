package goWallabag

import (
	"net/http"
	"net/http/httptest"
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

func TestAuthQueryHttpFailed(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
	}))

	wallabag := NewWallabag(ts.URL+"/", ts.Client())
	req := AuthRequest{}
	got := AuthQuery(&wallabag.Client, req)
	want := "Failed to parse error response: EOF"

	if got.Error() != want {
		t.Errorf("Error auth got: %v want %v", got, want)
	}

	ts.Close()
	got = AuthQuery(&wallabag.Client, req)
	if got == nil {
		t.Error("Error auth must return an error")
	}
}

func TestAuthQueryJsonFailed(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("error wanted for json"))
	}))

	defer ts.Close()

	wallabag := NewWallabag(ts.URL+"/", ts.Client())
	req := AuthRequest{}
	got := AuthQuery(&wallabag.Client, req)
	want := "Auth Failed from json decode: invalid character 'e' looking for beginning of value"

	if got.Error() != want {
		t.Errorf("Error auth got: %v want %v", got, want)
	}
}

func TestAuthQuerySuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("{}"))
	}))

	defer ts.Close()

	wallabag := NewWallabag(ts.URL+"/", ts.Client())
	req := AuthRequest{}
	got := AuthQuery(&wallabag.Client, req)

	if got != nil {
		t.Errorf("Error auth got non nil value.")
	}
}
