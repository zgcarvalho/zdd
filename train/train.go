package train

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "math/rand"
	"path/filepath"

	"github.com/thoj/go-galib"

	"bitbucket.org/jgcarvalho/zdd/ligand"
	"bitbucket.org/jgcarvalho/zdd/protein"
	"bitbucket.org/jgcarvalho/zdd/score"
	"github.com/gonum/optimize"
	"github.com/gonum/stat"

)

var params score.Parameters
var traindata []TrainData
var scores int

// func init() {
// 	params = score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
// 	fn, err := filepath.Abs("./traindata/data.json")
// 	if err != nil {
// 		fmt.Println("Arquivo com dados de treinamento", err)
// 		panic(err)
// 	}
// 	traindata = loadData(fn)
// }

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

func cost(params score.Parameters, traindata []TrainData) float64 {
	var totalScore float64
	pkdchan := make(chan float64, len(traindata))
	rankchan := make(chan float64, len(traindata))
	exp := make([]float64,len(traindata))
 	obs := make([]float64,len(traindata))
	pkdScore := 0.0
	rankScore := 0.0
	for i := 0; i < len(traindata); i++ {
		go func(i int) {
			protein := protein.LoadMol2("./traindata/" + traindata[i].Receptor)
			pos := ligand.LoadMol2("./traindata/" + traindata[i].Positive)
			total := params.Score(&protein, &pos)

			rk := 0.0
			for j := 0; j < len(traindata[i].Negatives); j++ {
				neg := ligand.LoadMol2("./traindata/" + traindata[i].Negatives[j])
				negTotal := params.Score(&protein, &neg)
				if negTotal >= total {
					rk += 0.1
				}
			}
			exp[i] = traindata[i].Pkd
			obs[i] = total
			pkdchan <- (traindata[i].Pkd - total) * (traindata[i].Pkd - total)
			rankchan <- rk
		}(i)
	}

	for i := 0; i < len(traindata); i++ {
		pkdScore += <-pkdchan
		rankScore += <-rankchan
	}
	corr := stat.Correlation(exp,obs,nil)
	// pkdScore = pkdScore
	// totalScore = pkdScore * rankScore
	totalScore = pkdScore/corr + (pkdScore/corr * rankScore)
	fmt.Printf("PKD %f - Rank %f - Corr %f- TOTAL %f\n", pkdScore, rankScore, corr, totalScore)
	return totalScore

}

func trainMain2() {
	method := &optimize.NelderMead{}
	// method.Shrink = 0.95
	// method.Contraction = 0.95
	// method.Reflection = 2.0
	// method.Expansion = 2.0
	method.SimplexSize = 0.5
	initX := []float64{1.5, 2.5, 3.6, 2.0, 3.0, 3.0, 4.0, 3.6, 2.0, 3.5, 3.3, 4.0, 3.5, 2.5, 3.0, 3.5, 3.4, 3.0, 3.0, 3.6, 3.0, 3.0}
	problem := &optimize.Problem{}
	problem.Func = func(x []float64) float64 {
		penal := (x[0] * 3.0) + 5.0
		metalDbest := x[1] * 1.0
		metalAlpha := x[2] * 0.2
		metalBeta := x[3] * 0.3
		repulsiveDbest := x[4] * 1.0
		repulsiveAlpha := x[5] * -0.01
		repulsiveBeta := x[6] * 0.3
		buriedDbest := x[7] * 1.0
		buriedAlpha := x[8] * 0.001
		buriedBeta := x[9] * 0.3
		hbondDbest := x[10] * 1.0
		hbondAlpha := x[11] * 0.1
		hbondBeta := x[12] * 0.3
		haDbest := x[13] * 1.0
		haAlpha := x[14] * 0.1
		haBeta := x[15] * 0.3
		harepDbest := x[16] * 1.0
		harepAlpha := x[17] * -0.002
		harepBeta := x[18] * 0.3
		npolarDbest := x[19] * 1.0
		npolarAlpha := x[20] * 0.01
		npolarBeta := x[21] * 0.3
		for k := range params {
			tmp := params[k]
			tmp.Penal = penal
			if tmp.Type == "metal" {
				tmp.Dbest = metalDbest
				tmp.Alpha = metalAlpha
				tmp.Beta = metalBeta
			} else if tmp.Type == "repulsive" {
				tmp.Dbest = repulsiveDbest
				tmp.Alpha = repulsiveAlpha
				tmp.Beta = repulsiveBeta
			} else if tmp.Type == "buried" {
				tmp.Dbest = buriedDbest
				tmp.Alpha = buriedAlpha
				tmp.Beta = buriedBeta
			} else if tmp.Type == "hbond" {
				tmp.Dbest = hbondDbest
				tmp.Alpha = hbondAlpha
				tmp.Beta = hbondBeta
			} else if tmp.Type == "ha" {
				tmp.Dbest = haDbest
				tmp.Alpha = haAlpha
				tmp.Beta = haBeta
			} else if tmp.Type == "ha-repulsive" {
				tmp.Dbest = harepDbest
				tmp.Alpha = harepAlpha
				tmp.Beta = harepBeta
			} else if tmp.Type == "npolar" {
				tmp.Dbest = npolarDbest
				tmp.Alpha = npolarAlpha
				tmp.Beta = npolarBeta
			} else {
				fmt.Println("Que tipo é esse?", tmp.Type)
			}
			params[k] = tmp
		}
		fmt.Println(x)
		return cost(params, traindata)
	}

	result, err := optimize.Local(*problem, initX, nil, method)
	if err != nil {
		fmt.Println("Erro minimização:", err)
	}
	fmt.Println("###RESULT:", result)
}

func trainMain() {

	gaparam := ga.GAParameter{
		Initializer: new(ga.GARandomInitializer),
		Selector:    ga.NewGATournamentSelector(0.2, 5),
		Breeder:     new(ga.GA2PointBreeder),
		Mutator:     ga.NewGAGaussianMutator(0.4, 0),
		PMutate:     0.5,
		PBreed:      0.2}

	gao := ga.NewGA(gaparam)
	genome := ga.NewFloatGenome(make([]float64, 22), sfcost, 6, 0)
	gao.Init(250, genome)
	gao.OptimizeUntil(func(best ga.GAGenome) bool {
		return best.Score() < 1e-3
	})
	best := gao.Best().(*ga.GAFloatGenome)
	fmt.Printf("%s = %f\n", best, best.Score())
	fmt.Printf("Calls to score = %d\n", scores)
}

func sfcost(g *ga.GAFloatGenome) float64 {
	penal := (g.Gene[0] * 3.0) + 5.0
	metalDbest := g.Gene[1] * 1.0
	metalAlpha := g.Gene[2] * 0.2
	metalBeta := g.Gene[3] * 0.3
	repulsiveDbest := g.Gene[4] * 1.0
	repulsiveAlpha := g.Gene[5] * -0.01
	repulsiveBeta := g.Gene[6] * 0.3
	buriedDbest := g.Gene[7] * 1.0
	buriedAlpha := g.Gene[8] * 0.001
	buriedBeta := g.Gene[9] * 0.3
	hbondDbest := g.Gene[10] * 1.0
	hbondAlpha := g.Gene[11] * 0.1
	hbondBeta := g.Gene[12] * 0.3
	haDbest := g.Gene[13] * 1.0
	haAlpha := g.Gene[14] * 0.1
	haBeta := g.Gene[15] * 0.3
	harepDbest := g.Gene[16] * 1.0
	harepAlpha := g.Gene[17] * -0.002
	harepBeta := g.Gene[18] * 0.3
	npolarDbest := g.Gene[19] * 1.0
	npolarAlpha := g.Gene[20] * 0.01
	npolarBeta := g.Gene[21] * 0.3
	for k := range params {
		tmp := params[k]
		tmp.Penal = penal
		if tmp.Type == "metal" {
			tmp.Dbest = metalDbest
			tmp.Alpha = metalAlpha
			tmp.Beta = metalBeta
		} else if tmp.Type == "repulsive" {
			tmp.Dbest = repulsiveDbest
			tmp.Alpha = repulsiveAlpha
			tmp.Beta = repulsiveBeta
		} else if tmp.Type == "buried" {
			tmp.Dbest = buriedDbest
			tmp.Alpha = buriedAlpha
			tmp.Beta = buriedBeta
		} else if tmp.Type == "hbond" {
			tmp.Dbest = hbondDbest
			tmp.Alpha = hbondAlpha
			tmp.Beta = hbondBeta
		} else if tmp.Type == "ha" {
			tmp.Dbest = haDbest
			tmp.Alpha = haAlpha
			tmp.Beta = haBeta
		} else if tmp.Type == "ha-repulsive" {
			tmp.Dbest = harepDbest
			tmp.Alpha = harepAlpha
			tmp.Beta = harepBeta
		} else if tmp.Type == "npolar" {
			tmp.Dbest = npolarDbest
			tmp.Alpha = npolarAlpha
			tmp.Beta = npolarBeta
		} else {
			fmt.Println("Que tipo é esse?", tmp.Type)
		}
		params[k] = tmp
	}
	fmt.Println(g.Gene)
	return cost(params, traindata)
}

func Train(method int) {
	params = score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
	fn, err := filepath.Abs("./traindata/data.json")
	if err != nil {
		fmt.Println("Arquivo com dados de treinamento", err)
		panic(err)
	}
	traindata = loadData(fn)
	switch method {
	case 1:
		trainMain()
	case 2:
		trainMain2()
	default:
		fmt.Println("Method unselected")
	}
}
