package router

import (
	"blog-service/controllers"
	"fmt"
	"net/http"
)

const (
	// V1 API versions
	V1 = "/api/v1"
	// Base paths
	postsPath = "/posts"
	// Detail path
	postDetailPath = "/posts/{id}"
)

// route defines the structure for an API endpoint.
// It contains all necessary information to register and handle an HTTP route.
type route struct {
	// HTTP method (GET, POST, PUT, DELETE)
	method string

	// URL path for the endpoint
	path string

	// HTTP handler function for this route
	handler http.HandlerFunc

	// API version this route belongs to
	version string

	// Human-readable name for the route, useful for documentation and logging
	name string
}

// createVersionPath combines an API version with a path to create a complete endpoint path.
//
// Parameters:
//   - version: The API version (e.g., "/api/v1")
//   - path: The endpoint path (e.g., "/posts")
//
// Returns:
//   - A complete versioned path (e.g., "/api/v1/posts")
func createVersionPath(version, path string) string {
	return fmt.Sprintf("%s%s", version, path)
}

// createPattern creates a complete route pattern by combining HTTP method and versioned path.
//
// Parameters:
//   - method: HTTP method (e.g., "GET", "POST")
//   - version: API version (e.g., "/api/v1")
//   - path: Endpoint path (e.g., "/posts")
//
// Returns:
//   - A complete route pattern (e.g., "GET /api/v1/posts")
func createPattern(method, version, path string) string {
	return fmt.Sprintf("%s %s", method, createVersionPath(version, path))
}

// Init initializes and returns a configured HTTP router with all API routes.
//
// This function sets up all available API endpoints for the blog service.
// Currently, it configures V1 routes only, but the structure allows for
// easy addition of new API versions.
//
// Parameters:
//   - ctrl: An instance of BlogController that implements all handler methods
//
// Returns:
//   - *http.ServeMux: A configured HTTP router with all routes registered
//
// Usage:
//
//	controller := controllers.NewBlogController(...)
//	router := router.Init(controller)
//	http.ListenAndServe(":8080", router)
func Init(blogCtrl controllers.BlogController) *http.ServeMux {
	mux := http.NewServeMux()

	// v1 routes
	v1Routes := []route{
		{
			method:  http.MethodGet,
			path:    postsPath,
			handler: blogCtrl.GetBlogList,
			version: V1,
			name:    "list posts",
		},
		{
			method:  http.MethodPost,
			path:    postsPath,
			handler: blogCtrl.CreateBlog,
			version: V1,
			name:    "create post",
		},
		{
			method:  http.MethodGet,
			path:    postDetailPath,
			handler: blogCtrl.GetBlog,
			version: V1,
			name:    "get post detail",
		},
		{
			method:  http.MethodPut,
			path:    postDetailPath,
			handler: blogCtrl.UpdateBlog,
			version: V1,
			name:    "update post detail",
		},
		{
			method:  http.MethodDelete,
			path:    postDetailPath,
			handler: blogCtrl.DeleteBlog,
			version: V1,
			name:    "delete post detail",
		},
	}

	// Register all V1 routes
	for _, route := range v1Routes {
		pattern := createPattern(route.method, route.version, route.path)
		mux.HandleFunc(pattern, route.handler)

	}
	return mux
}
