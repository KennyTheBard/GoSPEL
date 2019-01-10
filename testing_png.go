package main

import (
    "fmt"
    "strings"
    "os"
    "image"
    "image/jpeg"
    "image/png"
    lib "./lib"
)

func main() {

    img_file, _ := os.Open("logo.png")
    defer img_file.Close()
    img, format, _ := image.Decode(img_file)

    fmt.Println(format)
    card := lib.Rotate(lib.Copy(img), 90)

    Encode_image(card, "test_formats", format)

}

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
