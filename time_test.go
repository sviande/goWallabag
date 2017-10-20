package goWallabag

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeUnmarshalJSON(t *testing.T) {

	expected, _ := time.Parse(time.RFC3339, "2017-09-18T10:45:58+02:00")

	got := Time{}
	err := json.Unmarshal([]byte("\"2017-09-18T10:45:58+0200\""), &got)
	if err != nil {
		t.Error("Time error on unmarshal")
	}

	if !got.Equal(expected) {
		t.Errorf("Time got : %v expected : %v", got, expected)
	}

	got = Time{}
	err = json.Unmarshal([]byte("2006"), &got)
	if err == nil {
		t.Error(err)
		t.Errorf("Time bad format error expected")
	}
}
