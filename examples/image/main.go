package main

import (
	"fmt"

	_ "image/png"

	"github.com/jordanella/bracelet"
)

func main() {

	htmlContent := `
	<body>
		<sidebar style="direction: vertical;">
			<nav>
				<header>navigation</header>
				<item>Option 1</item>
				<item class="selected">Option 2</item>
				<item>Option 3</item>
			</nav>
			<img src="examples/norman.png" style="margin: 1 2; width: 40; height: 11; text-align: center;" />
		</sidebar>
		<content style="text-align: center;">But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness. No one rejects, dislikes, or avoids pleasure itself, because it is pleasure, but because those who do not know how to pursue pleasure rationally encounter consequences that are extremely painful. Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure. To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it? But who has any right to find fault with a man who chooses to enjoy a pleasure that has no annoying consequences, or one who avoids a pain that produces no resultant pleasure?</div>
		</content>
	</body>
    `

	cssContent := `
    body { width: 100; border: rounded #aa55aa; height: 20; vertical-align: top; }
    nav { width: 40; border: rounded #44ddff; direction: vertical; margin: 0 1; padding: 1 4; }
	header { text-transform: uppercase; font-style: bold; margin-bottom: 1; }
	item { margin-left: 1; }
	item.selected { color: #ff55dd; }
	content { width: 50; height: 18; margin: 1 2; vertical-align: center; }
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

	matchingNodes := bracelet.FindAll(root, "nav > item:not(.selected) > text")
	for _, node := range matchingNodes {
		(*node).SetContent("- " + (*node).GetContent())
	}
	matchingNode := bracelet.Find(root, "nav > item.selected > text")
	if matchingNode != nil {
		(*matchingNode).SetContent("+ " + (*matchingNode).GetContent())
	}

	fmt.Println(root.Serve())
}
