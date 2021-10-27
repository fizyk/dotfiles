package golang

import (
	"errors"
	"fmt"
	"github.com/fizyk/magex/file"
	"github.com/hashicorp/go-version"
	"github.com/magefile/mage/sh"
	"regexp"
)

const (
	golangPath string = "/usr/local/go"
	golangExecPath string = "/usr/local/go/bin/go"
	versionPatternGolang string = "go([\\d]+\\.[\\d]+\\.[\\d]+?)"
)

func golangVersion(path string) (*version.Version, error) {
	versionOutput, err := sh.Output(path, "version")
	if err != nil {
		return nil, err
	}
	fmt.Println(versionOutput)
	versionRegexp := regexp.MustCompile(versionPatternGolang)
	matchedVersions := versionRegexp.FindStringSubmatch(versionOutput)
	if len(matchedVersions) == 0 {
		return nil, errors.New("no version string found")
	}
	return version.NewVersion(matchedVersions[1])

}

func EnsureGolang(newVersion string) (*version.Version, error) {
	if file.Exists(golangExecPath) {
		return golangVersion(golangExecPath)
	}
	return nil, nil
}