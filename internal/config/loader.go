package config

func LoadConfig() *Config {
	return &Config{
		EntryPoint: "event-bridge",
		AppName:    "customer-change-event-consumer",
	}
}
