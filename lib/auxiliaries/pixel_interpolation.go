package auxiliaries

import (
    interp "../interpolation"
)

type Pixel struct {
    R, G, B, A uint32
}

/**
    Interpolate the value of the one on the
    segment delimited by the  2 pixels.
*/
func Pixel_linear_interpolation(px1, px2 Pixel, x float64) (Pixel) {
    r_aux := interp.Linear_interpolation(px1.R, px2.R, x)
    g_aux := interp.Linear_interpolation(px1.G, px2.G, x)
    b_aux := interp.Linear_interpolation(px1.B, px2.B, x)
    a_aux := interp.Linear_interpolation(px1.A, px2.A, x)

    return Pixel{r_aux, g_aux, b_aux, a_aux}
}

/**
    Interpolate the value of the one inside the
    square delimited by the 2 pixels.
*/
func Pixel_bilinear_interpolation(px11, px12, px21, px22 Pixel, x, y float64) (Pixel) {
    r_aux := interp.Bilinear_interpolation(px11.R, px12.R, px21.R, px22.R, x, y)
    g_aux := interp.Bilinear_interpolation(px11.G, px12.G, px21.G, px22.G, x, y)
    b_aux := interp.Bilinear_interpolation(px11.B, px12.B, px21.B, px22.B, x, y)
    a_aux := interp.Bilinear_interpolation(px11.A, px12.A, px21.A, px22.A, x, y)

return Pixel{r_aux, g_aux, b_aux, a_aux}
}
