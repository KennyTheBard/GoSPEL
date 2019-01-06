package auxiliaries

import (
    interp "../interpolation"
)

func Pixel_bilinear_interpolation(r11, g11, b11, a11, r12, g12, b12, a12, r21, g21, b21, a21, r22, g22, b22, a22 uint32, x, y float64) (uint32, uint32, uint32, uint32) {
    r_aux := interp.Bilinear_interpolation(r11, r12, r21, r22, x, y)
    g_aux := interp.Bilinear_interpolation(g11, g12, g21, g22, x, y)
    b_aux := interp.Bilinear_interpolation(b11, b12, b21, b22, x, y)
    a_aux := interp.Bilinear_interpolation(a11, a12, a21, a22, x, y)

    return r_aux, g_aux, b_aux, a_aux
}
