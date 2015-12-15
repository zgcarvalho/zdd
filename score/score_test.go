package score

import (
	"math"
	"testing"
)

func TestDist(t *testing.T) {
	d := dist([3]float64{0.0, 0.0, 0.0}, [3]float64{0.0, 0.0, 0.0})
	dexp := 0.0
	errmax := 0.0000000000000001
	if math.Abs(d-dexp) > errmax {
		t.Fail()
	}
	d = dist([3]float64{0.0, 0.0, 0.0}, [3]float64{1.0, 0.0, 0.0})
	dexp = 1.0
	if math.Abs(d-dexp) > errmax {
		t.Fail()
	}
	d = dist([3]float64{0.0, 0.0, 0.0}, [3]float64{1.0, 2.0, 3.0})
	dexp = math.Sqrt(14)
	if math.Abs(d-dexp) > errmax {
		t.Fail()
	}
}

func TestScore(t *testing.T) {
	sc := score(0, 3, -1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp := 30000.0
	errmax := 0.001
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(100000, 3, -1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp = 0.0
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(3.0000000000000000000001, 3, -1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp = -1500.0
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(2.9999999999999999999999, 3, -1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp = -1500.0
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(3, 3, -1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp = -1500.0
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(0, 3, -1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp = 30000.0
	errmax = 0.001
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(100000, 3, 1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp = 0.0
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(3.0000000000000000000001, 3, 1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp = 1500.0
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(2.9999999999999999999999, 3, 1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp = 1500.0
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
	sc = score(3, 3, 1500.0, 0.5, 30000.0, 1.0, 1.0, 1.0)
	scexp = 1500.0
	if math.Abs(sc-scexp) > errmax {
		t.Fail()
	}
}
