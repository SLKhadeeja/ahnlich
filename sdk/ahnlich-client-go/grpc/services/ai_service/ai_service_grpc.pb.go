// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: services/ai_service.proto

package ai_service

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	pipeline "github.com/deven96/ahnlich/sdk/ahnlich-client-go/grpc/ai/pipeline"
	query "github.com/deven96/ahnlich/sdk/ahnlich-client-go/grpc/ai/query"
	server "github.com/deven96/ahnlich/sdk/ahnlich-client-go/grpc/ai/server"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AIService_CreateStore_FullMethodName                   = "/services.ai_service.AIService/CreateStore"
	AIService_CreatePredIndex_FullMethodName               = "/services.ai_service.AIService/CreatePredIndex"
	AIService_CreateNonLinearAlgorithmIndex_FullMethodName = "/services.ai_service.AIService/CreateNonLinearAlgorithmIndex"
	AIService_GetKey_FullMethodName                        = "/services.ai_service.AIService/GetKey"
	AIService_GetPred_FullMethodName                       = "/services.ai_service.AIService/GetPred"
	AIService_GetSimN_FullMethodName                       = "/services.ai_service.AIService/GetSimN"
	AIService_Set_FullMethodName                           = "/services.ai_service.AIService/Set"
	AIService_DropPredIndex_FullMethodName                 = "/services.ai_service.AIService/DropPredIndex"
	AIService_DropNonLinearAlgorithmIndex_FullMethodName   = "/services.ai_service.AIService/DropNonLinearAlgorithmIndex"
	AIService_DelKey_FullMethodName                        = "/services.ai_service.AIService/DelKey"
	AIService_DropStore_FullMethodName                     = "/services.ai_service.AIService/DropStore"
	AIService_ListClients_FullMethodName                   = "/services.ai_service.AIService/ListClients"
	AIService_ListStores_FullMethodName                    = "/services.ai_service.AIService/ListStores"
	AIService_InfoServer_FullMethodName                    = "/services.ai_service.AIService/InfoServer"
	AIService_PurgeStores_FullMethodName                   = "/services.ai_service.AIService/PurgeStores"
	AIService_Ping_FullMethodName                          = "/services.ai_service.AIService/Ping"
	AIService_Pipeline_FullMethodName                      = "/services.ai_service.AIService/Pipeline"
)

// AIServiceClient is the client API for AIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AIServiceClient interface {
	// * Create methods *
	CreateStore(ctx context.Context, in *query.CreateStore, opts ...grpc.CallOption) (*server.Unit, error)
	CreatePredIndex(ctx context.Context, in *query.CreatePredIndex, opts ...grpc.CallOption) (*server.CreateIndex, error)
	CreateNonLinearAlgorithmIndex(ctx context.Context, in *query.CreateNonLinearAlgorithmIndex, opts ...grpc.CallOption) (*server.CreateIndex, error)
	// * Read methods *
	GetKey(ctx context.Context, in *query.GetKey, opts ...grpc.CallOption) (*server.Get, error)
	GetPred(ctx context.Context, in *query.GetPred, opts ...grpc.CallOption) (*server.Get, error)
	GetSimN(ctx context.Context, in *query.GetSimN, opts ...grpc.CallOption) (*server.GetSimN, error)
	// * Update methods *
	Set(ctx context.Context, in *query.Set, opts ...grpc.CallOption) (*server.Set, error)
	// * Delete methods *
	DropPredIndex(ctx context.Context, in *query.DropPredIndex, opts ...grpc.CallOption) (*server.Del, error)
	DropNonLinearAlgorithmIndex(ctx context.Context, in *query.DropNonLinearAlgorithmIndex, opts ...grpc.CallOption) (*server.Del, error)
	DelKey(ctx context.Context, in *query.DelKey, opts ...grpc.CallOption) (*server.Del, error)
	DropStore(ctx context.Context, in *query.DropStore, opts ...grpc.CallOption) (*server.Del, error)
	// * Ancillary info methods *
	ListClients(ctx context.Context, in *query.ListClients, opts ...grpc.CallOption) (*server.ClientList, error)
	ListStores(ctx context.Context, in *query.ListStores, opts ...grpc.CallOption) (*server.StoreList, error)
	InfoServer(ctx context.Context, in *query.InfoServer, opts ...grpc.CallOption) (*server.InfoServer, error)
	PurgeStores(ctx context.Context, in *query.PurgeStores, opts ...grpc.CallOption) (*server.Del, error)
	Ping(ctx context.Context, in *query.Ping, opts ...grpc.CallOption) (*server.Pong, error)
	// * Pipeline method for all methods *
	Pipeline(ctx context.Context, in *pipeline.AIRequestPipeline, opts ...grpc.CallOption) (*pipeline.AIResponsePipeline, error)
}

type aIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAIServiceClient(cc grpc.ClientConnInterface) AIServiceClient {
	return &aIServiceClient{cc}
}

func (c *aIServiceClient) CreateStore(ctx context.Context, in *query.CreateStore, opts ...grpc.CallOption) (*server.Unit, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Unit)
	err := c.cc.Invoke(ctx, AIService_CreateStore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) CreatePredIndex(ctx context.Context, in *query.CreatePredIndex, opts ...grpc.CallOption) (*server.CreateIndex, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.CreateIndex)
	err := c.cc.Invoke(ctx, AIService_CreatePredIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) CreateNonLinearAlgorithmIndex(ctx context.Context, in *query.CreateNonLinearAlgorithmIndex, opts ...grpc.CallOption) (*server.CreateIndex, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.CreateIndex)
	err := c.cc.Invoke(ctx, AIService_CreateNonLinearAlgorithmIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) GetKey(ctx context.Context, in *query.GetKey, opts ...grpc.CallOption) (*server.Get, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Get)
	err := c.cc.Invoke(ctx, AIService_GetKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) GetPred(ctx context.Context, in *query.GetPred, opts ...grpc.CallOption) (*server.Get, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Get)
	err := c.cc.Invoke(ctx, AIService_GetPred_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) GetSimN(ctx context.Context, in *query.GetSimN, opts ...grpc.CallOption) (*server.GetSimN, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.GetSimN)
	err := c.cc.Invoke(ctx, AIService_GetSimN_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) Set(ctx context.Context, in *query.Set, opts ...grpc.CallOption) (*server.Set, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Set)
	err := c.cc.Invoke(ctx, AIService_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) DropPredIndex(ctx context.Context, in *query.DropPredIndex, opts ...grpc.CallOption) (*server.Del, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Del)
	err := c.cc.Invoke(ctx, AIService_DropPredIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) DropNonLinearAlgorithmIndex(ctx context.Context, in *query.DropNonLinearAlgorithmIndex, opts ...grpc.CallOption) (*server.Del, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Del)
	err := c.cc.Invoke(ctx, AIService_DropNonLinearAlgorithmIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) DelKey(ctx context.Context, in *query.DelKey, opts ...grpc.CallOption) (*server.Del, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Del)
	err := c.cc.Invoke(ctx, AIService_DelKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) DropStore(ctx context.Context, in *query.DropStore, opts ...grpc.CallOption) (*server.Del, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Del)
	err := c.cc.Invoke(ctx, AIService_DropStore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) ListClients(ctx context.Context, in *query.ListClients, opts ...grpc.CallOption) (*server.ClientList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.ClientList)
	err := c.cc.Invoke(ctx, AIService_ListClients_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) ListStores(ctx context.Context, in *query.ListStores, opts ...grpc.CallOption) (*server.StoreList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.StoreList)
	err := c.cc.Invoke(ctx, AIService_ListStores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) InfoServer(ctx context.Context, in *query.InfoServer, opts ...grpc.CallOption) (*server.InfoServer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.InfoServer)
	err := c.cc.Invoke(ctx, AIService_InfoServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) PurgeStores(ctx context.Context, in *query.PurgeStores, opts ...grpc.CallOption) (*server.Del, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Del)
	err := c.cc.Invoke(ctx, AIService_PurgeStores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) Ping(ctx context.Context, in *query.Ping, opts ...grpc.CallOption) (*server.Pong, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(server.Pong)
	err := c.cc.Invoke(ctx, AIService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIServiceClient) Pipeline(ctx context.Context, in *pipeline.AIRequestPipeline, opts ...grpc.CallOption) (*pipeline.AIResponsePipeline, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pipeline.AIResponsePipeline)
	err := c.cc.Invoke(ctx, AIService_Pipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AIServiceServer is the server API for AIService service.
// All implementations must embed UnimplementedAIServiceServer
// for forward compatibility.
type AIServiceServer interface {
	// * Create methods *
	CreateStore(context.Context, *query.CreateStore) (*server.Unit, error)
	CreatePredIndex(context.Context, *query.CreatePredIndex) (*server.CreateIndex, error)
	CreateNonLinearAlgorithmIndex(context.Context, *query.CreateNonLinearAlgorithmIndex) (*server.CreateIndex, error)
	// * Read methods *
	GetKey(context.Context, *query.GetKey) (*server.Get, error)
	GetPred(context.Context, *query.GetPred) (*server.Get, error)
	GetSimN(context.Context, *query.GetSimN) (*server.GetSimN, error)
	// * Update methods *
	Set(context.Context, *query.Set) (*server.Set, error)
	// * Delete methods *
	DropPredIndex(context.Context, *query.DropPredIndex) (*server.Del, error)
	DropNonLinearAlgorithmIndex(context.Context, *query.DropNonLinearAlgorithmIndex) (*server.Del, error)
	DelKey(context.Context, *query.DelKey) (*server.Del, error)
	DropStore(context.Context, *query.DropStore) (*server.Del, error)
	// * Ancillary info methods *
	ListClients(context.Context, *query.ListClients) (*server.ClientList, error)
	ListStores(context.Context, *query.ListStores) (*server.StoreList, error)
	InfoServer(context.Context, *query.InfoServer) (*server.InfoServer, error)
	PurgeStores(context.Context, *query.PurgeStores) (*server.Del, error)
	Ping(context.Context, *query.Ping) (*server.Pong, error)
	// * Pipeline method for all methods *
	Pipeline(context.Context, *pipeline.AIRequestPipeline) (*pipeline.AIResponsePipeline, error)
	mustEmbedUnimplementedAIServiceServer()
}

// UnimplementedAIServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAIServiceServer struct{}

func (UnimplementedAIServiceServer) CreateStore(context.Context, *query.CreateStore) (*server.Unit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStore not implemented")
}
func (UnimplementedAIServiceServer) CreatePredIndex(context.Context, *query.CreatePredIndex) (*server.CreateIndex, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePredIndex not implemented")
}
func (UnimplementedAIServiceServer) CreateNonLinearAlgorithmIndex(context.Context, *query.CreateNonLinearAlgorithmIndex) (*server.CreateIndex, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNonLinearAlgorithmIndex not implemented")
}
func (UnimplementedAIServiceServer) GetKey(context.Context, *query.GetKey) (*server.Get, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}
func (UnimplementedAIServiceServer) GetPred(context.Context, *query.GetPred) (*server.Get, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPred not implemented")
}
func (UnimplementedAIServiceServer) GetSimN(context.Context, *query.GetSimN) (*server.GetSimN, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimN not implemented")
}
func (UnimplementedAIServiceServer) Set(context.Context, *query.Set) (*server.Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedAIServiceServer) DropPredIndex(context.Context, *query.DropPredIndex) (*server.Del, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropPredIndex not implemented")
}
func (UnimplementedAIServiceServer) DropNonLinearAlgorithmIndex(context.Context, *query.DropNonLinearAlgorithmIndex) (*server.Del, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropNonLinearAlgorithmIndex not implemented")
}
func (UnimplementedAIServiceServer) DelKey(context.Context, *query.DelKey) (*server.Del, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelKey not implemented")
}
func (UnimplementedAIServiceServer) DropStore(context.Context, *query.DropStore) (*server.Del, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropStore not implemented")
}
func (UnimplementedAIServiceServer) ListClients(context.Context, *query.ListClients) (*server.ClientList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClients not implemented")
}
func (UnimplementedAIServiceServer) ListStores(context.Context, *query.ListStores) (*server.StoreList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStores not implemented")
}
func (UnimplementedAIServiceServer) InfoServer(context.Context, *query.InfoServer) (*server.InfoServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoServer not implemented")
}
func (UnimplementedAIServiceServer) PurgeStores(context.Context, *query.PurgeStores) (*server.Del, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeStores not implemented")
}
func (UnimplementedAIServiceServer) Ping(context.Context, *query.Ping) (*server.Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAIServiceServer) Pipeline(context.Context, *pipeline.AIRequestPipeline) (*pipeline.AIResponsePipeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pipeline not implemented")
}
func (UnimplementedAIServiceServer) mustEmbedUnimplementedAIServiceServer() {}
func (UnimplementedAIServiceServer) testEmbeddedByValue()                   {}

// UnsafeAIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AIServiceServer will
// result in compilation errors.
type UnsafeAIServiceServer interface {
	mustEmbedUnimplementedAIServiceServer()
}

func RegisterAIServiceServer(s grpc.ServiceRegistrar, srv AIServiceServer) {
	// If the following call pancis, it indicates UnimplementedAIServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AIService_ServiceDesc, srv)
}

func _AIService_CreateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.CreateStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).CreateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_CreateStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).CreateStore(ctx, req.(*query.CreateStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_CreatePredIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.CreatePredIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).CreatePredIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_CreatePredIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).CreatePredIndex(ctx, req.(*query.CreatePredIndex))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_CreateNonLinearAlgorithmIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.CreateNonLinearAlgorithmIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).CreateNonLinearAlgorithmIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_CreateNonLinearAlgorithmIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).CreateNonLinearAlgorithmIndex(ctx, req.(*query.CreateNonLinearAlgorithmIndex))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.GetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_GetKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).GetKey(ctx, req.(*query.GetKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_GetPred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.GetPred)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).GetPred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_GetPred_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).GetPred(ctx, req.(*query.GetPred))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_GetSimN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.GetSimN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).GetSimN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_GetSimN_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).GetSimN(ctx, req.(*query.GetSimN))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.Set)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).Set(ctx, req.(*query.Set))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_DropPredIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.DropPredIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).DropPredIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_DropPredIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).DropPredIndex(ctx, req.(*query.DropPredIndex))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_DropNonLinearAlgorithmIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.DropNonLinearAlgorithmIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).DropNonLinearAlgorithmIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_DropNonLinearAlgorithmIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).DropNonLinearAlgorithmIndex(ctx, req.(*query.DropNonLinearAlgorithmIndex))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_DelKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.DelKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).DelKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_DelKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).DelKey(ctx, req.(*query.DelKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_DropStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.DropStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).DropStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_DropStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).DropStore(ctx, req.(*query.DropStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_ListClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.ListClients)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).ListClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_ListClients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).ListClients(ctx, req.(*query.ListClients))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_ListStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.ListStores)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).ListStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_ListStores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).ListStores(ctx, req.(*query.ListStores))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_InfoServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.InfoServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).InfoServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_InfoServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).InfoServer(ctx, req.(*query.InfoServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_PurgeStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.PurgeStores)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).PurgeStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_PurgeStores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).PurgeStores(ctx, req.(*query.PurgeStores))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).Ping(ctx, req.(*query.Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIService_Pipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pipeline.AIRequestPipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).Pipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_Pipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).Pipeline(ctx, req.(*pipeline.AIRequestPipeline))
	}
	return interceptor(ctx, in, info, handler)
}

// AIService_ServiceDesc is the grpc.ServiceDesc for AIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.ai_service.AIService",
	HandlerType: (*AIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStore",
			Handler:    _AIService_CreateStore_Handler,
		},
		{
			MethodName: "CreatePredIndex",
			Handler:    _AIService_CreatePredIndex_Handler,
		},
		{
			MethodName: "CreateNonLinearAlgorithmIndex",
			Handler:    _AIService_CreateNonLinearAlgorithmIndex_Handler,
		},
		{
			MethodName: "GetKey",
			Handler:    _AIService_GetKey_Handler,
		},
		{
			MethodName: "GetPred",
			Handler:    _AIService_GetPred_Handler,
		},
		{
			MethodName: "GetSimN",
			Handler:    _AIService_GetSimN_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _AIService_Set_Handler,
		},
		{
			MethodName: "DropPredIndex",
			Handler:    _AIService_DropPredIndex_Handler,
		},
		{
			MethodName: "DropNonLinearAlgorithmIndex",
			Handler:    _AIService_DropNonLinearAlgorithmIndex_Handler,
		},
		{
			MethodName: "DelKey",
			Handler:    _AIService_DelKey_Handler,
		},
		{
			MethodName: "DropStore",
			Handler:    _AIService_DropStore_Handler,
		},
		{
			MethodName: "ListClients",
			Handler:    _AIService_ListClients_Handler,
		},
		{
			MethodName: "ListStores",
			Handler:    _AIService_ListStores_Handler,
		},
		{
			MethodName: "InfoServer",
			Handler:    _AIService_InfoServer_Handler,
		},
		{
			MethodName: "PurgeStores",
			Handler:    _AIService_PurgeStores_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _AIService_Ping_Handler,
		},
		{
			MethodName: "Pipeline",
			Handler:    _AIService_Pipeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/ai_service.proto",
}
