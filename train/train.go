package train

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"bitbucket.org/jgcarvalho/zdd/ligand"
	"bitbucket.org/jgcarvalho/zdd/protein"
	"bitbucket.org/jgcarvalho/zdd/score"
)

type TrainData struct {
	Name      string
	Pkd       float64
	Receptor  string
	Positive  string
	Negatives []string
}

func loadData(fn string) []TrainData {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var traindata []TrainData
	err = json.Unmarshal(data, &traindata)
	if err != nil {
		fmt.Println("error:", err)
	}
	// cases := strings.Split(string(data), "\n")

	return traindata
}

func cost(params score.Parameters, traindata []TrainData) {
	for i := 0; i < len(traindata); i++ {
		protein := protein.LoadMol2(traindata[i].Receptor)
	}
	protein := protein.LoadMol2(*f_prot)
	// fmt.Println(protein)
	// ligand := ligand.LoadMol2("/home/jgcarvalho/pdbbind/core/ligands/10gs_ligand.mol2")
	ligand := ligand.LoadMol2(*f_lig)
	// fmt.Println(ligand)
	total := params.Score(&protein, &ligand)

}

func trainMain(params score.Parameters, traindata []TrainData) {

}

func Train() {
	fn, _ := filepath.Abs("./traindata/data.json")
	traindata := loadData(fn)
	fmt.Println(traindata)

}
