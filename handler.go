package main

import (
  //"fmt"
  "os"
  "image"
  "image/draw"
  "image/jpeg"
  "image/color"
)

type Filter struct {
  mat [3][3]float64
}

func apply_filter(img image.Image, start image.Point, end image.Point, f Filter) {
  // bounds := img.Bounds();

  n := 10
  for i := 0; i < n; i ++ {
    go func() {
      rank := i

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
    } ()
  }
}

func main() {
  img_file, _ := os.Open("test.jpg")
  defer img_file.Close()
  img, _, _ := image.Decode(img_file)

  card := image.NewRGBA(image.Rect(0, 0, 800, 800))
  draw.Draw(card, card.Bounds(), img, image.Point{0, 0}, draw.Src)

  f_elem := float64(1) / 9
  f := Filter { [3][3]float64 { {f_elem, f_elem, f_elem}, {f_elem, f_elem, f_elem}, {f_elem, f_elem, f_elem}}}

  apply_filter(card, image.Point {10, 10}, image.Point {400, 400}, f)

  rez, _ := os.Create("new.jpg")
  defer rez.Close()

  jpeg.Encode(rez, card, &jpeg.Options{jpeg.DefaultQuality})
}
