package main

import (
	"math"

	"./cauchy"
)

func main() {
	x_low := 0.
	x_high := 2.

	y_low := 0.

	h := 1e-4
	alpha := 0.5
	n := int(math.Ceil(math.Abs(x_high-x_low) / h))

	x_pos := make(cauchy.FloatVec64, 0)
	x_n := x_low
	for i := 0; i <= n; i++ {
		x_pos = append(x_pos, x_n)
		x_n += h
	}

	x_neg := make(cauchy.FloatVec64, 0)
	for i := 0; i <= n; i++ {
		x_neg = append(x_neg, -x_pos[i])
	}

	pc_pos := cauchy.PicardSolver(x_low, h, n)
	eu_pos := cauchy.EulerSolver(x_low, y_low, h, n)
	rk_pos := cauchy.RungeKuttaSolver(x_low, y_low, alpha, h, n)

	pc_neg := cauchy.PicardSolver(x_low, -h, n)
	eu_neg := cauchy.EulerSolver(x_low, y_low, -h, n)
	rk_neg := cauchy.RungeKuttaSolver(x_low, y_low, alpha, -h, n)

	cauchy.PrintTable(x_pos, pc_pos, eu_pos, rk_pos, h)
	cauchy.ExportPlot(x_pos, x_neg, pc_pos, pc_neg, eu_pos, eu_neg, rk_pos, rk_neg)
}
