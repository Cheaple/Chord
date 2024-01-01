// RPCs for Chord node comminications

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: chord/chord.proto

package chord

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Chord_Locate_FullMethodName           = "/Chord/Locate"
	Chord_Check_FullMethodName            = "/Chord/Check"
	Chord_GetPredecessor_FullMethodName   = "/Chord/GetPredecessor"
	Chord_GetSuccessorList_FullMethodName = "/Chord/GetSuccessorList"
	Chord_SetPredecessor_FullMethodName   = "/Chord/SetPredecessor"
	Chord_CheckKey_FullMethodName         = "/Chord/CheckKey"
	Chord_UploadFile_FullMethodName       = "/Chord/UploadFile"
	Chord_DownloadFile_FullMethodName     = "/Chord/DownloadFile"
)

// ChordClient is the client API for Chord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChordClient interface {
	// Locate target identifier in the Chord ring
	Locate(ctx context.Context, in *BytesMsg, opts ...grpc.CallOption) (*NodeEntry, error)
	// Check failure (for check_predecessor() function in the paper)
	Check(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*BoolMsg, error)
	// Get the target node's current predecessor
	GetPredecessor(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*NodeEntry, error)
	// Get the target node's successor list
	GetSuccessorList(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*NodeList, error)
	// Set the target node's predecessor (for notify() function in the paper)
	SetPredecessor(ctx context.Context, in *NodeEntry, opts ...grpc.CallOption) (*BoolMsg, error)
	// Check whether a key exists in the target node's buckets
	CheckKey(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (*BoolMsg, error)
	// Upload a file to the target node
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (Chord_UploadFileClient, error)
	// Download a file from the target node
	DownloadFile(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (Chord_DownloadFileClient, error)
}

type chordClient struct {
	cc grpc.ClientConnInterface
}

func NewChordClient(cc grpc.ClientConnInterface) ChordClient {
	return &chordClient{cc}
}

func (c *chordClient) Locate(ctx context.Context, in *BytesMsg, opts ...grpc.CallOption) (*NodeEntry, error) {
	out := new(NodeEntry)
	err := c.cc.Invoke(ctx, Chord_Locate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) Check(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*BoolMsg, error) {
	out := new(BoolMsg)
	err := c.cc.Invoke(ctx, Chord_Check_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) GetPredecessor(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*NodeEntry, error) {
	out := new(NodeEntry)
	err := c.cc.Invoke(ctx, Chord_GetPredecessor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) GetSuccessorList(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*NodeList, error) {
	out := new(NodeList)
	err := c.cc.Invoke(ctx, Chord_GetSuccessorList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) SetPredecessor(ctx context.Context, in *NodeEntry, opts ...grpc.CallOption) (*BoolMsg, error) {
	out := new(BoolMsg)
	err := c.cc.Invoke(ctx, Chord_SetPredecessor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) CheckKey(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (*BoolMsg, error) {
	out := new(BoolMsg)
	err := c.cc.Invoke(ctx, Chord_CheckKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (Chord_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chord_ServiceDesc.Streams[0], Chord_UploadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chordUploadFileClient{stream}
	return x, nil
}

type Chord_UploadFileClient interface {
	Send(*FileMsg) error
	CloseAndRecv() (*BoolMsg, error)
	grpc.ClientStream
}

type chordUploadFileClient struct {
	grpc.ClientStream
}

func (x *chordUploadFileClient) Send(m *FileMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chordUploadFileClient) CloseAndRecv() (*BoolMsg, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BoolMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chordClient) DownloadFile(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (Chord_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chord_ServiceDesc.Streams[1], Chord_DownloadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chordDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chord_DownloadFileClient interface {
	Recv() (*FileMsg, error)
	grpc.ClientStream
}

type chordDownloadFileClient struct {
	grpc.ClientStream
}

func (x *chordDownloadFileClient) Recv() (*FileMsg, error) {
	m := new(FileMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChordServer is the server API for Chord service.
// All implementations should embed UnimplementedChordServer
// for forward compatibility
type ChordServer interface {
	// Locate target identifier in the Chord ring
	Locate(context.Context, *BytesMsg) (*NodeEntry, error)
	// Check failure (for check_predecessor() function in the paper)
	Check(context.Context, *EmptyMsg) (*BoolMsg, error)
	// Get the target node's current predecessor
	GetPredecessor(context.Context, *EmptyMsg) (*NodeEntry, error)
	// Get the target node's successor list
	GetSuccessorList(context.Context, *EmptyMsg) (*NodeList, error)
	// Set the target node's predecessor (for notify() function in the paper)
	SetPredecessor(context.Context, *NodeEntry) (*BoolMsg, error)
	// Check whether a key exists in the target node's buckets
	CheckKey(context.Context, *StringMsg) (*BoolMsg, error)
	// Upload a file to the target node
	UploadFile(Chord_UploadFileServer) error
	// Download a file from the target node
	DownloadFile(*StringMsg, Chord_DownloadFileServer) error
}

// UnimplementedChordServer should be embedded to have forward compatible implementations.
type UnimplementedChordServer struct {
}

func (UnimplementedChordServer) Locate(context.Context, *BytesMsg) (*NodeEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Locate not implemented")
}
func (UnimplementedChordServer) Check(context.Context, *EmptyMsg) (*BoolMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedChordServer) GetPredecessor(context.Context, *EmptyMsg) (*NodeEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredecessor not implemented")
}
func (UnimplementedChordServer) GetSuccessorList(context.Context, *EmptyMsg) (*NodeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuccessorList not implemented")
}
func (UnimplementedChordServer) SetPredecessor(context.Context, *NodeEntry) (*BoolMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPredecessor not implemented")
}
func (UnimplementedChordServer) CheckKey(context.Context, *StringMsg) (*BoolMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckKey not implemented")
}
func (UnimplementedChordServer) UploadFile(Chord_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedChordServer) DownloadFile(*StringMsg, Chord_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}

// UnsafeChordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChordServer will
// result in compilation errors.
type UnsafeChordServer interface {
	mustEmbedUnimplementedChordServer()
}

func RegisterChordServer(s grpc.ServiceRegistrar, srv ChordServer) {
	s.RegisterService(&Chord_ServiceDesc, srv)
}

func _Chord_Locate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BytesMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Locate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chord_Locate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Locate(ctx, req.(*BytesMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chord_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Check(ctx, req.(*EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_GetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chord_GetPredecessor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetPredecessor(ctx, req.(*EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_GetSuccessorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetSuccessorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chord_GetSuccessorList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetSuccessorList(ctx, req.(*EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_SetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).SetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chord_SetPredecessor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).SetPredecessor(ctx, req.(*NodeEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_CheckKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).CheckKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chord_CheckKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).CheckKey(ctx, req.(*StringMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChordServer).UploadFile(&chordUploadFileServer{stream})
}

type Chord_UploadFileServer interface {
	SendAndClose(*BoolMsg) error
	Recv() (*FileMsg, error)
	grpc.ServerStream
}

type chordUploadFileServer struct {
	grpc.ServerStream
}

func (x *chordUploadFileServer) SendAndClose(m *BoolMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chordUploadFileServer) Recv() (*FileMsg, error) {
	m := new(FileMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chord_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StringMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChordServer).DownloadFile(m, &chordDownloadFileServer{stream})
}

type Chord_DownloadFileServer interface {
	Send(*FileMsg) error
	grpc.ServerStream
}

type chordDownloadFileServer struct {
	grpc.ServerStream
}

func (x *chordDownloadFileServer) Send(m *FileMsg) error {
	return x.ServerStream.SendMsg(m)
}

// Chord_ServiceDesc is the grpc.ServiceDesc for Chord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Chord",
	HandlerType: (*ChordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Locate",
			Handler:    _Chord_Locate_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Chord_Check_Handler,
		},
		{
			MethodName: "GetPredecessor",
			Handler:    _Chord_GetPredecessor_Handler,
		},
		{
			MethodName: "GetSuccessorList",
			Handler:    _Chord_GetSuccessorList_Handler,
		},
		{
			MethodName: "SetPredecessor",
			Handler:    _Chord_SetPredecessor_Handler,
		},
		{
			MethodName: "CheckKey",
			Handler:    _Chord_CheckKey_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _Chord_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _Chord_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chord/chord.proto",
}
