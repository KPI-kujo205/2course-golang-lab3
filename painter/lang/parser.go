package lang

import (
	"bufio"
	"errors"
	"image"
	"io"
	"strconv"
	"strings"

	"github.com/KPI-kujo205/2course-golang-lab3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	painter painter.StatePainter
}

// Parse парсить вхідний потік даних та повертає список операцій.
func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		commandLine := scanner.Text()
		err := p.parseCommand(commandLine)
		if err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	res := p.painter.GetOperationsList()

	return res, nil
}

// parseCommand парсить окрему команду та повертає операцію для неї.
func (p *Parser) parseCommand(commandLine string) error {
	fields := strings.Fields(commandLine)
	if len(fields) == 0 {
		_ = errors.New("empty command")
	}

	command := fields[0]
	args := fields[1:]

	switch command {
	case "white":
		p.painter.SetWhiteBg()
	case "green":
		p.painter.SetGreenBg()
	case "update":
		p.painter.Update()
	case "bgrect":
		if len(args) != 4 {
			return errors.New("invalid number of arguments for bgrect command")
		}
		x1, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		y1, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		x2, err := strconv.Atoi(args[2])
		if err != nil {
			return err
		}
		y2, err := strconv.Atoi(args[3])
		if err != nil {
			return err
		}
		p.painter.SetBgRectangle(image.Point{X: x1, Y: y1}, image.Point{X: x2, Y: y2})
		return nil
	case "figure":
		if len(args) != 2 {
			return errors.New("invalid number of arguments for figure command")
		}
		x, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return err
		}
		y, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			return err
		}
		p.painter.DrawTFigure(image.Point{X: int(x), Y: int(y)})
		return nil
	case "move":
		if len(args) != 2 {
			return errors.New("invalid number of arguments for move command")
		}
		x, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		p.painter.MoveTFigures(x, y)
		return nil
	case "reset":
		p.painter.ResetPainterState()
		p.painter.ResetBg()
		return nil
	default:
		return errors.New("unknown command: " + command)
	}
	return nil
}
