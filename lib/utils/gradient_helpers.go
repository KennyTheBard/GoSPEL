package utils

import (
    "math"
    "image"
    "image/color"
)


/**
    Returns the indexes of the elements
    closest to y in the ys array, or the
    same index if tha array has only 1
    element or y is contained in the array.

    The elements of the array are expected
    to be sorted in ascending order.
*/
func Search_interval(ys []int, y int) (int, int) {
    prev := 0

    for i := 1; i < len(ys); i++ {

        if ys[i] > y {
            return prev, i
        } else if ys[i] == y {
            return i, i
        }

        prev = i
    }


    return prev, prev
}


/**
    Returns the distance between 2 points.
*/
func Distance(a, b image.Point) float64 {
    x := float64(a.X - b.X)
    y := float64(a.Y - b.Y)
    return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}


/**
    Returns the sorted color points.
*/
func Sort_color_points(ys []int, vals []color.Color) {
    for i := 0; i < len(ys); i++ {
        for j := i + 1; j < len(ys); j++ {
            if ys[i] > ys [j] {
                aux_idx := ys[i]
                ys[i] = ys[j]
                ys[j] = aux_idx

                aux_col := vals[i]
                vals[i] = vals[j]
                vals[j] = aux_col
            }
        }
    }
}
