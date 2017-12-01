package object

import (
    "fmt"
)

const (
    INTEGER_OBJ = "INTEGER"
    BOOLEAN_OBJ = "BOOLEAN"
    NULL_OBJ = "NULL"
)

type ObjectType string

type Object interface {
    Type() ObjectType
    Inspect() string
}

type Integer struct {
    Value int64
}

type Boolean struct {
    Value bool
}

type NULL struct {}

func (i *Integer) Inspect() string {
    return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
    INTEGER_OBJ
}

func (b *Boolean) Inspect() string {
    return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
    BOOLEAN_OBJ
}

func (n *Null) Inspect() string {
    return "null"
}

func (n *Null) Type() ObjectType {
    NULL_OBJ
}
