package config

import (
	arg "github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
)

type Config struct {
	MetricsUrl     string `arg:"env:METRICS_URL"`
	DOMetadataUrl  string `arg:"env:DO_METADATA_URL"`
	PushgatewayUrl string `arg:"env:PUSHGATEWAY_URL"`
	PushInterval   uint64 `arg:"env:PUSH_INTERVAL"` //seconds
}

func NewConfig() Config {
	_ = godotenv.Load()

	c := Config{
		MetricsUrl:     "http://127.0.0.1:9100",
		DOMetadataUrl:  "http://169.254.169.254",
		PushgatewayUrl: "http://127.0.0.1:9091",
		PushInterval:   30,
	}

	arg.MustParse(&c)

	return c
}
