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

    card := image.NewRGBA(image.Rect(0, 0, 800, 800))
    lib.Rescale(img, card)

    lib.Apply_filter(card, image.Point {0, 0}, image.Point {800, 800}, f, 100)

    fmt.Println("done!")

    rez, _ := os.Create("new.jpg")
    defer rez.Close()

    jpeg.Encode(rez, card, &jpeg.Options{jpeg.DefaultQuality})
}
