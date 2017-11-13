package goWallabag

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestURLParam(t *testing.T) {
	urlValues := url.Values{}
	existParams := EntryExitsParams{}

	want := "queryUrlTest"

	paramSetterFunc := existParams.URL(want)
	paramSetterFunc(&urlValues)

	got := urlValues.Get("url")
	if got != want {
		t.Errorf("TestUrlParams want: %v, got: %v", want, got)
	}
}

func TestURLsParams(t *testing.T) {
	urlValues := url.Values{}
	existParams := EntryExitsParams{}

	want := []string{
		"one",
		"two",
	}

	paramSetterFunc := existParams.URLs(want)
	paramSetterFunc(&urlValues)

	got := urlValues["urls[]"]
	if got == nil || got[0] != want[0] || got[1] != want[1] {
		t.Errorf("TestUrlParams want: %v, got: %v", want, got)
	}
}

func TestEntryExistsRequest(t *testing.T) {
	client := WallabagClient{
		URL: "test.wallabag/",
	}

	request, err := EntryExistsRequest(client, nil)

	if err != nil {
		t.Error("Error on entry exists request")
		return
	}

	want := "test.wallabag/api/entries/exists.json?"
	got := request.URL.String()
	if got != want {
		t.Errorf("Entry Exists request failed want: %v got %v", want, got)
		return
	}

	request, err = EntryExistsRequest(client, func(values *url.Values) {
		values.Add("test", "param")
	})

	if err != nil {
		t.Error("Error on entry exists request")
		return
	}

	want = "test.wallabag/api/entries/exists.json?test=param"
	got = request.URL.String()
	if got != want {
		t.Errorf("Entry Exists request failed want: %v got %v", want, got)
		return
	}
}

func TestEntryExistsDefaultParser(t *testing.T) {
	failReader := strings.NewReader("Failed")
	_, err := EntryExistsDefaultParser(failReader)

	if err == nil {
		t.Error("Entry exists must fail")
	}

	successreader := strings.NewReader("{\"exists\": true}")
	got, err := EntryExistsDefaultParser(successreader)

	if err != nil {
		t.Errorf("Entry exists must not failed got: %v", err)
	}

	want := true
	if got.Exists != want {
		t.Errorf("Entry exists failed want: %v got %v", want, got.Exists)
	}
}

func TestEntriesExistsDefaultParser(t *testing.T) {
	failReader := strings.NewReader("Failed")
	_, err := EntriesExistsDefaultParser(failReader)

	if err == nil {
		t.Error("Entry exists must fail")
	}

	successreader := strings.NewReader("{\"url1\": true, \"url2\": false}")
	got, err := EntriesExistsDefaultParser(successreader)

	if err != nil {
		t.Errorf("Entry exists must not failed got: %v", err)
	}

	want := EntriesExists{
		"url1": true,
		"url2": false,
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Entry exists failed want: %v got %v", want, got)
	}
}
