package avroschema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructArrayOfObject(t *testing.T) {
	e := AvroSchema{Name: "testRecord", Type: "record", Fields: []*AvroSchema{
		{Name: "testArray", Type: "array", Items: AvroSchema{
			Name: "testObject", Type: "record", Fields: []*AvroSchema{
				{Name: "testString", Type: "string"},
			}},
		},
	}}
	ret, err := StructToJson(e)
	expected := `{ 
		"type": "record", "name":"testRecord", "fields": [
			{
				"name": "testArray", "type": "array", "items": {
					"type": "record", "name": "testObject", "fields": [
						{ "name": "testString", "type": "string" }
					]
				}
			}
		]}`
	assert.JSONEq(t, expected, ret)
	assert.Nil(t, err)
}
