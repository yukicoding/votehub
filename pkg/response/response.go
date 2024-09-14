package response

import (
	"encoding/json"
	"net/http"
)

// Response 是一个通用的 API 响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	headers http.Header `json:"-"`
}

func New(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// SetHeader 设置响应头
func (r *Response) SetHeader(key, value string) *Response {
	r.headers.Set(key, value)
	return r
}

// Write Json 写入 http.ResponseWriter
func (r *Response) WriteJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	// 应用自定义头部
	for key, values := range r.headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}

// SetCommonHeaders 设置通用的响应头
func SetCommonHeaders(r *Response) *Response {
	return r.
		SetHeader("X-Content-Type-Options", "nosniff").
		SetHeader("X-Frame-Options", "DENY").
		SetHeader("X-XSS-Protection", "1; mode=block").
		SetHeader("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
}

func Success(data interface{}) *Response {
	return New(http.StatusOK, "success", data)
}

func Error(code int, message string) *Response {
	return New(code, message, nil)
}

func ServerError(err error) *Response {
	return New(http.StatusInternalServerError, err.Error(), nil)
}
func BadRequest(message string) *Response {
	return New(http.StatusBadRequest, message, nil)
}

func Unauthorized(message string) *Response {
	return New(http.StatusUnauthorized, message, nil)
}

func NotFound(message string) *Response {
	return New(http.StatusNotFound, message, nil)
}
