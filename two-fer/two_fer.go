package twofer

import (
	"fmt"
	"strings"
)

func ShareWith(name string) string {
	name = strings.Trim(name, " ")

	if name == "" {
		name = "you"
	} else {
		name = strings.Title(name)
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
