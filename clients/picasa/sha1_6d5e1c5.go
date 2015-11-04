// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_sha1_6d5e1c5_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("Geoposition",
			[]types.Field{
				types.Field{"Latitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
				types.Field{"Longitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
			},
			types.Choices{},
		),
		types.MakeStructTypeRef("Georectangle",
			[]types.Field{
				types.Field{"TopLeft", types.MakeTypeRef(ref.Ref{}, 0), false},
				types.Field{"BottomRight", types.MakeTypeRef(ref.Ref{}, 0), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__mainPackageInFile_sha1_6d5e1c5_CachedRef = types.RegisterPackage(&p)
}

// Geoposition

type Geoposition struct {
	_Latitude  float32
	_Longitude float32

	ref *ref.Ref
}

func NewGeoposition() Geoposition {
	return Geoposition{
		_Latitude:  float32(0),
		_Longitude: float32(0),

		ref: &ref.Ref{},
	}
}

type GeopositionDef struct {
	Latitude  float32
	Longitude float32
}

func (def GeopositionDef) New() Geoposition {
	return Geoposition{
		_Latitude:  def.Latitude,
		_Longitude: def.Longitude,
		ref:        &ref.Ref{},
	}
}

func (s Geoposition) Def() (d GeopositionDef) {
	d.Latitude = s._Latitude
	d.Longitude = s._Longitude
	return
}

var __typeRefForGeoposition types.TypeRef
var __typeDefForGeoposition types.TypeRef

func (m Geoposition) TypeRef() types.TypeRef {
	return __typeRefForGeoposition
}

func init() {
	__typeRefForGeoposition = types.MakeTypeRef(__mainPackageInFile_sha1_6d5e1c5_CachedRef, 0)
	__typeDefForGeoposition = types.MakeStructTypeRef("Geoposition",
		[]types.Field{
			types.Field{"Latitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
			types.Field{"Longitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
		},
		types.Choices{},
	)
	types.RegisterStructBuilder(__typeRefForGeoposition, builderForGeoposition)
}

func (s Geoposition) InternalImplementation() types.Struct {
	// TODO: Remove this
	m := map[string]types.Value{
		"Latitude":  types.Float32(s._Latitude),
		"Longitude": types.Float32(s._Longitude),
	}
	return types.NewStruct(__typeRefForGeoposition, __typeDefForGeoposition, m)
}

func builderForGeoposition() chan types.Value {
	c := make(chan types.Value)
	s := Geoposition{ref: &ref.Ref{}}
	go func() {
		s._Latitude = float32((<-c).(types.Float32))
		s._Longitude = float32((<-c).(types.Float32))

		c <- s
	}()
	return c
}

func (s Geoposition) Equals(other types.Value) bool {
	return other != nil && __typeRefForGeoposition.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s Geoposition) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Geoposition) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForGeoposition.Chunks()...)
	return
}

func (s Geoposition) Latitude() float32 {
	return s._Latitude
}

func (s Geoposition) SetLatitude(val float32) Geoposition {
	s._Latitude = val
	s.ref = &ref.Ref{}
	return s
}

func (s Geoposition) Longitude() float32 {
	return s._Longitude
}

func (s Geoposition) SetLongitude(val float32) Geoposition {
	s._Longitude = val
	s.ref = &ref.Ref{}
	return s
}

// Georectangle

type Georectangle struct {
	_TopLeft     Geoposition
	_BottomRight Geoposition

	ref *ref.Ref
}

func NewGeorectangle() Georectangle {
	return Georectangle{
		_TopLeft:     NewGeoposition(),
		_BottomRight: NewGeoposition(),

		ref: &ref.Ref{},
	}
}

type GeorectangleDef struct {
	TopLeft     GeopositionDef
	BottomRight GeopositionDef
}

func (def GeorectangleDef) New() Georectangle {
	return Georectangle{
		_TopLeft:     def.TopLeft.New(),
		_BottomRight: def.BottomRight.New(),
		ref:          &ref.Ref{},
	}
}

func (s Georectangle) Def() (d GeorectangleDef) {
	d.TopLeft = s._TopLeft.Def()
	d.BottomRight = s._BottomRight.Def()
	return
}

var __typeRefForGeorectangle types.TypeRef
var __typeDefForGeorectangle types.TypeRef

func (m Georectangle) TypeRef() types.TypeRef {
	return __typeRefForGeorectangle
}

func init() {
	__typeRefForGeorectangle = types.MakeTypeRef(__mainPackageInFile_sha1_6d5e1c5_CachedRef, 1)
	__typeDefForGeorectangle = types.MakeStructTypeRef("Georectangle",
		[]types.Field{
			types.Field{"TopLeft", types.MakeTypeRef(__mainPackageInFile_sha1_6d5e1c5_CachedRef, 0), false},
			types.Field{"BottomRight", types.MakeTypeRef(__mainPackageInFile_sha1_6d5e1c5_CachedRef, 0), false},
		},
		types.Choices{},
	)
	types.RegisterStructBuilder(__typeRefForGeorectangle, builderForGeorectangle)
}

func (s Georectangle) InternalImplementation() types.Struct {
	// TODO: Remove this
	m := map[string]types.Value{
		"TopLeft":     s._TopLeft,
		"BottomRight": s._BottomRight,
	}
	return types.NewStruct(__typeRefForGeorectangle, __typeDefForGeorectangle, m)
}

func builderForGeorectangle() chan types.Value {
	c := make(chan types.Value)
	s := Georectangle{ref: &ref.Ref{}}
	go func() {
		s._TopLeft = (<-c).(Geoposition)
		s._BottomRight = (<-c).(Geoposition)

		c <- s
	}()
	return c
}

func (s Georectangle) Equals(other types.Value) bool {
	return other != nil && __typeRefForGeorectangle.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s Georectangle) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Georectangle) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForGeorectangle.Chunks()...)
	chunks = append(chunks, s._TopLeft.Chunks()...)
	chunks = append(chunks, s._BottomRight.Chunks()...)
	return
}

func (s Georectangle) TopLeft() Geoposition {
	return s._TopLeft
}

func (s Georectangle) SetTopLeft(val Geoposition) Georectangle {
	s._TopLeft = val
	s.ref = &ref.Ref{}
	return s
}

func (s Georectangle) BottomRight() Geoposition {
	return s._BottomRight
}

func (s Georectangle) SetBottomRight(val Geoposition) Georectangle {
	s._BottomRight = val
	s.ref = &ref.Ref{}
	return s
}