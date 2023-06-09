// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store.proto

package pb

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

type LoginPasswordData struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	LastModified         int64    `protobuf:"varint,4,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginPasswordData) Reset()         { *m = LoginPasswordData{} }
func (m *LoginPasswordData) String() string { return proto.CompactTextString(m) }
func (*LoginPasswordData) ProtoMessage()    {}
func (*LoginPasswordData) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{0}
}

func (m *LoginPasswordData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginPasswordData.Unmarshal(m, b)
}
func (m *LoginPasswordData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginPasswordData.Marshal(b, m, deterministic)
}
func (m *LoginPasswordData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginPasswordData.Merge(m, src)
}
func (m *LoginPasswordData) XXX_Size() int {
	return xxx_messageInfo_LoginPasswordData.Size(m)
}
func (m *LoginPasswordData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginPasswordData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginPasswordData proto.InternalMessageInfo

func (m *LoginPasswordData) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *LoginPasswordData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginPasswordData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LoginPasswordData) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

type TextData struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	LastModified         int64    `protobuf:"varint,3,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextData) Reset()         { *m = TextData{} }
func (m *TextData) String() string { return proto.CompactTextString(m) }
func (*TextData) ProtoMessage()    {}
func (*TextData) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{1}
}

func (m *TextData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextData.Unmarshal(m, b)
}
func (m *TextData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextData.Marshal(b, m, deterministic)
}
func (m *TextData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextData.Merge(m, src)
}
func (m *TextData) XXX_Size() int {
	return xxx_messageInfo_TextData.Size(m)
}
func (m *TextData) XXX_DiscardUnknown() {
	xxx_messageInfo_TextData.DiscardUnknown(m)
}

var xxx_messageInfo_TextData proto.InternalMessageInfo

func (m *TextData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *TextData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TextData) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

type BinaryData struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	LastModified         int64    `protobuf:"varint,3,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BinaryData) Reset()         { *m = BinaryData{} }
func (m *BinaryData) String() string { return proto.CompactTextString(m) }
func (*BinaryData) ProtoMessage()    {}
func (*BinaryData) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{2}
}

func (m *BinaryData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BinaryData.Unmarshal(m, b)
}
func (m *BinaryData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BinaryData.Marshal(b, m, deterministic)
}
func (m *BinaryData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinaryData.Merge(m, src)
}
func (m *BinaryData) XXX_Size() int {
	return xxx_messageInfo_BinaryData.Size(m)
}
func (m *BinaryData) XXX_DiscardUnknown() {
	xxx_messageInfo_BinaryData.DiscardUnknown(m)
}

var xxx_messageInfo_BinaryData proto.InternalMessageInfo

func (m *BinaryData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BinaryData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BinaryData) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

type BankCardData struct {
	Number               string   `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	ValidThru            string   `protobuf:"bytes,2,opt,name=valid_thru,json=validThru,proto3" json:"valid_thru,omitempty"`
	Cvv                  string   `protobuf:"bytes,3,opt,name=cvv,proto3" json:"cvv,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	LastModified         int64    `protobuf:"varint,5,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankCardData) Reset()         { *m = BankCardData{} }
func (m *BankCardData) String() string { return proto.CompactTextString(m) }
func (*BankCardData) ProtoMessage()    {}
func (*BankCardData) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{3}
}

func (m *BankCardData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankCardData.Unmarshal(m, b)
}
func (m *BankCardData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankCardData.Marshal(b, m, deterministic)
}
func (m *BankCardData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankCardData.Merge(m, src)
}
func (m *BankCardData) XXX_Size() int {
	return xxx_messageInfo_BankCardData.Size(m)
}
func (m *BankCardData) XXX_DiscardUnknown() {
	xxx_messageInfo_BankCardData.DiscardUnknown(m)
}

var xxx_messageInfo_BankCardData proto.InternalMessageInfo

func (m *BankCardData) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *BankCardData) GetValidThru() string {
	if m != nil {
		return m.ValidThru
	}
	return ""
}

func (m *BankCardData) GetCvv() string {
	if m != nil {
		return m.Cvv
	}
	return ""
}

func (m *BankCardData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BankCardData) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

type Key struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{4}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DataArray struct {
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataArray) Reset()         { *m = DataArray{} }
func (m *DataArray) String() string { return proto.CompactTextString(m) }
func (*DataArray) ProtoMessage()    {}
func (*DataArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{5}
}

func (m *DataArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataArray.Unmarshal(m, b)
}
func (m *DataArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataArray.Marshal(b, m, deterministic)
}
func (m *DataArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataArray.Merge(m, src)
}
func (m *DataArray) XXX_Size() int {
	return xxx_messageInfo_DataArray.Size(m)
}
func (m *DataArray) XXX_DiscardUnknown() {
	xxx_messageInfo_DataArray.DiscardUnknown(m)
}

var xxx_messageInfo_DataArray proto.InternalMessageInfo

func (m *DataArray) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type Value struct {
	// Types that are valid to be assigned to Kind:
	//	*Value_LoginPassword
	//	*Value_Text
	//	*Value_BinData
	//	*Value_CardData
	Kind                 isValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{6}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Kind interface {
	isValue_Kind()
}

type Value_LoginPassword struct {
	LoginPassword *LoginPasswordData `protobuf:"bytes,1,opt,name=login_password,json=loginPassword,proto3,oneof"`
}

type Value_Text struct {
	Text *TextData `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

type Value_BinData struct {
	BinData *BinaryData `protobuf:"bytes,3,opt,name=bin_data,json=binData,proto3,oneof"`
}

type Value_CardData struct {
	CardData *BankCardData `protobuf:"bytes,4,opt,name=card_data,json=cardData,proto3,oneof"`
}

func (*Value_LoginPassword) isValue_Kind() {}

func (*Value_Text) isValue_Kind() {}

func (*Value_BinData) isValue_Kind() {}

func (*Value_CardData) isValue_Kind() {}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Value) GetLoginPassword() *LoginPasswordData {
	if x, ok := m.GetKind().(*Value_LoginPassword); ok {
		return x.LoginPassword
	}
	return nil
}

func (m *Value) GetText() *TextData {
	if x, ok := m.GetKind().(*Value_Text); ok {
		return x.Text
	}
	return nil
}

func (m *Value) GetBinData() *BinaryData {
	if x, ok := m.GetKind().(*Value_BinData); ok {
		return x.BinData
	}
	return nil
}

func (m *Value) GetCardData() *BankCardData {
	if x, ok := m.GetKind().(*Value_CardData); ok {
		return x.CardData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_LoginPassword)(nil),
		(*Value_Text)(nil),
		(*Value_BinData)(nil),
		(*Value_CardData)(nil),
	}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{7}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LoginPasswordData)(nil), "pb.LoginPasswordData")
	proto.RegisterType((*TextData)(nil), "pb.TextData")
	proto.RegisterType((*BinaryData)(nil), "pb.BinaryData")
	proto.RegisterType((*BankCardData)(nil), "pb.BankCardData")
	proto.RegisterType((*Key)(nil), "pb.Key")
	proto.RegisterType((*DataArray)(nil), "pb.DataArray")
	proto.RegisterType((*Value)(nil), "pb.Value")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

func init() {
	proto.RegisterFile("store.proto", fileDescriptor_98bbca36ef968dfc)
}

var fileDescriptor_98bbca36ef968dfc = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x6e, 0x9a, 0xb4, 0x69, 0x4e, 0xda, 0x65, 0x1d, 0xfc, 0x09, 0x05, 0x21, 0xc6, 0x9b, 0xca,
	0x42, 0x84, 0x7a, 0x29, 0x08, 0xd6, 0x15, 0x0a, 0xab, 0x20, 0xe9, 0xea, 0x85, 0x37, 0x65, 0x92,
	0x8c, 0xed, 0xd0, 0x74, 0x12, 0x66, 0x26, 0x75, 0xf3, 0x0e, 0x3e, 0x83, 0x6f, 0xe0, 0xc3, 0xf8,
	0x46, 0x32, 0x93, 0x71, 0x5b, 0x76, 0x7b, 0xd1, 0x8b, 0xbd, 0x3b, 0xe7, 0x3b, 0xdf, 0x9c, 0xf3,
	0xe5, 0x3b, 0x27, 0xe0, 0x0b, 0x59, 0x72, 0x12, 0x57, 0xbc, 0x94, 0x25, 0xea, 0x56, 0xe9, 0x18,
	0x70, 0x2d, 0xd7, 0x6d, 0x1e, 0xfd, 0xb2, 0xe0, 0xd1, 0xa7, 0x72, 0x45, 0xd9, 0x17, 0x2c, 0xc4,
	0xcf, 0x92, 0xe7, 0x97, 0x58, 0x62, 0xf4, 0x18, 0x7a, 0x85, 0x02, 0x03, 0x2b, 0xb4, 0x26, 0x5e,
	0xd2, 0x26, 0x68, 0x0c, 0x83, 0xca, 0xb0, 0x82, 0xae, 0x2e, 0xdc, 0xe6, 0x28, 0x04, 0x3f, 0x27,
	0x22, 0xe3, 0xb4, 0x92, 0xb4, 0x64, 0x81, 0xad, 0xcb, 0x87, 0x10, 0x7a, 0x09, 0xa3, 0x02, 0x0b,
	0xb9, 0xdc, 0x96, 0x39, 0xfd, 0x41, 0x49, 0x1e, 0x38, 0xa1, 0x35, 0xb1, 0x93, 0xa1, 0x02, 0x3f,
	0x1b, 0x2c, 0x22, 0x30, 0xb8, 0x26, 0x37, 0x52, 0x8b, 0x40, 0xe0, 0xe4, 0x58, 0x62, 0xa3, 0x41,
	0xc7, 0x77, 0xc7, 0x74, 0x4f, 0x18, 0x63, 0x1f, 0x19, 0xb3, 0x02, 0x98, 0x51, 0x86, 0x79, 0x73,
	0x6f, 0xd0, 0xf0, 0x61, 0x07, 0xfd, 0xb6, 0x60, 0x38, 0xc3, 0x6c, 0xf3, 0x01, 0x1b, 0x67, 0x9f,
	0x42, 0x9f, 0xd5, 0xdb, 0x94, 0x70, 0xf3, 0x59, 0x26, 0x43, 0xcf, 0x01, 0x76, 0xb8, 0xa0, 0xf9,
	0x52, 0xae, 0x79, 0x6d, 0xc6, 0x79, 0x1a, 0xb9, 0x5e, 0xf3, 0x1a, 0x9d, 0x83, 0x9d, 0xed, 0x76,
	0xc6, 0x56, 0x15, 0xde, 0x15, 0xe8, 0x9c, 0x20, 0xb0, 0x77, 0x44, 0xe0, 0x33, 0xb0, 0xaf, 0x48,
	0xa3, 0xfa, 0x6f, 0x48, 0x63, 0x34, 0xa9, 0x30, 0x8a, 0xc1, 0x53, 0x82, 0xdf, 0x73, 0x8e, 0x1b,
	0xf4, 0x02, 0xfa, 0x3b, 0x5c, 0xd4, 0x44, 0x04, 0x56, 0x68, 0x4f, 0xfc, 0xa9, 0x17, 0x57, 0x69,
	0xfc, 0x4d, 0x21, 0x89, 0x29, 0x44, 0x7f, 0x2d, 0xe8, 0x69, 0x04, 0xbd, 0x83, 0x33, 0x7d, 0x2f,
	0xcb, 0xdb, 0x63, 0x51, 0x6d, 0xfd, 0xe9, 0x13, 0xf5, 0xe8, 0xde, 0xad, 0xcd, 0x3b, 0xc9, 0xa8,
	0x38, 0x04, 0x51, 0x04, 0x8e, 0x24, 0x37, 0x52, 0x9b, 0xe0, 0x4f, 0x87, 0xea, 0xd5, 0xff, 0x9b,
	0x98, 0x77, 0x12, 0x5d, 0x43, 0x17, 0x30, 0x48, 0x29, 0x5b, 0xea, 0xb5, 0xd9, 0x9a, 0x77, 0xa6,
	0x78, 0xfb, 0xa5, 0xce, 0x3b, 0x89, 0x9b, 0x52, 0xa6, 0x3d, 0x7f, 0x0d, 0x5e, 0x86, 0x79, 0xde,
	0xb2, 0x1d, 0xcd, 0x3e, 0xd7, 0xec, 0x83, 0xc5, 0xcc, 0x3b, 0xc9, 0x20, 0x33, 0xf1, 0xac, 0x0f,
	0xce, 0x86, 0xb2, 0x3c, 0x72, 0xa1, 0xf7, 0x71, 0x5b, 0xc9, 0x66, 0xfa, 0xc7, 0x02, 0x77, 0x21,
	0x4b, 0x8e, 0x57, 0x04, 0x4d, 0xc0, 0x5d, 0x90, 0xf6, 0x42, 0xf7, 0x36, 0x8c, 0x91, 0x0a, 0x13,
	0x22, 0xaa, 0x92, 0x09, 0xb2, 0x90, 0x58, 0xd6, 0x02, 0x85, 0xe0, 0x2c, 0x1a, 0x96, 0xb5, 0x34,
	0xdd, 0x68, 0x3c, 0x52, 0xe1, 0xde, 0xd7, 0x0b, 0x80, 0xaf, 0x55, 0x8e, 0x25, 0x39, 0xa5, 0xdd,
	0x2b, 0x80, 0x4b, 0x52, 0x10, 0x43, 0x76, 0x15, 0xe3, 0x8a, 0x34, 0xc7, 0xa8, 0xb3, 0xfe, 0x77,
	0x27, 0x7e, 0x5b, 0xa5, 0x69, 0x5f, 0xff, 0xe4, 0x6f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe5,
	0x3c, 0x33, 0x52, 0x03, 0x04, 0x00, 0x00,
}
