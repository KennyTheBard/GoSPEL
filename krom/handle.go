package krom

import (
    op "./operation"
    ctrl "./control"
    generics "./generics"
    macro "./macro"
)

/**
 *  Return the handle for the required function.
 */
func GetHandle(process string) (generics.Void) {
    switch process {
    case "copy":
        return op.CopyHandle
    case "load":
        return op.LoadHandle
    case "save":
        return op.SaveHandle
    case "point":
        return op.PointHandle
    case "rect":
        return op.RectangleHandle
    case "resize":
        return op.ResizeHandle
    case "rotate":
        return op.RotateHandle
    case "gen":
        return op.GeneratorHandle
    case "sizeof":
        return op.SizeofHandle
    case "filter":
        return op.ApplyFilterHandle
    case "modify":
        return op.ModifyColorsHandle
    case "merge":
        return op.MergeHandle
    case "crop":
        return op.CropHandle
    case "try":
        return ctrl.TryHandle
    case "print":
        return ctrl.PrintHandle
    case "let":
        return ctrl.LetHandle
    case "define":
        return ctrl.DefineHandle
    default:
        return macro.Macros.GetMacro(process)
    }
}
