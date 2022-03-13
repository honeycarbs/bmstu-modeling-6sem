package cauchy

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func PrintTable(x FloatVec64, pc FloatMtx64, eu, rk FloatVec64, h float64) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	table_skip := 0.05 / h
	num := 1

	t.AppendHeader(table.Row{"#", " X ", "picard-1st", "picard-2d", "picard-3d", "picard-4th", "euler", "runge-kutta"})
	for i := 0; i < len(pc[0]); i++ {
		if i%int(table_skip) == 0 {
			sx := fmt.Sprintf("%.5f", x[i])
			sp1 := fmt.Sprintf("%.8f", pc[0][i])
			sp2 := fmt.Sprintf("%.8f", pc[1][i])
			sp3 := fmt.Sprintf("%.8f", pc[2][i])
			sp4 := fmt.Sprintf("%.8f", pc[3][i])

			se := fmt.Sprintf("%.8f", eu[i])
			srk := fmt.Sprintf("%.8f", rk[i])

			t.AppendRow([]interface{}{num, sx, sp1, sp2, sp3, sp4, se, srk})
			num++
		}
	}
	t.Render()

	// fmt.Println(len(pc[0]))
}
