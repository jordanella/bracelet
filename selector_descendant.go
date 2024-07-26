package bracelet

type descendantSelector struct {
	Ancestor   Selector
	Descendant Selector
}

func (s *descendantSelector) String() string {
	return s.Ancestor.String() + " " + s.Descendant.String()
}

func (s *descendantSelector) Specificity() specificity {
	first := s.Ancestor.Specificity()
	second := s.Descendant.Specificity()
	return specificity{first[0] + second[0], first[1] + second[1], first[2] + second[2]}
}

func (s *descendantSelector) Matches(node *Node) bool {
	if !s.Descendant.Matches(node) {
		return false
	}
	parent := (*node).GetParent()
	for parent != nil {
		if s.Ancestor.Matches(parent) {
			return true
		}
		parent = (*parent).GetParent()
	}
	return false
}
