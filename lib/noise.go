package lib

import (
    "fmt"
    "math/rand"
    "time"
    "image"
    "image/draw"
    "image/color"
    utils "./utils"
    interp "./interpolation"
)

/**
    Apply the median filter on each pixel of the image img in
    the assigned area defined by the mask in the given radius.
*/
func Noise(img image.Image, mask image.Image, strength int) (image.Image) {
    bounds := img.Bounds()
    ret := Copy(img)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank
            seed := rand.NewSource(time.Now().UnixNano() + int64(rank))
            rand_gen := rand.New(seed)

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x ++ {
                    r, g, b, a := img.At(x, y).RGBA()

                    num := rand_gen.Intn(2 * strength) - strength
                    // fmt.Println(num)
                    fmt.Print()

                    r = uint32(utils.Clamp(0, (256 << 8) - 1, int32(r) + int32(num << 8)))
                    g = uint32(utils.Clamp(0, (256 << 8) - 1, int32(g) + int32(num << 8)))
                    b = uint32(utils.Clamp(0, (256 << 8) - 1, int32(b) + int32(num << 8)))
                    a = uint32(utils.Clamp(0, (256 << 8) - 1, int32(a) + int32(num << 8)))

                    r_aux, g_aux, b_aux, a_aux := ret.At(x, y).RGBA()
                    r_mask, g_mask, b_mask, a_mask := mask.At(x, y).RGBA()

                    // calculate the color modification through mask
                    fin_r := interp.Linear_interpolation(int32(r_aux), int32(r), float64(r_mask) / float64((256 << 8) - 1))
                    fin_g := interp.Linear_interpolation(int32(g_aux), int32(g), float64(g_mask) / float64((256 << 8) - 1))
                    fin_b := interp.Linear_interpolation(int32(b_aux), int32(b), float64(b_mask) / float64((256 << 8) - 1))
                    fin_a := interp.Linear_interpolation(int32(a_aux), int32(a), float64(a_mask) / float64((256 << 8) - 1))

                    ret.(draw.Image).Set(x, y, color.RGBA{uint8(fin_r >> 8), uint8(fin_g >> 8), uint8(fin_b >> 8), uint8(fin_a >> 8)})
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
