package entry

import "github.com/edwardfernando/godiary/model"

// Repository is interface for Entry repository
type Repository interface {
	CreateEntry(entry model.Entry) error
}
