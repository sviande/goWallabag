package goWallabag

import (
	"testing"
)

func TestGetHeader(t *testing.T) {
	authResponse := AuthResponse{
		TokenType:   "bearer",
		AccessToken: "123456",
	}

	expected := "Bearer 123456"
	got := authResponse.GetHeader()
	if expected != got {
		t.Errorf("Failed for auth GetHeader expected %v got %v", expected, got)
	}
}
