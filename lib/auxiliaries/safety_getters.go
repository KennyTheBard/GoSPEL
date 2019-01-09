package auxiliaries

import (
    "image"
)


/**
    Returns the color of the pixel or the closest pixel's
    color if the desired pixel is outside of the image.
*/
func Safe_Get_Color(img image.Image, x, y int) (uint32, uint32, uint32, uint32) {
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


/**
    Force the value val in [min, max].
*/
func Clamp(min, max uint32, val uint32) uint32 {
    if val <= min {
        return min
    } else if val >= max {
        return max
    } else {
        return val
    }
}
