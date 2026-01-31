package config

type Config struct {
	Port               string
	RateLimitPerSecond int
	CORSAllowedOrigins []string
}

func LoadConfig() *Config {
	return &Config{
		Port:               ":8080",
		RateLimitPerSecond: 30_000,
		CORSAllowedOrigins: []string{"*"},
	}
}
