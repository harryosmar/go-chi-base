package presentation

import (
	"encoding/json"
	"github.com/harryosmar/go-chi-base/core/errors"
	"net/http"
)

type ResponseContent struct {
	Status      bool                   `json:"status"`
	Message     string                 `json:"message,omitempty"`
	ErrorCode   string                 `json:"error_code,omitempty"`
	ErrorDetail string                 `json:"error_detail,omitempty"`
	Data        interface{}            `json:"data"`
	MetaData    map[string]interface{} `json:"metadata,omitempty"`
}

type ResponseEntity struct {
	StatusCode int               `json:"status"`
	Headers    map[string]string `json:"headers"`
	Content    ResponseContent   `json:"content"`
}

func NewResponseEntity() *ResponseEntity {
	return &ResponseEntity{
		Content: ResponseContent{
			MetaData: map[string]interface{}{},
		},
		Headers: map[string]string{},
	}
}

func (r *ResponseEntity) WithStatusCode(statusCode int) *ResponseEntity {
	r.StatusCode = statusCode
	return r
}

func (r *ResponseEntity) WithHeaders(headers map[string]string) *ResponseEntity {
	for k, v := range headers {
		r.Headers[k] = v
	}
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
	for k, v := range meta {
		r.Content.MetaData[k] = v
	}
	return r
}

type Paginator struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func (r *ResponseEntity) WithPaginator(paginator Paginator) *ResponseEntity {
	r.Content.MetaData["limit"] = paginator.Limit
	r.Content.MetaData["offset"] = paginator.Offset
	r.Content.MetaData["total"] = paginator.Total

	return r
}

func (r *ResponseEntity) WriteJson(w http.ResponseWriter) {
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
