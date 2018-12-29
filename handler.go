package main

import (
  "fmt"
  "os"
  "image"
  "image/draw"
  "image/jpeg"
  "image/color"
  "math"
)

type Filter struct {
  mat [3][3]float64
}

func apply_filter(img image.Image, start image.Point, end image.Point, f Filter) {

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

                          sum_r += float64(r) * f.mat[i + 1][j + 1]
                          sum_g += float64(g) * f.mat[i + 1][j + 1]
                          sum_b += float64(b) * f.mat[i + 1][j + 1]
                        }
                    }

                    _, _, _, alpha := img.At(x, y).RGBA();
                    img.(draw.Image).Set(x, y, color.RGBA{uint8(uint64(sum_r) >> 8), uint8(uint64(sum_g) >> 8), uint8(uint64(sum_b) >> 8), uint8(alpha >> 24)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }
}

func min(a, b uint32) uint32 {
    if a > b {
        return b
    } else {
        return a
    }
}

/**
    Return the interpoled value in a point between the 2
    v1 - value in first point
    v2 - value in second point
    proc - distance between the first point and
    the interest point raported to the whole distance
*/
func linear_interpolation(v1 uint32, v2 uint32, proc float64) uint32 {
    return uint32(math.Floor((float64(v1) * (1 - proc)) + (float64(v2) * proc)))
}

func scale_index(i int, r float64) int {
    return int(math.Floor(float64(i) * r))
}

func resize(orig image.Image, ret image.Image) {
    orig_bounds := orig.Bounds()
    ret_bounds := ret.Bounds()

    // ratio return to original
    height_ratio := float64(orig_bounds.Max.Y - orig_bounds.Min.Y) / float64(ret_bounds.Max.Y - ret_bounds.Min.Y)
    width_ratio := float64(orig_bounds.Max.X - orig_bounds.Min.X) / float64(ret_bounds.Max.X - ret_bounds.Min.X)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := ret_bounds.Min.Y + rank; y <= ret_bounds.Max.Y; y += n {
                for x := ret_bounds.Min.X; x <= ret_bounds.Max.X; x++ {

                     r11, g11, b11, a11 := orig.At(scale_index(x, width_ratio), scale_index(y, height_ratio)).RGBA()
                     r12, g12, b12, a12 := orig.At(scale_index(x + 1, width_ratio), scale_index(y, height_ratio)).RGBA()
                     r21, g21, b21, a21 := orig.At(scale_index(x, width_ratio), scale_index(y + 1, height_ratio)).RGBA()
                     r22, g22, b22, a22 := orig.At(scale_index(x + 1, width_ratio), scale_index(y + 1, height_ratio)).RGBA()

                     r_aux1 := linear_interpolation(r11, r12, 0.5)
                     g_aux1 := linear_interpolation(g11, g12, 0.5)
                     b_aux1 := linear_interpolation(b11, b12, 0.5)
                     a_aux1 := linear_interpolation(a11, a12, 0.5)

                     r_aux2 := linear_interpolation(r21, r22, 0.5)
                     g_aux2 := linear_interpolation(g21, g22, 0.5)
                     b_aux2 := linear_interpolation(b21, b22, 0.5)
                     a_aux2 := linear_interpolation(a21, a22, 0.5)

                     r_fin := linear_interpolation(r_aux1, r_aux2, 0.5)
                     g_fin := linear_interpolation(g_aux1, g_aux2, 0.5)
                     b_fin := linear_interpolation(b_aux1, b_aux2, 0.5)
                     a_fin := linear_interpolation(a_aux1, a_aux2, 0.5)

                     ret.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

}

func main() {
  img_file, _ := os.Open("test.jpg")
  defer img_file.Close()
  img, _, _ := image.Decode(img_file)

  // card := image.NewRGBA(image.Rect(0, 0, 800, 800))
  // draw.Draw(card, card.Bounds(), img, image.Point{0, 0}, draw.Src)

  // f_elem := float64(1) / 9
  // f := Filter { [3][3]float64 { {f_elem, f_elem, f_elem}, {f_elem, f_elem, f_elem}, {f_elem, f_elem, f_elem}}}

  // apply_filter(card, image.Point {10, 10}, image.Point {100, 100}, f)

  card := image.NewRGBA(image.Rect(0, 0, 1750, 400))
  resize(img, card)

  fmt.Println("done!")

  rez, _ := os.Create("new.jpg")
  defer rez.Close()

  jpeg.Encode(rez, card, &jpeg.Options{jpeg.DefaultQuality})
}
