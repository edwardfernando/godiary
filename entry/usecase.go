package entry

import (
	"github.com/edwardfernando/godiary/errorwrapper"
	"github.com/google/uuid"
)

// Usecase is interface for Entry usecase
type Usecase interface {
	PostEntry(request PostEntryRequest) *errorwrapper.ErrWrapper
}

// PostEntryRequest defines the required parameter to create a new entry
type PostEntryRequest struct {
	Title string
	Body  string
}

// PostEntryResponse defines the returned result
type PostEntryResponse struct {
	ID uuid.UUID
}
