package lib

import (
    "image"
    "image/color"
    "image/draw"
)

type Filter struct {
  Mat [3][3]float64
}

func Apply_filter(img image.Image, start image.Point, end image.Point, f Filter) {

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for x := start.X + rank; x <= end.X; x += n {
                for y := start.Y; y <= end.Y; y++ {
                    sum_r := float64(0)
                    sum_g := float64(0)
                    sum_b := float64(0)

                    for i := -1; i <= 1; i++ {
                        for j := -1; j <= 1; j++ {
                            // values are returned as uint16
                            r, g, b, _ := img.At(x + i, y + j).RGBA()

                          sum_r += float64(r) * f.Mat[i + 1][j + 1]
                          sum_g += float64(g) * f.Mat[i + 1][j + 1]
                          sum_b += float64(b) * f.Mat[i + 1][j + 1]
                        }
                    }

                    _, _, _, alpha := img.At(x, y).RGBA();
                    img.(draw.Image).Set(x, y, color.RGBA{uint8(uint64(sum_r) >> 8), uint8(uint64(sum_g) >> 8), uint8(uint64(sum_b) >> 8), uint8(alpha >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }
}
