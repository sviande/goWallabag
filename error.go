package goWallabag

//ErrorResponse struct represent error response from wallabag API
type ErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}
