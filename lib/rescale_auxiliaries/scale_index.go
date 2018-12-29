package rescale_auxiliaries

import "math"

func Scale_index(i int, r float64) int {
    return int(math.Floor(float64(i) * r))
}
