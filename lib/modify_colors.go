package lib

import (
    "image"
    "image/draw"
    "image/color"
    "math"
    utils "./utils"
)

type Modifier struct {
  Mat [4][4]float64
  Constant [4]float64
}

/**
    Modify each pixel of the image img with
    the equation given on its coresponding line
    in the mat 2D array. Can be used to amplify
    or minimize color, or create a grayscale
    with any channel ratio desired.
*/
func ModifyColors(img image.Image, m Modifier) (image.Image){
    bounds := img.Bounds()
    ret := Copy(img)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y < bounds.Max.Y; y += n {
                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()

                    rf := float64(r)
                    gf := float64(g)
                    bf := float64(b)
                    af := float64(a)

                    // calculate the color after color modification
                    r_fin := utils.Uclamp(0, (256 << 8) - 1, uint32(math.Floor(rf * m.Mat[0][0] + gf * m.Mat[0][1] + bf * m.Mat[0][2] + af * m.Mat[0][3] + m.Constant[0] )))
                    g_fin := utils.Uclamp(0, (256 << 8) - 1, uint32(math.Floor(rf * m.Mat[1][0] + gf * m.Mat[1][1] + bf * m.Mat[1][2] + af * m.Mat[1][3] + m.Constant[1] )))
                    b_fin := utils.Uclamp(0, (256 << 8) - 1, uint32(math.Floor(rf * m.Mat[2][0] + gf * m.Mat[2][1] + bf * m.Mat[2][2] + af * m.Mat[2][3] + m.Constant[2] )))
                    a_fin := utils.Uclamp(0, (256 << 8) - 1, uint32(math.Floor(rf * m.Mat[3][0] + gf * m.Mat[3][1] + bf * m.Mat[3][2] + af * m.Mat[3][3] + m.Constant[3] )))

                    ret.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
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
