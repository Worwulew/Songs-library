package utils

import "github.com/Worwulew/Songs-library/pkg/logger"

// LogResponseError sends Error response with logging error
func LogResponseError(logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, Error: %s", err,
	)
}
