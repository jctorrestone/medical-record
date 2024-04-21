package main

import (
	"log"
	"math"
	"time"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func run2(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var clkPatient widget.Clickable
	var edtAge widget.Editor
	edtAge.SingleLine = true
	var edtWeight widget.Editor
	edtWeight.SingleLine = true
	var edtHeight widget.Editor
	edtWeight.SingleLine = true
	var wedtDuration widget.Editor
	wedtDuration.SingleLine = true
	var wedtSymptoms widget.Editor
	var wedtIDX widget.Editor
	var examValues [11]widget.Bool
	var wedtTreatment widget.Editor
	var clkRecord widget.Clickable

	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			lblFiliation := material.Label(th, unit.Sp(16), "Filiación")
			lblFiliation.Alignment = text.Start
			lblFiliation.Font.Weight = font.Bold
			lblPatient := material.Label(th, unit.Sp(15), "---")
			btnPatient := material.Button(th, &clkPatient, "Seleccionar paciente")
			currentDate := time.Now().Format("02-01-2006")
			lblDate := material.Label(th, unit.Sp(15), currentDate)
			lblDate.Alignment = text.End
			lblAge := material.Label(th, unit.Sp(15), "Edad")
			medtAge := material.Editor(th, &edtAge, "")
			lblWeight := material.Label(th, unit.Sp(15), "Peso")
			medtWeight := material.Editor(th, &edtWeight, "")
			lblHeight := material.Label(th, unit.Sp(15), "Talla")
			medtHeight := material.Editor(th, &edtHeight, "")
			lblStory := material.Label(th, unit.Sp(16), "Historia de la enfermedad")
			lblStory.Alignment = text.Start
			lblStory.Font.Weight = font.Bold
			lblDuration := material.Label(th, unit.Sp(15), "Tiempo de enfermedad")
			medtDuration := material.Editor(th, &wedtDuration, "")
			lblSymptoms := material.Label(th, unit.Sp(15), "Síntomas principales")
			medtSymptoms := material.Editor(th, &wedtSymptoms, "")
			lblIDX := material.Label(th, unit.Sp(15), "IDX")
			medtIDX := material.Editor(th, &wedtIDX, "")
			lblExams := material.Label(th, unit.Sp(15), "Examenes solicitados")
			lblExams.Alignment = text.Start
			var chkExams [11]material.CheckBoxStyle
			chkExams[0] = material.CheckBox(th, &examValues[0], "Hg completo")
			chkExams[1] = material.CheckBox(th, &examValues[1], "G-V-C")
			chkExams[2] = material.CheckBox(th, &examValues[2], "Perfil hepático")
			chkExams[3] = material.CheckBox(th, &examValues[3], "Rx tórax")
			chkExams[4] = material.CheckBox(th, &examValues[4], "TEM de tórax y mediastino")
			chkExams[5] = material.CheckBox(th, &examValues[5], "Dímero D")
			chkExams[6] = material.CheckBox(th, &examValues[6], "PCR")
			chkExams[7] = material.CheckBox(th, &examValues[7], "Ferritina")
			chkExams[8] = material.CheckBox(th, &examValues[8], "TP y TPT")
			chkExams[9] = material.CheckBox(th, &examValues[9], "Procalcitonina")
			chkExams[10] = material.CheckBox(th, &examValues[10], "Fibrinógeno")
			lblTreatment := material.Label(th, unit.Sp(15), "Tratamiento")
			medtTreatment := material.Editor(th, &wedtTreatment, "")
			btnRecord := material.Button(th, &clkRecord, "Añadir")

			examFlex := layoutExams(gtx, chkExams)

			if clkPatient.Clicked() {
				go func() {
					w4 := app.NewWindow(
						app.Title("Seleccionar paciente"),
					)
					err := run4(w4)
					if err != nil {
						log.Fatal(err)
					}

				}()
			}

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, lblFiliation.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(gtx,
							layout.Flexed(1,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, lblPatient.Layout)
								},
							),
							layout.Flexed(1,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, btnPatient.Layout)
								},
							),
							layout.Flexed(1,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, lblDate.Layout)
								},
							),
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(gtx,
							layout.Flexed(1,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, lblAge.Layout)
								},
							),
							layout.Flexed(1,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return borderEditor.Layout(gtx, medtAge.Layout)
										},
									)
								},
							),
							layout.Flexed(1,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, lblWeight.Layout)
								},
							),
							layout.Flexed(1,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return borderEditor.Layout(gtx, medtWeight.Layout)
										},
									)
								},
							),
							layout.Flexed(1,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, lblHeight.Layout)
								},
							),
							layout.Flexed(1,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return borderEditor.Layout(gtx, medtHeight.Layout)
										},
									)
								},
							),
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, lblStory.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(gtx,
							layout.Flexed(0.3,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx, lblDuration.Layout)
								},
							),
							layout.Flexed(0.7,
								func(gtx layout.Context) layout.Dimensions {
									return marginFlex.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return borderEditor.Layout(gtx, medtDuration.Layout)
										},
									)
								},
							),
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, lblSymptoms.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return borderEditor.Layout(gtx, medtSymptoms.Layout)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, lblIDX.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return borderEditor.Layout(gtx, medtIDX.Layout)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, lblExams.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis: layout.Vertical,
								}.Layout(gtx, examFlex...)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, lblTreatment.Layout)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								return borderEditor.Layout(gtx, medtTreatment.Layout)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return marginFlex.Layout(gtx, btnRecord.Layout)
					},
				),
			)

			e.Frame(gtx.Ops)
		}
	}
	return nil
}

func layoutExams(gtx layout.Context, chkExams [11]material.CheckBoxStyle) []layout.FlexChild {
	parentsLen := float64(len(chkExams)) / 2
	parentsLen = math.Ceil(parentsLen)
	var parents = make([]layout.FlexChild, int(parentsLen))

	for i := range parents {
		first := i * 2
		second := i*2 + 1
		if i*2+1 < len(chkExams) {
			parents[i] = layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx,
						layout.Flexed(0.5, chkExams[first].Layout),
						layout.Flexed(0.5, chkExams[second].Layout),
					)
				},
			)
		} else {
			parents[i] = layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx,
						layout.Flexed(0.5, chkExams[first].Layout),
					)
				},
			)
		}
	}

	return parents
}
