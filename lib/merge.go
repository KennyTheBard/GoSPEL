package lib

import(
    "image"
    "image/draw"
    "image/color"
    aux "./auxiliaries"
)

/**
    Merge the over image with the trg image
    over the assigned area of the target image.
*/
func Merge(trg, over image.Image, area image.Rectangle) (image.Image) {
    // prepare the image to be returned
    trg_bounds := trg.Bounds()
    ret := image.Image(image.NewRGBA(trg_bounds))

    // prepare the over image to be merged with target image
    img := Rescale(over, area)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := trg_bounds.Min.Y + rank; y <= trg_bounds.Max.Y; y += n {
                for x := trg_bounds.Min.X; x <= trg_bounds.Max.X; x++ {

                    if y >= area.Min.Y && y <= area.Max.Y && x >= area.Min.X && x <= area.Max.X {
                        r1, g1, b1, a1 := trg.At(x, y).RGBA()
                        r2, g2, b2, a2 := img.At(x, y).RGBA()

                        px1 := aux.Pixel{r1, g1, b1, a1}
                        px2 := aux.Pixel{r2, g2, b2, a2}

                        proc := float64(a2) / float64((256 << 8) - 1)

                        fin := aux.Pixel_linear_interpolation(px1, px2, proc)

                        ret.(draw.Image).Set(x, y, color.RGBA{uint8(fin.R >> 8), uint8(fin.G >> 8), uint8(fin.B >> 8), uint8(255)})

                    } else {
                        r, g, b, a := trg.At(x, y).RGBA()
                        ret.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
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
