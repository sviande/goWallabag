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

	if err != nil {
		return version, err
	}

	return version, nil
}

func GetVersion(w Wallabag) (string, error) {
	versionRequest, err := http.NewRequest(
		http.MethodGet,
		w.URL+versionPathURL,
		nil,
	)

	if err != nil {
		return "", errors.Wrap(err, "Version error during get request creation")
	}

	var resp *http.Response
	resp, err = w.Do(versionRequest)

	if err != nil {
		return "", errors.Wrap(err, "Version error during get")
	}

	defer resp.Body.Close()

	return parseVersion(resp.Body)
}
