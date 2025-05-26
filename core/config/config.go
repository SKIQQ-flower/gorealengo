package config

import (
	"os"
	"sync"

	"github.com/skwb/realengo-conflict/core/log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Game   GameConfig   `toml:"game"`
	Window WindowConfig `toml:"window"`
}

type GameConfig struct {
	Debug    bool   `toml:"debug"`
	GameName string `toml:"name"`
}

type WindowConfig struct {
	WindowWidth    int32 `toml:"window_width"`
	WindowHeight   int32 `toml:"window_height"`
	ViewportWidth  int32 `toml:"viewport_width"`
	ViewportHeight int32 `toml:"viewport_height"`
}

var (
	once      sync.Once
	config    *Config
	loadError error
)

func LoadConfig() (*Config, error) {
	once.Do(func() {
		config, loadError = LoadConfigPath("./config.toml")
	})
	return config, loadError
}

func LoadConfigPath(path string) (*Config, error) {
	var conf Config
	data, err := os.ReadFile(path)
	if err != nil {
		log.Logger.Err(err).Msg("Unable to open config file (Does this file really exists?)")
		return nil, err
	}

	if _, err := toml.Decode(string(data), &conf); err != nil {
		log.Logger.Err(err).Msg("Unable to decode configs (Your computer is well?)")
		return nil, err
	}

	return &conf, nil
}
