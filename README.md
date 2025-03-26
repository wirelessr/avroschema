# Go Avro Schema Reflection

[![workflow](https://github.com/wirelessr/avroschema/actions/workflows/go.yml/badge.svg)](https://github.com/wirelessr/avroschema/actions/workflows/go.yml)
[![Lint](https://github.com/wirelessr/avroschema/actions/workflows/lint.yml/badge.svg)](https://github.com/wirelessr/avroschema/actions/workflows/lint.yml)
[![Go Report](https://goreportcard.com/badge/github.com/wirelessr/avroschema)](https://goreportcard.com/report/github.com/wirelessr/avroschema)

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
    { "name": "a_str_field", "type": "string" },
    { "name": "a_int_field", "type": "int" },
    { "name": "a_bool_field", "type": "boolean" },
    { "name": "a_float_field", "type": "float" },
    { "name": "a_double_field", "type": "double" }
  ]
}
```

## Advanced Configuration

The `Reflector` struct provides several configuration options:

```go
reflector := &avroschema.Reflector{
    BeBackwardTransitive: true,  // Make all fields optional
    EmitAllFields:        true,  // Include fields without tags
    SkipTagFieldNames:    false, // Use JSON/BSON tag names
    Mapper:               nil,   // Custom type mapper
    NameMapping:          nil,   // Custom name mapping
    Namespace:           "",     // Schema namespace
}
```

## Handling Nested Structures

The package supports nested structs and arrays:

```go
type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
}

type Person struct {
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Address Address  `json:"address"`
    Emails  []string `json:"emails"`
}
```

Resulting schema:

```json
{
  "name": "Person",
  "type": "record",
  "fields": [
    { "name": "name", "type": "string" },
    { "name": "age", "type": "int" },
    {
      "name": "address",
      "type": {
        "name": "Address",
        "type": "record",
        "fields": [
          { "name": "street", "type": "string" },
          { "name": "city", "type": "string" }
        ]
      }
    },
    { "name": "emails", "type": { "type": "array", "items": "string" } }
  ]
}
```

## Time Handling

Time values are automatically converted to timestamp-millis:

```go
type Event struct {
    ID        string    `json:"id"`
    Timestamp time.Time `json:"timestamp"`
}
```

```json
{
  "name": "Event",
  "type": "record",
  "fields": [
    { "name": "id", "type": "string" },
    { "name": "timestamp", "type": "long", "logicalType": "timestamp-millis" }
  ]
}
```

## Custom Type Mapping

You can implement custom type mapping:

```go
reflector.Mapper = func(t reflect.Type) any {
    if t == reflect.TypeOf(primitive.ObjectID{}) {
        return "string"
    }
    return nil // fall back to default mapping
}
```

## Optional Fields

Mark fields as optional with `,omitempty`:

```go
type User struct {
    Username string  `json:"username"`
    Email    *string `json:"email,omitempty"`
}
```

```json
{
  "name": "User",
  "type": "record",
  "fields": [
    { "name": "username", "type": "string" },
    { "name": "email", "type": ["null", "string"] }
  ]
}
```

## MongoDB ORM (mgm) Support

The popular MongoDB ORM, [mgm](https://github.com/Kamva/mgm), is supported:

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
