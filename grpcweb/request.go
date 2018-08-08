package grpcweb

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Request struct {
	endpoint string
	in, out  interface{}
}

func NewRequest(
	endpoint string,
	in proto.Message,
	out proto.Message,
) *Request {
	// desc, err := desc.LoadMessageDescriptorForMessage(out)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "invalid MessageDescriptor passed")
	// }
	return &Request{
		endpoint: endpoint,
		in:       in,
		out:      out,
	}
}

func ToEndpoint(pkg string, s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) string {
	return fmt.Sprintf("/%s.%s/%s", pkg, s.GetName(), m.GetName())
}
