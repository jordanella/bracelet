package bracelet

import "fmt"

type notSelector struct {
	Base     Selector
	Negation Selector
}

func (s *notSelector) Matches(node *Node) bool {
	return s.Base.Matches(node) && !s.Negation.Matches(node)
}

func (s *notSelector) Specificity() specificity {
	baseSpec := s.Base.Specificity()
	negSpec := s.Negation.Specificity()
	return specificity{
		baseSpec[0] + negSpec[0],
		baseSpec[1] + negSpec[1],
		baseSpec[2] + negSpec[2],
	}
}

func (s *notSelector) String() string {
	return fmt.Sprintf("%s:not(%s)", s.Base.String(), s.Negation.String())
}
