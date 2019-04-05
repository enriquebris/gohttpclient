package gohttpclient

import (
	"fmt"
	"net/http"
)

// ************************************************************************************************
// ** HTTPResponse
// ************************************************************************************************

type HTTPResponse interface {
	GetBody() string
	GetStatus() string
	GetStatusCode() int
	GetHeaders() map[string][]string
}

// ************************************************************************************************
// ** DefaultHTTPResponse
// ************************************************************************************************

type DefaultHTTPResponse struct {
	body       string
	statusCode int
	status     string
	headers    map[string][]string
}

func NewDefaultHTTPResponse(body string, statusCode int, status string, headers map[string][]string) *DefaultHTTPResponse {
	ret := &DefaultHTTPResponse{}
	ret.initialize(body, statusCode, status, headers)

	return ret
}

func (st *DefaultHTTPResponse) initialize(body string, statusCode int, status string, headers map[string][]string) {
	st.body = body
	st.statusCode = statusCode
	st.status = status
	st.headers = headers
}

func (st *DefaultHTTPResponse) GetBody() string {
	return st.body
}

func (st *DefaultHTTPResponse) GetStatus() string {
	return st.status
}

func (st *DefaultHTTPResponse) GetStatusCode() int {
	return st.statusCode
}

func (st *DefaultHTTPResponse) GetHeaders() map[string][]string {
	return st.headers
}

// ************************************************************************************************
// ** ResponseWriter
// ************************************************************************************************
// ** HTTP response writer
// ************************************************************************************************

type ResponseWriter interface {
	SetStatusCode(statusCode int)
	AddHeader(key string, value string)
	Print(message string)
	Printf(format string, a ...interface{})
}

// ************************************************************************************************
// ** DefaultResponseWriter
// ************************************************************************************************
// ** ResponseWriter implementation for native GO http
// ************************************************************************************************

func NewDefaultResponseWriter(w http.ResponseWriter) *DefaultResponseWriter {
	return &DefaultResponseWriter{
		w: w,
	}
}

type DefaultResponseWriter struct {
	w http.ResponseWriter
}

func (st *DefaultResponseWriter) SetStatusCode(statusCode int) {
	st.w.WriteHeader(statusCode)
}

func (st *DefaultResponseWriter) AddHeader(key string, value string) {
	st.w.Header().Add(key, value)
}

func (st *DefaultResponseWriter) Print(message string) {
	fmt.Fprint(st.w, message)
}

func (st *DefaultResponseWriter) Printf(format string, a ...interface{}) {
	fmt.Fprintf(st.w, format, a)
}
