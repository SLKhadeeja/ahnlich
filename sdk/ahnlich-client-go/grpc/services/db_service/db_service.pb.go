// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: services/db_service.proto

package db_service

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	pipeline "github.com/deven96/ahnlich/sdk/ahnlich-client-go/grpc/db/pipeline"
	query "github.com/deven96/ahnlich/sdk/ahnlich-client-go/grpc/db/query"
	server "github.com/deven96/ahnlich/sdk/ahnlich-client-go/grpc/db/server"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_db_service_proto protoreflect.FileDescriptor

var file_services_db_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x11, 0x64, 0x62, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x64, 0x62, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x64, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe0, 0x07, 0x0a, 0x09, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x15, 0x2e, 0x64, 0x62, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x0f, 0x2e, 0x64, 0x62, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x44, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x2e, 0x64,
	0x62, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x16, 0x2e, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x60, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65,
	0x61, 0x72, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x27, 0x2e, 0x64, 0x62, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x16, 0x2e, 0x64, 0x62, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x2a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x2e, 0x64, 0x62,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x1a, 0x0e, 0x2e,
	0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x12, 0x2c, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x12, 0x11, 0x2e, 0x64, 0x62, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x1a, 0x0e, 0x2e, 0x64, 0x62,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x53, 0x69, 0x6d, 0x4e, 0x12, 0x11, 0x2e, 0x64, 0x62, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x4e, 0x1a, 0x12, 0x2e, 0x64, 0x62, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x4e, 0x12, 0x24, 0x0a,
	0x03, 0x53, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x64, 0x62, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x53, 0x65, 0x74, 0x1a, 0x0e, 0x2e, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x0d, 0x44, 0x72, 0x6f, 0x70, 0x50, 0x72, 0x65, 0x64, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x2e, 0x64, 0x62, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x44, 0x72, 0x6f, 0x70, 0x50, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x0e, 0x2e,
	0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x12, 0x54, 0x0a,
	0x1b, 0x44, 0x72, 0x6f, 0x70, 0x4e, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x25, 0x2e, 0x64,
	0x62, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x4e, 0x6f, 0x6e, 0x4c,
	0x69, 0x6e, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x1a, 0x0e, 0x2e, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x2e,
	0x64, 0x62, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x1a,
	0x0e, 0x2e, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x12,
	0x2c, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x50, 0x72, 0x65, 0x64, 0x12, 0x11, 0x2e, 0x64, 0x62, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x50, 0x72, 0x65, 0x64, 0x1a, 0x0e, 0x2e,
	0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x12, 0x30, 0x0a,
	0x09, 0x44, 0x72, 0x6f, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x64, 0x62, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a,
	0x0e, 0x2e, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x12,
	0x3b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x15,
	0x2e, 0x64, 0x62, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x15, 0x2e, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x64, 0x62, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x1a, 0x14, 0x2e, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x64, 0x62, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x64, 0x62, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x27, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x64, 0x62, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0f, 0x2e, 0x64, 0x62, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x4b, 0x0a, 0x08, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1e, 0x2e, 0x64, 0x62, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x1f, 0x2e, 0x64, 0x62, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x39, 0x36, 0x2f, 0x61, 0x68,
	0x6e, 0x6c, 0x69, 0x63, 0x68, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x68, 0x6e, 0x6c, 0x69, 0x63,
	0x68, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3b, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_db_service_proto_goTypes = []any{
	(*query.CreateStore)(nil),                   // 0: db.query.CreateStore
	(*query.CreatePredIndex)(nil),               // 1: db.query.CreatePredIndex
	(*query.CreateNonLinearAlgorithmIndex)(nil), // 2: db.query.CreateNonLinearAlgorithmIndex
	(*query.GetKey)(nil),                        // 3: db.query.GetKey
	(*query.GetPred)(nil),                       // 4: db.query.GetPred
	(*query.GetSimN)(nil),                       // 5: db.query.GetSimN
	(*query.Set)(nil),                           // 6: db.query.Set
	(*query.DropPredIndex)(nil),                 // 7: db.query.DropPredIndex
	(*query.DropNonLinearAlgorithmIndex)(nil),   // 8: db.query.DropNonLinearAlgorithmIndex
	(*query.DelKey)(nil),                        // 9: db.query.DelKey
	(*query.DelPred)(nil),                       // 10: db.query.DelPred
	(*query.DropStore)(nil),                     // 11: db.query.DropStore
	(*query.ListClients)(nil),                   // 12: db.query.ListClients
	(*query.ListStores)(nil),                    // 13: db.query.ListStores
	(*query.InfoServer)(nil),                    // 14: db.query.InfoServer
	(*query.Ping)(nil),                          // 15: db.query.Ping
	(*pipeline.DBRequestPipeline)(nil),          // 16: db.pipeline.DBRequestPipeline
	(*server.Unit)(nil),                         // 17: db.server.Unit
	(*server.CreateIndex)(nil),                  // 18: db.server.CreateIndex
	(*server.Get)(nil),                          // 19: db.server.Get
	(*server.GetSimN)(nil),                      // 20: db.server.GetSimN
	(*server.Set)(nil),                          // 21: db.server.Set
	(*server.Del)(nil),                          // 22: db.server.Del
	(*server.ClientList)(nil),                   // 23: db.server.ClientList
	(*server.StoreList)(nil),                    // 24: db.server.StoreList
	(*server.InfoServer)(nil),                   // 25: db.server.InfoServer
	(*server.Pong)(nil),                         // 26: db.server.Pong
	(*pipeline.DBResponsePipeline)(nil),         // 27: db.pipeline.DBResponsePipeline
}
var file_services_db_service_proto_depIdxs = []int32{
	0,  // 0: services.db_service.DBService.CreateStore:input_type -> db.query.CreateStore
	1,  // 1: services.db_service.DBService.CreatePredIndex:input_type -> db.query.CreatePredIndex
	2,  // 2: services.db_service.DBService.CreateNonLinearAlgorithmIndex:input_type -> db.query.CreateNonLinearAlgorithmIndex
	3,  // 3: services.db_service.DBService.GetKey:input_type -> db.query.GetKey
	4,  // 4: services.db_service.DBService.GetPred:input_type -> db.query.GetPred
	5,  // 5: services.db_service.DBService.GetSimN:input_type -> db.query.GetSimN
	6,  // 6: services.db_service.DBService.Set:input_type -> db.query.Set
	7,  // 7: services.db_service.DBService.DropPredIndex:input_type -> db.query.DropPredIndex
	8,  // 8: services.db_service.DBService.DropNonLinearAlgorithmIndex:input_type -> db.query.DropNonLinearAlgorithmIndex
	9,  // 9: services.db_service.DBService.DelKey:input_type -> db.query.DelKey
	10, // 10: services.db_service.DBService.DelPred:input_type -> db.query.DelPred
	11, // 11: services.db_service.DBService.DropStore:input_type -> db.query.DropStore
	12, // 12: services.db_service.DBService.ListClients:input_type -> db.query.ListClients
	13, // 13: services.db_service.DBService.ListStores:input_type -> db.query.ListStores
	14, // 14: services.db_service.DBService.InfoServer:input_type -> db.query.InfoServer
	15, // 15: services.db_service.DBService.Ping:input_type -> db.query.Ping
	16, // 16: services.db_service.DBService.Pipeline:input_type -> db.pipeline.DBRequestPipeline
	17, // 17: services.db_service.DBService.CreateStore:output_type -> db.server.Unit
	18, // 18: services.db_service.DBService.CreatePredIndex:output_type -> db.server.CreateIndex
	18, // 19: services.db_service.DBService.CreateNonLinearAlgorithmIndex:output_type -> db.server.CreateIndex
	19, // 20: services.db_service.DBService.GetKey:output_type -> db.server.Get
	19, // 21: services.db_service.DBService.GetPred:output_type -> db.server.Get
	20, // 22: services.db_service.DBService.GetSimN:output_type -> db.server.GetSimN
	21, // 23: services.db_service.DBService.Set:output_type -> db.server.Set
	22, // 24: services.db_service.DBService.DropPredIndex:output_type -> db.server.Del
	22, // 25: services.db_service.DBService.DropNonLinearAlgorithmIndex:output_type -> db.server.Del
	22, // 26: services.db_service.DBService.DelKey:output_type -> db.server.Del
	22, // 27: services.db_service.DBService.DelPred:output_type -> db.server.Del
	22, // 28: services.db_service.DBService.DropStore:output_type -> db.server.Del
	23, // 29: services.db_service.DBService.ListClients:output_type -> db.server.ClientList
	24, // 30: services.db_service.DBService.ListStores:output_type -> db.server.StoreList
	25, // 31: services.db_service.DBService.InfoServer:output_type -> db.server.InfoServer
	26, // 32: services.db_service.DBService.Ping:output_type -> db.server.Pong
	27, // 33: services.db_service.DBService.Pipeline:output_type -> db.pipeline.DBResponsePipeline
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_db_service_proto_init() }
func file_services_db_service_proto_init() {
	if File_services_db_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_db_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_db_service_proto_goTypes,
		DependencyIndexes: file_services_db_service_proto_depIdxs,
	}.Build()
	File_services_db_service_proto = out.File
	file_services_db_service_proto_rawDesc = nil
	file_services_db_service_proto_goTypes = nil
	file_services_db_service_proto_depIdxs = nil
}
