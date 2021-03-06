// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple_av.proto

package simple_av

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

type BIG_CMD int32

const (
	BIG_CMD_INVALID_BIG_CMD BIG_CMD = 0
	BIG_CMD_SIMPLE_AV       BIG_CMD = 1
)

var BIG_CMD_name = map[int32]string{
	0: "INVALID_BIG_CMD",
	1: "SIMPLE_AV",
}

var BIG_CMD_value = map[string]int32{
	"INVALID_BIG_CMD": 0,
	"SIMPLE_AV":       1,
}

func (x BIG_CMD) String() string {
	return proto.EnumName(BIG_CMD_name, int32(x))
}

func (BIG_CMD) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{0}
}

type SUB_CMD int32

const (
	SUB_CMD_InvalidAv SUB_CMD = 0
	SUB_CMD_JoinRoom  SUB_CMD = 1
	SUB_CMD_ExitRoom  SUB_CMD = 2
	SUB_CMD_Upload    SUB_CMD = 3
	SUB_CMD_SendData  SUB_CMD = 4
)

var SUB_CMD_name = map[int32]string{
	0: "InvalidAv",
	1: "JoinRoom",
	2: "ExitRoom",
	3: "Upload",
	4: "SendData",
}

var SUB_CMD_value = map[string]int32{
	"InvalidAv": 0,
	"JoinRoom":  1,
	"ExitRoom":  2,
	"Upload":    3,
	"SendData":  4,
}

func (x SUB_CMD) String() string {
	return proto.EnumName(SUB_CMD_name, int32(x))
}

func (SUB_CMD) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{1}
}

type JoinRoomReq struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomReq) Reset()         { *m = JoinRoomReq{} }
func (m *JoinRoomReq) String() string { return proto.CompactTextString(m) }
func (*JoinRoomReq) ProtoMessage()    {}
func (*JoinRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{0}
}

func (m *JoinRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomReq.Unmarshal(m, b)
}
func (m *JoinRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomReq.Marshal(b, m, deterministic)
}
func (m *JoinRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomReq.Merge(m, src)
}
func (m *JoinRoomReq) XXX_Size() int {
	return xxx_messageInfo_JoinRoomReq.Size(m)
}
func (m *JoinRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomReq proto.InternalMessageInfo

func (m *JoinRoomReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *JoinRoomReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type JoinRoomRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomRsp) Reset()         { *m = JoinRoomRsp{} }
func (m *JoinRoomRsp) String() string { return proto.CompactTextString(m) }
func (*JoinRoomRsp) ProtoMessage()    {}
func (*JoinRoomRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{1}
}

func (m *JoinRoomRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomRsp.Unmarshal(m, b)
}
func (m *JoinRoomRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomRsp.Marshal(b, m, deterministic)
}
func (m *JoinRoomRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomRsp.Merge(m, src)
}
func (m *JoinRoomRsp) XXX_Size() int {
	return xxx_messageInfo_JoinRoomRsp.Size(m)
}
func (m *JoinRoomRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomRsp.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomRsp proto.InternalMessageInfo

type ExitRoomReq struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitRoomReq) Reset()         { *m = ExitRoomReq{} }
func (m *ExitRoomReq) String() string { return proto.CompactTextString(m) }
func (*ExitRoomReq) ProtoMessage()    {}
func (*ExitRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{2}
}

func (m *ExitRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitRoomReq.Unmarshal(m, b)
}
func (m *ExitRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitRoomReq.Marshal(b, m, deterministic)
}
func (m *ExitRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitRoomReq.Merge(m, src)
}
func (m *ExitRoomReq) XXX_Size() int {
	return xxx_messageInfo_ExitRoomReq.Size(m)
}
func (m *ExitRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_ExitRoomReq proto.InternalMessageInfo

func (m *ExitRoomReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *ExitRoomReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type ExitRoomRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitRoomRsp) Reset()         { *m = ExitRoomRsp{} }
func (m *ExitRoomRsp) String() string { return proto.CompactTextString(m) }
func (*ExitRoomRsp) ProtoMessage()    {}
func (*ExitRoomRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{3}
}

func (m *ExitRoomRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitRoomRsp.Unmarshal(m, b)
}
func (m *ExitRoomRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitRoomRsp.Marshal(b, m, deterministic)
}
func (m *ExitRoomRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitRoomRsp.Merge(m, src)
}
func (m *ExitRoomRsp) XXX_Size() int {
	return xxx_messageInfo_ExitRoomRsp.Size(m)
}
func (m *ExitRoomRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitRoomRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ExitRoomRsp proto.InternalMessageInfo

type UploadReq struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadReq) Reset()         { *m = UploadReq{} }
func (m *UploadReq) String() string { return proto.CompactTextString(m) }
func (*UploadReq) ProtoMessage()    {}
func (*UploadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{4}
}

func (m *UploadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadReq.Unmarshal(m, b)
}
func (m *UploadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadReq.Marshal(b, m, deterministic)
}
func (m *UploadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadReq.Merge(m, src)
}
func (m *UploadReq) XXX_Size() int {
	return xxx_messageInfo_UploadReq.Size(m)
}
func (m *UploadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadReq proto.InternalMessageInfo

func (m *UploadReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *UploadReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type UploadRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRsp) Reset()         { *m = UploadRsp{} }
func (m *UploadRsp) String() string { return proto.CompactTextString(m) }
func (*UploadRsp) ProtoMessage()    {}
func (*UploadRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{5}
}

func (m *UploadRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRsp.Unmarshal(m, b)
}
func (m *UploadRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRsp.Marshal(b, m, deterministic)
}
func (m *UploadRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRsp.Merge(m, src)
}
func (m *UploadRsp) XXX_Size() int {
	return xxx_messageInfo_UploadRsp.Size(m)
}
func (m *UploadRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRsp proto.InternalMessageInfo

type SendDataReq struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendDataReq) Reset()         { *m = SendDataReq{} }
func (m *SendDataReq) String() string { return proto.CompactTextString(m) }
func (*SendDataReq) ProtoMessage()    {}
func (*SendDataReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{6}
}

func (m *SendDataReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendDataReq.Unmarshal(m, b)
}
func (m *SendDataReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendDataReq.Marshal(b, m, deterministic)
}
func (m *SendDataReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendDataReq.Merge(m, src)
}
func (m *SendDataReq) XXX_Size() int {
	return xxx_messageInfo_SendDataReq.Size(m)
}
func (m *SendDataReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendDataReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendDataReq proto.InternalMessageInfo

func (m *SendDataReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *SendDataReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SendDataReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SendDataRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendDataRsp) Reset()         { *m = SendDataRsp{} }
func (m *SendDataRsp) String() string { return proto.CompactTextString(m) }
func (*SendDataRsp) ProtoMessage()    {}
func (*SendDataRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d847d8717ecca288, []int{7}
}

func (m *SendDataRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendDataRsp.Unmarshal(m, b)
}
func (m *SendDataRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendDataRsp.Marshal(b, m, deterministic)
}
func (m *SendDataRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendDataRsp.Merge(m, src)
}
func (m *SendDataRsp) XXX_Size() int {
	return xxx_messageInfo_SendDataRsp.Size(m)
}
func (m *SendDataRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendDataRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SendDataRsp proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("simple_av.BIG_CMD", BIG_CMD_name, BIG_CMD_value)
	proto.RegisterEnum("simple_av.SUB_CMD", SUB_CMD_name, SUB_CMD_value)
	proto.RegisterType((*JoinRoomReq)(nil), "simple_av.JoinRoomReq")
	proto.RegisterType((*JoinRoomRsp)(nil), "simple_av.JoinRoomRsp")
	proto.RegisterType((*ExitRoomReq)(nil), "simple_av.ExitRoomReq")
	proto.RegisterType((*ExitRoomRsp)(nil), "simple_av.ExitRoomRsp")
	proto.RegisterType((*UploadReq)(nil), "simple_av.UploadReq")
	proto.RegisterType((*UploadRsp)(nil), "simple_av.UploadRsp")
	proto.RegisterType((*SendDataReq)(nil), "simple_av.SendDataReq")
	proto.RegisterType((*SendDataRsp)(nil), "simple_av.SendDataRsp")
}

func init() { proto.RegisterFile("simple_av.proto", fileDescriptor_d847d8717ecca288) }

var fileDescriptor_d847d8717ecca288 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x6b, 0xf2, 0x40,
	0x10, 0x35, 0x46, 0xf4, 0xcb, 0x44, 0x71, 0x99, 0xaf, 0xb4, 0x92, 0x93, 0xe4, 0x24, 0x42, 0x3d,
	0xb4, 0x20, 0x3d, 0xf4, 0x12, 0xab, 0x94, 0x14, 0x95, 0x92, 0xa0, 0x87, 0x5e, 0xc2, 0x96, 0xcd,
	0x61, 0x41, 0xdd, 0x6d, 0x12, 0x43, 0x7f, 0x40, 0x7f, 0x78, 0xd9, 0xb4, 0x1b, 0x23, 0xe4, 0x62,
	0x6f, 0xbb, 0x6f, 0xf7, 0xbd, 0x37, 0x33, 0x6f, 0xa0, 0x9f, 0xf2, 0xbd, 0xdc, 0xc5, 0x11, 0xcd,
	0x27, 0x32, 0x11, 0x99, 0x40, 0xab, 0x04, 0xdc, 0x07, 0xb0, 0x5f, 0x04, 0x3f, 0x04, 0x42, 0xec,
	0x83, 0xf8, 0x03, 0x6f, 0xa0, 0x93, 0x08, 0xb1, 0x8f, 0x38, 0x1b, 0x18, 0x43, 0x63, 0x64, 0x06,
	0x6d, 0x75, 0xf5, 0x19, 0x12, 0x30, 0x8f, 0x9c, 0x0d, 0x9a, 0x05, 0xa8, 0x8e, 0x6e, 0xaf, 0xc2,
	0x4c, 0xa5, 0x12, 0x5a, 0x7c, 0xf2, 0xec, 0x6f, 0x42, 0x25, 0x33, 0x95, 0xee, 0x14, 0xac, 0x8d,
	0xdc, 0x09, 0xca, 0x2e, 0x94, 0xb1, 0x4b, 0x5e, 0x2a, 0xdd, 0x25, 0xd8, 0x61, 0x7c, 0x60, 0x73,
	0x9a, 0xd1, 0xcb, 0x64, 0x10, 0xa1, 0xc5, 0x68, 0x46, 0x07, 0xe6, 0xd0, 0x18, 0x75, 0x83, 0xe2,
	0xac, 0x2a, 0x2c, 0xd5, 0x52, 0x39, 0xbe, 0x85, 0xce, 0xcc, 0x7f, 0x8e, 0x9e, 0x56, 0x73, 0xfc,
	0x0f, 0x7d, 0x7f, 0xbd, 0xf5, 0x96, 0xfe, 0x3c, 0xfa, 0x85, 0x48, 0x03, 0x7b, 0x60, 0x85, 0xfe,
	0xea, 0x75, 0xb9, 0x88, 0xbc, 0x2d, 0x31, 0xc6, 0x6b, 0xe8, 0x84, 0x9b, 0x59, 0xf1, 0xbd, 0x07,
	0x96, 0x7f, 0xc8, 0xe9, 0x8e, 0x33, 0x2f, 0x27, 0x0d, 0xec, 0xc2, 0x3f, 0x3d, 0x42, 0x62, 0xa8,
	0x9b, 0x9e, 0x03, 0x69, 0x22, 0x40, 0xfb, 0xa7, 0x1d, 0x62, 0xaa, 0x17, 0xed, 0x4f, 0x5a, 0x77,
	0x5f, 0x4d, 0x20, 0xde, 0x91, 0x71, 0xb1, 0xe5, 0x2c, 0x16, 0x61, 0x9c, 0xe4, 0x71, 0x82, 0x8f,
	0x27, 0x29, 0xbc, 0x9e, 0x9c, 0x02, 0xaf, 0x84, 0xeb, 0xd4, 0xe2, 0xa9, 0x74, 0x1b, 0x8a, 0xad,
	0xad, 0xcf, 0xd8, 0x95, 0x44, 0x9d, 0x5a, 0xbc, 0x60, 0x4f, 0x75, 0xa9, 0x78, 0x55, 0xf9, 0x53,
	0x86, 0xe8, 0xd4, 0xa0, 0xda, 0x55, 0xb7, 0x75, 0xe6, 0x5a, 0x49, 0xce, 0xa9, 0xc5, 0x15, 0x7b,
	0x66, 0xbf, 0x9d, 0xd6, 0xf8, 0xbd, 0x5d, 0x2c, 0xf6, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x32, 0x2c, 0x67, 0x3f, 0xeb, 0x02, 0x00, 0x00,
}
