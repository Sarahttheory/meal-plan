package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port string `env:"PORT" env-default:"8080"`
	DB   struct {
		User    string `env:"DB_USER" env-required:"true"`
		Pass    string `env:"DB_PASS" env-required:"true"`
		Host    string `env:"DB_HOST" env-required:"localhost"`
		Port    string `env:"DB_PORT" env-required:"5432"`
		Name    string `env:"DB_NAME" env-required:"true"`
		SSLMode string `env:"DB_SSLMODE" env-default:"disable"`
	}
}

// чтобы дважды не грузить конфиг
var (
	cfg  *Config
	once sync.Once
)

func Load() *Config {
	once.Do(func() {
		cfg = &Config{}
		if err := cleanenv.ReadConfig(".env", cfg); err != nil {
			if err := cleanenv.ReadEnv(cfg); err != nil {
				log.Fatalf("Config error: %s", err)
			}
		}
	})
	return cfg
}

func (c *Config) GetDNS() string {
	return "postgres://" + c.DB.User + ":" + c.DB.Pass + "@" +
		c.DB.Host + ":" + c.DB.Port + "/" + c.DB.Name +
		"?sslmode=" + c.DB.SSLMode
}
