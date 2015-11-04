// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("Pitch",
			[]types.Field{
				types.Field{"X", types.MakePrimitiveTypeRef(types.Float64Kind), false},
				types.Field{"Z", types.MakePrimitiveTypeRef(types.Float64Kind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__mainPackageInFile_types_CachedRef = types.RegisterPackage(&p)
}

// Pitch

type Pitch struct {
	_X float64
	_Z float64

	ref *ref.Ref
}

func NewPitch() Pitch {
	return Pitch{
		_X: float64(0),
		_Z: float64(0),

		ref: &ref.Ref{},
	}
}

type PitchDef struct {
	X float64
	Z float64
}

func (def PitchDef) New() Pitch {
	return Pitch{
		_X:  def.X,
		_Z:  def.Z,
		ref: &ref.Ref{},
	}
}

func (s Pitch) Def() (d PitchDef) {
	d.X = s._X
	d.Z = s._Z
	return
}

var __typeRefForPitch types.TypeRef

func (m Pitch) TypeRef() types.TypeRef {
	return __typeRefForPitch
}

func init() {
	__typeRefForPitch = types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0)
	types.RegisterStruct(__typeRefForPitch, builderForPitch, readerForPitch)
}

func builderForPitch() chan types.Value {
	c := make(chan types.Value)
	go func() {
		s := Pitch{ref: &ref.Ref{}}
		s._X = float64((<-c).(types.Float64))
		s._Z = float64((<-c).(types.Float64))
		c <- s
	}()
	return c
}

func readerForPitch(v types.Value) chan types.Value {
	c := make(chan types.Value)
	go func() {
		s := v.(Pitch)
		c <- types.Float64(s._X)
		c <- types.Float64(s._Z)
	}()
	return c
}

func (s Pitch) Equals(other types.Value) bool {
	return other != nil && __typeRefForPitch.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s Pitch) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Pitch) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForPitch.Chunks()...)
	return
}

func (s Pitch) X() float64 {
	return s._X
}

func (s Pitch) SetX(val float64) Pitch {
	s._X = val
	s.ref = &ref.Ref{}
	return s
}

func (s Pitch) Z() float64 {
	return s._Z
}

func (s Pitch) SetZ(val float64) Pitch {
	s._Z = val
	s.ref = &ref.Ref{}
	return s
}

// ListOfMapOfStringToValue

type ListOfMapOfStringToValue struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfMapOfStringToValue() ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{types.NewList(), &ref.Ref{}}
}

type ListOfMapOfStringToValueDef []MapOfStringToValueDef

func (def ListOfMapOfStringToValueDef) New() ListOfMapOfStringToValue {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New()
	}
	return ListOfMapOfStringToValue{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfMapOfStringToValue) Def() ListOfMapOfStringToValueDef {
	d := make([]MapOfStringToValueDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(MapOfStringToValue).Def()
	}
	return d
}

func (l ListOfMapOfStringToValue) InternalImplementation() types.List {
	return l.l
}

func (l ListOfMapOfStringToValue) Equals(other types.Value) bool {
	return other != nil && __typeRefForListOfMapOfStringToValue.Equals(other.TypeRef()) && l.Ref() == other.Ref()
}

func (l ListOfMapOfStringToValue) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfMapOfStringToValue) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.TypeRef().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfMapOfStringToValue.
var __typeRefForListOfMapOfStringToValue types.TypeRef

func (m ListOfMapOfStringToValue) TypeRef() types.TypeRef {
	return __typeRefForListOfMapOfStringToValue
}

func init() {
	__typeRefForListOfMapOfStringToValue = types.MakeCompoundTypeRef(types.ListKind, types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.ValueKind)))
	types.RegisterFromValFunction(__typeRefForListOfMapOfStringToValue, func(v types.Value) types.Value {
		return ListOfMapOfStringToValue{v.(types.List), &ref.Ref{}}
	})
}

func (l ListOfMapOfStringToValue) Len() uint64 {
	return l.l.Len()
}

func (l ListOfMapOfStringToValue) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfMapOfStringToValue) Get(i uint64) MapOfStringToValue {
	return l.l.Get(i).(MapOfStringToValue)
}

func (l ListOfMapOfStringToValue) Slice(idx uint64, end uint64) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfMapOfStringToValue) Set(i uint64, val MapOfStringToValue) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfMapOfStringToValue) Append(v ...MapOfStringToValue) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfMapOfStringToValue) Insert(idx uint64, v ...MapOfStringToValue) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfMapOfStringToValue) Remove(idx uint64, end uint64) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfMapOfStringToValue) RemoveAt(idx uint64) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfMapOfStringToValue) fromElemSlice(p []MapOfStringToValue) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfMapOfStringToValueIterCallback func(v MapOfStringToValue, i uint64) (stop bool)

func (l ListOfMapOfStringToValue) Iter(cb ListOfMapOfStringToValueIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(MapOfStringToValue), i)
	})
}

type ListOfMapOfStringToValueIterAllCallback func(v MapOfStringToValue, i uint64)

func (l ListOfMapOfStringToValue) IterAll(cb ListOfMapOfStringToValueIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(MapOfStringToValue), i)
	})
}

type ListOfMapOfStringToValueFilterCallback func(v MapOfStringToValue, i uint64) (keep bool)

func (l ListOfMapOfStringToValue) Filter(cb ListOfMapOfStringToValueFilterCallback) ListOfMapOfStringToValue {
	nl := NewListOfMapOfStringToValue()
	l.IterAll(func(v MapOfStringToValue, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// MapOfStringToListOfPitch

type MapOfStringToListOfPitch struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToListOfPitch() MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToListOfPitchDef map[string]ListOfPitchDef

func (def MapOfStringToListOfPitchDef) New() MapOfStringToListOfPitch {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New())
	}
	return MapOfStringToListOfPitch{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToListOfPitch) Def() MapOfStringToListOfPitchDef {
	def := make(map[string]ListOfPitchDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(ListOfPitch).Def()
		return false
	})
	return def
}

func (m MapOfStringToListOfPitch) InternalImplementation() types.Map {
	return m.m
}

func (m MapOfStringToListOfPitch) Equals(other types.Value) bool {
	return other != nil && __typeRefForMapOfStringToListOfPitch.Equals(other.TypeRef()) && m.Ref() == other.Ref()
}

func (m MapOfStringToListOfPitch) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToListOfPitch) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.TypeRef().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToListOfPitch.
var __typeRefForMapOfStringToListOfPitch types.TypeRef

func (m MapOfStringToListOfPitch) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToListOfPitch
}

func init() {
	__typeRefForMapOfStringToListOfPitch = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeCompoundTypeRef(types.ListKind, types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0)))
	types.RegisterFromValFunction(__typeRefForMapOfStringToListOfPitch, func(v types.Value) types.Value {
		return MapOfStringToListOfPitch{v.(types.Map), &ref.Ref{}}
	})
}

func (m MapOfStringToListOfPitch) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToListOfPitch) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToListOfPitch) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToListOfPitch) Get(p string) ListOfPitch {
	return m.m.Get(types.NewString(p)).(ListOfPitch)
}

func (m MapOfStringToListOfPitch) MaybeGet(p string) (ListOfPitch, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewListOfPitch(), false
	}
	return v.(ListOfPitch), ok
}

func (m MapOfStringToListOfPitch) Set(k string, v ListOfPitch) MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToListOfPitch) Remove(p string) MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToListOfPitchIterCallback func(k string, v ListOfPitch) (stop bool)

func (m MapOfStringToListOfPitch) Iter(cb MapOfStringToListOfPitchIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(ListOfPitch))
	})
}

type MapOfStringToListOfPitchIterAllCallback func(k string, v ListOfPitch)

func (m MapOfStringToListOfPitch) IterAll(cb MapOfStringToListOfPitchIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(ListOfPitch))
	})
}

type MapOfStringToListOfPitchFilterCallback func(k string, v ListOfPitch) (keep bool)

func (m MapOfStringToListOfPitch) Filter(cb MapOfStringToListOfPitchFilterCallback) MapOfStringToListOfPitch {
	nm := NewMapOfStringToListOfPitch()
	m.IterAll(func(k string, v ListOfPitch) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// MapOfStringToString

type MapOfStringToString struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToString() MapOfStringToString {
	return MapOfStringToString{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToStringDef map[string]string

func (def MapOfStringToStringDef) New() MapOfStringToString {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), types.NewString(v))
	}
	return MapOfStringToString{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToString) Def() MapOfStringToStringDef {
	def := make(map[string]string)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(types.String).String()
		return false
	})
	return def
}

func (m MapOfStringToString) InternalImplementation() types.Map {
	return m.m
}

func (m MapOfStringToString) Equals(other types.Value) bool {
	return other != nil && __typeRefForMapOfStringToString.Equals(other.TypeRef()) && m.Ref() == other.Ref()
}

func (m MapOfStringToString) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToString) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.TypeRef().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToString.
var __typeRefForMapOfStringToString types.TypeRef

func (m MapOfStringToString) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToString
}

func init() {
	__typeRefForMapOfStringToString = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.StringKind))
	types.RegisterFromValFunction(__typeRefForMapOfStringToString, func(v types.Value) types.Value {
		return MapOfStringToString{v.(types.Map), &ref.Ref{}}
	})
}

func (m MapOfStringToString) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToString) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToString) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToString) Get(p string) string {
	return m.m.Get(types.NewString(p)).(types.String).String()
}

func (m MapOfStringToString) MaybeGet(p string) (string, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return "", false
	}
	return v.(types.String).String(), ok
}

func (m MapOfStringToString) Set(k string, v string) MapOfStringToString {
	return MapOfStringToString{m.m.Set(types.NewString(k), types.NewString(v)), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToString) Remove(p string) MapOfStringToString {
	return MapOfStringToString{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToStringIterCallback func(k string, v string) (stop bool)

func (m MapOfStringToString) Iter(cb MapOfStringToStringIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(types.String).String())
	})
}

type MapOfStringToStringIterAllCallback func(k string, v string)

func (m MapOfStringToString) IterAll(cb MapOfStringToStringIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(types.String).String())
	})
}

type MapOfStringToStringFilterCallback func(k string, v string) (keep bool)

func (m MapOfStringToString) Filter(cb MapOfStringToStringFilterCallback) MapOfStringToString {
	nm := NewMapOfStringToString()
	m.IterAll(func(k string, v string) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// MapOfStringToValue

type MapOfStringToValue struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToValue() MapOfStringToValue {
	return MapOfStringToValue{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToValueDef map[string]types.Value

func (def MapOfStringToValueDef) New() MapOfStringToValue {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v)
	}
	return MapOfStringToValue{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToValue) Def() MapOfStringToValueDef {
	def := make(map[string]types.Value)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v
		return false
	})
	return def
}

func (m MapOfStringToValue) InternalImplementation() types.Map {
	return m.m
}

func (m MapOfStringToValue) Equals(other types.Value) bool {
	return other != nil && __typeRefForMapOfStringToValue.Equals(other.TypeRef()) && m.Ref() == other.Ref()
}

func (m MapOfStringToValue) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToValue) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.TypeRef().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToValue.
var __typeRefForMapOfStringToValue types.TypeRef

func (m MapOfStringToValue) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToValue
}

func init() {
	__typeRefForMapOfStringToValue = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.ValueKind))
	types.RegisterFromValFunction(__typeRefForMapOfStringToValue, func(v types.Value) types.Value {
		return MapOfStringToValue{v.(types.Map), &ref.Ref{}}
	})
}

func (m MapOfStringToValue) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToValue) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToValue) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToValue) Get(p string) types.Value {
	return m.m.Get(types.NewString(p))
}

func (m MapOfStringToValue) MaybeGet(p string) (types.Value, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return types.Bool(false), false
	}
	return v, ok
}

func (m MapOfStringToValue) Set(k string, v types.Value) MapOfStringToValue {
	return MapOfStringToValue{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToValue) Remove(p string) MapOfStringToValue {
	return MapOfStringToValue{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToValueIterCallback func(k string, v types.Value) (stop bool)

func (m MapOfStringToValue) Iter(cb MapOfStringToValueIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueIterAllCallback func(k string, v types.Value)

func (m MapOfStringToValue) IterAll(cb MapOfStringToValueIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueFilterCallback func(k string, v types.Value) (keep bool)

func (m MapOfStringToValue) Filter(cb MapOfStringToValueFilterCallback) MapOfStringToValue {
	nm := NewMapOfStringToValue()
	m.IterAll(func(k string, v types.Value) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// ListOfPitch

type ListOfPitch struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfPitch() ListOfPitch {
	return ListOfPitch{types.NewList(), &ref.Ref{}}
}

type ListOfPitchDef []PitchDef

func (def ListOfPitchDef) New() ListOfPitch {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New()
	}
	return ListOfPitch{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfPitch) Def() ListOfPitchDef {
	d := make([]PitchDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(Pitch).Def()
	}
	return d
}

func (l ListOfPitch) InternalImplementation() types.List {
	return l.l
}

func (l ListOfPitch) Equals(other types.Value) bool {
	return other != nil && __typeRefForListOfPitch.Equals(other.TypeRef()) && l.Ref() == other.Ref()
}

func (l ListOfPitch) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfPitch) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.TypeRef().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfPitch.
var __typeRefForListOfPitch types.TypeRef

func (m ListOfPitch) TypeRef() types.TypeRef {
	return __typeRefForListOfPitch
}

func init() {
	__typeRefForListOfPitch = types.MakeCompoundTypeRef(types.ListKind, types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0))
	types.RegisterFromValFunction(__typeRefForListOfPitch, func(v types.Value) types.Value {
		return ListOfPitch{v.(types.List), &ref.Ref{}}
	})
}

func (l ListOfPitch) Len() uint64 {
	return l.l.Len()
}

func (l ListOfPitch) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfPitch) Get(i uint64) Pitch {
	return l.l.Get(i).(Pitch)
}

func (l ListOfPitch) Slice(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfPitch) Set(i uint64, val Pitch) ListOfPitch {
	return ListOfPitch{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfPitch) Append(v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfPitch) Insert(idx uint64, v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfPitch) Remove(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfPitch) RemoveAt(idx uint64) ListOfPitch {
	return ListOfPitch{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfPitch) fromElemSlice(p []Pitch) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfPitchIterCallback func(v Pitch, i uint64) (stop bool)

func (l ListOfPitch) Iter(cb ListOfPitchIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(Pitch), i)
	})
}

type ListOfPitchIterAllCallback func(v Pitch, i uint64)

func (l ListOfPitch) IterAll(cb ListOfPitchIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(Pitch), i)
	})
}

type ListOfPitchFilterCallback func(v Pitch, i uint64) (keep bool)

func (l ListOfPitch) Filter(cb ListOfPitchFilterCallback) ListOfPitch {
	nl := NewListOfPitch()
	l.IterAll(func(v Pitch, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}
