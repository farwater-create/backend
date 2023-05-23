package apikey

import (
	"testing"
)

func TestSecureRandomString(t *testing.T) {
	n := 10
	result := secureRandomString(n)
	if len(result) != n {
		t.Errorf("Expected string of length %d, but got %d", n, len(result))
	}
}
