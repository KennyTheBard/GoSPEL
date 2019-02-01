package utils

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
    Force the value val in [min, max] for int32.
*/
func Clamp(min, max, val int32) int32 {
    if val <= min {
        return min
    } else if val >= max {
        return max
    } else {
        return val
    }
}

/**
    Force the value val in [min, max] for uint32.
*/
func Uclamp(min, max, val uint32) uint32 {
    if val <= min {
        return min
    } else if val >= max {
        return max
    } else {
        return val
    }
}


/**
    Returns if a point is in the given rectangle.
*/
func In_rectangle(p image.Point, rect image.Rectangle) bool {
    if p.X < rect.Min.X {
        return false
    }

    if p.X > rect.Max.X {
        return false
    }

    if p.Y < rect.Min.Y {
        return false
    }

    if p.Y > rect.Max.Y {
        return false
    }

    return true
}


/**
    Returns if 2 rectangles are the same.
*/
func Equal_rectangles(r1, r2 image.Rectangle) bool {
    if r1.Min.X == r2.Min.X && r1.Min.Y == r2.Min.Y && r1.Max.X == r2.Max.X && r1.Max.Y == r2.Max.Y {
        return true
    } else {
        return false
    }
}
