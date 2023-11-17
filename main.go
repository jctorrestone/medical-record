package main

import (
	"image/color"
	"log"
	"os"
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

func main() {
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
	var storyList widget.List
	storyList.Axis = layout.Vertical
	var searchEditor widget.Editor
	var searchButton widget.Clickable
	var startButton widget.Clickable

	var searchValue string

	searchEditor.SingleLine = true

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			if searchButton.Clicked() {
				searchValue = searchEditor.Text()
				log.Println(searchValue)
			}

			if startButton.Clicked() {
				go func() {
					w2 := app.NewWindow(
						app.Title("Nuevo paciente"),
					)
					err := run2(w2)
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
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Left:   unit.Dp(35),
						}
						return margins.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								list := material.List(th, &storyList)
								return list.Layout(gtx, 15, func(gtx layout.Context, index int) layout.Dimensions {
									var buttonsample widget.Clickable
									button := material.Button(th, &buttonsample, "paciente "+strconv.Itoa(index))
									return layout.Inset{
										Bottom: unit.Dp(15),
									}.Layout(gtx, button.Layout)
								})
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
												margins := layout.Inset{
													Top:    unit.Dp(35),
													Bottom: unit.Dp(25),
													Left:   unit.Dp(35),
													Right:  unit.Dp(10),
												}

												border := widget.Border{
													Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
													CornerRadius: unit.Dp(2),
													Width:        unit.Dp(1),
												}

												return margins.Layout(gtx,
													func(gtx layout.Context) layout.Dimensions {
														return border.Layout(gtx,
															func(gtx layout.Context) layout.Dimensions {
																edit := material.Editor(th, &searchEditor, "Ingrese el nombre")
																return edit.Layout(gtx)
															},
														)
													},
												)
											},
										),
										layout.Flexed(0.3,
											func(gtx layout.Context) layout.Dimensions {
												margins := layout.Inset{
													Top:    unit.Dp(25),
													Bottom: unit.Dp(25),
													Left:   unit.Dp(10),
													Right:  unit.Dp(35),
												}
												return margins.Layout(gtx,
													func(gtx layout.Context) layout.Dimensions {
														ic, err := widget.NewIcon(icons.ActionSearch)
														if err != nil {
															log.Fatal(err)
														}
														icon := material.IconButton(th, &searchButton, ic, "Búsqueda")
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
									margins := layout.Inset{
										Top:    unit.Dp(25),
										Bottom: unit.Dp(25),
										Left:   unit.Dp(35),
										Right:  unit.Dp(35),
									}
									return margins.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											btn := material.Button(th, &startButton, "Añadir historia")
											//maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
											//title.Color = maroon
											//title.Alignment = text.Middle
											return btn.Layout(gtx)
										},
									)
								},
							),
						)
					},
				),
			)

			e.Frame(gtx.Ops)
		}
	}
}
