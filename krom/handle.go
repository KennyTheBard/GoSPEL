package krom

import (
    error "./error"
    op "./operation"
    ctrl "./control"
    generics "./generics"
)

/**
 *  Return the handle for the required function.
 */
func GetHandle(process string) (generics.Handle, error.Error) {
    switch process {
    case "copy":
        return op.CopyHandle, error.CreateNoError()
    case "load":
        return op.LoadHandle, error.CreateNoError()
    case "save":
        return op.SaveHandle, error.CreateNoError()
    case "point":
        return op.PointHandle, error.CreateNoError()
    case "rect":
        return op.RectangleHandle, error.CreateNoError()
    case "resize":
        return op.ResizeHandle, error.CreateNoError()
    case "rotate":
        return op.RotateHandle, error.CreateNoError()
    case "gen":
        return op.GeneratorHandle, error.CreateNoError()
    case "sizeof":
        return op.SizeofHandle, error.CreateNoError()
    case "filter":
        return op.ApplyFilterHandle, error.CreateNoError()
    case "modify":
        return op.ModifyColorsHandle, error.CreateNoError()
    case "merge":
        return op.MergeHandle, error.CreateNoError()
    case "crop":
        return op.CropHandle, error.CreateNoError()
    case "try":
        return ctrl.TryHandle, error.CreateNoError()
    case "print":
        return ctrl.PrintHandle, error.CreateNoError()
    case "let":
        return ctrl.LetHandle, error.CreateNoError()
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown handle name \"" + process + "\"!")
    }
}
