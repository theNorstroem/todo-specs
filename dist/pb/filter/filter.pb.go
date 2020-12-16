// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: filter/filter.proto

package filterpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

// Filter condition
type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Field
	Fld string `protobuf:"bytes,1,opt,name=fld,proto3" json:"fld,omitempty"`
	// The comparator like gt, eq,...
	Is string `protobuf:"bytes,2,opt,name=is,proto3" json:"is,omitempty"`
	// The value as string, parse this for your field
	Val string `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"`
	// And bracket with ors inside
	Aoc []*Condition `protobuf:"bytes,4,rep,name=aoc,proto3" json:"aoc,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_filter_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_filter_filter_proto_rawDescGZIP(), []int{0}
}

func (x *Condition) GetFld() string {
	if x != nil {
		return x.Fld
	}
	return ""
}

func (x *Condition) GetIs() string {
	if x != nil {
		return x.Is
	}
	return ""
}

func (x *Condition) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

func (x *Condition) GetAoc() []*Condition {
	if x != nil {
		return x.Aoc
	}
	return nil
}

// Filter root object
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Root bracket with ors inside, this is the most complex but most flexible way to define a filter
	Clause []*Condition `protobuf:"bytes,1,rep,name=clause,proto3" json:"clause,omitempty"`
	// Shortcut to set filter conditions without nesting.
	// It is up to you how the server handles the request.
	// Examples for a flat filter a,b,c:
	// - all active conditions *must* match (a && b && c).
	// - all conditions are handled as or (a || b || c).
	// - you build your own logic like (a && b) || c.
	Flat map[string]*Condition `protobuf:"bytes,2,rep,name=flat,proto3" json:"flat,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_filter_filter_proto_rawDescGZIP(), []int{1}
}

func (x *Filter) GetClause() []*Condition {
	if x != nil {
		return x.Clause
	}
	return nil
}

func (x *Filter) GetFlat() map[string]*Condition {
	if x != nil {
		return x.Flat
	}
	return nil
}

var File_filter_filter_proto protoreflect.FileDescriptor

var file_filter_filter_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x64, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x23,
	0x0a, 0x03, 0x61, 0x6f, 0x63, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03,
	0x61, 0x6f, 0x63, 0x22, 0xad, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x06, 0x63, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x66, 0x6c, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x6c, 0x61, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x74, 0x1a, 0x4a, 0x0a, 0x09, 0x46, 0x6c, 0x61, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x60, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x0b, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x4e, 0x6f, 0x72, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3b, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filter_filter_proto_rawDescOnce sync.Once
	file_filter_filter_proto_rawDescData = file_filter_filter_proto_rawDesc
)

func file_filter_filter_proto_rawDescGZIP() []byte {
	file_filter_filter_proto_rawDescOnce.Do(func() {
		file_filter_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_filter_filter_proto_rawDescData)
	})
	return file_filter_filter_proto_rawDescData
}

var file_filter_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_filter_filter_proto_goTypes = []interface{}{
	(*Condition)(nil), // 0: filter.Condition
	(*Filter)(nil),    // 1: filter.Filter
	nil,               // 2: filter.Filter.FlatEntry
}
var file_filter_filter_proto_depIdxs = []int32{
	0, // 0: filter.Condition.aoc:type_name -> filter.Condition
	0, // 1: filter.Filter.clause:type_name -> filter.Condition
	2, // 2: filter.Filter.flat:type_name -> filter.Filter.FlatEntry
	0, // 3: filter.Filter.FlatEntry.value:type_name -> filter.Condition
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_filter_filter_proto_init() }
func file_filter_filter_proto_init() {
	if File_filter_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filter_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_filter_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filter_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filter_filter_proto_goTypes,
		DependencyIndexes: file_filter_filter_proto_depIdxs,
		MessageInfos:      file_filter_filter_proto_msgTypes,
	}.Build()
	File_filter_filter_proto = out.File
	file_filter_filter_proto_rawDesc = nil
	file_filter_filter_proto_goTypes = nil
	file_filter_filter_proto_depIdxs = nil
}