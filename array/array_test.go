package array

import (
	"testing"
)

func TestIN(t *testing.T) {
	var a int8 = 1
	var b = []int8{1, 2}
	if IN(a, b) == false {
		t.Error("判断错误")
	}

	var c int8 = 1
	var d = []int{1, 2}
	if IN(c, d) == true {
		t.Error("判断错误")
	}

	type testStruct struct {
		Name string
	}

	var e = testStruct{Name: "test1"}
	var f = testStruct{Name: "test2"}
	var g = []testStruct{e, f}
	if IN(e, g) == false {
		t.Error("判断错误")
	}
	e.Name = "test3"
	if IN(e, g) == true {
		t.Error("判断错误")
	}
	var h = &testStruct{Name: "test1"}
	var i = &testStruct{Name: "test2"}
	var j = []*testStruct{h, i}
	if IN(h, j) == false {
		t.Error("判断错误")
	}
	h.Name = "test3"
	if IN(h, j) == false {
		t.Error("判断错误")
	}
	if IN(e, j) == true {
		t.Error("判断错误")
	}

	var testMap = make(map[int]int8, 0)
	testMap[1] = a
	if IN(a, testMap) == false {
		t.Error("判断错误")
	}

}

func TestSum(t *testing.T) {
	var a0 = []int{-1, -2}
	if a, err := Sum(a0); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var a1 = []int8{-1, -2}
	if a, err := Sum(a1); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var a2 = []int16{-1, -2}
	if a, err := Sum(a2); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var a3 = []int32{-1, -2}
	if a, err := Sum(a3); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var a4 = []int64{-1, -2}
	if a, err := Sum(a4); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}

	var u0 = []uint{1, 2}
	if a, err := Sum(u0); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var u1 = []uint8{1, 2}
	if a, err := Sum(u1); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var u2 = []uint16{1, 2}
	if a, err := Sum(u2); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var u3 = []uint32{1, 2}
	if a, err := Sum(u3); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var u4 = []uint64{1, 2}
	if a, err := Sum(u4); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}

	var ma0 = make(map[int]int, 2)
	ma0[1] = -1
	ma0[2] = -2
	if a, err := Sum(ma0); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var ma1 = make(map[int]int8, 2)
	ma1[1] = -1
	ma1[2] = -2
	if a, err := Sum(ma1); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var ma2 = make(map[int]int16, 2)
	ma2[1] = -1
	ma2[2] = -2
	if a, err := Sum(ma2); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var ma3 = make(map[int]int32, 2)
	ma3[1] = -1
	ma3[2] = -2
	if a, err := Sum(ma3); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var ma4 = make(map[int]int64, 2)
	ma4[1] = -1
	ma4[2] = -2
	if a, err := Sum(ma4); a != -3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}

	var mu0 = make(map[int]uint, 2)
	mu0[1] = 1
	mu0[2] = 2
	if a, err := Sum(mu0); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var mu1 = make(map[int]uint8, 2)
	mu1[1] = 1
	mu1[2] = 2
	if a, err := Sum(mu1); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var mu2 = make(map[int]uint16, 2)
	mu2[1] = 1
	mu2[2] = 2
	if a, err := Sum(mu2); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var mu3 = make(map[int]uint32, 2)
	mu3[1] = 1
	mu3[2] = 2
	if a, err := Sum(mu3); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}
	var mu4 = make(map[int]uint64, 2)
	mu4[1] = 1
	mu4[2] = 2
	if a, err := Sum(mu4); a != 3 || err != nil {
		t.Error("取整失败", a, err)
		return
	}

}
