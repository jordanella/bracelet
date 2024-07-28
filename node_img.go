package bracelet

import (
	"strconv"

	paintbrush "github.com/jordanella/go-ansi-paintbrush"
)

type ImgNode struct {
	Element
	canvas      *paintbrush.Canvas
	file_loaded bool
	updated     bool
}

func (n ImgNode) Create() NodeFactory {
	canvas := paintbrush.New()

	return func(tag string) Node {
		return &ImgNode{
			Element: NewElement(tag),
			canvas:  canvas,
		}
	}

}

func (n *ImgNode) SetAttribute(attr string, value string) {
	triggerAttributes := map[string]func(){
		"src":  func() { n.file_loaded = n.canvas.LoadImage(value) == nil },
		"font": func() { _ = n.canvas.LoadFont(value) },
	}
	if action, isTrigger := triggerAttributes[attr]; isTrigger {
		if n.GetAttribute(attr) != value {
			action()
			n.updated = false
		}
	}
	n.Attributes[attr] = value
}

func (n *ImgNode) SetProperty(attr string, value string) {
	triggerProperties := map[string]struct{}{
		"width":  {},
		"height": {},
	}
	if _, isTrigger := triggerProperties[attr]; isTrigger {
		if n.GetProperty(attr) != value {
			n.updated = false
		}
	}
	n.Properties[attr] = value
}

func (n *ImgNode) ConvertImage() {
	width, _ := strconv.Atoi(n.GetProperty("width"))
	height, _ := strconv.Atoi(n.GetProperty("height"))
	if !n.file_loaded {
		return
	}
	n.canvas.SetWidth(width)
	n.canvas.SetHeight(height)
	n.canvas.Paint()
	result := n.canvas.GetResult()
	n.SetContent(result)
	n.updated = true
}

func (n *ImgNode) Serve() string {
	if !n.updated {
		n.ConvertImage()
	}
	return n.Element.Serve()
}
