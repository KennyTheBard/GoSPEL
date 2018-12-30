package filter_auxiliaries

import (
    "image"
)

func Safe_At(img image.Image, x, y int) (uint32, uint32, uint32, uint32) {
    img_bounds := img.Bounds()

    if x < img_bounds.Min.X {
        x = img_bounds.Min.X
    }

    if y < img_bounds.Min.Y {
        y = img_bounds.Min.Y
    }

    if x > img_bounds.Max.X {
        x = img_bounds.Max.X
    }

    if y > img_bounds.Max.Y {
        y = img_bounds.Max.Y
    }

    return img.At(x, y).RGBA()
}
