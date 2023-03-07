package metrics

import (
	"bytes"
	"errors"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/s16rv/push-metrics-client/pkg/config"
	"github.com/s16rv/push-metrics-client/pkg/request"
)

type Metrics struct {
	Url            string
	MetricFamilies map[string]*dto.MetricFamily
}

func NewMetrics(config config.Config) *Metrics {
	return &Metrics{
		Url:            config.MetricsUrl,
		MetricFamilies: nil,
	}
}

func (m *Metrics) Parse() error {
	body, err := request.GetBody(m.Url)
	if err != nil {
		return err
	}

	var parser expfmt.TextParser
	m.MetricFamilies, err = parser.TextToMetricFamilies(bytes.NewReader(body))
	if err != nil {
		return err
	}

	return nil
}

func (m *Metrics) AppendLabels(labels []*dto.LabelPair) error {
	if m.MetricFamilies == nil {
		return errors.New("Metrics families doesn't exist, parse first")
	}

	for _, mf := range m.MetricFamilies {
		for _, mv := range mf.Metric {
			mv.Label = append(mv.Label, labels...)
		}
	}

	return nil
}

func (m *Metrics) Encode() (string, error) {
	if m.MetricFamilies == nil {
		return "", errors.New("Metrics families doesn't exist, parse first")
	}

	var buff bytes.Buffer
	textEncoder := expfmt.NewEncoder(&buff, expfmt.FmtText)

	encoded := ""
	for _, mf := range m.MetricFamilies {
		err := textEncoder.Encode(mf)
		if err != nil {
			return "", err
		}
		encoded += buff.String()
		buff.Reset()
	}

	return encoded, nil
}
