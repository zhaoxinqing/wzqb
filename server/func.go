package server

import (
	"fmt"
	"math"
)

const Pi = 3.14

// 在 Go 语言，函数也是一种值( C++ 函数对象、 Python 函数类似)，可以被传递。
// 跟其他普通值一样，函数也可以作为 参数 传递或作为 返回值 返回。

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func number() {
	var a uint8
	fmt.Println(a)
}
