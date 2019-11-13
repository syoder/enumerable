package enumerable

// ReduceIntToInt reduces a slice of int to int
func ReduceIntToInt(in []int, memo int, f func(int, int) int) int {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceIntToFloat64 reduces a slice of int to float64
func ReduceIntToFloat64(in []int, memo float64, f func(float64, int) float64) float64 {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceIntToBool reduces a slice of int to bool
func ReduceIntToBool(in []int, memo bool, f func(bool, int) bool) bool {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceIntToString reduces a slice of int to string
func ReduceIntToString(in []int, memo string, f func(string, int) string) string {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceIntToAnyValue reduces a slice of int to AnyValue
func ReduceIntToAnyValue(in []int, memo AnyValue, f func(AnyValue, int) AnyValue) AnyValue {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceFloat64ToInt reduces a slice of float64 to int
func ReduceFloat64ToInt(in []float64, memo int, f func(int, float64) int) int {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceFloat64ToFloat64 reduces a slice of float64 to float64
func ReduceFloat64ToFloat64(in []float64, memo float64, f func(float64, float64) float64) float64 {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceFloat64ToBool reduces a slice of float64 to bool
func ReduceFloat64ToBool(in []float64, memo bool, f func(bool, float64) bool) bool {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceFloat64ToString reduces a slice of float64 to string
func ReduceFloat64ToString(in []float64, memo string, f func(string, float64) string) string {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceFloat64ToAnyValue reduces a slice of float64 to AnyValue
func ReduceFloat64ToAnyValue(in []float64, memo AnyValue, f func(AnyValue, float64) AnyValue) AnyValue {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceBoolToInt reduces a slice of bool to int
func ReduceBoolToInt(in []bool, memo int, f func(int, bool) int) int {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceBoolToFloat64 reduces a slice of bool to float64
func ReduceBoolToFloat64(in []bool, memo float64, f func(float64, bool) float64) float64 {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceBoolToBool reduces a slice of bool to bool
func ReduceBoolToBool(in []bool, memo bool, f func(bool, bool) bool) bool {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceBoolToString reduces a slice of bool to string
func ReduceBoolToString(in []bool, memo string, f func(string, bool) string) string {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceBoolToAnyValue reduces a slice of bool to AnyValue
func ReduceBoolToAnyValue(in []bool, memo AnyValue, f func(AnyValue, bool) AnyValue) AnyValue {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceStringToInt reduces a slice of string to int
func ReduceStringToInt(in []string, memo int, f func(int, string) int) int {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceStringToFloat64 reduces a slice of string to float64
func ReduceStringToFloat64(in []string, memo float64, f func(float64, string) float64) float64 {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceStringToBool reduces a slice of string to bool
func ReduceStringToBool(in []string, memo bool, f func(bool, string) bool) bool {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceStringToString reduces a slice of string to string
func ReduceStringToString(in []string, memo string, f func(string, string) string) string {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceStringToAnyValue reduces a slice of string to AnyValue
func ReduceStringToAnyValue(in []string, memo AnyValue, f func(AnyValue, string) AnyValue) AnyValue {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceAnyValueToInt reduces a slice of AnyValue to int
func ReduceAnyValueToInt(in []AnyValue, memo int, f func(int, AnyValue) int) int {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceAnyValueToFloat64 reduces a slice of AnyValue to float64
func ReduceAnyValueToFloat64(in []AnyValue, memo float64, f func(float64, AnyValue) float64) float64 {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceAnyValueToBool reduces a slice of AnyValue to bool
func ReduceAnyValueToBool(in []AnyValue, memo bool, f func(bool, AnyValue) bool) bool {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceAnyValueToString reduces a slice of AnyValue to string
func ReduceAnyValueToString(in []AnyValue, memo string, f func(string, AnyValue) string) string {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

// ReduceAnyValueToAnyValue reduces a slice of AnyValue to AnyValue
func ReduceAnyValueToAnyValue(in []AnyValue, memo AnyValue, f func(AnyValue, AnyValue) AnyValue) AnyValue {
	for _, value := range in {
		memo = f(memo, value)
	}
	return memo
}

var reduceFuncs = []interface{}{
	ReduceIntToInt,
	ReduceIntToFloat64,
	ReduceIntToBool,
	ReduceIntToString,
	ReduceIntToAnyValue,
	ReduceFloat64ToInt,
	ReduceFloat64ToFloat64,
	ReduceFloat64ToBool,
	ReduceFloat64ToString,
	ReduceFloat64ToAnyValue,
	ReduceBoolToInt,
	ReduceBoolToFloat64,
	ReduceBoolToBool,
	ReduceBoolToString,
	ReduceBoolToAnyValue,
	ReduceStringToInt,
	ReduceStringToFloat64,
	ReduceStringToBool,
	ReduceStringToString,
	ReduceStringToAnyValue,
	ReduceAnyValueToInt,
	ReduceAnyValueToFloat64,
	ReduceAnyValueToBool,
	ReduceAnyValueToString,
	ReduceAnyValueToAnyValue,
}
