package bracelet

type specificity [3]int

func (s specificity) Less(other specificity) bool {
	for i := 0; i < 3; i++ {
		if s[i] != other[i] {
			return s[i] < other[i]
		}
	}
	return false
}

type MatchedRule struct {
	Specificity specificity
	Rule        Rule
}

func matchRule(node *Node, rule Rule) *MatchedRule {
	for _, selector := range rule.Selectors {
		if selector.Matches(node) {
			return &MatchedRule{
				Specificity: selector.Specificity(),
				Rule:        rule,
			}
		}
	}
	return nil
}

func matchingRules(node *Node, stylesheet []Rule) []MatchedRule {
	var matches []MatchedRule
	for i := range stylesheet {
		if match := matchRule(node, stylesheet[i]); match != nil {
			matches = append(matches, *match)
		}
	}
	return matches
}
