//go:build tools

package main

import (
	_ "github.com/conduitio/conduit-commons/paramgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "mvdan.cc/gofumpt"
)
