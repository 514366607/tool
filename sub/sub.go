package sub

import (
	"math"
)

// SubUint64 相加|相减
func SubUint64(a uint64, b int64) uint64 {
	if b < 0 {
		// 相减
		b = b - b*2 // 变成正数
		if int64(a) <= b {
			return 0
		}
		return a - uint64(b)
	}

	// 相加,防止溢出
	if a+uint64(b) < a {
		return math.MaxUint64
	}
	return a + uint64(b)
}

// SubInt8 相加|相减
func SubInt8(a int8, b int8) int8 {
	res := a + b
	if a > 0 && b > 0 && res < 0 {
		res = math.MaxInt8
	} else if a < 0 && b < 0 && res > 0 {
		res = math.MinInt8
	}
	return res
}

// SubInt 相加|相减
func SubInt(a int, b int) int {
	return int(SubInt32(int32(a), int32(b)))
}

// SubInt32 相加|相减
func SubInt32(a int32, b int32) int32 {
	res := a + b
	if a > 0 && b > 0 && res < 0 {
		res = math.MaxInt32
	} else if a < 0 && b < 0 && res > 0 {
		res = math.MinInt32
	}
	return res
}

// SubInt64 相加|相减
func SubInt64(a int64, b int64) int64 {
	res := a + b
	if a > 0 && b > 0 && res < 0 {
		res = math.MaxInt64
	} else if a < 0 && b < 0 && res > 0 {
		res = math.MinInt64
	}
	return res
}
