package generator

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type DrawTextOps struct {
	X int
	Y int

	Screen *image.RGBA
	Size   int
	Text   string
	Color  color.Color
}

func DrawText(ops DrawTextOps) error {
	marginX := 10
	marginY := 10

	f, err := opentype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size: float64(ops.Size),
		DPI:  72,
	})

	if err != nil {
		return err
	}

	point := fixed.Point26_6{
		X: fixed.I(ops.X + marginX),
		Y: fixed.I(ops.Y + marginY),
	}

	d := &font.Drawer{
		Dst:  ops.Screen,
		Src:  image.NewUniform(ops.Color),
		Face: face,
		Dot:  point,
	}

	d.DrawString(ops.Text)
	return nil
}
