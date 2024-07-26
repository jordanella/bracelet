package bracelet

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gorilla/css/scanner"
)

type Declaration struct {
	Name  string
	Value string
}

type Rule struct {
	Selectors    []Selector
	Declarations []Declaration
}

type Stylesheet struct {
	Rules []Rule
}

// ParseCSS parses a CSS string and returns a slice of Rules.
// It handles selectors, declarations, and nested rules.
func ParseCSS(cssContent string) ([]Rule, error) {

	stylesheet := []Rule{}
	s := scanner.New(cssContent)

	var State = struct{ Selector, Declaration, Value int }{0, 1, 2}
	currentState := State.Selector
	currentRule := Rule{}
	currentSelector := ""
	currentDeclaration := Declaration{}
	declarationString := ""

	for {
		token := s.Next()

		if token.Type == scanner.TokenEOF {
			break
		}

		switch token.Value {
		case "{":
			if currentState == State.Selector {
				selector, err := parseSelector(strings.TrimSpace(currentSelector))
				if err == nil {
					currentRule.Selectors = append(currentRule.Selectors, selector)
				} else {
					fmt.Printf("Error parsing selector: %v\n", err)
					return stylesheet, err
				}
				currentSelector = ""
				currentState = State.Declaration
			}
		case "}":
			if len(currentRule.Selectors) > 0 && len(currentRule.Declarations) > 0 {
				stylesheet = append(stylesheet, currentRule)
			}
			currentRule = Rule{}
			currentState = State.Selector
		case ":":
			if currentState == State.Declaration {
				currentDeclaration.Name = strings.TrimSpace(currentDeclaration.Name)
				currentState = State.Value
				declarationString = ""
			} else {
				declarationString += token.Value
			}
		case ";":
			if currentState == State.Value {
				currentDeclaration.Value = strings.TrimSpace(declarationString)
				currentRule.Declarations = append(currentRule.Declarations, currentDeclaration)
				currentDeclaration = Declaration{}
				declarationString = ""
				currentState = State.Declaration
			}
		default:
			switch currentState {
			case State.Selector:
				currentSelector += token.Value
			case State.Declaration:
				currentDeclaration.Name += token.Value
			case State.Value:
				declarationString += token.Value
			}
		}
	}

	return stylesheet, nil
}

// ParseInlineStyle parses an inline style string and returns a PropertyMap.
// The inline style string should be in the format "property: value; property: value;".
func ParseInlineStyle(inlineStyle string) map[string]string {
	properties := make(map[string]string)
	declarations := strings.Split(inlineStyle, ";")
	for _, declaration := range declarations {
		parts := strings.SplitN(declaration, ":", 2)
		if len(parts) == 2 {
			property := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			properties[property] = value
		}
	}
	return properties
}

// DetermineProperties calculates the final set of CSS properties for a given node,
// taking into account the stylesheet rules and any inline styles.
func DetermineProperties(node *Node, stylesheet []Rule) map[string]string {
	properties := make(map[string]string)

	rules := matchingRules(node, stylesheet)
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].Specificity.Less(rules[j].Specificity)
	})
	for _, matchedRule := range rules {
		for _, declaration := range matchedRule.Rule.Declarations {
			properties[declaration.Name] = declaration.Value
		}
	}

	attributes := (*node).GetAttributes()
	if inlineStyle, ok := attributes["style"]; ok {
		inlineProperties := ParseInlineStyle(inlineStyle)
		for property, value := range inlineProperties {
			properties[property] = value
		}
	}

	return properties
}

// ApplyStylesheet applies the given stylesheet to the node and all its descendants.
// It calculates and sets the final properties for each node in the tree.
func ApplyStylesheet(node *Node, stylesheet []Rule) {
	if node == nil {
		fmt.Println("Warning: nil node passed to ApplyStylesheet")
		return
	}

	properties := DetermineProperties(node, stylesheet)
	(*node).AddProperties(properties)

	for _, child := range (*node).GetChildren() {
		ApplyStylesheet(child, stylesheet)
	}
}
