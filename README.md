# Bracelet

Bracelet is a Go package that provides a flexible and powerful framework for parsing, manipulating, and rendering HTML-like structures with CSS-like styling. It's designed to be used in terminal-based user interfaces, leveraging the [lipgloss](https://github.com/charmbracelet/lipgloss) library for styling.

## Features

- Parse HTML-like structures into a node tree
- Apply CSS-like styling to nodes
- Flexible selector system for targeting specific nodes
- Support for pseudo-selectors like `:first-child`, `:last-child`, and `:nth-child(n)`
- Extensible and trivial to implement custom node elements
- Render styled nodes to string output suitable for terminal display

## Installation

To install Bracelet, use `go get`:

```bash
go get github.com/jordanella/bracelet
```

## Quick Start

Here's a simple example of how to use Bracelet:

```go
package main

import (
    "fmt"
    "github.com/jordanella/bracelet"
)

func main() {
    html := `
    <div>
        <h1>Hello, Bracelet!</h1>
        <p>This is a <em>styled</em> paragraph.</p>
    </div>
    `
    css := `
    h1 { color: blue; font-weight: bold; }
    p { margin-left: 2; }
    em { font-style: italic; color: red; }
    `

    root, _ := bracelet.ParseHTML(html)
    rules, _ := bracelet.ParseCSS(css)
    bracelet.ApplyStylesheet(&root, rules)

    fmt.Println(root.Serve())
}
```

## Documentation

For detailed documentation, run `godoc -http=:6060` and navigate to `http://localhost:6060/pkg/github.com/jordanella/bracelet/` in your web browser.

## Key Components

- `Node`: Interface representing an HTML-like element
- `Element`: Basic implementation of the `Node` interface
- `ParseHTML`: Function to parse HTML strings into a node tree
- `ParseCSS`: Function to parse CSS strings into rules
- `ApplyStylesheet`: Function to apply CSS rules to a node tree
- `Find` and `FindAll`: Functions to select nodes using CSS-like selectors
- `Serve`: Method to render a node and its children into a styled string

## Rendering

The `Serve` method is the core of Bracelet's rendering process. It's responsible for turning your styled nodes into strings that can be displayed in a terminal interface.

Here's how you typically use it:

```go
root, _ := bracelet.ParseHTML(htmlString)
rules, _ := bracelet.ParseCSS(cssString)
bracelet.ApplyStylesheet(&root, rules)

// Render the entire tree to a string
output := root.Serve()
fmt.Println(output)
```

The `Serve` method:

1. Applies all CSS properties to the node's content
2. Recursively renders child nodes
3. Joins child nodes according to the specified layout direction (horizontal or vertical)
4. Returns a fully styled string representation of the node and its children

This powerful method allows you to easily convert your HTML-like structures with CSS-like styling into terminal-ready output.

## Custom Node Elements

Bracelet makes it easy to implement custom node elements and register them for use in applications. This is particularly useful for making reusable components and when integrating with other libraries like BubbleTea. Here's a brief overview:

1. Create a new type that embeds `bracelet.Element`:

```go
type CustomNode struct {
    bracelet.Element
}
```

2. Implement the `Create()` method for your custom node:

```go
func (n CustomNode) Create() teatray.NodeFactory {
    return func(tag string) teatray.Node {
        return &CustomNode{
            Element: teatray.NewElement(tag)
        }
    }
}
```

If you would like to prepopulate properties, you can define the element yourself; NewElement is just an empty Element constructor.
```go
Element: Element{
    Tag:        tag,
    Classes:    []string{"default", "class", "definitions"},
    Attributes: map[string]string{"style": "margin: 0 2;" },
    Properties: PropertyMap{"direction": "vertical"},
    Children:   []*Node{},
}
```

All nodes created by Bracelet, whether through HTML parsing or custom node factories, 
are preprocessed to ensure their `Classes`, `Attributes`, `Properties`, and `Children` 
collections are initialized to empty (but non-nil) values. This ensures that operations 
on these collections will not cause nil pointer dereferences, even if the collections 
are empty.

3. Register your custom node:

```go
func init() {
    bracelet.RegisterNode("custom", &CustomNode{})
}
```

Now you can use your custom node type in your HTML-like structures:

```html
<custom>This is a custom node</custom>
```

This flexibility allows you to easily extend Bracelet's functionality and integrate it with other libraries like BubbleTea for creating rich, interactive terminal user interfaces.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

- [lipgloss](https://github.com/charmbracelet/lipgloss) for providing the underlying styling capabilities
- Inspired by web technologies and adapted for terminal-based interfaces