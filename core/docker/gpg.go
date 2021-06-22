package docker

import (
	"errors"
	"github.com/fizyk/dotfiles/core/command"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

const dockerAptPGPKeyfile string = "/usr/share/keyrings/docker-archive-keyring.gpg"
const dockerAptPGPKeyUri string = "https://download.docker.com/linux/ubuntu/gpg"
const rawPGPKeyFile string = "docker_apt_key_raw.gpg"

func PGP() error {
	if _, err := os.Stat(dockerAptPGPKeyfile); errors.Is(err, os.ErrNotExist) {
		// Download PGP file if not downloaded
		if _, err := os.Stat(rawPGPKeyFile); errors.Is(err, os.ErrNotExist) {
			resp, err := http.Get(dockerAptPGPKeyUri)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if body, err := ioutil.ReadAll(resp.Body); err != nil {
				return err
				// Write downloaded file locally
			} else if err := ioutil.WriteFile(rawPGPKeyFile, body, 0644); err != nil {
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
