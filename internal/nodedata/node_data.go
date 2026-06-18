package nodedata

type NodeData struct {
	ch    rune
	value int
	code  string
	left  *NodeData
	right *NodeData
}

func New(ch rune, value int, left *NodeData, right *NodeData) *NodeData {
	return &NodeData{ch: ch, value: value, left: left, right: right}
}

func (nd *NodeData) GetCh() rune {
	return nd.ch
}

func (nd *NodeData) GetValue() int {
	return nd.value
}

func (nd *NodeData) GetCode() string {
	return nd.code
}

func (nd *NodeData) GetLeft() *NodeData {
	return nd.left
}

func (nd *NodeData) GetRight() *NodeData {
	return nd.right
}

func (nd *NodeData) CompareTo(other *NodeData) int {
	if nd.value < other.value {
		return 1
	}
	if nd.value > other.value {
		return -1
	}
	return 0
}

func (nd *NodeData) SetCode() {
	nd.setCode("")
}

func (nd *NodeData) setCode(code string) {
	if nd.IsLeaf() {
		nd.code = code
		return
	}
	nd.left.setCode(code + "0")
	nd.right.setCode(code + "1")
}

func (nd *NodeData) IsLeaf() bool {
	return nd.left == nil && nd.right == nil
}
