package situation

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Generator interface {
	Operand(operandMin, operandMax int) string
	Operator(operators ...string) string
}

type generator struct {
	rng *rand.Rand
}

/*
	Operand
	Generate an Operand (number) between the min and max settings.
*/
func (g *generator) Operand(operandMin, operandMax int) string {
	result := ""
	length := 0

	// get the length of generated number
	// TODO: Improve the min + max algorithm
	if operandMax-operandMin == 0 {
		length = g.rng.Intn(operandMax) + 1
	} else {
		length = operandMin + g.rng.Intn(operandMax-operandMin)
	}

	log.Println("Length:", length)
	// generate random number
	for i := 0; i < length; i++ {
		n := g.rng.Intn(9)
		if length > 1 && i == 0 {
			n = g.rng.Intn(8) + 1 // prevent leading digit from being zero
		}
		result = fmt.Sprintf("%v%v", result, n)
	}
	return result
}

func (g *generator) Operator(operators ...string) string {
	idx := g.rng.Intn(len(operators))
	return operators[idx]
}

func DefaultGenerator() Generator {
	return &generator{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
