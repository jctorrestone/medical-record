package main

import (
	"log"
	"strconv"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var medObj Medicine

func runAddTreatment(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var wlistMedicine widget.List
	wlistMedicine.Axis = layout.Vertical
	var wedtSearch widget.Editor
	var clkSearch widget.Clickable
	var clkAddMedicine widget.Clickable
	icon, _ := widget.NewIcon(icons.ActionSearch)
	var wedtQuantity widget.Editor
	wedtQuantity.SingleLine = true
	var wedtDosage widget.Editor
	wedtDosage.SingleLine = true
	var wedtFrequency widget.Editor
	wedtFrequency.SingleLine = true
	var wedtNote widget.Editor
	var clkAdd widget.Clickable
	lblMedicine := material.Label(th, unit.Sp(15), "---")

	medicines, err := getMedicinesByDesc("")
	if err != nil {
		log.Fatal(err)
	}

	var buttonList []*ListItem
	for i := range medicines {
		buttonList = append(buttonList, &ListItem{
			Id:   medicines[i].ID,
			Text: medicines[i].Name + " " + strconv.FormatInt(medicines[i].Dose, 10) + medicines[i].Unit,
		})
	}

	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			mlistMedicine := material.List(th, &wlistMedicine)
			medtSearch := material.Editor(th, &wedtSearch, "Ingrese búsqueda")
			iconSearch := material.IconButton(th, &clkSearch, icon, "Búsqueda")
			btnAddMedicine := material.Button(th, &clkAddMedicine, "Añadir medicina")
			lblQuantity := material.Label(th, unit.Sp(15), "Cantidad")
			lblDosage := material.Label(th, unit.Sp(15), "Dosis")
			lblFrequency := material.Label(th, unit.Sp(15), "Frecuencia")
			medtQuantity := material.Editor(th, &wedtQuantity, "")
			medtDosage := material.Editor(th, &wedtDosage, "")
			medtFrequency := material.Editor(th, &wedtFrequency, "Nro. horas")
			medtNote := material.Editor(th, &wedtNote, "Añada un comentario")
			btnAdd := material.Button(th, &clkAdd, "Añadir")

			if medObj.ID != 0 {
				lblMedicine.Text = medObj.Name
			}

			if clkSearch.Clicked() {
				searchValue := wedtSearch.Text()
				medicines, err := getMedicinesByDesc(searchValue)
				if err != nil {
					log.Fatal(err)
				}
				buttonList = nil
				for i := range medicines {
					buttonList = append(buttonList, &ListItem{
						Id:   medicines[i].ID,
						Text: medicines[i].Name + " " + strconv.FormatInt(medicines[i].Dose, 10) + medicines[i].Unit,
					})
				}
			}

			if clkAddMedicine.Clicked() {
				w.Perform(system.ActionClose)
			}

			if clkAdd.Clicked() {
				quantityNum, _ := strconv.ParseInt(wedtQuantity.Text(), 10, 64)
				dosageNum, _ := strconv.ParseFloat(wedtDosage.Text(), 64)
				frequencyNum, _ := strconv.ParseInt(wedtFrequency.Text(), 10, 64)

				treatment := Treatment{
					MedicineID:  medObj.ID,
					MedicineObj: medObj,
					Quantity:    quantityNum,
					Dosage:      dosageNum,
					Frequency:   frequencyNum,
					Note:        wedtNote.Text(),
				}
				treatments = append(treatments, treatment)
				w.Perform(system.ActionClose)
			}

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Flexed(0.6,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(gtx,
							layout.Flexed(0.6,
								func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{
										Axis: layout.Vertical,
									}.Layout(gtx,
										layout.Flexed(0.8,
											func(gtx layout.Context) layout.Dimensions {
												return marginFlex.Layout(gtx,
													func(gtx layout.Context) layout.Dimensions {
														return mlistMedicine.Layout(gtx, len(buttonList), listItems(th, buttonList, sendMedicine))
													},
												)
											},
										),
										layout.Flexed(0.2,
											func(gtx layout.Context) layout.Dimensions {
												return marginFlex.Layout(gtx, lblMedicine.Layout)
											},
										),
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
												return marginFlex.Layout(gtx, btnAddMedicine.Layout)
											},
										),
									)
								},
							),
						)
					},
				),
				layout.Flexed(0.1,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(gtx,
							layout.Flexed(0.2,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, lblQuantity.Layout)
								},
							),
							layout.Flexed(0.2,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return borderEditor.Layout(gtx, medtQuantity.Layout)
										},
									)
								},
							),
							layout.Flexed(0.2,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, lblDosage.Layout)
								},
							),
							layout.Flexed(0.2,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return borderEditor.Layout(gtx, medtDosage.Layout)
										},
									)
								},
							),
							layout.Flexed(0.2,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, lblFrequency.Layout)
								},
							),
							layout.Flexed(0.2,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return borderEditor.Layout(gtx, medtFrequency.Layout)
										},
									)
								},
							),
						)
					},
				),
				layout.Flexed(0.2,
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return borderEditor.Layout(gtx, medtNote.Layout)
							},
						)
					},
				),
				layout.Flexed(0.1,
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

func sendMedicine(medicineID int64) {
	medicine, err := getMedicineById(medicineID)
	if err != nil {
		log.Fatal(err)
	}
	medObj = medicine
}
