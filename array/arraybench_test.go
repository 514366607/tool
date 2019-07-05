package array

import (
	"testing"
)

//go test -test.bench=".*"
func Benchmark_SliceIN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		flag := IN(4, []int{1, 2, 3, 4})
		if flag == false {
			b.Error("Run False")
		}
	}
}

func Benchmark_MapIN(b *testing.B) {
	var testMap = make(map[string]int8, 0)
	testMap["test1"] = 1
	testMap["test2"] = 2
	testMap["test3"] = 3
	testMap["test4"] = 4
	for i := 0; i < b.N; i++ {
		flag := IN(int8(4), testMap)
		if flag == false {
			b.Error("Run False")
		}
	}
}

func Benchmark_SliceSum(b *testing.B) {
	var testMap = []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		if sum, err := Sum(testMap); sum != 10 || err != nil {
			b.Error("Run False", sum, err)
		}
	}
}

func Benchmark_MapSum(b *testing.B) {
	var testMap = make(map[string]int8, 0)
	testMap["test1"] = 1
	testMap["test2"] = 2
	testMap["test3"] = 3
	testMap["test4"] = 4
	for i := 0; i < b.N; i++ {
		if sum, err := Sum(testMap); sum != 10 || err != nil {
			b.Error("Run False", sum, err)
		}
	}
}
