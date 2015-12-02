package main

import (
	// "flag"
	// "fmt"

	// "bitbucket.org/jgcarvalho/zdd/ligand"
	// "bitbucket.org/jgcarvalho/zdd/protein"
	"bitbucket.org/jgcarvalho/zdd/score"
	"bitbucket.org/jgcarvalho/zdd/train"
)

func main() {
	// f_prot := flag.String("p", "", "protein mol2")
	// f_lig := flag.String("l", "", "ligand mol2")
	// flag.Parse()
	params := score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
	// // fmt.Println(params)
	// // protein := protein.LoadMol2("/home/jgcarvalho/pdbbind/core/proteinmol2/10gs_protein.mol2")
	// protein := protein.LoadMol2(*f_prot)
	// // fmt.Println(protein)
	// // ligand := ligand.LoadMol2("/home/jgcarvalho/pdbbind/core/ligands/10gs_ligand.mol2")
	// ligand := ligand.LoadMol2(*f_lig)
	// // fmt.Println(ligand)
	// total := params.Score(&protein, &ligand)
	// fmt.Println(total)
	train.Train(params)
}
