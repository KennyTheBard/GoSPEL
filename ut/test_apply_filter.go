package ut

import (
    "os"
    "strings"
    "strconv"
    "image"
    "image/jpeg"
    "image/color"
    lib "../lib"
)

func Test_apply_filter(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := image.Image(aux)

    file_num := 1
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    apply_filter_test_1(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    apply_filter_test_2(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    apply_filter_test_3(img, curr_output)
}

func apply_filter_test_1(img image.Image, output string) {
    bounds := img.Bounds()

    grd := lib.CircularGradient(800, []int{0, 200, 400}, []color.Color{ color.RGBA{0, 0, 0, 0}, color.RGBA{128, 0, 128, 128}, color.RGBA{255, 128, 0, 255}})
    grd = lib.Resize(grd, bounds)
    ret := lib.ApplyFilter(img, grd, lib.Filter{ [][]float64{{1.0/9, 1.0/9, 1.0/9}, {1.0/9, 1.0/9, 1.0/9}, {1.0/9, 1.0/9, 1.0/9}} }, 20)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func apply_filter_test_2(img image.Image, output string) {
    bounds := img.Bounds()

    grd := lib.CircularGradient(800, []int{0, 200, 400}, []color.Color{ color.RGBA{0, 0, 0, 0}, color.RGBA{128, 0, 128, 128}, color.RGBA{255, 128, 0, 255}})
    grd = lib.Resize(grd, bounds)
    ret := lib.ApplyFilter(img, grd, lib.Filter{ [][]float64{{1.0/16, 2.0/16, 1.0/16}, {2.0/16, 4.0/16, 2.0/16}, {1.0/16, 2.0/16, 1.0/16}} }, 20)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func apply_filter_test_3(img image.Image, output string) {
    bounds := img.Bounds()

    grd := lib.CircularGradient(800, []int{0, 200, 400}, []color.Color{ color.RGBA{0, 0, 0, 0}, color.RGBA{128, 0, 128, 128}, color.RGBA{255, 128, 0, 255}})
    grd = lib.Resize(grd, bounds)
    ret := lib.ApplyFilter(img, grd, lib.Filter{ [][]float64{{1, 0, -1}, {0, 0, 0}, {-1, 0, 1}} }, 5)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
