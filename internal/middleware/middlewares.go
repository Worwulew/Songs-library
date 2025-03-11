package middleware

import (
	"github.com/Worwulew/Songs-library/config"
	"github.com/Worwulew/Songs-library/pkg/logger"
)

type MiddlewareManager struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewMiddlewareManager(cfg *config.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, logger: logger}
}
