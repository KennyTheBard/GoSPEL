package utils

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
    r_aux := uint32(interp.LinearInterpolation(int32(px1.R), int32(px2.R), x))
    g_aux := uint32(interp.LinearInterpolation(int32(px1.G), int32(px2.G), x))
    b_aux := uint32(interp.LinearInterpolation(int32(px1.B), int32(px2.B), x))
    a_aux := uint32(interp.LinearInterpolation(int32(px1.A), int32(px2.A), x))

    return Pixel{r_aux, g_aux, b_aux, a_aux}
}


/**
    Interpolate the value of the one inside the
    square delimited by the 2 pixels.
*/
func Pixel_bilinear_interpolation(px11, px12, px21, px22 Pixel, x, y float64) (Pixel) {
    r_aux := uint32(interp.BilinearInterpolation(int32(px11.R), int32(px12.R), int32(px21.R), int32(px22.R), x, y))
    g_aux := uint32(interp.BilinearInterpolation(int32(px11.G), int32(px12.G), int32(px21.G), int32(px22.G), x, y))
    b_aux := uint32(interp.BilinearInterpolation(int32(px11.B), int32(px12.B), int32(px21.B), int32(px22.B), x, y))
    a_aux := uint32(interp.BilinearInterpolation(int32(px11.A), int32(px12.A), int32(px21.A), int32(px22.A), x, y))

return Pixel{r_aux, g_aux, b_aux, a_aux}
}
