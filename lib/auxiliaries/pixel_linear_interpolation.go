package auxiliaries

import (
    interp "../interpolation"
)

func Pixel_linear_interpolation(r1, g1, b1, a1, r2, g2, b2, a2 uint32, x float64) (uint32, uint32, uint32, uint32) {
    r_aux := interp.Linear_interpolation(r1, r2, x)
    g_aux := interp.Linear_interpolation(g1, g2, x)
    b_aux := interp.Linear_interpolation(b1, b2, x)
    a_aux := interp.Linear_interpolation(a1, a2, x)

    return r_aux, g_aux, b_aux, a_aux
}
