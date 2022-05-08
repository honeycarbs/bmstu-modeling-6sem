package shoot

import (
	"fmt"
	"math"
	cyl "runge/cylinder"
)

func Shoot(c cyl.Cylinder) float64 {
	var (
		left  float64 = 1e-2
		right float64 = 1
	)

	f_left := Psi(left, c)
	f_right := Psi(right, c)

	// fmt.Println(f_left, f_right)
	// if f_left > f_right {
	// 	tmp := f_left
	// 	f_left = f_right
	// 	f_right = tmp
	// }
	fmt.Println(f_left, f_right)

	if math.Abs(f_left) < 1e-6 {
		return left
	}
	if math.Abs(f_right) < 1e-6 {
		return right
	}

	mid := (left + right) / 2
	for math.Abs((right-left)/mid) > c.Eps {
		// fmt.Println(left, right)

		f_mid := Psi(mid, c)
		if math.Abs(f_mid) < 1e-6 {
			return mid
		}
		if (f_left * f_mid) < 0 {
			right = mid
		} else {
			left = mid
		}
		mid = (left + right) / 2
		// fmt.Println(mid)
	}
	return mid
}
