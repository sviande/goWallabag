package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

const versionPathURL = "api/version.json"

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

	version := ""
	err = json.NewDecoder(resp.Body).Decode(&version)

	if err != nil {
		return "", err
	}

	return version, nil
}
