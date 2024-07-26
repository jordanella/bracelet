package bracelet

type childSelector struct {
	Parent Selector
	Child  Selector
}

func (s *childSelector) String() string {
	return s.Parent.String() + " > " + s.Child.String()
}

func (s *childSelector) Specificity() specificity {
	first := s.Parent.Specificity()
	second := s.Child.Specificity()
	return specificity{first[0] + second[0], first[1] + second[1], first[2] + second[2]}
}

func (s *childSelector) Matches(node *Node) bool {

	childMatches := s.Child.Matches(node)

	if !childMatches {
		return false
	}

	parent := (*node).GetParent()
	if parent == nil {
		return false
	}

	parentMatches := s.Parent.Matches(parent)

	return parentMatches
}
