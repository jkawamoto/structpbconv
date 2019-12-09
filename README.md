# structpbconv
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Converts [structpb.Struct](https://github.com/golang/protobuf/blob/master/ptypes/struct/struct.pb.go)
to another structure.

## Usage
For example, let us assume to convert Payload of a [logging.Entry](https://godoc.org/cloud.google.com/go/logging#Entry) to
an ActivityPayload, which is a basic log format in
[Google Compute Engine](https://cloud.google.com/compute/).

The type of Payload of logging.Entry is `interface{}` but it can be casted
to `*structpb.Struct` in most cases.

The following code converts the instance of `*structpb.Struct` to an
instance of ActivityPayload, which is also defined in that code.

Note that to specify a field name, use `structpb` tag.

```go
import (
	"github.com/golang/protobuf/ptypes/struct"
	"github.com/jkawamoto/structpbconv"
)


type ActivityPayload struct {
	EventTimestampUs string `structpb:"event_timestamp_us"`
	EventType        string `structpb:"event_type"`
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
}


func NewActivityPayload(payload *structpb.Struct) *ActivityPayload {
	var res ActivityPayload
	structpbconv.Convert(payload, &res)
	return &res
}
```


## License
This software is released under the MIT License, see [LICENSE](LICENSE).
