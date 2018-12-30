package lib

import (
    "image"
    "image/color"
    "image/draw"
)

type Filter struct {
  Mat [][]float64
}

func Safe_At(img image.Image, x, y int) (uint32, uint32, uint32, uint32) {
    img_bounds := img.Bounds()

    if x < img_bounds.Min.X {
        x = img_bounds.Min.X
    }

    if y < img_bounds.Min.Y {
        y = img_bounds.Min.Y
    }

    if x > img_bounds.Max.X {
        x = img_bounds.Max.X
    }

    if y > img_bounds.Max.Y {
        y = img_bounds.Max.Y
    }

    return img.At(x, y).RGBA()
}

func Apply_filter(img image.Image, start image.Point, end image.Point, f Filter, strength int) {

    aux_img := image.Image(image.NewRGBA(image.Rect(img.Bounds().Min.X, img.Bounds().Min.Y, img.Bounds().Max.X, img.Bounds().Max.X)))

    for i := 0; i < strength; i++ {
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
                        sum_a := float64(0)

                        for i := - len(f.Mat) / 2; i <= len(f.Mat) / 2 + len(f.Mat) % 2 - 1; i++ {
                            for j := - len(f.Mat[i + len(f.Mat) / 2]) / 2; j <= len(f.Mat[i + len(f.Mat) / 2]) / 2 + len(f.Mat[i + len(f.Mat) / 2]) % 2 - 1; j++ {
                                // values are returned as uint32
                                r, g, b, a := Safe_At(img, x + j, y + i)
                                //r, g, b, a := img.At(x + i, y + i).RGBA()

                                sum_r += float64(r) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                                sum_g += float64(g) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                                sum_b += float64(b) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                                sum_a += float64(a) * f.Mat[i + len(f.Mat) / 2][j + len(f.Mat[i + len(f.Mat) / 2]) / 2]
                            }
                        }

                        aux_img.(draw.Image).Set(x, y, color.RGBA{uint8(uint64(sum_r) >> 8), uint8(uint64(sum_g) >> 8), uint8(uint64(sum_b) >> 8), uint8(uint64(sum_a) >> 8)})
                    }
                }

                done <- true;
            } ()
        }

        for i := 0; i < n; i++ {
            <-done
        }

        aux_img, img = img, aux_img
    }
}
