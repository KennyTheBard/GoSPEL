package ut

import (
    "os"
    "strings"
    "strconv"
    "image"
    "image/jpeg"
    lib "../lib"
    trans "../lib/generators/transformations"
)

func Test_transform(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := image.Image(aux)

    file_num := 1
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    transform_test_1(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    transform_test_2(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    transform_test_3(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    transform_test_4(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    transform_test_5(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    transform_test_6(img, curr_output)
}

func transform_test_1(img image.Image, output string) {
    ret := lib.Transform(img, trans.MirrorFunction(trans.VERTICAL_MODE))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func transform_test_2(img image.Image, output string) {
    ret := lib.Transform(img, trans.MirrorFunction(trans.HORIZONTAL_MODE))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func transform_test_3(img image.Image, output string) {
    ret := lib.Transform(img, trans.MirrorFunction(trans.HORIZONTAL_MODE))
    ret = lib.Transform(ret, trans.MirrorFunction(trans.VERTICAL_MODE))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func transform_test_4(img image.Image, output string) {
    ret := lib.Transform(img, trans.MirrorFunction(trans.VERTICAL_MODE))
    ret = lib.Transform(ret, trans.MirrorFunction(trans.HORIZONTAL_MODE))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func transform_test_5(img image.Image, output string) {
    ret := lib.Transform(img, trans.MirrorFunction(trans.VERTICAL_MODE))
    ret = lib.Transform(ret, trans.MirrorFunction(trans.VERTICAL_MODE))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func transform_test_6(img image.Image, output string) {
    ret := lib.Transform(img, trans.MirrorFunction(trans.HORIZONTAL_MODE))
    ret = lib.Transform(ret, trans.MirrorFunction(trans.HORIZONTAL_MODE))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
