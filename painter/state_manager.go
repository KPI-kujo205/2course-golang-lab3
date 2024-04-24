package painter

import "image"

type StatePainter struct {
	setBgColour  Operation
	bgRectangle  *BgRectangle
	tFigures     []*NewTFigure
	moveTFigures []Operation
	updateOp     Operation
}

func (sp *StatePainter) SetWhiteBg() {
	sp.setBgColour = OperationFunc(WhiteFill)
}

func (sp *StatePainter) SetGreenBg() {
	sp.setBgColour = OperationFunc(GreenFill)
}

func (sp *StatePainter) SetBgRectangle(topLeft, bottomRight image.Point) {
	sp.bgRectangle = &BgRectangle{
		TopLeftPoint:     topLeft,
		BottomRightPoint: bottomRight,
	}
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

func (sp *StatePainter) ResetPainterState() {
	sp.setBgColour = nil
	sp.bgRectangle = nil
	sp.tFigures = nil
	sp.moveTFigures = nil
	sp.updateOp = nil
}

func (sp *StatePainter) Update() {
	sp.updateOp = UpdateOp
}

func (sp *StatePainter) ResetBg() {
	sp.setBgColour = OperationFunc(Reset)
}

func (sp *StatePainter) ResetUpdateOp() {
	if sp.updateOp != nil {
		sp.updateOp = nil
	}

	if sp.setBgColour == nil {
		sp.ResetBg()
	}
}
