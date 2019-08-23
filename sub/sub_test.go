package sub

import (
	"math"
	"testing"
)

func Test_Int8(t *testing.T) {
	var a = Int8(1, 2)
	if a != 3 {
		t.Fatal("添加失败", a)
		return
	}

	a = Int8(3, -2)
	if a != 1 {
		t.Fatal("减少失败", a)
		return
	}

	a = Int8(-3, -2)
	if a != -5 {
		t.Fatal("减少失败", a)
		return
	}

	a = Int8(math.MaxInt8, math.MaxInt8)
	if a != math.MaxInt8 {
		t.Fatal("添加失败", a)
		return
	}

	a = Int8(math.MaxInt8, 2)
	if a != math.MaxInt8 {
		t.Fatal("添加失败", a)
		return
	}

	a = Int8(math.MinInt8, -3)
	if a != math.MinInt8 {
		t.Fatal("添加失败", a)
		return
	}
}

func Test_Uint8(t *testing.T) {
	var a = Uint8(1, 2)
	if a != 3 {
		t.Fatal("添加失败", a)
		return
	}

	a = Uint8(1, -2)
	if a != 0 {
		t.Fatal("减少失败", a)
		return
	}

	a = Uint8(3, -2)
	if a != 1 {
		t.Fatal("减少失败", a)
		return
	}

	a = Uint8(math.MaxInt8+1, math.MaxInt8)
	if a != math.MaxUint8 {
		t.Fatal("添加失败", a)
		return
	}

	a = Uint8(math.MaxInt8+2, math.MaxInt8)
	if a != math.MaxUint8 {
		t.Fatal("添加失败", a)
		return
	}

	a = Uint8(math.MaxInt8, math.MaxInt8)
	if a == math.MaxUint8 {
		t.Fatal("添加失败", a)
		return
	}
}

func Test_Int(t *testing.T) {
	var a = Int(1, 2)
	if a != 3 {
		t.Fatal("添加失败", a)
		return
	}

	a = Int(3, -2)
	if a != 1 {
		t.Fatal("减少失败", a)
		return
	}

	a = Int(-3, -2)
	if a != -5 {
		t.Fatal("减少失败", a)
		return
	}

	a = Int(math.MaxInt32, math.MaxInt32)
	if a != math.MaxInt32 {
		t.Fatal("添加失败", a)
		return
	}

	a = Int(math.MaxInt32, 2)
	if a != math.MaxInt32 {
		t.Fatal("添加失败", a)
		return
	}

	a = Int(math.MinInt32, -3)
	if a != math.MinInt32 {
		t.Fatal("添加失败", a)
		return
	}
}

func Test_Uint(t *testing.T) {
	var a = Uint(1, 2)
	if a != 3 {
		t.Fatal("添加失败", a)
		return
	}

	a = Uint(1, -2)
	if a != 0 {
		t.Fatal("减少失败", a)
		return
	}

	a = Uint(3, -2)
	if a != 1 {
		t.Fatal("减少失败", a)
		return
	}

	a = Uint(math.MaxInt32+1, math.MaxInt32)
	if a != math.MaxUint32 {
		t.Fatal("添加失败", a)
		return
	}

	a = Uint(math.MaxInt32+2, math.MaxInt32)
	if a != math.MaxUint32 {
		t.Fatal("添加失败", a)
		return
	}

	a = Uint(math.MaxInt32, math.MaxInt32)
	if a == math.MaxUint32 {
		t.Fatal("添加失败", a)
		return
	}
}

func Test_Int64(t *testing.T) {
	var a = Int64(1, 2)
	if a != 3 {
		t.Fatal("添加失败", a)
		return
	}

	a = Int64(3, -2)
	if a != 1 {
		t.Fatal("减少失败", a)
		return
	}

	a = Int64(-3, -2)
	if a != -5 {
		t.Fatal("减少失败", a)
		return
	}

	a = Int64(math.MaxInt64, math.MaxInt64)
	if a != math.MaxInt64 {
		t.Fatal("添加失败", a)
		return
	}

	a = Int64(math.MaxInt64, 2)
	if a != math.MaxInt64 {
		t.Fatal("添加失败", a)
		return
	}

	a = Int64(math.MinInt64, -3)
	if a != math.MinInt64 {
		t.Fatal("添加失败", a)
		return
	}
}

func Test_Uint64(t *testing.T) {
	var a = Uint64(1, 2)
	if a != 3 {
		t.Fatal("添加失败", a)
		return
	}

	a = Uint64(1, -2)
	if a != 0 {
		t.Fatal("减少失败", a)
		return
	}

	a = Uint64(3, -2)
	if a != 1 {
		t.Fatal("减少失败", a)
		return
	}

	a = Uint64(math.MaxInt64+1, math.MaxInt64)
	if a != math.MaxUint64 {
		t.Fatal("添加失败", a)
		return
	}

	a = Uint64(math.MaxInt64+2, math.MaxInt64)
	if a != math.MaxUint64 {
		t.Fatal("添加失败", a)
		return
	}

	a = Uint64(math.MaxInt64, math.MaxInt64)
	if a == math.MaxUint64 {
		t.Fatal("添加失败", a)
		return
	}

}
