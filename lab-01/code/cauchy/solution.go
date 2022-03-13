package cauchy

import (
	"math"
)

func first_approx(x float64) float64 {
	return x * x * x / 3
}

func sec_approx(x float64) float64 {
	return first_approx(x) + math.Pow(x, 7)/63
}

func third_approx(x float64) float64 {
	return sec_approx(x) + 2*math.Pow(x, 11)/(3*7*9*11) + math.Pow(x, 15)/(9*9*7*7*15)
}

func fourth_approx(x float64) float64 {
	f := 4 * math.Pow(x, 15) / (3 * 3 * 7 * 9 * 11 * 15)
	s1 := 4 * math.Pow(x, 19) / (3 * 7 * 7 * 9 * 9 * 11 * 19)
	s2 := 2 * math.Pow(x, 19) / (3 * 9 * 9 * 7 * 7 * 15 * 19)
	t1 := 2 * math.Pow(x, 23) / (9 * 9 * 9 * 7 * 7 * 7 * 15 * 23)
	t2 := 2 * math.Pow(x, 23) / (3 * 3 * 7 * 7 * 9 * 9 * 11 * 11 * 23)
	fr := 4 * math.Pow(x, 27) / (3 * 7 * 7 * 7 * 9 * 9 * 9 * 11 * 15 * 27)
	fv := math.Pow(x, 31) / (9 * 9 * 9 * 9 * 7 * 7 * 7 * 7 * 15 * 15 * 31)

	return third_approx(x) + f + s1 + s2 + t1 + t2 + fr + fv
}

func PicardSolver(x0, h float64, n int) FloatMtx64 {
	values := MakeFloatMtx64(4, 0)

	// table_skip := 0.5 / h
	// fmt.Println(table_skip)

	for i := 0; i <= n; i++ {
		values[0] = append(values[0], first_approx(x0))
		values[1] = append(values[1], sec_approx(x0))
		values[2] = append(values[2], third_approx(x0))
		values[3] = append(values[3], fourth_approx(x0))

		x0 += h
		// x_prev = x0
	}

	return values
}

func EulerSolver(x0, y0, h float64, n int) FloatVec64 {
	values := make(FloatVec64, 0)

	for i := 0; i <= n; i++ {
		values = append(values, y0)
		y0 += h * domain(x0, y0)
		x0 += h
	}

	return values
}

func RungeKuttaSolver(x0, y0, alpha, h float64, n int) FloatVec64 {
	values := make(FloatVec64, 0)

	for i := 0; i <= n; i++ {
		values = append(values, y0)
		y0 += h * ((1-alpha)*domain(x0, y0) + alpha*domain(x0+h/2/alpha, y0+h*domain(x0, y0)/2/alpha))
		x0 += h
	}

	return values
}
