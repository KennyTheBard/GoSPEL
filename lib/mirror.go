package lib

import (
    "image"
    "image/draw"
    "image/color"
)

const (
    HORIZONTAL_MODE = 0
    VERTICAL_MODE = 1
)

type Reflection func (x, y int, bounds image.Rectangle) (int, int)

/**
    Returns the pixel mirrored with the given mirror function.
*/
func get_reflection(x, y int, bounds image.Rectangle, reflectionFunc Reflection) (int, int) {
    return reflectionFunc(x, y, bounds)
}

/**
    Returns the image mirrored in the given mode.
*/
func Mirror(img image.Image, mode int) (image.Image) {
    bounds := img.Bounds()
    ret := image.Image(image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)))

    var reflectionFunc Reflection

    switch mode {
        case HORIZONTAL_MODE:
            reflectionFunc = func (x, y int, bounds image.Rectangle) (int, int) {
                return x, bounds.Max.Y - (y - bounds.Min.Y)
            }
            break

        case VERTICAL_MODE:
            reflectionFunc = func (x, y int, bounds image.Rectangle) (int, int) {
                return bounds.Max.X - (x - bounds.Min.X), y
            }
            break

        default:
            reflectionFunc = func (x, y int, bounds image.Rectangle) (int, int) {
                return x, y
            }
    }

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    new_x, new_y := get_reflection(x, y, bounds, reflectionFunc)

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
