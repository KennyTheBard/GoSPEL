package lib

import (
    "image"
    "image/color"
    "image/draw"
    utils "./utils"
)

type Filter struct {
  Mat [][]float64
}

/**
    Apply the filter f on each pixel of the image img in the assigned area.
*/
func ApplyFilter(img image.Image, f Filter) (image.Image) {
    bounds := img.Bounds()
    trg := CreateImage(bounds, color.RGBA{0, 0, 0, 0})

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
                            r, g, b, a := utils.Safe_Get_Color(img, x + j, y + i)

                            sum_r += float64(r) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                            sum_g += float64(g) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                            sum_b += float64(b) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                            sum_a += float64(a) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                        }
                    }

                    fin_r := int32(sum_r)
                    fin_g := int32(sum_g)
                    fin_b := int32(sum_b)
                    fin_a := int32(sum_a)

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
