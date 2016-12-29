//
// conv.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

// Package structpbconv provides a method which converts a structpb.Struct instance
// to another given structure instance.
package structpbconv

import (
	"reflect"
	"strings"

	"github.com/golang/protobuf/ptypes/struct"
)

// tagKey defines a structure tag name for ConvertStructPB.
const tagKey = "structpb"

// Convert converts a structpb.Struct object to a concrete object.
func Convert(src *structpb.Struct, dest interface{}) error {

	r := reflect.Indirect(reflect.ValueOf(dest))
	for i := 0; i < r.NumField(); i++ {

		target := r.Field(i)
		targetType := r.Type().Field(i)

		name := targetType.Tag.Get(tagKey)
		if name == "" {
			name = strings.ToLower(targetType.Name)
		}

		if v, ok := src.GetFields()[name]; ok {
			switch t := v.GetKind().(type) {
			case *structpb.Value_BoolValue:
				target.SetBool(t.BoolValue)
			case *structpb.Value_ListValue:
				target.Set(reflect.ValueOf(t.ListValue))
			case *structpb.Value_NullValue:
				target.Set(reflect.ValueOf(t.NullValue))
			case *structpb.Value_NumberValue:
				target.Set(reflect.ValueOf(t.NumberValue))
			case *structpb.Value_StringValue:
				target.SetString(t.StringValue)
			case *structpb.Value_StructValue:
				Convert(t.StructValue, target.Addr().Interface())
			}

		}

	}

	return nil

}
