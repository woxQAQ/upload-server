// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: v1/task_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TaskService_GetTaskDetail_FullMethodName  = "/woxqaq.v1.TaskService/GetTaskDetail"
	TaskService_ListTaskDetail_FullMethodName = "/woxqaq.v1.TaskService/ListTaskDetail"
	TaskService_SubmitTask_FullMethodName     = "/woxqaq.v1.TaskService/SubmitTask"
	TaskService_Approve_FullMethodName        = "/woxqaq.v1.TaskService/Approve"
	TaskService_Reject_FullMethodName         = "/woxqaq.v1.TaskService/Reject"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	GetTaskDetail(ctx context.Context, in *GetTaskDetailRequest, opts ...grpc.CallOption) (*GetTaskDetailResponse, error)
	ListTaskDetail(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*ListTaskResponse, error)
	SubmitTask(ctx context.Context, in *SubmitTaskRequest, opts ...grpc.CallOption) (*SubmitTaskResponse, error)
	Approve(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error)
	Reject(ctx context.Context, in *RejectRequest, opts ...grpc.CallOption) (*RejectResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) GetTaskDetail(ctx context.Context, in *GetTaskDetailRequest, opts ...grpc.CallOption) (*GetTaskDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_GetTaskDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListTaskDetail(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*ListTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_ListTaskDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) SubmitTask(ctx context.Context, in *SubmitTaskRequest, opts ...grpc.CallOption) (*SubmitTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_SubmitTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Approve(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApproveResponse)
	err := c.cc.Invoke(ctx, TaskService_Approve_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Reject(ctx context.Context, in *RejectRequest, opts ...grpc.CallOption) (*RejectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RejectResponse)
	err := c.cc.Invoke(ctx, TaskService_Reject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility.
type TaskServiceServer interface {
	GetTaskDetail(context.Context, *GetTaskDetailRequest) (*GetTaskDetailResponse, error)
	ListTaskDetail(context.Context, *ListTaskRequest) (*ListTaskResponse, error)
	SubmitTask(context.Context, *SubmitTaskRequest) (*SubmitTaskResponse, error)
	Approve(context.Context, *ApproveRequest) (*ApproveResponse, error)
	Reject(context.Context, *RejectRequest) (*RejectResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskServiceServer struct{}

func (UnimplementedTaskServiceServer) GetTaskDetail(context.Context, *GetTaskDetailRequest) (*GetTaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskDetail not implemented")
}
func (UnimplementedTaskServiceServer) ListTaskDetail(context.Context, *ListTaskRequest) (*ListTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTaskDetail not implemented")
}
func (UnimplementedTaskServiceServer) SubmitTask(context.Context, *SubmitTaskRequest) (*SubmitTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTask not implemented")
}
func (UnimplementedTaskServiceServer) Approve(context.Context, *ApproveRequest) (*ApproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Approve not implemented")
}
func (UnimplementedTaskServiceServer) Reject(context.Context, *RejectRequest) (*RejectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reject not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}
func (UnimplementedTaskServiceServer) testEmbeddedByValue()                     {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	// If the following call pancis, it indicates UnimplementedTaskServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_GetTaskDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTaskDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetTaskDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTaskDetail(ctx, req.(*GetTaskDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListTaskDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListTaskDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_ListTaskDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListTaskDetail(ctx, req.(*ListTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_SubmitTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).SubmitTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_SubmitTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).SubmitTask(ctx, req.(*SubmitTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Approve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Approve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_Approve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Approve(ctx, req.(*ApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Reject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Reject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_Reject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Reject(ctx, req.(*RejectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "woxqaq.v1.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTaskDetail",
			Handler:    _TaskService_GetTaskDetail_Handler,
		},
		{
			MethodName: "ListTaskDetail",
			Handler:    _TaskService_ListTaskDetail_Handler,
		},
		{
			MethodName: "SubmitTask",
			Handler:    _TaskService_SubmitTask_Handler,
		},
		{
			MethodName: "Approve",
			Handler:    _TaskService_Approve_Handler,
		},
		{
			MethodName: "Reject",
			Handler:    _TaskService_Reject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/task_service.proto",
}
