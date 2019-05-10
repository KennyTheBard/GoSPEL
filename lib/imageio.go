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
    fout, err_create := os.Create(strings.Join([]string{name, format}, "."))
    if err_create != nil {
        return nil
    }
    defer fout.Close()

    switch format {
    case "png":
        if png.Encode(fout, img) != nil {
            return nil
        }
    case "jpeg", "jpg":
        if jpeg.Encode(fout, img, &jpeg.Options{jpeg.DefaultQuality}) != nil {
            return nil
        }
    }

    return img
}
