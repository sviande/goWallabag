package goWallabag

import (
	"net/http"
)

type Wallabag struct {
	Client http.Client
	URL    string
}

func (w Wallabag) Do(r *http.Request) (*http.Response, error) {
	return w.Client.Do(r)
}
