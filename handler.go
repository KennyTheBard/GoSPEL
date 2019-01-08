package main

import (
    "fmt"
    "os"
    "time"
    "image"
    "image/jpeg"
    "image/color"
    lib "./lib"
    ut "./ut"
)

func main() {
    img_file, _ := os.Open("test.jpg")
    defer img_file.Close()
    img, _, _ := image.Decode(img_file)

    start := time.Now()

    card := lib.Copy_image(img)
    card = lib.Rescale(card, image.Rect(0, 0, 800, 800))
    //lib.Apply_filter(card, image.Rect(200, 200, 600, 600), lib.Filter{ [][]float64{{1.0/9, 1.0/9, 1.0/9}, {1.0/9, 1.0/9, 1.0/9}, {1.0/9, 1.0/9, 1.0/9}} }, 50)
    //lib.Modify_colors(card, image.Rect(200, 200, 600, 600), [4][5]float64{{-1, 0, 0, 0, (256 << 8) - 1}, {0, -1, 0, 0, (256 << 8) - 1}, {0, 0, -1, 0, (256 << 8) - 1}, {0, 0, 0, 1}})
    grd := lib.Circular_gradient(800, 800, []int{0, 200, 400}, []color.Color{ color.RGBA{255, 0, 0, 0}, color.RGBA{128, 0, 0, 128}, color.RGBA{0, 0, 0, 255}})
    card = lib.Merge(card, grd, image.Rect(0, 0, 800, 800))

    elapsed := time.Since(start)

    ut.Test_crop("test.jpg", "test_results/crop_test")

    fmt.Printf("Done in %s!\n", elapsed)

    rez, _ := os.Create("new.jpg")
    defer rez.Close()

    jpeg.Encode(rez, card, &jpeg.Options{jpeg.DefaultQuality})

}
