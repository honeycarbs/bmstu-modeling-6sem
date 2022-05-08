package solver

import c "gaussian/cylinder"

func Solve(cyl c.Cylinder) ([]float64, []float64) {
	var (
		h      = 1e-6
		z_half = h / 2
	)
	M0 := cyl.Kappa(z_half)*z_half + cyl.C*cyl.R*h*h/8*cyl.Takeover(z_half)*z_half
	K0 := -cyl.Kappa(z_half)*z_half + cyl.C*cyl.R*h*h/8*cyl.Takeover(z_half)*z_half
	P0 := cyl.C * cyl.R * h * h / 4 * cyl.Takeover(z_half) * cyl.PlankFunction(z_half) * z_half
	xi := []float64{0}
	etha := []float64{0}

	xi = append(xi, -K0/M0)
	etha = append(etha, P0/M0)

	for i := 1; i < int(1/h)+1; i++ {
		a_n := cyl.Kappa(z_half) * z_half
		z_half += h / 2
		b_n := a_n + cyl.C*cyl.Takeover(z_half)*z_half*h*h*cyl.R
		f_n := cyl.C * cyl.Takeover(z_half) * cyl.PlankFunction(z_half) * h * h * z_half * cyl.R
		z_half += h / 2
		c_n := cyl.Kappa(z_half) * z_half
		b_n += c_n
		xi = append(xi, c_n/(b_n-a_n*xi[i]))
		etha = append(etha, (f_n+a_n*etha[i])/(b_n-a_n*xi[i]))
	}
	y := make([]float64, len(xi))
	z_half = 1 - h/2
	MN := -cyl.Kappa(z_half)*z_half + cyl.R*cyl.C*h*h/8*cyl.Takeover(z_half)*z_half
	KN := cyl.Kappa(z_half)*z_half + cyl.M*cyl.C*h/2 + cyl.R*cyl.C*h*h/8*cyl.Takeover(z_half)*z_half + cyl.R*cyl.C*h*h/4*cyl.Takeover(1)
	PN := cyl.C * cyl.R * h * h / 4 * (cyl.Takeover(z_half)*cyl.PlankFunction(z_half)*z_half + cyl.Takeover(1)*cyl.PlankFunction(1))

	y[len(y)-1] = (PN - MN*etha[len(etha)-1]) / (MN*xi[len(xi)-1] + KN)
	for i := len(y) - 2; i >= 0; i-- {
		y[i] = y[i+1]*xi[i+1] + etha[i+1]
	}
	z := make([]float64, len(y))
	for i := 0; i < len(z); i++ {
		z[i] = 0 + float64(i)*h
	}

	return y, z
}

func GetFDerivatives(y, z []float64, cyl c.Cylinder) []float64 {
	h := 1e-6

	dy := make([]float64, 0)
	dy = append(dy, (-3*y[0]+4*y[1]-y[2])/2/h)

	for i := 1; i < len(y)-1; i++ {
		dy = append(dy, (y[i+1]-y[i-1])/2/h)
	}
	dy = append(dy, (3*y[len(y)-1]-4*y[len(y)-2]+y[len(y)-3])/2/h)

	f := make([]float64, len(z))
	for i := 0; i < len(f); i++ {
		f[i] = -cyl.C / 3 / cyl.R / cyl.Takeover(z[i]) * dy[i]
	}
	f[0] = 0
	return f
}

func GetFIntegrals(u, z []float64, cyl c.Cylinder) []float64 {
	f := make([]float64, len(z))
	h := 1e-6
	s := 0.
	z_curr := 0.
	z_next := h
	for i := 1; i < len(f); i++ {
		s += h / 2 * (cyl.Takeover(z_curr)*(cyl.PlankFunction(z_curr)-u[i-1])*z_curr + cyl.Takeover(z_next)*(cyl.PlankFunction(z_next)-u[i])*z_next)
		f[i] = cyl.C * cyl.R / z_next * s
		z_curr = z_next
		z_next += h
	}

	return f
}

func U_p_show(cyl c.Cylinder) ([]float64, []float64) {
	z := 0.
	h := 1e-6
	z_s := []float64{0}
	u_s := []float64{cyl.PlankFunction(z)}
	for z < 1 {
		z += h
		z_s = append(z_s, z)
		u_s = append(u_s, cyl.PlankFunction(z))
	}

	return u_s, z_s
}
