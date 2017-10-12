package goWallabag

import (
	"net/url"
	"strconv"
	"strings"
)

type ParamsSetter func(*url.Values)

type EntryParams struct {
}

func (e EntryParams) IsStarred(isStarred bool) ParamsSetter {
	value := "1"
	if !isStarred {
		value = "0"
	}

	return func(params *url.Values) {
		params.Set("starred", value)
	}
}

func (e EntryParams) IsArchive(isArchive bool) ParamsSetter {
	value := "1"
	if !isArchive {
		value = "0"
	}

	return func(params *url.Values) {
		params.Set("archive", value)
	}
}

func (e EntryParams) Sort(sort string) ParamsSetter {
	return func(params *url.Values) {
		params.Set("sort", sort)
	}
}

func (e EntryParams) Order(order string) ParamsSetter {
	return func(params *url.Values) {
		params.Set("order", order)
	}
}

func (e EntryParams) Page(page int) ParamsSetter {
	return func(params *url.Values) {
		params.Set("page", strconv.Itoa(page))
	}
}

func (e EntryParams) PerPage(perPage int) ParamsSetter {
	return func(params *url.Values) {
		params.Set("perPage", strconv.Itoa(perPage))
	}
}

func (e EntryParams) WithTags(tags []string) ParamsSetter {
	return func(params *url.Values) {
		params.Set("tags", strings.Join(tags, ","))
	}
}
