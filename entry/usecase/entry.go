package usecase

import (
	"github.com/edwardfernando/godiary/entry"
)

type entryUsecase struct {
	ChallengeRepo entry.Repository
}

// NewEntryUsecase initialises the usecase for Entry
func NewEntryUsecase(repository entry.Repository) entry.Usecase {
	return &entryUsecase{
		ChallengeRepo: repository,
	}
}
