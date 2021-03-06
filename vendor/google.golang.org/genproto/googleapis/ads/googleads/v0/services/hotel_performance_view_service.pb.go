// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/hotel_performance_view_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Request message for [HotelPerformanceViewService.GetHotelPerformanceView][google.ads.googleads.v0.services.HotelPerformanceViewService.GetHotelPerformanceView].
type GetHotelPerformanceViewRequest struct {
	// Resource name of the Hotel Performance View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHotelPerformanceViewRequest) Reset()         { *m = GetHotelPerformanceViewRequest{} }
func (m *GetHotelPerformanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetHotelPerformanceViewRequest) ProtoMessage()    {}
func (*GetHotelPerformanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hotel_performance_view_service_91d056d973d5a9fa, []int{0}
}
func (m *GetHotelPerformanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHotelPerformanceViewRequest.Unmarshal(m, b)
}
func (m *GetHotelPerformanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHotelPerformanceViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetHotelPerformanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHotelPerformanceViewRequest.Merge(dst, src)
}
func (m *GetHotelPerformanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetHotelPerformanceViewRequest.Size(m)
}
func (m *GetHotelPerformanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHotelPerformanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHotelPerformanceViewRequest proto.InternalMessageInfo

func (m *GetHotelPerformanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHotelPerformanceViewRequest)(nil), "google.ads.googleads.v0.services.GetHotelPerformanceViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HotelPerformanceViewServiceClient is the client API for HotelPerformanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotelPerformanceViewServiceClient interface {
	// Returns the requested Hotel Performance View in full detail.
	GetHotelPerformanceView(ctx context.Context, in *GetHotelPerformanceViewRequest, opts ...grpc.CallOption) (*resources.HotelPerformanceView, error)
}

type hotelPerformanceViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewHotelPerformanceViewServiceClient(cc *grpc.ClientConn) HotelPerformanceViewServiceClient {
	return &hotelPerformanceViewServiceClient{cc}
}

func (c *hotelPerformanceViewServiceClient) GetHotelPerformanceView(ctx context.Context, in *GetHotelPerformanceViewRequest, opts ...grpc.CallOption) (*resources.HotelPerformanceView, error) {
	out := new(resources.HotelPerformanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.HotelPerformanceViewService/GetHotelPerformanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelPerformanceViewServiceServer is the server API for HotelPerformanceViewService service.
type HotelPerformanceViewServiceServer interface {
	// Returns the requested Hotel Performance View in full detail.
	GetHotelPerformanceView(context.Context, *GetHotelPerformanceViewRequest) (*resources.HotelPerformanceView, error)
}

func RegisterHotelPerformanceViewServiceServer(s *grpc.Server, srv HotelPerformanceViewServiceServer) {
	s.RegisterService(&_HotelPerformanceViewService_serviceDesc, srv)
}

func _HotelPerformanceViewService_GetHotelPerformanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelPerformanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelPerformanceViewServiceServer).GetHotelPerformanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.HotelPerformanceViewService/GetHotelPerformanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelPerformanceViewServiceServer).GetHotelPerformanceView(ctx, req.(*GetHotelPerformanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotelPerformanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.HotelPerformanceViewService",
	HandlerType: (*HotelPerformanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotelPerformanceView",
			Handler:    _HotelPerformanceViewService_GetHotelPerformanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/hotel_performance_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/hotel_performance_view_service.proto", fileDescriptor_hotel_performance_view_service_91d056d973d5a9fa)
}

var fileDescriptor_hotel_performance_view_service_91d056d973d5a9fa = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xe2, 0xd4,
	0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0x62, 0xfd, 0x8c, 0xfc, 0x92, 0xd4, 0x9c, 0xf8, 0x82, 0xd4, 0xa2,
	0xb4, 0xfc, 0xa2, 0xdc, 0xc4, 0xbc, 0xe4, 0xd4, 0xf8, 0xb2, 0xcc, 0xd4, 0xf2, 0x78, 0xa8, 0xbc,
	0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x02, 0x44, 0xaf, 0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0xdc,
	0x18, 0xbd, 0x32, 0x03, 0x3d, 0x98, 0x31, 0x52, 0x76, 0xb8, 0x2c, 0x2a, 0x4a, 0x2d, 0xce, 0x2f,
	0x2d, 0xc2, 0x6d, 0x13, 0xc4, 0x06, 0x29, 0x19, 0x98, 0xfe, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc,
	0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0xac, 0x92, 0x2b, 0x97, 0x9c, 0x7b, 0x6a,
	0x89, 0x07, 0xc8, 0x80, 0x00, 0x84, 0xfe, 0xb0, 0xcc, 0xd4, 0xf2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x65, 0x2e, 0x5e, 0x98, 0x4d, 0xf1, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x3c, 0x30, 0x41, 0xbf, 0xc4, 0xdc, 0x54, 0xa3, 0x4f, 0x8c, 0x5c, 0xd2,
	0xd8, 0x0c, 0x09, 0x86, 0xf8, 0x42, 0xe8, 0x12, 0x23, 0x97, 0x38, 0x0e, 0x7b, 0x84, 0x1c, 0xf4,
	0x08, 0x85, 0x81, 0x1e, 0x7e, 0x27, 0x4a, 0x99, 0xe3, 0x34, 0x01, 0x1e, 0x46, 0x7a, 0xd8, 0xf4,
	0x2b, 0xd9, 0x34, 0x5d, 0x7e, 0x32, 0x99, 0xc9, 0x4c, 0xc8, 0x04, 0x14, 0x9e, 0xd5, 0x28, 0xde,
	0xb4, 0x4d, 0x2e, 0x2d, 0x2e, 0xc9, 0xcf, 0x4d, 0x2d, 0x2a, 0xd6, 0xd7, 0x82, 0x04, 0x30, 0x9a,
	0xe6, 0x5a, 0xa7, 0x07, 0x8c, 0x5c, 0x2a, 0xc9, 0xf9, 0xb9, 0x04, 0x9d, 0xef, 0xa4, 0x80, 0x27,
	0x68, 0x02, 0x40, 0xd1, 0x10, 0xc0, 0x18, 0xe5, 0x01, 0x35, 0x25, 0x3d, 0x3f, 0x27, 0x31, 0x2f,
	0x5d, 0x2f, 0xbf, 0x28, 0x5d, 0x3f, 0x3d, 0x35, 0x0f, 0x1c, 0x49, 0xb0, 0x68, 0x2f, 0xc8, 0x2c,
	0xc6, 0x9d, 0xdc, 0xac, 0x61, 0x8c, 0x45, 0x4c, 0xcc, 0xee, 0x8e, 0x8e, 0xab, 0x98, 0x14, 0xdc,
	0x21, 0x06, 0x3a, 0xa6, 0x14, 0xeb, 0x41, 0x98, 0x20, 0x56, 0x98, 0x81, 0x1e, 0xd4, 0xe2, 0xe2,
	0x53, 0x30, 0x25, 0x31, 0x8e, 0x29, 0xc5, 0x31, 0x70, 0x25, 0x31, 0x61, 0x06, 0x31, 0x30, 0x25,
	0x49, 0x6c, 0x60, 0x07, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x98, 0x29, 0x9f, 0xee,
	0x02, 0x00, 0x00,
}
