// Code generated by protoc-gen-go.
// source: server/pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	server/pps/pps.proto

It has these top-level messages:
	StartPodRequest
	StartPodResponse
	FinishPodRequest
	FinishPodResponse
	ContinuePodRequest
	ContinuePodResponse
*/
package pps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/protobuf"
import fuse "github.com/pachyderm/pachyderm/src/server/pfs/fuse"
import _ "github.com/pachyderm/pachyderm/src/client/pfs"
import pps1 "github.com/pachyderm/pachyderm/src/client/pps"

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

type StartPodRequest struct {
	Job     *pps1.Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	PodName string    `protobuf:"bytes,2,opt,name=pod_name,json=podName" json:"pod_name,omitempty"`
}

func (m *StartPodRequest) Reset()                    { *m = StartPodRequest{} }
func (m *StartPodRequest) String() string            { return proto.CompactTextString(m) }
func (*StartPodRequest) ProtoMessage()               {}
func (*StartPodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartPodRequest) GetJob() *pps1.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type StartPodResponse struct {
	ChunkID      string              `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId" json:"chunk_id,omitempty"`
	Transform    *pps1.Transform     `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	CommitMounts []*fuse.CommitMount `protobuf:"bytes,3,rep,name=commit_mounts,json=commitMounts" json:"commit_mounts,omitempty"`
}

func (m *StartPodResponse) Reset()                    { *m = StartPodResponse{} }
func (m *StartPodResponse) String() string            { return proto.CompactTextString(m) }
func (*StartPodResponse) ProtoMessage()               {}
func (*StartPodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StartPodResponse) GetTransform() *pps1.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *StartPodResponse) GetCommitMounts() []*fuse.CommitMount {
	if m != nil {
		return m.CommitMounts
	}
	return nil
}

type FinishPodRequest struct {
	ChunkID string `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId" json:"chunk_id,omitempty"`
	PodName string `protobuf:"bytes,2,opt,name=pod_name,json=podName" json:"pod_name,omitempty"`
	Success bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
}

func (m *FinishPodRequest) Reset()                    { *m = FinishPodRequest{} }
func (m *FinishPodRequest) String() string            { return proto.CompactTextString(m) }
func (*FinishPodRequest) ProtoMessage()               {}
func (*FinishPodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type FinishPodResponse struct {
	// If fail is true, the pod is expected to exit with a non-zero code
	// so that k8s knows to reschedule the pod.
	Fail bool `protobuf:"varint,1,opt,name=fail" json:"fail,omitempty"`
}

func (m *FinishPodResponse) Reset()                    { *m = FinishPodResponse{} }
func (m *FinishPodResponse) String() string            { return proto.CompactTextString(m) }
func (*FinishPodResponse) ProtoMessage()               {}
func (*FinishPodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ContinuePodRequest struct {
	ChunkID string `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId" json:"chunk_id,omitempty"`
	PodName string `protobuf:"bytes,2,opt,name=pod_name,json=podName" json:"pod_name,omitempty"`
}

func (m *ContinuePodRequest) Reset()                    { *m = ContinuePodRequest{} }
func (m *ContinuePodRequest) String() string            { return proto.CompactTextString(m) }
func (*ContinuePodRequest) ProtoMessage()               {}
func (*ContinuePodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ContinuePodResponse struct {
	// if exit is true, the pod is expected to abandon its work and exit.
	// this happens if the chunk that the pod is working on has been assigned
	// to another pod.
	Exit bool `protobuf:"varint,1,opt,name=exit" json:"exit,omitempty"`
}

func (m *ContinuePodResponse) Reset()                    { *m = ContinuePodResponse{} }
func (m *ContinuePodResponse) String() string            { return proto.CompactTextString(m) }
func (*ContinuePodResponse) ProtoMessage()               {}
func (*ContinuePodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*StartPodRequest)(nil), "pps.StartPodRequest")
	proto.RegisterType((*StartPodResponse)(nil), "pps.StartPodResponse")
	proto.RegisterType((*FinishPodRequest)(nil), "pps.FinishPodRequest")
	proto.RegisterType((*FinishPodResponse)(nil), "pps.FinishPodResponse")
	proto.RegisterType((*ContinuePodRequest)(nil), "pps.ContinuePodRequest")
	proto.RegisterType((*ContinuePodResponse)(nil), "pps.ContinuePodResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for InternalPodAPI service

type InternalPodAPIClient interface {
	StartPod(ctx context.Context, in *StartPodRequest, opts ...grpc.CallOption) (*StartPodResponse, error)
	ContinuePod(ctx context.Context, in *ContinuePodRequest, opts ...grpc.CallOption) (*ContinuePodResponse, error)
	FinishPod(ctx context.Context, in *FinishPodRequest, opts ...grpc.CallOption) (*FinishPodResponse, error)
}

type internalPodAPIClient struct {
	cc *grpc.ClientConn
}

func NewInternalPodAPIClient(cc *grpc.ClientConn) InternalPodAPIClient {
	return &internalPodAPIClient{cc}
}

func (c *internalPodAPIClient) StartPod(ctx context.Context, in *StartPodRequest, opts ...grpc.CallOption) (*StartPodResponse, error) {
	out := new(StartPodResponse)
	err := grpc.Invoke(ctx, "/pps.InternalPodAPI/StartPod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalPodAPIClient) ContinuePod(ctx context.Context, in *ContinuePodRequest, opts ...grpc.CallOption) (*ContinuePodResponse, error) {
	out := new(ContinuePodResponse)
	err := grpc.Invoke(ctx, "/pps.InternalPodAPI/ContinuePod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalPodAPIClient) FinishPod(ctx context.Context, in *FinishPodRequest, opts ...grpc.CallOption) (*FinishPodResponse, error) {
	out := new(FinishPodResponse)
	err := grpc.Invoke(ctx, "/pps.InternalPodAPI/FinishPod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InternalPodAPI service

type InternalPodAPIServer interface {
	StartPod(context.Context, *StartPodRequest) (*StartPodResponse, error)
	ContinuePod(context.Context, *ContinuePodRequest) (*ContinuePodResponse, error)
	FinishPod(context.Context, *FinishPodRequest) (*FinishPodResponse, error)
}

func RegisterInternalPodAPIServer(s *grpc.Server, srv InternalPodAPIServer) {
	s.RegisterService(&_InternalPodAPI_serviceDesc, srv)
}

func _InternalPodAPI_StartPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPodAPIServer).StartPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pps.InternalPodAPI/StartPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPodAPIServer).StartPod(ctx, req.(*StartPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalPodAPI_ContinuePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContinuePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPodAPIServer).ContinuePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pps.InternalPodAPI/ContinuePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPodAPIServer).ContinuePod(ctx, req.(*ContinuePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalPodAPI_FinishPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPodAPIServer).FinishPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pps.InternalPodAPI/FinishPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPodAPIServer).FinishPod(ctx, req.(*FinishPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalPodAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pps.InternalPodAPI",
	HandlerType: (*InternalPodAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartPod",
			Handler:    _InternalPodAPI_StartPod_Handler,
		},
		{
			MethodName: "ContinuePod",
			Handler:    _InternalPodAPI_ContinuePod_Handler,
		},
		{
			MethodName: "FinishPod",
			Handler:    _InternalPodAPI_FinishPod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/pps/pps.proto",
}

func init() { proto.RegisterFile("server/pps/pps.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0x5d, 0xcb, 0xd3, 0x30,
	0x18, 0xb5, 0x56, 0x5d, 0xfb, 0x54, 0x5f, 0xdf, 0x37, 0x4e, 0xad, 0xf5, 0x66, 0xf4, 0xc6, 0x0a,
	0xb2, 0x42, 0x05, 0x41, 0xf0, 0xc6, 0x0d, 0xc4, 0x0d, 0x94, 0x11, 0xbd, 0xf2, 0x66, 0xf4, 0x23,
	0xdd, 0xa2, 0x6d, 0x12, 0x9b, 0x54, 0xf4, 0x77, 0xf8, 0xbb, 0xfc, 0x4f, 0x92, 0x6c, 0x5d, 0x6b,
	0x87, 0xde, 0x78, 0xd1, 0x92, 0xe7, 0x3c, 0x27, 0xe7, 0xe4, 0xe4, 0x03, 0xa6, 0x92, 0x34, 0xdf,
	0x48, 0x13, 0x0b, 0x21, 0xf5, 0x37, 0x17, 0x0d, 0x57, 0x1c, 0xd9, 0x42, 0xc8, 0xe0, 0xf1, 0x8e,
	0xf3, 0x5d, 0x45, 0x62, 0x03, 0x65, 0x6d, 0x19, 0x93, 0x5a, 0xa8, 0x1f, 0x07, 0x46, 0x10, 0x74,
	0xf3, 0x4a, 0x19, 0x97, 0xad, 0x24, 0xe6, 0x77, 0xec, 0x4d, 0xf3, 0x8a, 0x12, 0xa6, 0x4c, 0x4f,
	0x94, 0x72, 0x8c, 0x0e, 0x9d, 0xc2, 0xb7, 0x70, 0xf7, 0x83, 0x4a, 0x1b, 0xb5, 0xe1, 0x05, 0x26,
	0x5f, 0x5b, 0x22, 0x15, 0x0a, 0xc0, 0xfe, 0xcc, 0x33, 0xdf, 0x9a, 0x59, 0x91, 0x97, 0x38, 0x73,
	0xcd, 0x5d, 0xf3, 0x0c, 0x6b, 0x10, 0x3d, 0x02, 0x47, 0xf0, 0x62, 0xcb, 0xd2, 0x9a, 0xf8, 0xd7,
	0x67, 0x56, 0xe4, 0xe2, 0x89, 0xe0, 0xc5, 0xfb, 0xb4, 0x26, 0xe1, 0x4f, 0x0b, 0x2e, 0x7b, 0x29,
	0x29, 0x38, 0x93, 0x44, 0xf3, 0xf3, 0x7d, 0xcb, 0xbe, 0x6c, 0x69, 0x61, 0x04, 0x5d, 0x3c, 0x31,
	0xf5, 0xaa, 0x40, 0xcf, 0xc0, 0x55, 0x4d, 0xca, 0x64, 0xc9, 0x9b, 0xda, 0x68, 0x79, 0xc9, 0x85,
	0x31, 0xfb, 0xd8, 0xa1, 0xb8, 0x27, 0xa0, 0x17, 0x70, 0x27, 0xe7, 0x75, 0x4d, 0xd5, 0xb6, 0xe6,
	0x2d, 0x53, 0xd2, 0xb7, 0x67, 0x76, 0xe4, 0x25, 0x57, 0x73, 0x93, 0x7b, 0x69, 0x5a, 0xef, 0x74,
	0x07, 0xdf, 0xce, 0xfb, 0x42, 0x86, 0x19, 0x5c, 0xbe, 0xa1, 0x8c, 0xca, 0xfd, 0x20, 0xe0, 0x3f,
	0x16, 0xf5, 0xf7, 0x7c, 0xc8, 0x87, 0x89, 0x6c, 0xf3, 0x9c, 0x48, 0xed, 0x6d, 0x45, 0x0e, 0xee,
	0xca, 0xf0, 0x09, 0x5c, 0x0d, 0x3c, 0x8e, 0xc9, 0x11, 0xdc, 0x28, 0x53, 0x5a, 0x19, 0x03, 0x07,
	0x9b, 0x71, 0xb8, 0x06, 0xb4, 0xe4, 0x4c, 0x51, 0xd6, 0x92, 0xff, 0x5d, 0x4e, 0xf8, 0x14, 0xee,
	0xfd, 0xa1, 0xd5, 0xdb, 0x92, 0xef, 0x54, 0x75, 0xb6, 0x7a, 0x9c, 0xfc, 0xb2, 0xe0, 0x62, 0xc5,
	0x14, 0x69, 0x58, 0x5a, 0x6d, 0x78, 0xf1, 0x7a, 0xb3, 0x42, 0x2f, 0xc1, 0xe9, 0xce, 0x0a, 0x4d,
	0xcd, 0xae, 0x8f, 0x6e, 0x41, 0x70, 0x7f, 0x84, 0x1e, 0xf4, 0xc3, 0x6b, 0x68, 0x01, 0xde, 0xc0,
	0x18, 0x3d, 0x34, 0xbc, 0xf3, 0x58, 0x81, 0x7f, 0xde, 0x38, 0x69, 0xbc, 0x02, 0xf7, 0xb4, 0x63,
	0xe8, 0xe0, 0x34, 0x3e, 0xa5, 0xe0, 0xc1, 0x18, 0xee, 0x66, 0x2f, 0x6e, 0x7e, 0xd2, 0xef, 0x23,
	0xbb, 0x65, 0x6e, 0xf0, 0xf3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xa9, 0x80, 0x41, 0x43,
	0x03, 0x00, 0x00,
}
