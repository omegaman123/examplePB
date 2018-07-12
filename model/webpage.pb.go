// Code generated by protoc-gen-go. DO NOT EDIT.
// source: webpage.proto

package webPage

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

type List_ItemType int32

const (
	List_INTS    List_ItemType = 0
	List_STRINGS List_ItemType = 1
	List_BOOLS   List_ItemType = 2
)

var List_ItemType_name = map[int32]string{
	0: "INTS",
	1: "STRINGS",
	2: "BOOLS",
}
var List_ItemType_value = map[string]int32{
	"INTS":    0,
	"STRINGS": 1,
	"BOOLS":   2,
}

func (x List_ItemType) String() string {
	return proto.EnumName(List_ItemType_name, int32(x))
}
func (List_ItemType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_webpage_8dec348df474d859, []int{0, 0}
}

type List struct {
	Title                string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	NumIterms            int32         `protobuf:"varint,2,opt,name=numIterms,proto3" json:"numIterms,omitempty"`
	Type                 List_ItemType `protobuf:"varint,3,opt,name=type,proto3,enum=webPage.List_ItemType" json:"type,omitempty"`
	Items                []*List_Item  `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *List) Reset()         { *m = List{} }
func (m *List) String() string { return proto.CompactTextString(m) }
func (*List) ProtoMessage()    {}
func (*List) Descriptor() ([]byte, []int) {
	return fileDescriptor_webpage_8dec348df474d859, []int{0}
}
func (m *List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_List.Unmarshal(m, b)
}
func (m *List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_List.Marshal(b, m, deterministic)
}
func (dst *List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List.Merge(dst, src)
}
func (m *List) XXX_Size() int {
	return xxx_messageInfo_List.Size(m)
}
func (m *List) XXX_DiscardUnknown() {
	xxx_messageInfo_List.DiscardUnknown(m)
}

var xxx_messageInfo_List proto.InternalMessageInfo

func (m *List) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *List) GetNumIterms() int32 {
	if m != nil {
		return m.NumIterms
	}
	return 0
}

func (m *List) GetType() List_ItemType {
	if m != nil {
		return m.Type
	}
	return List_INTS
}

func (m *List) GetItems() []*List_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type List_Item struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Switch               bool     `protobuf:"varint,3,opt,name=switch,proto3" json:"switch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *List_Item) Reset()         { *m = List_Item{} }
func (m *List_Item) String() string { return proto.CompactTextString(m) }
func (*List_Item) ProtoMessage()    {}
func (*List_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_webpage_8dec348df474d859, []int{0, 0}
}
func (m *List_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_List_Item.Unmarshal(m, b)
}
func (m *List_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_List_Item.Marshal(b, m, deterministic)
}
func (dst *List_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List_Item.Merge(dst, src)
}
func (m *List_Item) XXX_Size() int {
	return xxx_messageInfo_List_Item.Size(m)
}
func (m *List_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_List_Item.DiscardUnknown(m)
}

var xxx_messageInfo_List_Item proto.InternalMessageInfo

func (m *List_Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *List_Item) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *List_Item) GetSwitch() bool {
	if m != nil {
		return m.Switch
	}
	return false
}

type ListofLists struct {
	List                 []*List  `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListofLists) Reset()         { *m = ListofLists{} }
func (m *ListofLists) String() string { return proto.CompactTextString(m) }
func (*ListofLists) ProtoMessage()    {}
func (*ListofLists) Descriptor() ([]byte, []int) {
	return fileDescriptor_webpage_8dec348df474d859, []int{1}
}
func (m *ListofLists) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListofLists.Unmarshal(m, b)
}
func (m *ListofLists) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListofLists.Marshal(b, m, deterministic)
}
func (dst *ListofLists) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListofLists.Merge(dst, src)
}
func (m *ListofLists) XXX_Size() int {
	return xxx_messageInfo_ListofLists.Size(m)
}
func (m *ListofLists) XXX_DiscardUnknown() {
	xxx_messageInfo_ListofLists.DiscardUnknown(m)
}

var xxx_messageInfo_ListofLists proto.InternalMessageInfo

func (m *ListofLists) GetList() []*List {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*List)(nil), "webPage.List")
	proto.RegisterType((*List_Item)(nil), "webPage.List.Item")
	proto.RegisterType((*ListofLists)(nil), "webPage.ListofLists")
	proto.RegisterEnum("webPage.List_ItemType", List_ItemType_name, List_ItemType_value)
}

func init() { proto.RegisterFile("webpage.proto", fileDescriptor_webpage_8dec348df474d859) }

var fileDescriptor_webpage_8dec348df474d859 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xd1, 0x6a, 0x83, 0x30,
	0x14, 0x86, 0x17, 0x1b, 0x5b, 0x3d, 0xd2, 0x21, 0x87, 0x51, 0x64, 0xec, 0xc2, 0x79, 0x25, 0x63,
	0xc8, 0xe8, 0xde, 0x60, 0x37, 0xc3, 0x51, 0xda, 0x11, 0x7d, 0x01, 0x1d, 0x67, 0x9d, 0xd0, 0xa8,
	0x98, 0x14, 0xe9, 0x2b, 0xec, 0xa9, 0x47, 0x52, 0xcb, 0x18, 0xbb, 0x09, 0xf9, 0xfe, 0xfc, 0xfc,
	0x7c, 0x04, 0x96, 0x23, 0xd5, 0x7d, 0xb5, 0xa7, 0xac, 0x1f, 0x3a, 0xdd, 0xe1, 0x62, 0xa4, 0xfa,
	0xbd, 0xda, 0x53, 0xf2, 0xed, 0x00, 0xdf, 0x34, 0x4a, 0xe3, 0x0d, 0xb8, 0xba, 0xd1, 0x07, 0x8a,
	0x58, 0xcc, 0x52, 0x5f, 0x9c, 0x01, 0xef, 0xc0, 0x6f, 0x8f, 0x32, 0xd7, 0x34, 0x48, 0x15, 0x39,
	0x31, 0x4b, 0x5d, 0xf1, 0x1b, 0xe0, 0x03, 0x70, 0x7d, 0xea, 0x29, 0x9a, 0xc5, 0x2c, 0xbd, 0x5e,
	0xaf, 0xb2, 0x69, 0x34, 0x33, 0x83, 0x59, 0xae, 0x49, 0x96, 0xa7, 0x9e, 0x84, 0xed, 0x60, 0x0a,
	0x6e, 0xa3, 0x49, 0xaa, 0x88, 0xc7, 0xb3, 0x34, 0x58, 0xe3, 0xff, 0xb2, 0x38, 0x17, 0x6e, 0xdf,
	0x80, 0x1b, 0x44, 0x04, 0xde, 0x56, 0xf2, 0x22, 0x64, 0xef, 0xb8, 0x82, 0x79, 0x7b, 0x94, 0x35,
	0x0d, 0x93, 0xcc, 0x44, 0x26, 0x57, 0x63, 0xa3, 0x3f, 0xbe, 0xac, 0x8b, 0x27, 0x26, 0x4a, 0x1e,
	0xc1, 0xbb, 0x78, 0xa0, 0x07, 0x3c, 0xdf, 0x96, 0x45, 0x78, 0x85, 0x01, 0x2c, 0x8a, 0x52, 0xe4,
	0xdb, 0xd7, 0x22, 0x64, 0xe8, 0x83, 0xfb, 0xb2, 0xdb, 0x6d, 0x8a, 0xd0, 0x49, 0x9e, 0x20, 0x30,
	0x36, 0xdd, 0xa7, 0x39, 0x15, 0xde, 0x03, 0x3f, 0x34, 0x4a, 0x47, 0xcc, 0x1a, 0x2f, 0xff, 0x18,
	0x0b, 0xfb, 0x54, 0xcf, 0xed, 0x77, 0x3e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x2a, 0xcb,
	0xc1, 0x5f, 0x01, 0x00, 0x00,
}
