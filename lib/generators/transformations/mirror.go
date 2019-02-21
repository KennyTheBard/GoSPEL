package transformations

import (
    "image"
    lib "../.."
)

const (
    HORIZONTAL_MODE = 0
    VERTICAL_MODE = 1
)

func MirrorFunction(mode int) (lib.TransformFunction) {
    var retFunc lib.TransformFunction
    switch mode {
        case HORIZONTAL_MODE:
            retFunc = func (x, y int, bounds image.Rectangle) (int, int) {
                return x, bounds.Max.Y - (y - bounds.Min.Y)
            }
            break

        case VERTICAL_MODE:
            retFunc = func (x, y int, bounds image.Rectangle) (int, int) {
                return bounds.Max.X - (x - bounds.Min.X), y
            }
            break

        default:
            retFunc = func (x, y int, bounds image.Rectangle) (int, int) {
                return x, y
            }
    }

    return retFunc
}
