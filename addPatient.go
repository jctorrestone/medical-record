package main

import (
	"log"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func run3(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var wedtName, wedtLname widget.Editor
	wedtName.SingleLine = true
	wedtLname.SingleLine = true
	var clkAdd widget.Clickable
	var gender widget.Enum
	gender.Value = "1"

	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			lblName := material.Label(th, unit.Sp(15), "Nombres")
			medtName := material.Editor(th, &wedtName, "")
			lblLname := material.Label(th, unit.Sp(15), "Apellidos")
			medtLname := material.Editor(th, &wedtLname, "")
			lblGender := material.Label(th, unit.Sp(15), "Sexo")
			btnAdd := material.Button(th, &clkAdd, "Añadir")

			if clkAdd.Clicked() {
				pat := Patient{
					Name:     wedtName.Text(),
					Lastname: wedtLname.Text(),
					Gender:   gender.Value,
				}
				id, err := addPatient(pat)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("Se añadio: %d", id)
				w.Perform(system.ActionClose)
			}

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, lblName.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return borderEditor.Layout(gtx, medtName.Layout)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, lblLname.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return borderEditor.Layout(gtx, medtLname.Layout)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, lblGender.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, material.RadioButton(th, &gender, "1", "masculino").Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, material.RadioButton(th, &gender, "0", "femenino").Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, btnAdd.Layout)
					},
				),
			)

			e.Frame(gtx.Ops)
		}
	}
	return nil
}
