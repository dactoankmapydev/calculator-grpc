// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: service_calculator.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_service_calculator_proto protoreflect.FileDescriptor

var file_service_calculator_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb5,
	0x02, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x6d, 0x12, 0x34, 0x0a,
	0x0f, 0x53, 0x75, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x18, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x4e, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x4e, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x38, 0x0a, 0x07,
	0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_calculator_proto_goTypes = []interface{}{
	(*SumRequest)(nil),      // 0: pb.SumRequest
	(*PNDRequest)(nil),      // 1: pb.PNDRequest
	(*AverageRequest)(nil),  // 2: pb.AverageRequest
	(*FindMaxRequest)(nil),  // 3: pb.FindMaxRequest
	(*SumResponse)(nil),     // 4: pb.SumResponse
	(*PNDResponse)(nil),     // 5: pb.PNDResponse
	(*AverageResponse)(nil), // 6: pb.AverageResponse
	(*FindMaxResponse)(nil), // 7: pb.FindMaxResponse
}
var file_service_calculator_proto_depIdxs = []int32{
	0, // 0: pb.CalculatorService.Sum:input_type -> pb.SumRequest
	0, // 1: pb.CalculatorService.SumWithDeadline:input_type -> pb.SumRequest
	1, // 2: pb.CalculatorService.PrimeNumberDecomposition:input_type -> pb.PNDRequest
	2, // 3: pb.CalculatorService.Average:input_type -> pb.AverageRequest
	3, // 4: pb.CalculatorService.FindMax:input_type -> pb.FindMaxRequest
	4, // 5: pb.CalculatorService.Sum:output_type -> pb.SumResponse
	4, // 6: pb.CalculatorService.SumWithDeadline:output_type -> pb.SumResponse
	5, // 7: pb.CalculatorService.PrimeNumberDecomposition:output_type -> pb.PNDResponse
	6, // 8: pb.CalculatorService.Average:output_type -> pb.AverageResponse
	7, // 9: pb.CalculatorService.FindMax:output_type -> pb.FindMaxResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_calculator_proto_init() }
func file_service_calculator_proto_init() {
	if File_service_calculator_proto != nil {
		return
	}
	file_calculator_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_calculator_proto_goTypes,
		DependencyIndexes: file_service_calculator_proto_depIdxs,
	}.Build()
	File_service_calculator_proto = out.File
	file_service_calculator_proto_rawDesc = nil
	file_service_calculator_proto_goTypes = nil
	file_service_calculator_proto_depIdxs = nil
}
