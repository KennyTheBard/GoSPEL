package main

import (
    "fmt"
    "os"
    "image"
    "image/jpeg"
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

    card := image.Image(image.NewRGBA(image.Rect(0, 0, 800, 800)))
    lib.Rescale(img, card)

    lib.Apply_filter(card, image.Point {150, 150}, image.Point {650, 650}, f, 5)

    lib.Add_color(card, 0, 25, 0, 0)

    card = lib.Mirror(card, lib.HORIZONTAL_MODE)

    fmt.Println("done!")

    rez, _ := os.Create("new.jpg")
    defer rez.Close()

    jpeg.Encode(rez, card, &jpeg.Options{jpeg.DefaultQuality})
}
