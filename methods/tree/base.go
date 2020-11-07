package tree

type BTNode struct {
	Data  uint32
	Bf    int8
	Left  *BTNode
	Right *BTNode
}

func (n *BTNode) Init(data uint32) {
	n.Data = data
	n.Bf = 0
}
