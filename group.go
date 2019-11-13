package enumerable

import (
	"reflect"
)

// Group groups the slice by the return value of `keyFunc`. It returns
// a map where the keys are the return values of `keyFunc`, and the
// values are slices of the same type as the given slice
func Group(slice AnySlice, keyFunc IterFunc) AnyValue {
	mustBeSlice(slice)
	mustBeIterFuncForSlice(keyFunc, slice)

	fVal := reflect.ValueOf(keyFunc)
	sVal := reflect.ValueOf(slice)
	sType := sVal.Type()

	mapType := reflect.MapOf(fVal.Type().Out(0), sType)
	groups := reflect.MakeMap(mapType)

	count := sVal.Len()
	var notPresent reflect.Value
	for i := 0; i < count; i++ {
		item := sVal.Index(i)
		key := fVal.Call([]reflect.Value{item})[0]
		list := groups.MapIndex(key)
		if list == notPresent {
			list = reflect.MakeSlice(sType, 0, count)
		}
		groups.SetMapIndex(key, reflect.Append(list, item))
	}
	return groups.Interface()
}
