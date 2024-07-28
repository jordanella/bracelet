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
	"color":            PropColor,
	"background-color": PropBackgroundColor,
	"font-weight":      PropFontWeight,
	"text-transform":   PropTextTransform,
	"font-style":       PropFontStyle,
	"text-decoration":  PropTextDecoration,
	"margin":           PropMargin,
	"margin-top":       PropMarginTop,
	"margin-bottom":    PropMarginBottom,
	"margin-left":      PropMarginLeft,
	"margin-right":     PropMarginRight,
	"padding":          PropPadding,
	"padding-top":      PropPaddingTop,
	"padding-bottom":   PropPaddingBottom,
	"padding-left":     PropPaddingLeft,
	"padding-right":    PropPaddingRight,
	"border":           PropBorder,
	"border-top":       PropBorderTop,
	"border-bottom":    PropBorderBottom,
	"border-left":      PropBorderLeft,
	"border-right":     PropBorderRight,
	"width":            PropWidth,
	"height":           PropHeight,
	"text-align":       PropTextAlign,
	"vertical-align":   PropVerticalAlign,
	"indent":           PropIndent,
	"text-indent":      PropIndent,
	"word-spacing":     PropWordSpacing,
}

// ApplyProperty looks up the appropriate PropertyFunction and applies it to a node's content and style.
// If the property is not recognized, no changes are made to the Node.
func ApplyProperty(node *Node, property string, value string) {
	if propFunc, ok := PropertyFunctions[property]; ok {
		content, style := propFunc(value)((*node).GetContent(), (*node).GetStyle())
		(*node).SetContent(content)
		(*node).SetStyle(style)
	}
}

// PropColor returns a PropertyFunction that sets the foreground color of the text.
func PropColor(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		return content, style.Foreground(lipgloss.Color(value))
	}
}

// PropBackgroundColor returns a PropertyFunction that sets the background color of the text.
func PropBackgroundColor(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		return content, style.Background(lipgloss.Color(value))
	}
}

// PropFontWeight returns a PropertyFunction that sets the font weight.
// Currently, it only supports making text bold.
func PropFontWeight(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if value == "bold" {
			return content, style.Bold(true)
		}
		return content, style
	}
}

// PropTextTransform returns a PropertyFunction that transforms the text content.
// Supports uppercase, lowercase, and capitalize.
func PropTextTransform(value string) PropertyFunction {
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

// PropFontStyle returns a PropertyFunction that sets the font style.
// Supports italic and bold.
func PropFontStyle(value string) PropertyFunction {
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

// PropTextDecoration returns a PropertyFunction that sets text decoration.
// Supports underline and line-through.
func PropTextDecoration(value string) PropertyFunction {
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

// PropMargin returns a PropertyFunction that sets all margins.
// Accepts up to four space-separated values.
func PropMargin(value string) PropertyFunction {
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

// PropMarginLeft returns a PropertyFunction that sets the left margin.
func PropMarginLeft(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.MarginLeft(v)
		}
		return content, style
	}
}

// PropMarginRight returns a PropertyFunction that sets the right margin.
func PropMarginRight(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.MarginRight(v)
		}
		return content, style
	}
}

// PropMarginTop returns a PropertyFunction that sets the top margin.
func PropMarginTop(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.MarginTop(v)
		}
		return content, style
	}
}

// PropMarginBottom returns a PropertyFunction that sets the bottom margin.
func PropMarginBottom(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.MarginBottom(v)
		}
		return content, style
	}
}

// PropPadding returns a PropertyFunction that sets all padding.
// Accepts up to four space-separated values.
func PropPadding(value string) PropertyFunction {
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

// PropPaddingLeft returns a PropertyFunction that sets the left padding.
func PropPaddingLeft(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.PaddingLeft(v)
		}
		return content, style
	}
}

// PropPaddingRight returns a PropertyFunction that sets the right padding.
func PropPaddingRight(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.PaddingRight(v)
		}
		return content, style
	}
}

// PropPaddingTop returns a PropertyFunction that sets the top padding.
func PropPaddingTop(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return content, style.PaddingTop(v)
		}
		return content, style
	}
}

// PropPaddingBottom returns a PropertyFunction that sets the bottom padding.
func PropPaddingBottom(value string) PropertyFunction {
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

// PropBorder returns a PropertyFunction that sets all borders.
// Accepts style, color, and width arguments.
func PropBorder(value string) PropertyFunction {
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

// PropBorderTop returns a PropertyFunction that sets the top border.
// Accepts style, color, and width arguments.
func PropBorderTop(value string) PropertyFunction {
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

// PropBorderBottom returns a PropertyFunction that sets the bottom border.
// Accepts style, color, and width arguments.
func PropBorderBottom(value string) PropertyFunction {
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

// PropBorderLeft returns a PropertyFunction that sets the left border.
// Accepts style, color, and width arguments.
func PropBorderLeft(value string) PropertyFunction {
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

// PropBorderRight returns a PropertyFunction that sets the right border.
// Accepts style, color, and width arguments.
func PropBorderRight(value string) PropertyFunction {
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

// PropWidth returns a PropertyFunction that sets the width of the element.
func PropWidth(width string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		w, _ := strconv.Atoi(width)
		return content, style.Width(w)
	}
}

// PropHeight returns a PropertyFunction that sets the height of the element.
func PropHeight(height string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		h, _ := strconv.Atoi(height)
		return content, style.Height(h)
	}
}

// PropTextAlign returns a PropertyFunction that sets the horizontal text alignment.
// Supports left, center, and right.
func PropTextAlign(value string) PropertyFunction {
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

// PropVerticalAlign returns a PropertyFunction that sets the vertical alignment.
// Supports top, center, and bottom.
func PropVerticalAlign(value string) PropertyFunction {
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

// PropIndent returns a PropertyFunction that sets the text indentation.
func PropIndent(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		indent, _ := strconv.Atoi(value)
		return content, style.MarginLeft(style.GetMarginLeft() + indent)
	}
}

// PropWordSpacing returns a PropertyFunction that sets the spacing between words.
func PropWordSpacing(value string) PropertyFunction {
	return func(content string, style lipgloss.Style) (string, lipgloss.Style) {
		spacing, _ := strconv.Atoi(value)
		content = strings.Join(strings.Split(content, " "), strings.Repeat(" ", spacing))
		return content, style
	}
}
