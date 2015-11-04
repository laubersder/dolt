// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("Struct",
			[]types.Field{
				types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_CachedRef = types.RegisterPackage(&p)
}

// Struct

type Struct struct {
	_s string
	_b bool

	ref *ref.Ref
}

func NewStruct() Struct {
	return Struct{
		_s: "",
		_b: false,

		ref: &ref.Ref{},
	}
}

type StructDef struct {
	S string
	B bool
}

func (def StructDef) New() Struct {
	return Struct{
		_s:  def.S,
		_b:  def.B,
		ref: &ref.Ref{},
	}
}

func (s Struct) Def() (d StructDef) {
	d.S = s._s
	d.B = s._b
	return
}

var __typeRefForStruct types.TypeRef

func (m Struct) TypeRef() types.TypeRef {
	return __typeRefForStruct
}

func init() {
	__typeRefForStruct = types.MakeTypeRef(__genPackageInFile_struct_CachedRef, 0)
	types.RegisterStruct(__typeRefForStruct, builderForStruct, readerForStruct)
}

func builderForStruct() chan types.Value {
	c := make(chan types.Value)
	go func() {
		s := Struct{ref: &ref.Ref{}}
		s._s = (<-c).(types.String).String()
		s._b = bool((<-c).(types.Bool))
		c <- s
	}()
	return c
}

func readerForStruct(v types.Value) chan types.Value {
	c := make(chan types.Value)
	go func() {
		s := v.(Struct)
		c <- types.NewString(s._s)
		c <- types.Bool(s._b)
	}()
	return c
}

func (s Struct) Equals(other types.Value) bool {
	return other != nil && __typeRefForStruct.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s Struct) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Struct) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForStruct.Chunks()...)
	return
}

func (s Struct) S() string {
	return s._s
}

func (s Struct) SetS(val string) Struct {
	s._s = val
	s.ref = &ref.Ref{}
	return s
}

func (s Struct) B() bool {
	return s._b
}

func (s Struct) SetB(val bool) Struct {
	s._b = val
	s.ref = &ref.Ref{}
	return s
}

// ListOfStruct

type ListOfStruct struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfStruct() ListOfStruct {
	return ListOfStruct{types.NewList(), &ref.Ref{}}
}

type ListOfStructDef []StructDef

func (def ListOfStructDef) New() ListOfStruct {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New()
	}
	return ListOfStruct{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfStruct) Def() ListOfStructDef {
	d := make([]StructDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(Struct).Def()
	}
	return d
}

func (l ListOfStruct) InternalImplementation() types.List {
	return l.l
}

func (l ListOfStruct) Equals(other types.Value) bool {
	return other != nil && __typeRefForListOfStruct.Equals(other.TypeRef()) && l.Ref() == other.Ref()
}

func (l ListOfStruct) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfStruct) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.TypeRef().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfStruct.
var __typeRefForListOfStruct types.TypeRef

func (m ListOfStruct) TypeRef() types.TypeRef {
	return __typeRefForListOfStruct
}

func init() {
	__typeRefForListOfStruct = types.MakeCompoundTypeRef(types.ListKind, types.MakeTypeRef(__genPackageInFile_struct_CachedRef, 0))
	types.RegisterFromValFunction(__typeRefForListOfStruct, func(v types.Value) types.Value {
		return ListOfStruct{v.(types.List), &ref.Ref{}}
	})
}

func (l ListOfStruct) Len() uint64 {
	return l.l.Len()
}

func (l ListOfStruct) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfStruct) Get(i uint64) Struct {
	return l.l.Get(i).(Struct)
}

func (l ListOfStruct) Slice(idx uint64, end uint64) ListOfStruct {
	return ListOfStruct{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfStruct) Set(i uint64, val Struct) ListOfStruct {
	return ListOfStruct{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfStruct) Append(v ...Struct) ListOfStruct {
	return ListOfStruct{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfStruct) Insert(idx uint64, v ...Struct) ListOfStruct {
	return ListOfStruct{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfStruct) Remove(idx uint64, end uint64) ListOfStruct {
	return ListOfStruct{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfStruct) RemoveAt(idx uint64) ListOfStruct {
	return ListOfStruct{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfStruct) fromElemSlice(p []Struct) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfStructIterCallback func(v Struct, i uint64) (stop bool)

func (l ListOfStruct) Iter(cb ListOfStructIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(Struct), i)
	})
}

type ListOfStructIterAllCallback func(v Struct, i uint64)

func (l ListOfStruct) IterAll(cb ListOfStructIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(Struct), i)
	})
}

type ListOfStructFilterCallback func(v Struct, i uint64) (keep bool)

func (l ListOfStruct) Filter(cb ListOfStructFilterCallback) ListOfStruct {
	nl := NewListOfStruct()
	l.IterAll(func(v Struct, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}
