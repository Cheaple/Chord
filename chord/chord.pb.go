// RPCs for Chord node comminications

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: chord/chord.proto

package chord

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

type EmptyMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMsg) Reset() {
	*x = EmptyMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_chord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMsg) ProtoMessage() {}

func (x *EmptyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_chord_chord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMsg.ProtoReflect.Descriptor instead.
func (*EmptyMsg) Descriptor() ([]byte, []int) {
	return file_chord_chord_proto_rawDescGZIP(), []int{0}
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier []byte `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_chord_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_chord_chord_proto_msgTypes[1]
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
	return file_chord_chord_proto_rawDescGZIP(), []int{1}
}

func (x *Id) GetIdentifier() []byte {
	if x != nil {
		return x.Identifier
	}
	return nil
}

type NodeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier []byte `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *NodeEntry) Reset() {
	*x = NodeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_chord_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeEntry) ProtoMessage() {}

func (x *NodeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_chord_chord_proto_msgTypes[2]
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
	return file_chord_chord_proto_rawDescGZIP(), []int{2}
}

func (x *NodeEntry) GetIdentifier() []byte {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *NodeEntry) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type NodeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*NodeEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *NodeList) Reset() {
	*x = NodeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_chord_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeList) ProtoMessage() {}

func (x *NodeList) ProtoReflect() protoreflect.Message {
	mi := &file_chord_chord_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeList.ProtoReflect.Descriptor instead.
func (*NodeList) Descriptor() ([]byte, []int) {
	return file_chord_chord_proto_rawDescGZIP(), []int{3}
}

func (x *NodeList) GetEntries() []*NodeEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type BoolMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BoolMsg) Reset() {
	*x = BoolMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_chord_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolMsg) ProtoMessage() {}

func (x *BoolMsg) ProtoReflect() protoreflect.Message {
	mi := &file_chord_chord_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolMsg.ProtoReflect.Descriptor instead.
func (*BoolMsg) Descriptor() ([]byte, []int) {
	return file_chord_chord_proto_rawDescGZIP(), []int{4}
}

func (x *BoolMsg) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type KeyMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KeyMsg) Reset() {
	*x = KeyMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_chord_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyMsg) ProtoMessage() {}

func (x *KeyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_chord_chord_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyMsg.ProtoReflect.Descriptor instead.
func (*KeyMsg) Descriptor() ([]byte, []int) {
	return file_chord_chord_proto_rawDescGZIP(), []int{5}
}

func (x *KeyMsg) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type FileMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FileMsg) Reset() {
	*x = FileMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_chord_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMsg) ProtoMessage() {}

func (x *FileMsg) ProtoReflect() protoreflect.Message {
	mi := &file_chord_chord_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMsg.ProtoReflect.Descriptor instead.
func (*FileMsg) Descriptor() ([]byte, []int) {
	return file_chord_chord_proto_rawDescGZIP(), []int{6}
}

func (x *FileMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_chord_chord_proto protoreflect.FileDescriptor

var file_chord_chord_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x22,
	0x24, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x23,
	0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6c, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x1a, 0x0a, 0x06, 0x4b, 0x65, 0x79, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x1d, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xdc,
	0x01, 0x0a, 0x05, 0x43, 0x68, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x03, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x09, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x1a, 0x09, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d,
	0x73, 0x67, 0x12, 0x27, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x12, 0x09, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x1a,
	0x0a, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x09, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x1a, 0x09, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64,
	0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x0a, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x1a, 0x09, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x12, 0x1d,
	0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x12, 0x07, 0x2e, 0x4b, 0x65, 0x79,
	0x4d, 0x73, 0x67, 0x1a, 0x08, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x4d, 0x73, 0x67, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chord_chord_proto_rawDescOnce sync.Once
	file_chord_chord_proto_rawDescData = file_chord_chord_proto_rawDesc
)

func file_chord_chord_proto_rawDescGZIP() []byte {
	file_chord_chord_proto_rawDescOnce.Do(func() {
		file_chord_chord_proto_rawDescData = protoimpl.X.CompressGZIP(file_chord_chord_proto_rawDescData)
	})
	return file_chord_chord_proto_rawDescData
}

var file_chord_chord_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_chord_chord_proto_goTypes = []interface{}{
	(*EmptyMsg)(nil),  // 0: EmptyMsg
	(*Id)(nil),        // 1: Id
	(*NodeEntry)(nil), // 2: NodeEntry
	(*NodeList)(nil),  // 3: NodeList
	(*BoolMsg)(nil),   // 4: BoolMsg
	(*KeyMsg)(nil),    // 5: KeyMsg
	(*FileMsg)(nil),   // 6: FileMsg
}
var file_chord_chord_proto_depIdxs = []int32{
	2, // 0: NodeList.entries:type_name -> NodeEntry
	1, // 1: Chord.Locate:input_type -> Id
	0, // 2: Chord.Check:input_type -> EmptyMsg
	0, // 3: Chord.GetPredecessor:input_type -> EmptyMsg
	0, // 4: Chord.GetSuccessorList:input_type -> EmptyMsg
	2, // 5: Chord.SetPredecessor:input_type -> NodeEntry
	5, // 6: Chord.CheckKey:input_type -> KeyMsg
	2, // 7: Chord.Locate:output_type -> NodeEntry
	0, // 8: Chord.Check:output_type -> EmptyMsg
	2, // 9: Chord.GetPredecessor:output_type -> NodeEntry
	3, // 10: Chord.GetSuccessorList:output_type -> NodeList
	0, // 11: Chord.SetPredecessor:output_type -> EmptyMsg
	4, // 12: Chord.CheckKey:output_type -> BoolMsg
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chord_chord_proto_init() }
func file_chord_chord_proto_init() {
	if File_chord_chord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chord_chord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMsg); i {
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
		file_chord_chord_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_chord_chord_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_chord_chord_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeList); i {
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
		file_chord_chord_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolMsg); i {
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
		file_chord_chord_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyMsg); i {
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
		file_chord_chord_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMsg); i {
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
			RawDescriptor: file_chord_chord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chord_chord_proto_goTypes,
		DependencyIndexes: file_chord_chord_proto_depIdxs,
		MessageInfos:      file_chord_chord_proto_msgTypes,
	}.Build()
	File_chord_chord_proto = out.File
	file_chord_chord_proto_rawDesc = nil
	file_chord_chord_proto_goTypes = nil
	file_chord_chord_proto_depIdxs = nil
}
