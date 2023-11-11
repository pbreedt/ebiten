package images

import (
	_ "embed"
)

var (
	//go:embed grass.png
	Grass_Tile []byte

	//go:embed soil.png
	Soil_Tile []byte
)
