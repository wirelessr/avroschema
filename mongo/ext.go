package mongo

import (
	"reflect"

	"github.com/wirelessr/avroschema"
)

func MgmExtension(t reflect.Type) interface{} {
	switch t.Name() {
	case "DateTime": // primitive.DateTime
		return &avroschema.AvroSchema{Type: "long", LogicalType: "timestamp-millis"}
	case "ObjectID": // primitive.ObjectID
		return "string"
	case "DefaultModel": // mgm.DefaultModel
		return []*avroschema.AvroSchema{
			{Name: "_id", Type: "string"},
			{Name: "created_at", Type: "long", LogicalType: "timestamp-millis"},
			{Name: "updated_at", Type: "long", LogicalType: "timestamp-millis"},
		}
	case "M": // bson.M
		return "string"
	}
	return nil
}
