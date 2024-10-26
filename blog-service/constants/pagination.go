package constants

const (
	DefaultPage     int = 1
	DefaultPageSize int = 20
	MaxPageSize     int = 100

	DefaultSortBy    string = SortOrderDesc
	DefaultSortOrder string = "created_at"

	SortOrderAsc  string = "asc"
	SortOrderDesc string = "desc"
)

// ValidSortFields maps entity names to their valid sort fields
var ValidSortFields = map[string][]string{
	"blog": {
		"title",
		"created_at",
		"updated_at",
		"author_id",
	},
}
