package docker

import (
	"github.com/fizyk/dotfiles/core/docker"
	magexTime "github.com/fizyk/magex/time"
	"time"
)

// Install install docker and it's apt repositories:
// https://docs.docker.com/engine/install/ubuntu/
func Install() error {
	defer magexTime.MeasureTime(time.Now(), "docker:install")
	err := docker.PGP()
	if err != nil {
		return err
	}
	return docker.InstallDocker()
}
