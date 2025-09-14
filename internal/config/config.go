package config

import "os"

type Config struct {
	Addr    string
	DataDir string
}

func Load() Config {
	addr := os.Getenv("API_ADDR")
	if addr == "" { addr = ":8080" }

	data := os.Getenv("DATA_DIR")
	if data == "" { data = "data" }

	return Config{Addr: addr, DataDir: data}
}
