// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: api/proto/v1/films.proto

package v1

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

// FilmServiceClient is the client API for FilmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilmServiceClient interface {
	CreateFilm(ctx context.Context, in *Film, opts ...grpc.CallOption) (*FilmResponse, error)
	GetFilm(ctx context.Context, in *FilmRequest, opts ...grpc.CallOption) (*Film, error)
	UpdateFilm(ctx context.Context, in *Film, opts ...grpc.CallOption) (*Film, error)
	DeleteFilm(ctx context.Context, in *FilmRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type filmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilmServiceClient(cc grpc.ClientConnInterface) FilmServiceClient {
	return &filmServiceClient{cc}
}

func (c *filmServiceClient) CreateFilm(ctx context.Context, in *Film, opts ...grpc.CallOption) (*FilmResponse, error) {
	out := new(FilmResponse)
	err := c.cc.Invoke(ctx, "/films.FilmService/CreateFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceClient) GetFilm(ctx context.Context, in *FilmRequest, opts ...grpc.CallOption) (*Film, error) {
	out := new(Film)
	err := c.cc.Invoke(ctx, "/films.FilmService/GetFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceClient) UpdateFilm(ctx context.Context, in *Film, opts ...grpc.CallOption) (*Film, error) {
	out := new(Film)
	err := c.cc.Invoke(ctx, "/films.FilmService/UpdateFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceClient) DeleteFilm(ctx context.Context, in *FilmRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/films.FilmService/DeleteFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilmServiceServer is the server API for FilmService service.
// All implementations must embed UnimplementedFilmServiceServer
// for forward compatibility
type FilmServiceServer interface {
	CreateFilm(context.Context, *Film) (*FilmResponse, error)
	GetFilm(context.Context, *FilmRequest) (*Film, error)
	UpdateFilm(context.Context, *Film) (*Film, error)
	DeleteFilm(context.Context, *FilmRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedFilmServiceServer()
}

// UnimplementedFilmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFilmServiceServer struct {
}

func (UnimplementedFilmServiceServer) CreateFilm(context.Context, *Film) (*FilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFilm not implemented")
}
func (UnimplementedFilmServiceServer) GetFilm(context.Context, *FilmRequest) (*Film, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilm not implemented")
}
func (UnimplementedFilmServiceServer) UpdateFilm(context.Context, *Film) (*Film, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFilm not implemented")
}
func (UnimplementedFilmServiceServer) DeleteFilm(context.Context, *FilmRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFilm not implemented")
}
func (UnimplementedFilmServiceServer) mustEmbedUnimplementedFilmServiceServer() {}

// UnsafeFilmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilmServiceServer will
// result in compilation errors.
type UnsafeFilmServiceServer interface {
	mustEmbedUnimplementedFilmServiceServer()
}

func RegisterFilmServiceServer(s grpc.ServiceRegistrar, srv FilmServiceServer) {
	s.RegisterService(&FilmService_ServiceDesc, srv)
}

func _FilmService_CreateFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Film)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmServiceServer).CreateFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.FilmService/CreateFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmServiceServer).CreateFilm(ctx, req.(*Film))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmService_GetFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmServiceServer).GetFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.FilmService/GetFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmServiceServer).GetFilm(ctx, req.(*FilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmService_UpdateFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Film)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmServiceServer).UpdateFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.FilmService/UpdateFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmServiceServer).UpdateFilm(ctx, req.(*Film))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmService_DeleteFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmServiceServer).DeleteFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.FilmService/DeleteFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmServiceServer).DeleteFilm(ctx, req.(*FilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FilmService_ServiceDesc is the grpc.ServiceDesc for FilmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "films.FilmService",
	HandlerType: (*FilmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFilm",
			Handler:    _FilmService_CreateFilm_Handler,
		},
		{
			MethodName: "GetFilm",
			Handler:    _FilmService_GetFilm_Handler,
		},
		{
			MethodName: "UpdateFilm",
			Handler:    _FilmService_UpdateFilm_Handler,
		},
		{
			MethodName: "DeleteFilm",
			Handler:    _FilmService_DeleteFilm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/films.proto",
}
