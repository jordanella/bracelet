package main

import (
	"fmt"

	"github.com/jordanella/bracelet"
)

func main() {
	htmlContent := `
		<body>
			<p>
				Text     with an      <span>example of an inline</span> style.
			</p>
		</body>
		`

	cssContent := `
		body { width: 50; border: rounded #aa55aa; }
		p { margin: 1 2; word-spacing: 2; }
		span { font-style: bold; }
	`

	root, err := bracelet.ParseHTML(htmlContent)
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return
	}

	stylesheet, err := bracelet.ParseCSS(cssContent)
	if err != nil {
		fmt.Printf("Error parsing CSS: %v\n", err)
		return
	}

	bracelet.ApplyStylesheet(&root, stylesheet)

	fmt.Println(root.Serve())
}
