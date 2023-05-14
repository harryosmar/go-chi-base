package codes

import (
	"fmt"
)

type CodeErr int

//go:generate mockgen -destination=mocks/mock_CodeErrEntity.go -package=mocks . CodeErrEntity
type CodeErrEntity interface {
	Error() string
	Code() string
	HttpStatus() int
	String() string
	Detail() string
}

type codeErrEntity struct {
	code    string
	message string
	status  int
}

const (
	defaultErrMessage = "Terjadi kesalahan sistem, mohon cobalah beberapa saat lagi."
)

var codeErrMap = map[CodeErr]codeErrEntity{
	ErrGeneral:           {code: "ERRSVC5000", status: 500, message: defaultErrMessage},
	ErrInvalidStatusCode: {code: "ERRSVC5001", status: 500, message: "invalid status code %d"},
	ErrValidation:        {code: "ERRSVC4000", status: 400, message: "validation error"},
}

const (
	ErrGeneral CodeErr = iota
	ErrInvalidStatusCode
	ErrValidation
)

func (c CodeErr) Error() string {
	return fmt.Sprintf("[%s] %s", c.Code(), c.String())
}

func (c CodeErr) Code() string {
	return codeErrMap[c].code
}

func (c CodeErr) HttpStatus() int {
	return codeErrMap[c].status
}

func (c CodeErr) String() string {
	return codeErrMap[c].message
}

func (c CodeErr) Detail() string {
	return ""
}

type codeErrf struct {
	err    CodeErr
	detail string
	params []interface{}
}

func NewCodeErrf(err CodeErr) *codeErrf {
	return &codeErrf{err: err}
}

func (c *codeErrf) SetDetail(detail string) *codeErrf {
	c.detail = detail
	return c
}

func (c *codeErrf) Paramf(a ...interface{}) *codeErrf {
	c.params = a
	return c
}

func (c *codeErrf) Error() string {
	return c.err.Error()
}

func (c *codeErrf) Code() string {
	return c.err.Code()
}

func (c *codeErrf) HttpStatus() int {
	return c.err.HttpStatus()
}

func (c *codeErrf) String() string {
	if c.params == nil {
		return c.err.String()
	}
	return fmt.Sprintf("%s", fmt.Sprintf(c.err.String(), c.params...))
}

func (c *codeErrf) Detail() string {
	return c.detail
}
