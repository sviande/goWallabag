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

//AuthPathURL url path for auth
const AuthPathURL = "oauth/v2/token"

//AuthRequest reprensent an auth request
type AuthRequest struct {
	GrantType    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

//GetURLValues Return url.Values required for auth
func (r AuthRequest) GetURLValues() url.Values {
	urlValues := url.Values{}
	urlValues.Set("grant_type", r.GrantType)
	urlValues.Set("client_id", r.ClientID)
	urlValues.Set("client_secret", r.ClientSecret)
	urlValues.Set("username", r.Username)
	urlValues.Set("password", r.Password)

	return urlValues
}

//AuthResponse represent auth response
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

//GetHeader return header for next request
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

//AuthQuery query wallabag backend for an auth token
func AuthQuery(w *WallabagClient, authRequest AuthRequest) error {

	resp, err := w.Client.PostForm(w.URL+AuthPathURL, authRequest.GetURLValues())

	if err != nil {
		return errors.Wrap(err, "Auth Failed from http")
	}

	if resp.StatusCode != http.StatusOK {
		parsedError := w.ParseError(resp.StatusCode, resp.Body)
		return parsedError
	}

	defer deferClose(resp.Body)

	w.auth = AuthResponse{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&w.auth)

	if err != nil {
		return errors.Wrap(err, "Auth Failed from json decode")
	}

	return nil
}
