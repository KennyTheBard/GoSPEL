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
func GetHandle(process string) (generics.Handle, Evaluation, error.Error) {
    switch process {
    case "copy":
        return op.CopyHandle, reverseEvaluation, error.CreateNoError()
    case "load":
        return op.LoadHandle, reverseEvaluation, error.CreateNoError()
    case "save":
        return op.SaveHandle, reverseEvaluation, error.CreateNoError()
    case "point":
        return op.PointHandle, reverseEvaluation, error.CreateNoError()
    case "rect":
        return op.RectangleHandle, reverseEvaluation, error.CreateNoError()
    case "resize":
        return op.ResizeHandle, reverseEvaluation, error.CreateNoError()
    case "rotate":
        return op.RotateHandle, reverseEvaluation, error.CreateNoError()
    case "gen":
        return op.GeneratorHandle, reverseEvaluation, error.CreateNoError()
    case "sizeof":
        return op.SizeofHandle, reverseEvaluation, error.CreateNoError()
    case "filter":
        return op.ApplyFilterHandle, reverseEvaluation, error.CreateNoError()
    case "modify":
        return op.ModifyColorsHandle, reverseEvaluation, error.CreateNoError()
    case "merge":
        return op.MergeHandle, reverseEvaluation, error.CreateNoError()
    case "crop":
        return op.CropHandle, reverseEvaluation, error.CreateNoError()
    case "try":
        return ctrl.TryHandle, normalEvaluation, error.CreateNoError()
    case "print":
        return ctrl.PrintHandle, normalEvaluation, error.CreateNoError()
    case "let":
        return ctrl.LetHandle, normalEvaluation, error.CreateNoError()
    default:
        return nil, nil, error.CreateError(error.UnknownHandle,
            "Unknown handle name \"" + process + "\"!")
    }
}
