package train

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"

	"github.com/thoj/go-galib"

	"bitbucket.org/jgcarvalho/zdd/ligand"
	"bitbucket.org/jgcarvalho/zdd/optimize"
	"bitbucket.org/jgcarvalho/zdd/protein"
	"bitbucket.org/jgcarvalho/zdd/score"
	"github.com/gonum/stat"
)

var params score.Parameters
var trainset []TrainItem
var scores int

type TrainData struct {
	Name      string
	Energy    float64
	Receptor  string
	Positive  string
	Negatives []string
}

type TrainItem struct {
	Name      string
	Energy    float64
	Receptor  protein.Protein
	Positive  ligand.Ligand
	Negatives []ligand.Ligand
}

func loadData(fn string) []TrainItem {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var traindata []TrainData
	err = json.Unmarshal(data, &traindata)
	if err != nil {
		fmt.Println("error:", err)
	}
	trainset := make([]TrainItem, len(traindata))
	for i := 0; i < len(trainset); i++ {
		trainset[i].Name = traindata[i].Name
		trainset[i].Energy = traindata[i].Energy
		trainset[i].Receptor = protein.LoadMol2("./traindata/" + traindata[i].Receptor)
		trainset[i].Positive = ligand.LoadMol2("./traindata/" + traindata[i].Positive)
		trainset[i].Negatives = make([]ligand.Ligand, len(traindata[i].Negatives))
		for j := 0; j < len(trainset[i].Negatives); j++ {
			trainset[i].Negatives[j] = ligand.LoadMol2("./traindata/" + traindata[i].Negatives[j])
		}
	}
	// cases := strings.Split(string(data), "\n")

	return trainset
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
	penal := (g.Gene[0] * 15000.0) + 30000.0
	metalDbest := g.Gene[1] * 1.0
	metalAlpha := g.Gene[2] * -4500.0
	metalBeta := g.Gene[3] * 0.3
	repulsiveDbest := g.Gene[4] * 1.0
	repulsiveAlpha := g.Gene[5] * 50.0
	repulsiveBeta := g.Gene[6] * 0.3
	buriedDbest := g.Gene[7] * 1.0
	buriedAlpha := g.Gene[8] * -5.0
	buriedBeta := g.Gene[9] * 0.3
	hbondDbest := g.Gene[10] * 1.0
	hbondAlpha := g.Gene[11] * 0.1
	hbondBeta := g.Gene[12] * -1500.0
	haDbest := g.Gene[13] * 1.0
	haAlpha := g.Gene[14] * -1000.0
	haBeta := g.Gene[15] * 0.3
	harepDbest := g.Gene[16] * 1.0
	harepAlpha := g.Gene[17] * 100.0
	harepBeta := g.Gene[18] * 0.3
	npolarDbest := g.Gene[19] * 1.0
	npolarAlpha := g.Gene[20] * -50.0
	npolarBeta := g.Gene[21] * 0.3
	for k := range params.Inter {
		tmp := params.Inter[k]
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
		params.Inter[k] = tmp
	}
	fmt.Println(g.Gene)
	return cost(params, trainset)
}

func cost(params score.Parameters, trainset []TrainItem) float64 {
	var totalScore float64
	enerchan := make(chan float64, len(trainset))
	rankchan := make(chan float64, len(trainset))
	exp := make([]float64, len(trainset))
	obs := make([]float64, len(trainset))
	enerScore := 0.0
	rankScore := 0.0
	for i := 0; i < len(trainset); i++ {
		go func(i int) {
			// protein := protein.LoadMol2("./traindata/" + traindata[i].Receptor)
			protein := trainset[i].Receptor
			// pos := ligand.LoadMol2("./traindata/" + traindata[i].Positive)
			pos := trainset[i].Positive
			total := params.Score(&protein, &pos)

			rk := 0.0
			for j := 0; j < len(trainset[i].Negatives); j++ {
				// neg := ligand.LoadMol2("./traindata/" + traindata[i].Negatives[j])
				neg := trainset[i].Negatives[j]
				negTotal := params.Score(&protein, &neg)
				if negTotal <= total {
					rk += 1.0
				}
			}
			exp[i] = trainset[i].Energy
			obs[i] = total
			enerchan <- ((trainset[i].Energy - total) * (trainset[i].Energy - total))
			rankchan <- rk
		}(i)
	}

	for i := 0; i < len(trainset); i++ {
		enerScore += <-enerchan
		rankScore += <-rankchan
	}
	corr := stat.Correlation(exp, obs, nil)

	// totalScore = enerScore/(corr*corr) + (enerScore / (corr * corr) * rankScore)
	totalScore = enerScore*(2000.0-999.0*corr) + (math.Sqrt(enerScore) * (2000.0 - 1999.0*corr) * rankScore * 1000000.0)
	// fmt.Printf("PKD %f - Rank %f - Corr %f - TOTAL %f\n", enerScore, rankScore, corr, totalScore)
	fmt.Printf("Energy %f - Rank %f - Corr %f - TOTAL %f\n", math.Sqrt(enerScore/float64(len(trainset))), rankScore, corr, totalScore)
	return totalScore
}

func cost2(params score.Parameters, trainset []TrainItem) float64 {
	var totalScore float64
	enerchan := make(chan float64, len(trainset))
	// rankchan := make(chan float64, len(trainset))
	exp := make([]float64, len(trainset))
	obs := make([]float64, len(trainset))
	enerScore := 0.0
	// rankScore := 0.0
	for i := 0; i < len(trainset); i++ {
		go func(i int) {
			// protein := protein.LoadMol2("./traindata/" + traindata[i].Receptor)
			protein := trainset[i].Receptor
			// pos := ligand.LoadMol2("./traindata/" + traindata[i].Positive)
			pos := trainset[i].Positive
			total := params.Score(&protein, &pos)

			// rk := 0.0
			// for j := 0; j < len(trainset[i].Negatives); j++ {
			// 	// neg := ligand.LoadMol2("./traindata/" + traindata[i].Negatives[j])
			// 	neg := trainset[i].Negatives[j]
			// 	negTotal := params.Score(&protein, &neg)
			// 	if negTotal <= total {
			// 		rk += 1.0
			// 	}
			// }
			exp[i] = trainset[i].Energy
			obs[i] = total
			enerchan <- ((trainset[i].Energy - total) * (trainset[i].Energy - total))
			// rankchan <- rk
		}(i)
	}

	for i := 0; i < len(trainset); i++ {
		enerScore += <-enerchan
		// rankScore += <-rankchan
	}
	corr := stat.Correlation(exp, obs, nil)

	// totalScore = enerScore/(corr*corr) + (enerScore / (corr * corr) * rankScore)
	// totalScore = enerScore*(2000.0-1999.0*corr) + (math.Sqrt(enerScore) * (2000.0 - 1999.0*corr) * rankScore * 1000000.0)
	totalScore = enerScore * (2000.0 - 1999.0*corr)
	// fmt.Printf("PKD %f - Rank %f - Corr %f - TOTAL %f\n", enerScore, rankScore, corr, totalScore)
	fmt.Printf("Energy %f - Corr %f - TOTAL %f\n", math.Sqrt(enerScore/float64(len(trainset))), corr, totalScore)
	return totalScore
}

func trainMain2() {
	method := &optimize.NelderMead{}
	method.Shrink = 0.95
	method.Contraction = 0.95
	method.Reflection = 2.0
	method.Expansion = 2.0
	method.SimplexSize = 1.0
	initX := []float64{0.5, 1.5, 3.0, 2.0, 2.0, 3.0, 4.0, 3.5, 2.0, 3.0, 2.0, 4.5, 3.5, 2.5, 3.0, 3.5, 3.5, 3.0, 3.0, 3.5, 4.0, 4.0}
	problem := &optimize.Problem{}
	problem.Func = func(x []float64) float64 {
		penal := (x[0] * 15000.0) + 30000.0
		metalDbest := x[1] * 1.0
		metalAlpha := x[2] * -4500.0
		metalBeta := x[3] * 0.3
		repulsiveDbest := x[4] * 1.0
		repulsiveAlpha := x[5] * 50.0
		repulsiveBeta := x[6] * 0.3
		buriedDbest := x[7] * 1.0
		buriedAlpha := x[8] * -5.0
		buriedBeta := x[9] * 0.3
		hbondDbest := x[10] * 1.0
		hbondAlpha := x[11] * 0.1
		hbondBeta := x[12] * -1500.0
		haDbest := x[13] * 1.0
		haAlpha := x[14] * -1000.0
		haBeta := x[15] * 0.3
		harepDbest := x[16] * 1.0
		harepAlpha := x[17] * 100.0
		harepBeta := x[18] * 0.3
		npolarDbest := x[19] * 1.0
		npolarAlpha := x[20] * -50.0
		npolarBeta := x[21] * 0.3
		for k := range params.Inter {
			tmp := params.Inter[k]
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
			params.Inter[k] = tmp
		}
		fmt.Println(x)
		return cost(params, trainset)
	}

	result, err := optimize.Local(*problem, initX, nil, method)
	if err != nil {
		fmt.Println("Erro minimização:", err)
	}
	fmt.Println("###RESULT:", result)
}

func trainMain3() {
	method := &optimize.CMAES{}
	// initX := []float64{0.5, 1.5, 3.0, 2.0, 2.0, 3.0, 4.0, 3.5, 2.0, 3.0, 2.0, 4.5, 3.5, 2.5, 3.0, 3.5, 3.5, 3.0, 3.0, 3.5, 4.0, 4.0}
	initX := []float64{0.44439817749423544, 1.861654501650633, 3.111182936447595, 1.662812245597922, 1.859229396738397, 2.596714801210651, 3.6612720109641352, 3.4187361480475005, 1.815317344258246, 3.1602223435925882, 2.0340215689027836, 4.008073061952763, 3.834808533112463, 2.349631328558668, 3.1409644294863917, 3.3930732169666062, 3.3577472259547574, 2.7405711507780564, 2.7396958878034745, 3.802195324120163, 4.555449200453119, 4.335640603745485, 0.0, 0.0, 0.0, 0.0, 0.0}
	problem := &optimize.Problem{}
	problem.Func = func(x []float64) float64 {
		penal := (x[0] * 15000.0) + 30000.0
		metalDbest := x[1] * 1.0
		metalAlpha := x[2] * -4500.0
		metalBeta := x[3] * 0.3
		repulsiveDbest := x[4] * 1.0
		repulsiveAlpha := x[5] * 50.0
		repulsiveBeta := x[6] * 0.3
		buriedDbest := x[7] * 1.0
		buriedAlpha := x[8] * -5.0
		buriedBeta := x[9] * 0.3
		hbondDbest := x[10] * 1.0
		hbondAlpha := x[11] * 0.1
		hbondBeta := x[12] * -1500.0
		haDbest := x[13] * 1.0
		haAlpha := x[14] * -1000.0
		haBeta := x[15] * 0.3
		harepDbest := x[16] * 1.0
		harepAlpha := x[17] * 100.0
		harepBeta := x[18] * 0.3
		npolarDbest := x[19] * 1.0
		npolarAlpha := x[20] * -50.0
		npolarBeta := x[21] * 0.3
		params.Wnp = x[22] * 1000.0
		params.Enp = x[23]
		params.Wp = x[24] * 1000.0
		params.Ep = x[25]
		// params.Wrb = x[26]
		// params.Erb = x[27]
		params.C = x[26] * 1000.0
		for k := range params.Inter {
			tmp := params.Inter[k]
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
			params.Inter[k] = tmp
		}
		fmt.Println(x)
		return cost(params, trainset)
	}

	result, err := optimize.Local(*problem, initX, nil, method)
	if err != nil {
		fmt.Println("Erro minimização:", err)
	}
	fmt.Println("###RESULT:", result)
}

func trainMain4() {
	method := &optimize.CMAES{}
	// initX := []float64{0.5, 1.5, 3.0, 2.0, 2.0, 3.0, 4.0, 3.5, 2.0, 3.0, 2.0, 4.5, 3.5, 2.5, 3.0, 3.5, 3.5, 3.0, 3.0, 3.5, 4.0, 4.0}
	initX := []float64{0.44439817749423544, 1.861654501650633, 3.111182936447595, 1.662812245597922, 1.859229396738397, 2.596714801210651, 3.6612720109641352, 3.4187361480475005, 1.815317344258246, 3.1602223435925882, 2.0340215689027836, 4.008073061952763, 3.834808533112463, 2.349631328558668, 3.1409644294863917, 3.3930732169666062, 3.3577472259547574, 2.7405711507780564, 2.7396958878034745, 3.802195324120163, 4.555449200453119, 4.335640603745485, 0.0, 0.0, 0.0, 0.0, 0.0}
	problem := &optimize.Problem{}
	problem.Func = func(x []float64) float64 {
		penal := (x[0] * 15000.0) + 30000.0
		metalDbest := x[1] * 1.0
		metalAlpha := x[2] * -4500.0
		metalBeta := x[3] * 0.3
		repulsiveDbest := x[4] * 1.0
		repulsiveAlpha := x[5] * 50.0
		repulsiveBeta := x[6] * 0.3
		buriedDbest := x[7] * 1.0
		buriedAlpha := x[8] * -5.0
		buriedBeta := x[9] * 0.3
		hbondDbest := x[10] * 1.0
		hbondAlpha := x[11] * 0.1
		hbondBeta := x[12] * -1500.0
		haDbest := x[13] * 1.0
		haAlpha := x[14] * -1000.0
		haBeta := x[15] * 0.3
		harepDbest := x[16] * 1.0
		harepAlpha := x[17] * 100.0
		harepBeta := x[18] * 0.3
		npolarDbest := x[19] * 1.0
		npolarAlpha := x[20] * -50.0
		npolarBeta := x[21] * 0.3
		params.Wnp = x[22] * 1000.0
		params.Enp = x[23]
		params.Wp = x[24] * 1000.0
		params.Ep = x[25]
		// params.Wrb = x[26]
		// params.Erb = x[27]
		params.C = x[26] * 1000.0
		for k := range params.Inter {
			tmp := params.Inter[k]
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
			params.Inter[k] = tmp
		}
		fmt.Println(x)
		return cost2(params, trainset)
	}

	result, err := optimize.Local(*problem, initX, nil, method)
	if err != nil {
		fmt.Println("Erro minimização:", err)
	}
	fmt.Println("###RESULT:", result)
}

func Train(method int) {
	params = score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
	fn, err := filepath.Abs("./traindata/data.json")
	if err != nil {
		fmt.Println("Arquivo com dados de treinamento", err)
		panic(err)
	}
	trainset = loadData(fn)
	switch method {
	case 1:
		trainMain()
	case 2:
		trainMain2()
	case 3:
		trainMain3()
	case 4:
		trainMain4()
	default:
		fmt.Println("Method unselected")
	}
}
