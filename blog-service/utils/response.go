package utils

import (
	"blog-service/models"
	"encoding/json"
	"net/http"
)

func RespondWithSuccess(w http.ResponseWriter, code int, payload interface{}) {
	response := models.NewBaseResponse()
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
	response := models.NewBaseResponse()
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
func CreatePageMetadata(req models.BaseRequest, totalItems int64) models.PageMetaData {
	totalPages := (int(totalItems) + req.PageSize - 1) / req.PageSize

	return models.PageMetaData{
		CurrentPage: req.Page,
		PageSize:    req.PageSize,
		TotalCount:  totalItems,
		TotalPages:  totalPages,
		HasNext:     req.Page < totalPages,
	}
}
