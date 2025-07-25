// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: ai/pipeline.proto

package pipeline

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	query "github.com/deven96/ahnlich/sdk/ahnlich-client-go/grpc/ai/query"
	server "github.com/deven96/ahnlich/sdk/ahnlich-client-go/grpc/ai/server"
	info "github.com/deven96/ahnlich/sdk/ahnlich-client-go/grpc/shared/info"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AIQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//
	//	*AIQuery_CreateStore
	//	*AIQuery_GetPred
	//	*AIQuery_GetSimN
	//	*AIQuery_CreatePredIndex
	//	*AIQuery_CreateNonLinearAlgorithmIndex
	//	*AIQuery_DropPredIndex
	//	*AIQuery_DropNonLinearAlgorithmIndex
	//	*AIQuery_Set
	//	*AIQuery_DelKey
	//	*AIQuery_DropStore
	//	*AIQuery_GetKey
	//	*AIQuery_InfoServer
	//	*AIQuery_ListClients
	//	*AIQuery_ListStores
	//	*AIQuery_PurgeStores
	//	*AIQuery_Ping
	Query isAIQuery_Query `protobuf_oneof:"query"`
}

func (x *AIQuery) Reset() {
	*x = AIQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIQuery) ProtoMessage() {}

func (x *AIQuery) ProtoReflect() protoreflect.Message {
	mi := &file_ai_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIQuery.ProtoReflect.Descriptor instead.
func (*AIQuery) Descriptor() ([]byte, []int) {
	return file_ai_pipeline_proto_rawDescGZIP(), []int{0}
}

func (m *AIQuery) GetQuery() isAIQuery_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *AIQuery) GetCreateStore() *query.CreateStore {
	if x, ok := x.GetQuery().(*AIQuery_CreateStore); ok {
		return x.CreateStore
	}
	return nil
}

func (x *AIQuery) GetGetPred() *query.GetPred {
	if x, ok := x.GetQuery().(*AIQuery_GetPred); ok {
		return x.GetPred
	}
	return nil
}

func (x *AIQuery) GetGetSimN() *query.GetSimN {
	if x, ok := x.GetQuery().(*AIQuery_GetSimN); ok {
		return x.GetSimN
	}
	return nil
}

func (x *AIQuery) GetCreatePredIndex() *query.CreatePredIndex {
	if x, ok := x.GetQuery().(*AIQuery_CreatePredIndex); ok {
		return x.CreatePredIndex
	}
	return nil
}

func (x *AIQuery) GetCreateNonLinearAlgorithmIndex() *query.CreateNonLinearAlgorithmIndex {
	if x, ok := x.GetQuery().(*AIQuery_CreateNonLinearAlgorithmIndex); ok {
		return x.CreateNonLinearAlgorithmIndex
	}
	return nil
}

func (x *AIQuery) GetDropPredIndex() *query.DropPredIndex {
	if x, ok := x.GetQuery().(*AIQuery_DropPredIndex); ok {
		return x.DropPredIndex
	}
	return nil
}

func (x *AIQuery) GetDropNonLinearAlgorithmIndex() *query.DropNonLinearAlgorithmIndex {
	if x, ok := x.GetQuery().(*AIQuery_DropNonLinearAlgorithmIndex); ok {
		return x.DropNonLinearAlgorithmIndex
	}
	return nil
}

func (x *AIQuery) GetSet() *query.Set {
	if x, ok := x.GetQuery().(*AIQuery_Set); ok {
		return x.Set
	}
	return nil
}

func (x *AIQuery) GetDelKey() *query.DelKey {
	if x, ok := x.GetQuery().(*AIQuery_DelKey); ok {
		return x.DelKey
	}
	return nil
}

func (x *AIQuery) GetDropStore() *query.DropStore {
	if x, ok := x.GetQuery().(*AIQuery_DropStore); ok {
		return x.DropStore
	}
	return nil
}

func (x *AIQuery) GetGetKey() *query.GetKey {
	if x, ok := x.GetQuery().(*AIQuery_GetKey); ok {
		return x.GetKey
	}
	return nil
}

func (x *AIQuery) GetInfoServer() *query.InfoServer {
	if x, ok := x.GetQuery().(*AIQuery_InfoServer); ok {
		return x.InfoServer
	}
	return nil
}

func (x *AIQuery) GetListClients() *query.ListClients {
	if x, ok := x.GetQuery().(*AIQuery_ListClients); ok {
		return x.ListClients
	}
	return nil
}

func (x *AIQuery) GetListStores() *query.ListStores {
	if x, ok := x.GetQuery().(*AIQuery_ListStores); ok {
		return x.ListStores
	}
	return nil
}

func (x *AIQuery) GetPurgeStores() *query.PurgeStores {
	if x, ok := x.GetQuery().(*AIQuery_PurgeStores); ok {
		return x.PurgeStores
	}
	return nil
}

func (x *AIQuery) GetPing() *query.Ping {
	if x, ok := x.GetQuery().(*AIQuery_Ping); ok {
		return x.Ping
	}
	return nil
}

type isAIQuery_Query interface {
	isAIQuery_Query()
}

type AIQuery_CreateStore struct {
	CreateStore *query.CreateStore `protobuf:"bytes,1,opt,name=create_store,json=createStore,proto3,oneof"`
}

type AIQuery_GetPred struct {
	GetPred *query.GetPred `protobuf:"bytes,2,opt,name=get_pred,json=getPred,proto3,oneof"`
}

type AIQuery_GetSimN struct {
	GetSimN *query.GetSimN `protobuf:"bytes,3,opt,name=get_sim_n,json=getSimN,proto3,oneof"`
}

type AIQuery_CreatePredIndex struct {
	CreatePredIndex *query.CreatePredIndex `protobuf:"bytes,4,opt,name=create_pred_index,json=createPredIndex,proto3,oneof"`
}

type AIQuery_CreateNonLinearAlgorithmIndex struct {
	CreateNonLinearAlgorithmIndex *query.CreateNonLinearAlgorithmIndex `protobuf:"bytes,5,opt,name=create_non_linear_algorithm_index,json=createNonLinearAlgorithmIndex,proto3,oneof"`
}

type AIQuery_DropPredIndex struct {
	DropPredIndex *query.DropPredIndex `protobuf:"bytes,6,opt,name=drop_pred_index,json=dropPredIndex,proto3,oneof"`
}

type AIQuery_DropNonLinearAlgorithmIndex struct {
	DropNonLinearAlgorithmIndex *query.DropNonLinearAlgorithmIndex `protobuf:"bytes,7,opt,name=drop_non_linear_algorithm_index,json=dropNonLinearAlgorithmIndex,proto3,oneof"`
}

type AIQuery_Set struct {
	Set *query.Set `protobuf:"bytes,8,opt,name=set,proto3,oneof"`
}

type AIQuery_DelKey struct {
	DelKey *query.DelKey `protobuf:"bytes,9,opt,name=del_key,json=delKey,proto3,oneof"`
}

type AIQuery_DropStore struct {
	DropStore *query.DropStore `protobuf:"bytes,10,opt,name=drop_store,json=dropStore,proto3,oneof"`
}

type AIQuery_GetKey struct {
	GetKey *query.GetKey `protobuf:"bytes,11,opt,name=get_key,json=getKey,proto3,oneof"`
}

type AIQuery_InfoServer struct {
	InfoServer *query.InfoServer `protobuf:"bytes,12,opt,name=info_server,json=infoServer,proto3,oneof"`
}

type AIQuery_ListClients struct {
	ListClients *query.ListClients `protobuf:"bytes,13,opt,name=list_clients,json=listClients,proto3,oneof"`
}

type AIQuery_ListStores struct {
	ListStores *query.ListStores `protobuf:"bytes,14,opt,name=list_stores,json=listStores,proto3,oneof"`
}

type AIQuery_PurgeStores struct {
	PurgeStores *query.PurgeStores `protobuf:"bytes,15,opt,name=purge_stores,json=purgeStores,proto3,oneof"`
}

type AIQuery_Ping struct {
	Ping *query.Ping `protobuf:"bytes,16,opt,name=ping,proto3,oneof"`
}

func (*AIQuery_CreateStore) isAIQuery_Query() {}

func (*AIQuery_GetPred) isAIQuery_Query() {}

func (*AIQuery_GetSimN) isAIQuery_Query() {}

func (*AIQuery_CreatePredIndex) isAIQuery_Query() {}

func (*AIQuery_CreateNonLinearAlgorithmIndex) isAIQuery_Query() {}

func (*AIQuery_DropPredIndex) isAIQuery_Query() {}

func (*AIQuery_DropNonLinearAlgorithmIndex) isAIQuery_Query() {}

func (*AIQuery_Set) isAIQuery_Query() {}

func (*AIQuery_DelKey) isAIQuery_Query() {}

func (*AIQuery_DropStore) isAIQuery_Query() {}

func (*AIQuery_GetKey) isAIQuery_Query() {}

func (*AIQuery_InfoServer) isAIQuery_Query() {}

func (*AIQuery_ListClients) isAIQuery_Query() {}

func (*AIQuery_ListStores) isAIQuery_Query() {}

func (*AIQuery_PurgeStores) isAIQuery_Query() {}

func (*AIQuery_Ping) isAIQuery_Query() {}

type AIRequestPipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queries []*AIQuery `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
}

func (x *AIRequestPipeline) Reset() {
	*x = AIRequestPipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIRequestPipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIRequestPipeline) ProtoMessage() {}

func (x *AIRequestPipeline) ProtoReflect() protoreflect.Message {
	mi := &file_ai_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIRequestPipeline.ProtoReflect.Descriptor instead.
func (*AIRequestPipeline) Descriptor() ([]byte, []int) {
	return file_ai_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *AIRequestPipeline) GetQueries() []*AIQuery {
	if x != nil {
		return x.Queries
	}
	return nil
}

type AIServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*AIServerResponse_Unit
	//	*AIServerResponse_Pong
	//	*AIServerResponse_ClientList
	//	*AIServerResponse_StoreList
	//	*AIServerResponse_InfoServer
	//	*AIServerResponse_Set
	//	*AIServerResponse_Get
	//	*AIServerResponse_GetSimN
	//	*AIServerResponse_Del
	//	*AIServerResponse_CreateIndex
	//	*AIServerResponse_Error
	Response isAIServerResponse_Response `protobuf_oneof:"response"`
}

func (x *AIServerResponse) Reset() {
	*x = AIServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIServerResponse) ProtoMessage() {}

func (x *AIServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ai_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIServerResponse.ProtoReflect.Descriptor instead.
func (*AIServerResponse) Descriptor() ([]byte, []int) {
	return file_ai_pipeline_proto_rawDescGZIP(), []int{2}
}

func (m *AIServerResponse) GetResponse() isAIServerResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *AIServerResponse) GetUnit() *server.Unit {
	if x, ok := x.GetResponse().(*AIServerResponse_Unit); ok {
		return x.Unit
	}
	return nil
}

func (x *AIServerResponse) GetPong() *server.Pong {
	if x, ok := x.GetResponse().(*AIServerResponse_Pong); ok {
		return x.Pong
	}
	return nil
}

func (x *AIServerResponse) GetClientList() *server.ClientList {
	if x, ok := x.GetResponse().(*AIServerResponse_ClientList); ok {
		return x.ClientList
	}
	return nil
}

func (x *AIServerResponse) GetStoreList() *server.StoreList {
	if x, ok := x.GetResponse().(*AIServerResponse_StoreList); ok {
		return x.StoreList
	}
	return nil
}

func (x *AIServerResponse) GetInfoServer() *server.InfoServer {
	if x, ok := x.GetResponse().(*AIServerResponse_InfoServer); ok {
		return x.InfoServer
	}
	return nil
}

func (x *AIServerResponse) GetSet() *server.Set {
	if x, ok := x.GetResponse().(*AIServerResponse_Set); ok {
		return x.Set
	}
	return nil
}

func (x *AIServerResponse) GetGet() *server.Get {
	if x, ok := x.GetResponse().(*AIServerResponse_Get); ok {
		return x.Get
	}
	return nil
}

func (x *AIServerResponse) GetGetSimN() *server.GetSimN {
	if x, ok := x.GetResponse().(*AIServerResponse_GetSimN); ok {
		return x.GetSimN
	}
	return nil
}

func (x *AIServerResponse) GetDel() *server.Del {
	if x, ok := x.GetResponse().(*AIServerResponse_Del); ok {
		return x.Del
	}
	return nil
}

func (x *AIServerResponse) GetCreateIndex() *server.CreateIndex {
	if x, ok := x.GetResponse().(*AIServerResponse_CreateIndex); ok {
		return x.CreateIndex
	}
	return nil
}

func (x *AIServerResponse) GetError() *info.ErrorResponse {
	if x, ok := x.GetResponse().(*AIServerResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isAIServerResponse_Response interface {
	isAIServerResponse_Response()
}

type AIServerResponse_Unit struct {
	Unit *server.Unit `protobuf:"bytes,1,opt,name=unit,proto3,oneof"`
}

type AIServerResponse_Pong struct {
	Pong *server.Pong `protobuf:"bytes,2,opt,name=pong,proto3,oneof"`
}

type AIServerResponse_ClientList struct {
	ClientList *server.ClientList `protobuf:"bytes,3,opt,name=client_list,json=clientList,proto3,oneof"`
}

type AIServerResponse_StoreList struct {
	StoreList *server.StoreList `protobuf:"bytes,4,opt,name=store_list,json=storeList,proto3,oneof"`
}

type AIServerResponse_InfoServer struct {
	InfoServer *server.InfoServer `protobuf:"bytes,5,opt,name=info_server,json=infoServer,proto3,oneof"`
}

type AIServerResponse_Set struct {
	Set *server.Set `protobuf:"bytes,6,opt,name=set,proto3,oneof"`
}

type AIServerResponse_Get struct {
	Get *server.Get `protobuf:"bytes,7,opt,name=get,proto3,oneof"`
}

type AIServerResponse_GetSimN struct {
	GetSimN *server.GetSimN `protobuf:"bytes,8,opt,name=get_sim_n,json=getSimN,proto3,oneof"`
}

type AIServerResponse_Del struct {
	Del *server.Del `protobuf:"bytes,9,opt,name=del,proto3,oneof"`
}

type AIServerResponse_CreateIndex struct {
	CreateIndex *server.CreateIndex `protobuf:"bytes,10,opt,name=create_index,json=createIndex,proto3,oneof"`
}

type AIServerResponse_Error struct {
	Error *info.ErrorResponse `protobuf:"bytes,11,opt,name=error,proto3,oneof"`
}

func (*AIServerResponse_Unit) isAIServerResponse_Response() {}

func (*AIServerResponse_Pong) isAIServerResponse_Response() {}

func (*AIServerResponse_ClientList) isAIServerResponse_Response() {}

func (*AIServerResponse_StoreList) isAIServerResponse_Response() {}

func (*AIServerResponse_InfoServer) isAIServerResponse_Response() {}

func (*AIServerResponse_Set) isAIServerResponse_Response() {}

func (*AIServerResponse_Get) isAIServerResponse_Response() {}

func (*AIServerResponse_GetSimN) isAIServerResponse_Response() {}

func (*AIServerResponse_Del) isAIServerResponse_Response() {}

func (*AIServerResponse_CreateIndex) isAIServerResponse_Response() {}

func (*AIServerResponse_Error) isAIServerResponse_Response() {}

type AIResponsePipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*AIServerResponse `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *AIResponsePipeline) Reset() {
	*x = AIResponsePipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_pipeline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIResponsePipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIResponsePipeline) ProtoMessage() {}

func (x *AIResponsePipeline) ProtoReflect() protoreflect.Message {
	mi := &file_ai_pipeline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIResponsePipeline.ProtoReflect.Descriptor instead.
func (*AIResponsePipeline) Descriptor() ([]byte, []int) {
	return file_ai_pipeline_proto_rawDescGZIP(), []int{3}
}

func (x *AIResponsePipeline) GetResponses() []*AIServerResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

var File_ai_pipeline_proto protoreflect.FileDescriptor

var file_ai_pipeline_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x69, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x69, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x1a, 0x0e, 0x61, 0x69, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x61, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x07, 0x0a, 0x07, 0x41, 0x49, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x00, 0x52,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x08,
	0x67, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x07, 0x67, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x09,
	0x67, 0x65, 0x74, 0x5f, 0x73, 0x69, 0x6d, 0x5f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x6d, 0x4e, 0x48, 0x00, 0x52, 0x07, 0x67, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x4e, 0x12, 0x47, 0x0a,
	0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65,
	0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x73, 0x0a, 0x21, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x6e, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x5f, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x48, 0x00, 0x52, 0x1d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x41, 0x0a, 0x0f, 0x64,
	0x72, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x44, 0x72, 0x6f, 0x70, 0x50, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x48, 0x00, 0x52,
	0x0d, 0x64, 0x72, 0x6f, 0x70, 0x50, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x6d,
	0x0a, 0x1f, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x6e, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x61,
	0x72, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x4e, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72,
	0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x48, 0x00,
	0x52, 0x1b, 0x64, 0x72, 0x6f, 0x70, 0x4e, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a,
	0x03, 0x73, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x69, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x03, 0x73, 0x65, 0x74,
	0x12, 0x2b, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c,
	0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a,
	0x0a, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x72, 0x6f,
	0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x00, 0x52, 0x09, 0x64, 0x72, 0x6f, 0x70, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x06, 0x67, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x37, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0a, 0x69,
	0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0c, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x69, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x48, 0x00, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x0c, 0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x50, 0x75, 0x72, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x70,
	0x75, 0x72, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x69, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x42, 0x07, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x43, 0x0a, 0x11, 0x41, 0x49, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2e,
	0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x69, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x49,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0xa6,
	0x04, 0x0a, 0x10, 0x41, 0x49, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x6e,
	0x69, 0x74, 0x48, 0x00, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x6f,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x0a, 0x69, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x03,
	0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x03, 0x73, 0x65, 0x74,
	0x12, 0x22, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x61, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x00, 0x52,
	0x03, 0x67, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x69, 0x6d, 0x5f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x4e, 0x48, 0x00, 0x52, 0x07, 0x67,
	0x65, 0x74, 0x53, 0x69, 0x6d, 0x4e, 0x12, 0x22, 0x0a, 0x03, 0x64, 0x65, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x03, 0x64, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x12, 0x41, 0x49, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3b, 0x0a,
	0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x61, 0x69, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41,
	0x49, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x39, 0x36,
	0x2f, 0x61, 0x68, 0x6e, 0x6c, 0x69, 0x63, 0x68, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x68, 0x6e,
	0x6c, 0x69, 0x63, 0x68, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x61, 0x69, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x3b,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ai_pipeline_proto_rawDescOnce sync.Once
	file_ai_pipeline_proto_rawDescData = file_ai_pipeline_proto_rawDesc
)

func file_ai_pipeline_proto_rawDescGZIP() []byte {
	file_ai_pipeline_proto_rawDescOnce.Do(func() {
		file_ai_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_ai_pipeline_proto_rawDescData)
	})
	return file_ai_pipeline_proto_rawDescData
}

var file_ai_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ai_pipeline_proto_goTypes = []any{
	(*AIQuery)(nil),                             // 0: ai.pipeline.AIQuery
	(*AIRequestPipeline)(nil),                   // 1: ai.pipeline.AIRequestPipeline
	(*AIServerResponse)(nil),                    // 2: ai.pipeline.AIServerResponse
	(*AIResponsePipeline)(nil),                  // 3: ai.pipeline.AIResponsePipeline
	(*query.CreateStore)(nil),                   // 4: ai.query.CreateStore
	(*query.GetPred)(nil),                       // 5: ai.query.GetPred
	(*query.GetSimN)(nil),                       // 6: ai.query.GetSimN
	(*query.CreatePredIndex)(nil),               // 7: ai.query.CreatePredIndex
	(*query.CreateNonLinearAlgorithmIndex)(nil), // 8: ai.query.CreateNonLinearAlgorithmIndex
	(*query.DropPredIndex)(nil),                 // 9: ai.query.DropPredIndex
	(*query.DropNonLinearAlgorithmIndex)(nil),   // 10: ai.query.DropNonLinearAlgorithmIndex
	(*query.Set)(nil),                           // 11: ai.query.Set
	(*query.DelKey)(nil),                        // 12: ai.query.DelKey
	(*query.DropStore)(nil),                     // 13: ai.query.DropStore
	(*query.GetKey)(nil),                        // 14: ai.query.GetKey
	(*query.InfoServer)(nil),                    // 15: ai.query.InfoServer
	(*query.ListClients)(nil),                   // 16: ai.query.ListClients
	(*query.ListStores)(nil),                    // 17: ai.query.ListStores
	(*query.PurgeStores)(nil),                   // 18: ai.query.PurgeStores
	(*query.Ping)(nil),                          // 19: ai.query.Ping
	(*server.Unit)(nil),                         // 20: ai.server.Unit
	(*server.Pong)(nil),                         // 21: ai.server.Pong
	(*server.ClientList)(nil),                   // 22: ai.server.ClientList
	(*server.StoreList)(nil),                    // 23: ai.server.StoreList
	(*server.InfoServer)(nil),                   // 24: ai.server.InfoServer
	(*server.Set)(nil),                          // 25: ai.server.Set
	(*server.Get)(nil),                          // 26: ai.server.Get
	(*server.GetSimN)(nil),                      // 27: ai.server.GetSimN
	(*server.Del)(nil),                          // 28: ai.server.Del
	(*server.CreateIndex)(nil),                  // 29: ai.server.CreateIndex
	(*info.ErrorResponse)(nil),                  // 30: shared.info.ErrorResponse
}
var file_ai_pipeline_proto_depIdxs = []int32{
	4,  // 0: ai.pipeline.AIQuery.create_store:type_name -> ai.query.CreateStore
	5,  // 1: ai.pipeline.AIQuery.get_pred:type_name -> ai.query.GetPred
	6,  // 2: ai.pipeline.AIQuery.get_sim_n:type_name -> ai.query.GetSimN
	7,  // 3: ai.pipeline.AIQuery.create_pred_index:type_name -> ai.query.CreatePredIndex
	8,  // 4: ai.pipeline.AIQuery.create_non_linear_algorithm_index:type_name -> ai.query.CreateNonLinearAlgorithmIndex
	9,  // 5: ai.pipeline.AIQuery.drop_pred_index:type_name -> ai.query.DropPredIndex
	10, // 6: ai.pipeline.AIQuery.drop_non_linear_algorithm_index:type_name -> ai.query.DropNonLinearAlgorithmIndex
	11, // 7: ai.pipeline.AIQuery.set:type_name -> ai.query.Set
	12, // 8: ai.pipeline.AIQuery.del_key:type_name -> ai.query.DelKey
	13, // 9: ai.pipeline.AIQuery.drop_store:type_name -> ai.query.DropStore
	14, // 10: ai.pipeline.AIQuery.get_key:type_name -> ai.query.GetKey
	15, // 11: ai.pipeline.AIQuery.info_server:type_name -> ai.query.InfoServer
	16, // 12: ai.pipeline.AIQuery.list_clients:type_name -> ai.query.ListClients
	17, // 13: ai.pipeline.AIQuery.list_stores:type_name -> ai.query.ListStores
	18, // 14: ai.pipeline.AIQuery.purge_stores:type_name -> ai.query.PurgeStores
	19, // 15: ai.pipeline.AIQuery.ping:type_name -> ai.query.Ping
	0,  // 16: ai.pipeline.AIRequestPipeline.queries:type_name -> ai.pipeline.AIQuery
	20, // 17: ai.pipeline.AIServerResponse.unit:type_name -> ai.server.Unit
	21, // 18: ai.pipeline.AIServerResponse.pong:type_name -> ai.server.Pong
	22, // 19: ai.pipeline.AIServerResponse.client_list:type_name -> ai.server.ClientList
	23, // 20: ai.pipeline.AIServerResponse.store_list:type_name -> ai.server.StoreList
	24, // 21: ai.pipeline.AIServerResponse.info_server:type_name -> ai.server.InfoServer
	25, // 22: ai.pipeline.AIServerResponse.set:type_name -> ai.server.Set
	26, // 23: ai.pipeline.AIServerResponse.get:type_name -> ai.server.Get
	27, // 24: ai.pipeline.AIServerResponse.get_sim_n:type_name -> ai.server.GetSimN
	28, // 25: ai.pipeline.AIServerResponse.del:type_name -> ai.server.Del
	29, // 26: ai.pipeline.AIServerResponse.create_index:type_name -> ai.server.CreateIndex
	30, // 27: ai.pipeline.AIServerResponse.error:type_name -> shared.info.ErrorResponse
	2,  // 28: ai.pipeline.AIResponsePipeline.responses:type_name -> ai.pipeline.AIServerResponse
	29, // [29:29] is the sub-list for method output_type
	29, // [29:29] is the sub-list for method input_type
	29, // [29:29] is the sub-list for extension type_name
	29, // [29:29] is the sub-list for extension extendee
	0,  // [0:29] is the sub-list for field type_name
}

func init() { file_ai_pipeline_proto_init() }
func file_ai_pipeline_proto_init() {
	if File_ai_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ai_pipeline_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AIQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ai_pipeline_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AIRequestPipeline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ai_pipeline_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AIServerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ai_pipeline_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AIResponsePipeline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_ai_pipeline_proto_msgTypes[0].OneofWrappers = []any{
		(*AIQuery_CreateStore)(nil),
		(*AIQuery_GetPred)(nil),
		(*AIQuery_GetSimN)(nil),
		(*AIQuery_CreatePredIndex)(nil),
		(*AIQuery_CreateNonLinearAlgorithmIndex)(nil),
		(*AIQuery_DropPredIndex)(nil),
		(*AIQuery_DropNonLinearAlgorithmIndex)(nil),
		(*AIQuery_Set)(nil),
		(*AIQuery_DelKey)(nil),
		(*AIQuery_DropStore)(nil),
		(*AIQuery_GetKey)(nil),
		(*AIQuery_InfoServer)(nil),
		(*AIQuery_ListClients)(nil),
		(*AIQuery_ListStores)(nil),
		(*AIQuery_PurgeStores)(nil),
		(*AIQuery_Ping)(nil),
	}
	file_ai_pipeline_proto_msgTypes[2].OneofWrappers = []any{
		(*AIServerResponse_Unit)(nil),
		(*AIServerResponse_Pong)(nil),
		(*AIServerResponse_ClientList)(nil),
		(*AIServerResponse_StoreList)(nil),
		(*AIServerResponse_InfoServer)(nil),
		(*AIServerResponse_Set)(nil),
		(*AIServerResponse_Get)(nil),
		(*AIServerResponse_GetSimN)(nil),
		(*AIServerResponse_Del)(nil),
		(*AIServerResponse_CreateIndex)(nil),
		(*AIServerResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ai_pipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ai_pipeline_proto_goTypes,
		DependencyIndexes: file_ai_pipeline_proto_depIdxs,
		MessageInfos:      file_ai_pipeline_proto_msgTypes,
	}.Build()
	File_ai_pipeline_proto = out.File
	file_ai_pipeline_proto_rawDesc = nil
	file_ai_pipeline_proto_goTypes = nil
	file_ai_pipeline_proto_depIdxs = nil
}
