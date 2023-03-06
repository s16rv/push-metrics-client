package main

import (
	"fmt"

	"github.com/s16rv/push-metrics-client/pkg/config"
	"github.com/s16rv/push-metrics-client/pkg/metrics"
)

func main() {
	config := config.NewConfig()

	metrics := metrics.NewMetrics(config)
	data, err := metrics.Parse()

	if err != nil {
		panic(err)
	}

	fmt.Println(data["node_memory_SwapFree_bytes"])

	for i, value := range data {
		fmt.Println(i)
		fmt.Println(value)
		break
	}
}
