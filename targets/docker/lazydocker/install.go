package lazydocker

import (
	"github.com/fizyk/dotfiles/core"
	"github.com/fizyk/dotfiles/core/lazydocker"
	"time"
)

// Install install/update lazydocker
func Install() error {
	defer core.MeasureTime(time.Now(), "docker:lazydocker:install")
	return lazydocker.Install()
}
