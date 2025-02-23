package mongo

import (
	"testing"

	"github.com/kamva/mgm/v3"
	"github.com/stretchr/testify/assert"
	"github.com/wirelessr/avroschema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMgmCommonTypes(t *testing.T) {
	type Book struct {
		// DefaultModel adds _id, created_at and updated_at fields to the Model.
		mgm.DefaultModel `bson:",inline"`
		Name             string             `json:"name" bson:"name"`
		Pages            int                `json:"pages" bson:"pages"`
		ObjId            primitive.ObjectID `json:"obj_id" bson:"obj_id"`
		ArrivedAt        primitive.DateTime `json:"arrived_at" bson:"arrived_at"`
		RefData          bson.M             `json:"ref_data" bson:"ref_data"`
		Author           []string           `json:"author" bson:"author"`
	}

	reflector := new(avroschema.Reflector)
	reflector.Mapper = MgmExtension

	expected := `{
		"name": "Book",
		"type": "record",
		"fields": [
			{ "name": "_id", "type": ["null", "string"] },
			{ "name": "created_at", "type": "long", "logicalType": "timestamp-millis" },
			{ "name": "updated_at", "type": "long", "logicalType": "timestamp-millis" },
			{ "name": "name", "type": "string" },
			{ "name": "pages", "type": "int" },
			{ "name": "obj_id", "type": "string" },
			{ "name": "arrived_at", "type": "long", "logicalType": "timestamp-millis" },
			{ "name": "ref_data", "type": "string" },
			{ "name": "author", "type": { "type": "array", "items": "string" }}
		]
	}`

	r, err := reflector.Reflect(Book{})
	assert.JSONEq(t, expected, r)
	assert.Nil(t, err)
}
