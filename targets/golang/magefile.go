package golang

import "github.com/magefile/mage/sh"

// Format formats go code
func Format() {
	sh.RunV("go", "fmt", "./...")
}

// Tidy tidies go.mod file
func Tidy() {
	sh.RunV("go", "mod", "tidy")
}

// Test run tests for go code
func Test() {
	sh.RunV("go", "test", "-race", "-coverpkg", "./...", "-coverprofile=coverage.txt", "-covermode=atomic", "./...")
}
