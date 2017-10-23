// Package golang implements the "golang" runtime.
package golang

import (
	"runtime"
	"strings"

	"github.com/apex/apex/function"
	"github.com/apex/apex/plugins/nodejs"
)

func init() {
	function.RegisterPlugin("golang", &Plugin{})
}

const (
	// Runtime for inference.
	Runtime = "golang"
)

// Plugin implementation.
type Plugin struct{}

// Open adds the shim and golang defaults.
func (p *Plugin) Open(fn *function.Function) error {
	if !strings.HasPrefix(fn.Runtime, "golang") {
		return nil
	}

	if fn.Hooks.Build == "" {
		if runtime.GOOS == "windows" {
			fn.Hooks.Build = "set GOOS=linux& set GOARCH=amd64& go build -o main"
		} else {
			fn.Hooks.Build = "GOOS=linux GOARCH=amd64 go build -o main *.go"
		}
	}

	fn.Shim = true
	fn.Runtime = nodejs.Runtime

	if fn.Handler == "" {
		fn.Handler = "index.handle"
	}

	if fn.Hooks.Clean == "" {
		if runtime.GOOS == "windows" {
			fn.Hooks.Clean = "del /F main"
		} else {
			fn.Hooks.Clean = "rm -f main"
		}
	}

	return nil
}
