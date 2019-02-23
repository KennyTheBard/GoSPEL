package transformations

import (
    "math"
    "image"
    lib "../.."
)

func SwirlFunc(factor float64) (lib.TransformFunction) {

    retFunc := func (x, y int, bounds image.Rectangle) (int, int) {
        mx := (bounds.Max.X + bounds.Min.X) / 2
        my := (bounds.Max.Y + bounds.Min.Y) / 2

        dx := x - mx
        dy := my - y

        var orig_angle float64

        if dx != 0 {
            orig_angle = math.Atan(math.Abs(float64(dy))/math.Abs(float64(dx)))
            if dx > 0 && dy < 0 {
                orig_angle = 2 * math.Pi - orig_angle
            } else if dx <= 0 && dy >= 0 {
                orig_angle = math.Pi - orig_angle
            } else if dx <= 0 && dy < 0 {
                orig_angle += math.Pi
            }
        } else {
            if dy >= 0 {
                orig_angle = 0.5 * math.Pi
            } else {
                orig_angle = 1.5 * math.Pi
            }
        }

        radius := math.Sqrt(float64(dx) * float64(dx) + float64(dy) * float64(dy))
        angle := orig_angle + factor * radius // 1 / (factor * radius + (4 / math.Pi))
        anglecos := math.Cos(angle)
        anglesin := math.Sin(angle)

        sx := int(math.Floor(radius * anglecos + 0.5))
        sy := int(math.Floor(radius * anglesin + 0.5))

		sx += mx
        sy += my

		sy = bounds.Max.Y - sy;

		// // Clamp the source to legal image pixel
		// if (srcX < 0) srcX = 0;
        //
		// else if (srcX >= width) srcX = width-1;
        //
		// if (srcY < 0) srcY = 0;
        //
		// else if (srcY >= height) srcY = height-1;



        // tx := int(math.Floor(float64(dx) * anglecos - float64(dy) * anglesin)) + mx
        // ty := int(math.Floor(float64(dx) * anglesin + float64(dy) * anglecos)) + my

        return sx, sy
    }

    return retFunc
}
