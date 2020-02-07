package handler

import (
	"net/http"

	"github.com/edwardfernando/godiary/contract"
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
	var req contract.EntryRequetContext
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, &contract.Response{
			Data: make(map[string]interface{}),
			Errors: []contract.ResponseErr{
				{
					Code:            "invalid requet",
					MessageTitle:    http.StatusText(http.StatusUnprocessableEntity),
					Message:         "Invalid entry request",
					MessageSeverity: "error",
				},
			},
		})
	}

	err := h.EntryUsecase.PostEntry(entry.PostEntryRequest{
		Title: req.Title,
		Body:  req.Body,
	})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, &contract.Response{
			Data: make(map[string]interface{}),
			Errors: []contract.ResponseErr{
				{
					Code:            "bad request",
					MessageTitle:    http.StatusText(http.StatusBadRequest),
					Message:         "Something bad happened",
					MessageSeverity: "error",
				},
			},
		})
	}

	return e.JSON(http.StatusOK, req)
}
