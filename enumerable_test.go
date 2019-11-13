package enumerable

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"
)

var numberWords = []string{"one", "two", "one", "three"}

func expectEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected `%#v` but got `%#v`", expected, actual)
	}
}

func TestMap(t *testing.T) {
	var items []int
	New(numberWords).Map(func(item string) int {
		return len(item)
	}).GetSlice(&items)
	expectEqual(t, []int{3, 3, 3, 5}, items)
}

func TestFilter(t *testing.T) {
	var items []string
	New(numberWords).Filter(func(item string) bool {
		return len(item) > 3
	}).GetSlice(&items)
	expectEqual(t, []string{"three"}, items)
}

func TestReduce(t *testing.T) {
	str := New(numberWords).Reduce("", func(memo, item string) string {
		return memo + item
	}).(string)
	expectEqual(t, "onetwoonethree", str)
}

func TestReduceToInt(t *testing.T) {
	sumLengths := New(numberWords).Map(func(item string) int {
		return len(item)
	}).ReduceToInt(0, func(memo int, item int) int {
		return memo + item
	})
	expectEqual(t, 14, sumLengths)
}
func TestReduceToFloat64(t *testing.T) {
	sumLengths := New(numberWords).Map(func(item string) int {
		return len(item)
	}).ReduceToFloat64(0.0, func(memo float64, item int) float64 {
		return memo + float64(item)
	})
	expectEqual(t, 14.0, sumLengths)
}

func TestFirst(t *testing.T) {
	var firstTwo []string
	New(numberWords).First(2).GetSlice(&firstTwo)
	expectEqual(t, []string{"one", "two"}, firstTwo)
}

func TestLast(t *testing.T) {
	var lastTwo []string
	New(numberWords).Last(2).GetSlice(&lastTwo)
	expectEqual(t, []string{"one", "three"}, lastTwo)
}

func TestFind(t *testing.T) {
	idx, found := New(numberWords).Find(func(item string) bool {
		return item == "three"
	})
	expectEqual(t, 3, idx)
	expectEqual(t, "three", found)
}

func TestAny(t *testing.T) {
	found := New(numberWords).Any(func(item string) bool {
		return item == "three"
	})
	expectEqual(t, true, found)

	found = New(numberWords).Any(func(item string) bool {
		return item == "eighteen"
	})
	expectEqual(t, false, found)
}

func TestUniq(t *testing.T) {
	var uniqStrings []string
	New(numberWords).Uniq().GetSlice(&uniqStrings)
	expectEqual(t, []string{"one", "two", "three"}, uniqStrings)
}

func TestFlatten(t *testing.T) {
	var letters []rune
	New(numberWords).Map(func(word string) []rune {
		return []rune(word)
	}).Flatten().GetSlice(&letters)
	expectEqual(t, []rune("onetwoonethree"), letters)
}

func TestSortBy(t *testing.T) {
	words := []string{"four", "thirteen", "one"}
	New(words).SortBy(func(word string) int {
		return len(word)
	}).GetSlice(&words)
	expectEqual(t, []string{"one", "four", "thirteen"}, words)
}

func TestReverse(t *testing.T) {
	var reversed []string
	numbersCopy := make([]string, len(numberWords))
	copy(numbersCopy, numberWords)
	New(numbersCopy).Reverse().GetSlice(&reversed)
	expectEqual(t, []string{"three", "one", "two", "one"}, reversed)
}

func TestMin(t *testing.T) {
	expectEqual(t, "one", New(numberWords).Min().(string))
}

func TestMax(t *testing.T) {
	expectEqual(t, "two", New(numberWords).Max().(string))
}

func TestGroup(t *testing.T) {
	groups := New(numberWords).Group(func(word string) int {
		return len(word)
	})
	expected := map[int][]string{
		3: []string{"one", "two", "one"},
		5: []string{"three"},
	}
	expectEqual(t, expected, groups)
}

func TestJoinAsString(t *testing.T) {
	expectEqual(t, "one, two, one, three", New(numberWords).JoinAsString(", "))

	floats := []float64{1.237, 2.672}
	joined := New(floats).JoinAsStringWithFormat("%.2f", ",")
	expectEqual(t, "1.24,2.67", joined)

	nonAlphabet := regexp.MustCompile("[^a-zA-Z]")
	files, _ := filepath.Glob("*.go")
	out := New(files).Map(func(file string) []string {
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		return nonAlphabet.Split(string(contents), -1)
	}).Flatten().Uniq().Sort().JoinAsString(", ")
	t.Error(out)
}

// ********* BENCHMARKS **********

var benchData = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
}

var benchDataInts = func() []int {
	length := 10000
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}
	return arr
}()

var benchFunc = func(n int) int {
	return n * 2
}

var stringLen = func(value string) int {
	return len(value)
}

type genericInt int

func BenchmarkMap(b *testing.B) {
	e := New(benchDataInts).Map(func(n int) genericInt {
		return genericInt(n)
	})
	for n := 0; n < b.N; n++ {
		e.Map(func(n genericInt) genericInt {
			return n * 2
		})
	}
}
func BenchmarkOptimizedMap(b *testing.B) {
	e := New(benchDataInts)
	for n := 0; n < b.N; n++ {
		e.Map(benchFunc)
	}
}

func BenchmarkGenericMap(b *testing.B) {
	e := New(benchDataInts).ToGeneric()
	for n := 0; n < b.N; n++ {
		e.GenericMap(func(item AnyValue) AnyValue {
			return item.(int) * 2
		})
	}
}

func BenchmarkMapHandrolled(b *testing.B) {
	for n := 0; n < b.N; n++ {
		output := make([]int, len(benchDataInts), len(benchDataInts))
		for i := 0; i < len(benchDataInts); i++ {
			output[i] = benchFunc(benchDataInts[i])
		}
	}
}

func BenchmarkReduce(b *testing.B) {
	genericInts := Map(benchDataInts, func(i int) genericInt {
		return genericInt(i)
	})
	for n := 0; n < b.N; n++ {
		Reduce(genericInts, 0, func(memo int, item genericInt) int {
			return memo + int(item)*2
		})
	}
}

func BenchmarkOptimizedReduce(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Reduce(benchDataInts, 0, func(memo, item int) int {
			return memo + item*2
		})
	}
}
