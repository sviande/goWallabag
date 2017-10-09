package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
)

type AuthRequest struct {
	GrantType    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

func (r AuthRequest) GetUrlValues() url.Values {
	urlValues := url.Values{}
	urlValues.Set("grant_type", r.GrantType)
	urlValues.Set("client_id", r.ClientID)
	urlValues.Set("client_secret", r.ClientSecret)
	urlValues.Set("username", r.Username)
	urlValues.Set("password", r.Password)

	return urlValues
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func (a AuthResponse) GetHeader() string {
	token := strings.ToUpper(string(a.TokenType[0]))
	token += a.TokenType[1:len(a.TokenType)]
	token += " " + a.AccessToken

	return token
}

func AuthQuery(w *Wallabag, authRequest AuthRequest) error {
	authURL := "oauth/v2/token"

	resp, err := w.Client.PostForm(w.URL+authURL, authRequest.GetUrlValues())

	if err != nil {
		return errors.Wrap(err, "Auth Failed from http")
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		errorResponse := ErrorResponse{}
		err = decoder.Decode(&errorResponse)

		if err != nil {
			return errors.Wrap(err, "Auth: Failed to parse error response")
		}

		return errors.Errorf(
			"Auth: return status code: %v with message:\n %v",
			resp.StatusCode,
			errorResponse,
		)
	}

	w.auth = AuthResponse{}
	err = decoder.Decode(&w.auth)

	if err != nil {
		return errors.Wrap(err, "Auth Failed from json decode")
	}

	return nil
}
