package docker

import (
	"errors"
	"github.com/fizyk/magex/command"
	"github.com/fizyk/magex/http"
	"os"
	"os/exec"
)

const dockerAptPGPKeyfile string = "/usr/share/keyrings/docker-archive-keyring.gpg"
const dockerAptPGPKeyUri string = "https://download.docker.com/linux/ubuntu/gpg"
const rawPGPKeyFile string = "docker_apt_key_raw.gpg"

// PGP Downloads and installs PGP key for docker apt repository
func PGP() error {
	if _, err := os.Stat(dockerAptPGPKeyfile); errors.Is(err, os.ErrNotExist) {
		// Download PGP file if not downloaded
		if _, err := os.Stat(rawPGPKeyFile); errors.Is(err, os.ErrNotExist) {
			if err := http.DownloadFile(dockerAptPGPKeyUri, rawPGPKeyFile); err != nil {
				return err
			}
		}
		// Dearmor the raw file with pgp and save it to final location
		mainCommand := exec.Command("cat", rawPGPKeyFile)
		pipeToCommand := exec.Command("sudo", "gpg", "--yes", "--dearmor", "-o", dockerAptPGPKeyfile)
		command.PipeCommands(mainCommand, pipeToCommand)
	}
	return nil
}
