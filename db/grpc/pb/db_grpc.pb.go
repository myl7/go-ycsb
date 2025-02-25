// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: pb/db.proto

package pb

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

// DBClient is the client API for DB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBClient interface {
	Read(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Val, error)
	Scan(ctx context.Context, in *KeyScan, opts ...grpc.CallOption) (DB_ScanClient, error)
	Put(ctx context.Context, in *KeyVal, opts ...grpc.CallOption) (*None, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*None, error)
}

type dBClient struct {
	cc grpc.ClientConnInterface
}

func NewDBClient(cc grpc.ClientConnInterface) DBClient {
	return &dBClient{cc}
}

func (c *dBClient) Read(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Val, error) {
	out := new(Val)
	err := c.cc.Invoke(ctx, "/db.DB/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Scan(ctx context.Context, in *KeyScan, opts ...grpc.CallOption) (DB_ScanClient, error) {
	stream, err := c.cc.NewStream(ctx, &DB_ServiceDesc.Streams[0], "/db.DB/Scan", opts...)
	if err != nil {
		return nil, err
	}
	x := &dBScanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DB_ScanClient interface {
	Recv() (*Val, error)
	grpc.ClientStream
}

type dBScanClient struct {
	grpc.ClientStream
}

func (x *dBScanClient) Recv() (*Val, error) {
	m := new(Val)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dBClient) Put(ctx context.Context, in *KeyVal, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/db.DB/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/db.DB/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBServer is the server API for DB service.
// All implementations must embed UnimplementedDBServer
// for forward compatibility
type DBServer interface {
	Read(context.Context, *Key) (*Val, error)
	Scan(*KeyScan, DB_ScanServer) error
	Put(context.Context, *KeyVal) (*None, error)
	Delete(context.Context, *Key) (*None, error)
	mustEmbedUnimplementedDBServer()
}

// UnimplementedDBServer must be embedded to have forward compatible implementations.
type UnimplementedDBServer struct {
}

func (UnimplementedDBServer) Read(context.Context, *Key) (*Val, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedDBServer) Scan(*KeyScan, DB_ScanServer) error {
	return status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (UnimplementedDBServer) Put(context.Context, *KeyVal) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedDBServer) Delete(context.Context, *Key) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDBServer) mustEmbedUnimplementedDBServer() {}

// UnsafeDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBServer will
// result in compilation errors.
type UnsafeDBServer interface {
	mustEmbedUnimplementedDBServer()
}

func RegisterDBServer(s grpc.ServiceRegistrar, srv DBServer) {
	s.RegisterService(&DB_ServiceDesc, srv)
}

func _DB_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DB/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Read(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Scan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KeyScan)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DBServer).Scan(m, &dBScanServer{stream})
}

type DB_ScanServer interface {
	Send(*Val) error
	grpc.ServerStream
}

type dBScanServer struct {
	grpc.ServerStream
}

func (x *dBScanServer) Send(m *Val) error {
	return x.ServerStream.SendMsg(m)
}

func _DB_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyVal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DB/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Put(ctx, req.(*KeyVal))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DB/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// DB_ServiceDesc is the grpc.ServiceDesc for DB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.DB",
	HandlerType: (*DBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _DB_Read_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _DB_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DB_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Scan",
			Handler:       _DB_Scan_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/db.proto",
}
