package handle

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

func CopyHandle(args []generics.Void) (generics.Void, error.Error) {
    if len(args) != 1 {
        return (nil, error.Error(error.InvalidNumberOfArguments, "Expected 1, received " + len(args) + "!"))
    }

    if reflect.TypeOf(args[0]) != reflect.TypeOf(image.Image{}) {
        return (nil, error.Error(error.InvalidArgumentType, "Expected image.Image, received " + reflect.TypeOf(args[0]) + "!"))
    }

    return lib.Copy(reflect.ValueOf(args[0]));
}
