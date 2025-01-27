// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/task/task.proto

package project_test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TaskService_ListTasks_FullMethodName   = "/TaskService/ListTasks"
	TaskService_GetTaskByID_FullMethodName = "/TaskService/GetTaskByID"
	TaskService_CreateTask_FullMethodName  = "/TaskService/CreateTask"
	TaskService_UpdateTask_FullMethodName  = "/TaskService/UpdateTask"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	ListTasks(ctx context.Context, opts ...grpc.CallOption) (TaskService_ListTasksClient, error)
	GetTaskByID(ctx context.Context, opts ...grpc.CallOption) (TaskService_GetTaskByIDClient, error)
	CreateTask(ctx context.Context, opts ...grpc.CallOption) (TaskService_CreateTaskClient, error)
	UpdateTask(ctx context.Context, opts ...grpc.CallOption) (TaskService_UpdateTaskClient, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) ListTasks(ctx context.Context, opts ...grpc.CallOption) (TaskService_ListTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[0], TaskService_ListTasks_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceListTasksClient{stream}
	return x, nil
}

type TaskService_ListTasksClient interface {
	Send(*TaskReq) error
	CloseAndRecv() (*TaskArr, error)
	grpc.ClientStream
}

type taskServiceListTasksClient struct {
	grpc.ClientStream
}

func (x *taskServiceListTasksClient) Send(m *TaskReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskServiceListTasksClient) CloseAndRecv() (*TaskArr, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TaskArr)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskServiceClient) GetTaskByID(ctx context.Context, opts ...grpc.CallOption) (TaskService_GetTaskByIDClient, error) {
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[1], TaskService_GetTaskByID_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceGetTaskByIDClient{stream}
	return x, nil
}

type TaskService_GetTaskByIDClient interface {
	Send(*TaskReq) error
	CloseAndRecv() (*Task, error)
	grpc.ClientStream
}

type taskServiceGetTaskByIDClient struct {
	grpc.ClientStream
}

func (x *taskServiceGetTaskByIDClient) Send(m *TaskReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskServiceGetTaskByIDClient) CloseAndRecv() (*Task, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskServiceClient) CreateTask(ctx context.Context, opts ...grpc.CallOption) (TaskService_CreateTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[2], TaskService_CreateTask_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceCreateTaskClient{stream}
	return x, nil
}

type TaskService_CreateTaskClient interface {
	Send(*TaskReq) error
	CloseAndRecv() (*Task, error)
	grpc.ClientStream
}

type taskServiceCreateTaskClient struct {
	grpc.ClientStream
}

func (x *taskServiceCreateTaskClient) Send(m *TaskReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskServiceCreateTaskClient) CloseAndRecv() (*Task, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskServiceClient) UpdateTask(ctx context.Context, opts ...grpc.CallOption) (TaskService_UpdateTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[3], TaskService_UpdateTask_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceUpdateTaskClient{stream}
	return x, nil
}

type TaskService_UpdateTaskClient interface {
	Send(*TaskReq) error
	CloseAndRecv() (*Task, error)
	grpc.ClientStream
}

type taskServiceUpdateTaskClient struct {
	grpc.ClientStream
}

func (x *taskServiceUpdateTaskClient) Send(m *TaskReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskServiceUpdateTaskClient) CloseAndRecv() (*Task, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	ListTasks(TaskService_ListTasksServer) error
	GetTaskByID(TaskService_GetTaskByIDServer) error
	CreateTask(TaskService_CreateTaskServer) error
	UpdateTask(TaskService_UpdateTaskServer) error
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) ListTasks(TaskService_ListTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedTaskServiceServer) GetTaskByID(TaskService_GetTaskByIDServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTaskByID not implemented")
}
func (UnimplementedTaskServiceServer) CreateTask(TaskService_CreateTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskServiceServer) UpdateTask(TaskService_UpdateTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_ListTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).ListTasks(&taskServiceListTasksServer{stream})
}

type TaskService_ListTasksServer interface {
	SendAndClose(*TaskArr) error
	Recv() (*TaskReq, error)
	grpc.ServerStream
}

type taskServiceListTasksServer struct {
	grpc.ServerStream
}

func (x *taskServiceListTasksServer) SendAndClose(m *TaskArr) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskServiceListTasksServer) Recv() (*TaskReq, error) {
	m := new(TaskReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TaskService_GetTaskByID_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).GetTaskByID(&taskServiceGetTaskByIDServer{stream})
}

type TaskService_GetTaskByIDServer interface {
	SendAndClose(*Task) error
	Recv() (*TaskReq, error)
	grpc.ServerStream
}

type taskServiceGetTaskByIDServer struct {
	grpc.ServerStream
}

func (x *taskServiceGetTaskByIDServer) SendAndClose(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskServiceGetTaskByIDServer) Recv() (*TaskReq, error) {
	m := new(TaskReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TaskService_CreateTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).CreateTask(&taskServiceCreateTaskServer{stream})
}

type TaskService_CreateTaskServer interface {
	SendAndClose(*Task) error
	Recv() (*TaskReq, error)
	grpc.ServerStream
}

type taskServiceCreateTaskServer struct {
	grpc.ServerStream
}

func (x *taskServiceCreateTaskServer) SendAndClose(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskServiceCreateTaskServer) Recv() (*TaskReq, error) {
	m := new(TaskReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TaskService_UpdateTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).UpdateTask(&taskServiceUpdateTaskServer{stream})
}

type TaskService_UpdateTaskServer interface {
	SendAndClose(*Task) error
	Recv() (*TaskReq, error)
	grpc.ServerStream
}

type taskServiceUpdateTaskServer struct {
	grpc.ServerStream
}

func (x *taskServiceUpdateTaskServer) SendAndClose(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskServiceUpdateTaskServer) Recv() (*TaskReq, error) {
	m := new(TaskReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTasks",
			Handler:       _TaskService_ListTasks_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetTaskByID",
			Handler:       _TaskService_GetTaskByID_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateTask",
			Handler:       _TaskService_CreateTask_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateTask",
			Handler:       _TaskService_UpdateTask_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/task/task.proto",
}
