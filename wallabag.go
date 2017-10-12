package goWallabag

import (
	"github.com/pkg/errors"
	"net/http"
)

type Wallabag struct {
	Client http.Client
	URL    string
	auth   AuthResponse
}

func (w Wallabag) Do(r *http.Request) (*http.Response, error) {
	if (w.auth == AuthResponse{}) {
		return nil, errors.New("No auth token please run AuthQuery before")
	}

	header, err := w.auth.GetHeader()

	if err != nil {
		return nil, err
	}

	r.Header.Set("Authorization", header)
	return w.Client.Do(r)
}
