package noop

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.headers.noop.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/headers/noop/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x2e, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0x2d, 0xd1, 0xcf, 0x48,
	0x4d, 0x4c, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xcb, 0xcf, 0x2f, 0xd0, 0x4f, 0xce, 0xcf, 0x4b, 0xcb,
	0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x82, 0x69, 0x2e, 0x4a, 0xd5, 0x83, 0x6b,
	0xd4, 0x83, 0x69, 0xd4, 0x83, 0x6a, 0xd4, 0x03, 0x69, 0x54, 0xe2, 0xe0, 0x62, 0x73, 0x06, 0xeb,
	0x75, 0x2a, 0xe0, 0x02, 0x59, 0xa7, 0x47, 0xbc, 0x5e, 0x27, 0x6e, 0x88, 0xce, 0x00, 0x90, 0xa5,
	0x51, 0x2c, 0x20, 0xa1, 0x55, 0x4c, 0x5a, 0x61, 0x46, 0x41, 0x89, 0x95, 0x7a, 0xce, 0x20, 0xfd,
	0x21, 0x70, 0xfd, 0x9e, 0x30, 0xfd, 0x1e, 0x50, 0xfd, 0x7e, 0xf9, 0xf9, 0x05, 0x49, 0x6c, 0x60,
	0xe7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x89, 0x52, 0x33, 0x81, 0xed, 0x00, 0x00, 0x00,
}