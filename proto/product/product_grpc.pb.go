// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package product

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

// ProductGapiClient is the client API for ProductGapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductGapiClient interface {
	InsertProduct(ctx context.Context, in *InsertProductRequest, opts ...grpc.CallOption) (*InsertProductResponse, error)
	GetListProducts(ctx context.Context, in *GetListProductRequest, opts ...grpc.CallOption) (*GetListProductResponse, error)
	GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*GetProductByIDResponse, error)
	UpdateProductByID(ctx context.Context, in *UpdateProductByIDRequest, opts ...grpc.CallOption) (*UpdateProductByIDResponse, error)
	DeleteProductByID(ctx context.Context, in *DeleteProductByIDRequest, opts ...grpc.CallOption) (*DeleteProductByIDResponse, error)
}

type productGapiClient struct {
	cc grpc.ClientConnInterface
}

func NewProductGapiClient(cc grpc.ClientConnInterface) ProductGapiClient {
	return &productGapiClient{cc}
}

func (c *productGapiClient) InsertProduct(ctx context.Context, in *InsertProductRequest, opts ...grpc.CallOption) (*InsertProductResponse, error) {
	out := new(InsertProductResponse)
	err := c.cc.Invoke(ctx, "/product.ProductGapi/InsertProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGapiClient) GetListProducts(ctx context.Context, in *GetListProductRequest, opts ...grpc.CallOption) (*GetListProductResponse, error) {
	out := new(GetListProductResponse)
	err := c.cc.Invoke(ctx, "/product.ProductGapi/GetListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGapiClient) GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*GetProductByIDResponse, error) {
	out := new(GetProductByIDResponse)
	err := c.cc.Invoke(ctx, "/product.ProductGapi/GetProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGapiClient) UpdateProductByID(ctx context.Context, in *UpdateProductByIDRequest, opts ...grpc.CallOption) (*UpdateProductByIDResponse, error) {
	out := new(UpdateProductByIDResponse)
	err := c.cc.Invoke(ctx, "/product.ProductGapi/UpdateProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGapiClient) DeleteProductByID(ctx context.Context, in *DeleteProductByIDRequest, opts ...grpc.CallOption) (*DeleteProductByIDResponse, error) {
	out := new(DeleteProductByIDResponse)
	err := c.cc.Invoke(ctx, "/product.ProductGapi/DeleteProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductGapiServer is the server API for ProductGapi service.
// All implementations must embed UnimplementedProductGapiServer
// for forward compatibility
type ProductGapiServer interface {
	InsertProduct(context.Context, *InsertProductRequest) (*InsertProductResponse, error)
	GetListProducts(context.Context, *GetListProductRequest) (*GetListProductResponse, error)
	GetProductByID(context.Context, *GetProductByIDRequest) (*GetProductByIDResponse, error)
	UpdateProductByID(context.Context, *UpdateProductByIDRequest) (*UpdateProductByIDResponse, error)
	DeleteProductByID(context.Context, *DeleteProductByIDRequest) (*DeleteProductByIDResponse, error)
	mustEmbedUnimplementedProductGapiServer()
}

// UnimplementedProductGapiServer must be embedded to have forward compatible implementations.
type UnimplementedProductGapiServer struct {
}

func (UnimplementedProductGapiServer) InsertProduct(context.Context, *InsertProductRequest) (*InsertProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertProduct not implemented")
}
func (UnimplementedProductGapiServer) GetListProducts(context.Context, *GetListProductRequest) (*GetListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListProducts not implemented")
}
func (UnimplementedProductGapiServer) GetProductByID(context.Context, *GetProductByIDRequest) (*GetProductByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}
func (UnimplementedProductGapiServer) UpdateProductByID(context.Context, *UpdateProductByIDRequest) (*UpdateProductByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductByID not implemented")
}
func (UnimplementedProductGapiServer) DeleteProductByID(context.Context, *DeleteProductByIDRequest) (*DeleteProductByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductByID not implemented")
}
func (UnimplementedProductGapiServer) mustEmbedUnimplementedProductGapiServer() {}

// UnsafeProductGapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductGapiServer will
// result in compilation errors.
type UnsafeProductGapiServer interface {
	mustEmbedUnimplementedProductGapiServer()
}

func RegisterProductGapiServer(s grpc.ServiceRegistrar, srv ProductGapiServer) {
	s.RegisterService(&ProductGapi_ServiceDesc, srv)
}

func _ProductGapi_InsertProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGapiServer).InsertProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductGapi/InsertProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGapiServer).InsertProduct(ctx, req.(*InsertProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGapi_GetListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGapiServer).GetListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductGapi/GetListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGapiServer).GetListProducts(ctx, req.(*GetListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGapi_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGapiServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductGapi/GetProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGapiServer).GetProductByID(ctx, req.(*GetProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGapi_UpdateProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGapiServer).UpdateProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductGapi/UpdateProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGapiServer).UpdateProductByID(ctx, req.(*UpdateProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGapi_DeleteProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGapiServer).DeleteProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductGapi/DeleteProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGapiServer).DeleteProductByID(ctx, req.(*DeleteProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductGapi_ServiceDesc is the grpc.ServiceDesc for ProductGapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductGapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductGapi",
	HandlerType: (*ProductGapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertProduct",
			Handler:    _ProductGapi_InsertProduct_Handler,
		},
		{
			MethodName: "GetListProducts",
			Handler:    _ProductGapi_GetListProducts_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _ProductGapi_GetProductByID_Handler,
		},
		{
			MethodName: "UpdateProductByID",
			Handler:    _ProductGapi_UpdateProductByID_Handler,
		},
		{
			MethodName: "DeleteProductByID",
			Handler:    _ProductGapi_DeleteProductByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/product/product.proto",
}
