// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/longrunning/operations.proto

/*
Package longrunning is a generated protocol buffer package.

It is generated from these files:
	google/longrunning/operations.proto

It has these top-level messages:
	Operation
	GetOperationRequest
	ListOperationsRequest
	ListOperationsResponse
	CancelOperationRequest
	DeleteOperationRequest
*/
package longrunning

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// The server-assigned name, which is only unique within the same service that
	// originally returns it. If you use the default HTTP mapping, the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Service-specific metadata associated with the operation.  It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.  Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata *google_protobuf1.Any `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If true, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	// The operation result, which can be either an `error` or a valid `response`.
	// If `done` == `false`, neither `error` nor `response` is set.
	// If `done` == `true`, exactly one of `error` or `response` is set.
	//
	// Types that are valid to be assigned to Result:
	//	*Operation_Error
	//	*Operation_Response
	Result isOperation_Result `protobuf_oneof:"result"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	Error *google_rpc.Status `protobuf:"bytes,4,opt,name=error,oneof"`
}
type Operation_Response struct {
	Response *google_protobuf1.Any `protobuf:"bytes,5,opt,name=response,oneof"`
}

func (*Operation_Error) isOperation_Result()    {}
func (*Operation_Response) isOperation_Result() {}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Operation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Operation) GetMetadata() *google_protobuf1.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Operation) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Operation) GetError() *google_rpc.Status {
	if x, ok := m.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Operation) GetResponse() *google_protobuf1.Any {
	if x, ok := m.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *Operation_Response:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Response); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Result has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 4: // result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_rpc.Status)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Error{msg}
		return true, err
	case 5: // result.response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Response{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Response:
		s := proto.Size(x.Response)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The request message for [Operations.GetOperation][google.longrunning.Operations.GetOperation].
type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetOperationRequest) Reset()                    { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()               {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsRequest struct {
	// The name of the operation collection.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// The standard list filter.
	Filter string `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	// The standard list page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The standard list page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListOperationsRequest) Reset()                    { *m = ListOperationsRequest{} }
func (m *ListOperationsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOperationsRequest) ProtoMessage()               {}
func (*ListOperationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListOperationsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListOperationsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListOperationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListOperationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsResponse struct {
	// A list of operations that matches the specified filter in the request.
	Operations []*Operation `protobuf:"bytes,1,rep,name=operations" json:"operations,omitempty"`
	// The standard List next-page token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListOperationsResponse) Reset()                    { *m = ListOperationsResponse{} }
func (m *ListOperationsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOperationsResponse) ProtoMessage()               {}
func (*ListOperationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListOperationsResponse) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *ListOperationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [Operations.CancelOperation][google.longrunning.Operations.CancelOperation].
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CancelOperationRequest) Reset()                    { *m = CancelOperationRequest{} }
func (m *CancelOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelOperationRequest) ProtoMessage()               {}
func (*CancelOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CancelOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.DeleteOperation][google.longrunning.Operations.DeleteOperation].
type DeleteOperationRequest struct {
	// The name of the operation resource to be deleted.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteOperationRequest) Reset()                    { *m = DeleteOperationRequest{} }
func (m *DeleteOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteOperationRequest) ProtoMessage()               {}
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Operation)(nil), "google.longrunning.Operation")
	proto.RegisterType((*GetOperationRequest)(nil), "google.longrunning.GetOperationRequest")
	proto.RegisterType((*ListOperationsRequest)(nil), "google.longrunning.ListOperationsRequest")
	proto.RegisterType((*ListOperationsResponse)(nil), "google.longrunning.ListOperationsResponse")
	proto.RegisterType((*CancelOperationRequest)(nil), "google.longrunning.CancelOperationRequest")
	proto.RegisterType((*DeleteOperationRequest)(nil), "google.longrunning.DeleteOperationRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Operations service

type OperationsClient interface {
	// Lists operations that match the specified filter in the request. If the
	// server doesn't support this method, it returns `UNIMPLEMENTED`.
	//
	// NOTE: the `name` binding below allows API services to override the binding
	// to use different resource name schemes, such as `users/*/operations`.
	ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error)
	// Gets the latest state of a long-running operation.  Clients can use this
	// method to poll the operation result at intervals as recommended by the API
	// service.
	GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error)
	// Deletes a long-running operation. This method indicates that the client is
	// no longer interested in the operation result. It does not cancel the
	// operation. If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.
	DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	// Starts asynchronous cancellation on a long-running operation.  The server
	// makes a best effort to cancel the operation, but success is not
	// guaranteed.  If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.  Clients can use
	// [Operations.GetOperation][google.longrunning.Operations.GetOperation] or
	// other methods to check whether the cancellation succeeded or whether the
	// operation completed despite cancellation. On successful cancellation,
	// the operation is not deleted; instead, it becomes an operation with
	// an [Operation.error][google.longrunning.Operation.error] value with a [google.rpc.Status.code][google.rpc.Status.code] of 1,
	// corresponding to `Code.CANCELLED`.
	CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type operationsClient struct {
	cc *grpc.ClientConn
}

func NewOperationsClient(cc *grpc.ClientConn) OperationsClient {
	return &operationsClient{cc}
}

func (c *operationsClient) ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error) {
	out := new(ListOperationsResponse)
	err := grpc.Invoke(ctx, "/google.longrunning.Operations/ListOperations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := grpc.Invoke(ctx, "/google.longrunning.Operations/GetOperation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/google.longrunning.Operations/DeleteOperation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/google.longrunning.Operations/CancelOperation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Operations service

type OperationsServer interface {
	// Lists operations that match the specified filter in the request. If the
	// server doesn't support this method, it returns `UNIMPLEMENTED`.
	//
	// NOTE: the `name` binding below allows API services to override the binding
	// to use different resource name schemes, such as `users/*/operations`.
	ListOperations(context.Context, *ListOperationsRequest) (*ListOperationsResponse, error)
	// Gets the latest state of a long-running operation.  Clients can use this
	// method to poll the operation result at intervals as recommended by the API
	// service.
	GetOperation(context.Context, *GetOperationRequest) (*Operation, error)
	// Deletes a long-running operation. This method indicates that the client is
	// no longer interested in the operation result. It does not cancel the
	// operation. If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.
	DeleteOperation(context.Context, *DeleteOperationRequest) (*google_protobuf2.Empty, error)
	// Starts asynchronous cancellation on a long-running operation.  The server
	// makes a best effort to cancel the operation, but success is not
	// guaranteed.  If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.  Clients can use
	// [Operations.GetOperation][google.longrunning.Operations.GetOperation] or
	// other methods to check whether the cancellation succeeded or whether the
	// operation completed despite cancellation. On successful cancellation,
	// the operation is not deleted; instead, it becomes an operation with
	// an [Operation.error][google.longrunning.Operation.error] value with a [google.rpc.Status.code][google.rpc.Status.code] of 1,
	// corresponding to `Code.CANCELLED`.
	CancelOperation(context.Context, *CancelOperationRequest) (*google_protobuf2.Empty, error)
}

func RegisterOperationsServer(s *grpc.Server, srv OperationsServer) {
	s.RegisterService(&_Operations_serviceDesc, srv)
}

func _Operations_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.longrunning.Operations/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).ListOperations(ctx, req.(*ListOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_GetOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).GetOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.longrunning.Operations/GetOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).GetOperation(ctx, req.(*GetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_DeleteOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).DeleteOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.longrunning.Operations/DeleteOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).DeleteOperation(ctx, req.(*DeleteOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_CancelOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).CancelOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.longrunning.Operations/CancelOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).CancelOperation(ctx, req.(*CancelOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.longrunning.Operations",
	HandlerType: (*OperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOperations",
			Handler:    _Operations_ListOperations_Handler,
		},
		{
			MethodName: "GetOperation",
			Handler:    _Operations_GetOperation_Handler,
		},
		{
			MethodName: "DeleteOperation",
			Handler:    _Operations_DeleteOperation_Handler,
		},
		{
			MethodName: "CancelOperation",
			Handler:    _Operations_CancelOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/longrunning/operations.proto",
}

func init() { proto.RegisterFile("google/longrunning/operations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xd3, 0x24, 0x4a, 0xa6, 0x40, 0xa4, 0x85, 0xba, 0xc6, 0x25, 0x22, 0x32, 0x08, 0x52,
	0xab, 0xb2, 0x21, 0xdc, 0x8a, 0x72, 0x20, 0x80, 0xda, 0x43, 0x25, 0x22, 0x97, 0x13, 0x97, 0x6a,
	0x9b, 0x4e, 0x2d, 0x0b, 0x67, 0xd7, 0xac, 0x37, 0xd0, 0x16, 0x55, 0x11, 0x1c, 0x38, 0x71, 0xe3,
	0x2f, 0xf8, 0x19, 0x0e, 0xfc, 0x02, 0x1f, 0x82, 0xbc, 0x76, 0x62, 0x93, 0x3a, 0x28, 0xb7, 0xf5,
	0xcc, 0x9b, 0x79, 0xf3, 0xde, 0xce, 0x1a, 0x1e, 0xf8, 0x9c, 0xfb, 0x21, 0xba, 0x21, 0x67, 0xbe,
	0x98, 0x30, 0x16, 0x30, 0xdf, 0xe5, 0x11, 0x0a, 0x2a, 0x03, 0xce, 0x62, 0x27, 0x12, 0x5c, 0x72,
	0x42, 0x52, 0x90, 0x53, 0x00, 0x99, 0xf7, 0xb2, 0x42, 0x1a, 0x05, 0x2e, 0x65, 0x8c, 0xcb, 0x62,
	0x85, 0x79, 0x37, 0xcb, 0xaa, 0xaf, 0x93, 0xc9, 0x99, 0x4b, 0xd9, 0x45, 0x96, 0xda, 0x5e, 0x4c,
	0xe1, 0x38, 0x92, 0xb3, 0xe4, 0x56, 0x96, 0x14, 0xd1, 0xc8, 0x8d, 0x25, 0x95, 0x93, 0xac, 0xa1,
	0xf5, 0x4b, 0x83, 0xe6, 0x9b, 0xd9, 0x5c, 0x84, 0x40, 0x95, 0xd1, 0x31, 0x1a, 0x5a, 0x47, 0xeb,
	0x36, 0x3d, 0x75, 0x26, 0x4f, 0xa0, 0x31, 0x46, 0x49, 0x4f, 0xa9, 0xa4, 0x46, 0xa5, 0xa3, 0x75,
	0x37, 0x7a, 0x77, 0x9c, 0x6c, 0xee, 0x19, 0x95, 0xf3, 0x82, 0x5d, 0x78, 0x73, 0x54, 0xd2, 0xe5,
	0x94, 0x33, 0x34, 0xd6, 0x3b, 0x5a, 0xb7, 0xe1, 0xa9, 0x33, 0xb1, 0xa1, 0x86, 0x42, 0x70, 0x61,
	0x54, 0x55, 0x0b, 0x32, 0x6b, 0x21, 0xa2, 0x91, 0x73, 0xa4, 0x06, 0x3a, 0x58, 0xf3, 0x52, 0x08,
	0xe9, 0x41, 0x43, 0x60, 0x1c, 0x71, 0x16, 0xa3, 0x51, 0x5b, 0xce, 0x78, 0xb0, 0xe6, 0xcd, 0x71,
	0x83, 0x06, 0xd4, 0x05, 0xc6, 0x93, 0x50, 0x5a, 0x3b, 0x70, 0x7b, 0x1f, 0xe5, 0x5c, 0x93, 0x87,
	0x1f, 0x26, 0x18, 0xcb, 0x32, 0x69, 0xd6, 0x14, 0x36, 0x0f, 0x83, 0x38, 0xc7, 0xc6, 0x8b, 0xe0,
	0x6a, 0xc1, 0x07, 0x1d, 0xea, 0x67, 0x41, 0x28, 0x51, 0x64, 0x2d, 0xb2, 0x2f, 0xb2, 0x0d, 0xcd,
	0x88, 0xfa, 0x78, 0x1c, 0x07, 0x97, 0xa8, 0x0c, 0xaa, 0x79, 0x8d, 0x24, 0x70, 0x14, 0x5c, 0x22,
	0x69, 0x03, 0xa8, 0xa4, 0xe4, 0xef, 0x91, 0x29, 0x43, 0x9a, 0x9e, 0x82, 0xbf, 0x4d, 0x02, 0xd6,
	0x14, 0xf4, 0xc5, 0x01, 0x52, 0x3d, 0xa4, 0x0f, 0x90, 0xaf, 0x8b, 0xa1, 0x75, 0xd6, 0xbb, 0x1b,
	0xbd, 0xb6, 0x73, 0x7d, 0x5f, 0x9c, 0x5c, 0x68, 0xa1, 0x80, 0x3c, 0x82, 0x16, 0xc3, 0x73, 0x79,
	0x5c, 0x20, 0xaf, 0x28, 0xf2, 0x9b, 0x49, 0x78, 0x38, 0x1f, 0x60, 0x17, 0xf4, 0x97, 0x94, 0x8d,
	0x30, 0x5c, 0xc9, 0xaf, 0x5d, 0xd0, 0x5f, 0x61, 0x88, 0x12, 0x57, 0x41, 0xf7, 0xbe, 0x57, 0x01,
	0x72, 0x65, 0xe4, 0x9b, 0x06, 0xb7, 0xfe, 0x15, 0x4b, 0x76, 0xca, 0x04, 0x95, 0xde, 0x88, 0x69,
	0xaf, 0x02, 0x4d, 0xbd, 0xb3, 0xda, 0x5f, 0x7f, 0xff, 0xf9, 0x51, 0xd9, 0x22, 0x9b, 0xee, 0xc7,
	0xa7, 0xee, 0xe7, 0x64, 0x96, 0x7e, 0x6e, 0xcd, 0x15, 0x39, 0x87, 0x1b, 0xc5, 0x05, 0x21, 0x8f,
	0xcb, 0x5a, 0x97, 0xac, 0x90, 0xf9, 0x7f, 0xff, 0xad, 0x8e, 0xa2, 0x35, 0x89, 0x51, 0x46, 0xeb,
	0xda, 0xf6, 0x15, 0xf9, 0x04, 0xad, 0x05, 0xff, 0x48, 0xa9, 0xae, 0x72, 0x93, 0x4d, 0xfd, 0xda,
	0x2b, 0x78, 0x9d, 0x3c, 0xf1, 0x19, 0xb1, 0xbd, 0x9c, 0xf8, 0x8b, 0x06, 0xad, 0x85, 0x7b, 0x2e,
	0x67, 0x2e, 0x5f, 0x86, 0xa5, 0xcc, 0xb6, 0x62, 0x7e, 0x68, 0xdd, 0x5f, 0xc6, 0xbc, 0x37, 0x52,
	0x0d, 0xf7, 0x34, 0x7b, 0x30, 0x05, 0x7d, 0xc4, 0xc7, 0x25, 0xa4, 0x83, 0x56, 0x7e, 0x87, 0xc3,
	0xa4, 0xff, 0x50, 0x7b, 0xd7, 0xcf, 0x60, 0x3e, 0x0f, 0x29, 0xf3, 0x1d, 0x2e, 0x7c, 0xd7, 0x47,
	0xa6, 0xd8, 0xdd, 0x34, 0x45, 0xa3, 0x20, 0x2e, 0xfe, 0x5d, 0x9f, 0x17, 0xce, 0x3f, 0x2b, 0x64,
	0x3f, 0xad, 0x3f, 0xe4, 0xcc, 0xf7, 0xd2, 0xe0, 0x49, 0x5d, 0x95, 0x3f, 0xfb, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x69, 0x4c, 0xa6, 0x3e, 0x9b, 0x05, 0x00, 0x00,
}