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

//GetEntryList retrieves entries from params
func (w Wallabag) GetEntryList(params ...ParamsSetter) (EntriesResponse, error) {
	URL := EntriesGetURL(w.Client, params...)
	return w.GetEntryListFromURL(URL)
}

//GetEntryListFromURL retrieve entries from url
func (w Wallabag) GetEntryListFromURL(URL string) (EntriesResponse, error) {
	request, err := EntriesListRequest(w.Client, URL)
	if err != nil {
		return EntriesResponse{}, errors.Wrap(err, "Error on GetEntryList")
	}

	return EntriesFromURL(w.Client, request, EntryListDefaultParser)
}
