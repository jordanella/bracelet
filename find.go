package bracelet

// Find searches for a single node matching the given CSS selector, starting from the root node.
// It returns a pointer to the first matching Node, or nil if no match is found.
func Find(root Node, selector string) *Node {
	parsedSelector, err := parseSelector(selector)
	if err != nil {
		return nil
	}

	var results []*Node
	results = traverseNodes(&root, parsedSelector, results, false)
	if len(results) == 1 {
		return results[0]
	}
	return nil
}

// FindAll searches for all nodes matching the given CSS selector, starting from the root node.
// It returns a slice of pointers to all matching Nodes.
func FindAll(root Node, selector string) []*Node {
	parsedSelector, err := parseSelector(selector)
	if err != nil {
		return []*Node{}
	}

	var results []*Node
	results = traverseNodes(&root, parsedSelector, results, true)
	return results
}

func traverseNodes(node *Node, selector Selector, results []*Node, all bool) []*Node {
	rule := Rule{
		Selectors: []Selector{selector},
	}

	if match := matchRule(node, rule); match != nil {
		results = append(results, node)
		if !all {
			return results
		}
	}

	for _, child := range (*node).GetChildren() {
		results = traverseNodes(child, selector, results, all)
		if !all && len(results) > 0 {
			return results
		}
	}
	if len(results) == 0 {
		return nil
	}

	return results
}
