package enumerable

import (
	"fmt"
	"reflect"
)

// IterFunc represents a function which will be called when iterating through a slice
type IterFunc interface{}

// AnySlice represents a slice with elements of any type
type AnySlice interface{}

// AnyValue represents any value
type AnyValue interface{}

func mustBeFunction(f IterFunc) reflect.Type {
	funcType := reflect.TypeOf(f)
	if funcType.Kind() != reflect.Func {
		panic(fmt.Errorf("enumerable: iterFunc isn't a function"))
	}
	return funcType
}

func mustBeIterFuncForSlice(f IterFunc, slice AnySlice) {
	funcType := mustBeFunction(f)

	sliceElemType := reflect.TypeOf(slice).Elem()
	if funcType.NumIn() != 1 || funcType.In(0) != sliceElemType {
		panic(fmt.Errorf("enumerable: iterFunc parameter type should be `%v`", sliceElemType))
	}

	if funcType.NumOut() != 1 {
		panic(fmt.Errorf("enumerable: iterFunc should return a value"))
	}
}

func mustBeIterFuncForSliceWithMemo(f IterFunc, memo interface{}, slice AnySlice) {
	funcType := mustBeFunction(f)

	memoType := reflect.TypeOf(memo)
	sliceElemType := reflect.TypeOf(slice).Elem()
	if funcType.NumIn() != 2 || funcType.In(0) != memoType || funcType.In(1) != sliceElemType {
		panic(fmt.Errorf("enumerable: iterFunc parameter types should be `%v,%v`", memoType, sliceElemType))
	}

	if funcType.NumOut() != 1 || funcType.Out(0) != memoType {
		panic(fmt.Errorf("enumerable: iterFunc return type should be `%v`", memoType))
	}
}

func mustBeIterFuncFilterForSlice(f IterFunc, slice AnySlice) {
	funcType := mustBeFunction(f)

	sliceElemType := reflect.TypeOf(slice).Elem()
	if funcType.NumIn() != 1 || funcType.In(0) != sliceElemType {
		panic(fmt.Errorf("enumerable: iterFunc parameter type should be `%v`", sliceElemType))
	}

	if funcType.NumOut() != 1 || funcType.Out(0).Kind() != reflect.Bool {
		panic(fmt.Errorf("enumerable: iterFunc should return a `bool`"))
	}
}

func mustBeSlice(slice AnySlice) {
	sliceType := reflect.TypeOf(slice)
	if sliceType.Kind() != reflect.Slice {
		panic(fmt.Errorf("enumerable: expected a slice but got a `%v`", sliceType))
	}
}

func makeOutputSlice(f IterFunc, len, cap int) reflect.Value {
	elemType := reflect.TypeOf(f).Out(0)
	return reflect.MakeSlice(reflect.SliceOf(elemType), len, cap)
}

func makeFilterSlice(slice AnySlice, len, cap int) reflect.Value {
	elemType := reflect.TypeOf(slice).Elem()
	return reflect.MakeSlice(reflect.SliceOf(elemType), len, cap)
}
