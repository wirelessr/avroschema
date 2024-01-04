package avroschema

import (
	"encoding/json"
	"reflect"
	"strings"
)

type AvroSchema struct {
	Name      string        `json:"name"`
	Type      interface{}   `json:"type"`
	Items     interface{}   `json:"items,omitempty"`
	Values    interface{}   `json:"values,omitempty"`
	Fields    []*AvroSchema `json:"fields,omitempty"`
	Namespace string        `json:"namespace,omitempty"`
	Doc       string        `json:"doc,omitempty"`
	Aliases   []string      `json:"aliases,omitempty"`
	Default   interface{}   `json:"default,omitempty"`
}

func handleRecord(v any) *AvroSchema {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	name := t.Name()
	tokens := strings.Split(name, ".")
	name = tokens[len(tokens)-1]

	ret := &AvroSchema{Name: name, Type: "record"}

	for i, n := 0, t.NumField(); i < n; i++ { // handle fields
		f := t.Field(i)

		jsonTag := f.Tag.Get("json")
		tokens := strings.Split(jsonTag, ",")
		jsonFieldName := tokens[0]

		switch f.Type.Kind() {
		case reflect.String:
			ret.Fields = append(ret.Fields, &AvroSchema{Name: jsonFieldName, Type: "string"})
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
			ret.Fields = append(ret.Fields, &AvroSchema{Name: jsonFieldName, Type: "int"})
		case reflect.Int64, reflect.Uint64:
			ret.Fields = append(ret.Fields, &AvroSchema{Name: jsonFieldName, Type: "long"})
		case reflect.Float32:
			ret.Fields = append(ret.Fields, &AvroSchema{Name: jsonFieldName, Type: "float"})
		case reflect.Float64:
			ret.Fields = append(ret.Fields, &AvroSchema{Name: jsonFieldName, Type: "double"})
		case reflect.Bool:
			ret.Fields = append(ret.Fields, &AvroSchema{Name: jsonFieldName, Type: "boolean"})

		default:
		}
	}
	return ret
}

func Reflect(v any) (string, error) {
	data := handleRecord(v)

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
