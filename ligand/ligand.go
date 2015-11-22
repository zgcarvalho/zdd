package ligand

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"bitbucket.org/jgcarvalho/zdd/atom"
)

type Ligand struct {
	Name  string
	Atoms []atom.Atom
}

func LoadMol2(fn string) Ligand{
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	atoms := make([]atom.Atom, 0)
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
			for _, v := range(atom.AtomNames) {
				if strings.Contains(string(line)[46:56], v) {
					name = v
					atoms = append(atoms, atom.Atom{name, [3]float64{x, y, z}})
					break
				}
			}
		}
		if strings.Contains(string(line), "@<TRIPOS>ATOM") {
			isAtom = true
		}
	}
	return Ligand{fn, atoms}
	// fmt.Println(atoms)
}
