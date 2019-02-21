package utils

import (
    "os"
    "strings"
    "image"
    "image/jpeg"
    "image/png"
)

func Encode_image(img image.Image, name, format string) {
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
}
