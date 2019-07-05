package common

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
