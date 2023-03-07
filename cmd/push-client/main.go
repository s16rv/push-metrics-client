package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	dto "github.com/prometheus/client_model/go"
	"github.com/s16rv/push-metrics-client/pkg/config"
	"github.com/s16rv/push-metrics-client/pkg/metrics"
)

func main() {
	config := config.NewConfig()

	m := metrics.NewMetrics(config)
	data, err := m.Parse()

	if err != nil {
		panic(err)
	}

	labels := []*dto.LabelPair{
		{
			Name:  proto.String("id"),
			Value: proto.String("123"),
		},
	}

	data = metrics.AppendLabels(data, labels)
	encoded, err := metrics.Encode(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded)
}
