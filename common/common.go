package common

import (
	"math/rand"
	"strconv"
	"time"
)

// WeightRand 从map里根据比率随机一个键值出来
func WeightRand(data map[int]int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	sum := 0
	for _, v := range data {
		sum += v
	}
	randNum := r.Intn(sum)

	var k int
	var v int
	for k, v = range data {
		if v >= randNum {
			break
		}
		randNum -= v
	}
	data = nil
	return k
}

// Now32 获取当前时间戳
func Now32() int {
	return int(time.Now().Unix())
}

// FormatStringToInt 格式化string为int64
func FormatStringToInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// FormatStringToFloat 格式化string为float64
func FormatStringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
