package lib

import (
    "os"
    "strings"
    "image"
    "image/jpeg"
    "image/png"
)

func DecodeImage(path string) (image.Image) {
    file, _ := os.Open(path)
    defer file.Close()

    img, _, _ := image.Decode(file)
    return img
}

func EncodeImage(img image.Image, name, format string) (image.Image) {
    fout, _ := os.Create(strings.Join([]string{name, format}, "."))
    defer fout.Close()

    switch format {
    case "png":
        png.Encode(fout, img)
        break;
    case "jpeg":
        jpeg.Encode(fout, img, &jpeg.Options{jpeg.DefaultQuality})
        break;
    }

    return img
}
