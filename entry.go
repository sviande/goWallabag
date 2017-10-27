package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
)

const entryGetAllPathURL = "api/entries.json"

//EntriesResponse represent struct for wallabag response
//Response with pagination
type EntriesResponse struct {
	Page     int
	Limit    int
	Pages    int
	Total    int
	Links    Links    `json:"_links"`
	Embedded Embedded `json:"_embedded"`
}

//Embedded represent array of Entry in wallabag response
type Embedded struct {
	Entries []Entry `json:"items"`
}

//Links List of link for pagination discovery
type Links struct {
	Self  Link `json:"self"`
	Last  Link `json:"last"`
	First Link `json:"first"`
	Next  Link `json:"next"`
}

//Link represent link for content discovery in EntriesResponse
type Link struct {
	Href string `json:"href"`
}

//Entry represent a Wallabag entry
type Entry struct {
	IsArchived  int      `json:"is_archived"`
	IsStarred   int      `json:"is_starred"`
	UserName    string   `json:"user_name"`
	UserEmail   string   `json:"user_email"`
	UserID      int      `json:"user_id"`
	Tags        TagList  `json:"tags"`
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

//EntriesRequest represent query var for request
type EntriesRequest url.Values

//EntryListParser function interface for parsing API response
type EntryListParser func(io.Reader) (EntriesResponse, error)

//EntryListDefaultParser parse Entries response from io.Reader
func EntryListDefaultParser(reader io.Reader) (EntriesResponse, error) {
	entries := EntriesResponse{}
	err := json.NewDecoder(reader).Decode(&entries)

	if err != nil {
		return EntriesResponse{}, err
	}

	return entries, nil
}

//EntriesGetURL return url for get Entries with query
func EntriesGetURL(w WallabagClient, options ...ParamsSetter) string {

	params := url.Values{}
	for _, opt := range options {
		opt(&params)
	}

	fullURL := w.URL + entryGetAllPathURL + "?" + params.Encode()

	return fullURL
}

//EntriesListRequest create an http request for fetching entries from API
func EntriesListRequest(w WallabagClient, fullURL string) (*http.Request, error) {
	return http.NewRequest(
		http.MethodGet,
		fullURL,
		nil,
	)
}

//EntriesFromURL fetch all entries from url
func EntriesFromURL(
	w WallabagClient,
	entryListRequest *http.Request,
	parser EntryListParser,
) (EntriesResponse, error) {
	resp, err := w.Do(entryListRequest)

	if err != nil {
		return EntriesResponse{}, errors.Wrap(err, "Failed to retrieve entries")
	}

	defer deferClose(resp.Body)

	return parser(resp.Body)
}
