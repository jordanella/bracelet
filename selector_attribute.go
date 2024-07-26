package bracelet

import (
	"errors"
	"fmt"
	"strings"
)

type attributeSelector struct {
	Name      string
	Value     string
	Operation attributeSelectorOperationType
}
type attributeSelectorOperationType int

var attributeOperation = struct {
	Exists, Exact, Contains, StartsWith, EndsWith, HyphenSeparated attributeSelectorOperationType
}{0, 1, 2, 3, 4, 5}

func (s *attributeSelector) String() string {
	switch s.Operation {
	case attributeOperation.Exists:
		return "[" + s.Name + "]"
	case attributeOperation.Exact:
		return "[" + s.Name + "=\"" + s.Value + "\"]"
	case attributeOperation.Contains:
		return "[" + s.Name + "*=\"" + s.Value + "\"]"
	case attributeOperation.StartsWith:
		return "[" + s.Name + "^=\"" + s.Value + "\"]"
	case attributeOperation.EndsWith:
		return "[" + s.Name + "$=\"" + s.Value + "\"]"
	case attributeOperation.HyphenSeparated:
		return "[" + s.Name + "|=\"" + s.Value + "\"]"
	default:
		return "[" + s.Name + "?=\"" + s.Value + "\"]"
	}
}

func (s *attributeSelector) Specificity() specificity {
	return specificity{0, 1, 0}
}

func (s *attributeSelector) Matches(node *Node) bool {
	value, exists := (*node).GetAttributes()[s.Name]
	if !exists {
		return false
	}

	switch s.Operation {
	case attributeOperation.Exists:
		return true
	case attributeOperation.Exact:
		return value == s.Value
	case attributeOperation.Contains:
		return strings.Contains(value, s.Value)
	case attributeOperation.StartsWith:
		return strings.HasPrefix(value, s.Value)
	case attributeOperation.EndsWith:
		return strings.HasSuffix(value, s.Value)
	case attributeOperation.HyphenSeparated:
		return value == s.Value || strings.HasPrefix(value, s.Value+"-")
	default:
		return false
	}
}

func parseAttributeSelector(tokens []string) (*attributeSelector, int, error) {
	if len(tokens) < 3 || tokens[0] != "[" || tokens[len(tokens)-1] != "]" {
		return nil, 0, errors.New("invalid attribute selector")
	}

	attrSelector := &attributeSelector{}
	attrParts := strings.Join(tokens[1:len(tokens)-1], "")

	if strings.Contains(attrParts, "=") {
		var opStr string
		parts := strings.SplitN(attrParts, "=", 2)
		attrSelector.Name = strings.TrimRight(parts[0], "^$*|")
		attrSelector.Value = strings.Trim(parts[1], "\"'")

		switch parts[0][len(parts[0])-1] {
		case '^':
			attrSelector.Operation = attributeOperation.StartsWith
			opStr = "^="
		case '$':
			attrSelector.Operation = attributeOperation.EndsWith
			opStr = "$="
		case '*':
			attrSelector.Operation = attributeOperation.Contains
			opStr = "*="
		case '|':
			attrSelector.Operation = attributeOperation.HyphenSeparated
			opStr = "|="
		default:
			attrSelector.Operation = attributeOperation.Exact
			opStr = "="
		}

		// Ensure the operation string is actually in the original selector
		if !strings.Contains(attrParts, opStr) {
			return nil, 0, fmt.Errorf("invalid attribute selector operation: %s", opStr)
		}
	} else {
		attrSelector.Name = attrParts
		attrSelector.Operation = attributeOperation.Exists
	}

	return attrSelector, len(tokens), nil
}
