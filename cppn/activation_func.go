/*


activation_func.go implementation of activation functions used in a network.

@licstart   The following is the entire license notice for
the Go code in this page.

Copyright (C) 2017 Jin Yeom

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

As additional permission under GNU GPL version 3 section 7, you
may distribute non-source (e.g., minimized or compacted) forms of
that code without the copy of the GNU GPL normally required by
section 4, provided you include this license notice and a URL
through which recipients can access the Corresponding Source.

@licend    The above is the entire license notice
for the Go code in this page.


*/

package cppn

import (
	"math"
	"math/rand"

	"github.com/jinyeom/neat"
)

var (
	ActivationSet = map[string]*neat.ActivationFunc{
		"identity": Identity(),
		"sigmoid":  Sigmoid(),
		"tanh":     Tanh(),
		"sin":      Sin(),
		"cos":      Cos(),
		"relu":     ReLU(),
		"log":      Log(),
		"exp":      Exp(),
		"abs":      Abs(),
		"square":   Square(),
		"cube":     Cube(),
		"gaussian": Gaussian(0.0, 1.0),
	}
)

// RandActivationFunc returns a random activation function from the
// ActivationSet.
func RandActivationFunc() *neat.ActivationFunc {
	afuncNames := make([]string, 0, len(ActivationSet))
	for name := range ActivationSet {
		afuncNames = append(afuncNames, name)
	}
	return ActivationSet[afuncNames[rand.Intn(len(afuncNames))]]
}

// Identity returns the identity function as an activation
// function. This function is only used for sensor nodes.
func Identity() *neat.ActivationFunc {
	return &neat.ActivationFunc{
		Name: "Identity",
		Fn: func(x float64) float64 {
			return x
		},
	}
}

// Sigmoid returns the sigmoid function as an activation function.
func Sigmoid() *neat.ActivationFunc {
	return &neat.ActivationFunc{
		Name: "Sigmoid",
		Fn: func(x float64) float64 {
			return 1.0 / (1.0 + math.Exp(-x))
		},
	}
}

// Tanh returns the hyperbolic tangent function as an activation function.
func Tanh() *neat.ActivationFunc {
	return &ActivationFunc{
		Name: "Tanh",
		Fn:   math.Tanh,
	}
}

// Sin returns the sin function as an activation function.
func Sin() *neat.ActivationFunc {
	return &ActivationFunc{
		Name: "Sine",
		Fn:   math.Sin,
	}
}

// Cos returns the cosine function as an activation function.
func Cos() *neat.ActivationFunc {
	return &ActivationFunc{
		Name: "Cosine",
		Fn:   math.Cos,
	}
}

// ReLU returns a rectifier linear unit as an activation function.
func ReLU() *neat.ActivationFunc {
	return &ActivationFunc{
		Name: "ReLU",
		Fn: func(x float64) float64 {
			return math.Max(x, 0.0)
		},
	}
}

// Log returns the log function as an activation function.
func Log() *neat.ActivationFunc {
	return &ActivationFunc{
		Name: "Log",
		Fn:   math.Log,
	}
}

// Exp returns the exponential function as an activation function.
func Exp() *ActivationFunc {
	return &ActivationFunc{
		Name: "Exp",
		Fn:   math.Exp,
	}
}

// Abs returns the absolute value function as an activation function.
func Abs() *ActivationFunc {
	return &ActivationFunc{
		Name: "Abs",
		Fn:   math.Abs,
	}
}

// Square returns the square function as an activation function.
func Square() *ActivationFunc {
	return &ActivationFunc{
		Name: "Square",
		Fn: func(x float64) float64 {
			return x * x
		},
	}
}

// Cube returns the cube function as an activation function.
func Cube() *ActivationFunc {
	return &ActivationFunc{
		Name: "Cube",
		Fn: func(x float64) float64 {
			return x * x * x
		},
	}
}

// Gaussian returns the Gaussian function as an activation function, given a
// mean and a standard deviation.
func Gaussian(mean, stdev float64) *ActivationFunc {
	return &ActivationFunc{
		Name: "Gaussian",
		Fn: func(x float64) float64 {
			return 1.0 / (stdev * math.Sqrt(2*math.Pi)) *
				math.Exp(math.Pow((x-mean)/stdev, 2.0)/-2.0)
		},
	}
}
