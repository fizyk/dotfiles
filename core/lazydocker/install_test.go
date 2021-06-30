package lazydocker

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type lazyDockerSuite struct {
	suite.Suite
}

func TestLazyDockerSuite(t *testing.T) {
	suite.Run(t, new(lazyDockerSuite))
}

// Test case matching a release for lazydocker
func (s *lazyDockerSuite) TestMatchRelease() {
	s.True(filter("download_Linux_x86_64.tar.gz"))
}

// Test case not matching a release for lazydocker
func (s *lazyDockerSuite) TestNotMatchRelease() {
	s.False(filter("download_Windows_x86_64.zip"))
}
