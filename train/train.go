package train

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
  "math"
	"math/rand"
	"bitbucket.org/jgcarvalho/zdd/ligand"
	"bitbucket.org/jgcarvalho/zdd/protein"
	"bitbucket.org/jgcarvalho/zdd/score"
	"github.com/gonum/optimize"
)

type SF struct{}

func (SF) Func(x []float64) float64 {
	params := score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
	fn, err := filepath.Abs("./traindata/data_teste.json")
	if err != nil {
		fmt.Println("Arquivo com dados de treinamento", err)
		panic(err)
	}
	traindata := loadData(fn)
	penal := x[0]
	metalDbest := x[1]
	metalAlpha := x[2]
	metalBeta := x[3]
	repulsiveDbest := x[4]
	repulsiveAlpha := x[5]
	repulsiveBeta := x[6]
	buriedDbest := x[7]
	buriedAlpha := x[8]
	buriedBeta := x[9]
	hbondDbest := x[10]
	hbondAlpha := x[11]
	hbondBeta := x[12]
	haDbest := x[13]
	haAlpha := x[14]
	haBeta := x[15]
	harepDbest := x[16]
	harepAlpha := x[17]
	harepBeta := x[18]
	npolarDbest := x[19]
	npolarAlpha := x[20]
	npolarBeta := x[21]
	for k,_ := range params {
		tmp := params[k]
		tmp.Penal = penal
		if tmp.Type == "metal" {
			tmp.Dbest = metalDbest
			tmp.Alpha = metalAlpha
			tmp.Beta = metalBeta
		}else if tmp.Type == "repulsive" {
			tmp.Dbest = repulsiveDbest
			tmp.Alpha = repulsiveAlpha
			tmp.Beta = repulsiveBeta
		}else if tmp.Type == "buried" {
			tmp.Dbest = buriedDbest
			tmp.Alpha = buriedAlpha
			tmp.Beta = buriedBeta
		}else if tmp.Type == "hbond" {
			tmp.Dbest = hbondDbest
			tmp.Alpha = hbondAlpha
			tmp.Beta = hbondBeta
		}else if tmp.Type == "ha" {
			tmp.Dbest = haDbest
			tmp.Alpha = haAlpha
			tmp.Beta = haBeta
		}else if tmp.Type == "ha-repulsive" {
			tmp.Dbest = harepDbest
			tmp.Alpha = harepAlpha
			tmp.Beta = harepBeta
		}else if tmp.Type == "npolar" {
			tmp.Dbest = npolarDbest
			tmp.Alpha = npolarAlpha
			tmp.Beta = npolarBeta
		}else{
			fmt.Println("Que tipo é esse?", tmp.Type)
		}
		params[k] = tmp
	}
	return cost(params,traindata)
}

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
	var score float64
	pkdScore := 0.0
	rankScore := 0.0
	for i := 0; i < len(traindata); i++ {
		protein := protein.LoadMol2("./traindata/"+ traindata[i].Receptor)
		pos := ligand.LoadMol2("./traindata/"+traindata[i].Positive)
		total := params.Score(&protein, &pos)
		//peso 10 para erros no pkd
		pkdScore += math.Abs(traindata[i].Pkd - total)
		// fmt.Println(traindata[i].Receptor)
		// fmt.Println(traindata[i].Positive)
		// fmt.Println(total)
		for j:= 0; j < len(traindata[i].Negatives); j++ {
			neg := ligand.LoadMol2("./traindata/"+traindata[i].Negatives[j])
			negTotal := params.Score(&protein, &neg)
			if negTotal >= total {
				rankScore += 1.0
			}
			// fmt.Println(traindata[i].Negatives[j])
			// fmt.Println(negTotal)
		}

	}
		fmt.Printf("PKD %f - Rank %f - TOTAL %f\n", pkdScore, rankScore, pkdScore+rankScore)
		score = pkdScore + rankScore
		return score

}

func genNewMainParams(params score.Parameters) score.Parameters{
	penal := rand.NormFloat64()*0.01 + params["Al_Br"].Penal
	metalDbest := rand.NormFloat64()*0.01 + params["Al_Br"].Dbest
	metalAlpha := rand.NormFloat64()*0.01 + params["Al_Br"].Alpha
	metalBeta := rand.NormFloat64()*0.01 + params["Al_Br"].Beta
	repulsiveDbest := rand.NormFloat64()*0.01 + params["Al_Ca"].Dbest
	repulsiveAlpha := rand.NormFloat64()*0.01 + params["Al_Ca"].Alpha
	repulsiveBeta := rand.NormFloat64()*0.01 + params["Al_Ca"].Beta
	buriedDbest := rand.NormFloat64()*0.01 + params["Al_C.1"].Dbest
	buriedAlpha := rand.NormFloat64()*0.01 + params["Al_C.1"].Alpha
	buriedBeta := rand.NormFloat64()*0.01 + params["Al_C.1"].Beta
	hbondDbest := rand.NormFloat64()*0.01 + params["N.1_O.3"].Dbest
	hbondAlpha := rand.NormFloat64()*0.01 + params["N.1_O.3"].Alpha
	hbondBeta := rand.NormFloat64()*0.01 + params["N.1_O.3"].Beta
	haDbest := rand.NormFloat64()*0.01 + params["Br_O.3"].Dbest
	haAlpha := rand.NormFloat64()*0.01 + params["Br_O.3"].Alpha
	haBeta := rand.NormFloat64()*0.01 + params["Br_O.3"].Beta
	harepDbest := rand.NormFloat64()*0.01 + params["Br_Br"].Dbest
	harepAlpha := rand.NormFloat64()*0.01 + params["Br_Br"].Alpha
	harepBeta := rand.NormFloat64()*0.01 + params["Br_Br"].Beta
	npolarDbest := rand.NormFloat64()*0.01 + params["C.1_C.1"].Dbest
	npolarAlpha := rand.NormFloat64()*0.01 + params["C.1_C.1"].Alpha
	npolarBeta := rand.NormFloat64()*0.01 + params["C.1_C.1"].Beta
	for k,_ := range params {
		tmp := params[k]
		tmp.Penal = penal
		if tmp.Type == "metal" {
			tmp.Dbest = metalDbest
			tmp.Alpha = metalAlpha
			tmp.Beta = metalBeta
		}else if tmp.Type == "repulsive" {
			tmp.Dbest = repulsiveDbest
			tmp.Alpha = repulsiveAlpha
			tmp.Beta = repulsiveBeta
		}else if tmp.Type == "buried" {
			tmp.Dbest = buriedDbest
			tmp.Alpha = buriedAlpha
			tmp.Beta = buriedBeta
		}else if tmp.Type == "hbond" {
			tmp.Dbest = hbondDbest
			tmp.Alpha = hbondAlpha
			tmp.Beta = hbondBeta
		}else if tmp.Type == "ha" {
			tmp.Dbest = haDbest
			tmp.Alpha = haAlpha
			tmp.Beta = haBeta
		}else if tmp.Type == "ha-repulsive" {
			tmp.Dbest = harepDbest
			tmp.Alpha = harepAlpha
			tmp.Beta = harepBeta
		}else if tmp.Type == "npolar" {
			tmp.Dbest = npolarDbest
			tmp.Alpha = npolarAlpha
			tmp.Beta = npolarBeta
		}else{
			fmt.Println("Que tipo é esse?", tmp.Type)
		}
		params[k] = tmp
	}
	return params
}
func trainMain2() {
	method := &optimize.NelderMead{}
	// method.Shrink = 0.95
	// method.Contraction = 0.95
	// method.Reflection = 2.0
	initX := []float64{10.0,3.2,0.7,1.0,3.2,-0.002,1.0,3.6,0.001,1.0,3.2,0.4,1.0,3.2,0.2,1.0,3.2,-0.001,3.2,3.6,0.01,1.0}

	problem := &optimize.Problem{}
	problem.Func = func(x []float64) float64 {
		params := score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
		fn, err := filepath.Abs("./traindata/data.json")
		if err != nil {
			fmt.Println("Arquivo com dados de treinamento", err)
			panic(err)
		}
		traindata := loadData(fn)
		penal := x[0]
		metalDbest := x[1]
		metalAlpha := x[2]
		metalBeta := x[3]
		repulsiveDbest := x[4]
		repulsiveAlpha := x[5]
		repulsiveBeta := x[6]
		buriedDbest := x[7]
		buriedAlpha := x[8]
		buriedBeta := x[9]
		hbondDbest := x[10]
		hbondAlpha := x[11]
		hbondBeta := x[12]
		haDbest := x[13]
		haAlpha := x[14]
		haBeta := x[15]
		harepDbest := x[16]
		harepAlpha := x[17]
		harepBeta := x[18]
		npolarDbest := x[19]
		npolarAlpha := x[20]
		npolarBeta := x[21]
		for k,_ := range params {
			tmp := params[k]
			tmp.Penal = penal
			if tmp.Type == "metal" {
				tmp.Dbest = metalDbest
				tmp.Alpha = metalAlpha
				tmp.Beta = metalBeta
			}else if tmp.Type == "repulsive" {
				tmp.Dbest = repulsiveDbest
				tmp.Alpha = repulsiveAlpha
				tmp.Beta = repulsiveBeta
			}else if tmp.Type == "buried" {
				tmp.Dbest = buriedDbest
				tmp.Alpha = buriedAlpha
				tmp.Beta = buriedBeta
			}else if tmp.Type == "hbond" {
				tmp.Dbest = hbondDbest
				tmp.Alpha = hbondAlpha
				tmp.Beta = hbondBeta
			}else if tmp.Type == "ha" {
				tmp.Dbest = haDbest
				tmp.Alpha = haAlpha
				tmp.Beta = haBeta
			}else if tmp.Type == "ha-repulsive" {
				tmp.Dbest = harepDbest
				tmp.Alpha = harepAlpha
				tmp.Beta = harepBeta
			}else if tmp.Type == "npolar" {
				tmp.Dbest = npolarDbest
				tmp.Alpha = npolarAlpha
				tmp.Beta = npolarBeta
			}else{
				fmt.Println("Que tipo é esse?", tmp.Type)
			}
			params[k] = tmp
		}
		return cost(params,traindata)
	}
	result, err := optimize.Local(*problem, initX,nil, method)
	if err != nil {
		fmt.Println("Erro minimização:",err)
	}
	fmt.Println("###RESULT:", result)
}

func trainMain(params score.Parameters, traindata []TrainData) {
	c := cost(params, traindata)
	fmt.Println(c)
	for i := 0; i < 20; i++ {
		pnew := genNewMainParams(params)
		newC := cost(pnew, traindata)
		fmt.Println(newC)
		if newC < c {
			c = newC
			params = pnew
			fmt.Println("solução melhor")
		}else{
			fmt.Println("solução pior")
		}
	}

}

func Train(params score.Parameters) {
	// fn, err := filepath.Abs("./traindata/data_teste.json")
	// if err != nil {
	// 	fmt.Println("Arquivo com dados de treinamento", err)
	// 	panic(err)
	// }
	// traindata := loadData(fn)
	// fmt.Println(traindata)
	// trainMain(params, traindata)
	trainMain2()

}
