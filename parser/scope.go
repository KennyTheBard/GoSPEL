package parser

import (
    "image"
)

type Scope struct {
    Images map[string]image.Image
    Data map[string]string
}
