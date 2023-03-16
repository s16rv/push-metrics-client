package main

import (
	"log"

	"github.com/jasonlvhit/gocron"
	"github.com/s16rv/push-metrics-client/pkg/config"
	"github.com/s16rv/push-metrics-client/pkg/metadata"
	"github.com/s16rv/push-metrics-client/pkg/metrics"
	"github.com/s16rv/push-metrics-client/pkg/pushgateway"
)

func task(config config.Config) {
	log.Println("Start pushing metrics")
	m := metrics.NewMetrics(config)
	err := m.Parse()
	if err != nil {
		log.Println(err)
		return
	}

	md := metadata.NewMetadata(config)
	labels, err := md.GetMetadataLabels()
	if err != nil {
		log.Println(err)
		return
	}

	err = m.AppendLabels(labels)
	if err != nil {
		log.Println(err)
		return
	}

	encoded, err := m.Encode()
	if err != nil {
		log.Println(err)
		return
	}

	pg := pushgateway.NewPushgateway(config)
	err = pg.PushMetrics(encoded)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Finished!")
}

func main() {
	config := config.NewConfig()

	s := gocron.NewScheduler()
	err := s.Every(config.PushInterval).Seconds().Do(task, config)
	if err != nil {
		log.Fatalln(err)
	}

	<-s.Start()
}
