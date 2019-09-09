package tools

import (
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

func Test_Abs(t *testing.T) {
	if Abs(-100) != 100 {
		t.Fatal("代码错误")
	}
	if Abs(100) != 100 {
		t.Fatal("代码错误")
	}
}

func Test_Now32(t *testing.T) {
	Now32()
}
