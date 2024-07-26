package bracelet

import (
	"errors"
	"fmt"
	"strings"
)

type Selector interface {
	Specificity() specificity
	Matches(node *Node) bool
	String() string
}

func parseSelector(input string) (Selector, error) {
	input = strings.TrimSpace(input)
	tokens := tokenizeSelector(input)
	return parseTokens(tokens)
}

func tokenizeSelector(input string) []string {
	input = strings.ReplaceAll(input, ">", " > ")
	input = strings.ReplaceAll(input, "+", " + ")
	input = strings.ReplaceAll(input, "~", " ~ ")
	input = strings.ReplaceAll(input, "[", " [ ")
	input = strings.ReplaceAll(input, "]", " ] ")
	return strings.Fields(input)
}

func parseTokens(tokens []string) (Selector, error) {
	if len(tokens) == 0 {
		return nil, errors.New("empty selector")
	}

	var currentSelector Selector

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token {
		case ">":
			if currentSelector == nil {
				return nil, errors.New("child selector cannot be at the start")
			}
			i++
			if i >= len(tokens) {
				return nil, errors.New("missing child selector")
			}
			parsedChildSelector, err := parseSimpleSelector(tokens[i])
			if err != nil {
				return nil, err
			}
			currentSelector = &childSelector{Parent: currentSelector, Child: parsedChildSelector}
		case "+":
			if currentSelector == nil {
				return nil, errors.New("adjacent sibling selector cannot be at the start")
			}
			i++
			if i >= len(tokens) {
				return nil, errors.New("missing adjacent sibling selector")
			}
			siblingSelector, err := parseSimpleSelector(tokens[i])
			if err != nil {
				return nil, err
			}
			currentSelector = &adjacentSiblingSelector{First: currentSelector, Second: siblingSelector}
		case "[":
			attrSelector, newIndex, err := parseAttributeSelector(tokens[i:])
			if err != nil {
				return nil, err
			}
			i = newIndex
			if currentSelector == nil {
				currentSelector = attrSelector
			} else {
				currentSelector = &descendantSelector{Ancestor: currentSelector, Descendant: attrSelector}
			}
		default:
			simpleSelector, err := parseSimpleSelector(token)
			if err != nil {
				return nil, err
			}
			if currentSelector == nil {
				currentSelector = simpleSelector
			} else {
				currentSelector = &descendantSelector{Ancestor: currentSelector, Descendant: simpleSelector}
			}
		}
	}

	return currentSelector, nil
}

func parsePseudoSelector(baseSelector Selector, pseudo string) (Selector, error) {
	switch {
	case pseudo == "first-child":
		return &firstChildSelector{baseSelector}, nil
	case pseudo == "last-child":
		return &lastChildSelector{baseSelector}, nil
	case strings.HasPrefix(pseudo, "nth-child("):
		n, err := parseNthChild(pseudo)
		if err != nil {
			return nil, err
		}
		return &nthChildSelector{baseSelector, n}, nil
	case strings.HasPrefix(pseudo, "not("):
		notContent := strings.TrimSuffix(strings.TrimPrefix(pseudo, "not("), ")")
		parsedNotSelector, err := parseSelector(notContent)
		if err != nil {
			return nil, fmt.Errorf("invalid :not selector: %s", err)
		}
		return &notSelector{Base: baseSelector, Negation: parsedNotSelector}, nil
	default:
		return nil, fmt.Errorf("unsupported pseudo-selector: %s", pseudo)
	}
}
