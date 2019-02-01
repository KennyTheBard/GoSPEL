package interpolation

import "math"

/**
    Return the interpoled value in a point between the 2
    v1 - value in first point
    v2 - value in second point
    x - [-1, 2]
*/
func CubicInterpolation(p0, p1, p2, p3 int32, x float64) int32 {
    p0f := float64(p0)
    p1f := float64(p1)
    p2f := float64(p2)
    p3f := float64(p3)

    x3 := float64(-1 * p0f / 2 + 3 * p1f / 2 - 3 * p2f / 2 + p3f / 2) * math.Pow(x, 3)
    x2 := float64(p0f - 5 * p1f / 2 + 2 * p2f - p3f / 2) * math.Pow(x, 2)
    x1 := float64(-1 * p0f / 2 + p2f / 2) * x
    x0 := p1f

    return int32(math.Floor(x3 + x2 + x1 + x0))
}
