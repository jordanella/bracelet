package bracelet

func init() {
	RegisterNode("text", &TextNode{})
	RegisterNode("img", &ImgNode{})
}

// RegisterNode registers a custom node type for a specific HTML tag.
// This allows users to extend the package with custom element implementations.
func RegisterNode(tag string, node Node) {
	customNodeFactories[tag] = node.Create()
}
