package tools

import (
	"math"
	"math/rand"
)

//正态分布公式
func NormalFloat64(x int64, miu int64, sigma int64) float64 {
	randomNormal := 1 / (math.Sqrt(2*math.Pi) * float64(sigma)) * math.Pow(math.E, -math.Pow(float64(x-miu), 2)/(2*math.Pow(float64(sigma), 2)))
	return randomNormal
}

//正态分布随机数
func RandomNormalInt64(min int, max int, miu int, sigma int64) int {
	if min >= max {
		return 0
	}
	if miu < min {
		miu = min
	}
	if miu > max {
		miu = max
	}
	var x int
	var y, dScope float64
	for {
		x = rand.Intn(max)
		y = NormalFloat64(int64(x), int64(miu), sigma) * 100000
		dScope = float64(rand.Intn(int(NormalFloat64(int64(miu), int64(miu), sigma) * 100000)))
		//注意下传的是两个miu
		if dScope <= y {
			break
		}
	}
	return int(x)
}
