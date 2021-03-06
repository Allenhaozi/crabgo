package register

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Context struct {
	echo.Context
	Config     *Config
	User       *SageUser
	SageHeader *SageHeader
	Workspace  *SageWorkspace
}

func (c Context) DebugModel() bool {
	return c.Config.GeneralConfig.CrabConfig.LogLevel == log.DebugLevel
}

func (c Context) GetNamespace() string {
	return c.Config.GeneralConfig.CrabConfig.Namespace
}

func (c Context) GetDockerSecret() string {
	return c.Config.GeneralConfig.CrabConfig.DockerSecret
}
