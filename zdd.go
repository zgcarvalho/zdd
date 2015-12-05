package main

import (
	"flag"
	"fmt"

	"bitbucket.org/jgcarvalho/zdd/ligand"
	"bitbucket.org/jgcarvalho/zdd/protein"
	"bitbucket.org/jgcarvalho/zdd/score"
	"bitbucket.org/jgcarvalho/zdd/train"
)

func main() {
	doTrain := flag.Bool("train", false, "train using genetic algorithm")
	doTrain2 := flag.Bool("train2", false, "train using nelder mead")
	doScore := flag.Bool("score", false, "score conformation")
	fprot := flag.String("p", "", "protein mol2")
	flig := flag.String("l", "", "ligand mol2")
	flag.Parse()

	if *doTrain {
		train.Train(1)
		return
	}else if *doTrain2{
		train.Train(2)
		return
	}else if *doScore{
			if *fprot == "" || *flig == ""{
				fmt.Println("Protein or ligand not set")
				return
			}
			params := score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
			prot := protein.LoadMol2(*fprot)
			lig := ligand.LoadMol2(*flig)
			total := params.Score(&prot, &lig)
			fmt.Println(total)
	}
	// params := score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
	// // fmt.Println(params)
	// // protein := protein.LoadMol2("/home/jgcarvalho/pdbbind/core/proteinmol2/10gs_protein.mol2")
	// protein := protein.LoadMol2(*f_prot)
	// // fmt.Println(protein)
	// // ligand := ligand.LoadMol2("/home/jgcarvalho/pdbbind/core/ligands/10gs_ligand.mol2")
	// ligand := ligand.LoadMol2(*f_lig)
	// // fmt.Println(ligand)
	// total := params.Score(&protein, &ligand)
	// fmt.Println(total)
	return
}
