package utils

import (
    "image"
)

/**
    Returns the values the image have to be shifted
with in order to have the image start from (0, 0).
*/
func Calculate_shift_factor(img image.Image) (image.Point) {
    return image.Point{0 - img.Bounds().Min.X, 0 - img.Bounds().Min.Y}
}
