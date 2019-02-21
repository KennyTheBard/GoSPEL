package lib

import (
    "image"
    "image/color"
    "image/draw"
)

/**
    Returns an image with the given size and the
    given color.
*/
func CreateImage(dimension image.Rectangle, col color.Color) (image.Image) {
    img := image.NewRGBA(dimension)
    draw.Draw(img, img.Bounds(), &image.Uniform{col}, dimension.Min, draw.Src)
    return img
}
