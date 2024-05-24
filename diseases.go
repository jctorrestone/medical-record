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

var diseaseList []*ListItem

func runDiseases(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var wlistDiseases widget.List
	var wlistSelected widget.List
	wlistDiseases.Axis = layout.Vertical
	var wedtSearch widget.Editor
	var clkSearch widget.Clickable
	icon, _ := widget.NewIcon(icons.ActionSearch)
	var clkAdd widget.Clickable
	var clkAccept widget.Clickable

	diseases, err := getDiseasesByDesc("")
	if err != nil {
		log.Fatal(err)
	}

	var buttonList []*ListItem
	for i := range diseases {
		buttonList = append(buttonList, &ListItem{
			Id:   diseases[i].ID,
			Text: diseases[i].Description,
		})
	}

	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			mlistDiseases := material.List(th, &wlistDiseases)
			mlistSelected := material.List(th, &wlistSelected)
			medtSearch := material.Editor(th, &wedtSearch, "¿Que enfermedad?")
			iconSearch := material.IconButton(th, &clkSearch, icon, "Búsqueda")
			btnAdd := material.Button(th, &clkAdd, "Añadir")
			btnAccept := material.Button(th, &clkAccept, "Aceptar")

			if clkSearch.Clicked() {
				searchValue := wedtSearch.Text()
				diseases, err := getDiseasesByDesc(searchValue)
				if err != nil {
					log.Fatal(err)
				}
				buttonList = nil
				for i := range diseases {
					buttonList = append(buttonList, &ListItem{
						Id:   diseases[i].ID,
						Text: diseases[i].Description,
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
								return mlistDiseases.Layout(gtx, len(buttonList), listItems(th, buttonList, sendDisease))
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
											return mlistSelected.Layout(gtx, len(diseaseList), listItems(th, diseaseList, removeDisease))
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

func sendDisease(diseaseID int64) {
	for i := range diseaseList {
		if diseaseList[i].Id == diseaseID {
			return
		}
	}
	disease, err := getDiseaseById(diseaseID)
	if err != nil {
		log.Fatal(err)
	}
	diseaseList = append(diseaseList, &ListItem{
		Id:   disease.ID,
		Text: disease.Description,
	})
}

func removeDisease(diseaseID int64) {
	for i := range diseaseList {
		if diseaseList[i].Id == diseaseID {
			diseaseList = slices.Delete(diseaseList, i, i+1)
			break
		}
	}
}
