package metadata

import "encoding/json"

type Metadata []byte

func Encode(m Metadata) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if err := json.Unmarshal(m, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func Decode(in map[string]interface{}) ([]byte, error) {
	return json.Marshal(in)
}
