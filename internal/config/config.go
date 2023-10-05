package config

import "time"

type Config struct {
	AppName     string
	EntryPoint  string
	EventBridge EventBridgeConfig
	Redis       RedisConfig
	DBConfig    DBConfig
	Server      ServerConfig
}

type EventBridgeConfig struct {
	Region       string
	EventBusName string
}
type RedisConfig struct {
	Address  string
	Password string
	DB       int
}
type DBConfig struct {
	InMemory     bool
	RetryOptions RetryOptions
	DynamoRegion string
}

type RetryOptions struct {
	MaxRetries int
	Delay      time.Duration
}

type ServerConfig struct {
	Port int
}
