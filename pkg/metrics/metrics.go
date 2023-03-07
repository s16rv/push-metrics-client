package metrics

import (
	"bytes"
	"io"
	"net/http"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/s16rv/push-metrics-client/pkg/config"
)

type Metrics struct {
	Url string
}

func NewMetrics(config config.Config) Metrics {
	return Metrics{
		Url: config.MetricsUrl,
	}
}

func (m Metrics) getBody() ([]byte, error) {
	resp, err := http.Get(m.Url)
	if err != nil {
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (m Metrics) Parse() (map[string]*dto.MetricFamily, error) {
	body, err := m.getBody()
	if err != nil {
		return nil, err
	}

	var parser expfmt.TextParser
	mf, err := parser.TextToMetricFamilies(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	return mf, nil
}

func Encode(mf map[string]*dto.MetricFamily) (string, error) {
	var buff bytes.Buffer
	textEncoder := expfmt.NewEncoder(&buff, expfmt.FmtText)

	encoded := ""
	for _, value := range mf {
		err := textEncoder.Encode(value)
		if err != nil {
			return "", err
		}
		encoded += buff.String()
		buff.Reset()
	}

	return encoded, nil
}

func AppendLabels(mf map[string]*dto.MetricFamily, labels []*dto.LabelPair) map[string]*dto.MetricFamily {
	for _, v := range mf {
		for _, mv := range v.Metric {
			mv.Label = append(mv.Label, labels...)
		}
	}

	return mf
}
