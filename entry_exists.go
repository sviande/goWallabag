package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
)

//EntryExistsURL url path for entries
const EntryExistsURL = "api/entries/exists.json"

//EntryExitsParams struct use as namespace for entry exits params
type EntryExitsParams struct {
}

//URL func setting params for filtering on stared entries
func (e EntryExitsParams) URL(queryURL string) ParamsSetter {
	return func(params *url.Values) {
		params.Set("url", queryURL)
	}
}

//URLs func setting params for filtering on stared entries
func (e EntryExitsParams) URLs(queryURLs []string) ParamsSetter {
	return func(params *url.Values) {
		for _, queryURL := range queryURLs {
			params.Add("urls[]", queryURL)
		}
	}
}

//EntryExists represent result for one URL
type EntryExists struct {
	Exists bool `json:"exists"`
}

//EntriesExists reprensent list of url with boolean
type EntriesExists map[string]bool

//EntryExistsURLWithParams create request on entryExists
func EntryExistsURLWithParams(options ...ParamsSetter) string {
	params := url.Values{}
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt(&params)
	}

	return EntryExistsURL + "?" + params.Encode()
}

//EntryExistsParser func interface for entry exists parser
type EntryExistsParser func(reader io.Reader) (EntryExists, error)

//EntryExistsDefaultParser parse EntryExists response from io.Reader
func EntryExistsDefaultParser(reader io.Reader) (EntryExists, error) {
	entryExists := EntryExists{}
	err := json.NewDecoder(reader).Decode(&entryExists)

	if err != nil {
		return EntryExists{}, errors.Wrap(err, "Failed to parse EntryExists")
	}

	return entryExists, nil
}

//EntryExistsFromRequest fetch EntryExists from API
func EntryExistsFromRequest(
	w WallabagClient,
	entryExistRequest *http.Request,
	parser EntryExistsParser,
) (EntryExists, error) {
	resp, err := w.Do(entryExistRequest)

	if err != nil {
		return EntryExists{}, errors.Wrap(err, "Failed to retrieve entry exists")
	}

	defer deferClose(resp.Body)

	return parser(resp.Body)
}

//EntriesExistsParser func interface for entry exists parser
type EntriesExistsParser func(reader io.Reader) (EntriesExists, error)

//EntriesExistsDefaultParser parse EntryExists response from io.Reader
func EntriesExistsDefaultParser(reader io.Reader) (EntriesExists, error) {
	entriesExists := EntriesExists{}
	err := json.NewDecoder(reader).Decode(&entriesExists)

	if err != nil {
		return EntriesExists{}, errors.Wrap(err, "Failed to parser EntriesExists")
	}

	return entriesExists, nil
}

//EntriesExistsFromRequest fetch EntryExists from API
func EntriesExistsFromRequest(
	w WallabagClient,
	entryExistsRequest *http.Request,
	parser EntriesExistsParser,
) (EntriesExists, error) {
	resp, err := w.Do(entryExistsRequest)

	if err != nil {
		return EntriesExists{}, errors.Wrap(err, "Failed to retrieve entries exists")
	}

	defer deferClose(resp.Body)

	return parser(resp.Body)
}
