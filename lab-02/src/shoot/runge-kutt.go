package shoot

import cyl "runge/cylinder"

func RungeKuttGetNext(zn, yn, tn, h float64, c cyl.Cylinder) (float64, float64) {
	if zn+h > 1 {
		h = 1 - zn
	}
	k1 := h * c.RadiantFlux(zn, yn, tn)
	q1 := h * c.Phi(zn, yn, tn)
	k2 := h * c.RadiantFlux(zn+h/2, yn+k1/2, tn+q1/2)
	q2 := h * c.Phi(zn+h/2, yn+k1/2, tn+q1/2)
	k3 := h * c.RadiantFlux(zn+h/2, yn+k2/2, tn+q2/2)
	q3 := h * c.Phi(zn+h/2, yn+k2/2, tn+q2/2)
	k4 := h * c.RadiantFlux(zn+h, yn+k3, tn+q3)
	q4 := h * c.Phi(zn+h, yn+k3, tn+q3)

	y_next := yn + (k1+2*k2+2*k3+k4)/6
	t_next := tn + (q1+2*q2+2*q3+q4)/6

	return y_next, t_next
}

func RungeKutt(x_i float64, c cyl.Cylinder) (float64, float64) {
	var (
		z float64 = 0
		t float64 = 0
	)

	y := x_i * c.PlankFunction(0)
	h := 1e-2

	for z < 1 {
		y, t = RungeKuttGetNext(z, y, t, h, c)
		z += h
	}
	return y, t
}

func Psi(x_i float64, c cyl.Cylinder) float64 {
	var (
		u float64 = 0
		F float64 = 0
	)
	u, F = RungeKutt(x_i, c)
	return F - c.M*c.C*u/2
}
