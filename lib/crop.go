package lib

import (
    "image"
    "image/draw"
    "image/color"
)

/**
    Returns a copy of the required area of the given image.
*/
func Crop_image(img image.Image, area image.Rectangle) (image.Image) {
    // prepare the image to be returned
    bounds := img.Bounds()
    ret_bounds := image.Rect(0, 0, area.Max.X - area.Min.X, area.Max.Y - area.Min.Y)
    ret := image.Image(image.NewRGBA(ret_bounds))

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := area.Min.Y + rank; y <= area.Max.Y; y += n {
                for x := area.Min.X; x <= area.Max.X; x++ {
                    if y >= bounds.Min.Y && y <= bounds.Max.Y && x >= bounds.Min.X && x <= bounds.Max.X {
                        r, g, b, a := img.At(x, y).RGBA()

                        ret.(draw.Image).Set(x - area.Min.X, y - area.Min.Y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
                    }
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
