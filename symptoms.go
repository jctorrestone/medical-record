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
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var symptomList []*ListItem

func runSymptoms(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var wlistSymptoms widget.List
	var wlistSelected widget.List
	wlistSymptoms.Axis = layout.Vertical
	var wedtSearch widget.Editor
	var clkSearch widget.Clickable
	icon, _ := widget.NewIcon(icons.ActionSearch)
	var clkAdd widget.Clickable
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

			mlistSymptoms := material.List(th, &wlistSymptoms)
			mlistSelected := material.List(th, &wlistSelected)
			medtSearch := material.Editor(th, &wedtSearch, "¿Que síntoma?")
			iconSearch := material.IconButton(th, &clkSearch, icon, "Búsqueda")
			btnAdd := material.Button(th, &clkAdd, "Añadir")
			btnAccept := material.Button(th, &clkAccept, "Aceptar")

			if clkSearch.Clicked() {
				searchValue := wedtSearch.Text()
				symptoms, err := getSymptomsByDesc(searchValue)
				if err != nil {
					log.Fatal(err)
				}
				buttonList = nil
				for i := range symptoms {
					buttonList = append(buttonList, &ListItem{
						Id:   symptoms[i].ID,
						Text: symptoms[i].Description,
					})
				}
			}

			if clkAdd.Clicked() {
				w.Perform(system.ActionClose)
			}

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
								return mlistSymptoms.Layout(gtx, len(buttonList), listItems(th, buttonList, sendSymptom))
							},
						)
					},
				),
				layout.Flexed(0.4,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Vertical,
							Spacing: layout.SpaceEnd,
						}.Layout(gtx,
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{
										Axis:      layout.Horizontal,
										Alignment: layout.Middle,
									}.Layout(gtx,
										layout.Flexed(0.7,
											func(gtx layout.Context) layout.Dimensions {
												return marginFlex.Layout(gtx,
													func(gtx layout.Context) layout.Dimensions {
														return borderEditor.Layout(gtx, medtSearch.Layout)
													},
												)
											},
										),
										layout.Flexed(0.3,
											func(gtx layout.Context) layout.Dimensions {
												return marginFlex.Layout(gtx, iconSearch.Layout)
											},
										),
									)
								},
							),
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return mlistSelected.Layout(gtx, len(symptomList), listItems(th, symptomList, removeSymptom))
										},
									)
								},
							),
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, btnAdd.Layout)
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
