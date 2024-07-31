package bracelet

type TextNode struct {
	Element
}

func (n TextNode) Create() NodeFactory {
	return func(tag string) Node {
		return &TextNode{
			Element: NewElement(tag),
		}
	}
}

func (n *TextNode) AddProperties(properties map[string]string) {
	inherit := []string{"color", "font-weight", "text-transform", "text-align", "background-color", "width", "height", "word-spacing"}
	parentProperties := (*n.GetParent()).GetProperties()

	for key, value := range properties {
		n.SetProperty(key, value)
	}

	for _, inheritable := range inherit {
		if value, exists := parentProperties[inheritable]; exists {
			if _, exists := n.Properties[inheritable]; !exists {
				n.SetProperty(inheritable, value)
			}
		}
	}
}
