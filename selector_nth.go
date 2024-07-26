package bracelet

import (
	"fmt"
	"strconv"
	"strings"
)

type nthChildSelector struct {
	Selector
	N int
}

func (s *nthChildSelector) Matches(node *Node) bool {
	if !s.Selector.Matches(node) {
		return false
	}
	parent := (*node).GetParent()
	if parent == nil {
		return false
	}
	children := (*parent).GetChildren()
	for i, child := range children {
		if child == node {
			return i+1 == s.N
		}
	}
	return false
}

func (s *nthChildSelector) Specificity() specificity {
	spec := s.Selector.Specificity()
	return specificity{spec[0], spec[1] + 1, spec[2]}
}

func (s *nthChildSelector) String() string {
	return fmt.Sprintf("%s:nth-child(%d)", s.Selector.String(), s.N)
}

func parseNthChild(pseudo string) (int, error) {
	n, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(pseudo, "nth-child("), ")"))
	if err != nil {
		return 0, fmt.Errorf("invalid nth-child value: %s", pseudo)
	}
	return n, nil
}
