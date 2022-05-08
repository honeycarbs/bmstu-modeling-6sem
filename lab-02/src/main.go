package main

import (
	"fmt"
	c "runge/cylinder"
	s "runge/shoot"
)

func main() {
	cyli := c.Cylinder{
		M:     0.786,
		C:     3e10,
		R:     0.35,
		T_w:   2e3,
		T_0:   1e4,
		K_0:   8e-2,
		Order: 4,
		Eps:   1e-15}

	x_i := s.Shoot(cyli)
	fmt.Println(x_i)
	s.Show(x_i, cyli)
}
