// BSON library for Go
//
// Copyright (c) 2010-2012 - Gustavo Niemeyer <gustavo@niemeyer.net>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package bsonutils is an implementation of the BSON specification for Go:
//
//     http://bsonspec.org
//
// It was created as part of the mgo MongoDB driver for Go, but is standalone
// and may be used on its own without the driver.
package bsonutils

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/globalsign/mgo/bson"
)

// --------------------------------------------------------------------------
// The public API.

type undefined struct{}

const initialBufferSize = 64

func handleErr(err *error) {
	if r := recover(); r != nil {
		if _, ok := r.(runtime.Error); ok {
			panic(r)
		} else if _, ok := r.(externalPanic); ok {
			panic(r)
		} else if s, ok := r.(string); ok {
			*err = errors.New(s)
		} else if e, ok := r.(error); ok {
			*err = e
		} else {
			panic(r)
		}
	}
}

// Marshal serializes the in value, which may be a map or a struct value.
// In the case of struct values, only exported fields will be serialized,
// and the order of serialized fields will match that of the struct itself.
// The lowercased field name is used as the key for each exported field,
// but this behavior may be changed using the respective field tag.
// The tag may also contain flags to tweak the marshalling behavior for
// the field. The tag formats accepted are:
//
//     "[<key>][,<flag1>[,<flag2>]]"
//
//     `(...) bsonutils:"[<key>][,<flag1>[,<flag2>]]" (...)`
//
// The following flags are currently supported:
//
//     omitempty  Only include the field if it's not set to the zero
//                value for the type or to empty slices or maps.
//
//     minsize    Marshal an int64 value as an int32, if that's feasible
//                while preserving the numeric value.
//
//     inline     Inline the field, which must be a struct or a map,
//                causing all of its fields or keys to be processed as if
//                they were part of the outer struct. For maps, keys must
//                not conflict with the bsonutils keys of other struct fields.
//
// Some examples:
//
//     type T struct {
//         A bool
//         B int    "myb"
//         C string "myc,omitempty"
//         D string `bsonutils:",omitempty" json:"jsonkey"`
//         E int64  ",minsize"
//         F int64  "myf,omitempty,minsize"
//     }
//
func Marshal(in interface{}) (out []byte, err error) {
	return MarshalBuffer(in, make([]byte, 0, initialBufferSize))
}

// MarshalBuffer behaves the same way as Marshal, except that instead of
// allocating a new byte slice it tries to use the received byte slice and
// only allocates more memory if necessary to fit the marshaled value.
func MarshalBuffer(in interface{}, buf []byte) (out []byte, err error) {
	defer handleErr(&err)
	e := &encoder{buf}
	e.addDoc(reflect.ValueOf(in))
	return e.out, nil
}

// Unmarshal deserializes data from in into the out value.  The out value
// must be a map, a pointer to a struct, or a pointer to a bsonutils.D value.
// In the case of struct values, only exported fields will be deserialized.
// The lowercased field name is used as the key for each exported field,
// but this behavior may be changed using the respective field tag.
// The tag may also contain flags to tweak the marshalling behavior for
// the field. The tag formats accepted are:
//
//     "[<key>][,<flag1>[,<flag2>]]"
//
//     `(...) bsonutils:"[<key>][,<flag1>[,<flag2>]]" (...)`
//
// The following flags are currently supported during unmarshal (see the
// Marshal method for other flags):
//
//     inline     Inline the field, which must be a struct or a map.
//                Inlined structs are handled as if its fields were part
//                of the outer struct. An inlined map causes keys that do
//                not match any other struct field to be inserted in the
//                map rather than being discarded as usual.
//
// The target field or element types of out may not necessarily match
// the BSON values of the provided data.  The following conversions are
// made automatically:
//
// - Numeric types are converted if at least the integer part of the
//   value would be preserved correctly
// - Bools are converted to numeric types as 1 or 0
// - Numeric types are converted to bools as true if not 0 or false otherwise
// - Binary and string BSON data is converted to a string, array or byte slice
//
// If the value would not fit the type and cannot be converted, it's
// silently skipped.
//
// Pointer values are initialized when necessary.
func Unmarshal(in []byte, out interface{}) (err error) {
	if raw, ok := out.(*bson.Raw); ok {
		raw.Kind = 3
		raw.Data = in
		return nil
	}
	defer handleErr(&err)
	v := reflect.ValueOf(out)
	switch v.Kind() {
	case reflect.Ptr:
		fallthrough
	case reflect.Map:
		d := newDecoder(in)
		d.readDocTo(v)
		if d.i < len(d.in) {
			return errors.New("document is corrupted")
		}
	case reflect.Struct:
		return errors.New("unmarshal can't deal with struct values. Use a pointer")
	default:
		return errors.New("unmarshal needs a map or a pointer to a struct")
	}
	return nil
}

// --------------------------------------------------------------------------
// Maintain a mapping of keys to structure field indexes

type structInfo struct {
	FieldsMap  map[string]fieldInfo
	FieldsList []fieldInfo
	InlineMap  int
	Zero       reflect.Value
}

type fieldInfo struct {
	Key       string
	Num       int
	OmitEmpty bool
	MinSize   bool
	Inline    []int
}

var structMap = make(map[reflect.Type]*structInfo)
var structMapMutex sync.RWMutex

type externalPanic string

func (e externalPanic) String() string {
	return string(e)
}

func getStructInfo(st reflect.Type) (*structInfo, error) {
	structMapMutex.RLock()
	sinfo, found := structMap[st]
	structMapMutex.RUnlock()
	if found {
		return sinfo, nil
	}
	n := st.NumField()
	fieldsMap := make(map[string]fieldInfo)
	fieldsList := make([]fieldInfo, 0, n)
	inlineMap := -1
	for i := 0; i != n; i++ {
		field := st.Field(i)
		if field.PkgPath != "" && !field.Anonymous {
			continue // Private field
		}

		info := fieldInfo{Num: i}

		tag := field.Tag.Get("bson")

		// If there's no bsonutils/json tag available.
		if tag == "" {
			// If there's no tag, and also no tag: value splits (i.e. no colon)
			// then assume the entire tag is the value
			if !strings.Contains(string(field.Tag), ":") {
				tag = string(field.Tag)
			}
		}

		if tag == "-" {
			continue
		}

		inline := false
		fields := strings.Split(tag, ",")
		if len(fields) > 1 {
			for _, flag := range fields[1:] {
				switch flag {
				case "omitempty":
					info.OmitEmpty = true
				case "minsize":
					info.MinSize = true
				case "inline":
					inline = true
				default:
					msg := fmt.Sprintf("Unsupported flag %q in tag %q of type %s", flag, tag, st)
					panic(externalPanic(msg))
				}
			}
			tag = fields[0]
		}

		if inline {
			switch field.Type.Kind() {
			case reflect.Map:
				if inlineMap >= 0 {
					return nil, errors.New("Multiple ,inline maps in struct " + st.String())
				}
				if field.Type.Key() != reflect.TypeOf("") {
					return nil, errors.New("Option ,inline needs a map with string keys in struct " + st.String())
				}
				inlineMap = info.Num
			case reflect.Ptr:
				// allow only pointer to struct
				if kind := field.Type.Elem().Kind(); kind != reflect.Struct {
					return nil, errors.New("Option ,inline allows a pointer only to a struct, was given pointer to " + kind.String())
				}

				field.Type = field.Type.Elem()
				fallthrough
			case reflect.Struct:
				sinfo, err := getStructInfo(field.Type)
				if err != nil {
					return nil, err
				}
				for _, finfo := range sinfo.FieldsList {
					if _, found := fieldsMap[finfo.Key]; found {
						msg := "Duplicated key '" + finfo.Key + "' in struct " + st.String()
						return nil, errors.New(msg)
					}
					if finfo.Inline == nil {
						finfo.Inline = []int{i, finfo.Num}
					} else {
						finfo.Inline = append([]int{i}, finfo.Inline...)
					}
					fieldsMap[finfo.Key] = finfo
					fieldsList = append(fieldsList, finfo)
				}
			default:
				panic("Option ,inline needs a struct value or a pointer to a struct or map field")
			}
			continue
		}

		if tag != "" {
			info.Key = tag
		} else {
			info.Key = strings.ToLower(field.Name)
		}

		if _, found = fieldsMap[info.Key]; found {
			msg := "Duplicated key '" + info.Key + "' in struct " + st.String()
			return nil, errors.New(msg)
		}

		fieldsList = append(fieldsList, info)
		fieldsMap[info.Key] = info
	}
	sinfo = &structInfo{
		fieldsMap,
		fieldsList,
		inlineMap,
		reflect.New(st).Elem(),
	}
	structMapMutex.Lock()
	structMap[st] = sinfo
	structMapMutex.Unlock()
	return sinfo, nil
}

// useOmitEmptyAsDefault defines if OmitEmpty should be used on all fields.
var useOmitEmptyAsDefault = false

// SetOmitEmptyAsDefault enable or disables tag omitempty for all fields.
func SetOmitEmptyAsDefault(state bool) {
	useOmitEmptyAsDefault = state
}

// IsOmitEmptyAsDefault check if omitempty is enable for all fields.
func IsOmitEmptyAsDefault() bool {
	return useOmitEmptyAsDefault
}
