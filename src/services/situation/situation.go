package situation

import (
	"fmt"
	reg "regexp"
	str "strconv"
)

// static const boost::regex operandRegex("^(\\-?)(\\d+)?(\\d|\\.)?(\\d)+");
// static const boost::regex operatorRegex("(\\+|\\-|\\(|\\)|/|\\*|x|\\^)");

type Situation interface {
	Parse(equation string)
	Generate(operandCnt, operandLengthMin, operandLengthMax int, operations string)
	SetCategories(categories []string)
	GetCategories() []string
	GetSource() string
	GetEquation() string
	GetSolution() float64
}

type situation struct {
	root          Node
	categories    []string
	source        string
	equation      string
	solution      float64
	operandRegEx  *reg.Regexp
	operatorRegEx *reg.Regexp
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
			panic(fmt.Sprintf("failed to convert %v to float \n", value))
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

// TODO
func (s *situation) buildTree(posNode, operandNode, operatorNode Node) Node {
	// build from source
	//  if(!position){
	// 	if(!operandNode && !operatorNode){
	// 		return new MathNode(NodeType::Operand, "0");
	// 	}
	// 	else if(operandNode && !operatorNode){
	// 		position = operandNode;
	// 	}
	// 	else{
	// 		operatorNode->setLeft(operandNode);
	// 		operandNode->setParent(operatorNode);
	// 		MathNode* placeholder = nullptr;
	// 		if(operatorNode->getText() == "+" || operatorNode->getText() == "-"){
	// 			placeholder = new MathNode(operatorNode, NodeType::Operand, "0");
	// 		}
	// 		else{
	// 			placeholder = new MathNode(operatorNode, NodeType::Operand, "1");
	// 		}
	// 		placeholder->setPlaceholder(true);
	// 		operatorNode->setRight(placeholder);
	// 		position = operatorNode;
	// 	}
	// }
	// else{
	// 	if(position->getRight()->getPlaceholder() && !operatorNode){
	// 		position->setRight(operandNode);
	// 		operandNode->setParent(position);
	// 		position->setPlaceholder(false);
	// 	}
	// 	else if(position->getRight()->getPlaceholder() && operatorNode){
	// 		MathNode* placeholder = nullptr;
	// 		if(operatorNode->getText() == "+" || operatorNode->getText() == "-"){
	// 			placeholder = new MathNode(operatorNode, NodeType::Operand, "0");
	// 			placeholder->setPlaceholder(true);
	// 			position->setRight(operandNode);
	// 			operandNode->setParent(position);
	// 			operatorNode->setRight(placeholder);
	// 			operatorNode->setLeft(position);
	// 			position->setParent(operatorNode);
	// 			return operatorNode;
	// 		}
	// 		else{
	// 			placeholder = new MathNode(operatorNode, NodeType::Operand, "1");
	// 		}
	// 	}
	// 	else{

	// 	}
	// }
	// return position;
}

// TODO
func (s *situation) Parse(equation string) {
	// this->equation = equation;

	// string operand = "0";
	// for(int i=0; i<(int)equation.length(); i++){
	// 	char charIndex = equation[i];

	// 	if(charIndex == ' '){
	// 		continue;
	// 	}

	// 	string tempOperand = operand;
	// 	tempOperand.push_back(charIndex);

	// 	if(operandCheck(tempOperand)){
	// 		operand.push_back(charIndex);
	// 		if(i == (int)equation.length()-1){
	// 			MathNode* operandNode = new MathNode(NodeType::Operand, operand);
	// 			this->root = buildTree(this->root, operandNode, nullptr);
	// 		}
	// 		else
	// 			continue;
	// 	}
	// 	else if(operatorCheck(charIndex)){
	// 		if(i == 0){
	// 			// leading character is operator
	// 			if(charIndex != '-'){
	// 				// ignore all leading operators not satisfied
	// 				continue;
	// 			}
	// 		}
	// 		MathNode* operandNode = new MathNode(NodeType::Operand, operand);
	// 		MathNode* operatorNode = new MathNode(NodeType::Operator, charIndex);
	// 		this->root = buildTree(this->root, operandNode, operatorNode);
	// 		operand = "0";
	// 	}
	// }
	// this->equation = this->buildEquation(this->root);
	// this->solution = this->solve(this->root);
}

// TODO
func (s *situation) Generate(operandCnt, operandMinLength, operandMaxLength int, operations string) {
	// if(this->root){
	// 	delete this->root;
	// 	this->root = nullptr;
	// }

	// for(int i = 0; i < operandCount; i++){
	// 	if (i == operandCount-1){
	// 		string operand_str = Generator::operand(operandLengthMin, operandLengthMax);
	// 		MathNode* operandNode = new MathNode(NodeType::Operand, operand_str);
	// 		this->root = this->buildTree(this->root, operandNode, nullptr);
	// 	}
	// 	else{
	// 		int op_index = rand() % operations.length();
	// 		char op = operations[op_index];
	// 		string operand_str = Generator::operand(operandLengthMin, operandLengthMax);
	// 		MathNode* operandNode = new MathNode(NodeType::Operand, operand_str);
	// 		MathNode* operatorNode = new MathNode(NodeType::Operator, op);
	// 		this->root = this->buildTree(this->root, operandNode, operatorNode);
	// 	}
	// }
	// this->equation = this->buildEquation(this->root);
	// this->solution = this->solve(this->root);
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

func (s *situation) GetEquation() string {
	return s.equation
}

func (s *situation) GetSolution() float64 {
	return s.solution
}

func DefaultSituation() Situation {
	return &situation{
		operandRegEx:  reg.MustCompile(`^(\\-?)(\\d+)?(\\d|\\.)?(\\d)+`),
		operatorRegEx: reg.MustCompile(`(\\+|\\-|\\(|\\)|/|\\*|x|\\^)`),
	}
}
