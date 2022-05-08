package shoot

import (
	"fmt"
	"image/color"
	"os"
	cyl "runge/cylinder"

	"github.com/jedib0t/go-pretty/table"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Show(x_i float64, c cyl.Cylinder) {
	var (
		z        float64 = 0
		t        float64 = 0
		h        float64 = 1e-4
		u        float64 = c.PlankFunction(z)
		t_writer         = table.NewWriter()
	)
	z_s := []float64{0}
	u_p_s := []float64{u}

	for z < 1 {
		z += h
		z_s = append(z_s, z)
		u_p_s = append(u_p_s, c.PlankFunction(z))
	}

	t_s := []float64{0}
	y_s := []float64{x_i * c.PlankFunction(0)}
	z = 0
	z_s = []float64{0}
	y := y_s[0]
	h = 1e-2

	for z < 1 {
		y, t = RungeKuttGetNext(z, y, t, h, c)
		z += h
		z_s = append(z_s, z)
		y_s = append(y_s, y)
		t_s = append(t_s, t)
	}

	fmt.Println(len(z_s), len(u_p_s))

	t_writer.SetOutputMirror(os.Stdout)
	t_writer.AppendHeader(table.Row{"z", "F(z)", " u(z) * 1e-7", "u_p(z) * 1e-6"})
	for i := 0; i < len(z_s); i++ {
		if i%10 == 0 {
			z_s_out := fmt.Sprintf("%4.2f", z_s[i])
			t_s_out := fmt.Sprintf("%8.2f", t_s[i])
			y_s_out := fmt.Sprintf("%6.2f", y_s[i]/1e-7)
			u_p_s_out := fmt.Sprintf("%6.2f", u_p_s[i*100]/1e-6)
			t_writer.AppendRow([]interface{}{z_s_out, t_s_out, y_s_out, u_p_s_out})
		}
	}

	t_writer.Render()

	fig_1 := plot.New()

	fig_1.Title.Text = "Зависимость F(z) от безразмерной координаты z"
	fig_1.X.Label.Text = "z"
	fig_1.Y.Label.Text = "F(z)"
	fig_1.Add(plotter.NewGrid())

	dots1 := makeDots(z_s, t_s)

	l1, err := plotter.NewLine(dots1)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	l1.LineStyle.Width = vg.Points(1)
	l1.LineStyle.Color = color.RGBA{B: 214, A: 91}

	fig_1.Add(l1)

	fig_1.Y.Max = 5000
	fig_1.X.Max = 1.1

	if err := fig_1.Save(10*vg.Inch, 4*vg.Inch, "assets/plot1.png"); err != nil {
		panic(err)
	}

	fig_2 := plot.New()
	fig_2.Title.Text = "Зависимость u(z) от безразмерной координаты z"
	fig_2.X.Label.Text = "z"
	fig_2.Y.Label.Text = "u(z)"
	fig_2.Add(plotter.NewGrid())

	dots2 := makeDots(z_s, y_s)

	l2, err := plotter.NewLine(dots2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	l2.LineStyle.Width = vg.Points(1)
	l2.LineStyle.Color = color.RGBA{B: 255, A: 255}

	fig_2.Add(l2)

	z = 0
	u = c.PlankFunction(z)
	h = 1e-4
	z_s = []float64{0}
	u_p_s = []float64{u}

	for z < 1 {
		z += h
		z_s = append(z_s, z)
		u_p_s = append(u_p_s, c.PlankFunction(z))
	}

	dots3 := makeDots(z_s, u_p_s)
	l3, err := plotter.NewLine(dots3)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	l3.LineStyle.Width = vg.Points(1)
	l3.LineStyle.Color = color.RGBA{B: 255, A: 90}

	fig_2.Add(l3)

	fig_2.Y.Max = 2.8e-6
	fig_2.X.Max = 1.1

	if err := fig_2.Save(10*vg.Inch, 4*vg.Inch, "assets/plot2.png"); err != nil {
		panic(err)
	}
}

func makeDots(xx, yy []float64) plotter.XYs {
	var conv plotter.XYs

	for i := 0; i < len(xx); i++ {
		d := plotter.XY{
			X: xx[i],
			Y: yy[i],
		}
		conv = append(conv, d)
	}

	return conv
}
