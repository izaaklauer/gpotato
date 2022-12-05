package config

type Gpotato struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// DefaultGpotatoConfig returns default config values
func DefaultGpotatoConfig() Gpotato {
	return Gpotato{
	HelloWorldMessage:
		"hello from the default config",
	}
}
