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

	got, err = parseVersion(strings.NewReader("2.3.3"))

	if err == nil {
		t.Errorf("err %s", err)
		return
	}
}
