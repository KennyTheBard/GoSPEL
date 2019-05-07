package krom

import (
    error "./error"
    handles "./handles"
    generics "./generics"
)

/**
 *  Return the handle for the required function.
 */
func GetHandle(process string) (generics.Handle, error.Error) {
    switch process {
    case "copy":
        return handles.CopyHandle, error.CreateNoError()
    case "load":
        return handles.LoadHandle, error.CreateNoError()
    case "save":
        return handles.SaveHandle, error.CreateNoError()
    case "point":
        return handles.PointHandle, error.CreateNoError()
    case "rect":
        return handles.RectangleHandle, error.CreateNoError()
    case "resize":
        return handles.ResizeHandle, error.CreateNoError()
    case "rotate":
        return handles.RotateHandle, error.CreateNoError()
    case "gen":
        return handles.GeneratorHandle, error.CreateNoError()
    case "sizeof":
        return handles.SizeofHandle, error.CreateNoError()
    case "filter":
        return handles.ApplyFilterHandle, error.CreateNoError()
    case "modify":
        return handles.ModifyColorsHandle, error.CreateNoError()
    case "merge":
        return handles.MergeHandle, error.CreateNoError()
    case "crop":
        return handles.CropHandle, error.CreateNoError()
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown handle name \"" + process + "\"!")
    }
}
