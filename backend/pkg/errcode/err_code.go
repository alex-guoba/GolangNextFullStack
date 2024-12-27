package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int      `json:"code" example:"100000"`
	Msg     string   `json:"msg" example:"collection not found"`
	Details []string `json:"details,omitempty" example:"collection format error,invalid parameter"`
	Status  int      `json:"-"`
}

var codes = map[int]string{}

func NewError(code int, msg string, status int) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("error code %d already exist", code))
	}
	codes[code] = msg
	if status == 0 {
		status = http.StatusInternalServerError
	}
	return &Error{Code: code, Msg: msg, Status: status}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d, detail: %s", e.Code, e.Msg)
}

func (e *Error) Msgf(args []any) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = []string{}
	newError.Details = append(newError.Details, details...)

	return &newError
}
