package config

import (
	"os"
	"strconv"

	// Load the .env file here in one places instead
	// of in each package.
	_ "github.com/joho/godotenv/autoload"
)

// Config holds system values like which version
// of the API to use.
type Config struct {
	APIKey     string
	APISecret  string
	APIUrl     string
	APIVersion string
	Debug      bool
}

// New will return a new config pointer.
func New() (c *Config) {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		debug = false
	}

	c = &Config{
		APIKey:     os.Getenv("CREX24_API_KEY"),
		APISecret:  os.Getenv("CREX24_API_SECRET"),
		APIUrl:     os.Getenv("CREX24_API_URL"),
		APIVersion: os.Getenv("CREX24_API_VERSION"),
		Debug:      debug,
	}
	return
}
