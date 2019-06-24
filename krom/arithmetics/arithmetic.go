package arithmetic

import (
    "reflect"
    "strconv"
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and evaluate a Lisp-style
 *  aritmethic expresion.
 *  Usage: define <name> (<code>)
 */
func ArithmeticHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected >= received {
        return nil, error.NumberArgumentsErrorAtLeast(expected, received)
    }

    // extract the equations
    var stack []float64
    var first, second float64
    for pos := len(raw_args) - 1; pos > 0; pos -= 1 {
        args[pos] = raw_args[pos].(generics.InterpreterTree)
        aux, err := args[pos].Interpret(scope.Clone())
        if err.Code != error.NoError {
            return nil, err
        }
        str, ok := aux.(string)
        if !ok {
            return nil, error.ArgumentTypeError(pos, "number as a string", reflect.TypeOf(aux).Name())
        }
        x, conv_err := strconv.ParseFloat(str, 64)
        if conv_err == nil {
            stack = append(stack, x)
        } else {
            first = stack[len(stack) - 2]
            second = stack[len(stack) - 1]
            stack = stack[:len(stack) - 2]

            stack = append(stack, executeOperation(str, first, second))
        }
    }

    // call the operation
    return strconv.FormatFloat(stack[0], 'E', -1, 64), error.CreateNoError()
}

/**
 *  Auxiliar function to calculate the result of the
 *  given operation on the given numbers.
 */
func executeOperation(op string, first, second float64) (float64) {
    switch op {
    case "+":
        return first + second
    case "-":
        return first - second
    case "*":
        return first * second
    case "/":
        return first / second
    default:
        return first
    }
}
