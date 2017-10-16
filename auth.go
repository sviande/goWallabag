package goWallabag

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
)

var errTokenType = errors.New("Impossible to create header no token type")
var errAccesToken = errors.New("Impossible to create header no access token")
const authPathURL =  "oauth/v2/token"

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

func (a AuthResponse) GetHeader() (string, error) {
	if a.TokenType == "" {
		return "", errTokenType
	}

	if a.AccessToken == "" {
		return "", errAccesToken
	}

	token := strings.ToUpper(string(a.TokenType[0]))
	token += a.TokenType[1:len(a.TokenType)]
	token += " " + a.AccessToken

	return token, nil
}

func AuthQuery(w *Wallabag, authRequest AuthRequest) error {

	resp, err := w.Client.PostForm(w.URL+authPathURL, authRequest.GetUrlValues())

	if err != nil {
		return errors.Wrap(err, "Auth Failed from http")
	}

	if resp.StatusCode != http.StatusOK {
		return w.ParseError(resp.StatusCode, resp.Body)
	}

	defer resp.Body.Close()
	w.auth = AuthResponse{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&w.auth)

	if err != nil {
		return errors.Wrap(err, "Auth Failed from json decode")
	}

	return nil
}
