//
// doc.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

// Package structpbconv provides a method which converts a structpb.Struct instance
// to another given structure instance.
//
// For example, let us assume to convert Payload of a logging.Entry to
// an ActivityPayload, which is a basic log format in
// Google Compute Engine.
//
// The type of Payload of logging.Entry is `interface{}` but it can be casted
// to `*structpb.Struct` in most cases.
//
// The following example converts the instance of `*structpb.Struct` to an
// instance of ActivityPayload, which is also defined in that code.
//
// Note that to specify a field name, use `structpb` tag.
package structpbconv
