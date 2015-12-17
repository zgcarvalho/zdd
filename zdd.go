package main

import (
	"flag"
	"fmt"

	"bitbucket.org/jgcarvalho/zdd/ligand"
	"bitbucket.org/jgcarvalho/zdd/protein"
	"bitbucket.org/jgcarvalho/zdd/score"
	"bitbucket.org/jgcarvalho/zdd/search"
	"bitbucket.org/jgcarvalho/zdd/train"
)

func main() {
	doTrain := flag.Bool("train", false, "train using genetic algorithm")
	doTrain2 := flag.Bool("train2", false, "train using nelder mead")
	doTrain3 := flag.Bool("train3", false, "CMA-ES")
	doTrain4 := flag.Bool("train4", false, "CMA-ES without ranking")
	doTrain5 := flag.Bool("train5", false, "CMA-ES corr and ranking")
	doScore := flag.Bool("score", false, "score conformation")
	doDockGlobal := flag.Bool("dockG", false, "global docking")
	fprot := flag.String("p", "", "protein mol2")
	flig := flag.String("l", "", "ligand mol2")
	flag.Parse()

	if *doTrain {
		train.Train(1)
		return
	} else if *doTrain2 {
		train.Train(2)
		return
	} else if *doTrain3 {
		train.Train(3)
		return
	} else if *doTrain4 {
		train.Train(4)
		return
	} else if *doTrain5 {
		train.Train(5)
		return
	} else if *doScore {
		if *fprot == "" || *flig == "" {
			fmt.Println("Protein or ligand not set")
			return
		}
		params := score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
		prot := protein.LoadMol2(*fprot)
		lig := ligand.LoadMol2(*flig)
		total := params.Score(&prot, &lig)
		fmt.Println(total)
	} else if *doDockGlobal {
		if *fprot == "" || *flig == "" {
			fmt.Println("Protein or ligand not set")
			return
		}
		prot := protein.LoadMol2(*fprot)
		lig := ligand.LoadMol2(*flig)
		search.Global(&prot, &lig)

	}

	// teste de otimizacao
	// method := &optimize.CMAES{}
	// initX := []float64{1, 2, 1, 1, 1, 1}
	// problem := &optimize.Problem{}
	// problem.Func = func(x []float64) (sum float64) {
	// 	if len(x) != 6 {
	// 		panic("dimension of the problem must be 6")
	// 	}
	//
	// 	for i := 1; i <= 13; i++ {
	// 		z := float64(i) / 10
	// 		y := math.Exp(-z) - 5*math.Exp(-10*z) + 3*math.Exp(-4*z)
	// 		f := x[2]*math.Exp(-x[0]*z) - x[3]*math.Exp(-x[1]*z) + x[5]*math.Exp(-x[4]*z) - y
	// 		sum += f * f
	// 	}
	// 	return sum
	// }
	// result, err := optimize.Local(*problem, initX, nil, method)
	// if err != nil {
	// 	fmt.Println("Erro minimização:", err)
	// }
	// fmt.Println("###RESULT:", result)

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
