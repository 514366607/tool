package sub

import (
	"math"
)

// Int8 相加|相减
func Int8(a int8, b int8) int8 {
	res := a + b
	if a > 0 && b > 0 && res < 0 {
		res = math.MaxInt8
	} else if a < 0 && b < 0 && res > 0 {
		res = math.MinInt8
	}
	return res
}

// Uint8 相加|相减
func Uint8(a uint8, b int8) uint8 {
	if b < 0 {
		// 相减
		b = b - b*2 // 变成正数
		if a <= uint8(b) {
			return 0
		}
		return a - uint8(b)
	}

	// 相加,防止溢出
	var newNum = a + uint8(b)
	if newNum < a {
		return math.MaxUint8
	}
	return newNum
}

// Int 相加|相减
func Int(a int, b int) int {
	return int(Int32(int32(a), int32(b)))
}

// Uint 相加|相减
func Uint(a int, b int) int {
	return int(Uint32(uint32(a), int32(b)))
}

// Int32 相加|相减
func Int32(a int32, b int32) int32 {
	res := a + b
	if a > 0 && b > 0 && res < 0 {
		res = math.MaxInt32
	} else if a < 0 && b < 0 && res > 0 {
		res = math.MinInt32
	}
	return res
}

// Uint32 相加|相减
func Uint32(a uint32, b int32) uint32 {
	if b < 0 {
		// 相减
		b = b - b*2 // 变成正数
		if a <= uint32(b) {
			return 0
		}
		return a - uint32(b)
	}

	// 相加,防止溢出
	var newNum = a + uint32(b)
	if newNum < a {
		return math.MaxUint32
	}
	return newNum
}

// Int64 相加|相减
func Int64(a int64, b int64) int64 {
	res := a + b
	if a > 0 && b > 0 && res < 0 {
		res = math.MaxInt64
	} else if a < 0 && b < 0 && res > 0 {
		res = math.MinInt64
	}
	return res
}

// Uint64 相加|相减
func Uint64(a uint64, b int64) uint64 {
	if b < 0 {
		// 相减
		b = b - b*2 // 变成正数
		if a <= uint64(b) {
			return 0
		}
		return a - uint64(b)
	}

	// 相加,防止溢出
	var newNum = a + uint64(b)
	if newNum < a {
		return math.MaxUint64
	}
	return newNum
}
