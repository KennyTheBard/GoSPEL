package lib

import (
    "image"
    "image/color"
    "image/draw"
)


func Create_image(dimension image.Rectangle, col color.Color) (image.Image) {
    img := image.NewRGBA(dimension)
    draw.Draw(img, img.Bounds(), &image.Uniform{col}, dimension.Min, draw.Src)
    return img
}
