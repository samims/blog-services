# Router Package

## Overview
The router package handles all HTTP routing for the blog service API. It implements a versioned routing system that makes it easy to manage different API versions and their endpoints.

## API Versions
Currently supported versions:
- V1 (`/api/v1`)

## Available Routes

### V1 Routes

| Method | Path               | Handler     | Description                 |
|--------|--------------------|-------------|-----------------------------|
| GET    | /api/v1/posts      | GetBlogList | List all blog posts         |
| POST   | /api/v1/posts      | CreateBlog  | Create a new blog post      |
| GET    | /api/v1/posts/{id} | GetBlog     | Get a specific blog post    |
| PUT    | /api/v1/posts/{id} | UpdateBlog  | Update a specific blog post |
| DELETE | /api/v1/posts/{id} | DeleteBlog  | Delete a specific blog post |

## Usage

```go

import (
    "blog-service/controllers"
    "blog-service/router"
    "net/http"
)

func main() {
    // Initialize your controller
    ctrl := controllers.NewBlogController(...)

    // Initialize the router
    r := router.Init(ctrl)

    // Start the server
    http.ListenAndServe(":8080", r)
}
```

### Adding New routes

```go
v1Routes := []route{
    // Existing routes...
    {
        method:  http.MethodGet,
        path:    "/new-endpoint",
        handler: ctrl.NewHandler,
        version: V1,
        name:    "new endpoint",
    },
}
```

### Adding New Versions
To add a new API version
*   Add a new version constant
* Create a new route slice for that version
* Register the routes in `Init()`



This documentation:
```
1. Explains the package's purpose
2. Details the structure and components
3. Provides usage examples
4. Includes instructions for maintenance and expansion
5. Follows Go's documentation conventions
6. Makes it easier for other developers to understand and work with the code
```
