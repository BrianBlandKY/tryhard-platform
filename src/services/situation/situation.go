package situation

import (
	"fmt"
	reg "regexp"
	str "strconv"
)

type Situation interface {
	Parse(equation string)
	Generate(operandCnt, operandLengthMin, operandLengthMax int, operations ...string)
	SetCategories(categories []string)
	GetCategories() []string
	GetSource() string
	GetProblem() string
	GetSolution() float64
}

type situation struct {
	root          Node
	categories    []string
	source        string
	problem       string
	solution      float64
	operandRegEx  *reg.Regexp
	operatorRegEx *reg.Regexp
	generator     Generator
}

func (s *situation) operandCheck(value string) bool {
	return s.operandRegEx.MatchString(value)
}

func (s *situation) operatorCheck(value string) bool {
	return s.operatorRegEx.MatchString(value)
}

func (s *situation) solve(pos Node) float64 {
	valueLeft := float64(0)
	valueRight := float64(0)
	value := float64(0)

	if pos.GetNodeType() == OPERAND {
		value, err := str.ParseFloat(pos.GetText(), 32)
		if err != nil {
			panic(fmt.Sprintf("err: %v", err))
		}
		return value
	}

	if pos.GetLeft() != nil {
		valueLeft = s.solve(pos.GetLeft())
	}

	if pos.GetRight() != nil {
		valueRight = s.solve(pos.GetRight())
	}

	if pos.GetText() == "+" {
		value = Add(valueLeft, valueRight)
	}

	if pos.GetText() == "-" {
		value = Subtract(valueLeft, valueRight)
	}
	return value
}

func (s *situation) build(pos Node) string {
	if pos == nil {
		return ""
	}

	if pos.GetNodeType() == OPERAND {
		return pos.GetText()
	}

	return fmt.Sprintf("%v %v %v",
		s.build(pos.GetLeft()),
		pos.GetText(),
		s.build(pos.GetRight()))
}

func (s *situation) buildTree(posNode, operandNode, operatorNode Node) Node {
	if posNode == nil {
		if operandNode == nil && operatorNode == nil {
			// What is this for?
			return BuildNode(nil, OPERAND, "0")
		} else if operandNode != nil && operatorNode == nil {
			posNode = operandNode
		} else {
			operatorNode.SetLeft(operandNode)
			operandNode.SetParent(operatorNode)
			var placeholder Node
			if operatorNode.GetText() == "+" || operatorNode.GetText() == "-" {
				// Why is text "0"?
				placeholder = BuildNode(operatorNode, OPERAND, "0")
			} else {
				// Why is text "1"?
				placeholder = BuildNode(operatorNode, OPERAND, "1")
			}
			placeholder.SetPlaceholder(true)
			operatorNode.SetRight(placeholder)
			posNode = operatorNode
		}
	} else {
		if posNode.GetRight().GetPlaceholder() && operatorNode == nil {
			posNode.SetRight(operandNode)
			operandNode.SetParent(posNode)
			posNode.SetPlaceholder(false)
		} else if posNode.GetRight().GetPlaceholder() && operatorNode != nil {
			var placeholder Node
			if operatorNode.GetText() == "+" || operatorNode.GetText() == "-" {
				placeholder = BuildNode(operatorNode, OPERAND, "0")
				placeholder.SetPlaceholder(true)
				posNode.SetRight(operatorNode)
				operandNode.SetParent(posNode)
				operatorNode.SetRight(placeholder)
				operatorNode.SetLeft(posNode)
				posNode.SetParent(operatorNode)
				return operatorNode
			}
			// What do we do with this guy?
			placeholder = BuildNode(operatorNode, OPERAND, "1")
		} else {
			// more?
		}
	}
	return posNode
}

func (s *situation) Parse(problem string) {
	s.problem = problem
	operand := "0"
	var charIdx string
	for i := 0; i < len(s.problem); i++ {
		charIdx = string(s.problem[i])

		if charIdx == " " {
			continue
		}

		tempOperand := operand
		tempOperand += charIdx

		if s.operandCheck(tempOperand) {
			tempOperand += charIdx
			if i == len(s.problem)-1 {
				operandNode := BuildNode(nil, OPERAND, operand)
				s.root = s.buildTree(s.root, operandNode, nil)
			} else {
				continue
			}
		} else if s.operatorCheck(charIdx) {
			if i == 0 {
				// the leading character is operator
				if charIdx != "-" {
					// ignore all other leading operators
					continue
				}
			}
			operandNode := BuildNode(nil, OPERAND, operand)
			operatorNode := BuildNode(nil, OPERATOR, charIdx)
			s.root = s.buildTree(s.root, operandNode, operatorNode)
			operand = "0"
		}
		s.problem = s.build(s.root)
		s.solution = s.solve(s.root)
	}
}

func (s *situation) Generate(operandCnt, operandMinLength, operandMaxLength int, operations ...string) {
	if s.root != nil {
		s.root = nil
	}

	for i := 0; i < operandCnt; i++ {
		if i == operandCnt-1 {
			// ?
			operandStr := s.generator.Operand(operandMinLength, operandMaxLength)
			operandNode := BuildNode(nil, OPERAND, operandStr)
			s.root = s.buildTree(s.root, operandNode, nil)
		} else {
			operator := s.generator.Operator(operations...)
			operand := s.generator.Operand(operandMinLength, operandMaxLength)
			operandNode := BuildNode(nil, OPERAND, operand)
			operatorNode := BuildNode(nil, OPERATOR, operator)
			s.root = s.buildTree(s.root, operandNode, operatorNode)
		}
		s.problem = s.build(s.root)
		s.solution = s.solve(s.root)
	}
}

func (s *situation) GetCategories() []string {
	return s.categories
}

func (s *situation) SetCategories(categories []string) {
	s.categories = categories
}

func (s *situation) GetSource() string {
	return s.source
}

func (s *situation) GetProblem() string {
	return s.problem
}

func (s *situation) GetSolution() float64 {
	return s.solution
}

func DefaultSituation() Situation {
	return &situation{
		operandRegEx:  reg.MustCompile(`^(\\-?)(\\d+)?(\\d|\\.)?(\\d)+`),
		operatorRegEx: reg.MustCompile(`(\\+|\\-|\\(|\\)|/|\\*|x|\\^)`),
		generator:     DefaultGenerator(),
	}
}
