package goWallabag

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestParseEntries(t *testing.T) {
	file := "test/entries.json"
	in, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Missing test file %s", file)
		return
	}

	got, err := EntryListParse(bytes.NewReader(in))
	if err != nil {
		t.Errorf("err %s", err)
		return
	}

	want := getWantedEntriesResponse()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Parse() got \n%q, want\n%q", got, want)
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
