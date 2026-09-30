package main

import "fmt"

type Config struct {
	Host    string
	Port    int
	Timeout int
}

func NewConfig(host string, port int) *Config {
	return &Config{
		Host:    host,
		Port:    port,
		Timeout: 30,
	}
}

func main() {
	cfg := NewConfig("localhost", 8080)
	fmt.Printf("%+v\n", *cfg)
}
