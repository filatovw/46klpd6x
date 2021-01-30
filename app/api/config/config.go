package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config of API server
type Config struct {
	Host           string
	Port           int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

// ConnectionString get <host:port> representation
func (c *Config) ConnectionString() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// Load config from environment variables
func Load() (*Config, error) {
	var err error
	config := &Config{
		Host:           "localhost",
		Port:           8000,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if val, ok := os.LookupEnv("API_HOST"); ok {
		config.Host = val
	}
	if val, ok := os.LookupEnv("API_PORT"); ok {
		config.Port, err = strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("failed to parse API_PORT: %w", err)
		}
	}
	if val, ok := os.LookupEnv("API_READ_TIMEOUT"); ok {
		readTimeout, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("failed to parse API_READ_TIMEOUT: %w", err)
		}
		config.ReadTimeout = time.Duration(readTimeout) * time.Second
	}
	if val, ok := os.LookupEnv("API_WRITE_TIMEOUT"); ok {
		writeTimeout, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("failed to parse API_WRITE_TIMEOUT: %w", err)
		}
		config.WriteTimeout = time.Duration(writeTimeout) * time.Second
	}
	if val, ok := os.LookupEnv("MAX_HEADER_BYTES"); ok {
		config.MaxHeaderBytes, err = strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("failed to parse MAX_HEADER_BYTES: %w", err)
		}
	}
	return config, nil
}
