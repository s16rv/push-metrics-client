package metadata

import (
	dto "github.com/prometheus/client_model/go"
	"github.com/s16rv/push-metrics-client/pkg/config"
	"github.com/s16rv/push-metrics-client/pkg/request"
	"google.golang.org/protobuf/proto"
)

const (
	LabelId            = "id"
	LabelName          = "name"
	LabelRegion        = "region"
	PathMetadataId     = "/metadata/v1/id"
	PathMetadataName   = "/metadata/v1/hostname"
	PathMetadataRegion = "/metadata/v1/region"
)

type Metadata struct {
	Url           string
	MetadataPaths []MetadataPath
}

type MetadataPath struct {
	Name string
	Path string
}

func NewMetadata(config config.Config) Metadata {
	return Metadata{
		Url: config.DOMetadataUrl,
		MetadataPaths: []MetadataPath{
			{
				Name: LabelId,
				Path: PathMetadataId,
			},
			{
				Name: LabelName,
				Path: PathMetadataName,
			},
			{
				Name: LabelRegion,
				Path: PathMetadataRegion,
			},
		},
	}
}

func (m Metadata) GetMetadataLabels() ([]*dto.LabelPair, error) {
	labels := []*dto.LabelPair{}
	for _, path := range m.MetadataPaths {
		body, err := request.GetBody(m.Url + path.Path)
		if err != nil {
			return []*dto.LabelPair{}, err
		}
		labels = append(labels, &dto.LabelPair{
			Name:  proto.String(path.Name),
			Value: proto.String(string(body)),
		})
	}
	return labels, nil
}
