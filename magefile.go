//+build mage

package main

import (
	"fmt"
	"github.com/fizyk/dotfiles/core/golang"
)

// Hello is a sample hello mage target
func Hello() {
	fmt.Println("Ehlo there!")
}

// Installs golangci-lint
func EnsureGolangCILint() error {
	return golang.EnsureGolangCILint("1.41.0")
}
