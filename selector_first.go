package bracelet

type firstChildSelector struct {
	Selector
}

func (s *firstChildSelector) Matches(node *Node) bool {
	if !s.Selector.Matches(node) {
		return false
	}
	parent := (*node).GetParent()
	if parent == nil {
		return false
	}
	return (*parent).GetChildren()[0] == node
}

func (s *firstChildSelector) Specificity() specificity {
	spec := s.Selector.Specificity()
	return specificity{spec[0], spec[1] + 1, spec[2]}
}

func (s *firstChildSelector) String() string {
	return s.Selector.String() + ":first-child"
}
