package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

var borderEditor = widget.Border{
	Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
	CornerRadius: unit.Dp(2),
	Width:        unit.Dp(1.5),
}

var marginList = layout.Inset{
	Bottom: unit.Dp(10),
}

var marginFlex = layout.Inset{
	Top:    unit.Dp(10),
	Bottom: unit.Dp(10),
	Left:   unit.Dp(20),
	Right:  unit.Dp(20),
}
