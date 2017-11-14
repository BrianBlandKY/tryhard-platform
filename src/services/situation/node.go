package situation

type NodeType int

const (
	OPERATOR = NodeType(iota)
	OPERAND
)

type Node interface {
	GetNodeType() NodeType
	GetText() string
	GetParent() Node
	GetLeft() Node
	GetRight() Node
	SetParent(Node)
	SetRight(Node)
	SetLeft(Node)
	SetPlaceholder(bool)
	GetPlaceholder() bool
}

type node struct {
	parent        Node
	left          Node
	right         Node
	nodeType      NodeType
	text          string
	isPlaceholder bool
}

func (n *node) GetNodeType() NodeType {
	return n.nodeType
}

func (n *node) GetText() string {
	return n.text
}

func (n *node) GetParent() Node {
	return n.parent
}

func (n *node) GetLeft() Node {
	return n.left
}

func (n *node) GetRight() Node {
	return n.right
}

func (n *node) SetParent(parent Node) {
	n.parent = parent
}

func (n *node) SetLeft(left Node) {
	n.left = left
}

func (n *node) SetRight(right Node) {
	n.right = right
}

func (n *node) SetPlaceholder(ph bool) {
	n.isPlaceholder = ph
}

func (n *node) GetPlaceholder() bool {
	return n.isPlaceholder
}

func BuildNode() Node {
	return nil
}
