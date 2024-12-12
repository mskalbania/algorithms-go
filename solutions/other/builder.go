package other

import (
	"errors"
)

const defaultPort = 8080

type Config struct {
	port int
}

func (c Config) Port() int {
	return c.port
}

type ConfigBuilder struct {
	Port *int
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{}
}

func (cb *ConfigBuilder) WithPort(port int) *ConfigBuilder {
	cb.Port = &port
	return cb
}

func (cb *ConfigBuilder) Build() (Config, error) {
	port := defaultPort
	if cb.Port != nil {
		port = *cb.Port
	}
	if port <= 0 {
		return Config{}, errors.New("port must be greater than 0")
	}
	return Config{port: port}, nil
}
