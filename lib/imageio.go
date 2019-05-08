package lib

import (
    "os"
    "strings"
    "image"
    "image/jpeg"
    "image/png"
)

func DecodeImage(path string) (image.Image, ) {
    file, err_open := os.Open(path)
    if err_open != nil {
        return nil
    }
    defer file.Close()

    img, _, err_decode := image.Decode(file)
    if err_decode != nil {
        return nil
    }
    return img
}

func EncodeImage(img image.Image, name, format string) (image.Image) {
    fout, _ := os.Create(strings.Join([]string{name, format}, "."))
    defer fout.Close()

    switch format {
    case "png":
        png.Encode(fout, img)
    case "jpeg":
        fallthrough
    case "jpg":
        jpeg.Encode(fout, img, &jpeg.Options{jpeg.DefaultQuality})
    }

    return img
}
