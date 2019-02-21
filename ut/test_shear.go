package ut

import (
    "os"
    "strings"
    "strconv"
    "image"
    "image/jpeg"
    lib "../lib"
)

func Test_shear(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := image.Image(aux)

    file_num := 1
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_1(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_2(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_3(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_4(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_5(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_6(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_7(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_8(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_9(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_10(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_11(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    shear_test_12(img, curr_output)
}

func shear_test_1(img image.Image, output string) {
    ret := lib.Shear(img, 0.25, lib.XSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_2(img image.Image, output string) {
    ret := lib.Shear(img, 0.25, lib.YSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_3(img image.Image, output string) {
    ret := lib.Shear(img, -0.25, lib.XSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_4(img image.Image, output string) {
    ret := lib.Shear(img, -0.25, lib.YSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_5(img image.Image, output string) {
    ret := lib.Shear(img, 0.25, lib.XSHEAR)
    ret = lib.Shear(ret, 0.25, lib.YSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_6(img image.Image, output string) {
    ret := lib.Shear(img, -0.25, lib.XSHEAR)
    ret = lib.Shear(ret, 0.25, lib.YSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_7(img image.Image, output string) {
    ret := lib.Shear(img, 0.25, lib.XSHEAR)
    ret = lib.Shear(ret, -0.25, lib.YSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_8(img image.Image, output string) {
    ret := lib.Shear(img, -0.25, lib.XSHEAR)
    ret = lib.Shear(ret, -0.25, lib.YSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_9(img image.Image, output string) {
    ret := lib.Shear(img, 0.25, lib.YSHEAR)
    ret = lib.Shear(ret, 0.25, lib.XSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_10(img image.Image, output string) {
    ret := lib.Shear(img, 0.25, lib.YSHEAR)
    ret = lib.Shear(ret, -0.25, lib.XSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_11(img image.Image, output string) {
    ret := lib.Shear(img, -0.25, lib.YSHEAR)
    ret = lib.Shear(ret, 0.25, lib.XSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func shear_test_12(img image.Image, output string) {
    ret := lib.Shear(img, -0.25, lib.YSHEAR)
    ret = lib.Shear(ret, -0.25, lib.XSHEAR)

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
