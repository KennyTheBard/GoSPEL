package krom

import (
    error "./error"
    handles "./handles"
    generics "./generics"
)

/**
 *  Return the handle for the required function.
 */
func GetHandle(process string) (generics.Handle, Evaluation, error.Error) {
    switch process {
    case "copy":
        return handles.CopyHandle, normalEvaluation, error.CreateNoError()
    case "load":
        return handles.LoadHandle, normalEvaluation, error.CreateNoError()
    case "save":
        return handles.SaveHandle, normalEvaluation, error.CreateNoError()
    case "point":
        return handles.PointHandle, normalEvaluation, error.CreateNoError()
    case "rect":
        return handles.RectangleHandle, normalEvaluation, error.CreateNoError()
    case "resize":
        return handles.ResizeHandle, normalEvaluation, error.CreateNoError()
    case "rotate":
        return handles.RotateHandle, normalEvaluation, error.CreateNoError()
    case "gen":
        return handles.GeneratorHandle, normalEvaluation, error.CreateNoError()
    case "sizeof":
        return handles.SizeofHandle, normalEvaluation, error.CreateNoError()
    case "filter":
        return handles.ApplyFilterHandle, normalEvaluation, error.CreateNoError()
    case "modify":
        return handles.ModifyColorsHandle, normalEvaluation, error.CreateNoError()
    case "merge":
        return handles.MergeHandle, normalEvaluation, error.CreateNoError()
    case "crop":
        return handles.CropHandle, normalEvaluation, error.CreateNoError()
    default:
        return nil, nil, error.CreateError(error.UnknownHandle,
            "Unknown handle name \"" + process + "\"!")
    }
}
