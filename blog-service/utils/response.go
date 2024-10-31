package utils

import (
	"encoding/json"
	"net/http"

	"blog-service/models/request"
	"blog-service/models/resp"
)

func RespondWithSuccess(w http.ResponseWriter, code int, payload interface{}) {
	response := resp.NewBaseResponse()
	response.Data = payload

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic("unable to encode response")
		return
	}
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	response := resp.NewBaseResponse()
	response.Success = false
	response.Error = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic("unable to encode response")
		return
	}
}

// CreatePageMetadata creates PageMetadata from request and total items
func CreatePageMetadata(req request.PaginationReq, totalItems int64) resp.PaginationResp {
	totalPages := (int(totalItems) + req.PageSize - 1) / req.PageSize
	return resp.PaginationResp{
		CurrentPage:    req.Page,
		PageSize:       req.PageSize,
		TotalPage:      totalPages,
		TotalItemCount: totalItems,
		HasMore:        req.Page < totalPages,
		IsFirstPage:    req.Page == 1,
		IsLastPage:     req.Page == totalPages,
	}
}
