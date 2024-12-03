package controllers

import (
	"blog-service/models/request"
	"encoding/json"
	"net/http"
	"strconv"

	"blog-service/logger"
	"blog-service/services"
)

type BlogController interface {
	GetBlogList(w http.ResponseWriter, r *http.Request)
	GetBlogByID(w http.ResponseWriter, r *http.Request)
	CreateBlog(w http.ResponseWriter, r *http.Request)
	UpdateBlog(w http.ResponseWriter, r *http.Request)
	DeleteBlog(w http.ResponseWriter, r *http.Request)
}

type blogController struct {
	svc services.BlogService
	l   *logger.AppLogger
}

// GetBlogList handles HTTP GET requests to retrieve a list of blogs with pagination.
func (b blogController) GetBlogList(w http.ResponseWriter, r *http.Request) {
	// Extract pagination parameters from the query string
	ctx := r.Context()
	b.l.Info(ctx, "Retrieving blog list ")
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	// Set default values for pagination
	page := 1      // Default to page 1
	pageSize := 10 // Default to page size of 10

	// Parse the page query parameter
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			b.l.Warn(ctx, "Invalid page number provided")
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	}

	// Parse the page size query parameter
	if pageSizeStr != "" {
		var err error
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 {
			b.l.Warn(ctx, "Invalid page size provided")
			http.Error(w, "Invalid page size", http.StatusBadRequest)
			return
		}
	}

	// Create a pagination request object
	pageReq := request.NewPaginationRequest(page, page, "", "")
	b.l.Info(ctx, "Retrieving blog list with page: %d, page size: %d", page, pageSize)

	// Call the service to get the blog list
	blogsResp, err := b.svc.GetAllBlogs(r.Context(), *pageReq)
	if err != nil {
		b.l.Error(ctx, "Error retrieving blog list: ", err)
		http.Error(w, "Failed to get blog list", http.StatusInternalServerError)
		return
	}

	// If no blogs found, return an empty response with a 204 No Content status
	if blogsResp == nil || len(blogsResp.Items) == 0 {
		b.l.Info(ctx, "No blogs found --")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set the response header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(blogsResp); err != nil {
		b.l.Errorf("Error encoding blog list to JSON: %s", err)
		http.Error(w, "Failed to encode blog list", http.StatusInternalServerError)
		return
	}

	b.l.Infof("Successfully retrieved %d blogs", len(blogsResp.Items))
}

// GetBlogByID retrieves a single blog by its ID.
func (b blogController) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	// Extract the blog ID from the URL path
	path := r.URL.Path
	b.l.Infof("Retrieving blog with ID: %s", path)
	writeLen, err := w.Write([]byte(path))
	if err != nil {
		b.l.Errorf("Error writing response: %s", err)
		return
	}
	b.l.Infof("Successfully wrote %d bytes", writeLen)

}

func (b blogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b blogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b blogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewBlogController(svc services.BlogService, l *logger.AppLogger) BlogController {
	return &blogController{
		svc: svc,
		l:   l,
	}
}
