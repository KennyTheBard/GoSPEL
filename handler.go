package main

import (
    "fmt"
    "os"
    "time"
    "image"
    "image/jpeg"
    //"image/png"
    lib "./lib"
)

func main() {
    img_file, _ := os.Open("humans.jpg")
    defer img_file.Close()
    img, _, _ := image.Decode(img_file)

    start := time.Now()

    card := lib.Copy_image(img)
    lib.Modify_colors(card, [4][5]float64{{-1, 0, 0, 0, (256 << 8) - 1}, {0, -1, 0, 0, (256 << 8) - 1}, {0, 0, -1, 0, (256 << 8) - 1}, {0, 0, 0, 1}})

    elapsed := time.Since(start)

    fmt.Printf("Done in %s!\n", elapsed)

    rez, _ := os.Create("new.jpg")
    defer rez.Close()

    jpeg.Encode(rez, card, &jpeg.Options{jpeg.DefaultQuality})
}
