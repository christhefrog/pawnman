package project

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Ready bool `json:"-"`

	Profiles map[string]Profile `json:"profiles"`
}

type Profile struct {
	CompilerVersion string   `json:"compiler_version"`
	Input           string   `json:"input"`
	Output          string   `json:"output"`
	Includes        []string `json:"includes"`
	Args            []string `json:"args"`
}

var config Config

func LoadConfig() (*Config, error) {
	if config.Ready {
		return &config, nil
	}

	path := "pawnctl.json"

	wd, _ := os.Getwd()
	path = fmt.Sprint(wd, "\\", path)

	file, err := os.ReadFile(path)
	if err == nil {
		err = json.Unmarshal(file, &config)
		if err != nil {
			return nil, err
		}
	}

	if config.Profiles == nil {
		config.Profiles = make(map[string]Profile)

		// config.Profiles[""] = Profile{
		// 	Includes: make([]string, 0),
		// 	Args:     make([]string, 0),
		// }
	}

	config.Ready = true

	return &config, nil
}

func (c *Config) Save() error {
	if !c.Ready {
		return errors.New("config handle isn't opened")
	}

	path := "pawnctl.json"

	wd, _ := os.Getwd()
	path = fmt.Sprint(wd, "\\", path)

	bytes, _ := json.MarshalIndent(c, "", "\t")

	os.WriteFile(path, bytes, 0664)

	return nil
}
