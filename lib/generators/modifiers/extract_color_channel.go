package modifiers

import (
    lib "../.."
)

const (
    RED_CHANNEL = 0
    GREEN_CHANNEL = 1
    BLUE_CHANNEL = 2
    ALPHA_CHANNEL = 3
)

func ExctractColorChannel(channel int) (lib.Modifier) {

    ignored := [4]float64{0, 0, 0, 0}

    switch channel {
    case RED_CHANNEL:
        return lib.Modifier{
                [4][4]float64{
                    [4]float64{1, 0, 0, 0},
                    ignored,
                    ignored,
                    ignored},
                ignored}
    case GREEN_CHANNEL:
        return lib.Modifier{
            [4][4]float64{
                    ignored,
                    [4]float64{0, 1, 0, 0},
                    ignored,
                    ignored},
                ignored}
    case BLUE_CHANNEL:
        return lib.Modifier{
            [4][4]float64{
                    ignored,
                    ignored,
                    [4]float64{0, 0, 1, 0},
                    ignored},
                ignored}
    default:
        return lib.Modifier{
            [4][4]float64{
                    ignored,
                    ignored,
                    ignored,
                    [4]float64{0, 0, 0, 1}},
                ignored}
    }
}
