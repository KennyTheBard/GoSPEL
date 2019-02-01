package main

import (
    "fmt"
    "os"
    "time"
    "image"
    "image/color"
    lib "./lib"
    utils "./lib/utils"
    //ut "./ut"
    filters "./lib/generators/filters"
)

func main() {
    img_file, _ := os.Open("humans.jpg")
    defer img_file.Close()
    img, image_format, _ := image.Decode(img_file)

    start := time.Now()

    card := lib.Copy(img)

    // card = lib.Rotate(lib.Resize(card, image.Rect(0, 0, 800, 800)), 275)

    area := lib.CreateImage(card.Bounds(), color.RGBA{255, 255, 255, 0})

    //card = lib.ModifyColors(card, area, lib.Modifier{[4][4]float64{{0, 1, 0, 0}, {0, 0, 1, 0}, {1, 0, 0, 0}, {0, 0, 0, 1}}, [4]float64{0, 0, 0, 0}})
    card = lib.Noise(card, area, lib.SALT_AND_PEPPER, 100, 0.05)
    //card = lib.Median(card, area, 3)
    /* f := */filters.BoxBlur(10)

    //card = lib.ApplyFilter(card, card.Bounds(), f, 1)
    elapsed := time.Since(start)


    fmt.Printf("Done in %s!\n", elapsed)

    utils.Encode_image(card, "new", image_format)

}
