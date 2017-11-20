package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
)

//WallabagClient struct used for requesting API
type WallabagClient struct {
	Client     *http.Client
	URL        string
	auth       AuthResponse
	NewRequest func(string, string, io.Reader) (*http.Request, error)
}

//Do func use for execute a request on API, need authResponse for token generation
func (w WallabagClient) Do(r *http.Request) (*http.Response, error) {
	if (w.auth == AuthResponse{}) {
		return nil, errors.New("No auth token please run AuthQuery before")
	}

	header, err := w.auth.GetHeader()

	if err != nil {
		return nil, err
	}

	r.Header.Set("Authorization", header)

	resp, err := w.Client.Do(r)

	if err != nil {
		return nil, errors.Wrapf(
			err,
			"Error durring request on URL %v:", r.URL.String(),
		)
	}

	if resp.StatusCode != http.StatusOK {
		parsedError := w.ParseError(resp.StatusCode, resp.Body)
		return nil, parsedError
	}

	return resp, nil
}

//ParseError func use for parsing error from API
func (w WallabagClient) ParseError(statusCode int, readCloser io.ReadCloser) error {
	defer deferClose(readCloser)

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

func deferClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Println(errors.Wrap(err, "Failed to close a resourcer"))
	}
}
