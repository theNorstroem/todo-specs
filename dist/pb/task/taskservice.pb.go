// Code generated by furo spectools. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: task/taskservice.proto

package taskpb

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_task_taskservice_proto protoreflect.FileDescriptor

var file_task_taskservice_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x61, 0x73,
	0x6b, 0x2f, 0x72, 0x65, 0x71, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xfa, 0x04, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x55, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x18, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x22, 0x06, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x3a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x55, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x12, 0x18, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2f, 0x7b, 0x74, 0x73, 0x6b, 0x7d, 0x12, 0x57, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x2a, 0x06, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x65, 0x0a, 0x0c, 0x46, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x46, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x14, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x74, 0x73, 0x6b, 0x7d, 0x3a, 0x66, 0x65, 0x72, 0x6d, 0x65,
	0x6e, 0x74, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x49, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x12, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x14, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x74,
	0x73, 0x6b, 0x7d, 0x12, 0x4b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x73, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x12, 0x6b, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12,
	0x18, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2a, 0x1a, 0x0c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x74, 0x73, 0x6b,
	0x7d, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x5a, 0x14, 0x32, 0x0c, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2f, 0x7b, 0x74, 0x73, 0x6b, 0x7d, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x5c, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x42, 0x10,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x68, 0x65, 0x4e, 0x6f, 0x72, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x73, 0x6d, 0x2f, 0x74, 0x6f, 0x64,
	0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_task_taskservice_proto_goTypes = []interface{}{
	(*CreateTasksRequest)(nil),     // 0: task.CreateTasksRequest
	(*DeleteTasksRequest)(nil),     // 1: task.DeleteTasksRequest
	(*DeleteAllTaskssRequest)(nil), // 2: task.DeleteAllTaskssRequest
	(*FermentTasksRequest)(nil),    // 3: task.FermentTasksRequest
	(*GetTasksRequest)(nil),        // 4: task.GetTasksRequest
	(*ListTaskssRequest)(nil),      // 5: task.ListTaskssRequest
	(*UpdateTasksRequest)(nil),     // 6: task.UpdateTasksRequest
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
	(*TaskEntity)(nil),             // 8: task.TaskEntity
	(*TaskCollection)(nil),         // 9: task.TaskCollection
}
var file_task_taskservice_proto_depIdxs = []int32{
	0, // 0: task.Tasks.CreateTasks:input_type -> task.CreateTasksRequest
	1, // 1: task.Tasks.DeleteTasks:input_type -> task.DeleteTasksRequest
	2, // 2: task.Tasks.DeleteAllTaskss:input_type -> task.DeleteAllTaskssRequest
	3, // 3: task.Tasks.FermentTasks:input_type -> task.FermentTasksRequest
	4, // 4: task.Tasks.GetTasks:input_type -> task.GetTasksRequest
	5, // 5: task.Tasks.ListTaskss:input_type -> task.ListTaskssRequest
	6, // 6: task.Tasks.UpdateTasks:input_type -> task.UpdateTasksRequest
	7, // 7: task.Tasks.CreateTasks:output_type -> google.protobuf.Empty
	7, // 8: task.Tasks.DeleteTasks:output_type -> google.protobuf.Empty
	7, // 9: task.Tasks.DeleteAllTaskss:output_type -> google.protobuf.Empty
	7, // 10: task.Tasks.FermentTasks:output_type -> google.protobuf.Empty
	8, // 11: task.Tasks.GetTasks:output_type -> task.TaskEntity
	9, // 12: task.Tasks.ListTaskss:output_type -> task.TaskCollection
	8, // 13: task.Tasks.UpdateTasks:output_type -> task.TaskEntity
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_task_taskservice_proto_init() }
func file_task_taskservice_proto_init() {
	if File_task_taskservice_proto != nil {
		return
	}
	file_task_reqmsgs_proto_init()
	file_task_task_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_taskservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_taskservice_proto_goTypes,
		DependencyIndexes: file_task_taskservice_proto_depIdxs,
	}.Build()
	File_task_taskservice_proto = out.File
	file_task_taskservice_proto_rawDesc = nil
	file_task_taskservice_proto_goTypes = nil
	file_task_taskservice_proto_depIdxs = nil
}
