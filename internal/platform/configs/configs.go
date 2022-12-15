package configs

import "log"

type AppConfig struct {
	// Port to listen
	Port string
	// Environment
	Env string
	// Logger
	Logger log.Logger
}
