package lib

import (
    "image"
    "image/draw"
    "image/color"
    utils "./utils"
)

func SelectColor(img image.Image, ref color.Color, threshold float64) (image.Image) {
    bounds := img.Bounds()
    ret := CreateImage(bounds, color.RGBA{0, 0, 0, 0})

    r_ref, g_ref, b_ref, _ := ref.RGBA()

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    r, g, b, _ := img.At(x, y).RGBA()

                    if float64(utils.Abs(int(r) - int(r_ref)) + utils.Abs(int(g) - int(g_ref)) + utils.Abs(int(b) - int(b_ref))) / float64((3 * 255) << 8)<= threshold {
                        ret.(draw.Image).Set(x, y, color.RGBA{255, 255, 255, 255})
                    } else {
                        ret.(draw.Image).Set(x, y, color.RGBA{0, 0, 0, 0})
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
