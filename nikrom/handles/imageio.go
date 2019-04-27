package handle

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

func LoadHandle(args []generics.Void) (generics.Void, error.Error) {
    if len(args) != 1 {
        return (nil, error.Error(error.InvalidNumberOfArguments, "Expected 1, received " + len(args) + "!"))
    }

    if reflect.TypeOf(args[0]) != reflect.TypeOf(string{}) {
        return (nil, error.Error(error.InvalidArgumentType, "Expected argument 1 of type string, received " + reflect.TypeOf(args[0]) + "!"))
    }

    return (lib.DecodeImage(reflect.ValueOf(args[0])), error.NoError())
}

func SaveHandle(args []generics.Void) (generics.Void, error.Error) {
    if len(args) != 3 {
        return (nil, error.Error(error.InvalidNumberOfArguments, "Expected 3, received " + len(args) + "!"))
    }

    if reflect.TypeOf(args[0]) != reflect.TypeOf(image.Image{}) {
        return (nil, error.Error(error.InvalidArgumentType, "Expected argument 1 of type image.Image, received " + reflect.TypeOf(args[0]) + "!"))
    }
    if reflect.TypeOf(args[1]) != reflect.TypeOf(string{}) {
        return (nil, error.Error(error.InvalidArgumentType, "Expected argument 2 of type string, received " + reflect.TypeOf(args[0]) + "!"))
    }
    if reflect.TypeOf(args[2]) != reflect.TypeOf(string{}) {
        return (nil, error.Error(error.InvalidArgumentType, "Expected argument 3 of type string, received " + reflect.TypeOf(args[0]) + "!"))
    }

    return (lib.EncodeImage(reflect.ValueOf(args[0])), error.NoError())
}
