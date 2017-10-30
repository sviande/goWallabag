package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

//Tag struct represent a wallabag tag
type Tag struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Slug  string `json:"slug"`
}

const tagsPathURL = "api/tags.json"

func parseTags(reader io.Reader) ([]Tag, error) {
	tags := make([]Tag, 0)
	err := json.NewDecoder(reader).Decode(&tags)

	return tags, errors.Wrap(err, "Failed to parse tags")
}

//GetTags fetch tag list from API
func GetTags(w WallabagClient) ([]Tag, error) {
	tagRequest, err := http.NewRequest(
		http.MethodGet,
		w.URL+tagsPathURL,
		nil,
	)

	if err != nil {
		return nil, errors.New("Failed to create request")
	}

	resp, err := w.Do(tagRequest)

	defer deferClose(resp.Body)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve response")
	}

	return parseTags(resp.Body)
}
