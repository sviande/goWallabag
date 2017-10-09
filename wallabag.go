package goWallabag

import (
	"net/http"
)

type Wallabag struct {
	Client http.Client
	URL    string
	auth   AuthResponse
}

func (w Wallabag) Do(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", w.auth.GetHeader())
	return w.Client.Do(r)
}
