package lib

import (
    "os"
    "strings"
    "strconv"
    "image"
    "image/jpeg"
)

func Crop_ut(input, output string) {
    img_file, _ := os.Open(input)
    defer img_file.Close()
    aux, _, _ := image.Decode(img_file)
    img := image.Image(aux)

    file_num := 0
    var curr_output string

    curr_output = strings.Join([]string{output, "_", strconv.Itoa(file_num), ".jpg"}, "")
    file_num ++
    crop_ut_1(img, curr_output)

}

func crop_ut_1(img image.Image, output string) {
    bounds := img.Bounds()
    width := bounds.Max.X - bounds.Min.X
    height := bounds.Max.Y - bounds.Min.Y

    ret := Crop_image(img, image.Rect(bounds.Min.X + width / 4, bounds.Min.Y + height / 4, bounds.Max.X - width / 4, bounds.Max.Y - height / 4))

    rez, _ := os.Create(output)
    defer rez.Close()
    jpeg.Encode(rez, ret, &jpeg.Options{jpeg.DefaultQuality})
}
