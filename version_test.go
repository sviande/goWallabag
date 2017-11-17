package goWallabag

import (
	"strings"
	"testing"
)

func TestParseVersion(t *testing.T) {
	jsonVersion := "\"2.3.3\""
	got, err := parseVersion(strings.NewReader(jsonVersion))

	if err != nil {
		t.Errorf("err %s", err)
		return
	}

	want := "2.3.3"

	if got != want {
		t.Errorf("Parse() got \n%q, want\n%q", got, want)
	}

	_, err = parseVersion(strings.NewReader("2.3.3"))

	if err == nil {
		t.Errorf("err %s", err)
		return
	}
}

func TestVersionRequest(t *testing.T) {
	client := WallabagClient{
		URL: "test.wallabag/",
	}

	req, err := VersionRequest(client)
	if err != nil {
		t.Errorf("Version get request failed %v", err)
	}
	want := "test.wallabag/api/version.json"
	got := req.URL.String()
	if got != want {
		t.Errorf("Entry get url failed want: %v got %v", want, got)
		return
	}
}
