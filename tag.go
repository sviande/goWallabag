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

//TagsPathURL Url path for tags
const TagsPathURL = "api/tags.json"

func parseTags(reader io.Reader) ([]Tag, error) {
	tags := make([]Tag, 0)
	err := json.NewDecoder(reader).Decode(&tags)

	return tags, errors.Wrap(err, "Failed to parse tags")
}

//GetTags fetch tag list from API
func GetTags(w WallabagClient, tagsRequest *http.Request) ([]Tag, error) {
	resp, err := w.Do(tagsRequest)

	defer deferClose(resp.Body)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve response")
	}

	return parseTags(resp.Body)
}
