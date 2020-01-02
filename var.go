package main

import (
	"github.com/ezbuy/expvarmon/pkg/varex"
)

// VarName represents variable name.
//
// It has dot-separated format, like "memstats.Alloc",
// but can be used in different forms, hence it's own type.
//
// It also can have optional "kind:" modifier, like "mem:" or "duration:"
type VarName = varex.VarName

// VarKind specifies special kinds of values, affects formatting.
type VarKind = varex.VarKind

// VarValue represents arbitrary value for variable.
type VarValue = varex.VarValue

const (
	KindDefault  = varex.KindDefault
	KindMemory   = varex.KindMemory
	KindDuration = varex.KindDuration
	KindString   = varex.KindString
)

var Format = varex.Format

var DottedFieldsToSliceEscaped = varex.DottedFieldsToSliceEscaped
