package lib

import (
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
func Median(img image.Image, mask image.Image, radius int) (image.Image) {
    bounds := img.Bounds()
    ret := Copy(img)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y < bounds.Max.Y; y += n {
                for x := bounds.Min.X; x < bounds.Max.X; x ++ {
                    sum_r := float64(0)
                    sum_g := float64(0)
                    sum_b := float64(0)
                    sum_a := float64(0)

                    // calculate the color after filter appliance
                    count := float64(0)
                    for i := - radius; i <= radius; i ++ {
                        for j := - radius; j <= radius; j ++ {
                            if (x != x + i || y != y + j) && utils.Inside_rectangle(image.Point{x + j, y + i}, img.Bounds()) {
                                count += 1
                                r, g, b, a := img.At(x + j, y + i).RGBA()

                                sum_r += float64(r)
                                sum_g += float64(g)
                                sum_b += float64(b)
                                sum_a += float64(a)
                            }
                        }
                    }

                    sum_r /= count
                    sum_g /= count
                    sum_b /= count
                    sum_a /= count

                    r_aux, g_aux, b_aux, a_aux := ret.At(x, y).RGBA()
                    r_mask, g_mask, b_mask, a_mask := mask.At(x, y).RGBA()

                    // calculate the color modification through mask
                    fin_r := interp.LinearInterpolation(int32(r_aux), int32(sum_r), float64(r_mask) / float64((256 << 8) - 1))
                    fin_g := interp.LinearInterpolation(int32(g_aux), int32(sum_g), float64(g_mask) / float64((256 << 8) - 1))
                    fin_b := interp.LinearInterpolation(int32(b_aux), int32(sum_b), float64(b_mask) / float64((256 << 8) - 1))
                    fin_a := interp.LinearInterpolation(int32(a_aux), int32(sum_a), float64(a_mask) / float64((256 << 8) - 1))

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
