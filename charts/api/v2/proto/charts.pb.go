// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/charts.proto

package charts

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

type Request struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Range                string   `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_41214dd822a49c8a, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Request) GetRange() string {
	if m != nil {
		return m.Range
	}
	return ""
}

type Chart struct {
	Points               []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	MinTime              int64    `protobuf:"varint,2,opt,name=minTime,proto3" json:"minTime,omitempty"`
	MaxTime              int64    `protobuf:"varint,3,opt,name=maxTime,proto3" json:"maxTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chart) Reset()         { *m = Chart{} }
func (m *Chart) String() string { return proto.CompactTextString(m) }
func (*Chart) ProtoMessage()    {}
func (*Chart) Descriptor() ([]byte, []int) {
	return fileDescriptor_41214dd822a49c8a, []int{1}
}

func (m *Chart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chart.Unmarshal(m, b)
}
func (m *Chart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chart.Marshal(b, m, deterministic)
}
func (m *Chart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chart.Merge(m, src)
}
func (m *Chart) XXX_Size() int {
	return xxx_messageInfo_Chart.Size(m)
}
func (m *Chart) XXX_DiscardUnknown() {
	xxx_messageInfo_Chart.DiscardUnknown(m)
}

var xxx_messageInfo_Chart proto.InternalMessageInfo

func (m *Chart) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *Chart) GetMinTime() int64 {
	if m != nil {
		return m.MinTime
	}
	return 0
}

func (m *Chart) GetMaxTime() int64 {
	if m != nil {
		return m.MaxTime
	}
	return 0
}

type Point struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Value                float32  `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_41214dd822a49c8a, []int{2}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Point) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Chart)(nil), "Chart")
	proto.RegisterType((*Point)(nil), "Point")
}

func init() { proto.RegisterFile("proto/charts.proto", fileDescriptor_41214dd822a49c8a) }

var fileDescriptor_41214dd822a49c8a = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x49, 0x43, 0xdc, 0x72, 0x65, 0x3a, 0x31, 0x44, 0x0c, 0x55, 0xf1, 0xd4, 0x29, 0x88,
	0xf6, 0x11, 0x18, 0x22, 0xb6, 0xca, 0xb0, 0x31, 0x85, 0x62, 0xc0, 0xa2, 0xf5, 0x15, 0xfb, 0x5c,
	0xf1, 0xf8, 0xc8, 0xe7, 0xb0, 0x75, 0xbb, 0xef, 0xbe, 0xd3, 0xaf, 0xdf, 0x06, 0x3c, 0x06, 0x62,
	0xba, 0xdf, 0x7d, 0x0d, 0x81, 0x63, 0x27, 0xa0, 0x37, 0x30, 0x35, 0xf6, 0x27, 0xd9, 0xc8, 0x88,
	0x70, 0x99, 0x92, 0x7b, 0x6f, 0xab, 0x65, 0xb5, 0xba, 0x32, 0x32, 0xe3, 0x0d, 0x34, 0x61, 0xf0,
	0x9f, 0xb6, 0x9d, 0xc8, 0xb2, 0x80, 0x7e, 0x85, 0xe6, 0x31, 0x87, 0xe0, 0x02, 0xd4, 0x91, 0x9c,
	0xe7, 0xd8, 0x56, 0xcb, 0x7a, 0x35, 0x5f, 0xab, 0x6e, 0x9b, 0xd1, 0x8c, 0x5b, 0x6c, 0x61, 0x7a,
	0x70, 0xfe, 0xc5, 0x1d, 0x4a, 0x40, 0x6d, 0xfe, 0x51, 0xcc, 0xf0, 0x2b, 0xa6, 0x1e, 0x4d, 0x41,
	0xfd, 0x00, 0x8d, 0x84, 0xe4, 0x3e, 0x9c, 0x7d, 0x25, 0x5e, 0xe6, 0xdc, 0xe7, 0x34, 0xec, 0x53,
	0x89, 0x9b, 0x98, 0x02, 0x6b, 0x02, 0x25, 0x7d, 0x22, 0x2e, 0x60, 0xd6, 0x5b, 0x7e, 0x66, 0xda,
	0x7d, 0xe3, 0xac, 0x1b, 0x5f, 0x76, 0xab, 0x3a, 0xd1, 0xfa, 0x02, 0xef, 0x60, 0xde, 0x5b, 0x7e,
	0xf2, 0x27, 0x1b, 0x99, 0xc2, 0xd9, 0x13, 0x0d, 0xd7, 0xbd, 0xe5, 0x2d, 0x05, 0xfe, 0xa0, 0xbd,
	0xa3, 0x73, 0x37, 0x6f, 0x4a, 0x3e, 0x6f, 0xf3, 0x17, 0x00, 0x00, 0xff, 0xff, 0x19, 0xd9, 0x70,
	0x23, 0x52, 0x01, 0x00, 0x00,
}
