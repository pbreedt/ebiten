package main

import (
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 300
	screenHeight = 300

	ballRadius                        = 15
	ballAccelerationConstant          = float64(0.0000000015)
	ballAccelerationSpeedUpMultiplier = float64(2)
	ballResistance                    = float64(0.975)
)

type Game struct {
	pressedKeys []ebiten.Key
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var (
	ballPositionX     = float64(screenWidth) / 2
	ballPositionY     = float64(screenHeight) / 2
	ballMovementX     = float64(0)
	ballMovementY     = float64(0)
	ballAccelerationX = float64(0)
	ballAccelerationY = float64(0)
	prevUpdateTime    = time.Now()
)

func (g *Game) Update() error {
	timeDelta := float64(time.Since(prevUpdateTime))
	prevUpdateTime = time.Now()

	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

	ballAccelerationX = 0
	ballAccelerationY = 0

	acc := ballAccelerationConstant

	for _, key := range g.pressedKeys {
		switch key.String() {
		case "Space":
			acc *= ballAccelerationSpeedUpMultiplier
		}
	}

	for _, key := range g.pressedKeys {
		switch key.String() {
		case "ArrowDown":
			ballAccelerationY = acc
		case "ArrowUp":
			ballAccelerationY = -acc
		case "ArrowRight":
			ballAccelerationX = acc
		case "ArrowLeft":
			ballAccelerationX = -acc
		}
	}

	ballMovementY += ballAccelerationY
	ballMovementX += ballAccelerationX

	ballMovementX *= ballResistance
	ballMovementY *= ballResistance

	ballPositionX += ballMovementX * timeDelta
	ballPositionY += ballMovementY * timeDelta

	const minX = ballRadius
	const minY = ballRadius
	const maxX = screenWidth - ballRadius
	const maxY = screenHeight - ballRadius

	if ballPositionX >= maxX || ballPositionX <= minX {
		if ballPositionX > maxX {
			ballPositionX = maxX
		} else if ballPositionX < minX {
			ballPositionX = minX
		}

		ballMovementX *= -1
	}

	if ballPositionY >= maxY || ballPositionY <= minY {
		if ballPositionY > maxY {
			ballPositionY = maxY
		} else if ballPositionY < minY {
			ballPositionY = minY
		}

		ballMovementY *= -1
	}

	return nil
}

var simpleShader *ebiten.Shader

func init() {
	var err error

	simpleShader, err = ebiten.NewShader([]byte(`
		package main

		func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
			return color
		}
	`))
	if err != nil {
		panic(err)
	}
}

func (g *Game) drawCircle(screen *ebiten.Image, x, y, radius float32, clr color.RGBA) {
	var path vector.Path

	path.MoveTo(x, y)
	path.Arc(x, y, radius, 0, math.Pi*2, vector.Clockwise)

	vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)

	redScaled := float32(clr.R) / 255
	greenScaled := float32(clr.G) / 255
	blueScaled := float32(clr.B) / 255
	alphaScaled := float32(clr.A) / 255

	for i := range vertices {
		v := &vertices[i]

		v.ColorR = redScaled
		v.ColorG = greenScaled
		v.ColorB = blueScaled
		v.ColorA = alphaScaled
	}

	screen.DrawTrianglesShader(vertices, indices, simpleShader, &ebiten.DrawTrianglesShaderOptions{
		FillRule: ebiten.EvenOdd,
	})
}

func (g *Game) Draw(screen *ebiten.Image) {
	purpleClr := color.RGBA{255, 0, 255, 255}

	g.drawCircle(
		screen,
		float32(ballPositionX),
		float32(ballPositionY),
		float32(ballRadius),
		purpleClr,
	)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("The Incredible Movable Ball")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
