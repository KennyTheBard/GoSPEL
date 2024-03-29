package ut

import (
    "os"
    "strings"
    "strconv"
    "image"
    "image/jpeg"
    lib "../lib"
)

func Test_resize(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := image.Image(aux)

    file_num := 1
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    resize_test_1(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    resize_test_2(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    resize_test_3(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    resize_test_4(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    resize_test_5(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    resize_test_6(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    resize_test_7(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    resize_test_8(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    resize_test_9(img, curr_output)

}

func resize_test_1(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X
    height := bounds.Max.Y - bounds.Min.Y

    ret := lib.Resize(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X - 3 * width / 4, bounds.Max.Y - 3 * height / 4))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func resize_test_2(img image.Image, output string) {
    bounds := img.Bounds()
    height := bounds.Max.Y - bounds.Min.Y

    ret := lib.Resize(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y - 3 * height / 4))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func resize_test_3(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X

    ret := lib.Resize(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X - 3 * width / 4, bounds.Max.Y))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func resize_test_4(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X
    height := bounds.Max.Y - bounds.Min.Y

    ret := lib.Resize(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X + 3 * width / 4, bounds.Max.Y + 3 * height / 4))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func resize_test_5(img image.Image, output string) {
    bounds := img.Bounds()
    height := bounds.Max.Y - bounds.Min.Y

    ret := lib.Resize(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y + 3 * height / 4))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func resize_test_6(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X

    ret := lib.Resize(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X + 3 * width / 4, bounds.Max.Y))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func resize_test_7(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X
    height := bounds.Max.Y - bounds.Min.Y

    ret := lib.Resize(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X + 3 * width, bounds.Max.Y + 3 * height))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func resize_test_8(img image.Image, output string) {
    bounds := img.Bounds()
    height := bounds.Max.Y - bounds.Min.Y

    ret := lib.Resize(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y + 3 * height))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func resize_test_9(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X

    ret := lib.Resize(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X + 3 * width, bounds.Max.Y))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
