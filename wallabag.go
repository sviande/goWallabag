package goWallabag

import (
	"github.com/pkg/errors"
	"net/http"
)

//Wallabag struct use for save authResponse, API URL and httpClient
type Wallabag struct {
	Client WallabagClient
}

//NewWallabag create new struct
func NewWallabag(URL string, client *http.Client) Wallabag {
	return Wallabag{
		Client: WallabagClient{
			URL:    URL,
			Client: client,
		},
	}
}

//Auth User
func (w *Wallabag) Auth(request AuthRequest) error {
	return AuthQuery(&w.Client, request)
}

//GetEntries retrieve entries from params
func (w Wallabag) GetEntries(params ...ParamsSetter) (EntriesResponse, error) {
	URL := EntriesGetURL(w.Client, params...)
	return w.GetEntriesFromURL(URL)
}

//GetEntriesFromURL retrieve entries from url
func (w Wallabag) GetEntriesFromURL(URL string) (EntriesResponse, error) {
	request, err := EntriesRequest(w.Client, URL)
	if err != nil {
		return EntriesResponse{}, errors.Wrap(err, "Error on GetEntriesFromUrl")
	}

	return EntriesFromRequest(w.Client, request, EntriesDefaultParser)
}

//EntryExists retrieve entry exists with one url
func (w Wallabag) EntryExists(params ...ParamsSetter) (EntryExists, error) {
	request, err := EntryExistsRequest(w.Client, params...)
	if err != nil {
		return EntryExists{}, errors.Wrap(err, "Failed to create request for entry exists")
	}

	return EntryExistsFromRequest(w.Client, request, EntryExistsDefaultParser)
}

//EntriesExists retrieve entry exists with multiple urls
func (w Wallabag) EntriesExists(params ...ParamsSetter) (EntriesExists, error) {
	request, err := EntryExistsRequest(w.Client, params...)
	if err != nil {
		return EntriesExists{}, errors.Wrap(err, "Failed to create request for entries exists")
	}

	return EntriesExistsFromRequest(w.Client, request, EntriesExistsDefaultParser)
}

//GetTags retrives all tags
func (w Wallabag) GetTags() ([]Tag, error) {
	request, err := TagsRequest(w.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create request for tags")
	}

	return GetTags(w.Client, request)
}

//GetVersion retrieve version
func (w Wallabag) GetVersion() (string, error) {
	request, err := VersionRequest(w.Client)
	if err != nil {
		return "", errors.Wrap(err, "Failed to create request for version")
	}

	return VersionFetch(w.Client, request)
}
