package utils_test

import (
	"fmt"

	"golang.org/x/exp/rand"
)

// testHelpers contains all helper functions and constants used in tests
type testHelpers struct {
	defaultSecret string
	domains       []string
}

// newTestHelpers creates a new instance of testHelpers
func newTestHelpers() *testHelpers {
	return &testHelpers{
		defaultSecret: "default-secret",
		domains: []string{
			"example.com",
		},
	}
}

// generateRandomEmail generates a random email for testing
func (h *testHelpers) generateRandomEmail(length int) string {
	return fmt.Sprintf("%s@%s", h.randomString(10), h.domains[rand.Intn(len(h.domains))])
}

// randomString generates a random string of specified length
func (h *testHelpers) randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz01234567890")
	s := make([]rune, length)

	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
