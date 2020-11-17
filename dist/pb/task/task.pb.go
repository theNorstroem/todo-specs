// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: task/task.proto

package taskpb

import (
	proto "github.com/golang/protobuf/proto"
	furo "github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo"
	person "github.com/theNorstroem/todo-specs/dist/pb/person"
	date "google.golang.org/genproto/googleapis/type/date"
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

// A checkbox item
type CheckboxItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// We use this field label the task
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Sample
	Done bool `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	// Add some notes
	Note string `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *CheckboxItem) Reset() {
	*x = CheckboxItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckboxItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckboxItem) ProtoMessage() {}

func (x *CheckboxItem) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckboxItem.ProtoReflect.Descriptor instead.
func (*CheckboxItem) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{0}
}

func (x *CheckboxItem) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CheckboxItem) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *CheckboxItem) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

// Use this to set a reference to a task
type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the referenced task.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Label of the referenced task
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// HTS for the initial search (works on root resources only)
	Link *furo.Link `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *Reference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reference) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Reference) GetLink() *furo.Link {
	if x != nil {
		return x.Link
	}
	return nil
}

// Collectioncontainer which holds a task.Reference
type ReferenceCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the data contains a task.Reference
	Entities []*ReferenceEntity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	// the Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta *furo.Meta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ReferenceCollection) Reset() {
	*x = ReferenceCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferenceCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferenceCollection) ProtoMessage() {}

func (x *ReferenceCollection) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferenceCollection.ProtoReflect.Descriptor instead.
func (*ReferenceCollection) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{2}
}

func (x *ReferenceCollection) GetEntities() []*ReferenceEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *ReferenceCollection) GetLinks() []*furo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *ReferenceCollection) GetMeta() *furo.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Entitycontainer which holds a task.Reference
type ReferenceEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the data contains a task.Reference
	Data *Reference `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// the Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta *furo.Meta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ReferenceEntity) Reset() {
	*x = ReferenceEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferenceEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferenceEntity) ProtoMessage() {}

func (x *ReferenceEntity) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferenceEntity.ProtoReflect.Descriptor instead.
func (*ReferenceEntity) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{3}
}

func (x *ReferenceEntity) GetData() *Reference {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ReferenceEntity) GetLinks() []*furo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *ReferenceEntity) GetMeta() *furo.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Tasks are essential at work.
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Task id (is a ULID).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// We use this field label the task
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Add some notes
	Note string `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	// The parent for this task
	Parent *Task `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	// A task can have a checklist (more work)
	Checklist []string `protobuf:"bytes,5,rep,name=checklist,proto3" json:"checklist,omitempty"`
	// Use this to give pressure.
	DueDate *date.Date `protobuf:"bytes,6,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	// Maybe we can benefit.
	RelatedTasks []*Reference `protobuf:"bytes,7,rep,name=related_tasks,json=relatedTasks,proto3" json:"related_tasks,omitempty"`
	// Sometimes a task is to hard for one person only.
	InvolvedPersons []*person.Reference `protobuf:"bytes,8,rep,name=involved_persons,json=involvedPersons,proto3" json:"involved_persons,omitempty"`
	// If something goes wrong, we should be able to make some one responsible for it. This is aligned with our CYA strategy.
	ResponsiblePerson *person.Reference `protobuf:"bytes,9,opt,name=responsible_person,json=responsiblePerson,proto3" json:"responsible_person,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{4}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Task) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Task) GetParent() *Task {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Task) GetChecklist() []string {
	if x != nil {
		return x.Checklist
	}
	return nil
}

func (x *Task) GetDueDate() *date.Date {
	if x != nil {
		return x.DueDate
	}
	return nil
}

func (x *Task) GetRelatedTasks() []*Reference {
	if x != nil {
		return x.RelatedTasks
	}
	return nil
}

func (x *Task) GetInvolvedPersons() []*person.Reference {
	if x != nil {
		return x.InvolvedPersons
	}
	return nil
}

func (x *Task) GetResponsiblePerson() *person.Reference {
	if x != nil {
		return x.ResponsiblePerson
	}
	return nil
}

// Collectioncontainer which holds a task.Task
type TaskCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the data contains a task.Task
	Entities []*TaskEntity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	// the Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta *furo.Meta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *TaskCollection) Reset() {
	*x = TaskCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCollection) ProtoMessage() {}

func (x *TaskCollection) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCollection.ProtoReflect.Descriptor instead.
func (*TaskCollection) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{5}
}

func (x *TaskCollection) GetEntities() []*TaskEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *TaskCollection) GetLinks() []*furo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *TaskCollection) GetMeta() *furo.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Entitycontainer which holds a task.Task
type TaskEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the data contains a task.Task
	Data *Task `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// the Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta *furo.Meta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *TaskEntity) Reset() {
	*x = TaskEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskEntity) ProtoMessage() {}

func (x *TaskEntity) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskEntity.ProtoReflect.Descriptor instead.
func (*TaskEntity) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{6}
}

func (x *TaskEntity) GetData() *Task {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TaskEntity) GetLinks() []*furo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *TaskEntity) GetMeta() *furo.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_task_task_proto protoreflect.FileDescriptor

var file_task_task_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x0f, 0x66, 0x75, 0x72, 0x6f, 0x2f, 0x66, 0x75,
	0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x6f,
	0x78, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x22, 0x5e, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x22, 0x8a, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72,
	0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1e, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75,
	0x72, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x78, 0x0a,
	0x0f, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0xf3, 0x02, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x64, 0x75, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x07,
	0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x3c, 0x0a,
	0x10, 0x69, 0x6e, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x12, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x11, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x80, 0x01,
	0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x20,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x22, 0x6e, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x42, 0x55, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x42, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x4e, 0x6f,
	0x72, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x73, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x73, 0x70,
	0x65, 0x63, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x3b, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_task_proto_rawDescOnce sync.Once
	file_task_task_proto_rawDescData = file_task_task_proto_rawDesc
)

func file_task_task_proto_rawDescGZIP() []byte {
	file_task_task_proto_rawDescOnce.Do(func() {
		file_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_task_proto_rawDescData)
	})
	return file_task_task_proto_rawDescData
}

var file_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_task_task_proto_goTypes = []interface{}{
	(*CheckboxItem)(nil),        // 0: task.CheckboxItem
	(*Reference)(nil),           // 1: task.Reference
	(*ReferenceCollection)(nil), // 2: task.ReferenceCollection
	(*ReferenceEntity)(nil),     // 3: task.ReferenceEntity
	(*Task)(nil),                // 4: task.Task
	(*TaskCollection)(nil),      // 5: task.TaskCollection
	(*TaskEntity)(nil),          // 6: task.TaskEntity
	(*furo.Link)(nil),           // 7: furo.Link
	(*furo.Meta)(nil),           // 8: furo.Meta
	(*date.Date)(nil),           // 9: google.type.Date
	(*person.Reference)(nil),    // 10: person.Reference
}
var file_task_task_proto_depIdxs = []int32{
	7,  // 0: task.Reference.link:type_name -> furo.Link
	3,  // 1: task.ReferenceCollection.entities:type_name -> task.ReferenceEntity
	7,  // 2: task.ReferenceCollection.links:type_name -> furo.Link
	8,  // 3: task.ReferenceCollection.meta:type_name -> furo.Meta
	1,  // 4: task.ReferenceEntity.data:type_name -> task.Reference
	7,  // 5: task.ReferenceEntity.links:type_name -> furo.Link
	8,  // 6: task.ReferenceEntity.meta:type_name -> furo.Meta
	4,  // 7: task.Task.parent:type_name -> task.Task
	9,  // 8: task.Task.due_date:type_name -> google.type.Date
	1,  // 9: task.Task.related_tasks:type_name -> task.Reference
	10, // 10: task.Task.involved_persons:type_name -> person.Reference
	10, // 11: task.Task.responsible_person:type_name -> person.Reference
	6,  // 12: task.TaskCollection.entities:type_name -> task.TaskEntity
	7,  // 13: task.TaskCollection.links:type_name -> furo.Link
	8,  // 14: task.TaskCollection.meta:type_name -> furo.Meta
	4,  // 15: task.TaskEntity.data:type_name -> task.Task
	7,  // 16: task.TaskEntity.links:type_name -> furo.Link
	8,  // 17: task.TaskEntity.meta:type_name -> furo.Meta
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_task_task_proto_init() }
func file_task_task_proto_init() {
	if File_task_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckboxItem); i {
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
		file_task_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_task_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferenceCollection); i {
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
		file_task_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferenceEntity); i {
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
		file_task_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_task_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCollection); i {
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
		file_task_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskEntity); i {
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
			RawDescriptor: file_task_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_task_task_proto_goTypes,
		DependencyIndexes: file_task_task_proto_depIdxs,
		MessageInfos:      file_task_task_proto_msgTypes,
	}.Build()
	File_task_task_proto = out.File
	file_task_task_proto_rawDesc = nil
	file_task_task_proto_goTypes = nil
	file_task_task_proto_depIdxs = nil
}
