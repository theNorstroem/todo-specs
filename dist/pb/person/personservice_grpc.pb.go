// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package personpb

import (
	context "context"
	signatures "github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo/signatures"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PersonsClient is the client API for Persons service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonsClient interface {
	// Use this to create new persons.
	CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*signatures.EmptyEntity, error)
	// We use this to disable a person in the list, we do not delete them.
	DeletePerson(ctx context.Context, in *DeletePersonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Fireing some persons can increase the performance of the other persons POST. Do not use this to much. The big downside is, that you have to assign their work to somone other.
	FirePerson(ctx context.Context, in *FirePersonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Returns a single person.
	GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*PersonEntity, error)
	// List persons with pagination.
	ListPersons(ctx context.Context, in *ListPersonsRequest, opts ...grpc.CallOption) (*PersonCollection, error)
	// Use this to update existing persons. PATCH is also supported
	UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*PersonEntity, error)
}

type personsClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonsClient(cc grpc.ClientConnInterface) PersonsClient {
	return &personsClient{cc}
}

func (c *personsClient) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*signatures.EmptyEntity, error) {
	out := new(signatures.EmptyEntity)
	err := c.cc.Invoke(ctx, "/person.Persons/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) DeletePerson(ctx context.Context, in *DeletePersonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/person.Persons/DeletePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) FirePerson(ctx context.Context, in *FirePersonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/person.Persons/FirePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*PersonEntity, error) {
	out := new(PersonEntity)
	err := c.cc.Invoke(ctx, "/person.Persons/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) ListPersons(ctx context.Context, in *ListPersonsRequest, opts ...grpc.CallOption) (*PersonCollection, error) {
	out := new(PersonCollection)
	err := c.cc.Invoke(ctx, "/person.Persons/ListPersons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*PersonEntity, error) {
	out := new(PersonEntity)
	err := c.cc.Invoke(ctx, "/person.Persons/UpdatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonsServer is the server API for Persons service.
// All implementations must embed UnimplementedPersonsServer
// for forward compatibility
type PersonsServer interface {
	// Use this to create new persons.
	CreatePerson(context.Context, *CreatePersonRequest) (*signatures.EmptyEntity, error)
	// We use this to disable a person in the list, we do not delete them.
	DeletePerson(context.Context, *DeletePersonRequest) (*emptypb.Empty, error)
	// Fireing some persons can increase the performance of the other persons POST. Do not use this to much. The big downside is, that you have to assign their work to somone other.
	FirePerson(context.Context, *FirePersonRequest) (*emptypb.Empty, error)
	// Returns a single person.
	GetPerson(context.Context, *GetPersonRequest) (*PersonEntity, error)
	// List persons with pagination.
	ListPersons(context.Context, *ListPersonsRequest) (*PersonCollection, error)
	// Use this to update existing persons. PATCH is also supported
	UpdatePerson(context.Context, *UpdatePersonRequest) (*PersonEntity, error)
	mustEmbedUnimplementedPersonsServer()
}

// UnimplementedPersonsServer must be embedded to have forward compatible implementations.
type UnimplementedPersonsServer struct {
}

func (UnimplementedPersonsServer) CreatePerson(context.Context, *CreatePersonRequest) (*signatures.EmptyEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedPersonsServer) DeletePerson(context.Context, *DeletePersonRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePerson not implemented")
}
func (UnimplementedPersonsServer) FirePerson(context.Context, *FirePersonRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FirePerson not implemented")
}
func (UnimplementedPersonsServer) GetPerson(context.Context, *GetPersonRequest) (*PersonEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (UnimplementedPersonsServer) ListPersons(context.Context, *ListPersonsRequest) (*PersonCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPersons not implemented")
}
func (UnimplementedPersonsServer) UpdatePerson(context.Context, *UpdatePersonRequest) (*PersonEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (UnimplementedPersonsServer) mustEmbedUnimplementedPersonsServer() {}

// UnsafePersonsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonsServer will
// result in compilation errors.
type UnsafePersonsServer interface {
	mustEmbedUnimplementedPersonsServer()
}

func RegisterPersonsServer(s *grpc.Server, srv PersonsServer) {
	s.RegisterService(&_Persons_serviceDesc, srv)
}

func _Persons_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.Persons/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).CreatePerson(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_DeletePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).DeletePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.Persons/DeletePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).DeletePerson(ctx, req.(*DeletePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_FirePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FirePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).FirePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.Persons/FirePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).FirePerson(ctx, req.(*FirePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.Persons/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).GetPerson(ctx, req.(*GetPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_ListPersons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPersonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).ListPersons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.Persons/ListPersons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).ListPersons(ctx, req.(*ListPersonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.Persons/UpdatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).UpdatePerson(ctx, req.(*UpdatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Persons_serviceDesc = grpc.ServiceDesc{
	ServiceName: "person.Persons",
	HandlerType: (*PersonsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePerson",
			Handler:    _Persons_CreatePerson_Handler,
		},
		{
			MethodName: "DeletePerson",
			Handler:    _Persons_DeletePerson_Handler,
		},
		{
			MethodName: "FirePerson",
			Handler:    _Persons_FirePerson_Handler,
		},
		{
			MethodName: "GetPerson",
			Handler:    _Persons_GetPerson_Handler,
		},
		{
			MethodName: "ListPersons",
			Handler:    _Persons_ListPersons_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _Persons_UpdatePerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person/personservice.proto",
}
