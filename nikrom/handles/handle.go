package handles

import (
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
func GetHandle(process string) (handle, error.Error) {
    switch process {
    case "copy":
        return CopyHandle, error.CreateNoError()
    case "load":
        return LoadHandle, error.CreateNoError()
    case "save":
        return SaveHandle, error.CreateNoError()
    case "point":
        return PointHandle, error.CreateNoError()
    case "rect":
        return RectangleHandle, error.CreateNoError()
    case "resize":
        return ResizeHandle, error.CreateNoError()
    case "rotate":
        return RotateHandle, error.CreateNoError()
    case "gen":
        return GeneratorHandle, error.CreateNoError()
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown handle name \"" + process + "\"!")
    }
}
