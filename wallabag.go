package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
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

func (w Wallabag) ParseError(statusCode int, readCloser io.ReadCloser) error {
	defer readCloser.Close()

	errorResponse := ErrorResponse{}
	err := json.NewDecoder(readCloser).Decode(&errorResponse)

	if err != nil {
		return errors.Wrap(err, "Failed to parse error response")
	}

	return errors.Errorf(
		"Return status code: %v with message:\n %v",
		statusCode,
		errorResponse,
	)
}
