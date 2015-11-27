package train

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	spsa "github.com/yanatan16/golang-spsa"
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

func Train() {
	fn, _ := filepath.Abs("./traindata/data.json")
	traindata := loadData(fn)
	fmt.Println(traindata)

	spsa := &spsa.SPSA{
		L:     spsa.AbsoluteSum,              // Loss Function
		C:     spsa.NoConstraints,            // Constraint Function
		Theta: spsa.Vector{1, 1, 1, 1, 1},    // Initial theta vector
		Ak:    spsa.StandardAk(1, 100, .602), // a tuned, A ~= n / 10, alpha = .602
		Ck:    spsa.StandardCk(.1, .101),     // c ~= std-dev(Loss function), gamma = .101
		Delta: spsa.Bernoulli{1},             // Perturbation Distribution
	}

	theta := spsa.Run(1000)
}
