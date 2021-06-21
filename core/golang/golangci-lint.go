package golang

import (
	"fmt"
	"github.com/fizyk/dotfiles/core"
	"github.com/fizyk/dotfiles/core/command"
	"github.com/hashicorp/go-version"
	"github.com/magefile/mage/sh"
	"os"
	"os/exec"
	"regexp"
	"time"
)

const versionRegexp string = "([\\d]+\\.[\\d]+(\\.[\\d]+)?)"

// Parse golangci-lint version
func golangCIVersion() (*version.Version, error) {
	versionOutput, error := sh.Output("golangci-lint", "version")
	if error != nil {
		return nil, error
	}
	versionRegexp := regexp.MustCompile(versionRegexp)
	matchedVersion := versionRegexp.FindString(versionOutput)
	return version.NewVersion(matchedVersion)
}

// Installs golangci-lint
func EnsureGolangCILint(newVersion string) error {
	defer core.MeasureTime(time.Now(), "ensureGolangCILint")
	installVersion, _ := version.NewVersion(newVersion)
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
	command.PipeCommands(c1, c2)

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
