package main

import (
	"fmt"

	"github.com/s16rv/push-metrics-client/pkg/config"
	"github.com/s16rv/push-metrics-client/pkg/metadata"
	"github.com/s16rv/push-metrics-client/pkg/metrics"
)

func main() {
	config := config.NewConfig()

	m := metrics.NewMetrics(config)
	err := m.Parse()
	if err != nil {
		panic(err)
	}

	md := metadata.NewMetadata(config)
	labels, err := md.GetMetadataLabels()
	if err != nil {
		panic(err)
	}

	err = m.AppendLabels(labels)
	if err != nil {
		panic(err)
	}

	encoded, err := m.Encode()
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded)
}
