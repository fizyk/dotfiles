//+build mage

package main

import (
	"fmt"
	"github.com/fizyk/dotfiles/core"
	"github.com/fizyk/dotfiles/core/golang"
	"github.com/fizyk/dotfiles/targets/docker"
	"github.com/magefile/mage/mg"

	// mage:import go
	_ "github.com/fizyk/dotfiles/targets/golang"
	// mage:import go:check
	_ "github.com/fizyk/dotfiles/targets/golang/check"
	// mage:import docker
	_ "github.com/fizyk/dotfiles/targets/docker"
	"time"
)

// Hello is a sample hello mage target
func Hello() {
	fmt.Println("Ehlo there!")
}

// Installs all required programs
func Install() {
	defer core.MeasureTime(time.Now(), "install")
	mg.Deps(EnsureGolangCILint)
	mg.Deps(docker.Install)
}

// Installs golangci-lint
func EnsureGolangCILint() error {
	defer core.MeasureTime(time.Now(), "ensureGolangCILint")
	return golang.EnsureGolangCILint("1.41.0")
}
