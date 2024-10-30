package utils_test

import (
	"crypto"
	"errors"
	"testing"
	"time"

	"auth-service/constants"
	"auth-service/utils"

	"github.com/golang-jwt/jwt"
)

var h = newTestHelpers()

// TestGenerateTokenWithCustomClaims  tests the GenerateTokenWithCustomClaims function
func TestGenerateTokenWithCustomClaims(t *testing.T) {
	email := h.generateRandomEmail(10) // Generate a random email from helper func

	claims := utils.NewTokenClaims(email, time.Now().UTC().Unix())

	token, err := utils.GenerateTokenWithCustomClaims(claims, h.defaultSecret, time.Now().UTC().Add(time.Hour*72).Unix())
	if err != nil {
		t.Errorf("GenerateTokenWithCustomClaims returned error: %v", err)
	}
	if token == "" {
		t.Fatal("expected a token, got an empty  string")
	}
}

// TestValidateToken_Success  tests the ValidateToken function with a valid token
func TestValidateToken_Success(t *testing.T) {
	email := h.generateRandomEmail(10)

	claims := utils.NewTokenClaims(email, time.Now().UTC().Unix())
	token, _ := utils.GenerateTokenWithCustomClaims(claims, h.defaultSecret, time.Now().UTC().Add(time.Hour*9).Unix())

	parsedClaims, err := utils.ValidateToken(token, h.defaultSecret)
	if err != nil {
		t.Errorf("ValidateToken returned error: %v", err)
	}

	if parsedClaims["email"] != email {
		t.Errorf("expected email to be %s, got %s", email, parsedClaims["email"])
	}
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	// create  a token that is expired
	email := h.generateRandomEmail(5)
	issuedAt := time.Now().UTC().Add(time.Hour * -72).Unix()
	expiresAt := time.Now().UTC().Add(time.Hour * -1).Unix()

	claims := utils.NewTokenClaims(email, issuedAt)
	token, _ := utils.GenerateTokenWithCustomClaims(claims, h.defaultSecret, expiresAt)

	_, err := utils.ValidateToken(token, h.defaultSecret)

	if err == nil {
		t.Fatal("Expected an error for expired token, got none")
	}
	//  check that the error is of Expired type
	if err.Error() != constants.TokenExpired {
		t.Errorf("Expected error to be %s, got %s", constants.TokenExpired, err)

	}

}

// TestValidateToken_MalformedToken  tests the ValidateToken function with a malformed token
func TestValidateToken_MalformedToken(t *testing.T) {
	// create  a malformed token
	malformedToken := "malformed.token.string"

	_, err := utils.ValidateToken(malformedToken, h.defaultSecret)

	if err == nil {
		t.Fatal("Expected an error for malformed token, got none")
	}

	//  check if the error str is constants.TokenMalformed
	if err.Error() != constants.TokenMalformed {
		t.Fatalf("Expected token malformed error, got %v", err)
	}
}

func TestValidateToken_Malformed(t *testing.T) {
	issuedAt := time.Now().UTC().Unix()
	expiresAt := time.Now().UTC().Add(time.Hour * 72).Unix()

	claims := utils.NewTokenClaims(h.generateRandomEmail(10), issuedAt)
	token, _ := utils.GenerateTokenWithCustomClaims(claims, h.defaultSecret, expiresAt)

	// tamper the token
	tamperedToken := token[7:] + "tampered"

	_, err := utils.ValidateToken(tamperedToken, h.defaultSecret)

	if err == nil {
		t.Fatal("Expected an error for tampered token, got none")
	}
	if err.Error() != constants.TokenMalformed {
		t.Fatalf("Expected token invalid error, got %v", err)
	}
}

// TestGenerateTokenWithCustomClaims_SigningError tests the signing error
func TestGenerateTokenWithCustomClaims_SigningError(t *testing.T) {
	// Setup
	issuedAt := time.Now().UTC().Unix()
	expiringAt := time.Now().UTC().Add(time.Minute * 2).Unix()
	claims := utils.NewTokenClaims(h.generateRandomEmail(10), issuedAt)

	// Store original signing method
	originalSigningMethod := jwt.SigningMethodHS256
	// Create a broken signing method
	jwt.SigningMethodHS256 = &jwt.SigningMethodHMAC{
		Name: "HS256",
		Hash: crypto.Hash(0),
	} // This will cause an error during signing

	// Defer the reset of the signing method
	defer func() {
		jwt.SigningMethodHS256 = originalSigningMethod
	}()

	// Act
	_, err := utils.GenerateTokenWithCustomClaims(claims, "any-secret", expiringAt)

	// Assert
	if err == nil {
		t.Error("Expected ErrSigningToken with modified signing method, got nil")
	}
	if err != nil && !errors.Is(err, utils.ErrSigningToken) {
		t.Errorf("Expected error to be %v, got %v", utils.ErrSigningToken, err)
	}
}

/*
1. GenerateTokenWithCustomClaims
3. GenerateTokenWithCustomClaims
*/

func BenchmarkGenerateTokenWithCustomClaims(b *testing.B) {
	issuedAt := time.Now().UTC().Unix()
	expiringAt := time.Now().UTC().Add(time.Minute * 2).Unix()
	claims := utils.NewTokenClaims(h.generateRandomEmail(10), issuedAt)

	b.ResetTimer() // reset before the benchmark starts
	for i := 0; i < b.N; i++ {
		_, err := utils.GenerateTokenWithCustomClaims(claims, h.defaultSecret, expiringAt)
		if err != nil {
			b.Errorf("Error generating token: %v", err)
		}
	}

}

var (
	testEmail   = h.generateRandomEmail(10)
	testSecret  = h.defaultSecret
	testExpires = time.Now().UTC().Add(time.Hour * 2).Unix() // Example expiration timestamp
)

// Benchmark for validating a token
func BenchmarkValidateToken(b *testing.B) {
	issuedAt := time.Now().Unix()
	claims := utils.NewTokenClaims(testEmail, issuedAt)
	token, err := utils.GenerateTokenWithCustomClaims(claims, testSecret, testExpires)
	if err != nil {
		b.Fatalf("Error generating token for benchmark: %v", err)
	}

	b.ResetTimer() // Reset the timer before the benchmark starts
	for i := 0; i < b.N; i++ {
		_, err := utils.ValidateToken(token, h.defaultSecret)
		if err != nil {
			b.Errorf("Error validating token: %v", err)
		}
	}
}
