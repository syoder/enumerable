package enumerable

import (
	"fmt"
	"reflect"
)

// JoinAsString returns a string with each item converted to string
// and joined with `sep`.
func JoinAsString(slice AnySlice, sep string) string {
	return JoinAsStringWithFormat(slice, "%v", sep)
}

// JoinAsStringWithFormat returns a string with each item converted to string
// using the given format string (see fmt.Printf) and joined with `sep`.
func JoinAsStringWithFormat(slice AnySlice, format, sep string) string {
	mustBeSlice(slice)
	sVal := reflect.ValueOf(slice)
	count := sVal.Len()

	var output string
	for i := 0; i < count; i++ {
		item := sVal.Index(i).Interface()
		if i == 0 {
			output += fmt.Sprintf(format, item)
		} else {
			output += fmt.Sprintf("%s"+format, sep, item)
		}
	}
	return output
}
