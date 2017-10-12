package goWallabag

import (
	"net/url"
	"testing"
)

var entryParam = EntryParams{}

func testSetter(t *testing.T, val string, key string, setter ParamsSetter) {
	expected := url.Values{}
	expected.Set(key, val)

	got := url.Values{}
	setter(&got)

	if got.Get(key) != expected.Get(key) {
		t.Errorf("For params %v got: %v expected: %v", key, got, expected)
	}
}

func TestIsStarred(t *testing.T) {
	const KeyStarred = "starred"

	testSetter(t, "0", KeyStarred, entryParam.IsStarred(false))
	testSetter(t, "1", KeyStarred, entryParam.IsStarred(true))
}

func TestIsArchive(t *testing.T) {
	const KeyArchive = "archive"

	testSetter(t, "0", KeyArchive, entryParam.IsArchive(false))
	testSetter(t, "1", KeyArchive, entryParam.IsArchive(true))
}

func TestSort(t *testing.T) {
	const KeySort = "sort"
	testSetter(t, "testSortString", KeySort, entryParam.Sort("testSortString"))
}

func TestOrder(t *testing.T) {
	const KeyOrder = "order"
	testSetter(t, "testOrderString", KeyOrder, entryParam.Order("testOrderString"))
}

func TestPage(t *testing.T) {
	const KeyPage = "page"
	testSetter(t, "123", KeyPage, entryParam.Page(123))
}

func TestPerPage(t *testing.T) {
	const KeyPerPage = "perPage"
	testSetter(t, "123", KeyPerPage, entryParam.PerPage(123))
}

func TestTags(t *testing.T) {
	const KeyTags = "tags"
	testSetter(t, "firstTag", KeyTags, entryParam.WithTags([]string{"firstTag"}))
	testSetter(t, "firstTag,SecondTag", KeyTags, entryParam.WithTags([]string{"firstTag", "SecondTag"}))
}
