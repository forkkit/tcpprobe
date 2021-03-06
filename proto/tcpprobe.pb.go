// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: tcpprobe.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr     string            `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Interval string            `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Labels   map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcpprobe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_tcpprobe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_tcpprobe_proto_rawDescGZIP(), []int{0}
}

func (x *Target) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Target) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

func (x *Target) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcpprobe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_tcpprobe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_tcpprobe_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics *_struct.Struct `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcpprobe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_tcpprobe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_tcpprobe_proto_rawDescGZIP(), []int{2}
}

func (x *Stats) GetMetrics() *_struct.Struct {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_tcpprobe_proto protoreflect.FileDescriptor

var file_tcpprobe_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x63, 0x70, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x30, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x31, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x32, 0x81, 0x01, 0x0a, 0x08, 0x54, 0x43, 0x50, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12,
	0x25, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x0e,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x24, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x22, 0x00, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x68, 0x72, 0x64, 0x61, 0x64, 0x72, 0x61, 0x64, 0x2f,
	0x74, 0x63, 0x70, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tcpprobe_proto_rawDescOnce sync.Once
	file_tcpprobe_proto_rawDescData = file_tcpprobe_proto_rawDesc
)

func file_tcpprobe_proto_rawDescGZIP() []byte {
	file_tcpprobe_proto_rawDescOnce.Do(func() {
		file_tcpprobe_proto_rawDescData = protoimpl.X.CompressGZIP(file_tcpprobe_proto_rawDescData)
	})
	return file_tcpprobe_proto_rawDescData
}

var file_tcpprobe_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tcpprobe_proto_goTypes = []interface{}{
	(*Target)(nil),         // 0: main.Target
	(*Response)(nil),       // 1: main.Response
	(*Stats)(nil),          // 2: main.Stats
	nil,                    // 3: main.Target.LabelsEntry
	(*_struct.Struct)(nil), // 4: google.protobuf.Struct
}
var file_tcpprobe_proto_depIdxs = []int32{
	3, // 0: main.Target.labels:type_name -> main.Target.LabelsEntry
	4, // 1: main.Stats.metrics:type_name -> google.protobuf.Struct
	0, // 2: main.TCPProbe.Add:input_type -> main.Target
	0, // 3: main.TCPProbe.Delete:input_type -> main.Target
	0, // 4: main.TCPProbe.Get:input_type -> main.Target
	1, // 5: main.TCPProbe.Add:output_type -> main.Response
	1, // 6: main.TCPProbe.Delete:output_type -> main.Response
	2, // 7: main.TCPProbe.Get:output_type -> main.Stats
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tcpprobe_proto_init() }
func file_tcpprobe_proto_init() {
	if File_tcpprobe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tcpprobe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
		file_tcpprobe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_tcpprobe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
			RawDescriptor: file_tcpprobe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tcpprobe_proto_goTypes,
		DependencyIndexes: file_tcpprobe_proto_depIdxs,
		MessageInfos:      file_tcpprobe_proto_msgTypes,
	}.Build()
	File_tcpprobe_proto = out.File
	file_tcpprobe_proto_rawDesc = nil
	file_tcpprobe_proto_goTypes = nil
	file_tcpprobe_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TCPProbeClient is the client API for TCPProbe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TCPProbeClient interface {
	Add(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *Target, opts ...grpc.CallOption) (TCPProbe_GetClient, error)
}

type tCPProbeClient struct {
	cc grpc.ClientConnInterface
}

func NewTCPProbeClient(cc grpc.ClientConnInterface) TCPProbeClient {
	return &tCPProbeClient{cc}
}

func (c *tCPProbeClient) Add(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/main.TCPProbe/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tCPProbeClient) Delete(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/main.TCPProbe/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tCPProbeClient) Get(ctx context.Context, in *Target, opts ...grpc.CallOption) (TCPProbe_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TCPProbe_serviceDesc.Streams[0], "/main.TCPProbe/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &tCPProbeGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TCPProbe_GetClient interface {
	Recv() (*Stats, error)
	grpc.ClientStream
}

type tCPProbeGetClient struct {
	grpc.ClientStream
}

func (x *tCPProbeGetClient) Recv() (*Stats, error) {
	m := new(Stats)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TCPProbeServer is the server API for TCPProbe service.
type TCPProbeServer interface {
	Add(context.Context, *Target) (*Response, error)
	Delete(context.Context, *Target) (*Response, error)
	Get(*Target, TCPProbe_GetServer) error
}

// UnimplementedTCPProbeServer can be embedded to have forward compatible implementations.
type UnimplementedTCPProbeServer struct {
}

func (*UnimplementedTCPProbeServer) Add(context.Context, *Target) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedTCPProbeServer) Delete(context.Context, *Target) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedTCPProbeServer) Get(*Target, TCPProbe_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterTCPProbeServer(s *grpc.Server, srv TCPProbeServer) {
	s.RegisterService(&_TCPProbe_serviceDesc, srv)
}

func _TCPProbe_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TCPProbeServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TCPProbe/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TCPProbeServer).Add(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _TCPProbe_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TCPProbeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TCPProbe/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TCPProbeServer).Delete(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _TCPProbe_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Target)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TCPProbeServer).Get(m, &tCPProbeGetServer{stream})
}

type TCPProbe_GetServer interface {
	Send(*Stats) error
	grpc.ServerStream
}

type tCPProbeGetServer struct {
	grpc.ServerStream
}

func (x *tCPProbeGetServer) Send(m *Stats) error {
	return x.ServerStream.SendMsg(m)
}

var _TCPProbe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.TCPProbe",
	HandlerType: (*TCPProbeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _TCPProbe_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TCPProbe_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _TCPProbe_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tcpprobe.proto",
}
