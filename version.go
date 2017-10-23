package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

const versionPathURL = "api/version.json"

func parseVersion(reader io.Reader) (string, error) {
	version := ""
	err := json.NewDecoder(reader).Decode(&version)

	return version, err
}

//VersionRequest generate http.Request for fetching version from API
func VersionRequest(w Wallabag) (*http.Request, error) {
	return http.NewRequest(
		http.MethodGet,
		w.URL+versionPathURL,
		nil,
	)
}

//VersionFetch fetch version API
func VersionFetch(w Wallabag, versionRequest *http.Request) (string, error) {
	var resp *http.Response
	resp, err := w.Do(versionRequest)

	if err != nil {
		return "", errors.Wrap(err, "Version error during get")
	}

	defer deferClose(resp.Body)

	return parseVersion(resp.Body)
}
