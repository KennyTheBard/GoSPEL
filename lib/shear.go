package lib

import(
    "image"
    "image/draw"
    "image/color"
    "math"
    interp "./interpolation"
    utils "./utils"
)

const (
    XSHEAR = 0
    YSHEAR = 1
)

func Shear(img image.Image, shear float64, mode int) (image.Image) {
    switch mode {
    case XSHEAR:
        return xshear(img, shear)
    case YSHEAR:
        return yshear(img, shear)
    default:
        return Copy(img)
    }
}


/**
    Returns an image sheared on X with the given shear factor.
*/
func xshear(img image.Image, shear float64) (image.Image) {
    bounds := img.Bounds()

    // calculate some shearing parameters
    shear_factor := int(math.Round(shear * float64(bounds.Max.X - bounds.Min.X) * 2))
    height := bounds.Max.Y - bounds.Min.Y + 1

    // initialize the new image
    new_bounds := image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X + utils.Abs(shear_factor), bounds.Max.Y)
    ret := image.Image(image.NewRGBA(new_bounds))
    ret = ScaleOpacity(ret, 0)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y < bounds.Max.Y; y += n {
                // calculate padding
                proc := float64(y - bounds.Min.Y) / float64(height)
                padding := int(interp.LERP(0, int32(shear_factor), proc))

                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()

                    ret.(draw.Image).Set(x + padding - utils.Min(shear_factor, 0), y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    return ret
}


/**
    Returns an image sheared on X with the given shear factor.
*/
func yshear(img image.Image, shear float64) (image.Image) {
    bounds := img.Bounds()

    // calculate some shearing parameters
    shear_factor := int(math.Round(shear * float64(bounds.Max.Y - bounds.Min.Y) * 2))
    width := bounds.Max.X - bounds.Min.X + 1

    // initialize the new image
    new_bounds := image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y + utils.Abs(shear_factor))
    ret := image.Image(image.NewRGBA(new_bounds))
    ret = ScaleOpacity(ret, 0)

    // calculate padding
    paddings := make([]int, width)
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
        proc := float64(x - bounds.Min.X) / float64(width)
        paddings[x - bounds.Min.X] = int(interp.LERP(0, int32(shear_factor), proc))
    }

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y < bounds.Max.Y; y += n {
                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()

                    ret.(draw.Image).Set(x, y + paddings[x - bounds.Min.X] - utils.Min(shear_factor, 0), color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    ret = Shift(ret, utils.Calculate_shift_factor(ret))
    return ret
}
