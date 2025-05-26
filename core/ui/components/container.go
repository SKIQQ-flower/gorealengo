package components

import "github.com/skwb/realengo-conflict/core/ui"

type Layout int

const (
	Vertical Layout = iota
	Horizontal
)

type Container struct {
	X, Y       int32
	Padding    int32
	Spacing    int32
	Children   []ui.UIElement
	Layout     Layout
	Background *NinePatch
}

func (c *Container) SetBackground(n *NinePatch) {
	c.Background = n
	c.recalculateLayout()
}

func (c *Container) Add(child ui.UIElement) {
	c.Children = append(c.Children, child)
	c.recalculateLayout()
}

func (c *Container) SetPosition(x, y int32) {
	c.X = x
	c.Y = y
	c.recalculateLayout()
}

func (c *Container) Draw() {
	if c.Background != nil {
		w, h := c.GetSize()
		c.Background.SetDest(c.X, c.Y, w, h)
		c.Background.Draw()
	}
	for _, child := range c.Children {
		child.Draw()
	}
}

func (c *Container) GetSize() (int32, int32) {
	var contentW, contentH int32
	for _, child := range c.Children {
		cw, ch := child.GetSize()
		if c.Layout == Vertical {
			contentH += ch + c.Spacing
			if cw > contentW {
				contentW = cw
			}
		} else {
			contentW += cw + c.Spacing
			if ch > contentH {
				contentH = ch
			}
		}
	}
	if len(c.Children) > 0 {
		if c.Layout == Vertical {
			contentH -= c.Spacing
		} else {
			contentW -= c.Spacing
		}
	}
	totalW := contentW + 2*c.Padding
	totalH := contentH + 2*c.Padding
	return totalW, totalH
}

func (c *Container) recalculateLayout() {
	cursorX := c.X + c.Padding
	cursorY := c.Y + c.Padding
	for _, child := range c.Children {
		child.SetPosition(cursorX, cursorY)
		cw, ch := child.GetSize()
		if c.Layout == Vertical {
			cursorY += ch + c.Spacing
		} else {
			cursorX += cw + c.Spacing
		}
	}
}
