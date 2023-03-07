package pushgateway

import (
	"github.com/s16rv/push-metrics-client/pkg/config"
	"github.com/s16rv/push-metrics-client/pkg/request"
)

type Pushgateway struct {
	Url      string
	Job      string
	Instance string
}

func NewPushgateway(config config.Config, instance string) Pushgateway {
	return Pushgateway{
		Url:      config.PushgatewayUrl,
		Job:      "pushgateway",
		Instance: instance,
	}
}

func (p Pushgateway) PushMetrics(metrics string) error {
	url := p.Url + "/metrics/job/" + p.Job + "/instance/" + p.Instance
	return request.PostText(url, metrics)
}
