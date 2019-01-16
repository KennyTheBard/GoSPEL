package parser

import (
    "fmt"
    "strings"
)

func FindBracePair(str string, begin int) (int, int) {
    brace := 0
    start := -1
    end := -1

    for i := begin; i < len(str); i++ {
        if str[i] == '{' {
            brace ++
            if start < 0 {
                start = i
            }
        }

        if str[i] == '}' {
            brace --
            if brace == 0 {
                end = i
                break
            }
        }
    }

    return start, end
}

func NextCommandEnd(str string, begin int) int {
    brace := 0

    for i := begin; i < len(str); i++ {
        fmt.Print("")
        if str[i] == '{' {
            brace ++
        }

        if str[i] == '}' {
            brace --
        }

        if brace == 0 && str[i] == ',' {
            return i
        }
    }

    return len(str)
}

func CountCommands(str string) int {
    brace := 0
    count := 0

    if len(str) > 0 && str[0] != '{' {
        count ++
    }

    for _, ch := range str {
        if ch == '{' {
            brace ++
        }

        if ch == '}' {
            brace --
        }

        if brace == 0 && ch == ',' {
            count ++
        }
    }

    return count
}

func CountBracesPairs(str string) int {
    brace := 0
    count := 0

    for _, ch := range str {
        if ch == '{' {
            brace ++
        }

        if ch == '}' {
            brace --
            if brace == 0 {
                count ++
            }
        }
    }

    return count
}

func StandardizeString(str string) string {
    return strings.Join(strings.Fields(str), " ")
}
