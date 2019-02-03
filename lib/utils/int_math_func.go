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
