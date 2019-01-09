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

    card := lib.Copy(img)
    card = lib.Resize(card, image.Rect(0, 0, 800, 800))
    grd := lib.Circular_gradient(800, []int{0, 200, 400}, []color.Color{ color.RGBA{0, 0, 0, 0}, color.RGBA{128, 0, 128, 128}, color.RGBA{255, 128, 0, 255}})
    //card = lib.Apply_filter(card, grd, lib.Filter{ [][]float64{{1.0/9, 1.0/9, 1.0/9}, {1.0/9, 1.0/9, 1.0/9}, {1.0/9, 1.0/9, 1.0/9}} }, 100)
    card = lib.Modify_colors(card, grd, [4][5]float64{{-1, 0, 0, 0, (256 << 8) - 1}, {0, -1, 0, 0, (256 << 8) - 1}, {0, 0, -1, 0, (256 << 8) - 1}, {0, 0, 0, 1, 0}})
    //card = lib.Merge(card, grd, image.Rect(0, 0, 400, 200))
    //ut.Test_crop("test.jpg", "test_results/crop_test")
    // ut.Test_resize("test.jpg", "test_results/resize_test")
    // ut.Test_mirror("test.jpg", "test_results/mirror_test")
    // ut.Test_copy("test.jpg", "test_results/copy_test")
    ut.Test_gradient("test.jpg", "test_results/gradient_test")

    elapsed := time.Since(start)


    fmt.Printf("Done in %s!\n", elapsed)

    rez, _ := os.Create("new.jpg")
    defer rez.Close()

    jpeg.Encode(rez, card, &jpeg.Options{jpeg.DefaultQuality})

}
