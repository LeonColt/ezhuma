package ezhuma_test

import (
	"net/http"
	"testing"

	"github.com/LeonColt/ez"
	"github.com/LeonColt/ezhuma"
	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func getBuilder(code ez.ErrorCode, message string) *ez.Error {
	return &ez.Error{
		Code:    code,
		Message: message,
	}
}

func TestHandleError(t *testing.T) {
	{
		err := ezhuma.HandleError(nil)
		require.Nil(t, err)
	}

	{
		err := getBuilder(ez.ErrorCodeOk, "OK")
		humaErr := ezhuma.HandleError(err)
		require.Nil(t, humaErr)
	}
	{
		err := getBuilder(ez.ErrorCodeNotFound, "Item was not found")
		humaErr := ezhuma.HandleError(err)
		he := humaErr.(huma.StatusError)
		require.Equal(t, http.StatusNotFound, he.GetStatus())
		require.Equal(t, "Item was not found", he.Error())
	}
	{
		err := getBuilder(ez.ErrorCodeInternal, "an error occurred")
		humaErr := ezhuma.HandleError(err)
		he := humaErr.(huma.StatusError)
		require.Equal(t, http.StatusInternalServerError, he.GetStatus())
		require.Equal(t, "an error occurred", he.Error())
	}
	{
		err := getBuilder(ez.ErrorCodeUnauthenticated, "unauthenticated")
		humaErr := ezhuma.HandleError(err)
		he := humaErr.(huma.StatusError)
		require.Equal(t, http.StatusUnauthorized, he.GetStatus())
		require.Equal(t, "unauthenticated", he.Error())
	}
}
