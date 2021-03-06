// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/authorization/credential/credential.proto

package credential

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

// CredentialClient is the client API for Credential service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CredentialClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*MainResponse, error)
	FindByID(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*MainResponse, error)
	FindByKey(ctx context.Context, in *FindByKeyRequest, opts ...grpc.CallOption) (*MainResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MainResponse, error)
	Delete(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*MainResponse, error)
}

type credentialClient struct {
	cc grpc.ClientConnInterface
}

func NewCredentialClient(cc grpc.ClientConnInterface) CredentialClient {
	return &credentialClient{cc}
}

func (c *credentialClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/credential.Credential/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/credential.Credential/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialClient) FindByID(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/credential.Credential/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialClient) FindByKey(ctx context.Context, in *FindByKeyRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/credential.Credential/FindByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/credential.Credential/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialClient) Delete(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/credential.Credential/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialServer is the server API for Credential service.
// All implementations must embed UnimplementedCredentialServer
// for forward compatibility
type CredentialServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	Store(context.Context, *StoreRequest) (*MainResponse, error)
	FindByID(context.Context, *FindRequest) (*MainResponse, error)
	FindByKey(context.Context, *FindByKeyRequest) (*MainResponse, error)
	Update(context.Context, *UpdateRequest) (*MainResponse, error)
	Delete(context.Context, *FindRequest) (*MainResponse, error)
	mustEmbedUnimplementedCredentialServer()
}

// UnimplementedCredentialServer must be embedded to have forward compatible implementations.
type UnimplementedCredentialServer struct {
}

func (UnimplementedCredentialServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedCredentialServer) Store(context.Context, *StoreRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedCredentialServer) FindByID(context.Context, *FindRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedCredentialServer) FindByKey(context.Context, *FindByKeyRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKey not implemented")
}
func (UnimplementedCredentialServer) Update(context.Context, *UpdateRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCredentialServer) Delete(context.Context, *FindRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCredentialServer) mustEmbedUnimplementedCredentialServer() {}

// UnsafeCredentialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CredentialServer will
// result in compilation errors.
type UnsafeCredentialServer interface {
	mustEmbedUnimplementedCredentialServer()
}

func RegisterCredentialServer(s grpc.ServiceRegistrar, srv CredentialServer) {
	s.RegisterService(&Credential_ServiceDesc, srv)
}

func _Credential_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.Credential/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credential_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.Credential/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credential_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.Credential/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).FindByID(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credential_FindByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).FindByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.Credential/FindByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).FindByKey(ctx, req.(*FindByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credential_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.Credential/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credential_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.Credential/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).Delete(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Credential_ServiceDesc is the grpc.ServiceDesc for Credential service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Credential_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "credential.Credential",
	HandlerType: (*CredentialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Credential_Fetch_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Credential_Store_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _Credential_FindByID_Handler,
		},
		{
			MethodName: "FindByKey",
			Handler:    _Credential_FindByKey_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Credential_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Credential_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/authorization/credential/credential.proto",
}
