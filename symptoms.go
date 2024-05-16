package main

import (
	"log"
	"slices"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var symptomList []*ListItem

func runSymptoms(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var wlistSymptoms widget.List
	var wlistSelected widget.List
	wlistSymptoms.Axis = layout.Vertical
	var clkAccept widget.Clickable

	symptoms, err := getSymptomsByDesc("")
	if err != nil {
		log.Fatal(err)
	}

	var buttonList []*ListItem
	for i := range symptoms {
		buttonList = append(buttonList, &ListItem{
			Id:   symptoms[i].ID,
			Text: symptoms[i].Description,
		})
	}

	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			btnAccept := material.Button(th, &clkAccept, "Aceptar")

			if clkAccept.Clicked() {
				w.Perform(system.ActionClose)
			}

			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx,
				layout.Flexed(0.6,
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return listItems(gtx, th, wlistSymptoms, buttonList, sendSymptom)
							},
						)
					},
				),
				layout.Flexed(0.4,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Vertical,
							Spacing: layout.SpaceStart,
						}.Layout(gtx,
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return listItems(gtx, th, wlistSelected, symptomList, removeSymptom)
										},
									)
								},
							),
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, btnAccept.Layout)
								},
							),
						)
					},
				),
			)

			e.Frame(gtx.Ops)
		}
	}
	return nil
}

func sendSymptom(symptomID int64) {
	for i := range symptomList {
		if symptomList[i].Id == symptomID {
			return
		}
	}
	symptom, err := getSymptomById(symptomID)
	if err != nil {
		log.Fatal(err)
	}
	symptomList = append(symptomList, &ListItem{
		Id:   symptom.ID,
		Text: symptom.Description,
	})
}

func removeSymptom(symptomID int64) {
	for i := range symptomList {
		if symptomList[i].Id == symptomID {
			symptomList = slices.Delete(symptomList, i, i+1)
			break
		}
	}
}
