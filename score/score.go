package score

import "math"

func dist(coord1 []float64, coord2 []float64) float64 {
	u := coord1[0] - coord2[0]
	v := coord1[1] - coord2[1]
	z := coord1[2] - coord2[2]
	return math.Sqrt(u*u + v*v + z*z)
}

func score(d, dbest, alpha, beta, penal, wa, wb, wpenal float64) float64 {
	// when d=0 we have the maximum penal
	penal = penal / 2.0
	value := 0.0
	if d < dbest {
		value = (wpenal*penal+wa*alpha)*(-math.Cos(math.Pi*d/dbest)) - (wpenal*penal - wa*alpha)
	} else {
		value = (2 * (wa * alpha)) * math.Exp((-1*math.Pow((d-dbest), 2))/(2*wb*beta))
	}
	return value
}
