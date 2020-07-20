// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: pkg/game/api/deathroll.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GameEvent int32

const (
	GameEvent_PlayerLost GameEvent = 0
	GameEvent_PlayerWon  GameEvent = 1
)

// Enum value maps for GameEvent.
var (
	GameEvent_name = map[int32]string{
		0: "PlayerLost",
		1: "PlayerWon",
	}
	GameEvent_value = map[string]int32{
		"PlayerLost": 0,
		"PlayerWon":  1,
	}
)

func (x GameEvent) Enum() *GameEvent {
	p := new(GameEvent)
	*p = x
	return p
}

func (x GameEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_game_api_deathroll_proto_enumTypes[0].Descriptor()
}

func (GameEvent) Type() protoreflect.EnumType {
	return &file_pkg_game_api_deathroll_proto_enumTypes[0]
}

func (x GameEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameEvent.Descriptor instead.
func (GameEvent) EnumDescriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{0}
}

//limit describes inital value for deathroll random's upper bound limit
type StartGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *StartGameRequest) Reset() {
	*x = StartGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameRequest) ProtoMessage() {}

func (x *StartGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameRequest.ProtoReflect.Descriptor instead.
func (*StartGameRequest) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{0}
}

func (x *StartGameRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type StartGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartGameResponse) Reset() {
	*x = StartGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameResponse) ProtoMessage() {}

func (x *StartGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameResponse.ProtoReflect.Descriptor instead.
func (*StartGameResponse) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{1}
}

//userId is unique user's discord identifier
type AddPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AddPlayerRequest) Reset() {
	*x = AddPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlayerRequest) ProtoMessage() {}

func (x *AddPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlayerRequest.ProtoReflect.Descriptor instead.
func (*AddPlayerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{2}
}

func (x *AddPlayerRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

//returns whether or not adding was successfull
type AddPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Added bool `protobuf:"varint,1,opt,name=added,proto3" json:"added,omitempty"`
}

func (x *AddPlayerResponse) Reset() {
	*x = AddPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlayerResponse) ProtoMessage() {}

func (x *AddPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlayerResponse.ProtoReflect.Descriptor instead.
func (*AddPlayerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{3}
}

func (x *AddPlayerResponse) GetAdded() bool {
	if x != nil {
		return x.Added
	}
	return false
}

type RollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RollRequest) Reset() {
	*x = RollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollRequest) ProtoMessage() {}

func (x *RollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollRequest.ProtoReflect.Descriptor instead.
func (*RollRequest) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{4}
}

func (x *RollRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

//describes roll parameters and limits, if roll wasnt successfull we should send an error via optional filed error
type RollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Low    int32  `protobuf:"varint,3,opt,name=low,proto3" json:"low,omitempty"`
	High   int32  `protobuf:"varint,4,opt,name=high,proto3" json:"high,omitempty"`
	Value  int32  `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	UserId uint64 `protobuf:"varint,6,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RollResponse) Reset() {
	*x = RollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollResponse) ProtoMessage() {}

func (x *RollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollResponse.ProtoReflect.Descriptor instead.
func (*RollResponse) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{5}
}

func (x *RollResponse) GetLow() int32 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *RollResponse) GetHigh() int32 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *RollResponse) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *RollResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetNextPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNextPlayerRequest) Reset() {
	*x = GetNextPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNextPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNextPlayerRequest) ProtoMessage() {}

func (x *GetNextPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNextPlayerRequest.ProtoReflect.Descriptor instead.
func (*GetNextPlayerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{6}
}

type GetNextPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetNextPlayerResponse) Reset() {
	*x = GetNextPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNextPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNextPlayerResponse) ProtoMessage() {}

func (x *GetNextPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNextPlayerResponse.ProtoReflect.Descriptor instead.
func (*GetNextPlayerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{7}
}

func (x *GetNextPlayerResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetGameEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGameEventsRequest) Reset() {
	*x = GetGameEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameEventsRequest) ProtoMessage() {}

func (x *GetGameEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameEventsRequest.ProtoReflect.Descriptor instead.
func (*GetGameEventsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{8}
}

type GetGameEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event  GameEvent `protobuf:"varint,1,opt,name=event,proto3,enum=api.GameEvent" json:"event,omitempty"`
	UserId uint64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetGameEventsResponse) Reset() {
	*x = GetGameEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_game_api_deathroll_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameEventsResponse) ProtoMessage() {}

func (x *GetGameEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_game_api_deathroll_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameEventsResponse.ProtoReflect.Descriptor instead.
func (*GetGameEventsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_game_api_deathroll_proto_rawDescGZIP(), []int{9}
}

func (x *GetGameEventsResponse) GetEvent() GameEvent {
	if x != nil {
		return x.Event
	}
	return GameEvent_PlayerLost
}

func (x *GetGameEventsResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_pkg_game_api_deathroll_proto protoreflect.FileDescriptor

var file_pkg_game_api_deathroll_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x65, 0x61, 0x74, 0x68, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x13, 0x0a,
	0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x29,
	0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x22, 0x25, 0x0a, 0x0b, 0x52, 0x6f, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x62, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6c,
	0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x16, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x2a, 0x0a, 0x09,
	0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4c, 0x6f, 0x73, 0x74, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x57, 0x6f, 0x6e, 0x10, 0x01, 0x32, 0xc4, 0x02, 0x0a, 0x0b, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x6c,
	0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42,
	0x0e, 0x5a, 0x0c, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_game_api_deathroll_proto_rawDescOnce sync.Once
	file_pkg_game_api_deathroll_proto_rawDescData = file_pkg_game_api_deathroll_proto_rawDesc
)

func file_pkg_game_api_deathroll_proto_rawDescGZIP() []byte {
	file_pkg_game_api_deathroll_proto_rawDescOnce.Do(func() {
		file_pkg_game_api_deathroll_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_game_api_deathroll_proto_rawDescData)
	})
	return file_pkg_game_api_deathroll_proto_rawDescData
}

var file_pkg_game_api_deathroll_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_game_api_deathroll_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_game_api_deathroll_proto_goTypes = []interface{}{
	(GameEvent)(0),                // 0: api.GameEvent
	(*StartGameRequest)(nil),      // 1: api.StartGameRequest
	(*StartGameResponse)(nil),     // 2: api.StartGameResponse
	(*AddPlayerRequest)(nil),      // 3: api.AddPlayerRequest
	(*AddPlayerResponse)(nil),     // 4: api.AddPlayerResponse
	(*RollRequest)(nil),           // 5: api.RollRequest
	(*RollResponse)(nil),          // 6: api.RollResponse
	(*GetNextPlayerRequest)(nil),  // 7: api.GetNextPlayerRequest
	(*GetNextPlayerResponse)(nil), // 8: api.GetNextPlayerResponse
	(*GetGameEventsRequest)(nil),  // 9: api.GetGameEventsRequest
	(*GetGameEventsResponse)(nil), // 10: api.GetGameEventsResponse
}
var file_pkg_game_api_deathroll_proto_depIdxs = []int32{
	0,  // 0: api.GetGameEventsResponse.event:type_name -> api.GameEvent
	1,  // 1: api.GameService.StartGame:input_type -> api.StartGameRequest
	3,  // 2: api.GameService.AddPlayer:input_type -> api.AddPlayerRequest
	7,  // 3: api.GameService.GetNextPlayer:input_type -> api.GetNextPlayerRequest
	5,  // 4: api.GameService.Roll:input_type -> api.RollRequest
	9,  // 5: api.GameService.GetGameEvents:input_type -> api.GetGameEventsRequest
	2,  // 6: api.GameService.StartGame:output_type -> api.StartGameResponse
	4,  // 7: api.GameService.AddPlayer:output_type -> api.AddPlayerResponse
	8,  // 8: api.GameService.GetNextPlayer:output_type -> api.GetNextPlayerResponse
	6,  // 9: api.GameService.Roll:output_type -> api.RollResponse
	10, // 10: api.GameService.GetGameEvents:output_type -> api.GetGameEventsResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_game_api_deathroll_proto_init() }
func file_pkg_game_api_deathroll_proto_init() {
	if File_pkg_game_api_deathroll_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_game_api_deathroll_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_game_api_deathroll_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_game_api_deathroll_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPlayerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_game_api_deathroll_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPlayerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_game_api_deathroll_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_game_api_deathroll_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_game_api_deathroll_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNextPlayerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_game_api_deathroll_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNextPlayerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_game_api_deathroll_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGameEventsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_game_api_deathroll_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGameEventsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_game_api_deathroll_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_game_api_deathroll_proto_goTypes,
		DependencyIndexes: file_pkg_game_api_deathroll_proto_depIdxs,
		EnumInfos:         file_pkg_game_api_deathroll_proto_enumTypes,
		MessageInfos:      file_pkg_game_api_deathroll_proto_msgTypes,
	}.Build()
	File_pkg_game_api_deathroll_proto = out.File
	file_pkg_game_api_deathroll_proto_rawDesc = nil
	file_pkg_game_api_deathroll_proto_goTypes = nil
	file_pkg_game_api_deathroll_proto_depIdxs = nil
}
