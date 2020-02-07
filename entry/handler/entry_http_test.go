package handler_test

import (
	"github.com/edwardfernando/godiary/entry/handler"
	"github.com/edwardfernando/godiary/entry/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewEntryHTTPHandler(t *testing.T) {
	usecase := new(mocks.Usecase)
	handler := handler.NewEntryHTTPHandler(usecase)

	require.NotNil(t, handler)

	usecase.AssertExpectations(t)
}
