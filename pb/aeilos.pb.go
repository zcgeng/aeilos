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

type GetStats struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStats) Reset()         { *m = GetStats{} }
func (m *GetStats) String() string { return proto.CompactTextString(m) }
func (*GetStats) ProtoMessage()    {}
func (*GetStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{6}
}

func (m *GetStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStats.Unmarshal(m, b)
}
func (m *GetStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStats.Marshal(b, m, deterministic)
}
func (m *GetStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStats.Merge(m, src)
}
func (m *GetStats) XXX_Size() int {
	return xxx_messageInfo_GetStats.Size(m)
}
func (m *GetStats) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStats.DiscardUnknown(m)
}

var xxx_messageInfo_GetStats proto.InternalMessageInfo

func (m *GetStats) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type Stats struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Score                int64    `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{7}
}

func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Stats) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

type ChatMsg struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMsg) Reset()         { *m = ChatMsg{} }
func (m *ChatMsg) String() string { return proto.CompactTextString(m) }
func (*ChatMsg) ProtoMessage()    {}
func (*ChatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{8}
}

func (m *ChatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMsg.Unmarshal(m, b)
}
func (m *ChatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMsg.Marshal(b, m, deterministic)
}
func (m *ChatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMsg.Merge(m, src)
}
func (m *ChatMsg) XXX_Size() int {
	return xxx_messageInfo_ChatMsg.Size(m)
}
func (m *ChatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMsg proto.InternalMessageInfo

func (m *ChatMsg) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ChatMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ChatMsg) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type ClientToServer struct {
	// Types that are valid to be assigned to Request:
	//	*ClientToServer_Touch
	//	*ClientToServer_GetArea
	//	*ClientToServer_GetStats
	//	*ClientToServer_ChatMsg
	Request              isClientToServer_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ClientToServer) Reset()         { *m = ClientToServer{} }
func (m *ClientToServer) String() string { return proto.CompactTextString(m) }
func (*ClientToServer) ProtoMessage()    {}
func (*ClientToServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{9}
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

type ClientToServer_GetStats struct {
	GetStats *GetStats `protobuf:"bytes,3,opt,name=getStats,proto3,oneof"`
}

type ClientToServer_ChatMsg struct {
	ChatMsg *ChatMsg `protobuf:"bytes,4,opt,name=chatMsg,proto3,oneof"`
}

func (*ClientToServer_Touch) isClientToServer_Request() {}

func (*ClientToServer_GetArea) isClientToServer_Request() {}

func (*ClientToServer_GetStats) isClientToServer_Request() {}

func (*ClientToServer_ChatMsg) isClientToServer_Request() {}

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

func (m *ClientToServer) GetGetStats() *GetStats {
	if x, ok := m.GetRequest().(*ClientToServer_GetStats); ok {
		return x.GetStats
	}
	return nil
}

func (m *ClientToServer) GetChatMsg() *ChatMsg {
	if x, ok := m.GetRequest().(*ClientToServer_ChatMsg); ok {
		return x.ChatMsg
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClientToServer) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClientToServer_Touch)(nil),
		(*ClientToServer_GetArea)(nil),
		(*ClientToServer_GetStats)(nil),
		(*ClientToServer_ChatMsg)(nil),
	}
}

type ServerToClient struct {
	// Types that are valid to be assigned to Response:
	//	*ServerToClient_Touch
	//	*ServerToClient_Area
	//	*ServerToClient_Update
	//	*ServerToClient_UpdateZeros
	//	*ServerToClient_Msg
	//	*ServerToClient_Stats
	Response             isServerToClient_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ServerToClient) Reset()         { *m = ServerToClient{} }
func (m *ServerToClient) String() string { return proto.CompactTextString(m) }
func (*ServerToClient) ProtoMessage()    {}
func (*ServerToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1d8450cb7e47f3, []int{10}
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
	Msg *ChatMsg `protobuf:"bytes,5,opt,name=msg,proto3,oneof"`
}

type ServerToClient_Stats struct {
	Stats *Stats `protobuf:"bytes,6,opt,name=stats,proto3,oneof"`
}

func (*ServerToClient_Touch) isServerToClient_Response() {}

func (*ServerToClient_Area) isServerToClient_Response() {}

func (*ServerToClient_Update) isServerToClient_Response() {}

func (*ServerToClient_UpdateZeros) isServerToClient_Response() {}

func (*ServerToClient_Msg) isServerToClient_Response() {}

func (*ServerToClient_Stats) isServerToClient_Response() {}

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

func (m *ServerToClient) GetMsg() *ChatMsg {
	if x, ok := m.GetResponse().(*ServerToClient_Msg); ok {
		return x.Msg
	}
	return nil
}

func (m *ServerToClient) GetStats() *Stats {
	if x, ok := m.GetResponse().(*ServerToClient_Stats); ok {
		return x.Stats
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServerToClient) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServerToClient_Touch)(nil),
		(*ServerToClient_Area)(nil),
		(*ServerToClient_Update)(nil),
		(*ServerToClient_UpdateZeros)(nil),
		(*ServerToClient_Msg)(nil),
		(*ServerToClient_Stats)(nil),
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
	proto.RegisterType((*GetStats)(nil), "pb.GetStats")
	proto.RegisterType((*Stats)(nil), "pb.Stats")
	proto.RegisterType((*ChatMsg)(nil), "pb.ChatMsg")
	proto.RegisterType((*ClientToServer)(nil), "pb.ClientToServer")
	proto.RegisterType((*ServerToClient)(nil), "pb.ServerToClient")
}

func init() { proto.RegisterFile("aeilos.proto", fileDescriptor_1a1d8450cb7e47f3) }

var fileDescriptor_1a1d8450cb7e47f3 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xf5, 0xc6, 0x76, 0x62, 0x8f, 0xd3, 0x7e, 0xfe, 0x16, 0x84, 0x22, 0x84, 0xd2, 0xe0, 0x03,
	0x84, 0x22, 0x7a, 0x08, 0x70, 0xe0, 0xd8, 0x44, 0x50, 0x23, 0x4a, 0x0b, 0x5b, 0x23, 0x5a, 0x6e,
	0x4e, 0xba, 0xa4, 0x15, 0x4e, 0x37, 0x78, 0x37, 0xa8, 0xe1, 0xc8, 0x0f, 0xe3, 0xa7, 0x21, 0xb4,
	0xb3, 0x9b, 0xd8, 0x45, 0x28, 0xdc, 0x66, 0xdf, 0x8c, 0x67, 0xde, 0x9b, 0x7d, 0x5e, 0x68, 0xe7,
	0xfc, 0xb2, 0x10, 0x72, 0x6f, 0x5e, 0x0a, 0x25, 0x68, 0x63, 0x3e, 0x4e, 0x7e, 0x10, 0xf0, 0x46,
	0xbc, 0x28, 0x68, 0x1b, 0xc8, 0x75, 0x87, 0xf4, 0x48, 0xdf, 0x65, 0xe4, 0x5a, 0x9f, 0x96, 0x9d,
	0x86, 0x39, 0x2d, 0xe9, 0x1d, 0xf0, 0xc7, 0x62, 0x36, 0x96, 0x1d, 0xb7, 0x47, 0xfa, 0x7e, 0xea,
	0x30, 0x73, 0xa4, 0x77, 0xa1, 0xf5, 0xb9, 0xc8, 0xa7, 0x1f, 0xd8, 0x61, 0xc7, 0xeb, 0x91, 0x7e,
	0x98, 0x3a, 0x6c, 0x05, 0xd0, 0x2e, 0x84, 0x8b, 0xab, 0x4c, 0x2c, 0x26, 0x17, 0xfc, 0xbc, 0xe3,
	0xf7, 0x48, 0x3f, 0x48, 0x1d, 0x56, 0x41, 0x43, 0x80, 0x40, 0xcf, 0xcd, 0x96, 0x73, 0x9e, 0x0c,
	0xc1, 0xdb, 0x2f, 0x79, 0xbe, 0x91, 0x43, 0x17, 0xfc, 0x09, 0x2f, 0x0a, 0xcd, 0xc1, 0xed, 0x47,
	0x83, 0x60, 0x6f, 0x3e, 0xde, 0xd3, 0x0d, 0x98, 0x81, 0x93, 0x2f, 0xd0, 0xc6, 0xd6, 0x8c, 0x7f,
	0x5d, 0x70, 0xa9, 0x36, 0xf6, 0x7a, 0x0c, 0xa1, 0xd2, 0xb5, 0x7a, 0x38, 0x6a, 0xda, 0x1e, 0x6c,
	0xe9, 0x7e, 0xd9, 0x0a, 0x64, 0x55, 0x9e, 0x52, 0xf0, 0x16, 0x92, 0x97, 0x46, 0x21, 0xc3, 0x38,
	0xf9, 0x0e, 0x5b, 0x76, 0x98, 0x9c, 0x8b, 0x2b, 0xc9, 0xe9, 0x6d, 0xf0, 0xe5, 0x44, 0x94, 0x1c,
	0x27, 0xfa, 0xcc, 0x1c, 0xe8, 0x3d, 0xf0, 0x34, 0x39, 0x1c, 0x5c, 0xa7, 0x8c, 0x28, 0x7d, 0x0e,
	0x91, 0xb2, 0x4d, 0x16, 0x85, 0xb2, 0x3c, 0x6e, 0xad, 0x79, 0x18, 0x18, 0xd9, 0xd4, 0xeb, 0x92,
	0x1e, 0x34, 0x4e, 0xcf, 0x36, 0xc9, 0x4b, 0xde, 0x83, 0xff, 0x89, 0x97, 0x42, 0xd2, 0x2e, 0x34,
	0x27, 0x42, 0x94, 0xe7, 0xb2, 0x43, 0x70, 0x69, 0x4d, 0xdd, 0xfc, 0xf4, 0x8c, 0x59, 0xb4, 0x62,
	0xdd, 0xa8, 0xb3, 0x5e, 0x09, 0x76, 0x6b, 0x82, 0x1f, 0x40, 0x70, 0xc0, 0xd5, 0x89, 0xca, 0x95,
	0xbe, 0xf5, 0x40, 0x63, 0x47, 0xf9, 0xcc, 0xc8, 0x0d, 0xd9, 0xfa, 0x9c, 0xbc, 0x00, 0xff, 0x9f,
	0x45, 0x37, 0xc7, 0xba, 0x76, 0x6c, 0xf2, 0x06, 0x5a, 0xa3, 0x8b, 0x5c, 0xbd, 0x95, 0xd3, 0x8d,
	0x1f, 0xc7, 0xe0, 0xce, 0xe4, 0x14, 0x3f, 0x0d, 0x99, 0x0e, 0x35, 0x5f, 0x75, 0x39, 0x33, 0x17,
	0xe9, 0x32, 0x8c, 0x93, 0x9f, 0x04, 0xb6, 0x47, 0xc5, 0x25, 0xbf, 0x52, 0x99, 0x38, 0xe1, 0xe5,
	0x37, 0x5e, 0xd2, 0x3e, 0xf8, 0xb8, 0x46, 0xec, 0x18, 0x0d, 0xe2, 0xda, 0xa2, 0xd1, 0x31, 0xda,
	0xd6, 0x58, 0x40, 0x13, 0x68, 0x4d, 0xb9, 0xd2, 0x8e, 0xb4, 0x37, 0x67, 0xf7, 0xa6, 0xed, 0x6d,
	0x13, 0x74, 0x17, 0x82, 0xa9, 0x5d, 0x08, 0x0e, 0x8e, 0x06, 0x6d, 0x5d, 0xb4, 0x5a, 0x52, 0xea,
	0xb0, 0x75, 0x9e, 0x3e, 0x84, 0xd6, 0xc4, 0x28, 0x43, 0x13, 0x45, 0x83, 0x08, 0x9d, 0x60, 0x20,
	0xdd, 0xd4, 0x66, 0x87, 0x21, 0xb4, 0x4a, 0x43, 0x26, 0xf9, 0x45, 0x60, 0xdb, 0x10, 0xcf, 0x84,
	0x11, 0x42, 0x1f, 0xdd, 0x14, 0xf0, 0x7f, 0xdd, 0x29, 0xe8, 0xc2, 0x4a, 0x41, 0x17, 0xbc, 0xbc,
	0xa2, 0x8f, 0xc6, 0xd3, 0xac, 0x53, 0x87, 0x21, 0x4e, 0x13, 0x68, 0x2e, 0xe6, 0xe7, 0xb9, 0xe2,
	0x96, 0xfb, 0xda, 0x9a, 0xa9, 0xc3, 0x6c, 0x86, 0x3e, 0x81, 0xc8, 0x44, 0xe8, 0x25, 0xcb, 0x3c,
	0xd4, 0x85, 0x08, 0xa4, 0x0e, 0xab, 0xe7, 0xe9, 0x8e, 0xb9, 0x17, 0xff, 0x6f, 0x02, 0xf1, 0x9a,
	0xee, 0x83, 0x2f, 0x71, 0x5d, 0xcd, 0xaa, 0xd3, 0x6a, 0x57, 0x26, 0xa3, 0xdf, 0x84, 0xd2, 0x6a,
	0xd9, 0xdd, 0x81, 0x70, 0xfd, 0x3b, 0xd2, 0x00, 0xbc, 0x57, 0x87, 0xaf, 0xdf, 0xc5, 0x8e, 0x89,
	0xf6, 0x0f, 0x62, 0xb2, 0xfb, 0x0c, 0xfe, 0xfb, 0xe3, 0x3f, 0xa1, 0x11, 0xb4, 0x46, 0xc7, 0x8c,
	0xbd, 0x1c, 0x65, 0xb1, 0x43, 0x43, 0xf0, 0x3f, 0xb2, 0xe3, 0xa3, 0x83, 0x98, 0xe8, 0x70, 0x7f,
	0x78, 0xcc, 0xb2, 0xb8, 0x31, 0x6e, 0xe2, 0xd3, 0xf7, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x78, 0xd9, 0x95, 0xa2, 0x0a, 0x05, 0x00, 0x00,
}
