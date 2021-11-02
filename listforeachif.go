package piscine

func IsPositive_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return true
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	novohead := l.Head
	for novohead != nil {
		if cond(novohead) == true {
			f(novohead)
		}
		novohead = novohead.Next
	}
}
