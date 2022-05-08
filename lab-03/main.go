package main

import (
	c "gaussian/cylinder"
	s "gaussian/solver"
)

func main() {
	cyli := c.Cylinder{
		M:     0.786,
		C:     3e10,
		R:     0.35,
		T_w:   2e3,
		T_0:   1e4,
		K_0:   9,
		Order: 4,
		Eps:   1e-4}

	u, z := s.Solve(cyli)
	// s.PlotShow(z, u, "Распределение плотности излучения U от z", "plot_uz", "z", "U")
	f := s.GetFDerivatives(u, z, cyli)
	// s.PlotShow(z, f, "Распределение плотности излучения F от z", "plot_fz", "z", "F")
	f1 := s.GetFIntegrals(u, z, cyli)
	// s.PlotShow(z, f1, "Распределение плотности излучения F1 от z", "plot_f1z", "z", "F1")
	u_s, z_s := s.U_p_show(cyli)
	// s.PlotShow(z_s, u_s, "Распределение равновесной плотности излучения U_p от z", "plot_upz", "z", "U_p")

	s.TableShow(z, u, f, f1, u_s, z_s)
}
