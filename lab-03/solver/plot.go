package solver

import (
	"fmt"
	"os"

	"github.com/Arafatk/glot"
	"github.com/jedib0t/go-pretty/table"
)

func TableShow(x, y, f, f1, u_s, z_s []float64) {
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)

	// s_i := 0
	// u_s_out := " "
	writer.AppendHeader(table.Row{"z", "u(z)", " F(z)", "F_1(z)", "u_p(z)"})

	for i := 0; i < len(x); i += 10000 {
		x_out := fmt.Sprintf("%15.3f", x[i])
		y_out := fmt.Sprintf("%15.3e", y[i])
		f_out := fmt.Sprintf("%15.3e", f[i])
		f1_out := fmt.Sprintf("%15.3e", f1[i])
		u_s_out := fmt.Sprintf("%15.3e", u_s[i])
		// if z_s[s_i] == 0 || int(x[i]) == 1 {

		// 	s_i += 1
		// } else {
		// 	u_s_out = " "
		// }
		writer.AppendRow([]interface{}{x_out, y_out, f_out, f1_out, u_s_out})
	}

	writer.Render()
}

func PlotShow(x, y []float64, title, fname, labelX, labelY string) {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	style := "lines"
	points := [][]float64{x, y}
	pointGroupName := title
	plot.AddPointGroup(pointGroupName, style, points)
	plot.SetXLabel(labelX)
	plot.SetYLabel(labelY)

	plot.SavePlot("assets/" + fname + ".png")
}
