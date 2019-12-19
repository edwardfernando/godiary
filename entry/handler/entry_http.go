package handler

import (
	"net/http"

	"github.com/edwardfernando/godiary/entry"
	"github.com/labstack/echo"
)

// EntryHTTPHandler is the handler for Entry endpoints
type EntryHTTPHandler struct {
	EntryUsecase entry.Usecase
}

// NewEntryHTTPHandler is the constructor for Entry handler
func NewEntryHTTPHandler(usecase entry.Usecase) *EntryHTTPHandler {
	return &EntryHTTPHandler{
		EntryUsecase: usecase,
	}
}

// PostEntry handles Entry creation
func (h *EntryHTTPHandler) PostEntry(e echo.Context) error {
	return e.JSON(http.StatusOK, nil)
}
