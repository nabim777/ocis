// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/app/provider/v1beta1/provider_api.proto

package providerv1beta1

import (
	context "context"
	fmt "fmt"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// REQUIRED.
// View mode.
type OpenInAppRequest_ViewMode int32

const (
	OpenInAppRequest_VIEW_MODE_INVALID OpenInAppRequest_ViewMode = 0
	// The resource can be opened but not downloaded.
	OpenInAppRequest_VIEW_MODE_VIEW_ONLY OpenInAppRequest_ViewMode = 1
	// The resource can be downloaded.
	OpenInAppRequest_VIEW_MODE_READ_ONLY OpenInAppRequest_ViewMode = 2
	// The resource can be downloaded and updated. The underlying application
	// MUST be a fully capable editor to support this mode.
	OpenInAppRequest_VIEW_MODE_READ_WRITE OpenInAppRequest_ViewMode = 3
	// The resource can be downloaded and updated, but must be shown in
	// preview mode. If the underlying application does not support a preview mode,
	// or if in a view-only mode users are not allowed to switch to edit mode,
	// then this mode MUST fall back to READ_WRITE.
	OpenInAppRequest_VIEW_MODE_PREVIEW OpenInAppRequest_ViewMode = 4
)

var OpenInAppRequest_ViewMode_name = map[int32]string{
	0: "VIEW_MODE_INVALID",
	1: "VIEW_MODE_VIEW_ONLY",
	2: "VIEW_MODE_READ_ONLY",
	3: "VIEW_MODE_READ_WRITE",
	4: "VIEW_MODE_PREVIEW",
}

var OpenInAppRequest_ViewMode_value = map[string]int32{
	"VIEW_MODE_INVALID":    0,
	"VIEW_MODE_VIEW_ONLY":  1,
	"VIEW_MODE_READ_ONLY":  2,
	"VIEW_MODE_READ_WRITE": 3,
	"VIEW_MODE_PREVIEW":    4,
}

func (x OpenInAppRequest_ViewMode) String() string {
	return proto.EnumName(OpenInAppRequest_ViewMode_name, int32(x))
}

func (OpenInAppRequest_ViewMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{0, 0}
}

type OpenInAppRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The resourceInfo to be opened. The gateway grpc message has a ref instead.
	ResourceInfo *v1beta11.ResourceInfo    `protobuf:"bytes,2,opt,name=resource_info,json=resourceInfo,proto3" json:"resource_info,omitempty"`
	ViewMode     OpenInAppRequest_ViewMode `protobuf:"varint,3,opt,name=view_mode,json=viewMode,proto3,enum=cs3.app.provider.v1beta1.OpenInAppRequest_ViewMode" json:"view_mode,omitempty"`
	// REQUIRED.
	// The access token this application provider will use when contacting
	// the storage provider to read and write.
	// Service implementors MUST make sure that the access token only grants
	// access to the requested resource.
	// Service implementors should use a ResourceId rather than a filepath to grant access, as
	// ResourceIds MUST NOT change when a resource is renamed.
	// The access token MUST be short-lived.
	// TODO(labkode): investigate token derivation techniques.
	AccessToken          string   `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenInAppRequest) Reset()         { *m = OpenInAppRequest{} }
func (m *OpenInAppRequest) String() string { return proto.CompactTextString(m) }
func (*OpenInAppRequest) ProtoMessage()    {}
func (*OpenInAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{0}
}

func (m *OpenInAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenInAppRequest.Unmarshal(m, b)
}
func (m *OpenInAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenInAppRequest.Marshal(b, m, deterministic)
}
func (m *OpenInAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenInAppRequest.Merge(m, src)
}
func (m *OpenInAppRequest) XXX_Size() int {
	return xxx_messageInfo_OpenInAppRequest.Size(m)
}
func (m *OpenInAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenInAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenInAppRequest proto.InternalMessageInfo

func (m *OpenInAppRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenInAppRequest) GetResourceInfo() *v1beta11.ResourceInfo {
	if m != nil {
		return m.ResourceInfo
	}
	return nil
}

func (m *OpenInAppRequest) GetViewMode() OpenInAppRequest_ViewMode {
	if m != nil {
		return m.ViewMode
	}
	return OpenInAppRequest_VIEW_MODE_INVALID
}

func (m *OpenInAppRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type OpenInAppResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The url that user agents will render to clients.
	// Usually the rendering happens by using HTML iframes or in separate browser tabs.
	AppUrl               *OpenInAppURL `protobuf:"bytes,3,opt,name=app_url,json=appUrl,proto3" json:"app_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *OpenInAppResponse) Reset()         { *m = OpenInAppResponse{} }
func (m *OpenInAppResponse) String() string { return proto.CompactTextString(m) }
func (*OpenInAppResponse) ProtoMessage()    {}
func (*OpenInAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{1}
}

func (m *OpenInAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenInAppResponse.Unmarshal(m, b)
}
func (m *OpenInAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenInAppResponse.Marshal(b, m, deterministic)
}
func (m *OpenInAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenInAppResponse.Merge(m, src)
}
func (m *OpenInAppResponse) XXX_Size() int {
	return xxx_messageInfo_OpenInAppResponse.Size(m)
}
func (m *OpenInAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenInAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenInAppResponse proto.InternalMessageInfo

func (m *OpenInAppResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OpenInAppResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenInAppResponse) GetAppUrl() *OpenInAppURL {
	if m != nil {
		return m.AppUrl
	}
	return nil
}

func init() {
	proto.RegisterEnum("cs3.app.provider.v1beta1.OpenInAppRequest_ViewMode", OpenInAppRequest_ViewMode_name, OpenInAppRequest_ViewMode_value)
	proto.RegisterType((*OpenInAppRequest)(nil), "cs3.app.provider.v1beta1.OpenInAppRequest")
	proto.RegisterType((*OpenInAppResponse)(nil), "cs3.app.provider.v1beta1.OpenInAppResponse")
}

func init() {
	proto.RegisterFile("cs3/app/provider/v1beta1/provider_api.proto", fileDescriptor_c007b70b037097fe)
}

var fileDescriptor_c007b70b037097fe = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x49, 0x3b, 0x95, 0xd5, 0x1d, 0x90, 0x19, 0xd0, 0x42, 0x35, 0xa4, 0xd2, 0x0b, 0x54,
	0x6d, 0xc8, 0x55, 0xdb, 0x07, 0x40, 0xe9, 0xda, 0x8b, 0x48, 0xdd, 0x1a, 0x99, 0xad, 0x13, 0xa8,
	0x52, 0xe4, 0xa5, 0x1e, 0x8a, 0xd8, 0xe2, 0x33, 0x3b, 0xe9, 0xc4, 0x15, 0x37, 0x3c, 0x01, 0x8f,
	0xc0, 0x25, 0x4f, 0xc0, 0x33, 0xf0, 0x54, 0x28, 0x8e, 0x13, 0xba, 0xa1, 0x6a, 0xbd, 0x8b, 0xcf,
	0xff, 0xfd, 0xc7, 0xf6, 0x7f, 0x1c, 0x74, 0x18, 0xaa, 0x41, 0x97, 0x01, 0x74, 0x41, 0x8a, 0x65,
	0xb4, 0xe0, 0xb2, 0xbb, 0xec, 0x5d, 0xf0, 0x84, 0xf5, 0xca, 0x42, 0xc0, 0x20, 0x22, 0x20, 0x45,
	0x22, 0xb0, 0x13, 0xaa, 0x01, 0x61, 0x00, 0xa4, 0xd0, 0x88, 0x81, 0x9b, 0x9d, 0xb5, 0x6d, 0x24,
	0x57, 0x22, 0x95, 0x21, 0x57, 0x79, 0x8f, 0xe6, 0x7e, 0x46, 0x4a, 0x08, 0x4b, 0x40, 0x25, 0x2c,
	0x49, 0x0b, 0xf5, 0x5d, 0xa6, 0xaa, 0x44, 0x48, 0xf6, 0x99, 0x3f, 0xdc, 0xeb, 0x75, 0x46, 0x27,
	0x5f, 0x81, 0xab, 0x12, 0xd1, 0xab, 0x5c, 0x6e, 0xff, 0xa8, 0x22, 0x7b, 0x0a, 0x3c, 0xf6, 0x62,
	0x17, 0x80, 0xf2, 0x9b, 0x94, 0xab, 0x04, 0xf7, 0x50, 0x4d, 0x00, 0xbb, 0x49, 0xb9, 0x63, 0xb5,
	0xac, 0x4e, 0xa3, 0xff, 0x8a, 0x64, 0x97, 0xca, 0x6d, 0xa6, 0x09, 0x99, 0x6a, 0x80, 0x1a, 0x10,
	0x4f, 0xd1, 0x93, 0x62, 0xe7, 0x20, 0x8a, 0x2f, 0x85, 0x53, 0xd1, 0xce, 0x03, 0xed, 0x34, 0x87,
	0xfd, 0x2f, 0x12, 0x42, 0x8d, 0xc5, 0x8b, 0x2f, 0x05, 0xdd, 0x91, 0x2b, 0x2b, 0xec, 0xa3, 0xfa,
	0x32, 0xe2, 0xb7, 0xc1, 0xb5, 0x58, 0x70, 0xa7, 0xda, 0xb2, 0x3a, 0x4f, 0xfb, 0x03, 0xb2, 0x2e,
	0x5b, 0x72, 0xff, 0x0a, 0x64, 0x16, 0xf1, 0xdb, 0x63, 0xb1, 0xe0, 0x74, 0x7b, 0x69, 0xbe, 0xf0,
	0x1b, 0xb4, 0xc3, 0xc2, 0x90, 0x2b, 0x15, 0x24, 0xe2, 0x0b, 0x8f, 0x9d, 0xad, 0x96, 0xd5, 0xa9,
	0xd3, 0x46, 0x5e, 0x3b, 0xcd, 0x4a, 0xed, 0xef, 0x16, 0xda, 0x2e, 0x9c, 0xf8, 0x25, 0xda, 0x9d,
	0x79, 0xe3, 0xf3, 0xe0, 0x78, 0x3a, 0x1a, 0x07, 0xde, 0xc9, 0xcc, 0x9d, 0x78, 0x23, 0xfb, 0x11,
	0xde, 0x43, 0xcf, 0xff, 0x95, 0xf5, 0xd7, 0xf4, 0x64, 0xf2, 0xd1, 0xb6, 0xee, 0x0a, 0x74, 0xec,
	0x8e, 0x72, 0xa1, 0x82, 0x1d, 0xf4, 0xe2, 0x9e, 0x70, 0x4e, 0xbd, 0xd3, 0xb1, 0x5d, 0xbd, 0xbb,
	0x85, 0x4f, 0xc7, 0xd9, 0xc2, 0xde, 0x6a, 0xff, 0xb6, 0xd0, 0xee, 0xca, 0x8d, 0x14, 0x88, 0x58,
	0x71, 0xdc, 0x45, 0xb5, 0xfc, 0x1d, 0x98, 0xa9, 0xec, 0xe9, 0x38, 0x24, 0x84, 0x65, 0x0a, 0x1f,
	0xb4, 0x4c, 0x0d, 0xb6, 0x32, 0xc6, 0xca, 0xa6, 0x63, 0x7c, 0x8f, 0x1e, 0x33, 0x80, 0x20, 0x95,
	0x57, 0x3a, 0xf3, 0x46, 0xff, 0xed, 0x06, 0x99, 0x9f, 0xd1, 0x09, 0xad, 0x31, 0x80, 0x33, 0x79,
	0xd5, 0x57, 0xa8, 0xe1, 0x1b, 0xd0, 0xf5, 0x3d, 0xbc, 0x40, 0xf5, 0x12, 0xc3, 0x07, 0x9b, 0xcf,
	0xaf, 0x79, 0xb8, 0x11, 0x9b, 0x27, 0x33, 0xfc, 0x86, 0xf6, 0x43, 0x71, 0xbd, 0xd6, 0x31, 0xb4,
	0xcb, 0x23, 0x41, 0xe4, 0x67, 0xcf, 0xde, 0xb7, 0x3e, 0x3d, 0x2b, 0x28, 0x03, 0xfd, 0xac, 0x54,
	0x8f, 0x5c, 0xff, 0x57, 0xc5, 0x39, 0x52, 0x03, 0xe2, 0x02, 0x90, 0xc2, 0x43, 0x66, 0xbd, 0x61,
	0x06, 0xfc, 0xd1, 0xd2, 0xdc, 0x05, 0x98, 0x17, 0xd2, 0xdc, 0x48, 0x17, 0x35, 0xfd, 0x33, 0x0d,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x27, 0x47, 0xc0, 0x89, 0x2a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProviderAPIClient is the client API for ProviderAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProviderAPIClient interface {
	// Returns the App URL and all necessary info to open a resource in an online editor.
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	OpenInApp(ctx context.Context, in *OpenInAppRequest, opts ...grpc.CallOption) (*OpenInAppResponse, error)
}

type providerAPIClient struct {
	cc *grpc.ClientConn
}

func NewProviderAPIClient(cc *grpc.ClientConn) ProviderAPIClient {
	return &providerAPIClient{cc}
}

func (c *providerAPIClient) OpenInApp(ctx context.Context, in *OpenInAppRequest, opts ...grpc.CallOption) (*OpenInAppResponse, error) {
	out := new(OpenInAppResponse)
	err := c.cc.Invoke(ctx, "/cs3.app.provider.v1beta1.ProviderAPI/OpenInApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderAPIServer is the server API for ProviderAPI service.
type ProviderAPIServer interface {
	// Returns the App URL and all necessary info to open a resource in an online editor.
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	OpenInApp(context.Context, *OpenInAppRequest) (*OpenInAppResponse, error)
}

// UnimplementedProviderAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProviderAPIServer struct {
}

func (*UnimplementedProviderAPIServer) OpenInApp(ctx context.Context, req *OpenInAppRequest) (*OpenInAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenInApp not implemented")
}

func RegisterProviderAPIServer(s *grpc.Server, srv ProviderAPIServer) {
	s.RegisterService(&_ProviderAPI_serviceDesc, srv)
}

func _ProviderAPI_OpenInApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenInAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderAPIServer).OpenInApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.app.provider.v1beta1.ProviderAPI/OpenInApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderAPIServer).OpenInApp(ctx, req.(*OpenInAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.app.provider.v1beta1.ProviderAPI",
	HandlerType: (*ProviderAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenInApp",
			Handler:    _ProviderAPI_OpenInApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/app/provider/v1beta1/provider_api.proto",
}
