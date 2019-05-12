package modifiers

import (
    lib "../.."
)

func Grayscale(red_ratio, green_ratio, blue_ratio float64) (lib.Modifier) {
    total := red_ratio + green_ratio + blue_ratio;

    rr := red_ratio / total
    gr := green_ratio / total
    br := blue_ratio / total

    return lib.Modifier{
        [4][4]float64{
            [4]float64{rr, gr, br, 0},
            [4]float64{rr, gr, br, 0},
            [4]float64{rr, gr, br, 0},
            [4]float64{0, 0, 0, 1}},
        [4]float64{0, 0, 0, 0}}
}
