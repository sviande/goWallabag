package goWallabag

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testWallabag(URL string) Wallabag {
	wallabag := NewWallabag(URL+"/", &http.Client{})
	wallabag.Client.auth = AuthResponse{
		AccessToken: "asd",
		TokenType:   "test",
	}

	return wallabag
}

func TestVersionFetchSuccess(t *testing.T) {
	want := "2.3.3"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		_ = encoder.Encode(want)
	}))
	defer ts.Close()

	wallabag := testWallabag(ts.URL)
	got, err := wallabag.GetVersion()

	if err != nil {
		t.Errorf("err %s", err)
	}

	if got != want {
		t.Errorf("Version failed got: %v, want: %v", got, want)
	}
}

func TestVersionFetchError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
	}))
	defer ts.Close()

	wallabag := testWallabag(ts.URL)
	_, err := wallabag.GetVersion()

	if err == nil {
		t.Errorf("Missing error")
	}

	wallabag.Client.NewRequest = func(method, URL string, reader io.Reader) (*http.Request, error) {
		return nil, errors.New("Test with error")
	}

	_, err = wallabag.GetVersion()
	if err == nil {
		t.Errorf("Missing error")
	}
}
