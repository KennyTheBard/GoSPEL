package lib

import (
    "image"
    "image/draw"
    "image/color"
    aux "./auxiliaries"
)

func Rescale(orig image.Image, ret image.Image) {
    orig_bounds := orig.Bounds()
    ret_bounds := ret.Bounds()

    // ratio return to original
    height_ratio := float64(orig_bounds.Max.Y - orig_bounds.Min.Y) / float64(ret_bounds.Max.Y - ret_bounds.Min.Y)
    width_ratio := float64(orig_bounds.Max.X - orig_bounds.Min.X) / float64(ret_bounds.Max.X - ret_bounds.Min.X)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := ret_bounds.Min.Y + rank; y <= ret_bounds.Max.Y; y += n {
                for x := ret_bounds.Min.X; x <= ret_bounds.Max.X; x++ {

                     r11, g11, b11, a11 := orig.At(aux.Scale_index(x, width_ratio), aux.Scale_index(y, height_ratio)).RGBA()
                     r12, g12, b12, a12 := orig.At(aux.Scale_index(x + 1, width_ratio), aux.Scale_index(y, height_ratio)).RGBA()
                     r21, g21, b21, a21 := orig.At(aux.Scale_index(x, width_ratio), aux.Scale_index(y + 1, height_ratio)).RGBA()
                     r22, g22, b22, a22 := orig.At(aux.Scale_index(x + 1, width_ratio), aux.Scale_index(y + 1, height_ratio)).RGBA()

                     r_fin, g_fin, b_fin, a_fin := aux.Pixel_bilinear_interpolation(r11, g11, b11, a11, r12, g12, b12, a12, r21, g21, b21, a21, r22, g22, b22, a22, 0.5, 0.5)

                     ret.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

}
