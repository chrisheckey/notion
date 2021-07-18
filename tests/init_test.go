package tests

import (
	_ "embed"
	"testing"
)

//go:embed token.txt
var token string

func testToken(t *testing.T) string {
	if token == "" {
		t.Fatal("empty test token")
	}
	return token
}
