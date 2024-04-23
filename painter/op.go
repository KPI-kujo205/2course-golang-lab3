package painter

import "C"
import (
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"image"
	"image/color"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

type BgRectangle struct {
	TopLeftPoint     image.Point
	BottomRightPoint image.Point
}

type NewTFigure struct {
	CentralPoint image.Point
}

type MoveNewTFigure struct {
	OffsetX int
	OffsetY int
	Targets []*NewTFigure
}

func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src)
}

func (bgrect *BgRectangle) Do(t screen.Texture) bool {
	t.Fill(
		image.Rect(bgrect.TopLeftPoint.X, bgrect.TopLeftPoint.Y, bgrect.BottomRightPoint.X, bgrect.BottomRightPoint.Y),
		color.Black, screen.Src,
	)
	return false
}

func (movtfigure *MoveNewTFigure) Do(t screen.Texture) bool {
	for i := range movtfigure.Targets {
		movtfigure.Targets[i].CentralPoint.X += movtfigure.OffsetX
		movtfigure.Targets[i].CentralPoint.Y += movtfigure.OffsetY
	}

	return false
}

func (ntfigure *NewTFigure) Do(t screen.Texture) bool {
	figureColor := color.RGBA{R: 255, G: 255, A: 1}
	t.Fill(
		image.Rect(
			ntfigure.CentralPoint.X-200,
			ntfigure.CentralPoint.Y-100,
			ntfigure.CentralPoint.X+200,
			ntfigure.CentralPoint.Y,
		),
		figureColor, draw.Src,
	)
	t.Fill(
		image.Rect(
			ntfigure.CentralPoint.X-50,
			ntfigure.CentralPoint.Y-100,
			ntfigure.CentralPoint.X+50,
			ntfigure.CentralPoint.Y+200,
		),
		figureColor, draw.Src,
	)
	return false
}
