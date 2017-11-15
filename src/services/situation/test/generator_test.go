package test

import (
	"strconv"
	"testing"
	sit "tryhard-platform/src/services/situation"
)

func TestOperandGenerator(t *testing.T) {
	gen := sit.DefaultGenerator()

	for i := 0; i < 100; i++ {
		valueStr := gen.Operand(1, 2)
		value, _ := strconv.Atoi(valueStr)

		if value > 99 || value < 0 {
			t.Fatalf("Invalid generator value %v range (0, 99)", value)
		}
	}

	for i := 0; i < 100; i++ {
		valueStr := gen.Operand(2, 4)
		value, _ := strconv.Atoi(valueStr)

		if value > 9999 || value < 10 {
			t.Fatalf("Invalid generator value %v range (99, 9999)", value)
		}
	}
}

func TestOperatorGenerator(t *testing.T) {
}
