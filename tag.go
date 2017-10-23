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

//TagList reprensent a array of Tag
type TagList []Tag

const tagsPathURL = "api/tags.json"

func parseTagList(reader io.Reader) (TagList, error) {
	tagList := TagList{}
	err := json.NewDecoder(reader).Decode(&tagList)

	if err != nil {
		return TagList{}, err
	}

	return tagList, nil
}

//GetTagList fetch tag list from API
func GetTagList(w Wallabag) (TagList, error) {
	tagRequest, err := http.NewRequest(
		http.MethodGet,
		w.URL+tagsPathURL,
		nil,
	)

	if err != nil {
		return TagList{}, errors.New("Failed to create request")
	}

	resp, err := w.Do(tagRequest)

	if err != nil {
		return TagList{}, errors.Wrap(err, "Failed to retrieve response")
	}

	return parseTagList(resp.Body)
}
