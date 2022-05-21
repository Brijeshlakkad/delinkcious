// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: news.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_LINK_ADDED   EventType = 0
	EventType_LINK_UPDATED EventType = 1
	EventType_LINK_DELETED EventType = 2
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "LINK_ADDED",
		1: "LINK_UPDATED",
		2: "LINK_DELETED",
	}
	EventType_value = map[string]int32{
		"LINK_ADDED":   0,
		"LINK_UPDATED": 1,
		"LINK_DELETED": 2,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_news_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_news_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{0}
}

type GetNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	StartToken string `protobuf:"bytes,2,opt,name=startToken,proto3" json:"startToken,omitempty"`
}

func (x *GetNewsRequest) Reset() {
	*x = GetNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsRequest) ProtoMessage() {}

func (x *GetNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsRequest.ProtoReflect.Descriptor instead.
func (*GetNewsRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{0}
}

func (x *GetNewsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetNewsRequest) GetStartToken() string {
	if x != nil {
		return x.StartToken
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType EventType              `protobuf:"varint,1,opt,name=eventType,proto3,enum=pb.EventType" json:"eventType,omitempty"`
	Username  string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Url       string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_LINK_ADDED
}

func (x *Event) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Event) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Event) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type GetNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events    []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	NextToken string   `protobuf:"bytes,2,opt,name=nextToken,proto3" json:"nextToken,omitempty"`
	Err       string   `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetNewsResponse) Reset() {
	*x = GetNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsResponse) ProtoMessage() {}

func (x *GetNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsResponse.ProtoReflect.Descriptor instead.
func (*GetNewsResponse) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{2}
}

func (x *GetNewsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *GetNewsResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

func (x *GetNewsResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_news_proto protoreflect.FileDescriptor

var file_news_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x9c, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x64,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x2a, 0x3f, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0x3c, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x34, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_news_proto_rawDescOnce sync.Once
	file_news_proto_rawDescData = file_news_proto_rawDesc
)

func file_news_proto_rawDescGZIP() []byte {
	file_news_proto_rawDescOnce.Do(func() {
		file_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_news_proto_rawDescData)
	})
	return file_news_proto_rawDescData
}

var file_news_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_news_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_news_proto_goTypes = []interface{}{
	(EventType)(0),                // 0: pb.EventType
	(*GetNewsRequest)(nil),        // 1: pb.GetNewsRequest
	(*Event)(nil),                 // 2: pb.Event
	(*GetNewsResponse)(nil),       // 3: pb.GetNewsResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_news_proto_depIdxs = []int32{
	0, // 0: pb.Event.eventType:type_name -> pb.EventType
	4, // 1: pb.Event.timestamp:type_name -> google.protobuf.Timestamp
	2, // 2: pb.GetNewsResponse.events:type_name -> pb.Event
	1, // 3: pb.News.GetNews:input_type -> pb.GetNewsRequest
	3, // 4: pb.News.GetNews:output_type -> pb.GetNewsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_news_proto_init() }
func file_news_proto_init() {
	if File_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsRequest); i {
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
		file_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsResponse); i {
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
			RawDescriptor: file_news_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_news_proto_goTypes,
		DependencyIndexes: file_news_proto_depIdxs,
		EnumInfos:         file_news_proto_enumTypes,
		MessageInfos:      file_news_proto_msgTypes,
	}.Build()
	File_news_proto = out.File
	file_news_proto_rawDesc = nil
	file_news_proto_goTypes = nil
	file_news_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NewsClient is the client API for News service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NewsClient interface {
	GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsResponse, error)
}

type newsClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsClient(cc grpc.ClientConnInterface) NewsClient {
	return &newsClient{cc}
}

func (c *newsClient) GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsResponse, error) {
	out := new(GetNewsResponse)
	err := c.cc.Invoke(ctx, "/pb.News/GetNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServer is the server API for News service.
type NewsServer interface {
	GetNews(context.Context, *GetNewsRequest) (*GetNewsResponse, error)
}

// UnimplementedNewsServer can be embedded to have forward compatible implementations.
type UnimplementedNewsServer struct {
}

func (*UnimplementedNewsServer) GetNews(context.Context, *GetNewsRequest) (*GetNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}

func RegisterNewsServer(s *grpc.Server, srv NewsServer) {
	s.RegisterService(&_News_serviceDesc, srv)
}

func _News_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.News/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetNews(ctx, req.(*GetNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _News_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.News",
	HandlerType: (*NewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNews",
			Handler:    _News_GetNews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news.proto",
}