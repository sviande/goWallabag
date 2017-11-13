package goWallabag

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestParseEntries(t *testing.T) {
	_, err := EntriesDefaultParser(strings.NewReader("Failed"))
	if err == nil {
		t.Errorf("Entry parser must failed")
	}

	file := "test/entries.json"
	in, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Missing test file %s", file)
		return
	}

	got, err := EntriesDefaultParser(bytes.NewReader(in))
	if err != nil {
		t.Errorf("err %s", err)
		return
	}

	want := getWantedEntriesResponse()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Parse() got \n%q, want\n%q", got, want)
	}
}

func TestEntriesGetUrl(t *testing.T) {
	client := WallabagClient{
		URL: "test.wallabag/",
	}

	got := EntriesGetURL(client, nil)
	want := "test.wallabag/api/entries.json?"
	if got != want {
		t.Errorf("Entry get url failed want: %v got %v", want, got)
		return
	}

	got = EntriesGetURL(client, func(values *url.Values) {
		values.Add("test", "param")
	})

	want = "test.wallabag/api/entries.json?test=param"
	if got != want {
		t.Errorf("Entry get url failed want: %v got %v", want, got)
		return
	}
}

func TestEntriesRequest(t *testing.T) {
	client := WallabagClient{
		URL: "test.wallabag/",
	}

	req, err := EntriesRequest(client, "url")
	if err != nil {
		t.Error("Entries Request must not failed")
	}

	if req.Method != http.MethodGet {
		t.Error("Entries Request must use get http method")
	}

	want := "url"
	got := req.URL.String()
	if got != want {
		t.Errorf("Entries Request url want: %v got: %v", want, got)
	}
}

func getWantedEntriesResponse() EntriesResponse {
	createdAt0, _ := time.Parse(TimeFormat, "2017-09-18T10:45:58+0200")
	updatedAt0, _ := time.Parse(TimeFormat, "2017-09-18T18:55:26+0200")

	createdAt1, _ := time.Parse(TimeFormat, "2017-09-18T10:45:49+0200")
	updatedAt1, _ := time.Parse(TimeFormat, "2017-09-18T18:47:44+0200")

	return EntriesResponse{
		Page:  1,
		Limit: 30,
		Total: 1438,
		Pages: 48,
		Links: Links{
			Self: Link{
				Href: "self",
			},
			Last: Link{
				Href: "last",
			},
			First: Link{
				Href: "first",
			},
			Next: Link{
				Href: "next",
			},
		},
		Embedded: Embedded{
			Entries: []Entry{
				Entry{
					IsArchived: 1,
					IsStarred:  0,
					UserName:   "example",
					UserEmail:  "example@gmail.com",
					UserID:     1,
					Tags: []Tag{
						Tag{
							ID:    3,
							Label: "Code",
							Slug:  "code",
						},
					},
					ID:          1788,
					Title:       "Title 1788",
					URL:         "https://blog.octo.com/le-demi-cercle-episode-3-communication-breakdown/",
					CreatedAt:   Time{createdAt0},
					UpdatedAt:   Time{updatedAt0},
					Annotations: []string{},
					MimeType:    "text/html",
					Language:    "fr-FR",
					ReadingTime: 8,
					DomainName:  "blog.octo.com",
					HTTPStatus:  "200",
					Links: Links{
						Self: Link{
							Href: "/api/entries/1788",
						},
					},
				},
				Entry{
					IsArchived:  0,
					IsStarred:   1,
					UserName:    "example",
					UserEmail:   "example@gmail.com",
					UserID:      1,
					Tags:        []Tag{},
					ID:          1787,
					Title:       "Le demi-cercle (épisode 2 — Voir / Avancer)",
					URL:         "https://blog.octo.com/le-demi-cercle-episode-2-voir-avancer/",
					CreatedAt:   Time{createdAt1},
					UpdatedAt:   Time{updatedAt1},
					Annotations: []string{},
					MimeType:    "text/html",
					Language:    "fr-FR",
					ReadingTime: 8,
					DomainName:  "blog.octo.com",
					HTTPStatus:  "200",
					Links: Links{
						Self: Link{
							Href: "/api/entries/1787",
						},
					},
				},
			},
		},
	}
}
