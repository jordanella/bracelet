package bracelet

type lastChildSelector struct {
	Selector
}

func (s *lastChildSelector) Matches(node *Node) bool {
	if !s.Selector.Matches(node) {
		return false
	}
	parent := (*node).GetParent()
	if parent == nil {
		return false
	}
	children := (*parent).GetChildren()
	return children[len(children)-1] == node
}

func (s *lastChildSelector) Specificity() specificity {
	spec := s.Selector.Specificity()
	return specificity{spec[0], spec[1] + 1, spec[2]}
}

func (s *lastChildSelector) String() string {
	return s.Selector.String() + ":last-child"
}
