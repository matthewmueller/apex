// Package golang implements the "golang" runtime.
package golang

import (
	"strings"

	"github.com/matthewmueller/apex/function"
)

func init() {
	function.RegisterPlugin("golang", &Plugin{})
}

const (
	// Runtime for inference.
	Runtime = "provided.al2"
)

// Plugin implementation.
type Plugin struct{}

// Open adds the shim and golang defaults.
func (p *Plugin) Open(fn *function.Function) error {
	if !strings.HasPrefix(fn.Runtime, "provided.al2") {
		return nil
	}

	if fn.Runtime == "golang" {
		fn.Runtime = Runtime
	}

	if fn.Hooks.Build == "" {
		fn.Hooks.Build = "GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap *.go"
	}

	if fn.Handler == "" {
		fn.Handler = "bootstrap"
	}

	if fn.Hooks.Clean == "" {
		fn.Hooks.Clean = "rm -f bootstrap"
	}

	return nil
}
