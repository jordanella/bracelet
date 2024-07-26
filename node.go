package bracelet

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Node represents an HTML element in the document tree.
type Node interface {
	// Serve renders the node and its children, returning the final string representation.
	Serve() string

	// Create returns a NodeFactory function for creating new instances of this node type.
	Create() NodeFactory

	// GetTag returns the HTML tag name of the node.
	GetTag() string

	// GetID returns the ID attribute of the node.
	GetID() string

	// GetStyle returns the lipgloss.Style associated with the node.
	GetStyle() lipgloss.Style

	// GetContent returns the text content of the node.
	GetContent() string

	// GetClasses returns a slice of CSS class names associated with the node.
	GetClasses() []string

	// GetAttributes returns a map of all attributes set on the node.
	GetAttributes() map[string]string

	// GetAttribute returns the value of a specific attribute by its name.
	// If the attribute doesn't exist, it returns an empty string.
	GetAttribute(string) string

	// GetProperties returns a PropertyMap containing all CSS properties applied to the node.
	GetProperties() map[string]string

	// GetProperty returns the value of a specific CSS property by its name.
	// If the property doesn't exist, it returns an empty string.
	GetProperty(string) string

	// GetParent returns a pointer to the parent node, or nil if this is the root node.
	GetParent() *Node

	// GetChildren returns a slice of pointers to all child nodes.
	GetChildren() []*Node

	// SetTag sets the HTML tag name of the node.
	SetTag(string)

	// SetID sets the ID attribute of the node.
	SetID(string)

	// SetStyle sets the lipgloss.Style for the node, which determines its appearance.
	SetStyle(lipgloss.Style)

	// SetContent sets the text content of the node.
	SetContent(string)

	// SetClasses sets the CSS class names for the node.
	SetClasses([]string)

	// AddClass adds one more CSS class names for the node.
	AddClass(...string)

	// RemoveClass removes one or more classes from the node's classes slice.
	RemoveClass(...string)

	// HasClass returns true if the node has the specified class, false otherwise.
	HasClass(string) bool

	// SetAttributes sets all attributes for the node, replacing any existing attributes.
	SetAttributes(map[string]string)

	// AddAttributes sets attributes for the node, adding to and updating existing attributes.
	AddAttributes(map[string]string)

	// SetAttribute sets a single attribute on the node. If the attribute already exists, its value is replaced.
	SetAttribute(string, string)

	// RemoveAttribute removes an attribute from the node's attributes map.
	RemoveAttribute(...string)

	// HasAttribute returns true if the node has the specified attribute, false otherwise.
	HasAttribute(key string) bool

	// SetProperties sets all CSS properties for the node, replacing any existing properties.
	SetProperties(map[string]string)

	// AddProperties sets CSS properties for the node, adding to and updating existing properties.
	AddProperties(map[string]string)

	// SetProperty sets a single CSS property on the node. If the property already exists, its value is replaced.
	SetProperty(string, string)

	// RemoveProperty removes one or more properties from the node's properties map.
	RemoveProperty(...string)

	// HasProperty returns true if the node has the specified property, false otherwise.
	HasProperty(string) bool

	// SetParent sets the parent node of this node.
	SetParent(*Node)

	// SetChildren sets all child nodes of this node, replacing any existing children.
	SetChildren([]*Node)

	// AddChild adds a new child node to this node.
	AddChild(*Node)

	// SetChild sets a child node at a specific index. If the index is out of range, the child is appended.
	SetChild(int, *Node)
}

type NodeFactory func(tag string) Node

var customNodeFactories = make(map[string]NodeFactory)

func createNode(tag string) Node {
	if factory, ok := customNodeFactories[tag]; ok {
		node := factory(tag)
		if node != nil {
			return ensureInitialized(node)
		}
		fmt.Printf("Warning: Custom factory for tag '%s' returned nil\n", tag)
	}
	return ensureInitialized(Element{}.Create()(tag))
}

func ensureInitialized(node Node) Node {
	if node.GetClasses() == nil {
		node.SetClasses([]string{})
	}
	if node.GetAttributes() == nil {
		node.SetAttributes(make(map[string]string))
	}
	if node.GetProperties() == nil {
		node.SetProperties(make(map[string]string))
	}
	if node.GetChildren() == nil {
		node.SetChildren([]*Node{})
	}
	return node
}
