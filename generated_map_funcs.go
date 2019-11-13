package enumerable

// MapIntToInt maps a slice of int to int
func MapIntToInt(in []int, f func(int) int) []int {
	out := make([]int, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapIntToFloat64 maps a slice of int to float64
func MapIntToFloat64(in []int, f func(int) float64) []float64 {
	out := make([]float64, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapIntToBool maps a slice of int to bool
func MapIntToBool(in []int, f func(int) bool) []bool {
	out := make([]bool, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapIntToString maps a slice of int to string
func MapIntToString(in []int, f func(int) string) []string {
	out := make([]string, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapIntToAnyValue maps a slice of int to AnyValue
func MapIntToAnyValue(in []int, f func(int) AnyValue) []AnyValue {
	out := make([]AnyValue, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapFloat64ToInt maps a slice of float64 to int
func MapFloat64ToInt(in []float64, f func(float64) int) []int {
	out := make([]int, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapFloat64ToFloat64 maps a slice of float64 to float64
func MapFloat64ToFloat64(in []float64, f func(float64) float64) []float64 {
	out := make([]float64, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapFloat64ToBool maps a slice of float64 to bool
func MapFloat64ToBool(in []float64, f func(float64) bool) []bool {
	out := make([]bool, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapFloat64ToString maps a slice of float64 to string
func MapFloat64ToString(in []float64, f func(float64) string) []string {
	out := make([]string, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapFloat64ToAnyValue maps a slice of float64 to AnyValue
func MapFloat64ToAnyValue(in []float64, f func(float64) AnyValue) []AnyValue {
	out := make([]AnyValue, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapBoolToInt maps a slice of bool to int
func MapBoolToInt(in []bool, f func(bool) int) []int {
	out := make([]int, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapBoolToFloat64 maps a slice of bool to float64
func MapBoolToFloat64(in []bool, f func(bool) float64) []float64 {
	out := make([]float64, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapBoolToBool maps a slice of bool to bool
func MapBoolToBool(in []bool, f func(bool) bool) []bool {
	out := make([]bool, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapBoolToString maps a slice of bool to string
func MapBoolToString(in []bool, f func(bool) string) []string {
	out := make([]string, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapBoolToAnyValue maps a slice of bool to AnyValue
func MapBoolToAnyValue(in []bool, f func(bool) AnyValue) []AnyValue {
	out := make([]AnyValue, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapStringToInt maps a slice of string to int
func MapStringToInt(in []string, f func(string) int) []int {
	out := make([]int, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapStringToFloat64 maps a slice of string to float64
func MapStringToFloat64(in []string, f func(string) float64) []float64 {
	out := make([]float64, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapStringToBool maps a slice of string to bool
func MapStringToBool(in []string, f func(string) bool) []bool {
	out := make([]bool, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapStringToString maps a slice of string to string
func MapStringToString(in []string, f func(string) string) []string {
	out := make([]string, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapStringToAnyValue maps a slice of string to AnyValue
func MapStringToAnyValue(in []string, f func(string) AnyValue) []AnyValue {
	out := make([]AnyValue, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapAnyValueToInt maps a slice of AnyValue to int
func MapAnyValueToInt(in []AnyValue, f func(AnyValue) int) []int {
	out := make([]int, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapAnyValueToFloat64 maps a slice of AnyValue to float64
func MapAnyValueToFloat64(in []AnyValue, f func(AnyValue) float64) []float64 {
	out := make([]float64, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapAnyValueToBool maps a slice of AnyValue to bool
func MapAnyValueToBool(in []AnyValue, f func(AnyValue) bool) []bool {
	out := make([]bool, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapAnyValueToString maps a slice of AnyValue to string
func MapAnyValueToString(in []AnyValue, f func(AnyValue) string) []string {
	out := make([]string, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

// MapAnyValueToAnyValue maps a slice of AnyValue to AnyValue
func MapAnyValueToAnyValue(in []AnyValue, f func(AnyValue) AnyValue) []AnyValue {
	out := make([]AnyValue, len(in))
	for i, value := range in {
		out[i] = f(value)
	}
	return out
}

var mapFuncs = []interface{}{
	MapIntToInt,
	MapIntToFloat64,
	MapIntToBool,
	MapIntToString,
	MapIntToAnyValue,
	MapFloat64ToInt,
	MapFloat64ToFloat64,
	MapFloat64ToBool,
	MapFloat64ToString,
	MapFloat64ToAnyValue,
	MapBoolToInt,
	MapBoolToFloat64,
	MapBoolToBool,
	MapBoolToString,
	MapBoolToAnyValue,
	MapStringToInt,
	MapStringToFloat64,
	MapStringToBool,
	MapStringToString,
	MapStringToAnyValue,
	MapAnyValueToInt,
	MapAnyValueToFloat64,
	MapAnyValueToBool,
	MapAnyValueToString,
	MapAnyValueToAnyValue,
}
