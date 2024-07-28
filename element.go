package bracelet

import "github.com/charmbracelet/lipgloss"

type Element struct {
	Tag        string
	ID         string
	Style      lipgloss.Style
	Content    string
	Classes    []string
	Attributes map[string]string
	Properties map[string]string
	Parent     *Node
	Children   []*Node
}

// Serve renders the Element and its children into a styled string.
// It applies all CSS properties to the Element's content and recursively
// renders child Elements. The method handles both leaf nodes (Elements with
// no children) and container nodes (Elements with children).
//
// For leaf nodes, it applies styling to the Element's content.
// For container nodes, it recursively renders children and joins them
// according to the specified layout direction (horizontal or vertical).
//
// The final output is a string that represents the fully styled Element,
// ready for display in a terminal interface.
func (e *Element) Serve() string {
	style := lipgloss.NewStyle()
	content := e.GetContent()
	for key, value := range e.GetProperties() {
		if function, exists := PropertyFunctions[key]; exists {
			content, style = function(value)(content, style)
		}
	}
	children := e.GetChildren()
	if len(children) != 0 {
		var contents []string
		for _, child := range children {
			contents = append(contents, (*child).Serve())
		}
		direction := e.GetProperty("direction")
		switch direction {
		case "horizontal":
			content = lipgloss.JoinHorizontal(style.GetAlignHorizontal(), contents...)
		case "vertical":
			content = lipgloss.JoinVertical(style.GetAlignVertical(), contents...)
		default:
			content = lipgloss.JoinHorizontal(style.GetAlignHorizontal(), contents...)
		}
	}
	return style.Render(content)
}

// NewElement creates a new Element with all fields properly initialized
func NewElement(tag string) Element {
	return Element{
		Tag:        tag,
		Classes:    []string{},
		Attributes: make(map[string]string),
		Properties: make(map[string]string),
		Children:   []*Node{},
	}
}

// Create creates a new Element with all fields properly initialized
func (e Element) Create() NodeFactory {
	return func(tag string) Node {
		element := NewElement(tag)
		return &element
	}
}

// GetTag returns the HTML tag name of the node.
func (n *Element) GetTag() string { return n.Tag }

// GetID returns the ID attribute of the node.
func (n *Element) GetID() string { return n.ID }

// GetStyle returns the lipgloss.Style associated with the node.
func (n *Element) GetStyle() lipgloss.Style { return n.Style }

// GetContent returns the text content of the node.
func (n *Element) GetContent() string { return n.Content }

// GetClasses returns a slice of CSS class names associated with the node.
func (n *Element) GetClasses() []string { return n.Classes }

// GetAttributes returns a map of all attributes set on the node.
func (n *Element) GetAttribute(attr string) string {
	if val, ok := n.Attributes[attr]; ok {
		return val
	} else {
		return ""
	}
}

// GetAttribute returns the value of a specific attribute by its name.
// If the attribute doesn't exist, it returns an empty string.
func (n *Element) GetAttributes() map[string]string { return n.Attributes }

// GetProperties returns a PropertyMap containing all CSS properties applied to the node.

func (n *Element) GetProperties() map[string]string { return n.Properties }

// GetProperty returns the value of a specific CSS property by its name.
// If the property doesn't exist, it returns an empty string.
func (n *Element) GetProperty(prop string) string {
	if val, ok := n.Properties[prop]; ok {
		return val
	} else {
		return ""
	}
}

// GetParent returns a pointer to the parent node, or nil if this is the root node.
func (n *Element) GetParent() *Node { return n.Parent }

// GetChildren returns a slice of pointers to all child nodes.
func (n *Element) GetChildren() []*Node { return n.Children }

// SetTag sets the HTML tag name of the node.
func (n *Element) SetTag(tag string) { n.Tag = tag }

// SetID sets the ID attribute of the node.
func (n *Element) SetID(id string) { n.ID = id }

// SetStyle sets the lipgloss.Style for the node, which determines its appearance.
func (n *Element) SetStyle(style lipgloss.Style) { n.Style = style }

// SetContent sets the text content of the node.
func (n *Element) SetContent(content string) { n.Content = content }

// SetClasses sets the CSS class names for the node.
func (n *Element) SetClasses(classes []string) { n.Classes = classes }

// AddClass adds one more CSS class names for the node.
func (n *Element) AddClass(classes ...string) { n.Classes = append(n.Classes, classes...) }

// RemoveClass removes one or more classes from the node's Classes slice.
func (e *Element) RemoveClass(classes ...string) {
	var found bool
	newClasses := make([]string, 0, len(e.Classes))
	for _, c := range e.Classes {
		found = false
		for _, class := range classes {
			if c == class {
				found = true
			}
		}
		if !found {
			newClasses = append(newClasses, c)
		}
	}
	e.Classes = newClasses
}

// HasClass returns true if the node has the specified class, false otherwise.
func (e *Element) HasClass(class string) bool {
	for _, c := range e.Classes {
		if c == class {
			return true
		}
	}
	return false
}

// SetAttributes sets all attributes for the node, replacing any existing attributes.
func (n *Element) SetAttributes(attributes map[string]string) {
	n.Attributes = map[string]string{}
	n.AddAttributes(attributes)
}

// SetAttribute sets a single attribute on the node. If the attribute already exists, its value is replaced.
func (n *Element) SetAttribute(attr string, value string) {
	n.Attributes[attr] = value
}

// AddAttributes sets attributes for the node, adding to and updating existing attributes.
func (n *Element) AddAttributes(attributes map[string]string) {
	for key, value := range attributes {
		n.SetAttribute(key, value)
	}
}

// RemoveAttribute removes an attribute from the node's Attributes map.
func (e *Element) RemoveAttribute(keys ...string) {
	for _, key := range keys {
		delete(e.Attributes, key)
	}
}

// HasAttribute returns true if the node has the specified attribute, false otherwise.
func (e *Element) HasAttribute(key string) bool {
	_, exists := e.Attributes[key]
	return exists
}

// SetProperties sets all CSS properties for the node, replacing any existing properties.
func (n *Element) SetProperties(properties map[string]string) {
	n.Properties = map[string]string{}
	n.AddProperties(properties)
}

// SetProperty sets a single CSS property on the node. If the property already exists, its value is replaced.
func (n *Element) SetProperty(prop string, value string) {
	n.Properties[prop] = value
}

// AddProperties sets CSS properties for the node, adding to and updating existing properties.
func (n *Element) AddProperties(properties map[string]string) {
	for key, value := range properties {
		n.SetProperty(key, value)
	}
}

// RemoveProperty removes one or more properties from the node's Properties map.
func (e *Element) RemoveProperty(keys ...string) {
	for _, key := range keys {
		delete(e.Properties, key)
	}
}

// HasProperty returns true if the node has the specified property, false otherwise.
func (e *Element) HasProperty(key string) bool {
	_, exists := e.Properties[key]
	return exists
}

// SetParent sets the parent node of this node.
func (n *Element) SetParent(parent *Node) { n.Parent = parent }

// SetChildren sets all child nodes of this node, replacing any existing children.
func (n *Element) SetChildren(children []*Node) { n.Children = children }

// AddChild adds a new child node to this node.
func (n *Element) AddChild(child *Node) { n.Children = append(n.Children, child) }

// SetChild sets a child node at a specific index. If the index is out of range, the child is appended.
func (n *Element) SetChild(index int, child *Node) {
	if index < 0 || index >= len(n.Children) {
		n.Children = append(n.Children, child)
	} else {
		n.Children[index] = child
	}
	if child != nil {
		nodePtr := Node(n)
		(*child).SetParent(&nodePtr)
	}
}
