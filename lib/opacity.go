package lib

import (
    "image"
    "image/draw"
    "image/color"
    utils "./utils"
)

/**
    Returns an image with the opacity ampified by
*/
func Scale_opacity(img image.Image, alpha_proc float64) (image.Image) {
    // early exit case
    if alpha_proc == 1 {
        return Copy(img)
    }

    // prepare the image to be returned
    bounds := img.Bounds()
    ret := image.Image(image.NewRGBA(bounds))

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()

                    a_fin := utils.Clamp(0, (256 << 8) - 1, uint32(alpha_proc * float64(a)))

                    ret.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a_fin >> 8)})
                }
            }

            done <- true;
        } ()
    }

    for i := 0; i < n; i++ {
        <-done
    }

    return ret
}


/**
    Returns the area of the image that contains all
pixels with alpha value greater than min_alpha.
*/
func Select_opaque(img image.Image, min_alpha uint32) (image.Rectangle) {
    var min_x, max_x, min_y, max_y int
    var found bool

    bounds := img.Bounds()

    // find minimum x with non-completly transparent property
    found = false
    for y := bounds.Min.Y; y <= bounds.Max.Y; y ++ {
        for x := bounds.Min.X; x <= bounds.Max.X; x ++ {
            _, _, _, a := img.At(x, y).RGBA()

            if a > min_alpha {
                min_y = y
                found = true;
                break
            }
        }

        if found {
            break
        }
    }

    // find maximum x with non-completly transparent property
    found = false
    for y := bounds.Max.Y; y >= bounds.Min.Y; y -- {
        for x := bounds.Min.X; x <= bounds.Max.X; x ++ {
            _, _, _, a := img.At(x, y).RGBA()

            if a > min_alpha {
                max_y = y
                found = true;
                break
            }
        }

        if found {
            break
        }
    }

    // find minimum y with non-completly transparent property
    found = false
    for x := bounds.Min.X; x <= bounds.Max.X; x ++ {
        for y := bounds.Min.Y; y <= bounds.Max.Y; y ++ {
            _, _, _, a := img.At(x, y).RGBA()

            if a > min_alpha {
                min_x = x
                found = true;
                break
            }
        }

        if found {
            break
        }
    }

    // find maximum y with non-completly transparent property
    found = false
    for x := bounds.Max.X; x >= bounds.Min.X; x -- {
        for y := bounds.Min.Y; y <= bounds.Max.Y; y ++ {
            _, _, _, a := img.At(x, y).RGBA()

            if a > min_alpha {
                max_x = x
                found = true;
                break
            }
        }

        if found {
            break
        }
    }

    return image.Rect(min_x, min_y, max_x, max_y)
}
