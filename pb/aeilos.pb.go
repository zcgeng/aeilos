// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aeilos.proto

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

type TouchType int32

const (
	TouchType_FLIP TouchType = 0
	TouchType_FLAG TouchType = 1
)

var TouchType_name = map[int32]string{
	0: "FLIP",
	1: "FLAG",
}

var TouchType_value = map[string]int32{
	"FLIP": 0,
	"FLAG": 1,
}

func (x TouchType) String() string {
	return proto.EnumName(TouchType_name, int32(x))
}

func (TouchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{0}
}

type TouchResultType int32

const (
	TouchResultType_CORRECT TouchResultType = 0
	TouchResultType_WRONG   TouchResultType = 1
	TouchResultType_ABORT   TouchResultType = 2
)

var TouchResultType_name = map[int32]string{
	0: "CORRECT",
	1: "WRONG",
	2: "ABORT",
}

var TouchResultType_value = map[string]int32{
	"CORRECT": 0,
	"WRONG":   1,
	"ABORT":   2,
}

func (x TouchResultType) String() string {
	return proto.EnumName(TouchResultType_name, int32(x))
}

func (TouchResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{1}
}

type Cell struct {
	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	// Types that are valid to be assigned to CellType:
	//	*Cell_Bombs
	//	*Cell_FlagURL
	//	*Cell_UnTouched
	CellType             isCell_CellType `protobuf_oneof:"CellType"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Cell) Reset()         { *m = Cell{} }
func (m *Cell) String() string { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()    {}
func (*Cell) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{0}
}

func (m *Cell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cell.Unmarshal(m, b)
}
func (m *Cell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cell.Marshal(b, m, deterministic)
}
func (m *Cell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell.Merge(m, src)
}
func (m *Cell) XXX_Size() int {
	return xxx_messageInfo_Cell.Size(m)
}
func (m *Cell) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell.DiscardUnknown(m)
}

var xxx_messageInfo_Cell proto.InternalMessageInfo

func (m *Cell) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Cell) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type isCell_CellType interface {
	isCell_CellType()
}

type Cell_Bombs struct {
	Bombs int32 `protobuf:"varint,3,opt,name=bombs,proto3,oneof"`
}

type Cell_FlagURL struct {
	FlagURL string `protobuf:"bytes,4,opt,name=flagURL,proto3,oneof"`
}

type Cell_UnTouched struct {
	UnTouched bool `protobuf:"varint,5,opt,name=unTouched,proto3,oneof"`
}

func (*Cell_Bombs) isCell_CellType() {}

func (*Cell_FlagURL) isCell_CellType() {}

func (*Cell_UnTouched) isCell_CellType() {}

func (m *Cell) GetCellType() isCell_CellType {
	if m != nil {
		return m.CellType
	}
	return nil
}

func (m *Cell) GetBombs() int32 {
	if x, ok := m.GetCellType().(*Cell_Bombs); ok {
		return x.Bombs
	}
	return 0
}

func (m *Cell) GetFlagURL() string {
	if x, ok := m.GetCellType().(*Cell_FlagURL); ok {
		return x.FlagURL
	}
	return ""
}

func (m *Cell) GetUnTouched() bool {
	if x, ok := m.GetCellType().(*Cell_UnTouched); ok {
		return x.UnTouched
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Cell) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Cell_Bombs)(nil),
		(*Cell_FlagURL)(nil),
		(*Cell_UnTouched)(nil),
	}
}

type Area struct {
	// 10 * 10 cells starting from (x, y) : 0,0 1,0 2,0 ... 9,0 0,-1 1,-1...
	X                    int64    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int64    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Cells                []*Cell  `protobuf:"bytes,3,rep,name=cells,proto3" json:"cells,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Area) Reset()         { *m = Area{} }
func (m *Area) String() string { return proto.CompactTextString(m) }
func (*Area) ProtoMessage()    {}
func (*Area) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{1}
}

func (m *Area) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Area.Unmarshal(m, b)
}
func (m *Area) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Area.Marshal(b, m, deterministic)
}
func (m *Area) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Area.Merge(m, src)
}
func (m *Area) XXX_Size() int {
	return xxx_messageInfo_Area.Size(m)
}
func (m *Area) XXX_DiscardUnknown() {
	xxx_messageInfo_Area.DiscardUnknown(m)
}

var xxx_messageInfo_Area proto.InternalMessageInfo

func (m *Area) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Area) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Area) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type TouchRequest struct {
	X                    int64     `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int64     `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	TouchType            TouchType `protobuf:"varint,3,opt,name=touchType,proto3,enum=pb.TouchType" json:"touchType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TouchRequest) Reset()         { *m = TouchRequest{} }
func (m *TouchRequest) String() string { return proto.CompactTextString(m) }
func (*TouchRequest) ProtoMessage()    {}
func (*TouchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{2}
}

func (m *TouchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TouchRequest.Unmarshal(m, b)
}
func (m *TouchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TouchRequest.Marshal(b, m, deterministic)
}
func (m *TouchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TouchRequest.Merge(m, src)
}
func (m *TouchRequest) XXX_Size() int {
	return xxx_messageInfo_TouchRequest.Size(m)
}
func (m *TouchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TouchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TouchRequest proto.InternalMessageInfo

func (m *TouchRequest) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TouchRequest) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *TouchRequest) GetTouchType() TouchType {
	if m != nil {
		return m.TouchType
	}
	return TouchType_FLIP
}

type TouchResponse struct {
	Score                int32           `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	Cell                 *Cell           `protobuf:"bytes,2,opt,name=cell,proto3" json:"cell,omitempty"`
	TouchResult          TouchResultType `protobuf:"varint,3,opt,name=touchResult,proto3,enum=pb.TouchResultType" json:"touchResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TouchResponse) Reset()         { *m = TouchResponse{} }
func (m *TouchResponse) String() string { return proto.CompactTextString(m) }
func (*TouchResponse) ProtoMessage()    {}
func (*TouchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{3}
}

func (m *TouchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TouchResponse.Unmarshal(m, b)
}
func (m *TouchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TouchResponse.Marshal(b, m, deterministic)
}
func (m *TouchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TouchResponse.Merge(m, src)
}
func (m *TouchResponse) XXX_Size() int {
	return xxx_messageInfo_TouchResponse.Size(m)
}
func (m *TouchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TouchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TouchResponse proto.InternalMessageInfo

func (m *TouchResponse) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *TouchResponse) GetCell() *Cell {
	if m != nil {
		return m.Cell
	}
	return nil
}

func (m *TouchResponse) GetTouchResult() TouchResultType {
	if m != nil {
		return m.TouchResult
	}
	return TouchResultType_CORRECT
}

type XY struct {
	X                    int64    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int64    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XY) Reset()         { *m = XY{} }
func (m *XY) String() string { return proto.CompactTextString(m) }
func (*XY) ProtoMessage()    {}
func (*XY) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{4}
}

func (m *XY) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XY.Unmarshal(m, b)
}
func (m *XY) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XY.Marshal(b, m, deterministic)
}
func (m *XY) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XY.Merge(m, src)
}
func (m *XY) XXX_Size() int {
	return xxx_messageInfo_XY.Size(m)
}
func (m *XY) XXX_DiscardUnknown() {
	xxx_messageInfo_XY.DiscardUnknown(m)
}

var xxx_messageInfo_XY proto.InternalMessageInfo

func (m *XY) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *XY) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type ClientToServer struct {
	// Types that are valid to be assigned to Request:
	//	*ClientToServer_Touch
	//	*ClientToServer_GetArea
	Request              isClientToServer_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ClientToServer) Reset()         { *m = ClientToServer{} }
func (m *ClientToServer) String() string { return proto.CompactTextString(m) }
func (*ClientToServer) ProtoMessage()    {}
func (*ClientToServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{5}
}

func (m *ClientToServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientToServer.Unmarshal(m, b)
}
func (m *ClientToServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientToServer.Marshal(b, m, deterministic)
}
func (m *ClientToServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientToServer.Merge(m, src)
}
func (m *ClientToServer) XXX_Size() int {
	return xxx_messageInfo_ClientToServer.Size(m)
}
func (m *ClientToServer) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientToServer.DiscardUnknown(m)
}

var xxx_messageInfo_ClientToServer proto.InternalMessageInfo

type isClientToServer_Request interface {
	isClientToServer_Request()
}

type ClientToServer_Touch struct {
	Touch *TouchRequest `protobuf:"bytes,1,opt,name=touch,proto3,oneof"`
}

type ClientToServer_GetArea struct {
	GetArea *XY `protobuf:"bytes,2,opt,name=getArea,proto3,oneof"`
}

func (*ClientToServer_Touch) isClientToServer_Request() {}

func (*ClientToServer_GetArea) isClientToServer_Request() {}

func (m *ClientToServer) GetRequest() isClientToServer_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ClientToServer) GetTouch() *TouchRequest {
	if x, ok := m.GetRequest().(*ClientToServer_Touch); ok {
		return x.Touch
	}
	return nil
}

func (m *ClientToServer) GetGetArea() *XY {
	if x, ok := m.GetRequest().(*ClientToServer_GetArea); ok {
		return x.GetArea
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClientToServer) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClientToServer_Touch)(nil),
		(*ClientToServer_GetArea)(nil),
	}
}

type ServerToClient struct {
	// Types that are valid to be assigned to Response:
	//	*ServerToClient_Touch
	//	*ServerToClient_Area
	//	*ServerToClient_Msg
	Response             isServerToClient_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ServerToClient) Reset()         { *m = ServerToClient{} }
func (m *ServerToClient) String() string { return proto.CompactTextString(m) }
func (*ServerToClient) ProtoMessage()    {}
func (*ServerToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{6}
}

func (m *ServerToClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerToClient.Unmarshal(m, b)
}
func (m *ServerToClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerToClient.Marshal(b, m, deterministic)
}
func (m *ServerToClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerToClient.Merge(m, src)
}
func (m *ServerToClient) XXX_Size() int {
	return xxx_messageInfo_ServerToClient.Size(m)
}
func (m *ServerToClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerToClient.DiscardUnknown(m)
}

var xxx_messageInfo_ServerToClient proto.InternalMessageInfo

type isServerToClient_Response interface {
	isServerToClient_Response()
}

type ServerToClient_Touch struct {
	Touch *TouchResponse `protobuf:"bytes,1,opt,name=touch,proto3,oneof"`
}

type ServerToClient_Area struct {
	Area *Area `protobuf:"bytes,2,opt,name=area,proto3,oneof"`
}

type ServerToClient_Msg struct {
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3,oneof"`
}

func (*ServerToClient_Touch) isServerToClient_Response() {}

func (*ServerToClient_Area) isServerToClient_Response() {}

func (*ServerToClient_Msg) isServerToClient_Response() {}

func (m *ServerToClient) GetResponse() isServerToClient_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ServerToClient) GetTouch() *TouchResponse {
	if x, ok := m.GetResponse().(*ServerToClient_Touch); ok {
		return x.Touch
	}
	return nil
}

func (m *ServerToClient) GetArea() *Area {
	if x, ok := m.GetResponse().(*ServerToClient_Area); ok {
		return x.Area
	}
	return nil
}

func (m *ServerToClient) GetMsg() string {
	if x, ok := m.GetResponse().(*ServerToClient_Msg); ok {
		return x.Msg
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServerToClient) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServerToClient_Touch)(nil),
		(*ServerToClient_Area)(nil),
		(*ServerToClient_Msg)(nil),
	}
}

func init() {
	proto.RegisterEnum("pb.TouchType", TouchType_name, TouchType_value)
	proto.RegisterEnum("pb.TouchResultType", TouchResultType_name, TouchResultType_value)
	proto.RegisterType((*Cell)(nil), "pb.Cell")
	proto.RegisterType((*Area)(nil), "pb.Area")
	proto.RegisterType((*TouchRequest)(nil), "pb.TouchRequest")
	proto.RegisterType((*TouchResponse)(nil), "pb.TouchResponse")
	proto.RegisterType((*XY)(nil), "pb.XY")
	proto.RegisterType((*ClientToServer)(nil), "pb.ClientToServer")
	proto.RegisterType((*ServerToClient)(nil), "pb.ServerToClient")
}

func init() { proto.RegisterFile("aeilos.proto", fileDescriptor_1a1d8450cb7e47f3) }

var fileDescriptor_1a1d8450cb7e47f3 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xe3, 0x36, 0x5e, 0x9b, 0xd7, 0xae, 0x84, 0x07, 0x42, 0x11, 0x42, 0x25, 0xf2, 0x29,
	0x0c, 0xa9, 0x87, 0x02, 0x7f, 0x40, 0x5b, 0xc1, 0x8a, 0x34, 0x51, 0x64, 0x82, 0xb6, 0x1d, 0x9b,
	0x62, 0x4a, 0x25, 0xaf, 0x0e, 0x49, 0x8a, 0x56, 0xb8, 0xf1, 0x97, 0xa3, 0x67, 0x67, 0x6b, 0xd8,
	0xa1, 0xb7, 0xbc, 0xcf, 0xcf, 0xdf, 0xef, 0xf9, 0xb3, 0x03, 0xfd, 0xa5, 0xda, 0x68, 0x53, 0x8e,
	0xf2, 0xc2, 0x54, 0x06, 0x5b, 0x79, 0x26, 0xfe, 0x32, 0xf0, 0x67, 0x4a, 0x6b, 0xec, 0x03, 0xbb,
	0x8d, 0x58, 0xcc, 0x92, 0xb6, 0x64, 0xb7, 0x54, 0xed, 0xa3, 0x96, 0xab, 0xf6, 0xf8, 0x0c, 0x78,
	0x66, 0x6e, 0xb2, 0x32, 0x6a, 0xc7, 0x2c, 0xe1, 0x73, 0x4f, 0xba, 0x12, 0x9f, 0x43, 0xe7, 0xbb,
	0x5e, 0xae, 0xbf, 0xca, 0x8b, 0xc8, 0x8f, 0x59, 0x12, 0xcc, 0x3d, 0x79, 0x27, 0xe0, 0x10, 0x82,
	0xdd, 0x36, 0x35, 0xbb, 0xd5, 0x0f, 0xf5, 0x2d, 0xe2, 0x31, 0x4b, 0xba, 0x73, 0x4f, 0x1e, 0xa4,
	0x29, 0x40, 0x97, 0xb8, 0xe9, 0x3e, 0x57, 0x62, 0x0a, 0xfe, 0xa4, 0x50, 0xcb, 0xa3, 0x33, 0x0c,
	0x81, 0xaf, 0x94, 0xd6, 0x34, 0x43, 0x3b, 0xe9, 0x8d, 0xbb, 0xa3, 0x3c, 0x1b, 0x91, 0x81, 0x74,
	0xb2, 0xb8, 0x84, 0xbe, 0xb5, 0x96, 0xea, 0xe7, 0x4e, 0x95, 0xd5, 0x51, 0xaf, 0xd7, 0x10, 0x54,
	0xd4, 0x4b, 0x70, 0x7b, 0xa6, 0xc1, 0xf8, 0x94, 0xfc, 0xd2, 0x3b, 0x51, 0x1e, 0xd6, 0xc5, 0x6f,
	0x38, 0xad, 0x8d, 0xcb, 0xdc, 0x6c, 0x4b, 0x85, 0x4f, 0x81, 0x97, 0x2b, 0x53, 0x28, 0xeb, 0xce,
	0xa5, 0x2b, 0xf0, 0x05, 0xf8, 0x34, 0x88, 0x85, 0x34, 0xc7, 0xb3, 0x2a, 0xbe, 0x83, 0x5e, 0x55,
	0x9b, 0xec, 0x74, 0x55, 0x33, 0x9f, 0xdc, 0x33, 0x9d, 0x6c, 0xc9, 0xcd, 0x3e, 0x11, 0x43, 0xeb,
	0xea, 0xfa, 0xd8, 0x51, 0xc4, 0x06, 0x06, 0x33, 0xbd, 0x51, 0xdb, 0x2a, 0x35, 0x5f, 0x54, 0xf1,
	0x4b, 0x15, 0x98, 0x00, 0xb7, 0x16, 0x76, 0x47, 0x6f, 0x1c, 0x36, 0x20, 0x36, 0x19, 0xba, 0x3e,
	0xdb, 0x80, 0x02, 0x3a, 0x6b, 0x55, 0x51, 0xf2, 0xf5, 0xd4, 0x27, 0xd4, 0x7b, 0x75, 0x4d, 0xd7,
	0x58, 0x2f, 0x4c, 0x03, 0xe8, 0x14, 0x6e, 0x9f, 0xf8, 0x03, 0x03, 0x87, 0x48, 0x8d, 0x43, 0xe2,
	0xab, 0xff, 0x51, 0x8f, 0x9b, 0xe7, 0xb1, 0x59, 0x1d, 0x58, 0x43, 0xf0, 0x97, 0x07, 0x90, 0x8d,
	0x87, 0xfc, 0xe7, 0x9e, 0xb4, 0x3a, 0x22, 0xb4, 0x6f, 0xca, 0xb5, 0x0d, 0x86, 0x9e, 0x11, 0x15,
	0xf4, 0x44, 0x8a, 0xda, 0xe8, 0xec, 0x25, 0x04, 0xf7, 0xb7, 0x83, 0x5d, 0xf0, 0x3f, 0x5c, 0x7c,
	0xfc, 0x1c, 0x7a, 0xee, 0x6b, 0x72, 0x1e, 0xb2, 0xb3, 0xb7, 0xf0, 0xe8, 0x41, 0x94, 0xd8, 0x83,
	0xce, 0x6c, 0x21, 0xe5, 0xfb, 0x59, 0x1a, 0x7a, 0x18, 0x00, 0xbf, 0x94, 0x8b, 0x4f, 0xe7, 0x21,
	0xa3, 0xcf, 0xc9, 0x74, 0x21, 0xd3, 0xb0, 0x95, 0x9d, 0xd8, 0x3f, 0xe1, 0xcd, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x31, 0x60, 0x35, 0x85, 0x19, 0x03, 0x00, 0x00,
}