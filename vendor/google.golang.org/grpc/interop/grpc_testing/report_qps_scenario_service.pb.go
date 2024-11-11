// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// An integration test service that covers all the method signature permutations
// of unary/streaming requests/responses.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: grpc/testing/report_qps_scenario_service.proto

package grpc_testing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_grpc_testing_report_qps_scenario_service_proto protoreflect.FileDescriptor

var file_grpc_testing_report_qps_scenario_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x71, 0x70, 0x73, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1a,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5e, 0x0a, 0x18, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x51, 0x70, 0x73, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x32, 0x0a, 0x0f, 0x69, 0x6f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x1d, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x51, 0x70, 0x73, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_grpc_testing_report_qps_scenario_service_proto_goTypes = []any{
	(*ScenarioResult)(nil), // 0: grpc.testing.ScenarioResult
	(*Void)(nil),           // 1: grpc.testing.Void
}
var file_grpc_testing_report_qps_scenario_service_proto_depIdxs = []int32{
	0, // 0: grpc.testing.ReportQpsScenarioService.ReportScenario:input_type -> grpc.testing.ScenarioResult
	1, // 1: grpc.testing.ReportQpsScenarioService.ReportScenario:output_type -> grpc.testing.Void
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_testing_report_qps_scenario_service_proto_init() }
func file_grpc_testing_report_qps_scenario_service_proto_init() {
	if File_grpc_testing_report_qps_scenario_service_proto != nil {
		return
	}
	file_grpc_testing_control_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_testing_report_qps_scenario_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_testing_report_qps_scenario_service_proto_goTypes,
		DependencyIndexes: file_grpc_testing_report_qps_scenario_service_proto_depIdxs,
	}.Build()
	File_grpc_testing_report_qps_scenario_service_proto = out.File
	file_grpc_testing_report_qps_scenario_service_proto_rawDesc = nil
	file_grpc_testing_report_qps_scenario_service_proto_goTypes = nil
	file_grpc_testing_report_qps_scenario_service_proto_depIdxs = nil
}
