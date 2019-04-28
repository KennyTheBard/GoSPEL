package lib

import (
    "image"
    "image/color"
    "image/draw"
    utils "./utils"
    interp "./interpolation"
)

type Filter struct {
  Mat [][]float64
}

/**
    Apply the filter f on each pixel of the image img in the assigned area.
*/
func ApplyFilter(img image.Image, mask image.Image, f Filter) (image.Image) {
    bounds := img.Bounds()
    trg := Copy(img)
    aux_img := Copy(img)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    sum_r := float64(0)
                    sum_g := float64(0)
                    sum_b := float64(0)
                    sum_a := float64(0)

                    // calculate the color after filter appliance
                    for i := - len(f.Mat) / 2; i <= len(f.Mat) / 2 + len(f.Mat) % 2 - 1; i++ {
                        for j := - len(f.Mat[i + len(f.Mat) / 2]) / 2; j <= len(f.Mat[i + len(f.Mat) / 2]) / 2 + len(f.Mat[i + len(f.Mat) / 2]) % 2 - 1; j++ {
                            // values are returned as uint32
                            r, g, b, a := utils.Safe_Get_Color(aux_img, x + j, y + i)

                            sum_r += float64(r) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                            sum_g += float64(g) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                            sum_b += float64(b) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                            sum_a += float64(a) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                        }
                    }

                    r_aux, g_aux, b_aux, a_aux := aux_img.At(x, y).RGBA()
                    r_mask, g_mask, b_mask, a_mask := mask.At(x, y).RGBA()

                    // calculate the color modification through mask
                    fin_r := interp.LinearInterpolation(int32(r_aux), int32(sum_r), float64(r_mask) / float64((256 << 8) - 1))
                    fin_g := interp.LinearInterpolation(int32(g_aux), int32(sum_g), float64(g_mask) / float64((256 << 8) - 1))
                    fin_b := interp.LinearInterpolation(int32(b_aux), int32(sum_b), float64(b_mask) / float64((256 << 8) - 1))
                    fin_a := interp.LinearInterpolation(int32(a_aux), int32(sum_a), float64(a_mask) / float64((256 << 8) - 1))

                    trg.(draw.Image).Set(x, y, color.RGBA{uint8(fin_r >> 8), uint8(fin_g >> 8), uint8(fin_b >> 8), uint8(fin_a >> 8)})
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
