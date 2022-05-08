package cylinder

const PLANK_NUM = 3.084e-4
const PLANK_DEN = 4.799e4

type Cylinder struct {
	M     float64
	C     float64
	R     float64
	T_w   float64
	T_0   float64
	K_0   float64
	Order float64
	Eps   float64
}
