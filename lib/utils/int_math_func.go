package utils


/**
	Returns the absolute value of the input.
*/
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


/**
	Returns the minimum value between the 2 inputs.
*/
func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}


/**
	Returns the minimum value between the 2 inputs.
*/
func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

/**
    Force the value val in [min, max] for int32.
*/
func Clamp(min, max, val int32) int32 {
    if val <= min {
        return min
    } else if val >= max {
        return max
    } else {
        return val
    }
}

/**
    Force the value val in [min, max] for uint32.
*/
func Uclamp(min, max, val uint32) uint32 {
    if val <= min {
        return min
    } else if val >= max {
        return max
    } else {
        return val
    }
}
