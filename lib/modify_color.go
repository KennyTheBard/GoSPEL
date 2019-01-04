package lib

import (
    "image"
    "image/draw"
    "image/color"
    "math"
)

/**
    Force the value val in [min, max].
*/
func Clamp(min, max uint32, val uint32) uint32 {
    if val <= min {
        return min
    } else if val >= max {
        return max
    } else {
        return val
    }
}

func Amplify_color(img image.Image, r_mod, g_mod, b_mod, a_mod float64) {
    bounds := img.Bounds()

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {

                     r, g, b, a := img.At(x, y).RGBA()

                     r_fin := uint32(math.Floor(float64(r) * r_mod))
                     g_fin := uint32(math.Floor(float64(g) * g_mod))
                     b_fin := uint32(math.Floor(float64(b) * b_mod))
                     a_fin := uint32(math.Floor(float64(a) * a_mod))

                     img.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }
}

func Add_color(img image.Image, r_mod, g_mod, b_mod, a_mod uint32) {
    bounds := img.Bounds()

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {

                     r, g, b, a := img.At(x, y).RGBA()

                     r_fin := Clamp(0, (256 << 8) - 1, r + (r_mod << 8))
                     g_fin := Clamp(0, (256 << 8) - 1, g + (g_mod << 8))
                     b_fin := Clamp(0, (256 << 8) - 1, b + (b_mod << 8))
                     a_fin := Clamp(0, (256 << 8) - 1, a + (a_mod << 8))

                     img.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }
}

func Grayscale(img image.Image, r_mod, g_mod, b_mod float64) {
    bounds := img.Bounds()

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {

                     r, g, b, a := img.At(x, y).RGBA()

                     fin := uint32(math.Floor(float64(r) * r_mod + float64(g) * g_mod + float64(b) * b_mod))

                     img.(draw.Image).Set(x, y, color.RGBA{uint8(fin >> 8), uint8(fin >> 8), uint8(fin >> 8), uint8(a >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }
}
