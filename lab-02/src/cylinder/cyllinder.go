package cylinder

import "math"

func (c *Cylinder) T(z float64) float64 {
	return (c.T_w-c.T_0)*math.Pow(z, c.Order) + c.T_0
}

func (c *Cylinder) Takeover(z float64) float64 { // k
	return c.K_0 * (math.Pow(c.T(z)/300, 2))
}

func (c *Cylinder) PlankFunction(z float64) float64 { // u_p
	return PLANK_NUM / (math.Exp(PLANK_DEN/c.T(z)) - 1)
}

func (c *Cylinder) RadiantFlux(z, u, F float64) float64 { // f
	return -3.0 * c.R * c.Takeover(z) * F / c.C
}

func (c *Cylinder) Phi(z, u, F float64) float64 {
	var rv float64

	if math.Abs(z) < 1e-10 {
		rv = c.C * c.Takeover(z) * (c.PlankFunction(z) - u) * c.R / 2
	} else {
		rv = c.C*c.Takeover(z)*(c.PlankFunction(z)-u)*c.R - F/z
	}

	return rv
}
