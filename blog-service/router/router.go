package router

import (
	"fmt"
	"log"
	"net/http"

	"blog-service/controllers"
	"blog-service/middleware"
)

const (
	// V1 represents the first version of the API.
	V1 = "/api/v1"
	// blogsPath is the base path for blog-related endpoints.
	blogsPath = "/blogs"
	// blogDetailPath is the path for accessing a specific blog by its ID.
	blogDetailPath = "/blogs/{id}"
)

// Middleware defines a function type for HTTP middleware.
// It takes a http.Handler and returns a http.Handler.
type Middleware func(http.Handler) http.Handler

// route represents an API endpoint configuration.
// It contains the HTTP method, version, path, handler, and a human-readable name.
type route struct {
	method  string       // HTTP method (GET, POST, PUT, DELETE)
	version string       // API version this route belongs to
	path    string       // URL path for the endpoint
	handler http.Handler // HTTP handler function for this route
	name    string       // Human-readable name for the route
}

// createVersionPath constructs a complete endpoint path by combining the API version and the specified path.
func createVersionPath(version, path string) string {
	return fmt.Sprintf("%s%s", version, path)
}

// createPattern generates a complete route pattern by combining the HTTP method and the versioned path.
//
// Parameters:
//   - method: HTTP method (e.g., "GET", "POST")
//   - version: API version (e.g., "/api/v1")
//   - path: Endpoint path (e.g., "/blogs")
//
// Returns:
//   - A complete route pattern (e.g., "GET /api/v1/blogs").
func createPattern(method, version, path string) string {
	return fmt.Sprintf("%s %s", method, createVersionPath(version, path))
}

// Init initializes and returns a configured HTTP router with all API routes.
// This function sets up all available API endpoints for the blog service.
// Currently, it configures V1 routes only, but the structure allows for easy addition of new API versions.
//
// Parameters:
//   - blogCtrl: An instance of BlogController that implements all handler methods.
//
// Returns:
//   - *http.ServeMux: A configured HTTP router with all routes registered.
func Init(blogCtrl controllers.BlogController) *http.ServeMux {
	mux := http.NewServeMux()

	// Define the routes for version 1 of the API.
	routes := []route{
		{
			method:  http.MethodGet,
			path:    blogsPath,
			handler: Middleware(middleware.RequestIDMiddleware)(http.HandlerFunc(blogCtrl.GetBlogList)),
			version: V1,
			name:    "List Blogs",
		},
		{
			method:  http.MethodPost,
			path:    blogsPath,
			handler: Middleware(middleware.RequestIDMiddleware)(http.HandlerFunc(blogCtrl.CreateBlog)),
			version: V1,
			name:    "Create Blog",
		},
		{
			method:  http.MethodGet,
			path:    blogDetailPath,
			handler: Middleware(middleware.RequestIDMiddleware)(http.HandlerFunc(blogCtrl.GetBlogByID)),
			version: V1,
			name:    "Get Blog Detail",
		},
		{
			method:  http.MethodPut,
			path:    blogDetailPath,
			handler: Middleware(middleware.RequestIDMiddleware)(http.HandlerFunc(blogCtrl.UpdateBlog)),
			version: V1,
			name:    "Update Blog Detail",
		},
		{
			method:  http.MethodDelete,
			path:    blogDetailPath,
			handler: Middleware(middleware.RequestIDMiddleware)(http.HandlerFunc(blogCtrl.DeleteBlog)),
			version: V1,
			name:    "Delete Blog Detail",
		},
	}

	// Register all defined routes with the HTTP multiplexer.
	log.Print("\n\n")
	log.Println("Registering routes.....")
	for _, route := range routes {
		pattern := createPattern(route.method, route.version, route.path)
		log.Println(pattern)

		mux.Handle(pattern, route.handler)
	}

	return mux
}
