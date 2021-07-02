package docker

import (
	"fmt"
	"github.com/fizyk/magex/file"
	"github.com/magefile/mage/sh"
	"io/ioutil"
)

const dockerAptFile string = "/etc/apt/sources.list.d/docker.list"

// InstallDocker actually installs docker given that the apt repository is prepared
func InstallDocker() error {
	if !file.Exists(dockerAptFile) {
		releaseOutput, err := sh.Output("lsb_release", "-cs")
		if err != nil {
			return err
		}
		var dockerAptLine string = fmt.Sprintf("deb [arch=amd64 signed-by=%s] https://download.docker.com/linux/ubuntu %s stable\n", dockerAptPGPKeyfile, releaseOutput)
		if err := ioutil.WriteFile("docker.list", []byte(dockerAptLine), 0644); err != nil {
			return err
		}
		sh.Run("sudo", "mv", "docker.list", dockerAptFile)
		sh.Run("sudo", "chown", "root:root", dockerAptFile)
		sh.Run("sudo", "apt", "update", "-o", fmt.Sprintf("Dir::Etc::sourcelist=%s", dockerAptFile), "-o", "Dir::Etc::sourceparts=-", "-o", "APT::Get::List-Cleanup=0")
	}
	sh.RunV("sudo", "apt", "install", "docker-ce", "docker-ce-cli", "containerd.io")
	return nil
}
