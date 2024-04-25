package lang

import (
	"image"
	"strings"
	"testing"

	"github.com/KPI-kujo205/2course-golang-lab3/painter"
	"github.com/stretchr/testify/assert"
)

func TestParser_figure_operations(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedRes painter.Operation
	}{
		{
			name:        "Test update command",
			input:       "update",
			expectedRes: painter.UpdateOp,
		},
		{
			name:        "Test bgrect command",
			input:       "bgrect 0.1 0.2 0.3 0.4",
			expectedRes: &painter.BgRectangle{TopLeftPoint: image.Point{X: 80, Y: 160}, BottomRightPoint: image.Point{X: 240, Y: 320}},
		},
		{
			name:        "Test figure command",
			input:       "figure 0.5 0.6",
			expectedRes: &painter.NewTFigure{CentralPoint: image.Point{X: 400, Y: 480}},
		},
		{
			name:        "Test move command",
			input:       "move 0.7 0.8",
			expectedRes: &painter.MoveNewTFigure{OffsetX: 560, OffsetY: 640},
		},
		{
			name:        "Test invalid command",
			input:       "invalid",
			expectedRes: nil,
		},
		{
			name:        "Test too many arguments",
			input:       "figure 0.5 0.6 0,7",
			expectedRes: nil,
		},
	}

	for _, test := range tests {
		p := &Parser{}
		t.Run(test.name, func(t *testing.T) {
			res, err := p.Parse(strings.NewReader(test.input))
			if test.expectedRes == nil {
				assert.Error(t, err)
			} else {
				if err != nil {
					t.Errorf("Expected %v, got: %v", test.expectedRes, res)
				}
				assert.Equal(t, res[0], test.expectedRes)
			}
		})
	}
}

func TestParse_colors_and_reset(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedRes painter.Operation
	}{
		{
			name:        "Test white command",
			input:       "white",
			expectedRes: painter.OperationFunc(painter.WhiteFill),
		},
		{
			name:        "Test green command",
			input:       "green",
			expectedRes: painter.OperationFunc(painter.GreenFill),
		},
		{
			name:        "Test reset command",
			input:       "reset",
			expectedRes: painter.OperationFunc(painter.Reset),
		},
		{
			name:        "Test bad command",
			input:       "reset white",
			expectedRes: nil,
		},
	}

	for _, test := range tests {
		p := &Parser{}
		t.Run(test.name, func(t *testing.T) {
			res, err := p.Parse(strings.NewReader(test.input))
			if test.expectedRes == nil {
				assert.Error(t, err)
			} else {
				if err != nil {
					t.Error(err)
				}
				assert.IsType(t, res[0], test.expectedRes)
			}
		})
	}
}
