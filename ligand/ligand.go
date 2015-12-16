package ligand

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"bitbucket.org/jgcarvalho/zdd/atom"
)

type Ligand struct {
	Name     string
	Atoms    []atom.Atom
	Angles   [3]float64
	NPatoms  int
	Patoms   int
	RotBonds int
}

func (l Ligand) Center() [3]float64 {
	return l.center()
}

func (l Ligand) center() [3]float64 {
	var center [3]float64
	max := [3]float64{math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64}
	min := [3]float64{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}

	for i := 0; i < len(l.Atoms); i++ {
		if l.Atoms[i].Coord[0] > max[0] {
			max[0] = l.Atoms[i].Coord[0]
		}
		if l.Atoms[i].Coord[1] > max[1] {
			max[1] = l.Atoms[i].Coord[1]
		}
		if l.Atoms[i].Coord[2] > max[2] {
			max[2] = l.Atoms[i].Coord[2]
		}
		if l.Atoms[i].Coord[0] < min[0] {
			min[0] = l.Atoms[i].Coord[0]
		}
		if l.Atoms[i].Coord[1] < min[1] {
			min[1] = l.Atoms[i].Coord[1]
		}
		if l.Atoms[i].Coord[2] < min[2] {
			min[2] = l.Atoms[i].Coord[2]
		}
	}
	center[0] = (max[0] - min[0]) / 2.0
	center[1] = (max[1] - min[1]) / 2.0
	center[2] = (max[2] - min[2]) / 2.0
	return center
}

func (l *Ligand) trans(x, y, z float64) {
	c := l.center()
	dx := x - c[0]
	dy := y - c[1]
	dz := z - c[2]
	for i := 0; i < len(l.Atoms); i++ {
		l.Atoms[i].Coord[0] = l.Atoms[i].Coord[0] + dx
		l.Atoms[i].Coord[1] = l.Atoms[i].Coord[1] + dy
		l.Atoms[i].Coord[2] = l.Atoms[i].Coord[2] + dz
	}
}

func dot(mrot [3][3]float64, coord [3]float64) (new_coord [3]float64) {
	new_coord[0] = mrot[0][0]*coord[0] + mrot[0][1]*coord[1] + mrot[0][2]*coord[2]
	new_coord[1] = mrot[1][0]*coord[0] + mrot[1][1]*coord[1] + mrot[1][2]*coord[2]
	new_coord[2] = mrot[2][0]*coord[0] + mrot[2][1]*coord[1] + mrot[2][2]*coord[2]
	return new_coord
}

func rotX(coord [3]float64, a float64) [3]float64 {
	sin_a, cos_a := math.Sincos(a)
	m_rot := [3][3]float64{[3]float64{1.0, 0.0, 0.0}, [3]float64{0.0, cos_a, sin_a}, [3]float64{0.0, -sin_a, cos_a}}
	return dot(m_rot, coord)
}

func rotY(coord [3]float64, b float64) [3]float64 {
	sin_b, cos_b := math.Sincos(b)
	m_rot := [3][3]float64{[3]float64{cos_b, 0.0, -sin_b}, [3]float64{0.0, 1.0, 0.0}, [3]float64{sin_b, 0.0, cos_b}}
	return dot(m_rot, coord)
}

func rotZ(coord [3]float64, c float64) [3]float64 {
	sin_c, cos_c := math.Sincos(c)
	m_rot := [3][3]float64{[3]float64{cos_c, sin_c, 0.0}, [3]float64{-sin_c, cos_c, 0.0}, [3]float64{0.0, 0.0, 1.0}}
	return dot(m_rot, coord)
}

func (l *Ligand) spin(a, b, c float64) {
	current := l.center()
	l.trans(0.0, 0.0, 0.0)
	for i := 0; i < len(l.Atoms); i++ {
		l.Atoms[i].Coord = rotX(l.Atoms[i].Coord, a)
		l.Atoms[i].Coord = rotY(l.Atoms[i].Coord, b)
		l.Atoms[i].Coord = rotZ(l.Atoms[i].Coord, c)
	}
	l.Angles[0] = a
	l.Angles[1] = b
	l.Angles[2] = c
	l.trans(current[0], current[1], current[2])
}

func (l *Ligand) Move(x, y, z, a, b, c float64) {
	l.trans(x, y, z)
	// l.spin(a, b, c)
}

func LoadMol2(fn string) Ligand {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	atoms := make([]atom.Atom, 0)
	npAtoms := 0
	pAtoms := 0
	bf := bufio.NewReader(f)
	isAtom := false
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
		if strings.Contains(string(line), "@<TRIPOS>BOND") {
			isAtom = false
		}
		if isAtom {
			x, _ := strconv.ParseFloat(strings.TrimSpace(string(line)[16:26]), 64)
			y, _ := strconv.ParseFloat(strings.TrimSpace(string(line)[26:36]), 64)
			z, _ := strconv.ParseFloat(strings.TrimSpace(string(line)[36:46]), 64)
			//TODO extrair o tipo do atomo, problema que o numero do residuo gruda no nome
			name := ""
			for _, v := range atom.AtomNames {
				if strings.Contains(string(line)[46:56], v) {
					name = v
					atoms = append(atoms, atom.Atom{name, [3]float64{x, y, z}})
					_, isNp := atom.NonPolar[name]
					if isNp {
						npAtoms += 1
					} else {
						pAtoms += 1
					}
					break
				}
			}
		}
		if strings.Contains(string(line), "@<TRIPOS>ATOM") {
			isAtom = true
		}
	}

	return Ligand{fn, atoms, [3]float64{0.0, 0.0, 0.0}, npAtoms, pAtoms, 0}
	// fmt.Println(atoms)
}
