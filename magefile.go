//+build mage

package main

import (
	"fmt"
	"github.com/fizyk/dotfiles/core"
	"github.com/fizyk/dotfiles/core/golang"
	"github.com/fizyk/dotfiles/targets/docker"
	"github.com/fizyk/dotfiles/targets/docker/lazydocker"
	"github.com/magefile/mage/mg"
	"runtime"
	"strings"

	// mage:import go
	_ "github.com/fizyk/dotfiles/targets/golang"
	// mage:import go:check
	_ "github.com/fizyk/dotfiles/targets/golang/check"
	// mage:import docker
	_ "github.com/fizyk/dotfiles/targets/docker"
	// mage:import docker:lazy
	_ "github.com/fizyk/dotfiles/targets/docker/lazydocker"
	"time"
)

// Hello is a sample hello mage target
func Hello() error {
	fmt.Println("Ehlo there!")
	fmt.Printf("You're running: %v at %v\n", strings.Title(runtime.GOOS), runtime.GOARCH)
	return nil
}

// Installs all required programs
func Install() {
	defer core.MeasureTime(time.Now(), "install")
	mg.Deps(EnsureGolangCILint, docker.Install)
	//mg.Deps(docker.Install)
	mg.Deps(lazydocker.Install)
}

// Installs golangci-lint
func EnsureGolangCILint() error {
	defer core.MeasureTime(time.Now(), "ensureGolangCILint")
	return golang.EnsureGolangCILint("1.41.0")
}
