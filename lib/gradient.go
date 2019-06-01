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
    Coord float64
    Color color.Color
}

type LinearGradient struct {
    Angle float64
    Cores []ColorCore
}

func (grd LinearGradient) DrawPixel(p image.Point, bounds image.Rectangle) (color.Color) {
    k := image.Point{bounds.Max.X,
        int32(math.Round(bounds.Max.X * math.Tan(grd.Angle)))}
    dotProd := p.X * k.X + p.Y * k.X
    len := p.X * p.X + p.Y * p.Y
    proj := image.Point{int32((dotProd * p.X) / len), int32((dotProd * p.Y) / len)}

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

// UNUSED
func ClosestPoints(gmap GradientMap, center image.Point) []ColorCore {
    aux := make([]ColorCore, len(gmap.Cores))
    for i, core := range gmap.Cores {
        aux[i] = core
    }

    for i := 0; i < len(aux); i++ {
        for j := i + 1; j < len(aux); j++ {
            if utils.Distance(center, aux[i].Point) > utils.Distance(center, aux[j].Point) {
                swap_aux := aux[i]
                aux[i] = aux[j]
                aux[j] = swap_aux
            }
        }
    }

    last := 3
    if len(aux) < 3 {
        last = len(aux)
    }

    return aux[:last]
}

/**
 *  Creates a gradient image given the size and color map.
 */
func Gradient(bounds image.Rectangle, gmap GradientMap) (image.Image) {
    ret := image.Image(image.NewRGBA(bounds))

    n := 10
    done := make(chan bool, n)

    for p := 0; p < n; p ++ {
        aux_rank := p
        go func() {
            rank := aux_rank

            for y := bounds.Min.Y + rank; y < bounds.Max.Y; y += n {
                for x := bounds.Min.X; x < bounds.Max.X; x ++ {
                    center := image.Point{x, y}
                    close := ClosestPoints(gmap, center)

                    var totalDist float64
                    var dists [3]float64
                    for pos := range close {
                        dists[pos] = utils.MinDistancePointToSegment(center,
                            close[(pos + 1) % 3].Point,  close[(pos + 2) % 3].Point)
                        totalDist += dists[pos]
                    }

                    var r_aux, g_aux, b_aux, a_aux float64
                    for pos, core := range close {
                        weight := dists[pos] / totalDist
                        r, g, b, a := core.Color.RGBA()

                        r_aux += weight * float64(r)
                        g_aux += weight * float64(g)
                        b_aux += weight * float64(b)
                        a_aux += weight * float64(a)
                    }


                    r_fin := uint32(math.Round(r_aux))
                    g_fin := uint32(math.Round(g_aux))
                    b_fin := uint32(math.Round(b_aux))
                    a_fin := uint32(math.Round(a_aux))

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

// /**
//     Create an image with dimensions bounds
//     with a descending gradient build around
//     the values vals given for each point of
//     height in the ys array.
//
//     Points of height are expected to be given
//     in ascending order, from 0 to bounds.Max.Y.
// */
// func LinearGradient(bounds image.Rectangle, ys []int, vals []color.Color) (image.Image) {
//     // early exit case
//     if len(ys) != len(vals) {
//         return CreateImage(bounds, color.RGBA{0, 0, 0, 0})
//     }
//
//     // sort the color points
//     utils.Sort_color_points(ys, vals)
//
//     img := image.Image(image.NewRGBA(bounds))
//
//     // first block
//     r, g, b, a := vals[0].RGBA()
//     for y := bounds.Min.Y; y <= ys[0]; y++ {
//         for x := bounds.Min.X; x < bounds.Max.X; x++ {
//             img.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
//         }
//     }
//
//     // middle blocks
//     n := len(ys) - 1
//     curr := ys[0]
//     var prev int
//
//     // for each block
//     for k := 1; k <= n; k++ {
//         prev = curr
//         curr = ys[k]
//
//         r1, g1, b1, a1 := vals[k - 1].RGBA()
//         r2, g2, b2, a2 := vals[k].RGBA()
//
//         px1 := utils.Pixel{r1, g1, b1, a1}
//         px2 := utils.Pixel{r2, g2, b2, a2}
//
//         for y := prev + 1; y <= curr; y++ {
//             for x := bounds.Min.X; x < bounds.Max.X; x++ {
//                 proc := float64(y - prev) / float64(curr - prev)
//
//                 fin := utils.Pixel_linear_interpolation(px1, px2, proc)
//
//                 img.(draw.Image).Set(x, y, color.RGBA{uint8(fin.R >> 8), uint8(fin.G >> 8), uint8(fin.B >> 8), uint8(fin.A >> 8)})
//             }
//         }
//
//     }
//
//
//     // last block
//     r, g, b, a = vals[n].RGBA()
//     for y := ys[n] + 1; y < bounds.Max.Y; y++ {
//         for x := bounds.Min.X; x < bounds.Max.X; x++ {
//             img.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
//         }
//     }
//
//     return img;
// }
//
//
// /**
//     Returns the image with the dimensions width
//     and height with a descending gradient build
//     around the center of the image with the values
//     vals given for each point of height in the
//     ys array.
//
//     Points of height are expected to be given
//     in ascending order, from 0 to height / 2.
// */
// func CircularGradient(size int, ys []int, vals []color.Color) (image.Image) {
//     // early exit case
//     if len(ys) != len(vals) {
//         return CreateImage(image.Rect(0, 0, size, size), color.RGBA{0, 0, 0, 0})
//     }
//
//     // sort the color points
//     utils.Sort_color_points(ys, vals)
//
//     cx := size / 2
//     cy := size / 2
//     bounds := image.Rect(0, 0, size, size)
//     img := image.Image(image.NewRGBA(bounds))
//
//     n := 10
//     done := make(chan bool, n)
//
//     for p := 0; p < n; p ++ {
//         aux_rank := p
//         go func() {
//             rank := aux_rank
//
//             for y := bounds.Min.Y + rank; y < bounds.Max.Y; y += n {
//                 for x := bounds.Min.X; x < bounds.Max.X; x++ {
//                     // determine the color interval
//                     dst := int(math.Round(utils.Distance(float64(x), float64(y), float64(cx), float64(cy))))
//                     k1, k2 := utils.Search_interval(ys, dst)
//
//                     if k1 == k2 {
//                         r, g, b, a := vals[k1].RGBA()
//                         img.(draw.Image).Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
//                         continue
//                     }
//
//                     // prepare data for interpolation
//                     r1, g1, b1, a1 := vals[k1].RGBA()
//                     r2, g2, b2, a2 := vals[k2].RGBA()
//                     px1 := utils.Pixel{r1, g1, b1, a1}
//                     px2 := utils.Pixel{r2, g2, b2, a2}
//
//                     // interpolation & pixel writing
//                     proc := float64(dst - ys[k1]) / float64(ys[k2] - ys[k1])
//
//                     fin := utils.Pixel_linear_interpolation(px1, px2, proc)
//
//                     //fmt.Println(dst, "intre", ys[k1], "si", ys[k2], "iar proc este", proc, "iar fin.A este", fin.A)
//
//                     img.(draw.Image).Set(x, y, color.RGBA{uint8(fin.R >> 8), uint8(fin.G >> 8), uint8(fin.B >> 8), uint8(fin.A >> 8)})
//                 }
//             }
//
//             done <- true;
//         } ()
//     }
//
//     for i := 0; i < n; i++ {
//         <-done
//     }
//
//     return img;
// }
