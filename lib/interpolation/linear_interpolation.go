package interpolation

import "math"

/**
    Return the interpoled value in a point between the 2
    v1 - value in first point
    v2 - value in second point
    x - [0, 1]
*/
func LinearInterpolation(v1, v2 int32, x float64) int32 {
    return int32(math.Floor((float64(v1) * (1 - x)) + (float64(v2) * x)))
}

func BilinearInterpolation(p11, p12, p21, p22 int32, x, y float64) int32 {
    p1 := LinearInterpolation(p11, p12, x)
    p2 := LinearInterpolation(p21, p22, x)

    p := LinearInterpolation(p1, p2, y)

    return p
}
