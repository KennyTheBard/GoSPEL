package interpolation

import "math"

/**
    Return the interpoled value in a point between the 2
    v1 - value in first point
    v2 - value in second point
    proc - distance between the first point and
    the interest point raported to the whole distance
*/
func Linear_interpolation(v1 uint32, v2 uint32, proc float64) uint32 {
    return uint32(math.Floor((float64(v1) * (1 - proc)) + (float64(v2) * proc)))
}
