// File: models/errors.go

package models

import "errors"

// Common errors that can occur when working with blog models
var (
	ErrInvalidTitle     = errors.New("invalid blog title")
	ErrInvalidContent   = errors.New("invalid blog content")
	ErrInvalidAuthor    = errors.New("invalid author ID")
	ErrInvalidPageSize  = errors.New("invalid page size")
	ErrInvalidPage      = errors.New("invalid page number")
	ErrInvalidSortField = errors.New("invalid sort field")
	ErrInvalidSortOrder = errors.New("invalid sort order")
)
