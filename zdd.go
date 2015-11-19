package main

import (
	"bitbucket.org/jgcarvalho/zdd/ligand"
	"bitbucket.org/jgcarvalho/zdd/score"
)

func main() {
	params := score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
	// fmt.Println(params)
	protein := protein.LoadMol2()
	ligand := ligand.LoadMol2()
	total := params.Score(protein, ligand)

}
