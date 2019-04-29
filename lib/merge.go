package lib

import(
    "image"
    "image/draw"
    "image/color"
    utils "./utils"
)

/**
    Merge the over image with the trg image starting
    at the anchor point in the target image.
*/
func Merge(trg, over image.Image, anchor image.Point) (image.Image) {
    // prepare the image to be returned
    trg_bounds := trg.Bounds()
    over_bounds := over.Bounds()
    ret := Copy(trg)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank
            MIN := func (a, b int) int {
                if a > b {
                    return b
                }
                return a
            }
            MAX := func (a, b int) int {
                if a < b {
                    return b
                }
                return a
            }

            start_y := MAX(trg_bounds.Min.Y, over_bounds.Min.Y + anchor.Y)
            end_y := MIN(trg_bounds.Max.Y, over_bounds.Max.Y + anchor.Y)
            start_x := MAX(trg_bounds.Min.X, over_bounds.Min.X + anchor.X)
            end_x := MIN(trg_bounds.Max.X, over_bounds.Max.X + anchor.X)

            for y := start_y + rank; y <= end_y; y += n {
                for x := start_x; x <= end_x; x ++ {
                    r1, g1, b1, a1 := trg.At(x, y).RGBA()
                    r2, g2, b2, a2 := over.At(x - anchor.X, y - anchor.Y).RGBA()

                    px1 := utils.Pixel{r1, g1, b1, a1}
                    px2 := utils.Pixel{r2, g2, b2, a2}

                    proc := float64(a2) / float64((256 << 8) - 1)

                    fin := utils.Pixel_linear_interpolation(px1, px2, proc)

                    ret.(draw.Image).Set(x, y, color.RGBA{uint8(fin.R >> 8), uint8(fin.G >> 8), uint8(fin.B >> 8), uint8(255)})

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
