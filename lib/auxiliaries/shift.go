package auxiliaries

import(
    "image"
    "image/draw"
    "image/color"
)

/**
    Returns a copy of the given image with the bounds
    shifted with shift_factor.
*/
func Shift(img image.Image, shift_factor image.Point) (image.Image) {
    // prepare the image to be returned
    bounds := img.Bounds()
    new_bounds := image.Rect(bounds.Min.X + shift_factor.X, bounds.Min.Y + shift_factor.Y, bounds.Max.X + shift_factor.X, bounds.Max.Y + shift_factor.Y)
    ret := image.Image(image.NewRGBA(new_bounds))

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()

                    ret.(draw.Image).Set(x + shift_factor.X, y + shift_factor.Y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    return ret

}
