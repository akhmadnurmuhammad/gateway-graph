// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/user/user/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	common "xti-gateway-graph-go/proto/common"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
	// rpc ClientStream(stream ClientStreamRequest) returns (ClientStreamResponse) {
	// option (google.api.http) = {
	//   delete: "/api/v1/user/user/delete",
	//   body: "*"
	// };
	// }
	// rpc ServerStream(ServerStreamRequest) returns (stream ServerStreamResponse) {
	// option (google.api.http) = {
	//   delete: "/api/v1/user/user/delete",
	//   body: "*"
	// };
	// }
	// rpc BidiStream(stream BidiStreamRequest) returns (stream BidiStreamResponse) {
	// option (google.api.http) = {
	//   delete: "/api/v1/user/user/delete",
	//   body: "*"
	// };
	// }
	Store(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*MainResponse, error)
	Fetch(ctx context.Context, in *common.FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error)
	Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*MainResponse, error)
	Delete(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/user.User/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Store(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/user.User/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Fetch(ctx context.Context, in *common.FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/user.User/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/user.User/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/user.User/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*MainResponse, error) {
	out := new(MainResponse)
	err := c.cc.Invoke(ctx, "/user.User/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Call(context.Context, *CallRequest) (*CallResponse, error)
	// rpc ClientStream(stream ClientStreamRequest) returns (ClientStreamResponse) {
	// option (google.api.http) = {
	//   delete: "/api/v1/user/user/delete",
	//   body: "*"
	// };
	// }
	// rpc ServerStream(ServerStreamRequest) returns (stream ServerStreamResponse) {
	// option (google.api.http) = {
	//   delete: "/api/v1/user/user/delete",
	//   body: "*"
	// };
	// }
	// rpc BidiStream(stream BidiStreamRequest) returns (stream BidiStreamResponse) {
	// option (google.api.http) = {
	//   delete: "/api/v1/user/user/delete",
	//   body: "*"
	// };
	// }
	Store(context.Context, *UserRequest) (*MainResponse, error)
	Fetch(context.Context, *common.FetchRequest) (*FetchResponse, error)
	FindByID(context.Context, *FindByIDRequest) (*MainResponse, error)
	Update(context.Context, *UserUpdateRequest) (*MainResponse, error)
	Delete(context.Context, *FindByIDRequest) (*MainResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Call(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedUserServer) Store(context.Context, *UserRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedUserServer) Fetch(context.Context, *common.FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedUserServer) FindByID(context.Context, *FindByIDRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedUserServer) Update(context.Context, *UserUpdateRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServer) Delete(context.Context, *FindByIDRequest) (*MainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Store(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Fetch(ctx, req.(*common.FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FindByID(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Update(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Delete(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _User_Call_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _User_Store_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _User_Fetch_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _User_FindByID_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _User_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _User_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/user/user.proto",
}
