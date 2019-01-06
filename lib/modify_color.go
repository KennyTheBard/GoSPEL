package lib

import (
    "image"
    "image/draw"
    "image/color"
    "math"
)

/**
    Force the value val in [min, max].
*/
func Clamp(min, max uint32, val uint32) uint32 {
    if val <= min {
        return min
    } else if val >= max {
        return max
    } else {
        return val
    }
}

func Modify_colors(img image.Image, mat [4][4]float64) {
    bounds := img.Bounds()

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {

                     r, g, b, a := img.At(x, y).RGBA()

                     rf := float64(r)
                     gf := float64(g)
                     bf := float64(b)
                     af := float64(a)

                     r_fin := Clamp(0, (256 << 8) - 1, uint32(math.Floor(rf * mat[0][0] + gf * mat[0][1] + bf * mat[0][2] + af * mat[0][3] )))
                     g_fin := Clamp(0, (256 << 8) - 1, uint32(math.Floor(rf * mat[1][0] + gf * mat[1][1] + bf * mat[1][2] + af * mat[1][3] )))
                     b_fin := Clamp(0, (256 << 8) - 1, uint32(math.Floor(rf * mat[2][0] + gf * mat[2][1] + bf * mat[2][2] + af * mat[2][3] )))
                     a_fin := Clamp(0, (256 << 8) - 1, uint32(math.Floor(rf * mat[3][0] + gf * mat[3][1] + bf * mat[3][2] + af * mat[3][3] )))

                     img.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }
}
