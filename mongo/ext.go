package mongo

import (
	"reflect"

	"github.com/wirelessr/avroschema"
)

func MgmExtension(t reflect.Type) any {
	switch t.Name() {
	case "DateTime": // primitive.DateTime
		return &avroschema.AvroSchema{Type: "long", LogicalType: "timestamp-millis"}
	case "ObjectID": // primitive.ObjectID
		return "string"
	case "M": // bson.M
		return "string"
	}
	return nil
}
