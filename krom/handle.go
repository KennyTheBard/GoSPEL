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
        return handles.CopyHandle, reverseEvaluation, error.CreateNoError()
    case "load":
        return handles.LoadHandle, reverseEvaluation, error.CreateNoError()
    case "save":
        return handles.SaveHandle, reverseEvaluation, error.CreateNoError()
    case "point":
        return handles.PointHandle, reverseEvaluation, error.CreateNoError()
    case "rect":
        return handles.RectangleHandle, reverseEvaluation, error.CreateNoError()
    case "resize":
        return handles.ResizeHandle, reverseEvaluation, error.CreateNoError()
    case "rotate":
        return handles.RotateHandle, reverseEvaluation, error.CreateNoError()
    case "gen":
        return handles.GeneratorHandle, reverseEvaluation, error.CreateNoError()
    case "sizeof":
        return handles.SizeofHandle, reverseEvaluation, error.CreateNoError()
    case "filter":
        return handles.ApplyFilterHandle, reverseEvaluation, error.CreateNoError()
    case "modify":
        return handles.ModifyColorsHandle, reverseEvaluation, error.CreateNoError()
    case "merge":
        return handles.MergeHandle, reverseEvaluation, error.CreateNoError()
    case "crop":
        return handles.CropHandle, reverseEvaluation, error.CreateNoError()
    default:
        return nil, nil, error.CreateError(error.UnknownHandle,
            "Unknown handle name \"" + process + "\"!")
    }
}
