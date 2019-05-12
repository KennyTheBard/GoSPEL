package krom

import (
    "os"
    "io/ioutil"
    "reflect"
    generics "./generics"
    error "./error"
)

/**
 *  Structure for holding filenames.
 */
type Files struct {
    paths []string
}

/**
 *  Add a new file to the list.
 */
func (fs *Files) Import(filepath string) {
    fs.paths = append(fs.paths, filepath)
}

/**
 *  Check if a file is already imported.
 */
func (fs Files) isImported(filepath string) bool {
    for _, path := range fs.paths {
        if filepath == path {
            return true
        }
    }
    return false
}

var ImportedFiles Files

/**
 *  Imports a file by filepath and execute all the code inside.
 *  Should be used mainly for bringing defines into global space.
 *  Usage: import <filepath>
 */
func Import(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the file name
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    filepath, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }

    // in case there might be a cyclic import
    if ImportedFiles.isImported(filepath) {
        return nil, error.CreateNoError()
    }

    // attempt to open the file
    file, err_open := os.Open(filepath)
    if err_open != nil {
        return nil, error.CreateError(error.FileError,
            "The file " + filepath + "could not be imported!")
    }
    defer file.Close()

    // add the file to imported list
    (&ImportedFiles).Import(filepath)

    // parse and interpret
    bs, _ := ioutil.ReadAll(file)
    script := string(bs)
    forrest := BuildForrest(script)
    return nil, ExecuteAll(forrest, []generics.Void{})
}
