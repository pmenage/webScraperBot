package config

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// Config contains the Telegram Bot Key
type Config struct {
	TelegramKey string `json:"telegram_key"`
}

// FromYAML reads from a YAML file
func FromYAML(filepath string) Config {
	byt, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(errors.Wrap(err, "Could not load YAML file: "+filepath))
	}
	var c Config
	if err := yaml.Unmarshal(byt, &c); err != nil {
		panic(errors.Wrap(err, "Could not parse YAML file: "+filepath))
	}
	return c
}
