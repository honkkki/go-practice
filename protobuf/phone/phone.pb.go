// Code generated by protoc-gen-go. DO NOT EDIT.
// source: phone.proto

//包名，通过protoc生成时go文件时

package phone

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

//手机类型
//枚举类型第一个字段必须为0
type PhoneType int32

const (
	PhoneType_HOME PhoneType = 0
	PhoneType_WORK PhoneType = 1
)

var PhoneType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
}

var PhoneType_value = map[string]int32{
	"HOME": 0,
	"WORK": 1,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}

func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0e585c54dbe5137, []int{0}
}

//手机
type Phone struct {
	Type                 PhoneType `protobuf:"varint,1,opt,name=type,proto3,enum=phone.PhoneType" json:"type,omitempty"`
	Number               string    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e585c54dbe5137, []int{0}
}

func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (m *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(m, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_HOME
}

func (m *Phone) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

//人
type Person struct {
	//后面的数字表示标识号
	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//repeated表示可重复
	//可以有多个手机
	Phone                []*Phone `protobuf:"bytes,3,rep,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e585c54dbe5137, []int{1}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetPhone() []*Phone {
	if m != nil {
		return m.Phone
	}
	return nil
}

//联系簿
type ContactBook struct {
	Person               []*Person `protobuf:"bytes,1,rep,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ContactBook) Reset()         { *m = ContactBook{} }
func (m *ContactBook) String() string { return proto.CompactTextString(m) }
func (*ContactBook) ProtoMessage()    {}
func (*ContactBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e585c54dbe5137, []int{2}
}

func (m *ContactBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactBook.Unmarshal(m, b)
}
func (m *ContactBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactBook.Marshal(b, m, deterministic)
}
func (m *ContactBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactBook.Merge(m, src)
}
func (m *ContactBook) XXX_Size() int {
	return xxx_messageInfo_ContactBook.Size(m)
}
func (m *ContactBook) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactBook.DiscardUnknown(m)
}

var xxx_messageInfo_ContactBook proto.InternalMessageInfo

func (m *ContactBook) GetPerson() []*Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func init() {
	proto.RegisterEnum("phone.PhoneType", PhoneType_name, PhoneType_value)
	proto.RegisterType((*Phone)(nil), "phone.Phone")
	proto.RegisterType((*Person)(nil), "phone.Person")
	proto.RegisterType((*ContactBook)(nil), "phone.ContactBook")
}

func init() { proto.RegisterFile("phone.proto", fileDescriptor_c0e585c54dbe5137) }

var fileDescriptor_c0e585c54dbe5137 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x4f, 0x4f, 0x4b, 0x87, 0x40,
	0x14, 0x6c, 0xfd, 0xb3, 0xe4, 0xb3, 0x44, 0xde, 0x21, 0xf6, 0x96, 0x48, 0x81, 0x74, 0xf0, 0x60,
	0x7d, 0x82, 0x42, 0x08, 0x22, 0x94, 0x25, 0xe8, 0xac, 0xb9, 0x90, 0x84, 0xbb, 0x8b, 0x6d, 0x07,
	0xbf, 0x7d, 0xf8, 0xdc, 0xe2, 0x77, 0x9b, 0xd9, 0x99, 0xd9, 0x99, 0x07, 0xa9, 0xfd, 0x34, 0x5a,
	0xd5, 0x76, 0x35, 0xce, 0x60, 0x4c, 0xa4, 0x6c, 0x21, 0xee, 0x77, 0x80, 0x37, 0x10, 0xb9, 0xcd,
	0x2a, 0xc1, 0x0a, 0x56, 0x65, 0x4d, 0x5e, 0x1f, 0x5e, 0xd2, 0xde, 0x36, 0xab, 0x24, 0xa9, 0x78,
	0x05, 0x5c, 0xff, 0x2c, 0xa3, 0x5a, 0x45, 0x50, 0xb0, 0x2a, 0x91, 0x9e, 0x95, 0x3d, 0xf0, 0x5e,
	0xad, 0xdf, 0x46, 0x63, 0x06, 0xc1, 0x3c, 0xd1, 0x2f, 0xb1, 0x0c, 0xe6, 0x09, 0x11, 0x22, 0x3d,
	0x2c, 0xca, 0xfb, 0x09, 0x63, 0x09, 0x47, 0xbb, 0x08, 0x8b, 0xb0, 0x4a, 0x9b, 0x8b, 0xd3, 0x32,
	0xe9, 0x87, 0x3d, 0x40, 0xfa, 0x64, 0xb4, 0x1b, 0x3e, 0xdc, 0xa3, 0x31, 0x5f, 0x78, 0x0b, 0xdc,
	0x52, 0x81, 0x60, 0x94, 0xb9, 0xfc, 0xcb, 0xd0, 0xa3, 0xf4, 0xe2, 0xdd, 0x35, 0x24, 0xff, 0x93,
	0xf1, 0x1c, 0xa2, 0xe7, 0xee, 0xb5, 0xcd, 0xcf, 0x76, 0xf4, 0xde, 0xc9, 0x97, 0x9c, 0x8d, 0x9c,
	0xae, 0xbf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x73, 0xf9, 0xcb, 0x0c, 0x01, 0x00, 0x00,
}