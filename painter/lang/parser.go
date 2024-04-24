package lang

import (
	"bufio"
	"errors"
	"fmt"
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

// CommandFunc представляє функцію, яка обробляє команду.
type CommandFunc func(p *Parser, args []string) error

// CommandMap містить відображення команд на функції, які їх обробляють.
var CommandMap = map[string]CommandFunc{
	"white":  handleWhite,
	"green":  handleGreen,
	"update": handleUpdate,
	"bgrect": handleBgRect,
	"figure": handleFigure,
	"move":   handleMove,
	"reset":  handleReset,
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
		return errors.New("empty command")
	}

	command := fields[0]
	args := fields[1:]

	if cmdFunc, ok := CommandMap[command]; ok {
		return cmdFunc(p, args)
	}

	return errors.New("unknown command: " + command)
}

func handleWhite(p *Parser, args []string) error {
	if len(args) != 0 {
		return errors.New("white command requires 0 arguments")
	}
	p.painter.SetWhiteBg()
	return nil
}

func handleGreen(p *Parser, args []string) error {
	if len(args) != 0 {
		return errors.New("green command requires 0 arguments")
	}
	p.painter.SetGreenBg()
	return nil
}

func handleUpdate(p *Parser, args []string) error {
	if len(args) != 0 {
		return errors.New("update command requires 0 arguments")
	}
	p.painter.Update()
	return nil
}

func handleBgRect(p *Parser, args []string) error {
	if len(args) != 4 {
		return errors.New("invalid number of arguments for bgrect command")
	}
	x1, err := convertToAbsoluteCoords(args[0])
	if err != nil {
		return err
	}
	y1, err := convertToAbsoluteCoords(args[1])
	if err != nil {
		return err
	}
	x2, err := convertToAbsoluteCoords(args[2])
	if err != nil {
		return err
	}
	y2, err := convertToAbsoluteCoords(args[3])
	if err != nil {
		return err
	}
	p.painter.SetBgRectangle(image.Point{X: x1, Y: y1}, image.Point{X: x2, Y: y2})
	return nil
}

func handleFigure(p *Parser, args []string) error {
	if len(args) != 2 {
		return errors.New("invalid number of arguments for figure command")
	}
	x, err := convertToAbsoluteCoords(args[0])
	if err != nil {
		return err
	}
	y, err := convertToAbsoluteCoords(args[1])
	if err != nil {
		return err
	}
	p.painter.DrawTFigure(image.Point{X: x, Y: y})
	return nil
}

func handleMove(p *Parser, args []string) error {
	if len(args) != 2 {
		return errors.New("invalid number of arguments for move command")
	}
	x, err := convertToAbsoluteCoords(args[0])
	if err != nil {
		return err
	}
	y, err := convertToAbsoluteCoords(args[1])
	if err != nil {
		return err
	}
	p.painter.MoveTFigures(x, y)
	return nil
}

func handleReset(p *Parser, args []string) error {
	if len(args) != 0 {
		return errors.New("reset command requires 0 arguments")
	}
	p.painter.ResetPainterState()
	p.painter.ResetBg()
	return nil
}

func convertToAbsoluteCoords(s string) (int, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot convert these coordinates: %s", s)
	}

	return int(f * 800), nil
}
