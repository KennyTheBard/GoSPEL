package lib

import(
    "image"
    "image/draw"
    "image/color"
    aux "./auxiliaries"
)

/**
    Returns an image with the opacity ampified by
*/
func Modify_opacity(img image.Image, alpha_proc float64) (image.Image) {
    // prepare the image to be returned
    bounds := img.Bounds()
    ret := image.Image(image.NewRGBA(bounds))

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()

                    a_fin := aux.Clamp(0, (256 << 8) - 1, uint32(alpha_proc * float64(a)))

                    ret.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a_fin >> 8)})
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
