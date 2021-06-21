//+build mage

package main

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/magefile/mage/sh"
	"io"
	"os"
	"os/exec"
	"regexp"
	"time"
)

const versionRegexp string = "([\\d]+\\.[\\d]+(\\.[\\d]+)?)"

// Hello is a sample hello mage target
func Hello() {
	fmt.Println("Ehlo there!")
}

// Installs golangci-lint
func EnsureGolangCILint() error {
	defer measureTime(time.Now(), "ensureGolangCILint")
	installVersion, _ := version.NewVersion("v1.41.0")
	fmt.Println("Start")
	gopath := os.Getenv("GOPATH")
	existingVersion, error := golangCIVersion()
	if error == nil {
		if existingVersion.Equal(installVersion) {
			fmt.Printf("GolangCI-Lint %s already installed\n", installVersion.String())
			return nil
		}
	}

	c1 := exec.Command("curl", "-sSfL", "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh")
	c2 := exec.Command("sh", "-s", "--", "-b", fmt.Sprintf("%s/bin", gopath), fmt.Sprintf("v%s", installVersion.String()))
	runCommandAndPipe(c1, c2)


	if newlyInstalledVersion, error := golangCIVersion(); error != nil {
		return error
	} else {
		fmt.Print("GolangCI-Lint got ")
		if existingVersion == nil {
			fmt.Printf("installed at version %s\n", newlyInstalledVersion.String())
		} else if existingVersion.GreaterThan(newlyInstalledVersion) {
			fmt.Printf("downgraded from %s to %s\n", existingVersion.String(), newlyInstalledVersion.String())
		} else {
			fmt.Printf("upgraded from %s to %s\n", existingVersion.String(), newlyInstalledVersion.String())
		}
	}
	return nil
}

func measureTime(start time.Time, banner string) {
	end := time.Now()
	fmt.Printf("%s took %.2f seconds\n", banner, end.Sub(start).Seconds())
}

func golangCIVersion() (*version.Version, error) {
	versionOutput, error := sh.Output("golangci-lint", "version")
	if error != nil {
		return nil, error
	}
	versionRegexp := regexp.MustCompile(versionRegexp)
	matchedVersion := versionRegexp.FindString(versionOutput)
	return version.NewVersion(matchedVersion)
}

// runCommandAndPipe runs first command and pipes it's output to second one
func runCommandAndPipe(mainCommand *exec.Cmd, pipeToCommand *exec.Cmd) {
	fmt.Printf("%s | %s \n", mainCommand.String(), pipeToCommand.String())
	read, write := io.Pipe()
	// Main command will write
	mainCommand.Stdout = write

	// pipeCommand will read
	pipeToCommand.Stdin = read
	mainCommand.Start()
	pipeToCommand.Start()
	// Wait for the mainCommand to finish
	mainCommand.Wait()
	// Close write io
	write.Close()
	// Wait for the pipeCommand to finish
	pipeToCommand.Wait()
}
