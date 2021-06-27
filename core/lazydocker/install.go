package lazydocker

import (
	"errors"
	"fmt"
	"github.com/fizyk/dotfiles/core/github"
	"github.com/fizyk/dotfiles/core/http"
	"github.com/hashicorp/go-version"
	"github.com/magefile/mage/sh"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var architectureMap map[string]string = map[string]string{
	"amd64": "x86_64",
}

var extensionMap map[string]string = map[string]string{
	"linux":   "tar.gz",
	"windows": ".zip",
}

const (
	tarFile string = "lazydocker.tar.gz"
	binPath string = "/usr/local/bin"
	exec    string = "lazydocker"
)

// filter filters asset name to correctly choose one for given Operating System and Architecture
func filter(downloadName string) bool {
	architecture := architectureMap[runtime.GOARCH]
	searchFor := fmt.Sprintf("%s_%s.tar.gz", strings.Title(runtime.GOOS), architecture)
	return strings.Contains(downloadName, searchFor)
}

const versionRegexpPattern string = "([\\d]+\\.[\\d]+(\\.[\\d]+)?)"

// lazydockerVersion detects currently installed lazydocker version.
func lazydockerVersion() (*version.Version, error) {
	versionOutput, err := sh.Output(fmt.Sprintf("%s/%s", binPath, exec), "--version")
	if err != nil {
		return nil, err
	}
	versionLineRegexp := regexp.MustCompile(fmt.Sprintf("Version: %s", versionRegexpPattern))
	versionLine := versionLineRegexp.FindString(versionOutput)
	versionRegexp := regexp.MustCompile(versionRegexpPattern)
	fmt.Printf("%v\n", versionRegexp.FindString(versionLine))

	return version.NewVersion(versionRegexp.FindString(versionLine))
}

// isInstalled checks whether lazydocker is installed within system
func isInstalled() bool {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", binPath, exec)); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// Install performs all steps that lazydocker's shell script would perform.
// Skips if the lazydocker is already installed in newest version
func Install() error {
	downloadURL, version, err := github.Latest("jesseduffield", "lazydocker", filter)
	if err != nil {
		return err
	}
	if isInstalled() {
		installedVersion, err := lazydockerVersion()
		if err != nil {
			return err
		}
		if installedVersion.Equal(version) {
			fmt.Println("Newest version already installed.")
			return nil
		}
		fmt.Printf("lazydocker %s is installed, will update to: %s\n", installedVersion.String(), version.String())
	} else {
		fmt.Printf("Will install %s version of lazydocker", version.String())
	}
	if err := http.DownloadFile(downloadURL, tarFile); err != nil {
		return err
	}
	if err := sh.Run("tar", "zxvf", tarFile, exec); err != nil {
		return err
	}
	if err := sh.Run("sudo", "mv", "-f", exec, binPath); err != nil {
		return err
	}
	os.Remove(tarFile)
	installedVersion, err := lazydockerVersion()
	if err != nil {
		return err
	}
	fmt.Printf("Lazydocker is installed at version %s", installedVersion.String())
	return nil
}
