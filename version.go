package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

//VersionPathURL url path for version
const VersionPathURL = "api/version.json"

func parseVersion(reader io.Reader) (string, error) {
	version := ""
	err := json.NewDecoder(reader).Decode(&version)

	return version, err
}

//VersionFetch fetch version API
func VersionFetch(w WallabagClient, versionRequest *http.Request) (string, error) {
	var resp *http.Response
	resp, err := w.Do(versionRequest)

	if err != nil {
		return "", errors.Wrap(err, "Version error during get")
	}

	defer deferClose(resp.Body)

	return parseVersion(resp.Body)
}
