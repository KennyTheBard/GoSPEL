package lib

import (
    "fmt"
    "math"
    "image"
    "image/draw"
    "image/color"
    aux "./auxiliaries"
)

func Rotate(img image.Image, angle float64) image.Image {
    bounds := img.Bounds()
    width, height := Rectangle_size(bounds)

    // calculate angle in radians
    rad := angle * math.Pi / 180

    aux_image := aux.Shift(img, image.Point{-width/2, -height/2})
    bounds = aux_image.Bounds()
    width, height = Rectangle_size(bounds)

    // calculate new bounds
    min_x, _ := Rotate_point(image.Point{-width/2, -height/2}, rad)
    _, min_y := Rotate_point(image.Point{width/2, -height/2}, rad)
    max_x, _ := Rotate_point(image.Point{width/2, height/2}, rad)
    _, max_y := Rotate_point(image.Point{-width/2, height/2}, rad)

    ret_bounds := image.Rect(int(math.Floor(min_x)), int(math.Floor(min_y)), int(math.Floor(max_x)), int(math.Floor(max_y)))
    ret_width, ret_height := Rectangle_size(ret_bounds)
    ret_bounds = image.Rect(0, -ret_height, ret_width, ret_height)

    ret := image.Image(image.NewRGBA(ret_bounds))


    fmt.Println(ret_bounds)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := ret_bounds.Min.Y + rank; y <= ret_bounds.Max.Y; y += n {
                for x := ret_bounds.Min.X; x <= ret_bounds.Max.X; x++ {
                    xx := x //+ ret_width / 2
                    yy := y //- int(float64(ret_height) * math.Cos(Angle2Rad(angle)) / 2)

                    r11, g11, b11, a11 := Get_rotated_color(img, image.Point{xx, yy}, -angle).RGBA()
                    r12, g12, b12, a12 := Get_rotated_color(img, image.Point{xx + 1, yy}, -angle).RGBA()
                    r21, g21, b21, a21 := Get_rotated_color(img, image.Point{xx, yy + 1}, -angle).RGBA()
                    r22, g22, b22, a22 := Get_rotated_color(img, image.Point{xx + 1, yy + 1}, -angle).RGBA()

                    px11 := aux.Pixel{r11, g11, b11, a11}
                    px12 := aux.Pixel{r12, g12, b12, a12}
                    px21 := aux.Pixel{r21, g21, b21, a21}
                    px22 := aux.Pixel{r22, g22, b22, a22}

                    fin := aux.Pixel_bilinear_interpolation(px11, px12, px21, px22, 0.5, 0.5)

                    ret.(draw.Image).Set(x, y, color.RGBA{uint8(fin.R >> 8), uint8(fin.G >> 8), uint8(fin.B >> 8), uint8(fin.A >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    ret = Crop(ret, Select_opaque(ret))

    return ret
}

func Rotate_point(p image.Point, rad float64) (float64, float64) {

    x := math.Cos(rad) * float64(p.X) + math.Sin(rad) * float64(p.Y)
    y := -math.Sin(rad) * float64(p.X) + math.Cos(rad) * float64(p.Y)

    // fmt.Println(p, "fata de", c, "este", x, y)

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
