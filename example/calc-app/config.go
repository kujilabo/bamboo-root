package main

import (
	"embed"
	"os"

	_ "embed"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"

	"github.com/kujilabo/bamboo-root/helper"
	"github.com/kujilabo/bamboo-root/internal"
)

var (
	Validator = validator.New()
)

type AppConfig struct {
	Name string `yaml:"name" validate:"required"`
}

type Config struct {
	App     *AppConfig                            `yaml:"app" validate:"required"`
	Workers map[string]*helper.WorkerClientConfig `yaml:"workers" validate:"required"`
	Trace   *helper.TraceConfig                   `yaml:"trace" validate:"required"`
	Log     *helper.LogConfig                     `yaml:"log" validate:"required"`
}

//go:embed debug.yml
var config embed.FS

func LoadConfig(env string) (*Config, error) {
	filename := env + ".yml"
	confContent, err := config.ReadFile(filename)
	if err != nil {
		return nil, internal.Errorf("config.ReadFile. filename: %s, err: %w", filename, err)
	}

	confContent = []byte(os.ExpandEnv(string(confContent)))
	conf := &Config{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		return nil, internal.Errorf("yaml.Unmarshal. filename: %s, err: %w", filename, err)
	}

	if err := Validator.Struct(conf); err != nil {
		return nil, internal.Errorf("Validator.Structl. filename: %s, err: %w", filename, err)
	}

	return conf, nil
}
