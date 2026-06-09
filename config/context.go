package config

import (
	"os"
	"strconv"
	"time"
)

const (
	EnvTimeout     = "TAIKUN_API_TIMEOUT"
	DefaultTimeout = 120
)

var APITimeout time.Duration

func ParseTimeout() time.Duration {
	raw := os.Getenv(EnvTimeout)
	if raw == "" {
		return time.Duration(DefaultTimeout) * time.Second
	}

	seconds, err := strconv.Atoi(raw)
	if err != nil || seconds <= 0 {
		return time.Duration(DefaultTimeout) * time.Second
	}

	return time.Duration(seconds) * time.Second
}
