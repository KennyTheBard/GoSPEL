package main

import (
    "fmt"
    "os"
    "time"
    "image"
    "strings"
    //"image/color"
    lib "./lib"
    utils "./lib/utils"
    filters "./lib/generators/filters"
    gg "github.com/fogleman/gg"
)



func main() {

    img_file, _ := os.Open("test.jpg")
    defer img_file.Close()
    img, image_format, _ := image.Decode(img_file)

    start := time.Now()

    card := lib.Copy(img)
    f := filters.BoxBlur(10)
    card = lib.Apply_filter(card, card.Bounds(), f, 1)

    label := "This is a blured image"
    context := gg.NewContextForImage(card)

    if err := context.LoadFontFace("/usr/share/fonts/truetype/lato/Lato-Bold.ttf", 96); err != nil {
		panic(err)
    } else {
        fmt.Println("The label has been drawn")
    }

    context.SetRGB(0, 0, 0)
    context.DrawString(label, 100, 100)
    context.Fill()
    card = context.Image()

    elapsed := time.Since(start)
    fmt.Printf("Done in %s!\n", elapsed)

    utils.Encode_image(card, "new", image_format)

}
