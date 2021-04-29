package codec

import (
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/encoding"
)

const Name = "mycodec"

func init() {
	encoding.RegisterCodec(demoCodec{})
}

type demoCodec struct {
}

func (demoCodec) Marshal(v interface{}) ([]byte, error) {
	fmt.Println("demo encode:", v)
	return json.Marshal(v)
}

func (demoCodec) Unmarshal(data []byte, v interface{}) error {
	fmt.Println("demo decode:", data, string(data))
	return json.Unmarshal(data, v)
}

func (demoCodec) Name() string {
	return Name
}
