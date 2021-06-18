//+build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"io"
	"os"
	"os/exec"
)

// Hello is a sample hello mage target
func Hello() {
	fmt.Println("Ehlo there!")
}

// Installs golangci-lint
func EnsureGolangCILint() error {
	// TODO: Pass version as a parameter
	// TODO: if version is already installed, do not install, print it's installed
	// TODO: Check what's the newest version if not passed and install newest version then (which check if installed)
	// TODO: measure time of the installation
	fmt.Println("Start")
	sh.RunV("pwd")
	//gopath, _ := sh.Output("go", "env", "GOPATH")
	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)
	// TODO: Have structs for commands and args
	c1 := exec.Command("curl", "-sSfL", "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh")
	c2 := exec.Command("sh", "-s", "--", "-b", fmt.Sprintf("%s/bin", gopath), "v1.41.0")
	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r
	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	fmt.Println("End")
	version, error := sh.Output("golangci-lint", "version")
	if error != nil {
		return error
	}
	// TODO: If updated, print it's updated
	fmt.Println("Installed golangci-lint!")
	fmt.Println(version)
	return nil
}
