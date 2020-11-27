package config

import "bufio"

type Config struct {
	vals bufio.Reader
}

func NewConfig() *Config {
	return &Config{}
}
