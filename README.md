# Avro Schema

![workflow](https://github.com/wirelessr/avroschema/actions/workflows/go.yml/badge.svg)
![Lint](https://github.com/wirelessr/avroschema/actions/workflows/lint.yml/badge.svg)

This package can be used to generate [Avro Schemas](https://avro.apache.org/docs/1.11.1/specification/) from Go types through reflection.

## Basic Usage

The following Go type:

```go
type Entity struct {
		AStrField    *string  `json:"a_str_field"`
		AIntField    *int     `json:"a_int_field"`
		ABoolField   *bool    `json:"a_bool_field"`
		AFloatField  *float32 `json:"a_float_field"`
		ADoubleField *float64 `json:"a_double_field"`
	}
```

Results in following JSON Schema:

```go
import "github.com/wirelessr/avroschema"

avroschema.Reflect(&Entity{})
```

```json
{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_str_field", "type": "string"},
      {"name": "a_int_field", "type": "int"},
      {"name": "a_bool_field", "type": "boolean"},
      {"name": "a_float_field", "type": "float"},
      {"name": "a_double_field", "type": "double"}
    ]
  }
```

## More Advanced Extensions

The popular MongoDB ORM, [mgm](https://github.com/Kamva/mgm), is supported.

The following Go type:

```go
type Book struct {
		mgm.DefaultModel `bson:",inline"`
		Name             string             `json:"name" bson:"name"`
		Pages            int                `json:"pages" bson:"pages"`
		ObjId            primitive.ObjectID `json:"obj_id" bson:"obj_id"`
		ArrivedAt        primitive.DateTime `json:"arrived_at" bson:"arrived_at"`
		RefData          bson.M             `json:"ref_data" bson:"ref_data"`
		Author           []string           `json:"author" bson:"author"`
	}
```

The type mappings can be customized by `Mapper`.

```go
import (
	"github.com/wirelessr/avroschema"
	"github.com/wirelessr/avroschema/mongo"
)

reflector := new(avroschema.Reflector)
reflector.Mapper = MgmExtension

reflector.Reflect(&Book{})
```

Results in following JSON Schema:

```json
{
		"name": "Book",
		"type": "record",
		"fields": [
			{ "name": "_id", "type": "string" },
			{ "name": "created_at", "type": "long", "logicalType": "timestamp-millis" },
			{ "name": "updated_at", "type": "long", "logicalType": "timestamp-millis" },
			{ "name": "name", "type": "string" },
			{ "name": "pages", "type": "int" },
			{ "name": "obj_id", "type": "string" },
			{ "name": "arrived_at", "type": "long", "logicalType": "timestamp-millis" },
			{ "name": "ref_data", "type": "string" },
			{ "name": "author", "type": "array", "items": "string" } 
		]
	}
```
