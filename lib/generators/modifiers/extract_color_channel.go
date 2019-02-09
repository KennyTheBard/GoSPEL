package modifiers

import (
    lib "../.."
)

const (
    RED_CHANNEL = 0
    GREEN_CHANNEL = 1
    BLUE_CHANNEL = 2
)

func ExctractColorChannel(channel int) (lib.Modifier) {

    switch channel {
    case RED_CHANNEL:
        return lib.Modifier{[4][4]float64{[4]float64{1, 0, 0, 0}, [4]float64{0, 0, 0, 0}, [4]float64{0, 0, 0, 0}, [4]float64{0, 0, 0, 0}}, [4]float64{0, 0, 0, 0}}
    case GREEN_CHANNEL:
        return lib.Modifier{[4][4]float64{[4]float64{0, 0, 0, 0}, [4]float64{0, 1, 0, 0}, [4]float64{0, 0, 0, 0}, [4]float64{0, 0, 0, 0}}, [4]float64{0, 0, 0, 0}}
    case BLUE_CHANNEL:
        return lib.Modifier{[4][4]float64{[4]float64{0, 0, 0, 0}, [4]float64{0, 0, 0, 0}, [4]float64{0, 0, 1, 0}, [4]float64{0, 0, 0, 0}}, [4]float64{0, 0, 0, 0}}
    default:
        return lib.Modifier{[4][4]float64{[4]float64{0, 0, 0, 0}, [4]float64{0, 0, 0, 0}, [4]float64{0, 0, 0, 0}, [4]float64{0, 0, 0, 0}}, [4]float64{0, 0, 0, 0}}
    }
}
