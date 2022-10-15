//go:build tools

package main

import (
	_ "github.com/golang/mock/mockgen"
	_ "golang.org/x/tools/cmd/godoc"
	_ "mvdan.cc/gofumpt"
)
