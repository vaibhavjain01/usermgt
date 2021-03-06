// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermgt/usermgt.proto

package usermgt_grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type NewUser struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUser) Reset()         { *m = NewUser{} }
func (m *NewUser) String() string { return proto.CompactTextString(m) }
func (*NewUser) ProtoMessage()    {}
func (*NewUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e0a379904753f4, []int{0}
}

func (m *NewUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUser.Unmarshal(m, b)
}
func (m *NewUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUser.Marshal(b, m, deterministic)
}
func (m *NewUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUser.Merge(m, src)
}
func (m *NewUser) XXX_Size() int {
	return xxx_messageInfo_NewUser.Size(m)
}
func (m *NewUser) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUser.DiscardUnknown(m)
}

var xxx_messageInfo_NewUser proto.InternalMessageInfo

func (m *NewUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewUser) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e0a379904753f4, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*NewUser)(nil), "usermgt.NewUser")
	proto.RegisterType((*User)(nil), "usermgt.User")
}

func init() { proto.RegisterFile("usermgt/usermgt.proto", fileDescriptor_17e0a379904753f4) }

var fileDescriptor_17e0a379904753f4 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2d, 0x4e, 0x2d,
	0xca, 0x4d, 0x2f, 0xd1, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x3e, 0x17, 0xbb, 0x5f, 0x6a, 0x79, 0x68, 0x71, 0x6a, 0x91, 0x90, 0x10, 0x17, 0x4b, 0x5e,
	0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc0, 0xc5, 0x9c,
	0x98, 0x9e, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x62, 0x2a, 0xd9, 0x70, 0xb1, 0x10,
	0xaf, 0x5a, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x19, 0x2c, 0xc0, 0x94, 0x99, 0x62, 0x64,
	0xcb, 0xc5, 0x0e, 0xd2, 0xed, 0x9b, 0x5e, 0x22, 0x64, 0xc4, 0xc5, 0xeb, 0x5c, 0x94, 0x9a, 0x58,
	0x92, 0x0a, 0xb3, 0x5f, 0x40, 0x0f, 0xe6, 0x46, 0xa8, 0x88, 0x14, 0x2f, 0x5c, 0x04, 0xc4, 0x55,
	0x62, 0x70, 0xd2, 0x8f, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0x2f, 0x4b, 0xcc, 0x4c, 0xca, 0x48, 0x2c, 0xcb, 0x4a, 0xcc, 0xcc, 0x33, 0x30, 0x84, 0x79, 0xd0,
	0x1a, 0x4a, 0xc7, 0xa7, 0x17, 0x15, 0x24, 0x27, 0xb1, 0x81, 0xbd, 0x6b, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xe3, 0xd7, 0x17, 0xe1, 0x07, 0x01, 0x00, 0x00,
}
