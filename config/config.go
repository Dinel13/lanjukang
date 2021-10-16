package config

// Appconfig holds the configuration for the app
type AppConfig struct {
	AppName    string
	AppVersion string
	JwtSecret  string
	Frontend   string
}
