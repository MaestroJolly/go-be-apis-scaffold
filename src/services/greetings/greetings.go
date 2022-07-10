package greetings

import (
	"fmt"
)

func Greetings(name string) string {
	return fmt.Sprintf("Hello %v !!!", name)
}
