package lib

import (
    "image"
    "image/draw"
    "image/color"
    aux "./gradient_auxiliaries"
)

func Linear_gradient(bounds image.Rectangle, ys []int, vals []color.Color) image.Image {
    img := image.Image(image.NewRGBA(bounds))

    // first block
    r, g, b, a := vals[0].RGBA()
    for y := bounds.Min.Y; y <= ys[0]; y++ {
        for x := bounds.Min.X; x <= bounds.Max.X; x++ {
            img.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
        }
    }

    // middle blocks
    n := len(ys) - 1
    curr := ys[0]
    var prev int

    // for each block
    for k := 1; k <= n; k++ {
        prev = curr
        curr = ys[k]

        r1, g1, b1, a1 := vals[k - 1].RGBA()
        r2, g2, b2, a2 := vals[k].RGBA()

        for y := prev + 1; y <= curr; y++ {
            for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                proc := float64(y) / float64(curr - prev)

                r_fin, g_fin, b_fin, a_fin := aux.Pixel_bilinear_interpolation(r1, g1, b1, a1, r2, g2, b2, a2, proc)

                img.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
            }
        }

    }


    // last block
    r, g, b, a = vals[n].RGBA()
    for y := ys[n] + 1; y <= bounds.Max.Y; y++ {
        for x := bounds.Min.X; x <= bounds.Max.X; x++ {
            img.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
        }
    }

    return img;
}
