package lib

import (
    "image"
    "image/draw"
    "image/color"
    interp "./interpolation"
    aux "./rescale_auxiliaries"
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

                     r_aux1 := interp.Linear_interpolation(r11, r12, 0.5)
                     g_aux1 := interp.Linear_interpolation(g11, g12, 0.5)
                     b_aux1 := interp.Linear_interpolation(b11, b12, 0.5)
                     a_aux1 := interp.Linear_interpolation(a11, a12, 0.5)

                     r_aux2 := interp.Linear_interpolation(r21, r22, 0.5)
                     g_aux2 := interp.Linear_interpolation(g21, g22, 0.5)
                     b_aux2 := interp.Linear_interpolation(b21, b22, 0.5)
                     a_aux2 := interp.Linear_interpolation(a21, a22, 0.5)

                     r_fin := interp.Linear_interpolation(r_aux1, r_aux2, 0.5)
                     g_fin := interp.Linear_interpolation(g_aux1, g_aux2, 0.5)
                     b_fin := interp.Linear_interpolation(b_aux1, b_aux2, 0.5)
                     a_fin := interp.Linear_interpolation(a_aux1, a_aux2, 0.5)

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
