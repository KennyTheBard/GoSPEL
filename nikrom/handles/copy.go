package handle

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

func CopyHandle(argc int, argv []generics.Void) (generics.Void, error.Error) {
    if argc != 1 {
        return (nil, error.Error(error.InvalidNumberOfArguments, "Expected 1, received " + argc + "!"))
    }

    if reflect.TypeOf(argv[1]) != reflect.TypeOf(image.Image{}) {
        return (nil, error.Error(error.InvalidArgumentType, "Expected image.Image, received " + reflect.TypeOf(argv[1]) + "!"))
    }

    return (generics.Void) lib.Copy(reflect.ValueOf(argv[i]));
}
