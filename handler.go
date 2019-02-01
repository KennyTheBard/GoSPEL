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

    // ut.Test_crop("test.jpg", "test_results/crop_test")
    // ut.Test_resize("test.jpg", "test_results/resize_test")
    // ut.Test_mirror("test.jpg", "test_results/mirror_test")
    // ut.Test_copy("test.jpg", "test_results/copy_test")
    // ut.Test_gradient("test.jpg", "test_results/gradient_test")
    // ut.Test_shear("test.jpg", "test_results/shear_test")
    // ut.Test_merge("test.jpg", "test_results/merge_test")
    // ut.Test_apply_filter("humans.jpg", "test_results/apply_filter_test")
    // ut.Test_modify_colors("humans.jpg", "test_results/modify_colors_test")
    // ut.Test_rotate("test.jpg", "test_results/rotate_test")
    // ut.Test_opacity("test.jpg", "test_results/opacity_test")

    area := lib.Create_image(card.Bounds(), color.RGBA{255, 255, 255, 0})
    // card = lib.Apply_filter(card, area, lib.Filter{ [][]float64{{1.0/16, 2.0/16, 1.0/16}, {2.0/16, 4.0/16, 2.0/16}, {1.0/16, 2.0/16, 1.0/16}} }, 10)
    // card = lib.Apply_filter(card, area, lib.Filter{ [][]float64{{-1, -1, -1}, {-1, 8, -1}, {-1, -1, -1}}}, 1)

    //card = lib.Modify_colors(card, area, lib.Modifier{[4][4]float64{{0, 1, 0, 0}, {0, 0, 1, 0}, {1, 0, 0, 0}, {0, 0, 0, 1}}, [4]float64{0, 0, 0, 0}})
    card = lib.Noise(card, area, 100)
    card = lib.Median(card, area, 3)
    /* f := */filters.BoxBlur(10)

    //card = lib.Apply_filter(card, card.Bounds(), f, 1)
    elapsed := time.Since(start)


    fmt.Printf("Done in %s!\n", elapsed)

    utils.Encode_image(card, "new", image_format)

}
