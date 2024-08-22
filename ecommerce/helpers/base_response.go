package helpers

import (
	"github.com/go-playground/validator/v10"
)

type BaseResponse[T any] struct {
	Message string         `json:"message"`
	Data    T              `json:"data"`
	Errors  map[string]any `json:"errors,omitempty"`
}

func NewSuccessResponse[T any](data T) BaseResponse[T] {
	return BaseResponse[T]{

		Message: "Success",
		Data:    data,
	}
}

func NewErrorResponse[T any](code int, message string) BaseResponse[T] {
	return BaseResponse[T]{
		// Code:    code,
		Message: message,
	}
}

func NewValidationResponse[T any](code int, message string, err error) BaseResponse[T] {

	var validationErrors map[string]any = make(map[string]any)

	if errors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errors {
			validationErrors[e.Field()] = e.Tag()
		}
	}

	return BaseResponse[T]{
		Message: message,
		Errors:  validationErrors,
	}
}

/**
* 1  = Success
* 2  = Error
* 3  = Validation Error
* 4  = Unauthorized
* 5  = Pending
* 6  = Stock Empty
* 7  =
 */
