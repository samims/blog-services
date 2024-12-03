// File: models/errors.go

package models

import "errors"

// Blog errors that can occur when working with blog models
var (
	ErrInvalidTitle     = errors.New("blog: invalid title")
	ErrInvalidContent   = errors.New("blog: invalid content")
	ErrInvalidAuthorID  = errors.New("blog: invalid author ID")
	ErrInvalidPageSize  = errors.New("blog: invalid page size")
	ErrInvalidPage      = errors.New("blog: invalid page number")
	ErrInvalidSortField = errors.New("blog: invalid sort field")
	ErrInvalidSortOrder = errors.New("blog: invalid sort order")
	ErrBlogNotFound     = errors.New("blog: not found")
	ErrDBOperation      = errors.New("blog: database operation failed")
	ErrBlogCreateFailed = errors.New("blog: creation failed")
)

// Validation errors that can occur during request validation

var (
	ErrInvalidRequest = errors.New("validation: invalid request")
)
