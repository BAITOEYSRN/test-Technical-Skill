package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomError struct {
	Err             error
	StatusCode      int
	Code            uuid.UUID
	Message         string
	ResponseMessage string
}

type BaseResponseData struct {
	Code              uuid.UUID `json:"code"`
	Message           string    `json:"message"`
	ResponseMessageEN string    `json:"response_message_en"`
	ResponseMessageTH string    `json:"response_message_th"`
}

type ResponseData struct {
	Code            uuid.UUID   `json:"code"`
	Message         string      `json:"message"`
	ResponseMessage string      `json:"response_message"`
	Result          interface{} `json:"result"`
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v (status code: %d)", e.Message, e.Err, e.StatusCode)
	}
	return fmt.Sprintf("%s (status code: %d)", e.Message, e.StatusCode)
}

func Unwrap(err error) *CustomError {
	if customErr, ok := err.(*CustomError); ok {
		return customErr
	}
	return nil
}

func Wrap(err error, statusCode int, code uuid.UUID, responseMessage string) error {
	return &CustomError{
		Message:         err.Error(),
		Err:             err,
		StatusCode:      statusCode,
		Code:            code,
		ResponseMessage: responseMessage,
	}
}

func ResponseErrorJsonWithCode(ctx *gin.Context, err error) error {
	customErr := Unwrap(err)
	return ResponseJsonWithCode(ctx,
		customErr.StatusCode,
		customErr.Code,
		customErr.Message,
		customErr.ResponseMessage,
		nil,
	)
}

func ResponseJsonWithCode(ctx *gin.Context, statusCode int, code uuid.UUID, message string, responseMessage string, responseObject interface{}) error {
	if responseObject == nil {
		responseObject = struct{}{}
	}
	result := ResponseData{
		Code:            code,
		Message:         message,
		ResponseMessage: responseMessage,
		Result:          responseObject,
	}
	ctx.JSON(statusCode, result)
	return nil
}
