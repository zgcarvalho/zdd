package score

import (
	// "fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"bufio"

	"bitbucket.org/jgcarvalho/zdd/ligand"
	"bitbucket.org/jgcarvalho/zdd/protein"
)

type Interaction struct {
	A      string
	B      string
	Type   string
	Dbest  float64
	Alpha  float64
	Beta   float64
	Penal  float64
	Wa     float64
	Wb     float64
	Wpenal float64
}

type Parameters struct {
	Inter map[string]Interaction
	Wnp   float64
	Enp   float64
	Wp    float64
	Ep    float64
	Wrb   float64
	Erb   float64
	C     float64
}

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
		value = (wpenal*penal+(wa*alpha*-1.0))*(math.Cos(math.Pi*d/dbest)) + (wpenal*penal + wa*alpha)
	} else {
		value = (2 * (wa * alpha)) * math.Exp((-1*math.Pow((d-dbest), 2))/(2*wb*beta))
	}
	return value
}

func LoadParams(fn string) Parameters {
	var params Parameters
	f, err := os.Open(fn)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	//Parameters
	params.Inter = make(map[string]Interaction)
	// count := 0
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
		d, _ := strconv.ParseFloat(data[3], 64)
		a, _ := strconv.ParseFloat(data[4], 64)
		b, _ := strconv.ParseFloat(data[5], 64)
		penal, _ := strconv.ParseFloat(data[6], 64)
		wa, _ := strconv.ParseFloat(data[7], 64)
		wb, _ := strconv.ParseFloat(data[8], 64)
		wpenal, _ := strconv.ParseFloat(data[9], 64)
		params.Inter[data[0]+"_"+data[1]] = Interaction{data[0], data[1], data[2], d, a, b, penal, wa, wb, wpenal}
		params.Inter[data[1]+"_"+data[0]] = Interaction{data[0], data[1], data[2], d, a, b, penal, wa, wb, wpenal}
		// params[count] = Interaction{data[0], data[1], d, a, b, penal, wa, wb, wpenal}
		// count++
	}
	return params
}

func (prm Parameters) Score(p *protein.Protein, l *ligand.Ligand) float64 {
	total := 0.0
	for _, la := range l.Atoms {
		for _, lp := range p.Atoms {
			c := lp.Name + "_" + la.Name
			total += score(dist(lp.Coord, la.Coord), prm.Inter[c].Dbest, prm.Inter[c].Alpha, prm.Inter[c].Beta, prm.Inter[c].Penal, prm.Inter[c].Wa, prm.Inter[c].Wb, prm.Inter[c].Wpenal)
		}
	}
	total += prm.Wnp * math.Pow(float64(l.NPatoms), prm.Enp)
	total += prm.Wp * math.Pow(float64(l.Patoms), prm.Ep)
	total += prm.Wrb * math.Pow(float64(l.RotBonds), prm.Erb)
	total += prm.C
	return total
}
