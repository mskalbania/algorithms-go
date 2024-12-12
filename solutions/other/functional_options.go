package other

import (
	"errors"
)

const httpDefaultPort = 8080

type (
	OptionsFunction func(*properties) error
	Client          struct{}
	properties      struct { //unexported struct holds properties
		port int
	}
)

func NewClient(options ...OptionsFunction) (*Client, error) {
	properties := new(properties)
	for _, option := range options {
		err := option(properties)
		if err != nil {
			return nil, err
		}
	}
	if properties.port == 0 { //default when unset
		properties.port = httpDefaultPort
	}
	return new(Client), nil
}

func WithPort(port int) OptionsFunction {
	return func(p *properties) error {
		if port < 0 { //validations where present
			return errors.New("port negative")
		}
		p.port = port
		return nil
	}
}
