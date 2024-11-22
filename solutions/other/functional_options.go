package other

import "errors"

const httpDefaultPort = 8080

type (
	OptionsFunction func(*Properties) error
	Client          struct{}
	Properties      struct {
		port int
	}
)

func NewClient(options ...OptionsFunction) (*Client, error) {
	properties := new(Properties)
	for _, option := range options {
		err := option(properties)
		if err != nil {
			return nil, err
		}
	}
	if properties.port == 0 {
		properties.port = httpDefaultPort
	}
	return new(Client), nil
}

func WithPort(port int) OptionsFunction {
	return func(p *Properties) error {
		if port < 0 {
			return errors.New("port negative")
		}
		p.port = port
		return nil
	}
}
