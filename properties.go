package bracelet

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type PropertyFunction func(string, lipgloss.Style) (string, lipgloss.Style)

// PropertyFunctions maps CSS property names to their corresponding PropertyFunction.
// Each PropertyFunction takes a property value as input and returns a function
// that applies that property to a node's content and style.
var PropertyFunctions = map[string]func(string) PropertyFunction{
	"color":            AttrColor,
	"background-color": AttrBackgroundColor,
	"font-weight":      AttrFontWeight,
	"text-transform":   AttrTextTransform,
	"font-style":       AttrFontStyle,
	"text-decoration":  AttrTextDecoration,
	"margin":           AttrMargin,
	"margin-top":       AttrMarginTop,
	"margin-bottom":    AttrMarginBottom,
	"margin-left":      AttrMarginLeft,
	"margin-right":     AttrMarginRight,
	"padding":          AttrPadding,
	"padding-top":      AttrPaddingTop,
	"padding-bottom":   AttrPaddingBottom,
	"padding-left":     AttrPaddingLeft,
	"padding-right":    AttrPaddingRight,
	"border":           AttrBorder,
	"border-top":       AttrBorderTop,
	"border-bottom":    AttrBorderBottom,
	"border-left":      AttrBorderLeft,
	"border-right":     AttrBorderRight,
	"width":            AttrWidth,
	"height":           AttrHeight,
	"text-align":       AttrTextAlign,
	"vertical-align":   AttrVerticalAlign,
	"indent":           AttrIndent,
	"text-indent":      AttrIndent,
	"word-spacing":     AttrWordSpacing,
}

// ApplyProperty looks up the appropriate PropertyFunction and applies it to a node's content and style.
// If the property is not recognized, no changes are made to the Node.
func ApplyProperty(node *Node, property string, value string) {
	if attrFunc, ok := PropertyFunctions[property]; ok {
		content, style := attrFunc(value)((*node).GetContent(), (*node).GetStyle())
		(*node).SetContent(content)
		(*node).SetStyle(style)
	}
}

// AttrColor returns a PropertyFunction that sets the foreground color of the text.
func AttrColor(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		return content, style.Foreground(lipgloss.Color(value))
	}
}

// AttrBackgroundColor returns a PropertyFunction that sets the background color of the text.
func AttrBackgroundColor(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		return content, style.Background(lipgloss.Color(value))
	}
}

// AttrFontWeight returns a PropertyFunction that sets the font weight.
// Currently, it only supports making text bold.
func AttrFontWeight(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if value == "bold" {
			return content, style.Bold(true)
		}
		return content, style
	}
}

// AttrTextTransform returns a PropertyFunction that transforms the text content.
// Supports uppercase, lowercase, and capitalize.
func AttrTextTransform(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		switch value {
		case "uppercase":
			return strings.ToUpper(content), style
		case "lowercase":
			return strings.ToLower(content), style
		case "capitalize":
			return strings.ToTitle(content), style
		default:
			return content, style
		}
	}
}

// AttrFontStyle returns a PropertyFunction that sets the font style.
// Supports italic and bold.
func AttrFontStyle(value string) PropertyFunction {
	var italic, bold, normal bool
	values := strings.Split(value, " ")
	for _, val := range values {
		if val == "italic" {
			italic = true
		}
		if val == "bold" {
			bold = true
		}
		if val == "normal" {
			normal = true
		}
	}
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if normal {
			style = style.Italic(false)
			style = style.Bold(false)
		}
		if italic {
			style = style.Italic(true)
		}
		if bold {
			style = style.Bold(true)
		}
		return content, style
	}
}

// AttrTextDecoration returns a PropertyFunction that sets text decoration.
// Supports underline and line-through.
func AttrTextDecoration(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		switch value {
		case "underline":
			return content, style.Underline(true)
		case "line-through":
			return content, style.Strikethrough(true)
		default:
			return content, style
		}
	}
}

// AttrMargin returns a PropertyFunction that sets all margins.
// Accepts up to four space-separated values.
func AttrMargin(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		parts := strings.Fields(value)
		var values []int
		for _, part := range parts {
			if v, err := strconv.Atoi(part); err == nil {
				values = append(values, v)
			}
		}
		return content, style.Margin(values...)
	}
}

// AttrMarginLeft returns a PropertyFunction that sets the left margin.
func AttrMarginLeft(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.MarginLeft(v)
		}
		return content, style
	}
}

// AttrMarginRight returns a PropertyFunction that sets the right margin.
func AttrMarginRight(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.MarginRight(v)
		}
		return content, style
	}
}

// AttrMarginTop returns a PropertyFunction that sets the top margin.
func AttrMarginTop(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.MarginTop(v)
		}
		return content, style
	}
}

// AttrMarginBottom returns a PropertyFunction that sets the bottom margin.
func AttrMarginBottom(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.MarginBottom(v)
		}
		return content, style
	}
}

// AttrPadding returns a PropertyFunction that sets all padding.
// Accepts up to four space-separated values.
func AttrPadding(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		parts := strings.Fields(value)
		var values []int
		for _, part := range parts {
			if v, err := strconv.Atoi(part); err == nil {
				values = append(values, v)
			}
		}
		return content, style.Padding(values...)
	}
}

// AttrPaddingLeft returns a PropertyFunction that sets the left padding.
func AttrPaddingLeft(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.PaddingLeft(v)
		}
		return content, style
	}
}

// AttrPaddingRight returns a PropertyFunction that sets the right padding.
func AttrPaddingRight(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.PaddingRight(v)
		}
		return content, style
	}
}

// AttrPaddingTop returns a PropertyFunction that sets the top padding.
func AttrPaddingTop(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.PaddingTop(v)
		}
		return content, style
	}
}

// AttrPaddingBottom returns a PropertyFunction that sets the bottom padding.
func AttrPaddingBottom(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.PaddingBottom(v)
		}
		return content, style
	}
}

func parseBorderStyle(value string) lipgloss.Border {
	var borderType lipgloss.Border
	switch strings.ToLower(value) {
	case "normal":
		borderType = lipgloss.NormalBorder()
	case "rounded":
		borderType = lipgloss.RoundedBorder()
	case "block":
		borderType = lipgloss.BlockBorder()
	case "double":
		borderType = lipgloss.DoubleBorder()
	case "hidden":
		borderType = lipgloss.HiddenBorder()
	case "inner", "innerhalf", "inner-half", "half":
		borderType = lipgloss.InnerHalfBlockBorder()
	case "outer", "outerhalf", "outer-half":
		borderType = lipgloss.OuterHalfBlockBorder()
	case "thick":
		borderType = lipgloss.ThickBorder()
	}
	return borderType
}

func parseBorderArgs(args []string) (bool, lipgloss.Border, lipgloss.Color, lipgloss.Color) {
	var (
		show       = true
		style      = lipgloss.NormalBorder()
		foreground = lipgloss.Color("")
		background = lipgloss.Color("")
	)

	for _, arg := range args {
		arg = strings.ToLower(arg)
		switch arg {
		case "true", "1", "on", "yes":
			show = true
			continue
		case "false", "0", "off", "no", "none":
			show = false
			continue
		case "normal", "rounded", "block", "double", "hidden", "inner", "innerhalf", "inner-half", "half", "outer", "outerhalf", "outer-half", "thick":
			style = parseBorderStyle(arg)
			continue
		default:
			if foreground == lipgloss.Color("") {
				foreground = lipgloss.Color(arg)
			} else {
				background = lipgloss.Color(arg)
			}
		}
	}

	return show, style, foreground, background
}

// AttrBorder returns a PropertyFunction that sets all borders.
// Accepts style, color, and width arguments.
func AttrBorder(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		args := strings.Fields(value)
		show, borderStyle, fg, bg := parseBorderArgs(args)

		if !show {
			return content, style.Border(lipgloss.HiddenBorder())
		}

		style = style.Border(borderStyle)
		if fg != lipgloss.Color("") {
			style = style.BorderForeground(fg)
		}
		if bg != lipgloss.Color("") {
			style = style.BorderBackground(bg)
		}

		return content, style
	}
}

// AttrBorderTop returns a PropertyFunction that sets the top border.
// Accepts style, color, and width arguments.
func AttrBorderTop(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		args := strings.Fields(value)
		show, borderStyle, fg, bg := parseBorderArgs(args)

		style = style.BorderTop(show)
		if show {
			style = style.BorderStyle(borderStyle)
			if fg != lipgloss.Color("") {
				style = style.BorderTopForeground(fg)
			}
			if bg != lipgloss.Color("") {
				style = style.BorderTopBackground(bg)
			}
		}

		return content, style
	}
}

// AttrBorderBottom returns a PropertyFunction that sets the bottom border.
// Accepts style, color, and width arguments.
func AttrBorderBottom(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		args := strings.Fields(value)
		show, borderStyle, fg, bg := parseBorderArgs(args)

		style = style.BorderBottom(show)
		if show {
			style = style.BorderStyle(borderStyle)
			if fg != lipgloss.Color("") {
				style = style.BorderBottomForeground(fg)
			}
			if bg != lipgloss.Color("") {
				style = style.BorderBottomBackground(bg)
			}
		}

		return content, style
	}
}

// AttrBorderLeft returns a PropertyFunction that sets the left border.
// Accepts style, color, and width arguments.
func AttrBorderLeft(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		args := strings.Fields(value)
		show, borderStyle, fg, bg := parseBorderArgs(args)

		style = style.BorderLeft(show)
		if show {
			style = style.BorderStyle(borderStyle)
			if fg != lipgloss.Color("") {
				style = style.BorderLeftForeground(fg)
			}
			if bg != lipgloss.Color("") {
				style = style.BorderLeftBackground(bg)
			}
		}

		return content, style
	}
}

// AttrBorderRight returns a PropertyFunction that sets the right border.
// Accepts style, color, and width arguments.
func AttrBorderRight(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		args := strings.Fields(value)
		show, borderStyle, fg, bg := parseBorderArgs(args)

		style = style.BorderRight(show)
		if show {
			style = style.BorderStyle(borderStyle)
			if fg != lipgloss.Color("") {
				style = style.BorderRightForeground(fg)
			}
			if bg != lipgloss.Color("") {
				style = style.BorderRightBackground(bg)
			}
		}

		return content, style
	}
}

// AttrWidth returns a PropertyFunction that sets the width of the element.
func AttrWidth(width string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		w, _ := strconv.Atoi(width)
		return content, style.Width(w)
	}
}

// AttrHeight returns a PropertyFunction that sets the height of the element.
func AttrHeight(height string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		h, _ := strconv.Atoi(height)
		return content, style.Height(h)
	}
}

// AttrTextAlign returns a PropertyFunction that sets the horizontal text alignment.
// Supports left, center, and right.
func AttrTextAlign(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		switch strings.ToLower(value) {
		case "left":
			return content, style.AlignHorizontal(lipgloss.Left)
		case "center":
			return content, style.AlignHorizontal(lipgloss.Center)
		case "right":
			return content, style.AlignHorizontal(lipgloss.Right)
		default:
			return content, style
		}
	}
}

// AttrVerticalAlign returns a PropertyFunction that sets the vertical alignment.
// Supports top, center, and bottom.
func AttrVerticalAlign(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		switch strings.ToLower(value) {
		case "top":
			return content, style.AlignVertical(lipgloss.Top)
		case "center":
			return content, style.AlignVertical(lipgloss.Center)
		case "bottom":
			return content, style.AlignVertical(lipgloss.Bottom)
		default:
			return content, style
		}
	}
}

// AttrIndent returns a PropertyFunction that sets the text indentation.
func AttrIndent(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		indent, _ := strconv.Atoi(value)
		return content, style.MarginLeft(style.GetMarginLeft() + indent)
	}
}

// AttrWordSpacing returns a PropertyFunction that sets the spacing between words.
func AttrWordSpacing(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		spacing, _ := strconv.Atoi(value)
		content = strings.Join(strings.Split(content, " "), strings.Repeat(" ", spacing))
		return content, style
	}
}
