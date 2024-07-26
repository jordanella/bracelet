package bracelet

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// ParseHTML parses an HTML string and returns the root Node of the resulting tree.
// It handles nested elements, attributes, and text nodes.
func ParseHTML(htmlContent string) (Node, error) {
	reader := strings.NewReader(htmlContent)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("html.Parse error: %v", err)
	}

	var buildTree func(*html.Node) (Node, error)
	buildTree = func(n *html.Node) (Node, error) {
		if n == nil {
			return nil, fmt.Errorf("nil html.Node encountered")
		}

		if n.Type == html.ElementNode {
			node := createNode(n.Data)
			if node == nil {
				return nil, fmt.Errorf("createNode returned nil for tag %s", n.Data)
			}

			for _, attr := range n.Attr {
				node.SetAttribute(attr.Key, attr.Val)
				if attr.Key == "id" {
					node.SetID(attr.Val)
				} else if attr.Key == "class" {
					node.AddClass(strings.Fields(attr.Val)...)
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				child, err := buildTree(c)
				if err != nil {
					return nil, fmt.Errorf("error building child node: %v", err)
				}
				if child != nil {
					childPtr := &child
					node.AddChild(childPtr)
					child.SetParent(&node)
				}
			}
			return node, nil
		} else if n.Type == html.TextNode {
			textContent := strings.TrimSpace(n.Data)
			if textContent != "" {
				textNode := createNode("text")
				if textNode == nil {
					return nil, fmt.Errorf("createNode returned nil for text element")
				}
				textNode.SetContent(textContent)
				return textNode, nil
			}
			return nil, nil
		} else if n.Type == html.DocumentNode {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				child, err := buildTree(c)
				if err != nil {
					return nil, fmt.Errorf("error building child node of document: %v", err)
				}
				if child != nil {
					return child, nil
				}
			}
		}
		return nil, nil
	}

	rootNode, err := buildTree(doc)
	if err != nil {
		return nil, fmt.Errorf("failed to build node tree: %v", err)
	}
	if rootNode == nil {
		return nil, fmt.Errorf("root node is nil after parsing")
	}
	return *Find(rootNode, "body"), nil
}

// PrintStyledHTML prints a styled representation of the HTML tree to the console.
// It includes computed styles for each element.
func PrintStyledHTML(node *Node, depth int) {
	if node == nil {
		return
	}

	indent := strings.Repeat("    ", depth)
	fmt.Printf("%s<%s", indent, (*node).GetTag())

	for key, value := range (*node).GetAttributes() {
		fmt.Printf(" %s=\"%s\"", key, value)
	}

	propertyList := []string{}
	for key, value := range (*node).GetProperties() {
		propertyList = append(propertyList, fmt.Sprintf("%s: %s", key, value))
	}
	fmt.Printf("> [Computed Styles] %s\n", strings.Join(propertyList, ", "))

	for _, child := range (*node).GetChildren() {
		PrintStyledHTML(child, depth+1)
	}

	fmt.Printf("%s</%s>\n", indent, (*node).GetTag())
}
