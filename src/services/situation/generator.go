package situation

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Generator interface {
	Operand(operandMin, operandMax int) string
	// OperandWithLeadingZero?
}

type generator struct {
}

/*
	Operand
	Generate an Operand (number) between the min and max settings.
*/
func (g *generator) Operand(operatorMin, operatorMax int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := ""

	// get the length of generated number
	length := r.Intn(operatorMax-operatorMin) + operatorMin
	log.Printf("Operand Length: %v \r\n", length)

	// generate random number
	for i := 0; i < length; i++ {
		n := r.Intn(9)
		if length > 1 && i == 0 {
			n = r.Intn(8) + 1 // prevent leading digit from being zero
		}
		log.Printf("Single Op: %v \r\n", n)
		result = fmt.Sprintf("%v%v", result, n)
		// result += ((char)(((int)'0')+(rand() % 9)));
	}
	log.Printf("Result: %v \r\n", result)
	return result
}

func DefaultGenerator() Generator {
	return &generator{}
}
