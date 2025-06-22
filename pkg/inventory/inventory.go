package inventory

import (
	"net/http"
)

type Package interface {
	ID() string
	Process() Process
	SoftTimeout() int
	Retry() int
	Timeout() int
	Retryable() bool
	Response() http.Response
}

type Process interface {
	Function() string
	Params() map[string]interface{}
	Return() interface{}
	JSON() []byte
}
