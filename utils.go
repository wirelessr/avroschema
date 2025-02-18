package avroschema

import "strings"

func GetNameAndOmit(jsonTag string) (string, bool, bool) {
	tags := strings.Split(jsonTag, ",")
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
	return name, optional, inline
}
