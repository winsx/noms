// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/nomdl/codegen/test/gen/sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __testPackageInFile_struct_with_imports_CachedRef = __testPackageInFile_struct_with_imports_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __testPackageInFile_struct_with_imports_Ref() ref.Ref {
	p := types.PackageDef{
		Types: types.ListOfTypeRefDef{

			types.MakeEnumTypeRef("E", "E1", "Ignored"),
			types.MakeStructTypeRef("ImportUser",
				[]types.Field{
					types.Field{"importedStruct", types.MakeTypeRef(ref.Parse("sha1-d64765d6f185e4e5f7d9e67d1fc4e084b38229f5"), 0), false},
					types.Field{"enum", types.MakeTypeRef(ref.Ref{}, 0), false},
				},
				types.Choices{},
			),
		},
	}.New()
	return types.RegisterPackage(&p)
}

// E

type E uint32

const (
	E1 E = iota
	Ignored
)

// ImportUser

type ImportUser struct {
	m types.Map
}

func NewImportUser() ImportUser {
	return ImportUser{types.NewMap(
		types.NewString("$type"), types.MakeTypeRef(__testPackageInFile_struct_with_imports_CachedRef, 1),
		types.NewString("importedStruct"), sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.NewD().NomsValue(),
		types.NewString("enum"), types.UInt32(0),
	)}
}

type ImportUserDef struct {
	ImportedStruct sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.DDef
	Enum           E
}

func (def ImportUserDef) New() ImportUser {
	return ImportUser{
		types.NewMap(
			types.NewString("$type"), types.MakeTypeRef(__testPackageInFile_struct_with_imports_CachedRef, 1),
			types.NewString("importedStruct"), def.ImportedStruct.New().NomsValue(),
			types.NewString("enum"), types.UInt32(def.Enum),
		)}
}

func (s ImportUser) Def() (d ImportUserDef) {
	d.ImportedStruct = sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.DFromVal(s.m.Get(types.NewString("importedStruct"))).Def()
	d.Enum = E(s.m.Get(types.NewString("enum")).(types.UInt32))
	return
}

var __typeRefForImportUser = types.MakeTypeRef(__testPackageInFile_struct_with_imports_CachedRef, 1)

func (m ImportUser) TypeRef() types.TypeRef {
	return __typeRefForImportUser
}

func init() {
	types.RegisterFromValFunction(__typeRefForImportUser, func(v types.Value) types.NomsValue {
		return ImportUserFromVal(v)
	})
}

func ImportUserFromVal(val types.Value) ImportUser {
	// TODO: Validate here
	return ImportUser{val.(types.Map)}
}

func (s ImportUser) NomsValue() types.Value {
	return s.m
}

func (s ImportUser) Equals(other types.Value) bool {
	if other, ok := other.(ImportUser); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s ImportUser) Ref() ref.Ref {
	return s.m.Ref()
}

func (s ImportUser) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s ImportUser) ImportedStruct() sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D {
	return sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.DFromVal(s.m.Get(types.NewString("importedStruct")))
}

func (s ImportUser) SetImportedStruct(val sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D) ImportUser {
	return ImportUser{s.m.Set(types.NewString("importedStruct"), val.NomsValue())}
}

func (s ImportUser) Enum() E {
	return E(s.m.Get(types.NewString("enum")).(types.UInt32))
}

func (s ImportUser) SetEnum(val E) ImportUser {
	return ImportUser{s.m.Set(types.NewString("enum"), types.UInt32(val))}
}

// ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D

type ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D struct {
	l types.List
}

func NewListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D() ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D{types.NewList()}
}

type ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DDef []sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.DDef

func (def ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DDef) New() ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New().NomsValue()
	}
	return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D{types.NewList(l...)}
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Def() ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DDef {
	d := make([]sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.DDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.DFromVal(l.l.Get(i)).Def()
	}
	return d
}

func ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DFromVal(val types.Value) ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	// TODO: Validate here
	return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D{val.(types.List)}
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) NomsValue() types.Value {
	return l.l
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Equals(other types.Value) bool {
	if other, ok := other.(ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D); ok {
		return l.l.Equals(other.l)
	}
	return false
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Ref() ref.Ref {
	return l.l.Ref()
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Chunks() (futures []types.Future) {
	futures = append(futures, l.TypeRef().Chunks()...)
	futures = append(futures, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D.
var __typeRefForListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D types.TypeRef

func (m ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) TypeRef() types.TypeRef {
	return __typeRefForListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D
}

func init() {
	__typeRefForListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D = types.MakeCompoundTypeRef("", types.ListKind, types.MakeTypeRef(ref.Parse("sha1-d64765d6f185e4e5f7d9e67d1fc4e084b38229f5"), 0))
	types.RegisterFromValFunction(__typeRefForListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D, func(v types.Value) types.NomsValue {
		return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DFromVal(v)
	})
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Len() uint64 {
	return l.l.Len()
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Get(i uint64) sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D {
	return sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.DFromVal(l.l.Get(i))
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Slice(idx uint64, end uint64) ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D{l.l.Slice(idx, end)}
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Set(i uint64, val sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D) ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D{l.l.Set(i, val.NomsValue())}
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Append(v ...sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D) ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Insert(idx uint64, v ...sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D) ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Remove(idx uint64, end uint64) ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D{l.l.Remove(idx, end)}
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) RemoveAt(idx uint64) ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	return ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D{(l.l.RemoveAt(idx))}
}

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) fromElemSlice(p []sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

type ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DIterCallback func(v sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D, i uint64) (stop bool)

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Iter(cb ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.DFromVal(v), i)
	})
}

type ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DIterAllCallback func(v sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D, i uint64)

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) IterAll(cb ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.DFromVal(v), i)
	})
}

type ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DFilterCallback func(v sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D, i uint64) (keep bool)

func (l ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D) Filter(cb ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_DFilterCallback) ListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D {
	nl := NewListOfsha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5_D()
	l.IterAll(func(v sha1_d64765d6f185e4e5f7d9e67d1fc4e084b38229f5.D, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}
