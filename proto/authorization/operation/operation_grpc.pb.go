// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/authorization/operation/operation.proto

package operation

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

// OperationClient is the client API for Operation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationClient interface {
	Fetch(ctx context.Context, in *FetchRequestOperation, opts ...grpc.CallOption) (*FetchResponseOperation, error)
	Store(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*MainResponseOperation, error)
	FindByID(ctx context.Context, in *FindByIDOperationRequest, opts ...grpc.CallOption) (*MainResponseOperation, error)
	Update(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*MainResponseOperation, error)
	Delete(ctx context.Context, in *FindByIDOperationRequest, opts ...grpc.CallOption) (*MainResponseOperation, error)
}

type operationClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationClient(cc grpc.ClientConnInterface) OperationClient {
	return &operationClient{cc}
}

func (c *operationClient) Fetch(ctx context.Context, in *FetchRequestOperation, opts ...grpc.CallOption) (*FetchResponseOperation, error) {
	out := new(FetchResponseOperation)
	err := c.cc.Invoke(ctx, "/operation.Operation/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) Store(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*MainResponseOperation, error) {
	out := new(MainResponseOperation)
	err := c.cc.Invoke(ctx, "/operation.Operation/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) FindByID(ctx context.Context, in *FindByIDOperationRequest, opts ...grpc.CallOption) (*MainResponseOperation, error) {
	out := new(MainResponseOperation)
	err := c.cc.Invoke(ctx, "/operation.Operation/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) Update(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*MainResponseOperation, error) {
	out := new(MainResponseOperation)
	err := c.cc.Invoke(ctx, "/operation.Operation/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) Delete(ctx context.Context, in *FindByIDOperationRequest, opts ...grpc.CallOption) (*MainResponseOperation, error) {
	out := new(MainResponseOperation)
	err := c.cc.Invoke(ctx, "/operation.Operation/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServer is the server API for Operation service.
// All implementations must embed UnimplementedOperationServer
// for forward compatibility
type OperationServer interface {
	Fetch(context.Context, *FetchRequestOperation) (*FetchResponseOperation, error)
	Store(context.Context, *OperationRequest) (*MainResponseOperation, error)
	FindByID(context.Context, *FindByIDOperationRequest) (*MainResponseOperation, error)
	Update(context.Context, *OperationRequest) (*MainResponseOperation, error)
	Delete(context.Context, *FindByIDOperationRequest) (*MainResponseOperation, error)
	mustEmbedUnimplementedOperationServer()
}

// UnimplementedOperationServer must be embedded to have forward compatible implementations.
type UnimplementedOperationServer struct {
}

func (UnimplementedOperationServer) Fetch(context.Context, *FetchRequestOperation) (*FetchResponseOperation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedOperationServer) Store(context.Context, *OperationRequest) (*MainResponseOperation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedOperationServer) FindByID(context.Context, *FindByIDOperationRequest) (*MainResponseOperation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedOperationServer) Update(context.Context, *OperationRequest) (*MainResponseOperation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOperationServer) Delete(context.Context, *FindByIDOperationRequest) (*MainResponseOperation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOperationServer) mustEmbedUnimplementedOperationServer() {}

// UnsafeOperationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationServer will
// result in compilation errors.
type UnsafeOperationServer interface {
	mustEmbedUnimplementedOperationServer()
}

func RegisterOperationServer(s grpc.ServiceRegistrar, srv OperationServer) {
	s.RegisterService(&Operation_ServiceDesc, srv)
}

func _Operation_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequestOperation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operation.Operation/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).Fetch(ctx, req.(*FetchRequestOperation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operation.Operation/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).Store(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operation.Operation/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).FindByID(ctx, req.(*FindByIDOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operation.Operation/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).Update(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operation.Operation/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).Delete(ctx, req.(*FindByIDOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Operation_ServiceDesc is the grpc.ServiceDesc for Operation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "operation.Operation",
	HandlerType: (*OperationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Operation_Fetch_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Operation_Store_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _Operation_FindByID_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Operation_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Operation_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/authorization/operation/operation.proto",
}