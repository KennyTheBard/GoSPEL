package handle

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

/**
 *  Declare a type for all handles.
 */
type handle func([]generics.Void) (generics.Void, error.Error)

/**
 *  Return the handle for the required function.
 */
func GetHandle(tree generics.Atom) (handle, error.Error) {
    switch tree.Process {
        "copy":
            return (CopyHandle, error.Error)
        default:
            return (nil, error.Error(error.UnknownHandle, "Unknown handle name \"" + tree.Process + "\"!"))

    }
}
