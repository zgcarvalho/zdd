package search

import (
	"fmt"
	"math/rand"
	"time"

	"bitbucket.org/jgcarvalho/zdd/ligand"
	"bitbucket.org/jgcarvalho/zdd/optimize"
	"bitbucket.org/jgcarvalho/zdd/protein"
	"bitbucket.org/jgcarvalho/zdd/score"
)

func nm(params *score.Parameters, l *ligand.Ligand, p *protein.Protein) {
	orig := l.Center()
	method := &optimize.NelderMead{}
	method.SimplexSize = 0.01
	initX := []float64{orig[0] + rand.Float64(), orig[1] + rand.Float64(), orig[2] + rand.Float64(), 0.0, 0.0, 0.0}
	problem := &optimize.Problem{}
	fmt.Println("INICIO", params.Score(p, l))
	fmt.Println("Center: ", orig)
	problem.Func = func(x []float64) float64 {
		l.Move(x[0], x[1], x[2], x[3]-l.Angles[0], x[4]-l.Angles[1], x[5]-l.Angles[2])
		fmt.Println(x)
		sc := params.Score(p, l)
		fmt.Println(sc)
		return sc
	}
	result, err := optimize.Local(*problem, initX, nil, method)
	if err != nil {
		fmt.Println("Erro minimização:", err)
	}
	fmt.Println("###RESULT:", result)
	fmt.Println("Location:", result.Location.X)
	l.Move(result.Location.X[0], result.Location.X[1], result.Location.X[2], result.Location.X[3], result.Location.X[4], result.Location.X[5])
	fmt.Println(params.Score(p, l))
	l.SavePDB("")
}

func Local() {}

func Global(p *protein.Protein, l *ligand.Ligand) {
	rand.Seed(time.Now().UnixNano())
	params := score.LoadParams("/home/jgcarvalho/gocode/src/bitbucket.org/jgcarvalho/zdd/params/INITPARAMS")
	for i := 0; i < 100; i++ {
		nm(&params, l, p)
	}

}
