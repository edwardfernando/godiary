package usecase

import (
	"github.com/edwardfernando/godiary/entry"
	"github.com/edwardfernando/godiary/errorwrapper"
	"github.com/edwardfernando/godiary/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (e *entryUsecase) PostEntry(request entry.PostEntryRequest) *errorwrapper.ErrWrapper {
	id, err := uuid.NewRandom()
	if err != nil {
		return &errorwrapper.ErrWrapper{
			ErrType: entry.GodiaryInternalServerError,
			Err:     errors.Wrapf(err, "failed to generate new uuid"),
		}
	}

	model := model.Entry{
		ID:    id,
		Title: request.Title,
		Body:  request.Body,
	}

	err = e.ChallengeRepo.CreateEntry(model)
	if err != nil {
		return &errorwrapper.ErrWrapper{
			ErrType: entry.GodiaryInternalServerError,
			Err:     errors.Wrapf(err, "entry usercase: failed to create an Entry"),
		}
	}

	return nil
}
