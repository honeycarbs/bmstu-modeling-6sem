package cylinder

import "math"

func (c *Cylinder) T(z float64) float64 {
	return (c.T_w-c.T_0)*math.Pow(z, c.Order) + c.T_0
}

func (c *Cylinder) Takeover(z float64) float64 { // k
	t_z := c.T(z)
	return c.K_0 * t_z * t_z / 90000
}

func (c *Cylinder) PlankFunction(z float64) float64 { // u_p
	return PLANK_NUM / (math.Exp(PLANK_DEN/c.T(z)) - 1)
}

func (c *Cylinder) Kappa(z float64) float64 {
	h := 1e-6
	return c.C * (c.Takeover(z+h/2) + c.Takeover(z-h/2)) / 6 / c.R / c.Takeover(z+h/2) / c.Takeover(z-h/2)
}
