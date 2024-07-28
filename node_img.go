package bracelet

import (
	"strconv"

	paintbrush "github.com/jordanella/go-ansi-paintbrush"
)

type ImgNode struct {
	Element
	paintbrush  paintbrush.AnsiArtInterface
	file_loaded bool
	updated     bool
}

func (n ImgNode) Create() NodeFactory {
	paintbrush := paintbrush.New()

	return func(tag string) Node {
		return &ImgNode{
			Element:    NewElement(tag),
			paintbrush: paintbrush,
		}
	}
}

func (n *ImgNode) SetAttribute(attr string, value string) {
	triggerAttributes := map[string]func(){
		"src":  func() { n.file_loaded = n.paintbrush.LoadImage(value) == nil },
		"font": func() { _ = n.paintbrush.LoadFont(value) },
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
		"width": {},
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
	if !n.file_loaded || width == 0 {
		return
	}
	n.paintbrush.SetWidth(width)
	n.paintbrush.Render()
	result := n.paintbrush.GetResultRaw()
	n.SetContent(result)
	n.updated = true
}

func (n *ImgNode) Serve() string {
	if !n.updated {
		n.ConvertImage()
	}
	return n.Element.Serve()
}
