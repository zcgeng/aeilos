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
	User                 string    `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
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

func (m *TouchRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
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

type Zeros struct {
	Coords               []*XY    `protobuf:"bytes,1,rep,name=coords,proto3" json:"coords,omitempty"`
	Score                int32    `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	User                 string   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Zeros) Reset()         { *m = Zeros{} }
func (m *Zeros) String() string { return proto.CompactTextString(m) }
func (*Zeros) ProtoMessage()    {}
func (*Zeros) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{5}
}

func (m *Zeros) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Zeros.Unmarshal(m, b)
}
func (m *Zeros) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Zeros.Marshal(b, m, deterministic)
}
func (m *Zeros) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Zeros.Merge(m, src)
}
func (m *Zeros) XXX_Size() int {
	return xxx_messageInfo_Zeros.Size(m)
}
func (m *Zeros) XXX_DiscardUnknown() {
	xxx_messageInfo_Zeros.DiscardUnknown(m)
}

var xxx_messageInfo_Zeros proto.InternalMessageInfo

func (m *Zeros) GetCoords() []*XY {
	if m != nil {
		return m.Coords
	}
	return nil
}

func (m *Zeros) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Zeros) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
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
	return fileDescriptor_1a1d8450cb7e47f3, []int{6}
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
	//	*ServerToClient_Update
	//	*ServerToClient_UpdateZeros
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
	return fileDescriptor_1a1d8450cb7e47f3, []int{7}
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

type ServerToClient_Update struct {
	Update *Cell `protobuf:"bytes,3,opt,name=update,proto3,oneof"`
}

type ServerToClient_UpdateZeros struct {
	UpdateZeros *Zeros `protobuf:"bytes,4,opt,name=updateZeros,proto3,oneof"`
}

type ServerToClient_Msg struct {
	Msg string `protobuf:"bytes,5,opt,name=msg,proto3,oneof"`
}

func (*ServerToClient_Touch) isServerToClient_Response() {}

func (*ServerToClient_Area) isServerToClient_Response() {}

func (*ServerToClient_Update) isServerToClient_Response() {}

func (*ServerToClient_UpdateZeros) isServerToClient_Response() {}

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

func (m *ServerToClient) GetUpdate() *Cell {
	if x, ok := m.GetResponse().(*ServerToClient_Update); ok {
		return x.Update
	}
	return nil
}

func (m *ServerToClient) GetUpdateZeros() *Zeros {
	if x, ok := m.GetResponse().(*ServerToClient_UpdateZeros); ok {
		return x.UpdateZeros
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
		(*ServerToClient_Update)(nil),
		(*ServerToClient_UpdateZeros)(nil),
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
	proto.RegisterType((*Zeros)(nil), "pb.Zeros")
	proto.RegisterType((*ClientToServer)(nil), "pb.ClientToServer")
	proto.RegisterType((*ServerToClient)(nil), "pb.ServerToClient")
}

func init() { proto.RegisterFile("aeilos.proto", fileDescriptor_1a1d8450cb7e47f3) }

var fileDescriptor_1a1d8450cb7e47f3 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcb, 0x8e, 0xd3, 0x4a,
	0x10, 0x75, 0x27, 0x76, 0x12, 0x97, 0x33, 0xb9, 0xbe, 0x05, 0x42, 0x16, 0x42, 0x21, 0xf2, 0xca,
	0x0c, 0x22, 0x8b, 0x00, 0x1f, 0x90, 0x44, 0x30, 0x46, 0x1a, 0x11, 0x68, 0x8c, 0x98, 0x61, 0x97,
	0x47, 0x13, 0x22, 0x3c, 0x69, 0xe3, 0xb6, 0xd1, 0x84, 0x25, 0x1f, 0xc8, 0x37, 0xa1, 0xae, 0x76,
	0xc6, 0x86, 0x45, 0x76, 0x55, 0xa7, 0xaa, 0xeb, 0x9c, 0x7a, 0x34, 0xf4, 0x97, 0x62, 0x97, 0x4a,
	0x35, 0xce, 0x72, 0x59, 0x48, 0x6c, 0x65, 0xab, 0xf0, 0x17, 0x03, 0x7b, 0x2e, 0xd2, 0x14, 0xfb,
	0xc0, 0x6e, 0x03, 0x36, 0x62, 0x51, 0x9b, 0xb3, 0x5b, 0xed, 0x1d, 0x82, 0x96, 0xf1, 0x0e, 0xf8,
	0x00, 0x9c, 0x95, 0xbc, 0x59, 0xa9, 0xa0, 0x3d, 0x62, 0x91, 0x13, 0x5b, 0xdc, 0xb8, 0xf8, 0x10,
	0xba, 0x5f, 0xd2, 0xe5, 0xf6, 0x23, 0xbf, 0x0c, 0xec, 0x11, 0x8b, 0xdc, 0xd8, 0xe2, 0x47, 0x00,
	0x87, 0xe0, 0x96, 0xfb, 0x44, 0x96, 0xeb, 0xaf, 0x62, 0x13, 0x38, 0x23, 0x16, 0xf5, 0x62, 0x8b,
	0xd7, 0xd0, 0x0c, 0xa0, 0xa7, 0x79, 0x93, 0x43, 0x26, 0xc2, 0x19, 0xd8, 0xd3, 0x5c, 0x2c, 0x4f,
	0x6a, 0x18, 0x82, 0xb3, 0x16, 0x69, 0xaa, 0x35, 0xb4, 0x23, 0x6f, 0xd2, 0x1b, 0x67, 0xab, 0xb1,
	0x2e, 0xc0, 0x0d, 0x1c, 0x7e, 0x83, 0x3e, 0x95, 0xe6, 0xe2, 0x7b, 0x29, 0x54, 0x71, 0xb2, 0xd6,
	0x53, 0x70, 0x0b, 0x9d, 0xab, 0xc9, 0xa9, 0xa7, 0xc1, 0xe4, 0x4c, 0xd7, 0x4b, 0x8e, 0x20, 0xaf,
	0xe3, 0x88, 0x60, 0x97, 0x4a, 0xe4, 0xa6, 0x43, 0x4e, 0x76, 0xf8, 0x13, 0xce, 0x2a, 0x32, 0x95,
	0xc9, 0xbd, 0x12, 0x78, 0x1f, 0x1c, 0xb5, 0x96, 0xb9, 0x20, 0x46, 0x87, 0x1b, 0x07, 0x1f, 0x81,
	0xad, 0xc5, 0x11, 0x71, 0x53, 0x32, 0xa1, 0xf8, 0x12, 0xbc, 0xa2, 0x2a, 0x52, 0xa6, 0x45, 0xa5,
	0xe3, 0xde, 0x9d, 0x0e, 0x03, 0x93, 0x9a, 0x66, 0x5e, 0x38, 0x82, 0xd6, 0xd5, 0xf5, 0xa9, 0xf6,
	0xc2, 0xf7, 0xe0, 0x7c, 0x16, 0xb9, 0x54, 0x38, 0x84, 0xce, 0x5a, 0xca, 0x7c, 0xa3, 0x02, 0x46,
	0x43, 0xeb, 0xe8, 0xe2, 0x57, 0xd7, 0xbc, 0x42, 0x6b, 0xd5, 0xad, 0xa6, 0xea, 0x63, 0xc3, 0xed,
	0x46, 0xc3, 0x3b, 0x18, 0xcc, 0xd3, 0x9d, 0xd8, 0x17, 0x89, 0xfc, 0x20, 0xf2, 0x1f, 0x22, 0xc7,
	0x08, 0x1c, 0x52, 0x45, 0x22, 0xbc, 0x89, 0xdf, 0xd0, 0x4d, 0x0b, 0xd0, 0x57, 0x42, 0x09, 0x18,
	0x42, 0x77, 0x2b, 0x0a, 0xbd, 0xe0, 0x6a, 0x10, 0x95, 0x0c, 0x7d, 0x2d, 0x55, 0x60, 0xe6, 0x42,
	0x37, 0x37, 0xef, 0xc2, 0xdf, 0x0c, 0x06, 0x86, 0x23, 0x91, 0x86, 0x13, 0x9f, 0xfc, 0xcd, 0xf5,
	0x7f, 0x73, 0x46, 0x34, 0xff, 0x9a, 0x6c, 0x08, 0xf6, 0xb2, 0x66, 0xa2, 0x91, 0x6b, 0x82, 0xd8,
	0xe2, 0x84, 0x63, 0x08, 0x9d, 0x32, 0xdb, 0x2c, 0x0b, 0xb3, 0xf7, 0xc6, 0x52, 0x62, 0x8b, 0x57,
	0x11, 0x7c, 0x06, 0x9e, 0xb1, 0x68, 0x8a, 0xb4, 0x78, 0x6f, 0xe2, 0xea, 0x44, 0x02, 0x62, 0x8b,
	0x37, 0xe3, 0x88, 0xd0, 0xbe, 0x51, 0x5b, 0xba, 0x71, 0xfd, 0x03, 0xb4, 0xa3, 0xaf, 0x3b, 0xaf,
	0xb4, 0x9d, 0x3f, 0x06, 0xf7, 0xee, 0xb0, 0xb0, 0x07, 0xf6, 0xeb, 0xcb, 0x37, 0xef, 0x7c, 0xcb,
	0x58, 0xd3, 0x0b, 0x9f, 0x9d, 0xbf, 0x80, 0xff, 0xfe, 0xd9, 0x38, 0x7a, 0xd0, 0x9d, 0x2f, 0x38,
	0x7f, 0x35, 0x4f, 0x7c, 0x0b, 0x5d, 0x70, 0x3e, 0xf1, 0xc5, 0xdb, 0x0b, 0x9f, 0x69, 0x73, 0x3a,
	0x5b, 0xf0, 0xc4, 0x6f, 0xad, 0x3a, 0xf4, 0x89, 0x9f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x66,
	0x8e, 0x6f, 0x5f, 0xd4, 0x03, 0x00, 0x00,
}
