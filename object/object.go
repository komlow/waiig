package object

import (
    "fmt"
    "bytes"
    "strings"
    "github.com/komlow/writing-an-interpreter-in-go/ast"
    "hash/fnv"
)

const (
    INTEGER_OBJ      = "INTEGER"
    BOOLEAN_OBJ      = "BOOLEAN"
    STRING_OBJ       = "STRING"
    NULL_OBJ         = "NULL"
    RETURN_VALUE_OBJ = "RETURN_VALUE"
    ERROR_OBJ        = "ERROR"
    FUNCTION_OBJ     = "FUNCTION"
    BUILTIN_OBJ      = "BUILTIN"
    ARRAY_OBJ        = "ARRAY"
    HASH_OBJ         = "HASH"
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

type String struct {
    Value string
}

type Null struct {}

type ReturnValue struct {
    Value Object
}

type Error struct {
    Message string
}

type Function struct {
    Parameters []*ast.Identifier
    Body       *ast.BlockStatement
    Env        *Environment
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
    Fn BuiltinFunction
}

type Array struct {
    Elements []Object
}

type HashKey struct {
    Type ObjectType
    Value uint64
}

type HashPair struct {
    Key   Object
    Value Object
}

type Hash struct {
    Pairs map[HashKey]HashPair
}

type Hashable interface {
    HashKey() HashKey
}

func (i *Integer) Inspect() string {
    return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
    return INTEGER_OBJ
}

func (i *Integer) HashKey() HashKey {
    return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (b *Boolean) Inspect() string {
    return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
    return BOOLEAN_OBJ
}

func (b *Boolean) HashKey() HashKey {
    var value uint64

    if b.Value {
        value = 1
    } else {
        value = 0
    }

    return HashKey{Type: b.Type(), Value: value}
}

func (s *String) Inspect() string {
    return s.Value
}

func (s *String) Type() ObjectType {
    return STRING_OBJ
}

func (s *String) HashKey() HashKey {
    h := fnv.New64a()
    h.Write([]byte(s.Value))

    return HashKey{Type: s.Type(), Value: h.Sum64()}
}

func (n *Null) Inspect() string {
    return "null"
}

func (n *Null) Type() ObjectType {
    return NULL_OBJ
}

func (rv *ReturnValue) Type() ObjectType {
    return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
    return rv.Value.Inspect()
}

func (e *Error) Type() ObjectType {
    return ERROR_OBJ
}

func (e *Error) Inspect() string {
    return "ERROR: " + e.Message
}

func (f *Function) Type() ObjectType {
    return FUNCTION_OBJ
}

func (f *Function) Inspect() string {
    var out bytes.Buffer

    params := []string{}
    for _, p := range f.Parameters {
        params = append(params, p.String())
    }

    out.WriteString("fn")
    out.WriteString("(")
    out.WriteString(strings.Join(params, ", "))
    out.WriteString(") {\n")
    out.WriteString(f.Body.String())
    out.WriteString("\n}")

    return out.String()
}

func (b *Builtin) Type() ObjectType {
    return BUILTIN_OBJ
}

func (b *Builtin) Inspect() string {
    return "builtin function"
}

func (ao *Array) Type() ObjectType {
    return ARRAY_OBJ
}

func (ao *Array) Inspect() string {
    var out bytes.Buffer

    elements := []string{}
    for _, e := range ao.Elements {
        elements = append(elements, e.Inspect())
    }

    out.WriteString("[")
    out.WriteString(strings.Join(elements, ", "))
    out.WriteString("]")

    return out.String()
}

func (h *Hash) Type() ObjectType {
    return HASH_OBJ
}

func (h *Hash) Inspect() string {
    var out bytes.Buffer

    pairs := []string{}
    for _, pair := range h.Pairs {
        pairs = append(pairs, fmt.Sprintf("%s: %s",
            pair.Key.Inspect(), pair.Value.Inspect()))
    }

    out.WriteString("{")
    out.WriteString(strings.Join(pairs, ", "))
    out.WriteString("}")

    return out.String()
}

