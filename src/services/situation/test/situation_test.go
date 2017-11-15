package test

import (
	"log"
	"testing"
	sit "tryhard-platform/src/services/situation"
)

func TestSituation(t *testing.T) {
	s := sit.DefaultSituation()

	s.Generate(2, 1, 2, "+", "-")
	log.Println("Generated:", s.GetProblem(), "=", s.GetSolution())

	s.Generate(2, 2, 2, "+")
	log.Println("Generated:", s.GetProblem(), "=", s.GetSolution())

	// Generator is broken
	// -
	// -
	// -
	// go test -v situation_test.go
}
