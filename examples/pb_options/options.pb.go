// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options.proto

package pb_options

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MyEnum int32

const (
	MyEnum_FOO MyEnum = 1
	MyEnum_BAR MyEnum = 2
)

var MyEnum_name = map[int32]string{
	1: "FOO",
	2: "BAR",
}

var MyEnum_value = map[string]int32{
	"FOO": 1,
	"BAR": 2,
}

func (x MyEnum) Enum() *MyEnum {
	p := new(MyEnum)
	*p = x
	return p
}

func (x MyEnum) String() string {
	return proto.EnumName(MyEnum_name, int32(x))
}

func (x *MyEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyEnum_value, data, "MyEnum")
	if err != nil {
		return err
	}
	*x = MyEnum(value)
	return nil
}

func (MyEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

type MyMessage struct {
	Foo                  *int32   `protobuf:"varint,1,opt,name=foo" json:"foo,omitempty"`
	Bar                  *string  `protobuf:"bytes,2,opt,name=bar" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
func (*MyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

func (m *MyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMessage.Unmarshal(m, b)
}
func (m *MyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMessage.Marshal(b, m, deterministic)
}
func (m *MyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMessage.Merge(m, src)
}
func (m *MyMessage) XXX_Size() int {
	return xxx_messageInfo_MyMessage.Size(m)
}
func (m *MyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MyMessage proto.InternalMessageInfo

func (m *MyMessage) GetFoo() int32 {
	if m != nil && m.Foo != nil {
		return *m.Foo
	}
	return 0
}

func (m *MyMessage) GetBar() string {
	if m != nil && m.Bar != nil {
		return *m.Bar
	}
	return ""
}

type RequestType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestType) Reset()         { *m = RequestType{} }
func (m *RequestType) String() string { return proto.CompactTextString(m) }
func (*RequestType) ProtoMessage()    {}
func (*RequestType) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{1}
}

func (m *RequestType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestType.Unmarshal(m, b)
}
func (m *RequestType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestType.Marshal(b, m, deterministic)
}
func (m *RequestType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestType.Merge(m, src)
}
func (m *RequestType) XXX_Size() int {
	return xxx_messageInfo_RequestType.Size(m)
}
func (m *RequestType) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestType.DiscardUnknown(m)
}

var xxx_messageInfo_RequestType proto.InternalMessageInfo

type ResponseType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseType) Reset()         { *m = ResponseType{} }
func (m *ResponseType) String() string { return proto.CompactTextString(m) }
func (*ResponseType) ProtoMessage()    {}
func (*ResponseType) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{2}
}

func (m *ResponseType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseType.Unmarshal(m, b)
}
func (m *ResponseType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseType.Marshal(b, m, deterministic)
}
func (m *ResponseType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseType.Merge(m, src)
}
func (m *ResponseType) XXX_Size() int {
	return xxx_messageInfo_ResponseType.Size(m)
}
func (m *ResponseType) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseType.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseType proto.InternalMessageInfo

var E_MyFileOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "my_file_option",
	Tag:           "bytes,50000,opt,name=my_file_option",
	Filename:      "options.proto",
}

var E_MyMessageOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         50001,
	Name:          "my_message_option",
	Tag:           "varint,50001,opt,name=my_message_option",
	Filename:      "options.proto",
}

var E_MyFieldOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*float32)(nil),
	Field:         50002,
	Name:          "my_field_option",
	Tag:           "fixed32,50002,opt,name=my_field_option",
	Filename:      "options.proto",
}

var E_MyEnumOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50003,
	Name:          "my_enum_option",
	Tag:           "varint,50003,opt,name=my_enum_option",
	Filename:      "options.proto",
}

var E_MyEnumValueOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         50004,
	Name:          "my_enum_value_option",
	Tag:           "varint,50004,opt,name=my_enum_value_option",
	Filename:      "options.proto",
}

var E_MyServiceOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*MyEnum)(nil),
	Field:         50005,
	Name:          "my_service_option",
	Tag:           "varint,50005,opt,name=my_service_option,enum=MyEnum",
	Filename:      "options.proto",
}

var E_MyMethodOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*MyMessage)(nil),
	Field:         50006,
	Name:          "my_method_option",
	Tag:           "bytes,50006,opt,name=my_method_option",
	Filename:      "options.proto",
}

func init() {
	proto.RegisterEnum("MyEnum", MyEnum_name, MyEnum_value)
	proto.RegisterType((*MyMessage)(nil), "MyMessage")
	proto.RegisterType((*RequestType)(nil), "RequestType")
	proto.RegisterType((*ResponseType)(nil), "ResponseType")
	proto.RegisterExtension(E_MyFileOption)
	proto.RegisterExtension(E_MyMessageOption)
	proto.RegisterExtension(E_MyFieldOption)
	proto.RegisterExtension(E_MyEnumOption)
	proto.RegisterExtension(E_MyEnumValueOption)
	proto.RegisterExtension(E_MyServiceOption)
	proto.RegisterExtension(E_MyMethodOption)
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_110d40819f1994f9) }

var fileDescriptor_110d40819f1994f9 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xc0, 0xeb, 0xb8, 0x21, 0xc9, 0x25, 0x76, 0x5d, 0x8b, 0xe1, 0xb0, 0xf8, 0xe3, 0x66, 0xa1,
	0x62, 0x70, 0xa5, 0x8c, 0x5e, 0xa0, 0x11, 0x14, 0x96, 0x28, 0xd2, 0x35, 0xca, 0x1a, 0xa5, 0xf5,
	0x4b, 0xb0, 0x74, 0xe7, 0x33, 0x3e, 0xbb, 0xc8, 0x2b, 0x03, 0x62, 0x64, 0x41, 0x62, 0x64, 0x64,
	0xae, 0x64, 0x31, 0xf3, 0x0d, 0xa0, 0xfc, 0xf9, 0x3c, 0xe8, 0xce, 0xb9, 0x34, 0x15, 0xe9, 0x76,
	0xf7, 0xde, 0xd3, 0x4f, 0xef, 0xfd, 0xde, 0x43, 0x16, 0x4f, 0xf3, 0x98, 0x27, 0x22, 0x48, 0x33,
	0x9e, 0x73, 0xcf, 0x5f, 0x72, 0xbe, 0xa4, 0x70, 0xa4, 0x7e, 0x67, 0xc5, 0xe2, 0x28, 0x02, 0x71,
	0x9e, 0xc5, 0x69, 0xce, 0xb3, 0xba, 0xa2, 0xff, 0x14, 0x75, 0x46, 0xe5, 0x08, 0x84, 0x98, 0x2f,
	0xc1, 0xbd, 0x87, 0xcc, 0x05, 0xe7, 0xd8, 0xf0, 0x8d, 0xc3, 0xe6, 0xb0, 0xf5, 0xa9, 0xc2, 0x3b,
	0x3b, 0x1f, 0x9f, 0x11, 0x19, 0x73, 0x1d, 0x64, 0x9e, 0xcd, 0x33, 0xdc, 0xf0, 0x8d, 0xc3, 0x0e,
	0x91, 0xcf, 0xb0, 0xf9, 0xa1, 0xc2, 0x57, 0x9d, 0xbe, 0x85, 0xba, 0x04, 0xde, 0x14, 0x20, 0xf2,
	0x49, 0x99, 0x42, 0xdf, 0x46, 0x3d, 0x02, 0x22, 0xe5, 0x89, 0x00, 0xf9, 0x7f, 0xf2, 0x18, 0xdd,
	0x19, 0x95, 0x2f, 0x92, 0x82, 0xb9, 0x36, 0x32, 0x4f, 0xc6, 0x63, 0xc7, 0xf0, 0x9a, 0x5f, 0x2a,
	0xfc, 0xbd, 0xe1, 0xb6, 0x90, 0x39, 0x3c, 0x26, 0x4e, 0xc3, 0xdb, 0xfd, 0x5c, 0x61, 0x63, 0x30,
	0x91, 0x8d, 0x9c, 0x42, 0x76, 0x11, 0x9f, 0x83, 0x7b, 0x8c, 0xda, 0xb2, 0xab, 0xfc, 0x35, 0x8f,
	0xdc, 0x5e, 0xb0, 0xc1, 0xf7, 0xac, 0x60, 0x13, 0xdf, 0xc7, 0x97, 0x15, 0x36, 0xdb, 0xdf, 0x76,
	0x2f, 0x2b, 0x6c, 0xb9, 0xdd, 0x53, 0xce, 0xc0, 0x17, 0x79, 0x16, 0x27, 0x4b, 0x6f, 0xf7, 0x6b,
	0x85, 0x8d, 0xf0, 0x39, 0xb2, 0x59, 0x39, 0x5b, 0xc4, 0x14, 0x66, 0xb5, 0x19, 0xf7, 0x7e, 0x50,
	0x3b, 0x09, 0xb4, 0x93, 0xe0, 0x24, 0xa6, 0x30, 0xae, 0xb5, 0xe1, 0x1f, 0xef, 0x4d, 0x35, 0x61,
	0x8f, 0x95, 0xd7, 0xe1, 0x70, 0x84, 0xf6, 0x59, 0x39, 0x63, 0xb5, 0x25, 0x0d, 0x7a, 0xf4, 0x1f,
	0x68, 0xa5, 0x51, 0xb3, 0x7e, 0x2a, 0x56, 0x93, 0xec, 0xb1, 0xf2, 0x46, 0x26, 0x7c, 0x89, 0xf6,
	0x54, 0x53, 0x40, 0x23, 0x0d, 0x7b, 0xb0, 0xa5, 0x2b, 0xa0, 0x91, 0x46, 0x5d, 0x29, 0x54, 0x83,
	0x58, 0xb2, 0xad, 0x75, 0x7c, 0x35, 0x1d, 0x24, 0x05, 0xbb, 0x7d, 0x3a, 0xe9, 0x5e, 0x63, 0x7e,
	0x29, 0x4c, 0x5b, 0x4e, 0x77, 0x1d, 0x0e, 0x27, 0xe8, 0xae, 0xa6, 0x5c, 0xcc, 0x69, 0xb1, 0x1e,
	0xf0, 0x60, 0x2b, 0x6b, 0x2a, 0x4b, 0x34, 0xf0, 0xb7, 0x02, 0x5a, 0x64, 0xbf, 0x06, 0x6e, 0xe4,
	0xc2, 0xa9, 0x72, 0x26, 0xea, 0x85, 0xde, 0xee, 0x6c, 0xb5, 0x71, 0x0d, 0xfc, 0xa3, 0x80, 0xf6,
	0xa0, 0x15, 0xd4, 0x57, 0x23, 0xe5, 0xdd, 0x28, 0x09, 0xa7, 0xc8, 0x51, 0xbb, 0x90, 0xb7, 0xa1,
	0xb1, 0x0f, 0xb7, 0xac, 0x42, 0xe6, 0x35, 0xf5, 0xaf, 0xa2, 0x76, 0x07, 0x28, 0x58, 0xdf, 0x3a,
	0xb1, 0x59, 0xb9, 0x59, 0x34, 0x74, 0xde, 0x55, 0xb8, 0xf7, 0x0a, 0x28, 0xe5, 0xfe, 0x5b, 0x9e,
	0xd1, 0xe8, 0xe0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x87, 0x17, 0x3d, 0x4c, 0x03, 0x00,
	0x00,
}
