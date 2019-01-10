package lib

import (
    "math"
    "image"
    "image/draw"
    "image/color"
    aux "./auxiliaries"
)

/**
    Create an image with dimensions bounds
    with a descending gradient build around
    the values vals given for each point of
    height in the ys array.

    Points of height are expected to be given
    in ascending order, from 0 to bounds.Max.Y.
*/
func Linear_gradient(bounds image.Rectangle, ys []int, vals []color.Color) image.Image {
    // early exit case
    if len(ys) != len(vals) {
        return Create_image(bounds, color.RGBA{0, 0, 0, 0})
    }

    // sort the color points
    aux.Sort_color_points(ys, vals)

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

        px1 := aux.Pixel{r1, g1, b1, a1}
        px2 := aux.Pixel{r2, g2, b2, a2}

        for y := prev + 1; y <= curr; y++ {
            for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                proc := float64(y - prev) / float64(curr - prev)

                fin := aux.Pixel_linear_interpolation(px1, px2, proc)

                img.(draw.Image).Set(x, y, color.RGBA{uint8(fin.R >> 8), uint8(fin.G >> 8), uint8(fin.B >> 8), uint8(fin.A >> 8)})
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


/**
    Returns the image with the dimensions width
    and height with a descending gradient build
    around the center of the image with the values
    vals given for each point of height in the
    ys array.

    Points of height are expected to be given
    in ascending order, from 0 to height / 2.
*/
func Circular_gradient(size int, ys []int, vals []color.Color) image.Image {
    // early exit case
    if len(ys) != len(vals) {
        return Create_image(image.Rect(0, 0, size, size), color.RGBA{0, 0, 0, 0})
    }

    // sort the color points
    aux.Sort_color_points(ys, vals)

    cx := size / 2
    cy := size / 2
    bounds := image.Rect(0, 0, size, size)
    img := image.Image(image.NewRGBA(bounds))

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    // determine the color interval
                    dst := int(math.Round(aux.Distance(float64(x), float64(y), float64(cx), float64(cy))))
                    k1, k2 := aux.Search_interval(ys, dst)

                    if k1 == k2 {
                        r, g, b, a := vals[k1].RGBA()
                        img.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
                        continue
                    }

                    // prepare data for interpolation
                    r1, g1, b1, a1 := vals[k1].RGBA()
                    r2, g2, b2, a2 := vals[k2].RGBA()
                    px1 := aux.Pixel{r1, g1, b1, a1}
                    px2 := aux.Pixel{r2, g2, b2, a2}

                    // interpolation & pixel writing
                    proc := float64(dst - ys[k1]) / float64(ys[k2] - ys[k1])

                    fin := aux.Pixel_linear_interpolation(px1, px2, proc)

                    //fmt.Println(dst, "intre", ys[k1], "si", ys[k2], "iar proc este", proc, "iar fin.A este", fin.A)

                    img.(draw.Image).Set(x, y, color.RGBA{uint8(fin.R >> 8), uint8(fin.G >> 8), uint8(fin.B >> 8), uint8(fin.A >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    return img;
}
