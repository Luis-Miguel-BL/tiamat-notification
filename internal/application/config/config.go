package config

import "time"

type Config struct {
	DBConfig DBConfig
}

type DBConfig struct {
	RetryOptions RetryOptions
	DynamoRegion string
}

type RetryOptions struct {
	MaxRetries int
	Delay      time.Duration
}
