package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.",
		math.Nextafter(2, 3))
}

// %gは%e(科学的記数法、例: -1234.456e+78)、
//     %f(指数なしの小数、例: 123.456)のどちらか出力の短い方
