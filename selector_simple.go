package bracelet

import "strings"

type simpleSelector struct {
	Tag                string
	ID                 string
	Classes            []string
	AttributeSelectors []attributeSelector
}

func (s *simpleSelector) String() string {
	result := s.Tag
	if s.ID != "" {
		result += "#" + s.ID
	}
	for _, class := range s.Classes {
		result += "." + class
	}
	return result
}

func (s *simpleSelector) Specificity() specificity {
	a := 0                                          // ID
	b := len(s.Classes) + len(s.AttributeSelectors) // Class and attribute selectors
	c := 0                                          // Type
	if s.ID != "" {
		a = 1
	}
	if s.Tag != "" {
		c = 1
	}
	return specificity{a, b, c}
}

func (s *simpleSelector) Matches(node *Node) bool {
	if s.Tag != "" && s.Tag != (*node).GetTag() {
		return false
	}
	if s.ID != "" && s.ID != (*node).GetID() {
		return false
	}

	contains := func(slice []string, item string) bool {
		for _, a := range slice {
			if a == item {
				return true
			}
		}
		return false
	}

	nodeClasses := (*node).GetClasses()
	for _, class := range s.Classes {
		if !contains(nodeClasses, class) {
			return false
		}
	}

	for _, attrSelector := range s.AttributeSelectors {
		if !attrSelector.Matches(node) {
			return false
		}
	}

	return true
}

func parseSimpleSelector(token string) (Selector, error) {
	selector := &simpleSelector{}
	parts := strings.Split(token, ":")

	if parts[0] != "" {
		mainParts := strings.Split(parts[0], ".")
		if len(mainParts) > 0 && mainParts[0] != "" {
			if mainParts[0][0] == '#' {
				selector.ID = mainParts[0][1:]
			} else {
				selector.Tag = mainParts[0]
			}
		}
		if len(mainParts) > 1 {
			selector.Classes = mainParts[1:]
		}
	}

	if len(parts) > 1 {
		return parsePseudoSelector(selector, parts[1])
	}

	return selector, nil
}
