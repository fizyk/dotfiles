package docker

import (
	"fmt"
	"github.com/fizyk/dotfiles/core"
	"github.com/magefile/mage/sh"
	"strings"
	"time"
)


func hasDocker(groups string, searchGroup string) bool {
	groupList := strings.Split(groups, " ")
	for _, group := range groupList {
		if group == searchGroup {
			return true
		}
	}
	return false
}

// Group makes sure user is already in a docker group
func Group() error {
	defer core.MeasureTime(time.Now(), "docker:group")
	if groups, err := sh.Output("groups"); err != nil {
		return err
	} else if hasDocker(groups, "docker") {
		fmt.Println("User already in `docker` group")
	} else {
		if err := sh.Run("sudo", "usermod", "-a", "-G", "docker", "$USER"); err != nil {
			return err
		}
		if err := sh.Run("newgroup", "docker"); err != nil {
			return err
		}
		if err := sh.Run("newgroup", "$USER"); err != nil {
			return err
		}
	}
	return nil
}