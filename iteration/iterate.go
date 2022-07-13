package iteration

import "strings"

func Iterate(char string, count int) (repeated string) {
	repeated = strings.Repeat(char, count)
	return
}
