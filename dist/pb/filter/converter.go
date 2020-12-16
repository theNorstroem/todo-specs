package filterpb

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
)

func (x *Filter) NewFromBase64(input string) {
	f, err := base64.StdEncoding.DecodeString(input)
	*x = Filter{}
	if err == nil {
		proto.Unmarshal(f, x)
	}
}
