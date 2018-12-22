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

func min(a, b uint32) uint32 {
    if a > b {
        return b
    } else {
        return a
    }
}

func resize(orig image.Image, ret image.Image) {
    orig_bounds := orig.Bounds()
    ret_bounds := ret.Bounds()

    // ratio return to original
    height_ratio := float64(ret_bounds.Max.Y - ret_bounds.Min.Y) / float64(orig_bounds.Max.Y - orig_bounds.Min.Y)
    width_ratio := float64(ret_bounds.Max.X - ret_bounds.Min.X) / float64(orig_bounds.Max.X - orig_bounds.Min.X)

    // n := 10
    // for p := 0; p < n; p ++ {
    //     go func() {
    //         rank := p

            for y := orig_bounds.Min.Y /*+ rank*/; y <= orig_bounds.Max.Y; y ++/*+= n*/ {
                for x := orig_bounds.Min.X; x <= orig_bounds.Max.X; x++ {
                    r, g, b, a := orig.At(y, x).RGBA()

                    // fmt.Print("[",y,"][",x,"] enters in riles ")
                    // fmt.Print("from ",uint64(math.Floor(float64(y) * height_ratio))," to ",uint64(math.Ceil(float64(y + 1) * height_ratio)) - 1, " ")
                    // fmt.Println("and from",uint64(math.Floor(float64(x) * width_ratio)),"to",uint64(math.Ceil(float64(x + 1) * width_ratio)) - 1)

                    for i := uint64(math.Floor(float64(y) * height_ratio)); i < uint64(math.Ceil(float64(y + 1) * height_ratio)); i++ {
                        for j := uint64(math.Floor(float64(x) * width_ratio)); j < uint64(math.Ceil(float64(x + 1) * width_ratio)); j++ {
                            nr, ng, nb, na := ret.At(int(i), int(b)).RGBA()

                            proc_h := (math.Min(float64(y + 1) * height_ratio, float64(i + 1)) - math.Max(float64(y) * height_ratio, float64(i))) / height_ratio
                            proc_w := (math.Min(float64(x + 1) * width_ratio, float64(j + 1)) - math.Max(float64(x) * width_ratio, float64(j))) / width_ratio

                            // if x == 5 && y == 5 {
                            //     fmt.Println("i =", i)
                            //     fmt.Println("j =", j)
                                // fmt.Println("height_ratio =", height_ratio)
                                // fmt.Println("width_ratio =", width_ratio)
                                // fmt.Println("proc_h =", proc_h)
                                // fmt.Println("proc_w =", proc_w)
                                // fmt.Println("math.Min(float64(y + 1) * height_ratio, float64(i + 1)) =", math.Min(float64(y + 1) * height_ratio, float64(i + 1)))
                                // fmt.Println("math.Max(float64(y) * height_ratio, float64(i) =", math.Max(float64(y) * height_ratio, float64(i)))
                                // fmt.Println("math.Min(float64(x + 1) * width_ratio, float64(j + 1)) =", math.Min(float64(x + 1) * width_ratio, float64(j + 1)))
                                // fmt.Println("math.Max(float64(x) * width_ratio, float64(j)) =", math.Max(float64(x) * width_ratio, float64(j)))
                                // fmt.Println()
                            // }
                            //fmt.Println(nr, "+", uint32(proc_h * proc_w * float64(r)))

                            nr += uint32(proc_h * proc_w * float64(r))
                            ng += uint32(proc_h * proc_w * float64(g))
                            nb += uint32(proc_h * proc_w * float64(b))
                            na += uint32(proc_h * proc_w * float64(a))

                            // nr = min(255 << 8 - 1, nr + uint32(proc_h * proc_w * float64(r)))
                            // ng = min(255 << 8 - 1, ng + uint32(proc_h * proc_w * float64(g)))
                            // nb = min(255 << 8 - 1, nb + uint32(proc_h * proc_w * float64(b)))
                            // na = min(255 << 8 - 1, na + uint32(proc_h * proc_w * float64(a)))

                            ret.(draw.Image).Set(int(i), int(j), color.RGBA{uint8(nr >> 8), uint8(ng >> 8), uint8(nb >> 8), uint8(na >> 24)})
                        }
                    }
                }
            }
    //     } ()
    // }

}

func main() {
  img_file, _ := os.Open("test.jpg")
  defer img_file.Close()
  img, _, _ := image.Decode(img_file)

  card := image.NewRGBA(image.Rect(0, 0, 800, 800))
  draw.Draw(card, card.Bounds(), img, image.Point{0, 0}, draw.Src)

  // f_elem := float64(1) / 9
  // f := Filter { [3][3]float64 { {f_elem, f_elem, f_elem}, {f_elem, f_elem, f_elem}, {f_elem, f_elem, f_elem}}}

  // apply_filter(card, image.Point {10, 10}, image.Point {100, 100}, f)

  card2 := image.NewRGBA(image.Rect(0, 0, 800, 800))
  resize(card, card2)

  fmt.Println("done!")

  rez, _ := os.Create("new.jpg")
  defer rez.Close()

  jpeg.Encode(rez, card2, &jpeg.Options{jpeg.DefaultQuality})
}
