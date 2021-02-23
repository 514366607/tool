package tools

import (
	"math"
	"testing"
)

func Test_FormatStringToInt(t *testing.T) {
	var s = "1"
	if a, _ := FormatStringToInt(s); a != 1 {
		t.Error("格式化错误")
	}

	s = "-2"
	if a, _ := FormatStringToInt(s); a != -2 {
		t.Error("格式化错误")
	}
}

func Test_FormatStringToFloat(t *testing.T) {
	var s = "1.1111111"
	if a, _ := FormatStringToFloat(s); a != 1.1111111 {
		t.Error("格式化错误")
	}

	s = "-2.222222222222222"
	if a, _ := FormatStringToFloat(s); a != -2.222222222222222 {
		t.Error("格式化错误")
	}
}

func Test_WeightRand(t *testing.T) {
	var weight = map[int]int{1: 10000, 2: 1}
	if WeightRand(weight) != 1 {
		t.Fatal("概率太低不可能")
	}
}

func Test_Now32(t *testing.T) {
	Now32()
}

func Test_Abs(t *testing.T) {
	if Abs(-1) != 1 {
		t.Fatal("Not Same")
	}
	if Abs(0) != 0 {
		t.Fatal("Not Same")
	}
	if Abs(-math.MaxInt32) != math.MaxInt32 {
		t.Fatal("Not Same")
	}
	if Abs(math.MaxInt32) != math.MaxInt32 {
		t.Fatal("Not Same")
	}
	if Abs(-math.MaxInt64) != math.MaxInt64 {
		t.Fatal("Not Same")
	}
	if Abs(math.MaxInt64) != math.MaxInt64 {
		t.Fatal("Not Same")
	}
}

/*
这两个测试说明了直接用math的转来转去也可以用，只是使用麻烦点，一丁丁的性能差异可以忽略
cpu: Intel(R) Core(TM) i5-8500 CPU @ 3.00GHz
Benchmark_Abs-6       	1000000000	         0.2612 ns/op	       0 B/op	       0 allocs/op
Benchmark_MathAbs-6   	1000000000	         0.2630 ns/op	       0 B/op	       0 allocs/op
*/

func Benchmark_Abs(b *testing.B) {
	var testInt64 int64 = -1
	for i := 0; i < b.N; i++ {
		if Abs(testInt64) != -testInt64 {
			b.Fatal("Not Same")
		}
	}
}

func Benchmark_MathAbs(b *testing.B) {
	var testInt64 int64 = -1
	for i := 0; i < b.N; i++ {
		if int64(math.Abs(float64(testInt64))) != -testInt64 {
			b.Fatal("Not Same")
		}
	}
}
