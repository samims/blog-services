package utils

import (
	"auth-service/constants"

	"errors"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

// GenerateTokenWithCustomClaims  generates a token with custom claims
func GenerateTokenWithCustomClaims(claims TokenClaims, secret string, expiresAt int64) (string, error) {
	// set  expiration time
	claims.ExpiresAt = expiresAt

	// create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with secret  key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

// ValidateToken checks the validity of the JWT token and returns the claims if valid
func ValidateToken(tokenString, secret string) (jwt.MapClaims, error) {
	// define key function
	validateSigningMethod := func(token *jwt.Token) (interface{}, error) {
		//  ensuring that the signing method is the  one we expect (e.g.,  HS256)
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.NewValidationError(constants.TokenUnExpectedSigningMethod, jwt.ValidationErrorUnverifiable)
		}
		//
		return []byte(secret), nil
	}

	//  parse the token using the defined keyF function validateSigningMethod
	// TODO: check the ParseWithClaims  function
	token, err := jwt.Parse(tokenString, validateSigningMethod)
	if err != nil {

		// check for specific validation errors
		var validationErr *jwt.ValidationError
		if errors.As(err, &validationErr) {
			if validationErr.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, jwt.NewValidationError(constants.TokenExpired, jwt.ValidationErrorExpired)
			}
			if validationErr.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(constants.TokenNotValidYet)
			}
			if validationErr.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(constants.TokenMalformed)
			}
		}
		return nil, err // return original error if it's not a validation
	}

	// check if the token is valid and extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.NewValidationError(constants.TokenInvalid, jwt.ValidationErrorMalformed)
}
