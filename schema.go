package avroschema

import "encoding/json"

type AvroSchema struct {
	Name        string        `json:"name,omitempty"`
	Type        interface{}   `json:"type"`
	Items       interface{}   `json:"items,omitempty"`
	Values      interface{}   `json:"values,omitempty"`
	Fields      []*AvroSchema `json:"fields,omitempty"`
	Namespace   string        `json:"namespace,omitempty"`
	Doc         string        `json:"doc,omitempty"`
	Aliases     []string      `json:"aliases,omitempty"`
	Default     interface{}   `json:"default,omitempty"`
	LogicalType string        `json:"logicalType,omitempty"`
}

func StructToJson(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
