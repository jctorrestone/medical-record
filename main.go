package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

func main() {
	connect()
	go func() {
		w := app.NewWindow(
			app.Title("Historias clínicas virtuales"),
		)
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var wlistRecord widget.List
	wlistRecord.Axis = layout.Vertical
	var wedtSearch widget.Editor
	wedtSearch.SingleLine = true
	var clkSearch widget.Clickable
	var clkAdd widget.Clickable

	var searchValue string

	records, err := getRecords()
	if err != nil {
		log.Fatal(err)
	}

	var buttonList []*ListItem
	for i := range records {
		buttonList = append(buttonList, &ListItem{
			Id:   records[i].ID,
			Text: records[i].Date + ", " + records[i].PatientOBj.Lastname + " " + records[i].PatientOBj.Name,
		})
	}

	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			medtSearch := material.Editor(th, &wedtSearch, "Ingrese el nombre")
			btnAdd := material.Button(th, &clkAdd, "Añadir historia")

			if clkSearch.Clicked() {
				searchValue = wedtSearch.Text()
				log.Println(searchValue)
				records, err := getRecordsByPatient(searchValue)
				if err != nil {
					log.Fatal(err)
				}
				buttonList = nil
				for i := range records {
					buttonList = append(buttonList, &ListItem{
						Id:   records[i].ID,
						Text: records[i].Date + ", " + records[i].PatientOBj.Lastname + " " + records[i].PatientOBj.Name,
					})
				}
			}

			if clkAdd.Clicked() {
				go func() {
					w := app.NewWindow(
						app.Title("Nueva historia clínica"),
						app.Size(unit.Dp(750), unit.Dp(740)),
					)
					err := runAddRecord(w)
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
								return listItems(gtx, th, wlistRecord, buttonList, sendMedicalRecord) //cambia la funcion
							},
							//layout.Spacer{Width: unit.Dp(25)}.Layout,
						)
					},
				),
				layout.Flexed(0.4,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis: layout.Vertical,
						}.Layout(gtx,
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{
										Axis: layout.Horizontal,
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
												return marginFlex.Layout(gtx,
													func(gtx layout.Context) layout.Dimensions {
														ic, err := widget.NewIcon(icons.ActionSearch)
														if err != nil {
															log.Fatal(err)
														}
														icon := material.IconButton(th, &clkSearch, ic, "Búsqueda")
														return icon.Layout(gtx)
													},
												)
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

func sendMedicalRecord(patientID int64) {

}
