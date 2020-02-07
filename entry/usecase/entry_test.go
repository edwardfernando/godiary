package usecase_test

import (
	"testing"

	"github.com/edwardfernando/godiary/entry/mocks"
	"github.com/edwardfernando/godiary/entry/usecase"
	"github.com/edwardfernando/godiary/model"
	"github.com/stretchr/testify/require"
)

func TestNewEntryUsecase(t *testing.T) {
	repository := new(mocks.Repository)

	entryUsecase := usecase.NewEntryUsecase(repository)
	require.NotNil(t, entryUsecase)

	repository.AssertExpectations(t)
}

func TestEntryUsecase_CreateEntry(t *testing.T) {
	testCases := []struct {
		name                  string
		expectedEntryResponse model.Entry
	}{
		{
			name: "Success - Entry saved to database",

		}
	}
}
