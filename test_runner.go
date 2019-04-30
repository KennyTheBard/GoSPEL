package main

import (
    "fmt"
    "time"
    ut "./ut"
)

type Test func(string, string)

func run_test(t Test, fin, dir_out, fout string, done chan<- string) {
    t(fin, dir_out + fout)
    done <- fout
}

func main() {

    tests := []Test{
        ut.Test_crop,
        ut.Test_resize,
        ut.Test_transform,
        ut.Test_copy,
        ut.Test_gradient,
        ut.Test_shear,
        ut.Test_merge,
        ut.Test_apply_filter,
        ut.Test_modify_colors,
        ut.Test_rotate,
        ut.Test_opacity}
    fin := "test.jpg"
    dir_out := "test_results/"
    fout := []string{
        "crop_test",
        "resize_test",
        "mirror_test",
        "copy_test",
        "gradient_test",
        "shear_test",
        "merge_test",
        "apply_filter_test",
        "modify_colors_test",
        "rotate_test",
        "opacity_test"}

    if len(tests) != len(fout) {
        fmt.Println("Number of tests does not coincide with number of output file names")
    }

    done := make(chan string, len(tests))

    start := time.Now()

    for i := 0; i < len(tests); i ++ {
        go run_test(tests[i], fin, dir_out, fout[i], done)
    }

    for i := 0; i < len(tests); i++ {
        str := <-done
        fmt.Println(str, "is done!")
    }

    elapsed := time.Since(start)

    fmt.Printf("Done in %s!\n", elapsed)

}
