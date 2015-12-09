// cma-es
package optimize

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/gonum/matrix/mat64"
)

type CMAES struct {
	N       int
	Xmean   []float64
	Sigma   float64
	Target  float64
	MaxEval int
	// Randn
	Lambda  int
	Mu      int
	Weights []float64
	Mueff   float64
	Cc      float64
	Cs      float64
	C1      float64
	Cmu     float64
	Damps   float64
	Pc      []float64
	Ps      []float64

	B         *mat64.Dense
	D         *mat64.Dense
	C         *mat64.Dense
	InvSqrtC  *mat64.Dense
	EigenEval int
	CountEval int
	ChiN      float64
	// FitVals []
	Population []Candidate
	// Best       Candidate
}

type Candidate struct {
	X   []float64
	Fit float64
}

func (c *CMAES) Init(loc *Location) (Operation, error) {
	c.N = len(loc.X)
	c.Xmean = make([]float64, c.N)
	copy(c.Xmean, loc.X)
	c.Sigma = 0.1          // input do usuario
	c.Target = 0.000000001 //input do usuario
	c.MaxEval = 1000       //input do usuario
	// c.Randn

	// Strategy parameter setting: Selection
	c.Lambda = 4 + int(3*math.Log(float64(c.N)))
	c.Mu = c.Lambda / 2
	c.Weights = make([]float64, c.Mu)
	sumWtmp := 0.0
	for i := 0; i < len(c.Weights); i++ {
		c.Weights[i] = math.Log(float64(c.Mu)+0.5) - math.Log(float64(i)+1.0)
		sumWtmp += c.Weights[i]
	}
	sumW := 0.0
	sumWSquare := 0.0
	for i := 0; i < len(c.Weights); i++ {
		c.Weights[i] /= sumWtmp
		sumW += c.Weights[i]
		sumWSquare += (c.Weights[i] * c.Weights[i])
	}
	c.Mueff = sumW / sumWSquare

	// Strategy parameter setting: Adaptation
	c.Cc = (4.0 + c.Mueff/float64(c.N)) / (float64(c.N) + 4.0 + 2.0*c.Mueff/float64(c.N))
	c.Cs = (c.Mueff + 2.0) / (float64(c.N) + c.Mueff + 5.0)
	c.C1 = 2.0 / ((float64(c.N)+1.3)*(float64(c.N)+1.3) + c.Mueff)
	c.Cmu = math.Min((1 - c.C1), (2.0 * (c.Mueff - 2.0 + 1.0/c.Mueff) / (((float64(c.N) + 2) * (float64(c.N) + 2)) + c.Mueff)))
	c.Damps = 2.0*c.Mueff/float64(c.Lambda) + 0.3 + c.Cs

	// Initialize dynamic (internal) state variables
	c.Pc = make([]float64, c.N)
	c.Ps = make([]float64, c.N)
	c.B = mat64.NewDense(c.N, c.N, nil)
	for i := 0; i < c.N; i++ {
		c.B.Set(i, i, 1.0)
	}
	c.D = mat64.NewDense(c.N, c.N, nil)
	for i := 0; i < c.N; i++ {
		c.D.Set(i, i, 1.0)
	}
	c.C = mat64.NewDense(c.N, c.N, nil)
	c.C.MulElem(c.D, c.D)
	c.C.Mul(c.B, c.C)
	c.C.Mul(c.C, c.B.T())
	// c.C = mat64.Mul(mat64.Mul(c.B, mat64.MulElem(c.D, c.D)), c.B.T())
	// c.C = stat.CovarianceMatrix(c.C, )
	// for i := 0; i < c.N; i++ {
	// 	c.C.Set(i, i, 1.0)
	// }
	c.InvSqrtC = mat64.NewDense(c.N, c.N, nil)
	c.InvSqrtC.Inverse(c.D)
	c.InvSqrtC.Mul(c.B, c.InvSqrtC)
	c.InvSqrtC.Mul(c.InvSqrtC, c.B.T())
	// for i := 0; i < c.N; i++ {
	// 	c.InvSqrtC.Set(i, i, 1.0)
	// }
	c.EigenEval = 0
	c.CountEval = 0
	c.ChiN = math.Pow(float64(c.N), 0.5*(1.0-1.0/(4.0*float64(c.N))+1.0/21.0*float64(c.N)*float64(c.N)))

	c.Population = make([]Candidate, c.Lambda)
	for i := 0; i < c.Lambda; i++ {
		c.Population[i].X = make([]float64, c.N)
	}
	// c.Fits = make([]float64, c.Lambda)
	// FitVals
	// best
	return FuncEvaluation, nil
}

var currLambda int

func (c *CMAES) Iterate(loc *Location) (Operation, error) {
	fmt.Println(loc.X)
	fmt.Println(loc.F)

	if currLambda < c.Lambda {
		fmt.Printf("Candidate %d of %d\n", currLambda+1, c.Lambda)
		copy(c.Population[currLambda].X, loc.X)
		c.Population[currLambda].Fit = loc.F
		currLambda += 1
	}
	if currLambda == c.Lambda {
		fmt.Printf("Generation %d\n", c.CountEval/c.Lambda)
		c.tell()
		// fmt.Println(c.Population)
		currLambda = 0
	}
	loc.X = c.ask()

	return FuncEvaluation, nil
}

// erro no ask -> não devo passar loc.X como x
func (c *CMAES) ask() []float64 {

	lenC, _ := c.C.Dims()
	if float64(c.CountEval-c.EigenEval) > float64(c.Lambda)/(c.C1+c.Cmu)/float64(lenC)/10.0 {

		c.EigenEval = c.CountEval

		c.C.Add(c.C, c.C.T())
		c.C.Scale(0.5, c.C)

		eig := mat64.Eigen(c.C, 0.000001)
		// fmt.Println("here")
		c.D = eig.D()
		c.B = eig.V
		// fmt.Println("# Entrou no loop")
		// fmt.Println("# ask c.B", c.B)
		// fmt.Println("# ask c.D", c.D)
		// fmt.Println("# ask c.C", c.C)
		// fmt.Println("# ask c.InvSqrtC", c.InvSqrtC)
		// fmt.Println("# here")
		for i := 0; i < c.N; i++ {
			if c.D.At(i, i) < 0 {
				c.D.Set(i, i, -1*math.Sqrt(-1*c.D.At(i, i)))
			} else {
				c.D.Set(i, i, math.Sqrt(c.D.At(i, i)))
			}

		}
		c.InvSqrtC.Inverse(c.D)
		c.InvSqrtC.Mul(c.B, c.InvSqrtC)
		c.InvSqrtC.Mul(c.InvSqrtC, c.B.T())
		// fmt.Println("Entrou no loop")
		// fmt.Println("ask c.B", c.B)
		// fmt.Println("ask c.D", c.D)
		// fmt.Println("ask c.C", c.C)
		// fmt.Println("ask c.InvSqrtC", c.InvSqrtC)
	}
	randn := make([]float64, c.N)
	for i := 0; i < c.N; i++ {
		randn[i] = c.D.At(i, i) * rand.NormFloat64()
	}
	tmp := mat64.NewDense(c.N, 1, randn)
	tmp.Mul(c.B, tmp)
	newL := make([]float64, c.N)
	for i := 0; i < c.N; i++ {
		newL[i] = c.Xmean[i] + c.Sigma*tmp.At(i, 0)
	}

	// fmt.Println("ask tmp", tmp)

	return newL
}

func (c *CMAES) tell() {
	c.CountEval += len(c.Population)
	xold := make([]float64, c.N)
	xnew := make([]float64, c.N)
	copy(xold, c.Xmean)

	sort.Sort(ByFit(c.Population))
	for i := 0; i < c.Mu; i++ {
		for j := 0; j < len(c.Xmean); j++ {
			xnew[j] += c.Population[i].X[j] * c.Weights[i]
		}
	}
	copy(c.Xmean, xnew)

	//  # Cumulation: update evolution paths
	y := make([]float64, c.N)
	for i := 0; i < len(y); i++ {
		y[i] = c.Xmean[i] - xold[i]
	}
	z := mat64.NewDense(c.N, 1, y)
	z.Mul(c.InvSqrtC, z)
	// fmt.Println("ESSE E O Z", z)

	// ps
	ctmp := math.Sqrt(c.Cs*(2.0-c.Cs)*c.Mueff) / c.Sigma
	for i := 0; i < len(c.Ps); i++ {
		c.Ps[i] -= c.Cs * c.Ps[i]
		c.Ps[i] += ctmp * z.At(i, 0)
	}
	// fmt.Println("Esse éo PS", c.Ps)

	// #hsig
	sumPs := 0.0
	for i := 0; i < len(c.Ps); i++ {
		sumPs += c.Ps[i] * c.Ps[i]
	}
	hsig := 0.0
	if (sumPs/(1.0-math.Pow((1.0-c.Cs), 2.0*float64(c.CountEval)/float64(c.Lambda))))/float64(c.N) < 2.0+4.0/(float64(c.N)+1.0) {
		hsig = 1.0
	}

	// pc
	ctmp = math.Sqrt(c.Cc*(2.0-c.Cc)*c.Mueff) / c.Sigma
	for i := 0; i < len(c.Pc); i++ {
		c.Pc[i] -= c.Cc * c.Pc[i]
		c.Pc[i] += ctmp * y[i] * hsig
	}
	// fmt.Println("Esse éo PC", c.Pc)

	// # Adapt covariance matrix C
	// c1a = self.c1 - (1-hsig**2) * self.c1 * self.cc * (2-self.cc)
	cmuij := 0.0
	c1a := c.C1 - (1.0-hsig)*c.C1*c.Cc*(2.0-c.Cc)
	for i := 0; i < c.N; i++ {
		for j := 0; j < c.N; j++ {
			cmuij = 0.0
			for k := 0; k < c.Mu; k++ {
				cmuij += c.Weights[k] * (c.Population[k].X[i] - xold[i]) * (c.Population[k].X[i] - xold[i])
			}
			cmuij /= c.Sigma * c.Sigma
			c.C.Set(i, j, c.C.At(i, j)+(-c1a-c.Cmu)*c.C.At(i, j)+c.C1*c.Pc[i]*c.Pc[j]+c.Cmu*cmuij)
		}
	}
	// fmt.Println("Esse é o novo C", c.C)
	// # Adapt step-size sigma with factor <= exp(0.6) \approx 1.82
	c.Sigma *= math.Exp(math.Min(0.6, c.Cs/c.Damps*(sumPs/float64(c.N)-1.0)/2.0))

}

type ByFit []Candidate

func (a ByFit) Len() int           { return len(a) }
func (a ByFit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFit) Less(i, j int) bool { return a[i].Fit < a[j].Fit }

func (*CMAES) Needs() struct {
	Gradient bool
	Hessian  bool
} {
	return struct {
		Gradient bool
		Hessian  bool
	}{false, false}
}
