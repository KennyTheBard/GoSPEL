package main

import (
    "fmt"
    "time"
    ut "./ut"
)

func main() {

    start := time.Now()

    ut.Test_crop("test.jpg", "test_results/crop_test")
    ut.Test_resize("test.jpg", "test_results/resize_test")
    ut.Test_mirror("test.jpg", "test_results/mirror_test")
    ut.Test_copy("test.jpg", "test_results/copy_test")
    ut.Test_gradient("test.jpg", "test_results/gradient_test")
    ut.Test_shear("test.jpg", "test_results/shear_test")
    ut.Test_merge("test.jpg", "test_results/merge_test")
    ut.Test_apply_filter("humans.jpg", "test_results/apply_filter_test")
    ut.Test_modify_colors("humans.jpg", "test_results/modify_colors_test")
    ut.Test_rotate("test.jpg", "test_results/rotate_test")
    ut.Test_opacity("test.jpg", "test_results/opacity_test")

    elapsed := time.Since(start)

    fmt.Printf("Done in %s!\n", elapsed)

}
