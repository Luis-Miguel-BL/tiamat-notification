package config

import "time"

type Config struct {
	AppName     string
	EntryPoint  string
	EventBridge EventBridgeConfig
	DBConfig    DBConfig
}

type EventBridgeConfig struct {
	Region       string
	EventBusName string
}
type DBConfig struct {
	RetryOptions RetryOptions
	DynamoRegion string
}

type RetryOptions struct {
	MaxRetries int
	Delay      time.Duration
}
