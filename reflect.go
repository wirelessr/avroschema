package avroschema

import (
	"reflect"
	"strings"
)

func reflectType(t reflect.Type) interface{} {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return "int"
	case reflect.Int64, reflect.Uint64:
		return "long"
	case reflect.Float32:
		return "float"
	case reflect.Float64:
		return "double"
	case reflect.Bool:
		return "boolean"
	case reflect.Array, reflect.Slice:
		return handleArray(t)
	case reflect.Struct:
		// TODO: handle special types, e.g. time.Time
		return handleRecord(t)
	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			// If the key is not a string, then treat the whole object as a string.
			return "string"
		}
		return handleMap(t)
	default:
		return "" // FIXME: no error handle
	}
}

func handleMap(t reflect.Type) *AvroSchema {
	return &AvroSchema{Type: "map", Values: reflectType(t.Elem())}
}

func handleArray(t reflect.Type) *AvroSchema {
	return &AvroSchema{Type: "array", Items: reflectType(t.Elem())}
}

func handleRecord(t reflect.Type) *AvroSchema {
	name := t.Name()
	tokens := strings.Split(name, ".")
	name = tokens[len(tokens)-1]

	ret := &AvroSchema{Name: name, Type: "record"}

	// reflect.Type: t & f.Type & f.Type.Elem() & f.Type.Key()
	for i, n := 0, t.NumField(); i < n; i++ { // handle fields
		f := t.Field(i)

		jsonTag := f.Tag.Get("json")
		tokens := strings.Split(jsonTag, ",")
		jsonFieldName := tokens[0]

		if jsonFieldName == "" {
			continue
		}

		ret.Fields = append(ret.Fields, reflectEx(f.Type, jsonFieldName))
	}
	return ret
}

func reflectEx(t reflect.Type, n string) *AvroSchema {
	ret := reflectType(t)
	if reflect.TypeOf(ret).Kind() == reflect.String {
		return &AvroSchema{Name: n, Type: ret}
	}

	result, ok := ret.(*AvroSchema)
	if !ok {
		return nil
	}
	result.Name = n
	return result
}

func Reflect(v any) (string, error) {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	data := handleRecord(t)

	return StructToJson(data)
}
