package painter

import "image"

type StatePainter struct {
	bgColour     Operation
	bgRectangle  *BgRectangle
	tFigures     []*NewTFigure
	moveTFigures []Operation
	updateOp     Operation
}

func (sp *StatePainter) SetWhiteBg() {
	sp.bgColour = OperationFunc(WhiteFill)
}

func (sp *StatePainter) SetGreenBg() {
	sp.bgColour = OperationFunc(GreenFill)
}

func (sp *StatePainter) SetBgRectangle(topLeft, bottomRight image.Point) {
	sp.bgRectangle = &BgRectangle{
		TopLeftPoint:     topLeft,
		BottomRightPoint: bottomRight,
	}
}

func (sp *StatePainter) Update() {
	sp.updateOp = UpdateOp
}

func (sp *StatePainter) MoveTFigures(offsetX, offsetY int) {
	moveTFigure := &MoveNewTFigure{
		OffsetX: offsetX,
		OffsetY: offsetY,
		Targets: sp.tFigures,
	}
	sp.moveTFigures = append(sp.moveTFigures, moveTFigure)
}

func (sp *StatePainter) DrawTFigure(centralPoint image.Point) {
	figure := &NewTFigure{CentralPoint: centralPoint}
	sp.tFigures = append(sp.tFigures, figure)
}
