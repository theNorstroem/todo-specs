// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: person/reqmsgs.proto

package personpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// request message for CreatePersons
type CreatePersonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with person.Person
	Body *Person `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CreatePersonsRequest) Reset() {
	*x = CreatePersonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_reqmsgs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonsRequest) ProtoMessage() {}

func (x *CreatePersonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_reqmsgs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonsRequest.ProtoReflect.Descriptor instead.
func (*CreatePersonsRequest) Descriptor() ([]byte, []int) {
	return file_person_reqmsgs_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePersonsRequest) GetBody() *Person {
	if x != nil {
		return x.Body
	}
	return nil
}

// request message for DeletePersons
type DeletePersonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with google.protobuf.Empty
	Body *emptypb.Empty `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// The query param pn stands for the id of a person.
	Pn string `protobuf:"bytes,2,opt,name=pn,proto3" json:"pn,omitempty"`
}

func (x *DeletePersonsRequest) Reset() {
	*x = DeletePersonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_reqmsgs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePersonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePersonsRequest) ProtoMessage() {}

func (x *DeletePersonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_reqmsgs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePersonsRequest.ProtoReflect.Descriptor instead.
func (*DeletePersonsRequest) Descriptor() ([]byte, []int) {
	return file_person_reqmsgs_proto_rawDescGZIP(), []int{1}
}

func (x *DeletePersonsRequest) GetBody() *emptypb.Empty {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *DeletePersonsRequest) GetPn() string {
	if x != nil {
		return x.Pn
	}
	return ""
}

// request message for FirePersons
type FirePersonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with google.protobuf.Empty
	Body *emptypb.Empty `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// The query param pn stands for the id of a person.
	Pn string `protobuf:"bytes,2,opt,name=pn,proto3" json:"pn,omitempty"`
}

func (x *FirePersonsRequest) Reset() {
	*x = FirePersonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_reqmsgs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirePersonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirePersonsRequest) ProtoMessage() {}

func (x *FirePersonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_reqmsgs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirePersonsRequest.ProtoReflect.Descriptor instead.
func (*FirePersonsRequest) Descriptor() ([]byte, []int) {
	return file_person_reqmsgs_proto_rawDescGZIP(), []int{2}
}

func (x *FirePersonsRequest) GetBody() *emptypb.Empty {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *FirePersonsRequest) GetPn() string {
	if x != nil {
		return x.Pn
	}
	return ""
}

// request message for GetPersons
type GetPersonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with google.protobuf.Empty
	Body *emptypb.Empty `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// The query param pn stands for the id of a person.
	Pn string `protobuf:"bytes,2,opt,name=pn,proto3" json:"pn,omitempty"`
}

func (x *GetPersonsRequest) Reset() {
	*x = GetPersonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_reqmsgs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonsRequest) ProtoMessage() {}

func (x *GetPersonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_reqmsgs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonsRequest.ProtoReflect.Descriptor instead.
func (*GetPersonsRequest) Descriptor() ([]byte, []int) {
	return file_person_reqmsgs_proto_rawDescGZIP(), []int{3}
}

func (x *GetPersonsRequest) GetBody() *emptypb.Empty {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *GetPersonsRequest) GetPn() string {
	if x != nil {
		return x.Pn
	}
	return ""
}

// request message for ListPersonss
type ListPersonssRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with google.protobuf.Empty
	Body *emptypb.Empty `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Use this to search for a person by text.
	Q string `protobuf:"bytes,2,opt,name=q,proto3" json:"q,omitempty"`
	// Use this field to filter the persons, this is not searching.
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Use this field to specify the ordering.
	OrderBy string `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Use this field to specify page to display.
	Page string `protobuf:"bytes,5,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListPersonssRequest) Reset() {
	*x = ListPersonssRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_reqmsgs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonssRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonssRequest) ProtoMessage() {}

func (x *ListPersonssRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_reqmsgs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonssRequest.ProtoReflect.Descriptor instead.
func (*ListPersonssRequest) Descriptor() ([]byte, []int) {
	return file_person_reqmsgs_proto_rawDescGZIP(), []int{4}
}

func (x *ListPersonssRequest) GetBody() *emptypb.Empty {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *ListPersonssRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *ListPersonssRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListPersonssRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListPersonssRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

// request message for UpdatePersons
type UpdatePersonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with person.Person
	Body *Person `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// The query param pn stands for the id of a person.
	Pn string `protobuf:"bytes,2,opt,name=pn,proto3" json:"pn,omitempty"`
	// Needed to patch a record
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdatePersonsRequest) Reset() {
	*x = UpdatePersonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_reqmsgs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePersonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersonsRequest) ProtoMessage() {}

func (x *UpdatePersonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_reqmsgs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersonsRequest.ProtoReflect.Descriptor instead.
func (*UpdatePersonsRequest) Descriptor() ([]byte, []int) {
	return file_person_reqmsgs_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePersonsRequest) GetBody() *Person {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *UpdatePersonsRequest) GetPn() string {
	if x != nil {
		return x.Pn
	}
	return ""
}

func (x *UpdatePersonsRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_person_reqmsgs_proto protoreflect.FileDescriptor

var file_person_reqmsgs_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x71, 0x6d, 0x73, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x52,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x70, 0x6e, 0x22, 0x50, 0x0a, 0x12, 0x46, 0x69, 0x72, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x70, 0x6e, 0x22, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x70, 0x6e, 0x22, 0x96, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x87,
	0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x70,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x62, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x66, 0x75, 0x72, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42,
	0x0c, 0x52, 0x65, 0x71, 0x6d, 0x73, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x4e,
	0x6f, 0x72, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x73, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x73,
	0x70, 0x65, 0x63, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_reqmsgs_proto_rawDescOnce sync.Once
	file_person_reqmsgs_proto_rawDescData = file_person_reqmsgs_proto_rawDesc
)

func file_person_reqmsgs_proto_rawDescGZIP() []byte {
	file_person_reqmsgs_proto_rawDescOnce.Do(func() {
		file_person_reqmsgs_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_reqmsgs_proto_rawDescData)
	})
	return file_person_reqmsgs_proto_rawDescData
}

var file_person_reqmsgs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_person_reqmsgs_proto_goTypes = []interface{}{
	(*CreatePersonsRequest)(nil),  // 0: person.CreatePersonsRequest
	(*DeletePersonsRequest)(nil),  // 1: person.DeletePersonsRequest
	(*FirePersonsRequest)(nil),    // 2: person.FirePersonsRequest
	(*GetPersonsRequest)(nil),     // 3: person.GetPersonsRequest
	(*ListPersonssRequest)(nil),   // 4: person.ListPersonssRequest
	(*UpdatePersonsRequest)(nil),  // 5: person.UpdatePersonsRequest
	(*Person)(nil),                // 6: person.Person
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
}
var file_person_reqmsgs_proto_depIdxs = []int32{
	6, // 0: person.CreatePersonsRequest.body:type_name -> person.Person
	7, // 1: person.DeletePersonsRequest.body:type_name -> google.protobuf.Empty
	7, // 2: person.FirePersonsRequest.body:type_name -> google.protobuf.Empty
	7, // 3: person.GetPersonsRequest.body:type_name -> google.protobuf.Empty
	7, // 4: person.ListPersonssRequest.body:type_name -> google.protobuf.Empty
	6, // 5: person.UpdatePersonsRequest.body:type_name -> person.Person
	8, // 6: person.UpdatePersonsRequest.update_mask:type_name -> google.protobuf.FieldMask
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_person_reqmsgs_proto_init() }
func file_person_reqmsgs_proto_init() {
	if File_person_reqmsgs_proto != nil {
		return
	}
	file_person_person_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_person_reqmsgs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonsRequest); i {
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
		file_person_reqmsgs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePersonsRequest); i {
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
		file_person_reqmsgs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirePersonsRequest); i {
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
		file_person_reqmsgs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonsRequest); i {
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
		file_person_reqmsgs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonssRequest); i {
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
		file_person_reqmsgs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePersonsRequest); i {
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
			RawDescriptor: file_person_reqmsgs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_reqmsgs_proto_goTypes,
		DependencyIndexes: file_person_reqmsgs_proto_depIdxs,
		MessageInfos:      file_person_reqmsgs_proto_msgTypes,
	}.Build()
	File_person_reqmsgs_proto = out.File
	file_person_reqmsgs_proto_rawDesc = nil
	file_person_reqmsgs_proto_goTypes = nil
	file_person_reqmsgs_proto_depIdxs = nil
}