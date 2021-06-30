package lazydocker

import (
	"github.com/fizyk/dotfiles/core/lazydocker"
	magexTime "github.com/fizyk/magex/time"
	"time"
)

// Install install/update lazydocker
func Install() error {
	defer magexTime.MeasureTime(time.Now(), "docker:lazydocker:install")
	return lazydocker.Install()
}
