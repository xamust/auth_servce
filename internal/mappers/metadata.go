package mappers

import (
	"fmt"
	"gitlab.com/xamops/auth/internal/interfaces"
	"gitlab.com/xamops/auth/pkg/metadata"
)

var _ interfaces.MetadataMappers = (*MetadataMappers)(nil)

type MetadataMappers struct {
}

func newMetadataMappers() *MetadataMappers {
	return &MetadataMappers{}
}

func (m *MetadataMappers) ByKey(key string, in []byte) (result interface{}, err error) {
	data, err := m.FromData(in)
	if err != nil {
		return nil, err
	}
	result, ok := data[key]
	if !ok {
		return result, fmt.Errorf("key %s not found", key)
	}
	return
}

func (m *MetadataMappers) FromData(in []byte) (map[string]interface{}, error) {
	return metadata.Encode(in)
}

func (m *MetadataMappers) ToData(in map[string]interface{}) ([]byte, error) {
	return metadata.Decode(in)
}
