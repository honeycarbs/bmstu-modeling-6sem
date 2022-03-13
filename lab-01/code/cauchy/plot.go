package cauchy

import (
	"math"

	"github.com/Arafatk/glot"
)

func getMaxValue(x, eu, rkut FloatVec64) int {
	max := 0

	for max = 0; max < len(eu) && math.Abs(eu[max]-rkut[max]) < 1e-2; max++ {
	}
	return max
}

func getFullLine(pos, neg FloatVec64) FloatVec64 {
	dst := make(FloatVec64, len(neg)+len(pos))

	j := 0
	for i := len(neg) - 1; i >= 0; i-- {
		dst[j] = neg[i]
		j++
	}
	for i := 0; i < len(pos); i++ {
		dst[j] = pos[i]
		j++
	}
	return dst
}

func ExportPlot(x_pos, x_neg FloatVec64, pc_pos, pc_neg FloatMtx64, eu_pos, eu_neg, rk_pos, rk_neg FloatVec64) {

	dimensions := 2
	persist := false
	debug := false

	plot, _ := glot.NewPlot(dimensions, persist, debug)
	plot.SetTitle("Cauchy problem solving")
	plot.SetXLabel("x")
	plot.SetYLabel("u(x)")

	plot.SetXrange(-2, 4)
	plot.SetYrange(-8, 8)

	style := "lines"

	max := getMaxValue(x_pos, eu_pos, rk_pos)

	x := getFullLine(x_pos[:max], x_neg[:max])
	eu := getFullLine(eu_pos[:max], eu_neg[:max])
	rk := getFullLine(rk_pos[:max], rk_neg[:max])
	pc1 := getFullLine(pc_pos[0][:max], pc_neg[0][:max])
	pc2 := getFullLine(pc_pos[1][:max], pc_neg[1][:max])
	pc3 := getFullLine(pc_pos[2][:max], pc_neg[2][:max])
	pc4 := getFullLine(pc_pos[3][:max], pc_neg[3][:max])

	eu_legend := "Euler"
	points_eu := [][]float64{x, eu}

	rk_legend := "Runge-Kutta"
	points_rk := [][]float64{x, rk}

	pc1_legend := "Picard 1-st"
	points_pc1 := [][]float64{x, pc1}

	pc2_legend := "Picard 2-nd"
	points_pc2 := [][]float64{x, pc2}

	pc3_legend := "Picard 3-rd"
	points_pc3 := [][]float64{x, pc3}

	pc4_legend := "Picard 4-th"
	points_pc4 := [][]float64{x, pc4}

	plot.AddPointGroup(eu_legend, style, points_eu)
	plot.AddPointGroup(rk_legend, style, points_rk)
	plot.AddPointGroup(pc1_legend, style, points_pc1)
	plot.AddPointGroup(pc2_legend, style, points_pc2)
	plot.AddPointGroup(pc3_legend, style, points_pc3)
	plot.AddPointGroup(pc4_legend, style, points_pc4)

	err := plot.SavePlot("plot.png")
	if err != nil {
		panic(err)
	}

}
