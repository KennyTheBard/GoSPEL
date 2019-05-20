package main

import (
    "fmt"
    "time"
    "image"
    "image/color"
    lib "./lib"
    //ut "./ut"
    // filters "./lib/generators/filters"
    // modifiers "./lib/generators/modifiers"
    // trans "./lib/generators/transformations"
)

func main() {
    img := lib.DecodeImage("test.jpg")

    start := time.Now()

    card := lib.Copy(img)

    // card = lib.Rotate(lib.Resize(card, image.Rect(0, 0, 800, 800)), 275)

    // card = lib.Resize(card, image.Rect(0, 0, 200, 200))
    // area := lib.CreateImage(card.Bounds(), color.RGBA{255, 255, 255, 0})
    // area := lib.SelectColor(card, color.RGBA{255, 50, 50, 255}, 0.3)
    // card = lib.ModifyColors(card, area, lib.Modifier{[4][4]float64{{0, 1, 0, 0}, {0, 0, 1, 0}, {1, 0, 0, 0}, {0, 0, 0, 1}}, [4]float64{0, 0, 0, 0}})
    // card = lib.AddHSV(card, area, 75, -0.3, -0.3)
    // card = lib.Noise(card, area, lib.SALT_AND_PEPPER, 100, 0.05)
    // card = lib.Median(card, area, 3)
    // f := filters.AxialBlur(25, 1)
    // f := filters.BoxBlur(9)
    // card = lib.ApplyFilter(card, card.Bounds(), f, 1)
    // card = lib.ModifyColors(card, card.Bounds(), modifiers.ExctractColorChannel(modifiers.RED_CHANNEL))

    //card = lib.Transform(card, trans.SwirlFunc(-0.001))
    b := card.Bounds()

    gmap := lib.GradientMap{ []lib.ColorCore{
                            lib.ColorCore{image.Point{0, 0}, color.RGBA{255, 0, 0, 255}},
                            lib.ColorCore{image.Point{100, 200}, color.RGBA{0, 255, 0, 255}},
                            lib.ColorCore{image.Point{350, 120}, color.RGBA{255, 0, 0, 255}},
                            lib.ColorCore{image.Point{700, 50}, color.RGBA{0, 255, 0, 255}},
                            lib.ColorCore{image.Point{145, 430}, color.RGBA{255, 0, 0, 255}},
                            lib.ColorCore{image.Point{600, 110}, color.RGBA{0, 255, 0, 255}},
                            lib.ColorCore{image.Point{80, 710}, color.RGBA{255, 0, 0, 255}},
                            lib.ColorCore{image.Point{500, 500}, color.RGBA{0, 255, 0, 255}},
                            lib.ColorCore{image.Point{b.Max.X, b.Max.Y}, color.RGBA{0, 0, 255, 255}}} }

    grd := lib.Gradient(b, gmap)
    // fmt.Println(grd)
    card = lib.Merge(card, grd, image.Point{0, 0})

    elapsed := time.Since(start)

    fmt.Printf("Done in %s!\n", elapsed)

    lib.EncodeImage(card, "new", "png")

}
