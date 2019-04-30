package ut

import (
    "os"
    "strings"
    "strconv"
    "image"
    "image/jpeg"
    lib "../lib"
)

func Test_modify_colors(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := image.Image(aux)

    file_num := 1
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    modify_colors_test_1(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    modify_colors_test_2(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    modify_colors_test_3(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    modify_colors_test_4(img, curr_output)
}

func modify_colors_test_1(img image.Image, output string) {
    ret := lib.ModifyColors(img, lib.Modifier{[4][4]float64{{-1, 0, 0, 0}, {0, -1, 0, 0}, {0, 0, -1, 0}, {0, 0, 0, 1}}, [4]float64{(256 << 8) - 1, (256 << 8) - 1, (256 << 8) - 1, (256 << 8) - 1}})

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func modify_colors_test_2(img image.Image, output string) {
    ret := lib.ModifyColors(img, lib.Modifier{[4][4]float64{{0.35, 0.35, 0.35, 0}, {0.35, 0.35, 0.35, 0}, {0.35, 0.35, 0.35, 0}, {0, 0, 0, 1}}, [4]float64{0, 0, 0, 0}})

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func modify_colors_test_3(img image.Image, output string) {
    ret := lib.ModifyColors(img, lib.Modifier{[4][4]float64{{0, 1, 0, 0}, {0, 0, 1, 0}, {1, 0, 0, 0}, {0, 0, 0, 1}}, [4]float64{0, 0, 0, 0}})

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func modify_colors_test_4(img image.Image, output string) {
    ret := lib.ModifyColors(img, lib.Modifier{[4][4]float64{{0, 0.2, 0.2, 0}, {0, 1, 0, 0}, {0.2, 0.2, 0, 0}, {0, 0, 0, 1}}, [4]float64{0, 0, 0, 0}})

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
