package ezhuma

import (
	"errors"
	"net/http"

	"github.com/LeonColt/ez"
	"github.com/danielgtaylor/huma/v2"
)

func ParseHumaError(in *ez.Error) error {
	switch in.Code {
	case ez.ErrorCodeOk:
		return nil
	case ez.ErrorCodeCancelled:
		return huma.NewError(499, in.Message, in.Err)
	case ez.ErrorCodeUnknown:
		return huma.Error500InternalServerError(in.Message, in.Err)
	case ez.ErrorCodeInvalidArgument:
		return huma.Error400BadRequest(in.Message, in.Err)
	case ez.ErrorCodeDeadlineExceeded:
		return huma.NewError(http.StatusRequestTimeout, in.Message, in.Err)
	case ez.ErrorCodeNotFound:
		return huma.Error404NotFound(in.Message, in.Err)
	case ez.ErrorCodeConflict:
		return huma.Error409Conflict(in.Message, in.Err)
	case ez.ErrorCodeNotAuthorized:
		return huma.Error403Forbidden(in.Message, in.Err)
	case ez.ErrorCodeResourceExhausted:
		return huma.Error429TooManyRequests(in.Message, in.Err)
	case ez.ErrorCodeFailedPrecondition:
		return huma.Error412PreconditionFailed(in.Message, in.Err)
	case ez.ErrorCodeAborted:
		return huma.Error409Conflict(in.Message, in.Err)
	case ez.ErrorCodeOutOfRange:
		return huma.NewError(http.StatusRequestedRangeNotSatisfiable, in.Message, in.Err)
	case ez.ErrorCodeUnimplemented:
		return huma.Error501NotImplemented(in.Message, in.Err)
	case ez.ErrorCodeInternal:
		return huma.Error500InternalServerError(in.Message, in.Err)
	case ez.ErrorCodeUnavailable:
		return huma.Error503ServiceUnavailable(in.Message, in.Err)
	case ez.ErrorCodeDataLoss:
		return huma.Error500InternalServerError(in.Message, in.Err)
	case ez.ErrorCodeUnauthenticated:
		return huma.Error401Unauthorized(in.Message, in.Err)
	}
	return huma.Error500InternalServerError(in.Message, in.Err)
}

func HandleError(err error) error {
	if err == nil {
		return nil
	}
	var humaErr *ez.Error
	if errors.As(err, &humaErr) {
		return ParseHumaError(humaErr)
	} else {
		return huma.Error500InternalServerError(err.Error(), err)
	}
}
