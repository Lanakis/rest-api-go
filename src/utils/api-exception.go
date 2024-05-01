package utils

import (
	"encoding/json"
	"fmt"
)

type ErrorFields map[string]string
type ErrorParams map[string]string

type AppError struct {
	Err              error       `json:"-"`
	Message          string      `json:"message,omitempty"`
	DeveloperMessage string      `json:"developer_message,omitempty"`
	Code             string      `json:"code,omitempty"`
	Fields           ErrorFields `json:"fields"`
	Params           ErrorParams `json:"params"`
}

func (e *AppError) WithFields(fields ErrorFields) {
	e.Fields = fields
}
func (e *AppError) WithParams(params ErrorParams) {
	e.Params = params
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}
func BadRequestError(message, developerMessage string) *AppError {
	return NewAppError(fmt.Errorf(message), message, developerMessage, "US-000001")
}
