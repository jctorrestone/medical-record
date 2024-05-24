package main

import (
	"log"
	"strconv"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var treatments []Treatment

func runTreatments(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var wlistTreatment widget.List
	wlistTreatment.Axis = layout.Vertical
	var clkAdd widget.Clickable
	var clkAccept widget.Clickable

	var buttonList []*ListItem

	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			mlistTreatment := material.List(th, &wlistTreatment)
			btnAdd := material.Button(th, &clkAdd, "Añadir")
			btnAccept := material.Button(th, &clkAccept, "Aceptar")

			if len(treatments) != len(buttonList) {
				buttonList = nil
				for i := range treatments {
					buttonList = append(buttonList, &ListItem{
						Text: treatments[i].MedicineObj.Name + ", " + strconv.FormatInt(treatments[i].Quantity, 10),
					})
				}
			}

			if clkAdd.Clicked() {
				go func() {
					w := app.NewWindow(
						app.Title("Seleccionar medicina"),
					)
					err := runAddTreatment(w)
					if err != nil {
						log.Fatal(err)
					}
				}()
			}

			if clkAccept.Clicked() {
				w.Perform(system.ActionClose)
			}

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Flexed(0.8,
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return mlistTreatment.Layout(gtx, len(buttonList), listItems(th, buttonList, sendTreatment))
							},
						)
					},
				),
				layout.Flexed(0.1,
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, btnAdd.Layout)
					},
				),
				layout.Flexed(0.1,
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, btnAccept.Layout)
					},
				),
			)

			e.Frame(gtx.Ops)
		}
	}

	return nil
}

func sendTreatment(treatmentID int64) {

}
