package bracelet

type adjacentSiblingSelector struct {
	First  Selector
	Second Selector
}

func (s *adjacentSiblingSelector) String() string {
	return s.First.String() + " + " + s.Second.String()
}

func (s *adjacentSiblingSelector) Specificity() specificity {
	first := s.First.Specificity()
	second := s.Second.Specificity()
	return specificity{first[0] + second[0], first[1] + second[1], first[2] + second[2]}
}

func (s *adjacentSiblingSelector) Matches(node *Node) bool {
	if !s.Second.Matches(node) {
		return false
	}
	parent := (*node).GetParent()
	if parent == nil {
		return false
	}
	siblings := (*parent).GetChildren()
	for i, sibling := range siblings {
		if sibling == node && i > 0 {
			return s.First.Matches(siblings[i-1])
		}
	}
	return false
}
