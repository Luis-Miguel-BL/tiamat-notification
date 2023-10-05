package config

func LoadConfig() *Config {
	return &Config{
		EntryPoint: "event-bridge",
		AppName:    "customer-change-event-consumer",
		Server:     ServerConfig{Port: 8080},
		DBConfig:   DBConfig{InMemory: true},
		Redis:      RedisConfig{Address: "127.0.0.1:6379", Password: "", DB: 0},
	}
}
