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

func Test_merge(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := image.Image(aux)

    file_num := 1
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    merge_test_1(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    merge_test_2(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    merge_test_3(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    merge_test_4(img, curr_output)


}


func merge_test_1(img image.Image, output string) {
    grd := lib.CircularGradient(800, []int{0, 200, 800}, []color.Color{color.RGBA{255, 0, 255, 255}, color.RGBA{255, 0, 255, 125}, color.RGBA{0, 0, 0, 0}})
    ret := lib.Merge(img, grd, img.Bounds())

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func merge_test_2(img image.Image, output string) {
    grd := lib.LinearGradient(image.Rect(0, 0, 200, 200), []int{0, 200, 800}, []color.Color{color.RGBA{255, 0, 255, 255}, color.RGBA{255, 0, 255, 125}, color.RGBA{0, 0, 0, 0}})
    ret := lib.Merge(img, grd, img.Bounds())

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func merge_test_3(img image.Image, output string) {
    ret := lib.Merge(img, img, image.Rect(100, 100, 300, 300))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func merge_test_4(img image.Image, output string) {
    ret := lib.Merge(img, img, image.Rect(-100, -100, 100, 100))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
