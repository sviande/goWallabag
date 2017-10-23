package goWallabag

import (
	"net/url"
	"strconv"
	"strings"
)

//ParamsSetter use for setting a params for an entry request
type ParamsSetter func(*url.Values)

//EntryParams struct use as namespace for entryParams
type EntryParams struct {
}

//IsStarred func setting params for filtering on stared entries
func (e EntryParams) IsStarred(isStarred bool) ParamsSetter {
	value := "1"
	if !isStarred {
		value = "0"
	}

	return func(params *url.Values) {
		params.Set("starred", value)
	}
}

//IsArchive func setting params for filtering on archive entries
func (e EntryParams) IsArchive(isArchive bool) ParamsSetter {
	value := "1"
	if !isArchive {
		value = "0"
	}

	return func(params *url.Values) {
		params.Set("archive", value)
	}
}

//Sort func setting params for sort
func (e EntryParams) Sort(sort string) ParamsSetter {
	return func(params *url.Values) {
		params.Set("sort", sort)
	}
}

//Order func setting params for order
func (e EntryParams) Order(order string) ParamsSetter {
	return func(params *url.Values) {
		params.Set("order", order)
	}
}

//Page func setting params for page
func (e EntryParams) Page(page int) ParamsSetter {
	return func(params *url.Values) {
		params.Set("page", strconv.Itoa(page))
	}
}

//PerPage func setting params for perpage
func (e EntryParams) PerPage(perPage int) ParamsSetter {
	return func(params *url.Values) {
		params.Set("perPage", strconv.Itoa(perPage))
	}
}

//WithTags func setting params for tags
func (e EntryParams) WithTags(tags []string) ParamsSetter {
	return func(params *url.Values) {
		params.Set("tags", strings.Join(tags, ","))
	}
}
