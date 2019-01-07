package lib

import (
    "image"
    "math"
    "fmt"
)

func Rotate(img image.Image, angle float64) image.Image {
    // calculate radians
    angle = -angle;
    rad := angle * math.Pi / 180

    // calculate shearing factors
    xsf := -math.Tan(rad/2)
    ysf := math.Sin(rad)

    fmt.Println("Grade =", -angle)
    fmt.Println("xsf =", xsf)
    fmt.Println("ysf =", ysf)

    // apply shearing
    ret := Shear(img, xsf, XSHEAR)
    ret = Shear(ret, ysf, YSHEAR)
    ret = Shear(ret, xsf, XSHEAR)

    return ret
}
