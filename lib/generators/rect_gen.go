package tools

import (
    "image"
)

func CreateRectangleByPoints(start_x, start_y, end_x, end_y int) (image.Rectangle) {
    return image.Rect(start_x, start_y, end_x, end_y)
}

func CreateRectangleBySize(width, height int) (image.Rectangle) {
    return image.Rect(0, 0, width, height)
}

func CreateSquare(size int) (image.Rectangle) {
    return image.Rect(0, 0, size, size)
}
