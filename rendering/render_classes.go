package rendering

import "net/http"

type RequestContext struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}
