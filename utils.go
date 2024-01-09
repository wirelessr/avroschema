package avroschema

import "strings"

func GetNameAndOmit(jsonTag string) (string, bool) {
	tags := strings.Split(jsonTag, ",")
	name := tags[0]

	for _, tag := range tags {
		if tag == "omitempty" {
			return name, true
		}
	}
	return name, false
}
