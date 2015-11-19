package score

import (
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"bufio"
)

type Interaction struct {
	A      string
	B      string
	Dbest  float64
	Alpha  float64
	Beta   float64
	Penal  float64
	Wa     float64
	Wb     float64
	Wpenal float64
}

type Parameters []Interaction

func dist(coord1 [3]float64, coord2 [3]float64) float64 {
	u := coord1[0] - coord2[0]
	v := coord1[1] - coord2[1]
	z := coord1[2] - coord2[2]
	return math.Sqrt(u*u + v*v + z*z)
}

func score(d, dbest, alpha, beta, penal, wa, wb, wpenal float64) float64 {
	// when d=0 we have the maximum penal
	penal = penal / 2.0
	alpha = alpha / 2.0
	value := 0.0
	if d < dbest {
		value = (wpenal*penal+wa*alpha)*(-math.Cos(math.Pi*d/dbest)) - (wpenal*penal - wa*alpha)
	} else {
		value = (2 * (wa * alpha)) * math.Exp((-1*math.Pow((d-dbest), 2))/(2*wb*beta))
	}
	return value
}

func LoadParams(fn string) Parameters {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	params := make([]Interaction, 861)
	count := 0
	bf := bufio.NewReader(f)
	bf.ReadLine() //skip header
	for {
		line, isPrefix, err := bf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", f.Name())
		}
		data := strings.Split(string(line), "\t")
		d, _ := strconv.ParseFloat(data[2], 64)
		a, _ := strconv.ParseFloat(data[3], 64)
		b, _ := strconv.ParseFloat(data[4], 64)
		penal, _ := strconv.ParseFloat(data[5], 64)
		wa, _ := strconv.ParseFloat(data[6], 64)
		wb, _ := strconv.ParseFloat(data[7], 64)
		wpenal, _ := strconv.ParseFloat(data[8], 64)
		params[count] = Interaction{data[0], data[1], d, a, b, penal, wa, wb, wpenal}
		count++
	}
	return params
}

// func Score(p *protein.Protein, l *ligand.Ligand) {
//
// }
