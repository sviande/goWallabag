package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
)

type EntriesResponse struct {
	Page     int
	Limit    int
	Pages    int
	Total    int
	Links    Links    `json:"_links"`
	Embedded Embedded `json:"_embedded"`
}

type Embedded struct {
	Entries []Entry `json:"items"`
}

type Links struct {
	Self  Link `json:"self"`
	Last  Link `json:"last"`
	First Link `json:"first"`
	Next  Link `json:"next"`
}

type Link struct {
	Href string `json:"href"`
}

type Entry struct {
	IsArchived  int      `json:"is_archived"`
	IsStarred   int      `json:"is_starred"`
	UserName    string   `json:"user_name"`
	UserEmail   string   `json:"user_email"`
	UserID      int      `json:"user_id"`
	Tags        []Tags   `json:"tags"`
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	CreatedAt   Time     `json:"created_at"`
	UpdatedAt   Time     `json:"updated_at"`
	Annotations []string `json:"annotations"`
	MimeType    string   `json:"mimetype"`
	Language    string   `json:"language"`
	ReadingTime int      `json:"reading_time"`
	DomainName  string   `json:"domain_name"`
	HTTPStatus  string   `json:"http_status"`
	Content     string   `json:"content"`
	Links       Links    `json:"_links"`
}

type EntriesRequest url.Values

func parseEntries(reader io.Reader) (EntriesResponse, error) {
	entries := EntriesResponse{}
	err := json.NewDecoder(reader).Decode(&entries)

	if err != nil {
		return EntriesResponse{}, err
	}

	return entries, nil
}

func EntriesGetURL(w Wallabag, options ...ParamsSetter) string {

	params := url.Values{}
	for _, opt := range options {
		opt(&params)
	}

	entriesURL := "api/entries.json"

	fullURL := w.URL + entriesURL + "?" + params.Encode()

	return fullURL
}

func EntriesFromURL(w Wallabag, fullURL string) (EntriesResponse, error) {
	entriesRequest, err := http.NewRequest(
		http.MethodGet,
		fullURL,
		nil,
	)

	resp, err := w.Do(entriesRequest)

	if err != nil {
		return EntriesResponse{}, errors.Wrap(err, "Failed to retrieve response")
	}

	return parseEntries(resp.Body)
}
