// RPCs for Chord node comminications

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: chord/rpc/chord.proto

package rpc

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_rpc_chord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_chord_rpc_chord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_chord_rpc_chord_proto_rawDescGZIP(), []int{0}
}

type NodeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *NodeEntry) Reset() {
	*x = NodeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_rpc_chord_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeEntry) ProtoMessage() {}

func (x *NodeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_chord_rpc_chord_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeEntry.ProtoReflect.Descriptor instead.
func (*NodeEntry) Descriptor() ([]byte, []int) {
	return file_chord_rpc_chord_proto_rawDescGZIP(), []int{1}
}

func (x *NodeEntry) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *NodeEntry) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_chord_rpc_chord_proto protoreflect.FileDescriptor

var file_chord_rpc_chord_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x6f, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x35, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x51, 0x0a, 0x05, 0x43, 0x68, 0x6f, 0x72, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chord_rpc_chord_proto_rawDescOnce sync.Once
	file_chord_rpc_chord_proto_rawDescData = file_chord_rpc_chord_proto_rawDesc
)

func file_chord_rpc_chord_proto_rawDescGZIP() []byte {
	file_chord_rpc_chord_proto_rawDescOnce.Do(func() {
		file_chord_rpc_chord_proto_rawDescData = protoimpl.X.CompressGZIP(file_chord_rpc_chord_proto_rawDescData)
	})
	return file_chord_rpc_chord_proto_rawDescData
}

var file_chord_rpc_chord_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chord_rpc_chord_proto_goTypes = []interface{}{
	(*Empty)(nil),     // 0: Empty
	(*NodeEntry)(nil), // 1: NodeEntry
}
var file_chord_rpc_chord_proto_depIdxs = []int32{
	0, // 0: Chord.GetPredecessor:input_type -> Empty
	0, // 1: Chord.GetSuccessor:input_type -> Empty
	1, // 2: Chord.GetPredecessor:output_type -> NodeEntry
	1, // 3: Chord.GetSuccessor:output_type -> NodeEntry
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chord_rpc_chord_proto_init() }
func file_chord_rpc_chord_proto_init() {
	if File_chord_rpc_chord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chord_rpc_chord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_chord_rpc_chord_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeEntry); i {
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
			RawDescriptor: file_chord_rpc_chord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chord_rpc_chord_proto_goTypes,
		DependencyIndexes: file_chord_rpc_chord_proto_depIdxs,
		MessageInfos:      file_chord_rpc_chord_proto_msgTypes,
	}.Build()
	File_chord_rpc_chord_proto = out.File
	file_chord_rpc_chord_proto_rawDesc = nil
	file_chord_rpc_chord_proto_goTypes = nil
	file_chord_rpc_chord_proto_depIdxs = nil
}