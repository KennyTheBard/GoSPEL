package lib

import (
    "image"
    "image/draw"
    "image/color"
)

type TransformFunction func (x, y int, bounds image.Rectangle) (int, int)

func Transform(img image.Image, transform TransformFunction) (image.Image) {
    bounds := img.Bounds()
    ret := image.Image(image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)))

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    new_x, new_y := transform(x, y, bounds)

                    r, g, b, a := img.At(x, y).RGBA()

                    ret.(draw.Image).Set(new_x, new_y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    return ret;
}
