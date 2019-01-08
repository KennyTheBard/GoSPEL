package lib

import (
    "image"
    "image/color"
    "image/draw"
    aux "./auxiliaries"
)

type Filter struct {
  Mat [][]float64
}

/**
    Apply the filter f on each pixel of the image img in the assigned area.
*/
func Apply_filter(img image.Image, area image.Rectangle, f Filter, strength int) {
    trg := Copy(img)

    for i := 0; i < strength; i++ {
        n := 10
        done := make(chan bool, n)

        for p := 0; p < n; p ++ {
            aux_rank := p
            go func() {
                rank := aux_rank

                for y := area.Min.Y + rank; y <= area.Max.Y; y += n {
                    for x := area.Min.X; x <= area.Max.X; x++ {
                        sum_r := float64(0)
                        sum_g := float64(0)
                        sum_b := float64(0)
                        sum_a := float64(0)

                        for i := - len(f.Mat) / 2; i <= len(f.Mat) / 2 + len(f.Mat) % 2 - 1; i++ {
                            for j := - len(f.Mat[i + len(f.Mat) / 2]) / 2; j <= len(f.Mat[i + len(f.Mat) / 2]) / 2 + len(f.Mat[i + len(f.Mat) / 2]) % 2 - 1; j++ {
                                // values are returned as uint32
                                r, g, b, a := aux.Safe_Get_Color(img, x + j, y + i)

                                sum_r += float64(r) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                                sum_g += float64(g) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                                sum_b += float64(b) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                                sum_a += float64(a) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                            }
                        }

                        trg.(draw.Image).Set(x, y, color.RGBA{uint8(uint64(sum_r) >> 8), uint8(uint64(sum_g) >> 8), uint8(uint64(sum_b) >> 8), uint8(uint64(sum_a) >> 8)})
                    }
                }

                done <- true;
            } ()
        }

        for i := 0; i < n; i++ {
            <-done
        }

        trg, img = img, trg
    }
}
