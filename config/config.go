package config

import "log"

// Appconfig holds the configuration for the app
type AppConfig struct {
	AppName    string
	AppVersion string
	AppPort    string
	InfoLog    *log.Logger
}
