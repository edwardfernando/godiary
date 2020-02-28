package usecase_test

import (
	"errors"
	"testing"

	"github.com/edwardfernando/godiary/entry"
	"github.com/edwardfernando/godiary/entry/mocks"
	"github.com/edwardfernando/godiary/entry/usecase"
	"github.com/edwardfernando/godiary/errorwrapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewEntryUsecase(t *testing.T) {
	repository := new(mocks.Repository)

	entryUsecase := usecase.NewEntryUsecase(repository)
	require.NotNil(t, entryUsecase)

	repository.AssertExpectations(t)
}

func TestEntryUsecase_CreateEntry(t *testing.T) {
	entryRequest := &entry.PostEntryRequest{
		Title: "my random test",
		Body:  "my body test",
	}

	testCases := []struct {
		name                           string
		entryRequest                   *entry.PostEntryRequest
		expectCreateEntryToDbFail      bool
		expectedCreateEntryResponseErr error
		expectedFailureResponse        *errorwrapper.ErrWrapper
	}{
		{
			name:                           "Success - Entry saved to database",
			entryRequest:                   entryRequest,
			expectCreateEntryToDbFail:      false,
			expectedCreateEntryResponseErr: nil,
			expectedFailureResponse:        nil,
		},
		{
			name:                           "Failed to save an Entry to database",
			entryRequest:                   entryRequest,
			expectCreateEntryToDbFail:      true,
			expectedCreateEntryResponseErr: errors.New("whooops"),
			expectedFailureResponse: &errorwrapper.ErrWrapper{
				ErrType: entry.GodiaryInternalServerError,
				Err:     errors.New("entry usercase: failed to create an Entry: whooops"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repository := new(mocks.Repository)
			repository.On("CreateEntry", mock.Anything).Return(tc.expectedCreateEntryResponseErr)

			u := usecase.NewEntryUsecase(repository)
			err := u.PostEntry(*tc.entryRequest)

			if tc.expectCreateEntryToDbFail {
				assert.Equal(t, tc.expectedFailureResponse.Err.Error(), err.Err.Error())
			} else {
				assert.Nil(t, err)
			}

			repository.AssertExpectations(t)
		})
	}
}
