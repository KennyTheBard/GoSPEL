package ut

import (
    "os"
    "strings"
    "strconv"
    "image"
    "image/jpeg"
    lib "../lib"
)

func Test_crop(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := image.Image(aux)

    file_num := 1
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    crop_test_1(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    crop_test_2(img, curr_output)

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    crop_test_3(img, curr_output)

}

func crop_test_1(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X
    height := bounds.Max.Y - bounds.Min.Y

    ret := lib.Crop(img, image.Rect(bounds.Min.X + width / 4, bounds.Min.Y + height / 4, bounds.Max.X - width / 4, bounds.Max.Y - height / 4))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func crop_test_2(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X
    height := bounds.Max.Y - bounds.Min.Y

    ret := lib.Crop(img, image.Rect(bounds.Min.X - width / 4, bounds.Min.Y - height / 4, bounds.Max.X - width / 4, bounds.Max.Y - height / 4))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}

func crop_test_3(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X
    height := bounds.Max.Y - bounds.Min.Y

    ret := lib.Crop(img, image.Rect(bounds.Min.X + width / 4, bounds.Min.Y + height / 4, bounds.Max.X + width / 4, bounds.Max.Y + height / 4))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
