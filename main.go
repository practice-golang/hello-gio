package main // import "hello-gio"
import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	// "gioui.org/font/gofont"
	"giofont"
)

var colors = make(map[string]color.NRGBA)

func toRectF(r image.Rectangle) f32.Rectangle {
	return f32.Rectangle{
		Min: f32.Point{X: float32(r.Min.X), Y: float32(r.Min.Y)},
		Max: f32.Point{X: float32(r.Max.X), Y: float32(r.Max.Y)},
	}
}

func loop(w *app.Window) error {
	th := material.NewTheme(giofont.Collection())
	var ops op.Ops

	var startButton widget.Clickable
	var btnMSG string = "바꿔!"
	var headerMSG string = "안녕, 지오\nHello, Gio"

	spacerNarrow := layout.Rigid(layout.Spacer{Height: unit.Dp(4)}.Layout)
	spacerWide := layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			myRect := image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
			paint.FillShape(gtx.Ops, colors["white"], clip.Rect(myRect).Op())
			return layout.Spacer{Height: unit.Dp(32)}.Layout(gtx)
		},
	)
	titleHeader := layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			// paint.ColorOp{Color: colors["cream"]}.Add(gtx.Ops)
			// paint.PaintOp{}.Add(gtx.Ops)

			l := material.H1(th, headerMSG)
			l.Color = colors["maroon"]
			l.Alignment = text.Middle

			// paint.Fill(gtx.Ops, colors["cream"]) // 전체
			// myRect := image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Constraints.Max.Y) // 위쪽 스페이서 제외한 영역
			myRect := image.Rect(0, 0, l.Layout(gtx).Size.X, l.Layout(gtx).Size.X) // 가로세로
			paint.FillShape(gtx.Ops, colors["cream"], clip.Rect(myRect).Op())

			return l.Layout(gtx)
		},
	)
	buttonArea := layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			margins := layout.Inset{
				Top:    unit.Dp(25),
				Bottom: unit.Dp(25),
				Right:  unit.Dp(35),
				Left:   unit.Dp(35),
			}

			return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &startButton, btnMSG)
				return btn.Layout(gtx)
			})
		},
	)

	myLayout := []layout.FlexChild{spacerNarrow, titleHeader, spacerWide, buttonArea}

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			if startButton.Clicked() {
				if btnMSG == "바꿔!" {
					btnMSG = "갈굼 멈춰!"
					headerMSG = "안돼, 안 바꿔줘\n바꿀 생각 없어. 빨리 돌아가"
				} else {
					btnMSG = "바꿔!"
					headerMSG = "안녕, 지오\nHello, Gio"
				}
			}

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceBetween,
				// Spacing: layout.SpaceEnd,
			}.Layout(gtx, myLayout...)

			e.Frame(gtx.Ops)
		}
	}
}

func init() {
	colors["white"] = color.NRGBA{255, 255, 255, 255}
	colors["maroon"] = color.NRGBA{127, 0, 0, 255}
	colors["cream"] = color.NRGBA{240, 240, 127, 255}
	colors["black"] = color.NRGBA{0, 0, 0, 255}
}

func main() {
	go func() {
		w := app.NewWindow()
		w.Invalidate() // Prevent eat system.FrameEvent
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
