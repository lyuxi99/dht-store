// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.0--rc2
// source: chord.proto

package proto

import (
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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Node) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type SuccessorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *SuccessorList) Reset() {
	*x = SuccessorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessorList) ProtoMessage() {}

func (x *SuccessorList) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessorList.ProtoReflect.Descriptor instead.
func (*SuccessorList) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{2}
}

func (x *SuccessorList) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type PutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key           []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expire        int64  `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
	InitiatorAddr string `protobuf:"bytes,4,opt,name=initiatorAddr,proto3" json:"initiatorAddr,omitempty"`
	Replication   int32  `protobuf:"varint,5,opt,name=replication,proto3" json:"replication,omitempty"`
}

func (x *PutReq) Reset() {
	*x = PutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutReq) ProtoMessage() {}

func (x *PutReq) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutReq.ProtoReflect.Descriptor instead.
func (*PutReq) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{3}
}

func (x *PutReq) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PutReq) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *PutReq) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

func (x *PutReq) GetInitiatorAddr() string {
	if x != nil {
		return x.InitiatorAddr
	}
	return ""
}

func (x *PutReq) GetReplication() int32 {
	if x != nil {
		return x.Replication
	}
	return 0
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{4}
}

func (x *GetReq) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Ok    bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{5}
}

func (x *GetResp) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *GetResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{6}
}

var File_chord_proto protoreflect.FileDescriptor

var file_chord_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x04, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x32, 0x0a, 0x0d, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x06, 0x50,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1a, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f,
	0x69, 0x64, 0x32, 0x80, 0x02, 0x0a, 0x05, 0x43, 0x68, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x0d,
	0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x09, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65,
	0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x26, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x44, 0x48, 0x54, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chord_proto_rawDescOnce sync.Once
	file_chord_proto_rawDescData = file_chord_proto_rawDesc
)

func file_chord_proto_rawDescGZIP() []byte {
	file_chord_proto_rawDescOnce.Do(func() {
		file_chord_proto_rawDescData = protoimpl.X.CompressGZIP(file_chord_proto_rawDescData)
	})
	return file_chord_proto_rawDescData
}

var file_chord_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_chord_proto_goTypes = []interface{}{
	(*Id)(nil),            // 0: proto.Id
	(*Node)(nil),          // 1: proto.Node
	(*SuccessorList)(nil), // 2: proto.SuccessorList
	(*PutReq)(nil),        // 3: proto.PutReq
	(*GetReq)(nil),        // 4: proto.GetReq
	(*GetResp)(nil),       // 5: proto.GetResp
	(*Void)(nil),          // 6: proto.Void
}
var file_chord_proto_depIdxs = []int32{
	1, // 0: proto.SuccessorList.nodes:type_name -> proto.Node
	0, // 1: proto.Chord.FindSuccessor:input_type -> proto.Id
	1, // 2: proto.Chord.Notify:input_type -> proto.Node
	6, // 3: proto.Chord.GetPredecessor:input_type -> proto.Void
	6, // 4: proto.Chord.Ping:input_type -> proto.Void
	3, // 5: proto.Chord.Put:input_type -> proto.PutReq
	4, // 6: proto.Chord.Get:input_type -> proto.GetReq
	1, // 7: proto.Chord.FindSuccessor:output_type -> proto.Node
	2, // 8: proto.Chord.Notify:output_type -> proto.SuccessorList
	1, // 9: proto.Chord.GetPredecessor:output_type -> proto.Node
	6, // 10: proto.Chord.Ping:output_type -> proto.Void
	6, // 11: proto.Chord.Put:output_type -> proto.Void
	5, // 12: proto.Chord.Get:output_type -> proto.GetResp
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chord_proto_init() }
func file_chord_proto_init() {
	if File_chord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_chord_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_chord_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessorList); i {
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
		file_chord_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutReq); i {
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
		file_chord_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_chord_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResp); i {
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
		file_chord_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_chord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chord_proto_goTypes,
		DependencyIndexes: file_chord_proto_depIdxs,
		MessageInfos:      file_chord_proto_msgTypes,
	}.Build()
	File_chord_proto = out.File
	file_chord_proto_rawDesc = nil
	file_chord_proto_goTypes = nil
	file_chord_proto_depIdxs = nil
}
