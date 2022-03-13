package cauchy

type FloatVec64 []float64

type FloatMtx64 []FloatVec64

func MakeFloatMtx64(l, h int) FloatMtx64 {
	m := make(FloatMtx64, l)
	for i := 0; i < h; i++ {
		m[i] = make(FloatVec64, h)
	}
	return m
}

func domain(x, u float64) float64 {
	return x*x + u*u
}
