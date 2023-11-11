package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pbreedt/ebiten/tiles2/images"
)

const (
	screenWidth  = 240
	screenHeight = 240
)

const (
	tileSize = 16
)

var (
	soilImg  *ebiten.Image
	grassImg *ebiten.Image
)

func init() {
	var err error
	soilImg, _, err = ebitenutil.NewImageFromFile("images/soil.png")
	if err != nil {
		log.Fatal(err)
	}

	// grassImg, _, err = ebitenutil.NewImageFromFile("grass.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	img, _, err := image.Decode(bytes.NewReader(images.Grass_Tile))
	if err != nil {
		log.Fatal(err)
	}
	grassImg = ebiten.NewImageFromImage(img)
}

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// imgSize := grassImg.Bounds().Dx()

	for x := 0; x <= 15; x++ {
		for y := 0; y <= 15; y++ {
			op := &ebiten.DrawImageOptions{}
			sx := float64(x * tileSize)
			sy := float64(y * tileSize)
			op.GeoM.Translate(sx, sy)
			if (x*y)%4 == 0 {
				screen.DrawImage(grassImg, op)
			} else {
				screen.DrawImage(soilImg, op)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Tiles (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
