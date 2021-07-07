//+build mage

package main

import (
	"fmt"
	"github.com/fizyk/dotfiles/core/golang"
	"github.com/fizyk/dotfiles/targets/docker"
	"github.com/fizyk/dotfiles/targets/docker/lazydocker"
	"github.com/fizyk/magex/magefiles/mage"
	magexTime "github.com/fizyk/magex/time"
	"github.com/magefile/mage/mg"
	"runtime"
	"strings"
	"time"

	// mage:import go
	_ "github.com/fizyk/magex/magefiles/golang"
	// mage:import go:check
	_ "github.com/fizyk/magex/magefiles/golang/check"
	// mage:import docker
	_ "github.com/fizyk/dotfiles/targets/docker"
	// mage:import docker:lazy
	_ "github.com/fizyk/dotfiles/targets/docker/lazydocker"
	// mage:import mage
	_ "github.com/fizyk/magex/magefiles/mage"
)

// Hello is a sample hello mage target
func Hello() error {
	fmt.Println("Ehlo there!")
	fmt.Printf("You're running: %v at %v\n", strings.Title(runtime.GOOS), runtime.GOARCH)
	return nil
}

// Installs all required programs
func Install() {
	defer magexTime.MeasureTime(time.Now(), "install")
	mg.Deps(EnsureGolangCILint, docker.Install, mage.Install)
	//mg.Deps(docker.Install)
	mg.Deps(lazydocker.Install)
}

// Installs golangci-lint
func EnsureGolangCILint() error {
	defer magexTime.MeasureTime(time.Now(), "ensureGolangCILint")
	return golang.EnsureGolangCILint("1.41.0")
}
