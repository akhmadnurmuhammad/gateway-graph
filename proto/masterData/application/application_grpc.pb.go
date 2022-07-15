// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/masterData/application/application.proto

package application

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

// ApplicationClient is the client API for Application service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationClient interface {
	Store(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*MainResponse, error)
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MainResponse, error)
	Delete(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error)
	UpdateActive(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error)
}

type applicationClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationClient(cc grpc.ClientConnInterface) ApplicationClient {
	return &applicationClient{cc}
}

func (c *applicationClient) Store(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/application.Application/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/application.Application/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/application.Application/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/application.Application/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Delete(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/application.Application/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) UpdateActive(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/application.Application/UpdateActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServer is the server API for Application service.
// All implementations must embed UnimplementedApplicationServer
// for forward compatibility
type ApplicationServer interface {
	Store(context.Context, *MasterRequest) (*MainResponse, error)
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	FindByID(context.Context, *FindByIDRequest) (*MainResponse, error)
	Update(context.Context, *UpdateRequest) (*MainResponse, error)
	Delete(context.Context, *FindByIDRequest) (*MainResponse, error)
	UpdateActive(context.Context, *FindByIDRequest) (*MainResponse, error)
	mustEmbedUnimplementedApplicationServer()
}

// UnimplementedApplicationServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationServer struct {
}

func (UnimplementedApplicationServer) Store(context.Context, *MasterRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedApplicationServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedApplicationServer) FindByID(context.Context, *FindByIDRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedApplicationServer) Update(context.Context, *UpdateRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedApplicationServer) Delete(context.Context, *FindByIDRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedApplicationServer) UpdateActive(context.Context, *FindByIDRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateActive not implemented")
}
func (UnimplementedApplicationServer) mustEmbedUnimplementedApplicationServer() {}

// UnsafeApplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationServer will
// result in compilation errors.
type UnsafeApplicationServer interface {
	mustEmbedUnimplementedApplicationServer()
}

func RegisterApplicationServer(s grpc.ServiceRegistrar, srv ApplicationServer) {
	s.RegisterService(&Application_ServiceDesc, srv)
}

func _Application_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Store(ctx, req.(*MasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).FindByID(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Delete(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_UpdateActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).UpdateActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/UpdateActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).UpdateActive(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Application_ServiceDesc is the grpc.ServiceDesc for Application service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Application_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "application.Application",
	HandlerType: (*ApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Application_Store_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _Application_Fetch_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _Application_FindByID_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Application_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Application_Delete_Handler,
		},
		{
			MethodName: "UpdateActive",
			Handler:    _Application_UpdateActive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/masterData/application/application.proto",
}