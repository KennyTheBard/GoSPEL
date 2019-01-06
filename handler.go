package main

import (
    "fmt"
    "os"
    "image"
    "image/jpeg"
    "image/color"
    lib "./lib"
)

func main() {
    img_file, _ := os.Open("test.jpg")
    defer img_file.Close()
    img, _, _ := image.Decode(img_file)

    // card := image.NewRGBA(image.Rect(0, 0, 800, 800))
    // draw.Draw(card, card.Bounds(), img, image.Point{0, 0}, draw.Src)

    f_elem := float64(1) / 9
    f := lib.Filter { [][]float64 { {f_elem, f_elem, f_elem}, {f_elem, f_elem, f_elem}, {f_elem, f_elem, f_elem}}}

    card := lib.Rescale(img, image.Rect(0, 0, 800, 800))

    lib.Apply_filter(card, image.Point {150, 150}, image.Point {650, 650}, f, 5)

    lib.Modify_colors(card, [4][4]float64{{0.35, 0, 0, 0}, {0.35, 0, 0, 0}, {0.35, 0.35, 0.3, 0}, {0, 0, 0, 1}})

    card = lib.Mirror(card, lib.HORIZONTAL_MODE)
    card = lib.Mirror(card, lib.VERTICAL_MODE)

    fmt.Println("done!")

    grd := lib.Linear_gradient(image.Rect(0, 0, 800, 800), []int{0, 700}, []color.Color{color.RGBA{255, 255, 255, 255}, color.RGBA{255, 0, 0, 125}})

    card = lib.Merge(card, grd, image.Rect(0, 0, 800, 800))

    rez, _ := os.Create("new.jpg")
    defer rez.Close()

    jpeg.Encode(rez, card, &jpeg.Options{jpeg.DefaultQuality})
}
