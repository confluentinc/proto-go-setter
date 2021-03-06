// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/person5.proto

package person

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/confluentinc/proto-go-setter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Person5 struct {
	Id           int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumbers map[string]string `protobuf:"bytes,3,rep,name=phone_numbers,json=phoneNumbers" json:"phone_numbers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Person5) Reset()                    { *m = Person5{} }
func (m *Person5) String() string            { return proto.CompactTextString(m) }
func (*Person5) ProtoMessage()               {}
func (*Person5) Descriptor() ([]byte, []int) { return fileDescriptorPerson5, []int{0} }

func (m *Person5) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person5) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person5) GetPhoneNumbers() map[string]string {
	if m != nil {
		return m.PhoneNumbers
	}
	return nil
}

func init() {
	proto.RegisterType((*Person5)(nil), "Person5")
}

func init() { proto.RegisterFile("example/person5.proto", fileDescriptorPerson5) }

var fileDescriptorPerson5 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2f, 0x48, 0x2d, 0x2a, 0xce, 0xcf, 0x33, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x97, 0x32, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce,
	0x4f, 0xa9, 0x4c, 0x2c, 0x4a, 0xac, 0xd4, 0x07, 0x4b, 0xe9, 0xa6, 0xe7, 0xeb, 0x16, 0xa7, 0x96,
	0x94, 0xa4, 0x16, 0xe9, 0x43, 0x28, 0x88, 0x0e, 0xa5, 0xad, 0x8c, 0x5c, 0xec, 0x01, 0x10, 0x33,
	0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98, 0x32, 0x53,
	0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0,
	0x6c, 0x21, 0x57, 0x2e, 0xde, 0x82, 0x8c, 0xfc, 0xbc, 0xd4, 0xf8, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4,
	0xa2, 0x62, 0x09, 0x66, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x29, 0x3d, 0xa8, 0x21, 0x7a, 0x01, 0x20,
	0x59, 0x3f, 0x88, 0xa4, 0x6b, 0x5e, 0x49, 0x51, 0xa5, 0x13, 0xcb, 0x87, 0x5f, 0xda, 0x8c, 0x41,
	0x3c, 0x05, 0x48, 0x12, 0x52, 0xf6, 0x5c, 0x82, 0x18, 0x0a, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53,
	0x2b, 0xc1, 0x0e, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x61,
	0x4e, 0x80, 0x70, 0xac, 0x98, 0x2c, 0x18, 0x9d, 0x38, 0xa2, 0xd8, 0x20, 0x5e, 0x4f, 0x62, 0x03,
	0x7b, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x86, 0x9a, 0xd2, 0xa2, 0x13, 0x01, 0x00, 0x00,
}
