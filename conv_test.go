//
// conv_test.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package structpbconv_test

import (
	"reflect"
	"testing"

	structpb "github.com/golang/protobuf/ptypes/struct"

	"github.com/jkawamoto/structpbconv"
)

type ActivityPayload struct {
	EventTimestampUs string `structpb:"event_timestamp_us"`
	EventType        string `structpb:"vent_type"`
	TraceID          string `structpb:"trace_id"`
	Actor            struct {
		User string
	}
	Resource struct {
		Zone string
		Type string
		ID   string
		Name string
	}
	Version      string
	EventSubtype string `structpb:"event_subtype"`
	Operation    struct {
		Zone string
		Type string
		ID   string
		Name string
	}
	Items []string
	Map   map[string]string
}

func TestConvert(t *testing.T) {
	src := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"items": {
				Kind: &structpb.Value_ListValue{
					ListValue: &structpb.ListValue{
						Values: []*structpb.Value{
							{
								Kind: &structpb.Value_StringValue{
									StringValue: "Hello World",
								},
							},
						},
					},
				},
			},
			"map": {
				Kind: &structpb.Value_StructValue{
					StructValue: &structpb.Struct{
						Fields: map[string]*structpb.Value{
							"key1": {
								Kind: &structpb.Value_StringValue{
									StringValue: "value1",
								},
							},
						},
					},
				},
			},
		},
	}
	var payload ActivityPayload
	if err := structpbconv.Convert(src, &payload); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(payload, ActivityPayload{
		Items: []string{"Hello World"},
		Map: map[string]string{
			"key1": "value1",
		},
	}) {
		t.Fatal("failed to convert")
	}
}

func ExampleConvert() {

	var payload interface{}
	var res *ActivityPayload

	switch s := payload.(type) {
	case *structpb.Struct:
		res = &ActivityPayload{}
		structpbconv.Convert(s, res)
	default:
		// Error: Given payload is not an instance of *structpb.Struct.
	}

}
