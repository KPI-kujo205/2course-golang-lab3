package painter

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"golang.org/x/exp/shiny/screen"
)

type testReceiver struct {
	lastTexture screen.Texture
}

func (tr *testReceiver) Update(t screen.Texture) {
	tr.lastTexture = t
}

type mockScreen struct{}

func (m mockScreen) NewBuffer(size image.Point) (screen.Buffer, error) {
	panic("implement me")
}

func (m mockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	return new(mockTexture), nil
}

func (m mockScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	panic("implement me")
}

type mockTexture struct {
	Colors []color.Color
}

func (m *mockTexture) Release() {}

func (m *mockTexture) Size() image.Point { return size }

func (m *mockTexture) Bounds() image.Rectangle {
	return image.Rectangle{Max: m.Size()}
}

func (m *mockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {}
func (m *mockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	m.Colors = append(m.Colors, src)
}

type Mock struct {
	mock.Mock
}

func (mockOperation *Mock) Do(t screen.Texture) bool {
	args := mockOperation.Called(t)
	return args.Bool(0)
}

func TestLoop_Start(t *testing.T) {
	mockScreen := mockScreen{}

	var l Loop
	l.Receiver = &testReceiver{}
	l.Start(mockScreen)

	assert.NotNil(t, l.next)
	assert.NotNil(t, l.prev)
}

func TestLoop_Post(t *testing.T) {
	mockScreen := mockScreen{}

	var l Loop
	l.Receiver = &testReceiver{}
	l.Start(mockScreen)

	l.Post(OperationFunc(func(tx screen.Texture) {}))
	l.Post(OperationFunc(func(tx screen.Texture) {}))

	assert.False(t, l.mq.Empty())
}

func TestLoop_StopAndWait(t *testing.T) {
	mockScreen := mockScreen{}

	var l Loop
	l.Receiver = &testReceiver{}
	l.Start(mockScreen)

	l.StopAndWait()
	assert.True(t, l.mq.Empty())
}
