package main

type Config struct {
	Port int    `toml:"port"`
	Env  string `toml:"env"`
	Db   struct {
		Dsn          string `toml:"dsn"`
		MaxOpenConns int    `toml:"maxOpenConns"`
		MaxIdleConns int    `toml:"maxIdleConns"`
		MaxIdleTime  string `toml:"maxIdleTime"`
	}
}

func NewConfig() Config {
	return Config{}
}
