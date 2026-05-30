package configs

import (
	"encoding/json"
	"os"
)

const ConfigFile = "/opt/orp/config.json"

type Config struct {
	JitsiBase string `json:"jitsi_base"`
}

func Load() (*Config, error) {
	cfg := &Config{
		JitsiBase: "meet.handyweb.org",
	}

	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func Save(cfg *Config) error {
	err := os.MkdirAll("/opt/orp", 0755)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigFile, data, 0644)
}
