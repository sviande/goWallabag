package goWallabag

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestParseTags(t *testing.T) {
	file := "test/tags.json"
	in, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Missing test file %s", file)
		return
	}

	got, err := parseTags(bytes.NewReader(in))
	if err != nil {
		t.Errorf("err %s", err)
		return
	}

	want := getWantedTag()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Parse() got \n%q, want\n%q", got, want)
	}

	_, err = parseTags(strings.NewReader("asd"))

	if err == nil {
		t.Errorf("err %s", err)
		return
	}
}

func getWantedTag() []Tag {
	return []Tag{
		Tag{
			ID:    8,
			Label: "Agile",
			Slug:  "agile",
		},
		Tag{
			ID:    32,
			Label: "Best practices",
			Slug:  "best-practices",
		},
		Tag{
			ID:    27,
			Label: "Sécurité",
			Slug:  "securite",
		},
		Tag{
			ID:    13,
			Label: "Web Perf",
			Slug:  "web-perf",
		},
	}
}
