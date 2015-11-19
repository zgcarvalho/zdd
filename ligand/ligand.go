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

func LoadMol2(fn string) {
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
		if strings.Contains(string(line), "@<TRIPOS>ATOM") {
			isAtom = true
		}
		if isAtom {
			x, _ := strconv.ParseFloat(string(line)[16:26], 64)
			y, _ := strconv.ParseFloat(string(line)[26:36], 64)
			z, _ := strconv.ParseFloat(string(line)[36:46], 64)
			//TODO extrair o tipo do atomo, problema que o numero do residuo gruda no nome
			name := string(line)[47:]
			atoms = append(atoms, atom.Atom{name, [3]float64{x, y, z}})
		}
	}
}
