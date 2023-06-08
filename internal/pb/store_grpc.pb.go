// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: store.proto

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

const (
	Storage_SetData_FullMethodName      = "/pb.Storage/SetData"
	Storage_GetData_FullMethodName      = "/pb.Storage/GetData"
	Storage_GetDataArray_FullMethodName = "/pb.Storage/GetDataArray"
	Storage_UpdateData_FullMethodName   = "/pb.Storage/UpdateData"
	Storage_DeleteData_FullMethodName   = "/pb.Storage/DeleteData"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	SetData(ctx context.Context, in *Value, opts ...grpc.CallOption) (*ResponseStatus, error)
	GetData(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	GetDataArray(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DataArray, error)
	UpdateData(ctx context.Context, in *Value, opts ...grpc.CallOption) (*ResponseStatus, error)
	DeleteData(ctx context.Context, in *Key, opts ...grpc.CallOption) (*ResponseStatus, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) SetData(ctx context.Context, in *Value, opts ...grpc.CallOption) (*ResponseStatus, error) {
	out := new(ResponseStatus)
	err := c.cc.Invoke(ctx, Storage_SetData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetData(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, Storage_GetData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetDataArray(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DataArray, error) {
	out := new(DataArray)
	err := c.cc.Invoke(ctx, Storage_GetDataArray_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) UpdateData(ctx context.Context, in *Value, opts ...grpc.CallOption) (*ResponseStatus, error) {
	out := new(ResponseStatus)
	err := c.cc.Invoke(ctx, Storage_UpdateData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeleteData(ctx context.Context, in *Key, opts ...grpc.CallOption) (*ResponseStatus, error) {
	out := new(ResponseStatus)
	err := c.cc.Invoke(ctx, Storage_DeleteData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	SetData(context.Context, *Value) (*ResponseStatus, error)
	GetData(context.Context, *Key) (*Value, error)
	GetDataArray(context.Context, *Empty) (*DataArray, error)
	UpdateData(context.Context, *Value) (*ResponseStatus, error)
	DeleteData(context.Context, *Key) (*ResponseStatus, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) SetData(context.Context, *Value) (*ResponseStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetData not implemented")
}
func (UnimplementedStorageServer) GetData(context.Context, *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedStorageServer) GetDataArray(context.Context, *Empty) (*DataArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataArray not implemented")
}
func (UnimplementedStorageServer) UpdateData(context.Context, *Value) (*ResponseStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateData not implemented")
}
func (UnimplementedStorageServer) DeleteData(context.Context, *Key) (*ResponseStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteData not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_SetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SetData(ctx, req.(*Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetData(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetDataArray_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetDataArray(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GetDataArray_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetDataArray(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_UpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).UpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_UpdateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).UpdateData(ctx, req.(*Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeleteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeleteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_DeleteData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeleteData(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetData",
			Handler:    _Storage_SetData_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _Storage_GetData_Handler,
		},
		{
			MethodName: "GetDataArray",
			Handler:    _Storage_GetDataArray_Handler,
		},
		{
			MethodName: "UpdateData",
			Handler:    _Storage_UpdateData_Handler,
		},
		{
			MethodName: "DeleteData",
			Handler:    _Storage_DeleteData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store.proto",
}
