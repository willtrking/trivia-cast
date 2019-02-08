// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stm/shoot_the_messenger.proto

package stm

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type STATE int32

const (
	STATE_UNKNOWN    STATE = 0
	STATE_RESPONDING STATE = 1
	STATE_READING    STATE = 2
	STATE_GUESSING   STATE = 3
	STATE_RESULTS    STATE = 4
)

var STATE_name = map[int32]string{
	0: "UNKNOWN",
	1: "RESPONDING",
	2: "READING",
	3: "GUESSING",
	4: "RESULTS",
}

var STATE_value = map[string]int32{
	"UNKNOWN":    0,
	"RESPONDING": 1,
	"READING":    2,
	"GUESSING":   3,
	"RESULTS":    4,
}

func (x STATE) String() string {
	return proto.EnumName(STATE_name, int32(x))
}

func (STATE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{0}
}

type Prompt struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Contents             string   `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prompt) Reset()         { *m = Prompt{} }
func (m *Prompt) String() string { return proto.CompactTextString(m) }
func (*Prompt) ProtoMessage()    {}
func (*Prompt) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{0}
}

func (m *Prompt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prompt.Unmarshal(m, b)
}
func (m *Prompt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prompt.Marshal(b, m, deterministic)
}
func (m *Prompt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prompt.Merge(m, src)
}
func (m *Prompt) XXX_Size() int {
	return xxx_messageInfo_Prompt.Size(m)
}
func (m *Prompt) XXX_DiscardUnknown() {
	xxx_messageInfo_Prompt.DiscardUnknown(m)
}

var xxx_messageInfo_Prompt proto.InternalMessageInfo

func (m *Prompt) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Prompt) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

type Response struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Response             string   `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Response) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type RoleFrame struct {
	Guesser              bool     `protobuf:"varint,1,opt,name=guesser,proto3" json:"guesser,omitempty"`
	Out                  bool     `protobuf:"varint,2,opt,name=out,proto3" json:"out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleFrame) Reset()         { *m = RoleFrame{} }
func (m *RoleFrame) String() string { return proto.CompactTextString(m) }
func (*RoleFrame) ProtoMessage()    {}
func (*RoleFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{2}
}

func (m *RoleFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleFrame.Unmarshal(m, b)
}
func (m *RoleFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleFrame.Marshal(b, m, deterministic)
}
func (m *RoleFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleFrame.Merge(m, src)
}
func (m *RoleFrame) XXX_Size() int {
	return xxx_messageInfo_RoleFrame.Size(m)
}
func (m *RoleFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleFrame.DiscardUnknown(m)
}

var xxx_messageInfo_RoleFrame proto.InternalMessageInfo

func (m *RoleFrame) GetGuesser() bool {
	if m != nil {
		return m.Guesser
	}
	return false
}

func (m *RoleFrame) GetOut() bool {
	if m != nil {
		return m.Out
	}
	return false
}

type GuessFrame struct {
	PromptId             string   `protobuf:"bytes,1,opt,name=prompt_id,json=promptId,proto3" json:"prompt_id,omitempty"`
	ResponseId           string   `protobuf:"bytes,2,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GuessFrame) Reset()         { *m = GuessFrame{} }
func (m *GuessFrame) String() string { return proto.CompactTextString(m) }
func (*GuessFrame) ProtoMessage()    {}
func (*GuessFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{3}
}

func (m *GuessFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuessFrame.Unmarshal(m, b)
}
func (m *GuessFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuessFrame.Marshal(b, m, deterministic)
}
func (m *GuessFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuessFrame.Merge(m, src)
}
func (m *GuessFrame) XXX_Size() int {
	return xxx_messageInfo_GuessFrame.Size(m)
}
func (m *GuessFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_GuessFrame.DiscardUnknown(m)
}

var xxx_messageInfo_GuessFrame proto.InternalMessageInfo

func (m *GuessFrame) GetPromptId() string {
	if m != nil {
		return m.PromptId
	}
	return ""
}

func (m *GuessFrame) GetResponseId() string {
	if m != nil {
		return m.ResponseId
	}
	return ""
}

type PromptResponseFrame struct {
	PromptId             string   `protobuf:"bytes,1,opt,name=prompt_id,json=promptId,proto3" json:"prompt_id,omitempty"`
	Response             string   `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromptResponseFrame) Reset()         { *m = PromptResponseFrame{} }
func (m *PromptResponseFrame) String() string { return proto.CompactTextString(m) }
func (*PromptResponseFrame) ProtoMessage()    {}
func (*PromptResponseFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{4}
}

func (m *PromptResponseFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromptResponseFrame.Unmarshal(m, b)
}
func (m *PromptResponseFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromptResponseFrame.Marshal(b, m, deterministic)
}
func (m *PromptResponseFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromptResponseFrame.Merge(m, src)
}
func (m *PromptResponseFrame) XXX_Size() int {
	return xxx_messageInfo_PromptResponseFrame.Size(m)
}
func (m *PromptResponseFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_PromptResponseFrame.DiscardUnknown(m)
}

var xxx_messageInfo_PromptResponseFrame proto.InternalMessageInfo

func (m *PromptResponseFrame) GetPromptId() string {
	if m != nil {
		return m.PromptId
	}
	return ""
}

func (m *PromptResponseFrame) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type PromptResponseInputFrame struct {
	PromptId             string   `protobuf:"bytes,1,opt,name=prompt_id,json=promptId,proto3" json:"prompt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromptResponseInputFrame) Reset()         { *m = PromptResponseInputFrame{} }
func (m *PromptResponseInputFrame) String() string { return proto.CompactTextString(m) }
func (*PromptResponseInputFrame) ProtoMessage()    {}
func (*PromptResponseInputFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{5}
}

func (m *PromptResponseInputFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromptResponseInputFrame.Unmarshal(m, b)
}
func (m *PromptResponseInputFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromptResponseInputFrame.Marshal(b, m, deterministic)
}
func (m *PromptResponseInputFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromptResponseInputFrame.Merge(m, src)
}
func (m *PromptResponseInputFrame) XXX_Size() int {
	return xxx_messageInfo_PromptResponseInputFrame.Size(m)
}
func (m *PromptResponseInputFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_PromptResponseInputFrame.DiscardUnknown(m)
}

var xxx_messageInfo_PromptResponseInputFrame proto.InternalMessageInfo

func (m *PromptResponseInputFrame) GetPromptId() string {
	if m != nil {
		return m.PromptId
	}
	return ""
}

type AllResponsesFrame struct {
	Responses            []*Response `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AllResponsesFrame) Reset()         { *m = AllResponsesFrame{} }
func (m *AllResponsesFrame) String() string { return proto.CompactTextString(m) }
func (*AllResponsesFrame) ProtoMessage()    {}
func (*AllResponsesFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{6}
}

func (m *AllResponsesFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllResponsesFrame.Unmarshal(m, b)
}
func (m *AllResponsesFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllResponsesFrame.Marshal(b, m, deterministic)
}
func (m *AllResponsesFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllResponsesFrame.Merge(m, src)
}
func (m *AllResponsesFrame) XXX_Size() int {
	return xxx_messageInfo_AllResponsesFrame.Size(m)
}
func (m *AllResponsesFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_AllResponsesFrame.DiscardUnknown(m)
}

var xxx_messageInfo_AllResponsesFrame proto.InternalMessageInfo

func (m *AllResponsesFrame) GetResponses() []*Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

type ReadResponseFrame struct {
	ResponseId           string   `protobuf:"bytes,1,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponseFrame) Reset()         { *m = ReadResponseFrame{} }
func (m *ReadResponseFrame) String() string { return proto.CompactTextString(m) }
func (*ReadResponseFrame) ProtoMessage()    {}
func (*ReadResponseFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{7}
}

func (m *ReadResponseFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponseFrame.Unmarshal(m, b)
}
func (m *ReadResponseFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponseFrame.Marshal(b, m, deterministic)
}
func (m *ReadResponseFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponseFrame.Merge(m, src)
}
func (m *ReadResponseFrame) XXX_Size() int {
	return xxx_messageInfo_ReadResponseFrame.Size(m)
}
func (m *ReadResponseFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponseFrame.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponseFrame proto.InternalMessageInfo

func (m *ReadResponseFrame) GetResponseId() string {
	if m != nil {
		return m.ResponseId
	}
	return ""
}

type PlayerFrame struct {
	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	// Types that are valid to be assigned to Frame:
	//	*PlayerFrame_State
	//	*PlayerFrame_Guess
	//	*PlayerFrame_PromptResponseInput
	//	*PlayerFrame_PromptResponse
	//	*PlayerFrame_RoleChange
	//	*PlayerFrame_AllResponses
	//	*PlayerFrame_ReadResponse
	Frame                isPlayerFrame_Frame `protobuf_oneof:"frame"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PlayerFrame) Reset()         { *m = PlayerFrame{} }
func (m *PlayerFrame) String() string { return proto.CompactTextString(m) }
func (*PlayerFrame) ProtoMessage()    {}
func (*PlayerFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{8}
}

func (m *PlayerFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerFrame.Unmarshal(m, b)
}
func (m *PlayerFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerFrame.Marshal(b, m, deterministic)
}
func (m *PlayerFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerFrame.Merge(m, src)
}
func (m *PlayerFrame) XXX_Size() int {
	return xxx_messageInfo_PlayerFrame.Size(m)
}
func (m *PlayerFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerFrame.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerFrame proto.InternalMessageInfo

func (m *PlayerFrame) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type isPlayerFrame_Frame interface {
	isPlayerFrame_Frame()
}

type PlayerFrame_State struct {
	State *GameState `protobuf:"bytes,2,opt,name=state,proto3,oneof"`
}

type PlayerFrame_Guess struct {
	Guess *GuessFrame `protobuf:"bytes,3,opt,name=guess,proto3,oneof"`
}

type PlayerFrame_PromptResponseInput struct {
	PromptResponseInput *PromptResponseInputFrame `protobuf:"bytes,4,opt,name=prompt_response_input,json=promptResponseInput,proto3,oneof"`
}

type PlayerFrame_PromptResponse struct {
	PromptResponse *PromptResponseFrame `protobuf:"bytes,5,opt,name=prompt_response,json=promptResponse,proto3,oneof"`
}

type PlayerFrame_RoleChange struct {
	RoleChange *RoleFrame `protobuf:"bytes,6,opt,name=role_change,json=roleChange,proto3,oneof"`
}

type PlayerFrame_AllResponses struct {
	AllResponses *AllResponsesFrame `protobuf:"bytes,7,opt,name=all_responses,json=allResponses,proto3,oneof"`
}

type PlayerFrame_ReadResponse struct {
	ReadResponse *ReadResponseFrame `protobuf:"bytes,8,opt,name=read_response,json=readResponse,proto3,oneof"`
}

func (*PlayerFrame_State) isPlayerFrame_Frame() {}

func (*PlayerFrame_Guess) isPlayerFrame_Frame() {}

func (*PlayerFrame_PromptResponseInput) isPlayerFrame_Frame() {}

func (*PlayerFrame_PromptResponse) isPlayerFrame_Frame() {}

func (*PlayerFrame_RoleChange) isPlayerFrame_Frame() {}

func (*PlayerFrame_AllResponses) isPlayerFrame_Frame() {}

func (*PlayerFrame_ReadResponse) isPlayerFrame_Frame() {}

func (m *PlayerFrame) GetFrame() isPlayerFrame_Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *PlayerFrame) GetState() *GameState {
	if x, ok := m.GetFrame().(*PlayerFrame_State); ok {
		return x.State
	}
	return nil
}

func (m *PlayerFrame) GetGuess() *GuessFrame {
	if x, ok := m.GetFrame().(*PlayerFrame_Guess); ok {
		return x.Guess
	}
	return nil
}

func (m *PlayerFrame) GetPromptResponseInput() *PromptResponseInputFrame {
	if x, ok := m.GetFrame().(*PlayerFrame_PromptResponseInput); ok {
		return x.PromptResponseInput
	}
	return nil
}

func (m *PlayerFrame) GetPromptResponse() *PromptResponseFrame {
	if x, ok := m.GetFrame().(*PlayerFrame_PromptResponse); ok {
		return x.PromptResponse
	}
	return nil
}

func (m *PlayerFrame) GetRoleChange() *RoleFrame {
	if x, ok := m.GetFrame().(*PlayerFrame_RoleChange); ok {
		return x.RoleChange
	}
	return nil
}

func (m *PlayerFrame) GetAllResponses() *AllResponsesFrame {
	if x, ok := m.GetFrame().(*PlayerFrame_AllResponses); ok {
		return x.AllResponses
	}
	return nil
}

func (m *PlayerFrame) GetReadResponse() *ReadResponseFrame {
	if x, ok := m.GetFrame().(*PlayerFrame_ReadResponse); ok {
		return x.ReadResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PlayerFrame) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PlayerFrame_State)(nil),
		(*PlayerFrame_Guess)(nil),
		(*PlayerFrame_PromptResponseInput)(nil),
		(*PlayerFrame_PromptResponse)(nil),
		(*PlayerFrame_RoleChange)(nil),
		(*PlayerFrame_AllResponses)(nil),
		(*PlayerFrame_ReadResponse)(nil),
	}
}

type BeginGameReq struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeginGameReq) Reset()         { *m = BeginGameReq{} }
func (m *BeginGameReq) String() string { return proto.CompactTextString(m) }
func (*BeginGameReq) ProtoMessage()    {}
func (*BeginGameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{9}
}

func (m *BeginGameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeginGameReq.Unmarshal(m, b)
}
func (m *BeginGameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeginGameReq.Marshal(b, m, deterministic)
}
func (m *BeginGameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeginGameReq.Merge(m, src)
}
func (m *BeginGameReq) XXX_Size() int {
	return xxx_messageInfo_BeginGameReq.Size(m)
}
func (m *BeginGameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BeginGameReq.DiscardUnknown(m)
}

var xxx_messageInfo_BeginGameReq proto.InternalMessageInfo

func (m *BeginGameReq) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type GameState struct {
	GameState            STATE     `protobuf:"varint,1,opt,name=game_state,json=gameState,proto3,enum=stm.STATE" json:"game_state,omitempty"`
	Prompts              []*Prompt `protobuf:"bytes,2,rep,name=prompts,proto3" json:"prompts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GameState) Reset()         { *m = GameState{} }
func (m *GameState) String() string { return proto.CompactTextString(m) }
func (*GameState) ProtoMessage()    {}
func (*GameState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4170d256edd2c8d, []int{10}
}

func (m *GameState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameState.Unmarshal(m, b)
}
func (m *GameState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameState.Marshal(b, m, deterministic)
}
func (m *GameState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameState.Merge(m, src)
}
func (m *GameState) XXX_Size() int {
	return xxx_messageInfo_GameState.Size(m)
}
func (m *GameState) XXX_DiscardUnknown() {
	xxx_messageInfo_GameState.DiscardUnknown(m)
}

var xxx_messageInfo_GameState proto.InternalMessageInfo

func (m *GameState) GetGameState() STATE {
	if m != nil {
		return m.GameState
	}
	return STATE_UNKNOWN
}

func (m *GameState) GetPrompts() []*Prompt {
	if m != nil {
		return m.Prompts
	}
	return nil
}

func init() {
	proto.RegisterEnum("stm.STATE", STATE_name, STATE_value)
	proto.RegisterType((*Prompt)(nil), "stm.Prompt")
	proto.RegisterType((*Response)(nil), "stm.Response")
	proto.RegisterType((*RoleFrame)(nil), "stm.RoleFrame")
	proto.RegisterType((*GuessFrame)(nil), "stm.GuessFrame")
	proto.RegisterType((*PromptResponseFrame)(nil), "stm.PromptResponseFrame")
	proto.RegisterType((*PromptResponseInputFrame)(nil), "stm.PromptResponseInputFrame")
	proto.RegisterType((*AllResponsesFrame)(nil), "stm.AllResponsesFrame")
	proto.RegisterType((*ReadResponseFrame)(nil), "stm.ReadResponseFrame")
	proto.RegisterType((*PlayerFrame)(nil), "stm.PlayerFrame")
	proto.RegisterType((*BeginGameReq)(nil), "stm.BeginGameReq")
	proto.RegisterType((*GameState)(nil), "stm.GameState")
}

func init() { proto.RegisterFile("stm/shoot_the_messenger.proto", fileDescriptor_a4170d256edd2c8d) }

var fileDescriptor_a4170d256edd2c8d = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x8d, 0x93, 0x26, 0xb1, 0xc7, 0x69, 0x9a, 0x6c, 0xf5, 0x7d, 0x58, 0x45, 0x15, 0x95, 0x25,
	0x68, 0x01, 0x29, 0x85, 0x50, 0x95, 0x2b, 0x24, 0xda, 0x12, 0x92, 0x40, 0x49, 0xab, 0x75, 0x2a,
	0x24, 0x24, 0x14, 0xb9, 0xf5, 0xd6, 0xa9, 0x64, 0x7b, 0x8d, 0x77, 0x73, 0x81, 0x78, 0x48, 0x5e,
	0x09, 0xed, 0x6e, 0xd7, 0x76, 0x93, 0xf2, 0x73, 0xe7, 0x99, 0x33, 0x67, 0x7c, 0x66, 0xce, 0x68,
	0x61, 0x9b, 0xf1, 0x78, 0x9f, 0xcd, 0x29, 0xe5, 0x33, 0x3e, 0x27, 0xb3, 0x98, 0x30, 0x46, 0x92,
	0x90, 0x64, 0xbd, 0x34, 0xa3, 0x9c, 0xa2, 0x1a, 0xe3, 0xb1, 0x7b, 0x00, 0x8d, 0xf3, 0x8c, 0xc6,
	0x29, 0x47, 0x6d, 0xa8, 0xde, 0x04, 0x8e, 0xb1, 0x63, 0xec, 0x59, 0xb8, 0x7a, 0x13, 0xa0, 0x2d,
	0x30, 0xaf, 0x68, 0xc2, 0x49, 0xc2, 0x99, 0x53, 0x95, 0xd9, 0x3c, 0x76, 0x0f, 0xc1, 0xc4, 0x84,
	0xa5, 0x34, 0x61, 0xe4, 0x3e, 0x5e, 0x76, 0x8b, 0x69, 0x9e, 0x8e, 0xdd, 0xd7, 0x60, 0x61, 0x1a,
	0x91, 0xf7, 0x99, 0x1f, 0x13, 0xe4, 0x40, 0x33, 0x5c, 0x08, 0x4d, 0x99, 0x64, 0x9b, 0x58, 0x87,
	0xa8, 0x03, 0x35, 0xba, 0xe0, 0x92, 0x6d, 0x62, 0xf1, 0xe9, 0x7e, 0x00, 0x18, 0x0a, 0x50, 0x31,
	0x1f, 0x82, 0x95, 0x4a, 0xd1, 0xb3, 0xfc, 0xcf, 0xa6, 0x4a, 0x8c, 0x03, 0xf4, 0x08, 0x6c, 0xfd,
	0x3f, 0x01, 0x2b, 0x09, 0xa0, 0x53, 0xe3, 0xc0, 0x9d, 0xc0, 0xa6, 0x1a, 0x59, 0x8f, 0xf0, 0x0f,
	0x4d, 0xff, 0x3c, 0x94, 0x73, 0xb7, 0xdf, 0x38, 0x49, 0x17, 0xfc, 0xef, 0x4d, 0xdd, 0xb7, 0xd0,
	0x3d, 0x8a, 0x22, 0xcd, 0xba, 0x9d, 0xed, 0x39, 0x58, 0xba, 0x33, 0x73, 0x8c, 0x9d, 0xda, 0x9e,
	0xdd, 0x5f, 0xef, 0x31, 0x1e, 0xf7, 0x74, 0x1d, 0x2e, 0x70, 0xf7, 0x00, 0xba, 0x98, 0xf8, 0xc1,
	0xdd, 0x41, 0x96, 0x16, 0x60, 0xac, 0x2c, 0xe0, 0x67, 0x0d, 0xec, 0xf3, 0xc8, 0xff, 0x4e, 0xb2,
	0x42, 0xa4, 0x0c, 0xcb, 0x22, 0x65, 0x62, 0x1c, 0xa0, 0x27, 0x50, 0x67, 0xdc, 0xe7, 0x6a, 0x6c,
	0xbb, 0xdf, 0x96, 0x5a, 0x86, 0x7e, 0x4c, 0x3c, 0x91, 0x1d, 0x55, 0xb0, 0x82, 0xd1, 0x2e, 0xd4,
	0xa5, 0x7d, 0x4e, 0x4d, 0xd6, 0x6d, 0xa8, 0xba, 0xdc, 0x33, 0x51, 0x28, 0x71, 0xe4, 0xc1, 0x7f,
	0xb7, 0x2b, 0x29, 0x54, 0x8a, 0x85, 0x39, 0x6b, 0x92, 0xb8, 0x2d, 0x89, 0xbf, 0x5b, 0xe8, 0xa8,
	0x82, 0x37, 0xd3, 0x55, 0x0c, 0x9d, 0xc0, 0xc6, 0x52, 0x53, 0xa7, 0x2e, 0xdb, 0x39, 0xf7, 0xb4,
	0xd3, 0x9d, 0xda, 0x77, 0x3b, 0xa1, 0x97, 0x60, 0x67, 0x34, 0x22, 0xb3, 0xab, 0xb9, 0x9f, 0x84,
	0xc4, 0x69, 0x94, 0x06, 0xce, 0xaf, 0x76, 0x54, 0xc1, 0x20, 0x8a, 0x4e, 0x64, 0x0d, 0x7a, 0x03,
	0xeb, 0x7e, 0x14, 0xcd, 0x0a, 0xc7, 0x9a, 0x92, 0xf4, 0xbf, 0x24, 0xad, 0x98, 0x3b, 0xaa, 0xe0,
	0x96, 0x5f, 0x4a, 0x0a, 0x7a, 0x46, 0xfc, 0xa0, 0x10, 0x6d, 0x96, 0xe8, 0x2b, 0xce, 0x0a, 0x7a,
	0x56, 0x4a, 0x1e, 0x37, 0xa1, 0x7e, 0x2d, 0x00, 0x77, 0x17, 0x5a, 0xc7, 0x24, 0xbc, 0x49, 0x84,
	0x2f, 0x98, 0x7c, 0x43, 0x0f, 0xa0, 0x99, 0x51, 0x1a, 0x17, 0x7e, 0x36, 0x44, 0x38, 0x0e, 0xdc,
	0xaf, 0x60, 0xe5, 0xde, 0xa1, 0xa7, 0x00, 0xa1, 0x1f, 0x93, 0x99, 0xf2, 0x57, 0x14, 0xb6, 0xfb,
	0x20, 0x7f, 0xed, 0x4d, 0x8f, 0xa6, 0x03, 0x6c, 0x85, 0x79, 0xe9, 0x63, 0x68, 0xaa, 0x65, 0x89,
	0xb7, 0x40, 0xdc, 0xa4, 0x5d, 0xde, 0xab, 0xc6, 0x9e, 0x9d, 0x42, 0x5d, 0x52, 0x91, 0x0d, 0xcd,
	0x8b, 0xc9, 0xc7, 0xc9, 0xd9, 0xe7, 0x49, 0xa7, 0x82, 0xda, 0x00, 0x78, 0xe0, 0x9d, 0x9f, 0x4d,
	0xde, 0x8d, 0x27, 0xc3, 0x8e, 0x21, 0x40, 0x3c, 0x38, 0x92, 0x41, 0x15, 0xb5, 0xc0, 0x1c, 0x5e,
	0x0c, 0x3c, 0x4f, 0x44, 0x35, 0x05, 0x79, 0x17, 0xa7, 0x53, 0xaf, 0xb3, 0xd6, 0xff, 0x01, 0x5d,
	0x4f, 0xbc, 0x5e, 0xd3, 0x39, 0xf9, 0xa4, 0xdf, 0x2e, 0xd4, 0x03, 0x2b, 0x1f, 0x15, 0x75, 0xa5,
	0x8a, 0xf2, 0xe8, 0x5b, 0x4b, 0x07, 0x8a, 0x0e, 0xa1, 0xa5, 0x6e, 0xdd, 0xe3, 0x19, 0xf1, 0x63,
	0xd4, 0x51, 0xc2, 0x8b, 0xf3, 0xdf, 0x5a, 0xc9, 0xec, 0x19, 0x2f, 0x8c, 0xe3, 0xee, 0x97, 0x8d,
	0x90, 0xee, 0xcb, 0x97, 0xf2, 0x72, 0x71, 0xbd, 0xcf, 0x78, 0x7c, 0xd9, 0x90, 0xd1, 0xab, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xef, 0xca, 0x91, 0x58, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShootTheMessengerClient is the client API for ShootTheMessenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShootTheMessengerClient interface {
	BeginGame(ctx context.Context, in *BeginGameReq, opts ...grpc.CallOption) (*GameState, error)
	PlayerStream(ctx context.Context, opts ...grpc.CallOption) (ShootTheMessenger_PlayerStreamClient, error)
}

type shootTheMessengerClient struct {
	cc *grpc.ClientConn
}

func NewShootTheMessengerClient(cc *grpc.ClientConn) ShootTheMessengerClient {
	return &shootTheMessengerClient{cc}
}

func (c *shootTheMessengerClient) BeginGame(ctx context.Context, in *BeginGameReq, opts ...grpc.CallOption) (*GameState, error) {
	out := new(GameState)
	err := c.cc.Invoke(ctx, "/stm.ShootTheMessenger/BeginGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shootTheMessengerClient) PlayerStream(ctx context.Context, opts ...grpc.CallOption) (ShootTheMessenger_PlayerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ShootTheMessenger_serviceDesc.Streams[0], "/stm.ShootTheMessenger/PlayerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &shootTheMessengerPlayerStreamClient{stream}
	return x, nil
}

type ShootTheMessenger_PlayerStreamClient interface {
	Send(*PlayerFrame) error
	Recv() (*PlayerFrame, error)
	grpc.ClientStream
}

type shootTheMessengerPlayerStreamClient struct {
	grpc.ClientStream
}

func (x *shootTheMessengerPlayerStreamClient) Send(m *PlayerFrame) error {
	return x.ClientStream.SendMsg(m)
}

func (x *shootTheMessengerPlayerStreamClient) Recv() (*PlayerFrame, error) {
	m := new(PlayerFrame)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShootTheMessengerServer is the server API for ShootTheMessenger service.
type ShootTheMessengerServer interface {
	BeginGame(context.Context, *BeginGameReq) (*GameState, error)
	PlayerStream(ShootTheMessenger_PlayerStreamServer) error
}

func RegisterShootTheMessengerServer(s *grpc.Server, srv ShootTheMessengerServer) {
	s.RegisterService(&_ShootTheMessenger_serviceDesc, srv)
}

func _ShootTheMessenger_BeginGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShootTheMessengerServer).BeginGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stm.ShootTheMessenger/BeginGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShootTheMessengerServer).BeginGame(ctx, req.(*BeginGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShootTheMessenger_PlayerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ShootTheMessengerServer).PlayerStream(&shootTheMessengerPlayerStreamServer{stream})
}

type ShootTheMessenger_PlayerStreamServer interface {
	Send(*PlayerFrame) error
	Recv() (*PlayerFrame, error)
	grpc.ServerStream
}

type shootTheMessengerPlayerStreamServer struct {
	grpc.ServerStream
}

func (x *shootTheMessengerPlayerStreamServer) Send(m *PlayerFrame) error {
	return x.ServerStream.SendMsg(m)
}

func (x *shootTheMessengerPlayerStreamServer) Recv() (*PlayerFrame, error) {
	m := new(PlayerFrame)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ShootTheMessenger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stm.ShootTheMessenger",
	HandlerType: (*ShootTheMessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BeginGame",
			Handler:    _ShootTheMessenger_BeginGame_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PlayerStream",
			Handler:       _ShootTheMessenger_PlayerStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stm/shoot_the_messenger.proto",
}
