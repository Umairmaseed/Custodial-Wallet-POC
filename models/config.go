package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type EnvConfig struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	ServerAddr string `mapstructure:"SERVER_ADDR"`
	Mode       string `mapstructure:"MODE"`
}

func (config *EnvConfig) Validate() error {
	return validation.ValidateStruct(config,
		validation.Field(&config.ServerPort, is.Port),
		validation.Field(&config.ServerAddr, validation.Required),
		validation.Field(&config.Mode, validation.In("debug", "release")),
	)
}
