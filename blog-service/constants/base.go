package constants

// Context Keys
type contextKey string

// RequestIDKey is the key used to store the request ID in the context.
const RequestIDKey contextKey = "request_id"

// SourceServiceKey is the key used to store the source service in the context.
const SourceServiceKey contextKey = "source_service"

// API Endpoints
const (
	// ApiV1 represents the base path for version 1 of the API.
	ApiV1 = "/api/v1"
	// BlogsPath is the base path for blog-related endpoints.
	BlogsPath = "/blogs"
	// BlogDetailPath is the path for accessing a specific blog by its ID.
	BlogDetailPath = "/blogs/{id}"
)

// Pagination Defaults
const (
	// DefaultPage is the default page number for paginated requests.
	DefaultPage = 1
	// DefaultPageSize is the default page size for paginated requests.
	DefaultPageSize = 20
	// MaxPageSize is the maximum page size allowed for paginated requests.
	MaxPageSize = 100
)

// Sorting Defaults
const (
	// SortOrderAsc represents the ascending sort order.
	SortOrderAsc = "asc"
	// SortOrderDesc represents the descending sort order.
	SortOrderDesc = "desc"
	// DefaultSortBy is the default field to sort by.
	DefaultSortBy = SortOrderDesc
	// DefaultSortOrder is the default sort order.
	DefaultSortOrder = "created_at"
)

// ValidSortFields maps entity names to their valid sort fields.
var ValidSortFields = map[string][]string{
	"blog": {
		"title",
		"created_at",
		"updated_at",
		"author_id",
	},
}

// Error Messages
const (
	// ErrInvalidEmailOrPass is the error message for invalid email or password.
	ErrInvalidEmailOrPass = "username or password error"

	// Token related errors.
	TokenUnExpectedSigningMethod = "unexpected signing method"
	TokenInvalidSigningMethod    = "invalid signing method"
	TokenInvalid                 = "invalid token"
	TokenExpired                 = "token expired"
	TokenNotValidYet             = "token is not valid yet"
	TokenMalformed               = "token is malformed"
	TokenInvalidIssuer           = "invalid issuer"

	// ErrBlogNotFound Blog related errors.
	ErrBlogNotFound = "blog not found"
)

// Environment Variables
const (
	AppDebug  = "APP_DEBUG"
	SecretKey = "SECRET_KEY"
	BuildEnv  = "BUILD_ENV"
	AppPort   = "APP_PORT"

	PostgresHost       = "POSTGRES_HOST"
	PostgresPort       = "POSTGRES_PORT"
	PostgresUser       = "POSTGRES_USER"
	PostgresPasswd     = "POSTGRES_PASSWORD"
	PostgresDBName     = "POSTGRES_DB"
	DatabaseSSLMode    = "DB_SSL_MODE"
	DatabaseDefaultSSL = "disable"
)

// Header related

const (
	HeaderRequestID = "X-Request-ID"
)
