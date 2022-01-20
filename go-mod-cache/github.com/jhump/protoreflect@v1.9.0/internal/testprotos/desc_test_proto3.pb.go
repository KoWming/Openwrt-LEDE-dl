// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.0
// source: desc_test_proto3.proto

package testprotos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	pkg "github.com/jhump/protoreflect/internal/testprotos/pkg"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Proto3Enum int32

const (
	Proto3Enum_UNKNOWN Proto3Enum = 0
	Proto3Enum_VALUE1  Proto3Enum = 1
	Proto3Enum_VALUE2  Proto3Enum = 2
)

// Enum value maps for Proto3Enum.
var (
	Proto3Enum_name = map[int32]string{
		0: "UNKNOWN",
		1: "VALUE1",
		2: "VALUE2",
	}
	Proto3Enum_value = map[string]int32{
		"UNKNOWN": 0,
		"VALUE1":  1,
		"VALUE2":  2,
	}
)

func (x Proto3Enum) Enum() *Proto3Enum {
	p := new(Proto3Enum)
	*p = x
	return p
}

func (x Proto3Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Proto3Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_desc_test_proto3_proto_enumTypes[0].Descriptor()
}

func (Proto3Enum) Type() protoreflect.EnumType {
	return &file_desc_test_proto3_proto_enumTypes[0]
}

func (x Proto3Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Proto3Enum.Descriptor instead.
func (Proto3Enum) EnumDescriptor() ([]byte, []int) {
	return file_desc_test_proto3_proto_rawDescGZIP(), []int{0}
}

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo    []Proto3Enum                                    `protobuf:"varint,1,rep,packed,name=foo,proto3,enum=testprotos.Proto3Enum" json:"foo,omitempty"`
	Bar    string                                          `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
	Baz    *TestMessage                                    `protobuf:"bytes,3,opt,name=baz,proto3" json:"baz,omitempty"`
	Snafu  *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,4,opt,name=snafu,proto3" json:"snafu,omitempty"`
	Flags  map[string]bool                                 `protobuf:"bytes,5,rep,name=flags,proto3" json:"flags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Others map[string]*TestMessage                         `protobuf:"bytes,6,rep,name=others,proto3" json:"others,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_test_proto3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desc_test_proto3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_desc_test_proto3_proto_rawDescGZIP(), []int{0}
}

func (x *TestRequest) GetFoo() []Proto3Enum {
	if x != nil {
		return x.Foo
	}
	return nil
}

func (x *TestRequest) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

func (x *TestRequest) GetBaz() *TestMessage {
	if x != nil {
		return x.Baz
	}
	return nil
}

func (x *TestRequest) GetSnafu() *TestMessage_NestedMessage_AnotherNestedMessage {
	if x != nil {
		return x.Snafu
	}
	return nil
}

func (x *TestRequest) GetFlags() map[string]bool {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *TestRequest) GetOthers() map[string]*TestMessage {
	if x != nil {
		return x.Others
	}
	return nil
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Atm *AnotherTestMessage `protobuf:"bytes,1,opt,name=atm,proto3" json:"atm,omitempty"`
	Vs  []int32             `protobuf:"varint,2,rep,packed,name=vs,proto3" json:"vs,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_test_proto3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_desc_test_proto3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_desc_test_proto3_proto_rawDescGZIP(), []int{1}
}

func (x *TestResponse) GetAtm() *AnotherTestMessage {
	if x != nil {
		return x.Atm
	}
	return nil
}

func (x *TestResponse) GetVs() []int32 {
	if x != nil {
		return x.Vs
	}
	return nil
}

var File_desc_test_proto3_proto protoreflect.FileDescriptor

var file_desc_test_proto3_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x10, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcb, 0x03, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x03, 0x62,
	0x61, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x03, 0x62, 0x61, 0x7a, 0x12, 0x50, 0x0a, 0x05, 0x73, 0x6e, 0x61, 0x66, 0x75, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x6e, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x05, 0x73, 0x6e, 0x61, 0x66, 0x75, 0x12, 0x38, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x1a,
	0x38, 0x0a, 0x0a, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x52, 0x0a, 0x0b, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x50, 0x0a,
	0x0c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x03, 0x61, 0x74, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x61, 0x74, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x76, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x02, 0x76, 0x73, 0x2a,
	0x31, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x31, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x32,
	0x10, 0x02, 0x32, 0xbc, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x44, 0x6f, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6a, 0x68, 0x75,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2e,
	0x64, 0x65, 0x73, 0x63, 0x2e, 0x42, 0x61, 0x72, 0x12, 0x46, 0x0a, 0x0f, 0x44, 0x6f, 0x53, 0x6f,
	0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x6c, 0x73, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x12, 0x52, 0x0a, 0x10, 0x44, 0x6f, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x41,
	0x67, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x6a, 0x68, 0x75, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x2e, 0x42,
	0x61, 0x72, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x12, 0x44, 0x6f, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x68, 0x75, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_desc_test_proto3_proto_rawDescOnce sync.Once
	file_desc_test_proto3_proto_rawDescData = file_desc_test_proto3_proto_rawDesc
)

func file_desc_test_proto3_proto_rawDescGZIP() []byte {
	file_desc_test_proto3_proto_rawDescOnce.Do(func() {
		file_desc_test_proto3_proto_rawDescData = protoimpl.X.CompressGZIP(file_desc_test_proto3_proto_rawDescData)
	})
	return file_desc_test_proto3_proto_rawDescData
}

var file_desc_test_proto3_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_desc_test_proto3_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_desc_test_proto3_proto_goTypes = []interface{}{
	(Proto3Enum)(0),      // 0: testprotos.Proto3Enum
	(*TestRequest)(nil),  // 1: testprotos.TestRequest
	(*TestResponse)(nil), // 2: testprotos.TestResponse
	nil,                  // 3: testprotos.TestRequest.FlagsEntry
	nil,                  // 4: testprotos.TestRequest.OthersEntry
	(*TestMessage)(nil),  // 5: testprotos.TestMessage
	(*TestMessage_NestedMessage_AnotherNestedMessage)(nil), // 6: testprotos.TestMessage.NestedMessage.AnotherNestedMessage
	(*AnotherTestMessage)(nil),                             // 7: testprotos.AnotherTestMessage
	(*pkg.Bar)(nil),                                        // 8: jhump.protoreflect.desc.Bar
}
var file_desc_test_proto3_proto_depIdxs = []int32{
	0,  // 0: testprotos.TestRequest.foo:type_name -> testprotos.Proto3Enum
	5,  // 1: testprotos.TestRequest.baz:type_name -> testprotos.TestMessage
	6,  // 2: testprotos.TestRequest.snafu:type_name -> testprotos.TestMessage.NestedMessage.AnotherNestedMessage
	3,  // 3: testprotos.TestRequest.flags:type_name -> testprotos.TestRequest.FlagsEntry
	4,  // 4: testprotos.TestRequest.others:type_name -> testprotos.TestRequest.OthersEntry
	7,  // 5: testprotos.TestResponse.atm:type_name -> testprotos.AnotherTestMessage
	5,  // 6: testprotos.TestRequest.OthersEntry.value:type_name -> testprotos.TestMessage
	1,  // 7: testprotos.TestService.DoSomething:input_type -> testprotos.TestRequest
	5,  // 8: testprotos.TestService.DoSomethingElse:input_type -> testprotos.TestMessage
	8,  // 9: testprotos.TestService.DoSomethingAgain:input_type -> jhump.protoreflect.desc.Bar
	1,  // 10: testprotos.TestService.DoSomethingForever:input_type -> testprotos.TestRequest
	8,  // 11: testprotos.TestService.DoSomething:output_type -> jhump.protoreflect.desc.Bar
	2,  // 12: testprotos.TestService.DoSomethingElse:output_type -> testprotos.TestResponse
	7,  // 13: testprotos.TestService.DoSomethingAgain:output_type -> testprotos.AnotherTestMessage
	2,  // 14: testprotos.TestService.DoSomethingForever:output_type -> testprotos.TestResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_desc_test_proto3_proto_init() }
func file_desc_test_proto3_proto_init() {
	if File_desc_test_proto3_proto != nil {
		return
	}
	file_desc_test1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_desc_test_proto3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_desc_test_proto3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_desc_test_proto3_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_desc_test_proto3_proto_goTypes,
		DependencyIndexes: file_desc_test_proto3_proto_depIdxs,
		EnumInfos:         file_desc_test_proto3_proto_enumTypes,
		MessageInfos:      file_desc_test_proto3_proto_msgTypes,
	}.Build()
	File_desc_test_proto3_proto = out.File
	file_desc_test_proto3_proto_rawDesc = nil
	file_desc_test_proto3_proto_goTypes = nil
	file_desc_test_proto3_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	DoSomething(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*pkg.Bar, error)
	DoSomethingElse(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingElseClient, error)
	DoSomethingAgain(ctx context.Context, in *pkg.Bar, opts ...grpc.CallOption) (TestService_DoSomethingAgainClient, error)
	DoSomethingForever(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingForeverClient, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) DoSomething(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*pkg.Bar, error) {
	out := new(pkg.Bar)
	err := c.cc.Invoke(ctx, "/testprotos.TestService/DoSomething", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) DoSomethingElse(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingElseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/testprotos.TestService/DoSomethingElse", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceDoSomethingElseClient{stream}
	return x, nil
}

type TestService_DoSomethingElseClient interface {
	Send(*TestMessage) error
	CloseAndRecv() (*TestResponse, error)
	grpc.ClientStream
}

type testServiceDoSomethingElseClient struct {
	grpc.ClientStream
}

func (x *testServiceDoSomethingElseClient) Send(m *TestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceDoSomethingElseClient) CloseAndRecv() (*TestResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) DoSomethingAgain(ctx context.Context, in *pkg.Bar, opts ...grpc.CallOption) (TestService_DoSomethingAgainClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/testprotos.TestService/DoSomethingAgain", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceDoSomethingAgainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_DoSomethingAgainClient interface {
	Recv() (*AnotherTestMessage, error)
	grpc.ClientStream
}

type testServiceDoSomethingAgainClient struct {
	grpc.ClientStream
}

func (x *testServiceDoSomethingAgainClient) Recv() (*AnotherTestMessage, error) {
	m := new(AnotherTestMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) DoSomethingForever(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingForeverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[2], "/testprotos.TestService/DoSomethingForever", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceDoSomethingForeverClient{stream}
	return x, nil
}

type TestService_DoSomethingForeverClient interface {
	Send(*TestRequest) error
	Recv() (*TestResponse, error)
	grpc.ClientStream
}

type testServiceDoSomethingForeverClient struct {
	grpc.ClientStream
}

func (x *testServiceDoSomethingForeverClient) Send(m *TestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceDoSomethingForeverClient) Recv() (*TestResponse, error) {
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	DoSomething(context.Context, *TestRequest) (*pkg.Bar, error)
	DoSomethingElse(TestService_DoSomethingElseServer) error
	DoSomethingAgain(*pkg.Bar, TestService_DoSomethingAgainServer) error
	DoSomethingForever(TestService_DoSomethingForeverServer) error
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) DoSomething(context.Context, *TestRequest) (*pkg.Bar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSomething not implemented")
}
func (*UnimplementedTestServiceServer) DoSomethingElse(TestService_DoSomethingElseServer) error {
	return status.Errorf(codes.Unimplemented, "method DoSomethingElse not implemented")
}
func (*UnimplementedTestServiceServer) DoSomethingAgain(*pkg.Bar, TestService_DoSomethingAgainServer) error {
	return status.Errorf(codes.Unimplemented, "method DoSomethingAgain not implemented")
}
func (*UnimplementedTestServiceServer) DoSomethingForever(TestService_DoSomethingForeverServer) error {
	return status.Errorf(codes.Unimplemented, "method DoSomethingForever not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_DoSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).DoSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testprotos.TestService/DoSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).DoSomething(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_DoSomethingElse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).DoSomethingElse(&testServiceDoSomethingElseServer{stream})
}

type TestService_DoSomethingElseServer interface {
	SendAndClose(*TestResponse) error
	Recv() (*TestMessage, error)
	grpc.ServerStream
}

type testServiceDoSomethingElseServer struct {
	grpc.ServerStream
}

func (x *testServiceDoSomethingElseServer) SendAndClose(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceDoSomethingElseServer) Recv() (*TestMessage, error) {
	m := new(TestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_DoSomethingAgain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pkg.Bar)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).DoSomethingAgain(m, &testServiceDoSomethingAgainServer{stream})
}

type TestService_DoSomethingAgainServer interface {
	Send(*AnotherTestMessage) error
	grpc.ServerStream
}

type testServiceDoSomethingAgainServer struct {
	grpc.ServerStream
}

func (x *testServiceDoSomethingAgainServer) Send(m *AnotherTestMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_DoSomethingForever_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).DoSomethingForever(&testServiceDoSomethingForeverServer{stream})
}

type TestService_DoSomethingForeverServer interface {
	Send(*TestResponse) error
	Recv() (*TestRequest, error)
	grpc.ServerStream
}

type testServiceDoSomethingForeverServer struct {
	grpc.ServerStream
}

func (x *testServiceDoSomethingForeverServer) Send(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceDoSomethingForeverServer) Recv() (*TestRequest, error) {
	m := new(TestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testprotos.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
			Handler:    _TestService_DoSomething_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoSomethingElse",
			Handler:       _TestService_DoSomethingElse_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DoSomethingAgain",
			Handler:       _TestService_DoSomethingAgain_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DoSomethingForever",
			Handler:       _TestService_DoSomethingForever_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "desc_test_proto3.proto",
}