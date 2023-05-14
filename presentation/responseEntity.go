package presentation

import (
	"encoding/json"
	codes "github.com/harryosmar/go-chi-base/errors"
	"net/http"
)

type ResponseContent struct {
	Status      bool                   `json:"status"`
	Message     string                 `json:"message"`
	ErrorCode   string                 `json:"error_code"`
	ErrorDetail string                 `json:"error_detail"`
	Data        interface{}            `json:"data"`
	MetaData    map[string]interface{} `json:"metadata"`
}

type ResponseEntity struct {
	StatusCode int               `json:"status"`
	Headers    map[string]string `json:"headers"`
	Content    ResponseContent   `json:"content"`
}

func NewResponseEntity() *ResponseEntity {
	return &ResponseEntity{Content: ResponseContent{}}
}

func (r *ResponseEntity) WithStatusCode(statusCode int) *ResponseEntity {
	r.StatusCode = statusCode
	return r
}

func (r *ResponseEntity) WithHeaders(headers map[string]string) *ResponseEntity {
	r.Headers = headers
	return r
}

func (r *ResponseEntity) WithData(data interface{}) *ResponseEntity {
	r.Content.Data = data
	return r
}

func (r *ResponseEntity) WithContentStatus(status bool) *ResponseEntity {
	r.Content.Status = status
	return r
}

func (r *ResponseEntity) WithMessage(message string) *ResponseEntity {
	r.Content.Message = message
	return r
}

func (r *ResponseEntity) WithErrorCode(errorCode string) *ResponseEntity {
	r.Content.ErrorCode = errorCode
	return r
}

func (r *ResponseEntity) WithErrorDetail(detail string) *ResponseEntity {
	r.Content.ErrorDetail = detail
	return r
}

func (r *ResponseEntity) WithMetaData(meta map[string]interface{}) *ResponseEntity {
	r.Content.MetaData = meta
	return r
}

func (r *ResponseEntity) Write(w http.ResponseWriter) {
	statusText := http.StatusText(r.StatusCode)
	if statusText == "" {
		ResponseErr(w, codes.NewCodeErrf(codes.ErrInvalidStatusCode).Paramf(r.StatusCode))
		return
	}

	w.WriteHeader(r.StatusCode)
	if r.Headers != nil {
		for k, v := range r.Headers {
			w.Header().Set(k, v)
		}
	}
	json.NewEncoder(w).Encode(r.Content)
}
