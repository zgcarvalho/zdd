package score

import (
	"math"
	"testing"
)

func TestDist(t *testing.T) {
	d := dist([]float64{0.0, 0.0, 0.0}, []float64{0.0, 0.0, 0.0})
	dexp := 0.0
	errmax := 0.0000000000000001
	if math.Abs(d-dexp) > errmax {
		t.Fail()
	}
	d = dist([]float64{0.0, 0.0, 0.0}, []float64{1.0, 0.0, 0.0})
	dexp = 1.0
	if math.Abs(d-dexp) > errmax {
		t.Fail()
	}
	d = dist([]float64{0.0, 0.0, 0.0}, []float64{1.0, 2.0, 3.0})
	dexp = math.Sqrt(14)
	if math.Abs(d-dexp) > errmax {
		t.Fail()
	}
}

func TestScore(t *testing.T) {
	sc := score(0, 3, 0.3, 0.5, 10, 0.6, 1, 0.8)
	scexp := -8.0
	errmax := 0.0000001
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(100000, 3, 0.3, 0.5, 10, 0.6, 1, 0.8)
	scexp = 0.0
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(3.0000000000000000000001, 3, 0.3, 0.5, 10, 0.6, 1, 0.8)
	scexp = 0.36
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(2.9999999999999999999999, 3, 0.3, 0.5, 10, 0.6, 1, 0.8)
	scexp = 0.36
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(3, 3, 0.3, 0.5, 10, 0.6, 1, 0.8)
	scexp = 0.36
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
}
