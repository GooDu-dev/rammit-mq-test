package types

import (
	"net/http"

	"github.com/GooDu-Dev/function-parser-go/pkg/inventory"
)

type Package struct {
	id          string
	process     inventory.Process
	softTimeout int
	retry       int
	timeout     int
	retryable   bool
	response    http.Response
}

func NewPackage(id string, process inventory.Process, softTimeout int, retry int, timeout int, retryable bool) inventory.Package {
	return &Package{
		id:          id,
		process:     process,
		softTimeout: softTimeout,
		retry:       retry,
		timeout:     timeout,
		retryable:   retryable,
	}
}

func (p *Package) ID() string {
	return p.id
}

func (p *Package) Process() inventory.Process {
	return p.process
}

func (p *Package) SoftTimeout() int {
	return p.softTimeout
}

func (p *Package) Retry() int {
	return p.retry
}

func (p *Package) Timeout() int {
	return p.timeout
}

func (p *Package) Retryable() bool {
	return p.retryable
}

func (p *Package) Response() http.Response {
	return p.response
}
