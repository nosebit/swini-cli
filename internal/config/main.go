package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//go:embed defaults.yml
var defaultConfig []byte

//go:embed development.yml
var devConfig []byte

type Config struct {
	DataDir string
	ApiUrl  string `mapstructure:"api_url"`
}

var cfg *Config

func Load() (*Config, error) {
	if cfg == nil {
		// Load embedded defaults into Viper
		viper.SetConfigType("yaml")
		if err := viper.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
			return nil, err
		}

		if os.Getenv("GO_ENV") == "development" {
			// Layer development config on top if in development mode
			if err := viper.MergeConfig(bytes.NewBuffer(devConfig)); err != nil {
				return nil, err
			}
		}

		// Then layer environment variables
		viper.AutomaticEnv()
		viper.SetEnvPrefix("SWINI")

		// Unmarshal into config struct
		if err := viper.Unmarshal(&cfg); err != nil {
			return nil, err
		}

		// Force DataDir to a static path regardless of config/env
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return nil, err
		}
		cfg.DataDir = home + "/.swini"

		// Now layer user config on top
		// home, _ := os.UserHomeDir()
		// viper.SetConfigName("config")
		// viper.SetConfigType("yaml")
		// viper.AddConfigPath(filepath.Join(home, ".swini"))
		// viper.AddConfigPath(".")
		// _ = viper.MergeInConfig() // merge instead of replace
	}

	return cfg, nil
}
