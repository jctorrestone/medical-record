package main

import (
	"log"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ListItem struct {
	Id    int64
	Text  string
	Click widget.Clickable
}

func run4(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var wlistPatient widget.List
	wlistPatient.Axis = layout.Vertical
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

			btnAdd := material.Button(th, &clkAdd, "Añadir")

			if clkAdd.Clicked() {
				go func() {
					w3 := app.NewWindow(
						app.Title("Añadir paciente"),
					)
					err := run3(w3)
					if err != nil {
						log.Fatal(err)
					}

				}()
			}

			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx,
				layout.Flexed(0.6,
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return listItems(gtx, th, wlistPatient, buttonList)
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

func listItems(gtx layout.Context, th *material.Theme, wlistItem widget.List, buttonList []*ListItem) layout.Dimensions {
	return material.List(th, &wlistItem).Layout(gtx, len(buttonList),
		func(gtx layout.Context, index int) layout.Dimensions {
			item := buttonList[index]
			for item.Click.Clicked() {
				log.Printf("My id is: %v\n", item.Id)
			}
			return marginList.Layout(gtx, material.Button(th, &item.Click, item.Text).Layout)
		},
	)
}
