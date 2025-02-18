package avroschema

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrimitiveType(t *testing.T) {
	type Entity struct {
		AStrField    string  `json:"a_str_field"`
		AIntField    int     `json:"a_int_field"`
		ABoolField   bool    `json:"a_bool_field"`
		AFloatField  float32 `json:"a_float_field"`
		ADoubleField float64 `json:"a_double_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_str_field", "type": "string"},
      {"name": "a_int_field", "type": "int"},
      {"name": "a_bool_field", "type": "boolean"},
      {"name": "a_float_field", "type": "float"},
      {"name": "a_double_field", "type": "double"}
    ]
  }`

	e := Entity{}

	// test for instance
	r1, err1 := Reflect(e)
	assert.JSONEq(t, expected, r1)
	assert.Nil(t, err1)

	// test for pointer
	r2, err2 := Reflect(&e)
	assert.JSONEq(t, expected, r2)
	assert.Nil(t, err2)
}

func TestPrimitivePointer(t *testing.T) {
	type Entity struct {
		AStrField    *string  `json:"a_str_field"`
		AIntField    *int     `json:"a_int_field"`
		ABoolField   *bool    `json:"a_bool_field"`
		AFloatField  *float32 `json:"a_float_field"`
		ADoubleField *float64 `json:"a_double_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_str_field", "type": "string"},
      {"name": "a_int_field", "type": "int"},
      {"name": "a_bool_field", "type": "boolean"},
      {"name": "a_float_field", "type": "float"},
      {"name": "a_double_field", "type": "double"}
    ]
  }`

	e := Entity{}

	// test for instance
	r1, err1 := Reflect(e)
	assert.JSONEq(t, expected, r1)
	assert.Nil(t, err1)

	// test for pointer
	r2, err2 := Reflect(&e)
	assert.JSONEq(t, expected, r2)
	assert.Nil(t, err2)
}

func TestArrayOfPrimitive(t *testing.T) {
	type Entity struct {
		ArrayStrField []string `json:"a_str_array_field"`
		ArrayIntField []int    `json:"a_int_array_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_str_array_field", "type": {"type": "array", "items": "string"}},
      {"name": "a_int_array_field", "type": {"type": "array", "items": "int"}}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestArrayOfPrimitivePointer(t *testing.T) {
	type Entity struct {
		ArrayStrField []*string `json:"a_str_array_field"`
		ArrayIntField []*int    `json:"a_int_array_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_str_array_field", "type": {"type": "array", "items": "string"}},
      {"name": "a_int_array_field", "type": {"type": "array", "items": "int"}}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestArrayOfObject(t *testing.T) {
	type Foo struct {
		Bar string `json:"bar"`
	}
	type Entity struct {
		ArrayObjectField []Foo  `json:"a_obj_array_field"`
		ArrayObjPtrField []*Foo `json:"a_obj_ptr_array_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_obj_array_field", "type": {
	    "type": "array", "items": {
          "name": "Foo", "type": "record", "fields": [{"name": "bar", "type": "string"}]
        }
	  }},
      {"name": "a_obj_ptr_array_field", "type": {
	    "type": "array", "items": {
          "name": "Foo", "type": "record", "fields": [{"name": "bar", "type": "string"}]
        }
	  }}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)

	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestMapOfPrimitive(t *testing.T) {
	type Entity struct {
		MapStrField map[string]string `json:"a_str_map_field"`
		MapIntField map[string]int    `json:"a_int_map_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_str_map_field", "type": {"type":"map", "values": "string"}},
      {"name": "a_int_map_field", "type": {"type":"map", "values": "int"}}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestInvalidMap(t *testing.T) {
	type Entity struct {
		MapStrField map[int]string `json:"a_invalid_map_field"`
		MapIntField map[string]int `json:"a_int_map_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_invalid_map_field", "type": "string"},
      {"name": "a_int_map_field", "type": {"type":"map", "values": "int"}}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestMapOfArray(t *testing.T) {
	type Foo struct {
		Bar string `json:"bar"`
	}
	type Entity struct {
		MapArrayField map[string][]Foo `json:"a_array_map_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_array_map_field", "type": {"type":"map", "values": {
        "type": "array", "items": {
					"name": "Foo", "type": "record", "fields": [{"name": "bar", "type": "string"}]
				}
      }}}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestInvalidMapInMap(t *testing.T) {
	type Foo struct {
		Bar string `json:"bar"`
	}
	type Entity struct {
		MapArrayField map[string]map[int]Foo `json:"a_array_map_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_array_map_field", "type": {"type":"map", "values": "string"}}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

// shouldn't the pointer field be marked optional?
func TestTimeType(t *testing.T) {
	type Entity struct {
		TimeField1 time.Time  `json:"time_field_1"`
		TimeField2 *time.Time `json:"time_field_2"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "time_field_1", "type": "long", "logicalType": "timestamp-millis"},
      {"name": "time_field_2", "type": "long", "logicalType": "timestamp-millis"}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestMapperToString(t *testing.T) {
	type Entity struct {
		ArrayField []int          `json:"a_int_array_field"`
		MapField   map[string]int `json:"a_int_map_field"`
	}

	expected := `{
		"name": "Entity",
		"type": "record",
		"fields": [
			{"name": "a_int_array_field", "type": "string"},
			{"name": "a_int_map_field", "type": "string"}
		]
	}`

	e := Entity{}

	reflactor := new(Reflector)
	reflactor.Mapper = func(t reflect.Type) any {
		return "string"
	}

	r, err := reflactor.Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)

	reflactor.Mapper = func(t reflect.Type) any {
		return nil
	}

	expected2 := `{
		"name": "Entity",
		"type": "record",
		"fields": [
			{"name": "a_int_array_field", "type": {"type":"array", "items": "int"}},
			{"name": "a_int_map_field", "type": {"type":"map", "values": "int"}}
		]
	}`

	r2, err2 := reflactor.Reflect(e)
	assert.JSONEq(t, expected2, r2)
	assert.Nil(t, err2)
}

func TestUnionPrimitiveType(t *testing.T) {
	type Entity struct {
		UnionField int `json:"union_field,omitempty"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "union_field", "type": ["null", "int"]}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestUnionRecordType(t *testing.T) {
	type Foo struct {
		Bar string `json:"bar"`
	}
	type Entity struct {
		UnionField Foo `json:"union_field,omitempty"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "union_field", "type": ["null", {
        "name": "Foo", "type": "record", "fields": [{"name": "bar", "type": "string"}]
      }]}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestUnionArrayType(t *testing.T) {
	type Entity struct {
		UnionArrayField []int `json:"union_array_field,omitempty"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "union_array_field", "type": ["null", {
        "type": "array", "items": "int"
      }]}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestBackwardTransitive(t *testing.T) {
	type Entity struct {
		AStrField    string  `json:"a_str_field"`
		AIntField    int     `json:"a_int_field"`
		ABoolField   bool    `json:"a_bool_field"`
		AFloatField  float32 `json:"a_float_field"`
		ADoubleField float64 `json:"a_double_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_str_field", "type": ["null", "string"]},
      {"name": "a_int_field", "type": ["null", "int"]},
      {"name": "a_bool_field", "type": ["null", "boolean"]},
      {"name": "a_float_field", "type": ["null", "float"]},
      {"name": "a_double_field", "type": ["null", "double"]}
    ]
  }`

	e := Entity{}

	// assign flag
	reflector := new(Reflector)
	reflector.BeBackwardTransitive = true

	r, err := reflector.Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestInterfaceOfMap(t *testing.T) {
	type Entity struct {
		AMapInterfaceField map[string]any `json:"a_map_interface_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_map_interface_field", "type": {"type":"map", "values": "string"}}
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)

}

func TestInlineRecordType(t *testing.T) {
	type Foo struct {
		Bar string `json:"bar"`
	}
	type Entity struct {
		Foo         `json:",inline"`
		Embedded    Foo  `json:"embedded"`
		EmbeddedOpt *Foo `json:"embedded_opt,omitempty"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
	  {"name": "bar", "type": "string"},
	  {"name": "embedded", "type":
	    {
          "name": "Foo", "type": "record", "fields": [{"name": "bar", "type": "string"}]
        }
      },
	  {"name": "embedded_opt", "type":
	    [
		  "null",
	      {
            "name": "Foo", "type": "record", "fields": [{"name": "bar", "type": "string"}]
          }
		]
      }
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}

func TestNestedRecordType(t *testing.T) {
	type Foo struct {
		Bar string `json:"bar"`
	}
	type Entity struct {
		EmbeddedField Foo `json:"embedded_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "embedded_field", "type":
	    {
          "name": "Foo", "type": "record", "fields": [{"name": "bar", "type": "string"}]
        }
      }
    ]
  }`

	e := Entity{}

	r, err := Reflect(e)
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}
