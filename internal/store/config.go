package store

type Config struct {
	DatabaseUrl string
}

func NewConfig(DatabaseUrl string) *Config {
	return &Config{DatabaseUrl}
}
