package generator

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/sunshineplan/imgconv"
)

type DrawIconOps struct {
	X int
	Y int

	Name   string
	Screen *image.RGBA
	Size   int
}

func DrawIcon(ops DrawIconOps) error {
	icon, err := getIconPng(ops)
	if err != nil {
		return err
	}

	// draw.Draw(ops.Screen, icon.Rect, icon, image.Point{}, draw.Over)
	draw.Draw(ops.Screen, image.Rect(
		ops.X, ops.Y,
		ops.Screen.Rect.Dx(), ops.Screen.Rect.Dy()),
		*icon, image.Point{},
		draw.Over,
	)

	return nil
}

func getIconPng(ops DrawIconOps) (*image.Image, error) {
	iconFile, err := os.Open("./assets/" + ops.Name + ".png")
	if err != nil {
		return nil, err
	}

	defer iconFile.Close()

	// Decode the image (from PNG to image.Image):
	src, _ := png.Decode(iconFile)
	dst := imgconv.Resize(src, &imgconv.ResizeOption{Width: ops.Size, Height: ops.Size})
	return &dst, nil
}
