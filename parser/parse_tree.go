package parser

import (
    "fmt"
    "strings"
    "strconv"
)

type ParseTree struct {
    Command string
    Children []*ParseTree
}

type ParseTreeInterface interface {
    Parse(str string) *ParseTreeInterface
    Resolve(scp Scope)
}

func (p *ParseTree) Parse(str string) *ParseTree {
    num_commands := CountCommands(str)

    if num_commands > 1 {
        p.Children = make([]*ParseTree, num_commands)

        prev := -1
        var curr int
        for i := 0; i < num_commands; i ++ {
            curr = NextCommandEnd(str, prev + 1)
            p.Children[i] = p.Children[i].Parse(str[prev + 1 : curr])
            prev = curr
        }

    } else {
        braces := CountBracesPairs(str)

        if braces != 0 {
            p.Children = make([]*ParseTree, braces)
            curr := 0

            start, end := FindBracePair(str, 0)
            for start > 0 && end > 0 {
                p.Children[curr] = p.Children[curr].Parse(str[start + 1 : end])
                replacer_name := strings.Join([]string{"__var", strconv.Itoa(curr), "__"},"")
                str = strings.Join([]string{ str[: start - 1], replacer_name, str[end + 1 :]}, "")

                start, end = FindBracePair(str, 0)
                curr++
            }
        }

        p.Command = StandardizeString(str)

    }

    return p
}

func (p *ParseTree) Resolve(scp Scope) {
    if len(p.Command) > 0 {
        fmt.Println(p.Command)
    } else {
        for _, child := range p.Children {
            child.Resolve(scp)
        }
    }
}

// strings.NewReplacer("{", " { ", "}", " } ", ";", " ; ", "\"", "")
