package utils

import "math"


/**
    Returns the index i scaled with ratio r.
*/
func Scale_index(i int, r float64) int {
    return int(math.Floor(float64(i) * r))
}
