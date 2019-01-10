package utils

import (
    "math"
    "image"
    "image/color"
)

func Rotate_point(p image.Point, rad float64) (float64, float64) {
    x := math.Cos(rad) * float64(p.X) + math.Sin(rad) * float64(p.Y)
    y := -math.Sin(rad) * float64(p.X) + math.Cos(rad) * float64(p.Y)

    return x, y
}

func Get_rotated_color(img image.Image, p image.Point, angle float64) (color.Color) {
    rad := Angle2Rad(angle)
    xf, yf := Rotate_point(p, rad)

    return img.At(int(math.Round(xf)), int(math.Round(yf)))
}

func Angle2Rad(angle float64) float64 {
    return angle * math.Pi / 180
}

func Rad2Angle(rad float64) float64 {
    return rad / math.Pi * 180
}

func Rectangle_center(rect image.Rectangle) (image.Point) {
    center_x := (rect.Max.X + rect.Min.X) / 2
    center_y := (rect.Max.Y + rect.Min.Y) / 2
    return image.Point{center_x, center_y}
}

func Rectangle_size(rect image.Rectangle) (int, int) {
    width := rect.Max.X - rect.Min.X
    height := rect.Max.Y - rect.Min.Y
    return width, height
}
