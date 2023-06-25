package config

func LoadConfig() *Config {
	return &Config{
		EntryPoint: "event-bridge",
	}
}
