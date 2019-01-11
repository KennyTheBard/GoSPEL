package lib

import (
    "math"
    "image"
    "image/draw"
    "image/color"
    utils "./utils"
)

/**
    Returns the given image rotated with the given angle.
*/
func Rotate(img image.Image, angle float64) image.Image {
    // for angles outside of [0, 90]
    scaled_image := img
    scaled_angle := angle

    for scaled_angle < 0 {
        scaled_angle += 360
    }
    for scaled_angle > 90 {
        scaled_image = Rotate(scaled_image, 90)
        scaled_angle -= 90;
    }
    if scaled_angle == 0 {
        return scaled_image
    }


    bounds := scaled_image.Bounds()
    width, height := utils.Rectangle_size(bounds)

    // calculate angle in radians
    rad := scaled_angle * math.Pi / 180

    aux_image := utils.Shift(scaled_image, utils.Calculate_shift_factor(scaled_image))
    bounds = aux_image.Bounds()
    width, height = utils.Rectangle_size(bounds)

    // calculate new bounds
    min_x, _ := utils.Rotate_point(image.Point{-width/2, -height/2}, rad)
    _, min_y := utils.Rotate_point(image.Point{width/2, -height/2}, rad)
    max_x, _ := utils.Rotate_point(image.Point{width/2, height/2}, rad)
    _, max_y := utils.Rotate_point(image.Point{-width/2, height/2}, rad)

    ret_bounds := image.Rect(int(math.Floor(min_x)), int(math.Floor(min_y)), int(math.Floor(max_x)), int(math.Floor(max_y)))
    ret_width, ret_height := utils.Rectangle_size(ret_bounds)
    ret_bounds = image.Rect(0, -ret_height, ret_width, ret_height)

    ret := image.Image(image.NewRGBA(ret_bounds))

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

                    r11, g11, b11, a11 := utils.Get_rotated_color(scaled_image, image.Point{xx, yy}, -scaled_angle).RGBA()
                    r12, g12, b12, a12 := utils.Get_rotated_color(scaled_image, image.Point{xx + 1, yy}, -scaled_angle).RGBA()
                    r21, g21, b21, a21 := utils.Get_rotated_color(scaled_image, image.Point{xx, yy + 1}, -scaled_angle).RGBA()
                    r22, g22, b22, a22 := utils.Get_rotated_color(scaled_image, image.Point{xx + 1, yy + 1}, -scaled_angle).RGBA()

                    px11 := utils.Pixel{r11, g11, b11, a11}
                    px12 := utils.Pixel{r12, g12, b12, a12}
                    px21 := utils.Pixel{r21, g21, b21, a21}
                    px22 := utils.Pixel{r22, g22, b22, a22}

                    fin := utils.Pixel_bilinear_interpolation(px11, px12, px21, px22, 0.5, 0.5)

                    ret.(draw.Image).Set(x, y, color.RGBA{uint8(fin.R >> 8), uint8(fin.G >> 8), uint8(fin.B >> 8), uint8(fin.A >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    ret = Crop(ret, Select_opaque(ret, 0))
    ret = utils.Shift(ret, utils.Calculate_shift_factor(ret))
    return ret
}
