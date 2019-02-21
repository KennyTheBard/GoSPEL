package lib

import (
    "math"
    "image"
    "image/draw"
    "image/color"
    utils "./utils"
    interp "./interpolation"
)

func RGB2HSV(r, g, b uint32) (int32, float64, float64) {
    r_sec := float64(r) / float64(255 << 8)
    g_sec := float64(g) / float64(255 << 8)
    b_sec := float64(b) / float64(255 << 8)

    var c_max, c_min, c_med float64

    if r_sec >= g_sec && r_sec >= b_sec {
        c_max = r_sec
    } else if g_sec >= r_sec && g_sec >= b_sec {
        c_max = g_sec
    } else {
        c_max = b_sec
    }

    if r_sec <= g_sec && r_sec <= b_sec {
        c_min = r_sec
    } else if g_sec <= r_sec && g_sec <= b_sec {
        c_min = g_sec
    } else {
        c_min = b_sec
    }

    if (c_max == g_sec && c_min == b_sec) || (c_min == g_sec && c_max == b_sec) {
        c_med = r_sec
    } else if (c_max == r_sec && c_min == b_sec) || (c_min == r_sec && c_max == b_sec) {
        c_med = g_sec
    } else {
        c_med = b_sec
    }

    delta := c_max - c_min
    delta2 := c_med - c_min

    // calculate hue
    var hue int32
    fractio := delta2 / delta
    if r >= g && g >= b {
        hue = int32(60 * fractio)
    } else if g > r && r >= b {
        hue = int32(60 * (2 - fractio))
    } else if g >= b && b > r {
        hue = int32(60 * (2 + fractio))
    } else if b > g && g > r {
        hue = int32(60 * (4 - fractio))
    } else if b > r && r >= g {
        hue = int32(60 * (4 + fractio))
    } else if r >= b && b > g {
        hue = int32(60 * (6 - fractio))
    }

    // calculate saturation
    var saturation float64
    if c_max == 0 {
        saturation = 0
    } else {
        saturation = delta / c_max
    }

    // calculate value
    value := c_max

    return hue, saturation, value
}

func HSV2RGB(hue int32, saturation float64, value float64) (uint32, uint32, uint32) {
    c := value * saturation
    aux_hue := float64(hue) / 60
    for aux_hue > 2 {
        aux_hue -= 2
    }
    x := c * (1 - math.Abs(aux_hue - 1))
    m := value - c

    var r_sec, g_sec, b_sec float64

    if hue < 60 {
        r_sec = c
        g_sec = x
        b_sec = 0

    } else if hue < 120 {
        r_sec = x
        g_sec = c
        b_sec = 0

    } else if hue < 180 {
        r_sec = 0
        g_sec = c
        b_sec = x

    } else if hue < 240 {
        r_sec = 0
        g_sec = x
        b_sec = c

    } else if hue < 300 {
        r_sec = x
        g_sec = 0
        b_sec = c

    } else {
        r_sec = c
        g_sec = 0
        b_sec = x

    }

    r, g, b := uint32((r_sec + m) * (255 << 8)), uint32((g_sec + m) * (255 << 8)), uint32((b_sec + m) * (255 << 8))

    return r, g, b
}

func pixelAddHSV(pixel color.Color, hue int32, saturation float64, value float64) (color.Color) {
    r, g, b, a := pixel.RGBA()

    h, s, v := RGB2HSV(r, g, b)

    h += hue
    s += saturation
    v += value

    for h < 0 {
        h += 360
    }
    for h >= 360 {
        h -= 360
    }

    s = utils.Fclamp(0, 1, s)
    v = utils.Fclamp(0, 1, v)

    r, g, b = HSV2RGB(h, s, v)

    return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

func AddHSV(img image.Image, mask image.Image, hue int32, saturation float64, value float64) (image.Image) {
    bounds := img.Bounds()
    ret := Copy(img)
    mask = Resize(mask, bounds)

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y <= bounds.Max.Y; y += n {
                for x := bounds.Min.X; x <= bounds.Max.X; x++ {
                    r, g, b, a := pixelAddHSV(img.At(x, y), hue, saturation, value).RGBA()

                    r_aux, g_aux, b_aux, a_aux := img.At(x, y).RGBA()
                    r_mask, g_mask, b_mask, a_mask := mask.At(x, y).RGBA()

                    // calculate the color modification through mask
                    r_fin := uint32(interp.LinearInterpolation(int32(r_aux), int32(r), float64(r_mask) / float64((256 << 8) - 1)))
                    g_fin := uint32(interp.LinearInterpolation(int32(g_aux), int32(g), float64(g_mask) / float64((256 << 8) - 1)))
                    b_fin := uint32(interp.LinearInterpolation(int32(b_aux), int32(b), float64(b_mask) / float64((256 << 8) - 1)))
                    a_fin := uint32(interp.LinearInterpolation(int32(a_aux), int32(a), float64(a_mask) / float64((256 << 8) - 1)))

                    ret.(draw.Image).Set(x, y, color.RGBA{uint8(r_fin >> 8), uint8(g_fin >> 8), uint8(b_fin >> 8), uint8(a_fin >> 8)})
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
