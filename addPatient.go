package main

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func run2(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var nameEditor, lnameEditor widget.Editor
	nameEditor.SingleLine = true
	lnameEditor.SingleLine = true
	var add widget.Clickable
	var gender widget.Enum
	gender.Value = "m"

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			name := material.Label(th, unit.Sp(15), "Nombres")
			nameInput := material.Editor(th, &nameEditor, "")
			lname := material.Label(th, unit.Sp(15), "Apellidos")
			lnameInput := material.Editor(th, &lnameEditor, "")
			inputBorder := widget.Border{
				Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
				CornerRadius: unit.Dp(2),
				Width:        unit.Dp(1),
			}
			margins := layout.Inset{
				Top:    unit.Dp(10),
				Bottom: unit.Dp(10),
				Left:   unit.Dp(100),
				Right:  unit.Dp(100),
			}
			genderLabel := material.Label(th, unit.Sp(15), "Sexo")
			addButton := material.Button(th, &add, "Añadir")

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx, name.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return inputBorder.Layout(gtx, nameInput.Layout)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx, lname.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return inputBorder.Layout(gtx, lnameInput.Layout)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx, genderLabel.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx, material.RadioButton(th, &gender, "m", "masculino").Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx, material.RadioButton(th, &gender, "f", "femenino").Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx, addButton.Layout)
					},
				),
			)

			e.Frame(gtx.Ops)
		}
	}
}
