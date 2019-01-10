package lib

import (
    "image"
    "image/draw"
    "image/color"
    "math"
    utils "./utils"
    interp "./interpolation"
)

/**
    Modify each pixel of the image img with
    the equation given on its coresponding line
    in the mat 2D array. Can be used to amplify
    or minimize color, or create a grayscale
    with any channel ratio desired.
*/
func Modify_colors(img image.Image, mask image.Image, mat [4][5]float64) (image.Image){
    bounds := img.Bounds()
    trg := Copy(img)
    mask = Resize(mask, bounds)

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

                    // calculate the color after color modification
                    r_fin := utils.Clamp(0, (256 << 8) - 1, uint32(math.Floor(rf * mat[0][0] + gf * mat[0][1] + bf * mat[0][2] + af * mat[0][3] + mat[0][4] )))
                    g_fin := utils.Clamp(0, (256 << 8) - 1, uint32(math.Floor(rf * mat[1][0] + gf * mat[1][1] + bf * mat[1][2] + af * mat[1][3] + mat[1][4] )))
                    b_fin := utils.Clamp(0, (256 << 8) - 1, uint32(math.Floor(rf * mat[2][0] + gf * mat[2][1] + bf * mat[2][2] + af * mat[2][3] + mat[2][4] )))
                    a_fin := utils.Clamp(0, (256 << 8) - 1, uint32(math.Floor(rf * mat[3][0] + gf * mat[3][1] + bf * mat[3][2] + af * mat[3][3] + mat[3][4] )))

                    r_aux, g_aux, b_aux, a_aux := img.At(x, y).RGBA()
                    r_mask, g_mask, b_mask, a_mask := mask.At(x, y).RGBA()

                    // calculate the color modification through mask
                    r_fin = uint32(interp.Linear_interpolation(int32(r_aux), int32(r_fin), float64(r_mask) / float64((256 << 8) - 1)))
                    g_fin = uint32(interp.Linear_interpolation(int32(g_aux), int32(g_fin), float64(g_mask) / float64((256 << 8) - 1)))
                    b_fin = uint32(interp.Linear_interpolation(int32(b_aux), int32(b_fin), float64(b_mask) / float64((256 << 8) - 1)))
                    a_fin = uint32(interp.Linear_interpolation(int32(a_aux), int32(a_fin), float64(a_mask) / float64((256 << 8) - 1)))

                    trg.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    return trg
}
