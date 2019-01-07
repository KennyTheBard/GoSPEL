package lib

import(
    "fmt"
    "image"
    "image/draw"
    "image/color"
    "math"
    interp "./interpolation"
)

/**
    Returns an image sheared with the given shear factor.
*/
func Shear(img image.Image, shear float64) (image.Image) {
    bounds := img.Bounds()

    // calculate some shearing parameters
    shear_factor := uint32(math.Round(shear * float64(bounds.Max.X - bounds.Min.X) * 2))
    height := bounds.Max.Y - bounds.Min.Y
    fmt.Println(bounds.Max.X, "->", bounds.Max.X + int(shear_factor))

    // initialize the new image
    new_bounds := image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X + int(shear_factor), bounds.Max.Y)
    ret := image.Image(image.NewRGBA(new_bounds))
    //ret = Modify_opacity(img, 0)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                proc := float64(y) / float64(height)
                padding := int(interp.Linear_interpolation(0, shear_factor, proc))

                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()

                    ret.(draw.Image).Set(x + padding, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
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
