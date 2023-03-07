package config

import (
	arg "github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPTimeout   int    `arg:"env:HTTP_TIMEOUT"`
	MetricsUrl    string `arg:"env:METRICS_URL"`
	DOMetadataUrl string `arg:"env:DO_METADATA_URL"`
}

func NewConfig() Config {
	_ = godotenv.Load()

	c := Config{
		HTTPTimeout:   5000,
		MetricsUrl:    "http://127.0.0.1:9100",
		DOMetadataUrl: "http://169.254.169.254",
	}

	arg.MustParse(&c)

	return c
}
