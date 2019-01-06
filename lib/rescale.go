package lib

import (
    "image"
    "image/draw"
    "image/color"
    aux "./auxiliaries"
)

/**
    Rescale the given image to perfectly fit the given area.
*/
func Rescale(orig image.Image, trg_bounds image.Rectangle) (image.Image) {
    orig_bounds := orig.Bounds()
    trg := image.Image(image.NewRGBA(trg_bounds))

    // ratio return to original
    height_ratio := float64(orig_bounds.Max.Y - orig_bounds.Min.Y) / float64(trg_bounds.Max.Y - trg_bounds.Min.Y)
    width_ratio := float64(orig_bounds.Max.X - orig_bounds.Min.X) / float64(trg_bounds.Max.X - trg_bounds.Min.X)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := trg_bounds.Min.Y + rank; y <= trg_bounds.Max.Y; y += n {
                for x := trg_bounds.Min.X; x <= trg_bounds.Max.X; x++ {

                     r11, g11, b11, a11 := orig.At(aux.Scale_index(x, width_ratio), aux.Scale_index(y, height_ratio)).RGBA()
                     r12, g12, b12, a12 := orig.At(aux.Scale_index(x + 1, width_ratio), aux.Scale_index(y, height_ratio)).RGBA()
                     r21, g21, b21, a21 := orig.At(aux.Scale_index(x, width_ratio), aux.Scale_index(y + 1, height_ratio)).RGBA()
                     r22, g22, b22, a22 := orig.At(aux.Scale_index(x + 1, width_ratio), aux.Scale_index(y + 1, height_ratio)).RGBA()

                     px11 := aux.Pixel{r11, g11, b11, a11}
                     px12 := aux.Pixel{r12, g12, b12, a12}
                     px21 := aux.Pixel{r21, g21, b21, a21}
                     px22 := aux.Pixel{r22, g22, b22, a22}

                     fin := aux.Pixel_bilinear_interpolation(px11, px12, px21, px22, 0.5, 0.5)

                     trg.(draw.Image).Set(x, y, color.RGBA{uint8(fin.R >> 8), uint8(fin.G >> 8), uint8(fin.B >> 8), uint8(fin.A >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    return trg
}
