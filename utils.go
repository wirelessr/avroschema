package avroschema

import "strings"

type structTag struct {
	Name     string
	Optional bool
	Inline   bool
}

func parseStructTag(tag string) *structTag {
	tags := strings.Split(tag, ",")
	name := tags[0]
	optional := false
	inline := false

	for _, tag := range tags {
		switch tag {
		case "omitempty":
			optional = true

		case "inline":
			inline = true
		}
	}
	return &structTag{name, optional, inline}
}
