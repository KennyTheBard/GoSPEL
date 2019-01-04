package lib

import (
    "image"
    "image/draw"
    "image/color"
    "math"
    interpolation "./interpolation"
)

func Linear_gradient(bounds image.Rectangle, ys []int, vals []image.Color) image.Image {
    img = image.NewRGBA(bounds)

    // first block
    {
        r, g, b, a := vals[0].RGBA()
        for y := bounds.Min.Y; y <= ys[0]; y++ {
            for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                img.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
            }
        }
    }

    // middle blocks
    {
        n := len(ys) - 1
        curr := ys[0]
        var prev int

        // for each block
        for k = 1; k <= n; k++ {
            prev = curr
            curr = ys[k]

            r1, g1, b1, a1 := vals[k - 1].RGBA()
            r2, g2, b2, a2 := vals[k].RGBA()

            for y := prev + 1; y <= curr; y++ {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    proc := float64(y) / float64(curr - prev)
                    r_fin, g_fin, b_fin, a_fin := Pixel_bilinear_interpolation(r1, g1, b1, a1, r2, g2, b2, a2, proc)
                    
                    img.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
                }
            }

        }
    }

    // last block
    {
        n := len(ys) - 1
        r, g, b, a := vals[n].RGBA()
        for y := ys[n] + 1; y <= bounds.Max.Y; y++ {
            for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                img.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
            }
        }
    }



    prev := bounds.Min.Y
    for curr := range ys {
        for y := prev + 1; y <= curr; y++ {
            r, g, b, a := img.At(x, y).RGBA()

            r_fin, g_fin, b_fin, a_fin := Pixel_bilinear_interpolation()

            for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                img.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
            }
        }

        prev = y
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
