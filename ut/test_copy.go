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

func Test_copy(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := image.Image(aux)

    file_num := 1
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    copy_test_1(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    copy_test_2(img, curr_output)
}

func copy_test_1(img image.Image, output string) {
    ret := lib.Copy(img)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}


func copy_test_2(img image.Image, output string) {
    grd := lib.Circular_gradient(800, []int{0, 200, 800}, []color.Color{color.RGBA{255, 0, 255, 255}, color.RGBA{255, 0, 255, 125}, color.RGBA{0, 0, 0, 0}})
    ret := lib.Merge(img, grd, img.Bounds())
    ret = lib.Copy(ret)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
