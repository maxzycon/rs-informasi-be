package errors

import (
	"errors"

	"github.com/maxzycon/rs-informasi-be/pkg/dto"
	"gorm.io/gorm"
)

var (
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("data not found")
	ErrNotFoundGorm   = gorm.ErrRecordNotFound

	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrStatusAlreadyCheckIn      = errors.New("invalid status already checkin")
	ErrStatusAlreadyCheckOut     = errors.New("invalid status already check out")
	ErrInvalidPassword           = errors.New("invalid password")
	ErrInvalidUsername           = errors.New("invalid username")
	ErrUsernameUsed              = errors.New("username already used")
	ErrBlockedAccount            = errors.New("account blocked")
	ErrUnparseableRequestBody    = errors.New("unpardonable request body error")
	ErrAuthTokenExpired          = errors.New("auth token expired")
	ErrNoPermission              = errors.New("doesn't have permission")
	ErrPasswordValidation        = errors.New("password must be at least 8 characters consists of uppercase, lowercase, and a digit")
	ErrInvalidQty                = errors.New("invalid qty, measure available qty >= qty input")
	ErrInvalidCheckoutQty        = errors.New("invalid qty, measure qty check out >= qty input")
)

var errorMapping = map[error]dto.ErrorResponse{
	// general errors
	ErrBadRequest:             {HTTPCode: 400, Code: 1001, Message: ErrBadRequest.Error()},
	ErrInternalServer:         {HTTPCode: 500, Code: 1002, Message: ErrInternalServer.Error()},
	ErrUnauthorized:           {HTTPCode: 401, Code: 1003, Message: ErrUnauthorized.Error()},
	ErrUnparseableRequestBody: {HTTPCode: 400, Code: 1004, Message: ErrUnparseableRequestBody.Error()},
	ErrNotFound:               {HTTPCode: 404, Code: 1005, Message: ErrNotFound.Error()},
	ErrStatusAlreadyCheckIn:   {HTTPCode: 404, Code: 1005, Message: ErrStatusAlreadyCheckIn.Error()},
	ErrStatusAlreadyCheckOut:  {HTTPCode: 404, Code: 1005, Message: ErrStatusAlreadyCheckOut.Error()},
	ErrNotFoundGorm:           {HTTPCode: 404, Code: 1006, Message: ErrNotFoundGorm.Error()},
	ErrForbidden:              {HTTPCode: 403, Code: 1007, Message: ErrForbidden.Error()},

	// business logic errors
	ErrInvalidUsernameOrPassword: {HTTPCode: 400, Code: 2001, Message: ErrInvalidUsernameOrPassword.Error()},
	ErrInvalidPassword:           {HTTPCode: 400, Code: 2001, Message: ErrInvalidPassword.Error()},
	ErrInvalidUsername:           {HTTPCode: 400, Code: 2001, Message: ErrInvalidUsername.Error()},
	ErrUsernameUsed:              {HTTPCode: 400, Code: 2001, Message: ErrUsernameUsed.Error()},
	ErrBlockedAccount:            {HTTPCode: 400, Code: 2001, Message: ErrBlockedAccount.Error()},

	ErrAuthTokenExpired:   {HTTPCode: 403, Code: 2002, Message: ErrAuthTokenExpired.Error()},
	ErrNoPermission:       {HTTPCode: 400, Code: 2003, Message: ErrNoPermission.Error()},
	ErrPasswordValidation: {HTTPCode: 400, Code: 2005, Message: ErrPasswordValidation.Error()},
	ErrInvalidQty:         {HTTPCode: 400, Code: 2006, Message: ErrInvalidQty.Error()},
	ErrInvalidCheckoutQty: {HTTPCode: 400, Code: 2006, Message: ErrInvalidCheckoutQty.Error()},
}

func GetErrorResponse(er error) (errRes dto.ErrorResponse) {
	errRes, found := errorMapping[er]
	if !found {
		errRes = errorMapping[ErrInternalServer]
	}
	return
}
