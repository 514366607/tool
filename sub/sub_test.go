package sub

import (
	"math"
	"testing"
)

func Test_SubUint64(t *testing.T) {
	var a = SubUint64(1, 2)
	if a != 3 {
		t.Error("添加失败", a)
		return
	}

	a = SubUint64(1, -2)
	if a != 0 {
		t.Error("减少失败", a)
		return
	}

	a = SubUint64(3, -2)
	if a != 1 {
		t.Error("减少失败", a)
		return
	}

	a = SubUint64(math.MaxInt64+1, math.MaxInt64)
	if a != math.MaxUint64 {
		t.Error("添加失败", a)
		return
	}

	a = SubUint64(math.MaxInt64+2, math.MaxInt64)
	if a != math.MaxUint64 {
		t.Error("添加失败", a)
		return
	}

	a = SubUint64(math.MaxInt64, math.MaxInt64)
	if a == math.MaxUint64 {
		t.Error("添加失败", a)
		return
	}

}

func Test_SubInt32(t *testing.T) {
	var a = SubInt32(1, 2)
	if a != 3 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt32(3, -2)
	if a != 1 {
		t.Error("减少失败", a)
		return
	}

	a = SubInt32(-3, -2)
	if a != -5 {
		t.Error("减少失败", a)
		return
	}

	a = SubInt32(math.MaxInt32, math.MaxInt32)
	if a != math.MaxInt32 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt32(math.MaxInt32, int32(2))
	if a != math.MaxInt32 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt32(math.MinInt32, -3)
	if a != math.MinInt32 {
		t.Error("添加失败", a)
		return
	}

}

func Test_SubInt8(t *testing.T) {
	var a = SubInt8(1, 2)
	if a != 3 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt8(3, -2)
	if a != 1 {
		t.Error("减少失败", a)
		return
	}

	a = SubInt8(-3, -2)
	if a != -5 {
		t.Error("减少失败", a)
		return
	}

	a = SubInt8(math.MaxInt8, math.MaxInt8)
	if a != math.MaxInt8 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt8(math.MaxInt8, 2)
	if a != math.MaxInt8 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt8(math.MinInt8, -3)
	if a != math.MinInt8 {
		t.Error("添加失败", a)
		return
	}

}

func Test_SubInt(t *testing.T) {
	var a = SubInt(1, 2)
	if a != 3 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt(3, -2)
	if a != 1 {
		t.Error("减少失败", a)
		return
	}

	a = SubInt(-3, -2)
	if a != -5 {
		t.Error("减少失败", a)
		return
	}

	a = SubInt(math.MaxInt32, math.MaxInt32)
	if a != math.MaxInt32 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt(math.MaxInt32, 2)
	if a != math.MaxInt32 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt(math.MinInt32, -3)
	if a != math.MinInt32 {
		t.Error("添加失败", a)
		return
	}

}

func Test_SubInt64(t *testing.T) {
	var a = SubInt64(1, 2)
	if a != 3 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt64(3, -2)
	if a != 1 {
		t.Error("减少失败", a)
		return
	}

	a = SubInt64(math.MaxInt64, math.MaxInt64)
	if a != math.MaxInt64 {
		t.Error("添加失败", a)
		return
	}

	a = SubInt64(math.MaxInt64, int64(2))
	if a != math.MaxInt64 {
		t.Error("添加失败", a)
		return
	}

}
