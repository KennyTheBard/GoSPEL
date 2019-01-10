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

func Test_opacity(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := lib.Copy(image.Image(aux))

    file_num := 1
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    opacity_test_1(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    opacity_test_2(img, curr_output)

}

func opacity_test_1(img image.Image, output string) {
    aux_img := lib.Scale_opacity(img, 0)
    ret := lib.Merge(img, aux_img, img.Bounds())
    ret = lib.Crop(ret, lib.Select_opaque(ret))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func opacity_test_2(img image.Image, output string) {
    aux_img := lib.Create_image(image.Rect(0, 0, img.Bounds().Max.X * 2, img.Bounds().Max.Y * 2), color.RGBA{255, 0, 0, 0})
    ret := lib.Merge(aux_img, img, img.Bounds())
    ret = lib.Crop(ret, lib.Select_opaque(ret))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
