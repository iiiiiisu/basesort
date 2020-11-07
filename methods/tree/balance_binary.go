package tree

type Rotation int

const (
	LeftRotation  Rotation = -1
	NoRotation    Rotation = 0
	RightRotation Rotation = 1
)

type BalanceBinary struct {
	Data   []uint32
	Length int32
	Root   *BTNode
}

func (s *BalanceBinary) SetData(data []uint32) {
	s.Data = data
	s.SetLength()
}

func (s *BalanceBinary) SetLength() {
	s.Length = int32(len(s.Data))
}

func (s *BalanceBinary) InsertNode(data uint32) {
	node := s.Root
	path := []*BTNode{}
	for {
		if data >= node.Data && node.Right != nil {
			path = append(path, node)
			node = node.Right
			continue
		}
		if data < node.Data && node.Left != nil {
			path = append(path, node)
			node = node.Left
			continue
		}
		if data >= node.Data {
			node.Right = &BTNode{}
			node.Right.Init(data)
			s.GetDepth(node)
		} else {
			node.Left = &BTNode{}
			node.Left.Init(data)
			s.GetDepth(node)
		}
		break
	}
	//for node.Left != nil && node.Right != nil {
	//	path = append(path, node)
	//	if data < node.Data {
	//		node = node.Left
	//	} else {
	//		node = node.Right
	//	}
	//}
	//if data < node.Data {
	//	if node.Left != nil {
	//		node.Right = &BTNode{}
	//		node.Right.Init(node.Data)
	//		node.Data = data
	//	} else {
	//		node.Left = &BTNode{}
	//		node.Left.Init(data)
	//		node.Bf = 1
	//	}
	//} else {
	//	if node.Right != nil {
	//		node.Left = &BTNode{}
	//		node.Left.Init(node.Data)
	//		node.Data = data
	//	} else {
	//		node.Right = &BTNode{}
	//		node.Right.Init(data)
	//		node.Bf = 1
	//	}
	//}
	length := len(path)
	for i := length - 1; i >= 0; i-- {
		node = path[i]
		r := s.GetRotation(node)
		if r == NoRotation {
			s.GetDepth(node)
			continue
		}
		if r == LeftRotation {
			s.LeftRotate(node)
		} else {
			s.RightRotate(node)
		}
		s.GetDepth(node)
	}
}

func (s *BalanceBinary) GetRotation(node *BTNode) Rotation {
	if node.Right == nil {
		if node.Left == nil {
			return NoRotation
		} else if node.Left.Bf == 0 {
			return NoRotation
		} else {
			return RightRotation
		}
	}
	if node.Left == nil {
		if node.Right.Bf == 0 {
			return NoRotation
		} else {
			return LeftRotation
		}
	}
	d := node.Right.Bf - node.Left.Bf
	if d > 1 {
		return LeftRotation
	} else if d < -1 {
		return RightRotation
	} else {
		return NoRotation
	}
}

func (s *BalanceBinary) LeftRotate(node *BTNode) {
	rNode := node.Right
	if rNode.Right == nil || rNode.Left != nil && rNode.Left.Bf > rNode.Right.Bf {
		temp := &BTNode{}
		temp.Init(rNode.Data)
		temp.Right = rNode.Right
		temp.Left = rNode.Left.Right
		s.GetDepth(temp)
		rNode.Right = temp
		rNode.Data = rNode.Left.Data
		rNode.Left = rNode.Left.Left
		s.GetDepth(rNode)
	}
	temp := &BTNode{}
	temp.Init(node.Data)
	temp.Left = node.Left
	temp.Right = node.Right.Left
	s.GetDepth(temp)
	node.Left = temp
	node.Data = node.Right.Data
	node.Right = node.Right.Right
	s.GetDepth(node)
}

func (s *BalanceBinary) RightRotate(node *BTNode) {
	lNode := node.Left
	if lNode.Left == nil || lNode.Right != nil && lNode.Right.Bf > lNode.Left.Bf {
		temp := &BTNode{}
		temp.Init(lNode.Data)
		temp.Left = lNode.Left
		temp.Right = lNode.Right.Left
		s.GetDepth(temp)
		lNode.Left = temp
		lNode.Data = lNode.Right.Data
		lNode.Right = lNode.Right.Right
		s.GetDepth(lNode)
	}
	temp := &BTNode{}
	temp.Init(node.Data)
	temp.Right = node.Right
	temp.Left = node.Left.Right
	s.GetDepth(temp)
	node.Right = temp
	node.Data = node.Left.Data
	node.Left = node.Left.Left
	s.GetDepth(node)
}

func (s *BalanceBinary) GetDepth(node *BTNode) {
	var left, right int8
	if node.Left == nil {
		left = -1
	} else {
		left = node.Left.Bf
	}
	if node.Right == nil {
		right = -1
	} else {
		right = node.Right.Bf
	}
	if left > right {
		node.Bf = left + 1
	} else {
		node.Bf = right + 1
	}
}

func (s *BalanceBinary) Sort() {
	var i int32
	s.Root = &BTNode{}
	s.Root.Init(s.Data[0])
	for i = 1; i < s.Length; i++ {
		s.InsertNode(s.Data[i])
	}
}

func (s *BalanceBinary) CheckResult() bool {
	var i int32
	s.InorderTraversal()
	for i = 1; i < s.Length; i++ {
		if s.Data[i] < s.Data[i-1] {
			return false
		}
	}
	return true
}

func (s *BalanceBinary) InorderTraversal() {
	node := s.Root
	length := 0
	pos := 0
	stack := make([]*BTNode, s.Root.Bf+1)
	for node != nil || length != 0 {
		for node != nil {
			stack[length] = node
			length++
			node = node.Left
			continue
		}
		if node == nil {
			length--
			node = stack[length]
		}
		s.Data[pos] = node.Data
		pos++
		node = node.Right
	}
}
