package lib

import (
    // "fmt"
    "math"
    "image"
    "image/draw"
    "image/color"
    utils "./utils"

)

type Gradient interface {
    DrawPixel(image.Point, image.Rectangle) (color.Color)
}

type ColorCore struct {
    Coord   float64
    Color   color.Color
}

type LinearGradient struct {
    Angle   float64
    Offset  image.Point
    Cores   []ColorCore
}

func iAbs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}

func DistanceToLine(P, A, B image.Point) float64 {
    aux_y := float64(B.Y - A.Y)
    aux_x := float64(B.X - A.X)
    length := math.Sqrt(aux_y * aux_y + aux_x * aux_x)
    return float64(iAbs((B.Y - A.Y) * P.X - (B.X - A.X) * P.Y + B.X * A.Y - B.Y * A.X)) / length
}

func (grd LinearGradient) DrawPixel(p image.Point, bounds image.Rectangle) (color.Color) {
    k := image.Point{bounds.Max.X,
        int(math.Round(float64(bounds.Max.X) * math.Tan(grd.Angle)))}
    dotProd := p.X * k.X + p.Y * k.X
    length := p.X * p.X + p.Y * p.Y
    proj := image.Point{int((dotProd * p.X) / length), int((dotProd * p.Y) / length)}

    total := k.X * k.X + k.Y * k.Y
    projLen := proj.X * proj.X + proj.Y * proj.Y

    val := float64(projLen) / float64(total)

    curr := 0
    for pos, core := range grd.Cores {
        if core.Coord > val {
            curr = pos
            break
        }
    }

    if curr == 0 {
        return grd.Cores[0].Color
    }

    if curr == len(grd.Cores) - 1 {
        return grd.Cores[len(grd.Cores) - 1].Color
    }

    r1, g1, b1, a1 := grd.Cores[curr - 1].Color.RGBA()
    p1 := utils.Pixel{r1, g1, b1, a1}
    r2, g2, b2, a2 := grd.Cores[curr].Color.RGBA()
    p2 := utils.Pixel{r2, g2, b2, a2}

    px := utils.Pixel_linear_interpolation(p1, p2,
        (val - grd.Cores[curr - 1].Coord) / (grd.Cores[curr].Coord - grd.Cores[curr - 1].Coord))

    return color.RGBA{uint8(px.R >> 8), uint8(px.G >> 8), uint8(px.B >> 8), uint8(px.A >> 8)}
}

/**
 *  Creates a gradient image given the size and color map.
 */
func ApplyGradient(bounds image.Rectangle, grd Gradient) (image.Image) {
    ret := image.Image(image.NewRGBA(bounds))

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y < bounds.Max.Y; y += n {
                for x := bounds.Min.X; x < bounds.Max.X; x ++ {
                    ret.(draw.Image).Set(x, y, grd.DrawPixel(image.Point{x, y}, bounds))
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
