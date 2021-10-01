package response

import "net/http"

// Response is a response from the assembly API. When using the HTTP API,
// API methods return *HTTPResponse values that implement Response.
type Response interface {
}

// HTTPResponse is a wrapped HTTP response from the SMSGlobal API with
// additional SMSGlobal specific response information parsed out. It
// implements Response.
type HTTPResponse struct {
	*http.Response
}

// newResponse create a new Response for the provided http.Response.
func NewResponse(r *http.Response) *HTTPResponse {
	return &HTTPResponse{Response: r}
}