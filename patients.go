package main

import (
	"log"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type ListItem struct {
	Id    int64
	Text  string
	Click widget.Clickable
}

var patWindow *app.Window

func runPatients(w *app.Window) error {
	patWindow = w
	th := material.NewTheme()
	var ops op.Ops

	var wlistPatient widget.List
	wlistPatient.Axis = layout.Vertical
	var wedtSearch widget.Editor
	var clkSearch widget.Clickable
	icon, _ := widget.NewIcon(icons.ActionSearch)
	var clkAdd widget.Clickable

	patients, err := getPatients()
	if err != nil {
		log.Fatal(err)
	}

	var buttonList []*ListItem
	for i := range patients {
		buttonList = append(buttonList, &ListItem{
			Id:   patients[i].ID,
			Text: patients[i].Lastname + ", " + patients[i].Name,
		})
	}

	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			mlistPatient := material.List(th, &wlistPatient)
			medtSearch := material.Editor(th, &wedtSearch, "Ingrese el nombre")
			iconSearch := material.IconButton(th, &clkSearch, icon, "Búsqueda")
			btnAdd := material.Button(th, &clkAdd, "Añadir")

			if clkSearch.Clicked() {
				searchValue := wedtSearch.Text()
				patients, err := getPatientsByName(searchValue)
				if err != nil {
					log.Fatal(err)
				}
				buttonList = nil
				for i := range patients {
					buttonList = append(buttonList, &ListItem{
						Id:   patients[i].ID,
						Text: patients[i].Lastname + ", " + patients[i].Name,
					})
				}
			}

			if clkAdd.Clicked() {
				go func() {
					w := app.NewWindow(
						app.Title("Añadir paciente"),
					)
					err := runAddPatient(w)
					if err != nil {
						log.Fatal(err)
					}

				}()
				w.Perform(system.ActionClose)
			}

			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx,
				layout.Flexed(0.6,
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return mlistPatient.Layout(gtx, len(buttonList), listItems(th, buttonList, sendPatient))
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
									return marginFlex.Layout(gtx, btnAdd.Layout)
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

func sendPatient(patientID int64) {
	patient, err := getPatientById(patientID)
	if err != nil {
		log.Fatal(err)
	}
	patObj = patient
	patWindow.Perform(system.ActionClose)
}

func listItems(th *material.Theme, buttonList []*ListItem, itemClicked func(int64)) layout.ListElement {
	return func(gtx layout.Context, index int) layout.Dimensions {
		item := buttonList[index]
		for item.Click.Clicked() {
			log.Printf("My id is: %v\n", item.Id)
			itemClicked(item.Id)
		}
		return marginList.Layout(gtx, material.Button(th, &item.Click, item.Text).Layout)
	}
}
