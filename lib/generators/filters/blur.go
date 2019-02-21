package filters

import (
    //"math"
    lib "../.."
)

func BoxBlur(size int) (lib.Filter) {
    mat := make([][]float64, size)
    for i := range mat {
        mat[i] = make([]float64, size)
    }

    coef := float64(size * size)
    for i := range mat {
        for j := range mat[i] {
            mat[i][j] = 1.0 / coef
        }
    }

    return lib.Filter{mat}
}

// func GaussianFilter(width, height int, xdev, ydev, theta float64) (lib.Filter) {
//     r := [][]float64{[]float64{math.Cos(theta), -math.Sin(theta)}, []float64{math.Sin(theta), math.Cos(theta)}}
//
//     mat := make([][]float64, height)
//     for i := range mat {
//         mat[i] = make([]float64, width)
//     }
//
//     for i := 0; i < height; i++ {
//         for j := 0; j < width; j++ {
//             a := float64(j - (width + 1) / 2)
//             b := float64(i - (height + 1) / 2)
//             u := []float64{r[0][0] * a + r[0][1] * b, r[1][0] * a + r[1][1] * b}
//             mat[i][j] = gauss(u[0], xdev) * gauss(u[1], ydev)
//         }
//     }
//
//     sum := 0.0
//     for i := 0; i < height; i++ {
//         for j := 0; j < width; j++ {
//             sum += mat[i][j] * mat[i][j]
//         }
//     }
//
//     for i := 0; i < height; i++ {
//         for j := 0; j < width; j++ {
//             mat[i][j] = mat[i][j] / math.Sqrt(sum)
//         }
//     }
//
//     return lib.Filter{mat}
// }
//
// func gauss(val float64, dev float64) (float64) {
//     e := math.Exp(-math.Pow(val, 2) / (2 * math.Pow(dev, 2)))
//     b := dev * math.Sqrt(2 * math.Pi)
//     return e / b
// }

func AxialBlur(radius, mode int) (lib.Filter) {
    size := radius * 2 + 1
    var mat [][]float64

    switch mode {
    case lib.HORIZONTAL_MODE:
        mat = make([][]float64, 1)
        for i := range mat {
            mat[i] = make([]float64, size)
        }
        break
    case lib.VERTICAL_MODE:
        mat = make([][]float64, size)
        for i := range mat {
            mat[i] = make([]float64, 1)
        }
        break
    default:
        mat = make([][]float64, 1)
        mat[0] = make([]float64, 1)
    }

    coef := float64(size)
    for i := range mat {
        for j := range mat[i] {
            mat[i][j] = 1.0 / coef
        }
    }

    return lib.Filter{mat}
}
