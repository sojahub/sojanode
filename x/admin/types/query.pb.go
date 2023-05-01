// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sojanode/admin/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ListAccountsRequest struct {
}

func (m *ListAccountsRequest) Reset()         { *m = ListAccountsRequest{} }
func (m *ListAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAccountsRequest) ProtoMessage()    {}
func (*ListAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e062bad86f8e9de, []int{0}
}
func (m *ListAccountsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAccountsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccountsRequest.Merge(m, src)
}
func (m *ListAccountsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccountsRequest proto.InternalMessageInfo

type ListAccountsResponse struct {
	Keys []*AdminAccount `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (m *ListAccountsResponse) Reset()         { *m = ListAccountsResponse{} }
func (m *ListAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAccountsResponse) ProtoMessage()    {}
func (*ListAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e062bad86f8e9de, []int{1}
}
func (m *ListAccountsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAccountsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccountsResponse.Merge(m, src)
}
func (m *ListAccountsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccountsResponse proto.InternalMessageInfo

func (m *ListAccountsResponse) GetKeys() []*AdminAccount {
	if m != nil {
		return m.Keys
	}
	return nil
}

type GetParamsRequest struct {
}

func (m *GetParamsRequest) Reset()         { *m = GetParamsRequest{} }
func (m *GetParamsRequest) String() string { return proto.CompactTextString(m) }
func (*GetParamsRequest) ProtoMessage()    {}
func (*GetParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e062bad86f8e9de, []int{2}
}
func (m *GetParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParamsRequest.Merge(m, src)
}
func (m *GetParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetParamsRequest proto.InternalMessageInfo

type GetParamsResponse struct {
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *GetParamsResponse) Reset()         { *m = GetParamsResponse{} }
func (m *GetParamsResponse) String() string { return proto.CompactTextString(m) }
func (*GetParamsResponse) ProtoMessage()    {}
func (*GetParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e062bad86f8e9de, []int{3}
}
func (m *GetParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParamsResponse.Merge(m, src)
}
func (m *GetParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetParamsResponse proto.InternalMessageInfo

func (m *GetParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*ListAccountsRequest)(nil), "sojanode.admin.v1.ListAccountsRequest")
	proto.RegisterType((*ListAccountsResponse)(nil), "sojanode.admin.v1.ListAccountsResponse")
	proto.RegisterType((*GetParamsRequest)(nil), "sojanode.admin.v1.GetParamsRequest")
	proto.RegisterType((*GetParamsResponse)(nil), "sojanode.admin.v1.GetParamsResponse")
}

func init() { proto.RegisterFile("sojanode/admin/v1/query.proto", fileDescriptor_3e062bad86f8e9de) }

var fileDescriptor_3e062bad86f8e9de = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcb, 0x4e, 0x02, 0x31,
	0x14, 0x9d, 0xf1, 0x41, 0x62, 0x71, 0x81, 0x15, 0x13, 0x32, 0x21, 0x0d, 0x19, 0xa3, 0xc1, 0xcd,
	0x54, 0xc6, 0x2f, 0x40, 0x63, 0x4c, 0x8c, 0x0b, 0xc5, 0xc4, 0x85, 0x1b, 0x53, 0x86, 0x32, 0x34,
	0x4a, 0xef, 0x40, 0x3b, 0x44, 0xfe, 0xc2, 0x4f, 0x72, 0xe9, 0x92, 0xa5, 0x4b, 0x03, 0x3f, 0x62,
	0x28, 0x1d, 0x82, 0x8c, 0xd1, 0xdd, 0xcd, 0x3d, 0x8f, 0x9e, 0xdb, 0x83, 0xaa, 0x4a, 0x74, 0x25,
	0x74, 0x38, 0x65, 0x9d, 0xbe, 0x90, 0x74, 0xd4, 0xa0, 0x83, 0x94, 0x0f, 0xc7, 0x41, 0x32, 0x04,
	0x0d, 0xb8, 0x64, 0xd1, 0xc0, 0xa0, 0xc1, 0xa8, 0xe1, 0x95, 0x63, 0x88, 0xc1, 0x80, 0x74, 0x3e,
	0x2d, 0x78, 0x5e, 0x35, 0x06, 0x88, 0x5f, 0x38, 0x65, 0x89, 0xa0, 0x4c, 0x4a, 0xd0, 0x4c, 0x0b,
	0x90, 0x2a, 0x43, 0x73, 0x6f, 0xe8, 0x71, 0xc2, 0x2d, 0xea, 0x1f, 0xa0, 0xfd, 0x1b, 0xa1, 0x74,
	0x33, 0x8a, 0x20, 0x95, 0x5a, 0xb5, 0xf8, 0x20, 0xe5, 0x4a, 0xfb, 0xd7, 0xa8, 0xfc, 0x73, 0xad,
	0x12, 0x90, 0x8a, 0xe3, 0x10, 0x6d, 0x3d, 0xf3, 0xb1, 0xaa, 0x6c, 0xd4, 0x36, 0xeb, 0xc5, 0x90,
	0x04, 0xeb, 0x09, 0x83, 0xe6, 0x7c, 0xb0, 0xb2, 0x96, 0xe1, 0xfa, 0x18, 0x95, 0xae, 0xb8, 0xbe,
	0x65, 0x43, 0xd6, 0x5f, 0xfa, 0x5f, 0xa2, 0xbd, 0x95, 0x9d, 0x35, 0x3f, 0x45, 0x85, 0xc4, 0x6c,
	0x2a, 0x6e, 0xcd, 0xad, 0x17, 0xc3, 0x4a, 0xde, 0xde, 0x2a, 0x2c, 0x2f, 0x7c, 0x77, 0xd1, 0xf6,
	0xdd, 0xfc, 0xc7, 0xf0, 0x13, 0xda, 0x5d, 0x0d, 0x8c, 0x8f, 0xf2, 0xda, 0x5f, 0xee, 0xf4, 0x8e,
	0xff, 0xa3, 0x2d, 0xa2, 0xf9, 0x0e, 0x7e, 0x40, 0x3b, 0xcb, 0xc4, 0xd8, 0xcf, 0xcb, 0xd6, 0x4f,
	0xf4, 0x0e, 0xff, 0xe4, 0x64, 0xbe, 0xe7, 0x17, 0x1f, 0x53, 0xe2, 0x4e, 0xa6, 0xc4, 0xfd, 0x9a,
	0x12, 0xf7, 0x6d, 0x46, 0x9c, 0xc9, 0x8c, 0x38, 0x9f, 0x33, 0xe2, 0x3c, 0x9e, 0xc4, 0x42, 0xf7,
	0xd2, 0x76, 0x10, 0x41, 0x9f, 0xde, 0x8b, 0x6e, 0xd4, 0x63, 0x42, 0xd2, 0xac, 0xcc, 0x57, 0x5b,
	0xa7, 0xe9, 0xb2, 0x5d, 0x30, 0x65, 0x9e, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x53, 0x9d, 0x55,
	0x9e, 0x50, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	GetParams(ctx context.Context, in *GetParamsRequest, opts ...grpc.CallOption) (*GetParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, "/sojanode.admin.v1.Query/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetParams(ctx context.Context, in *GetParamsRequest, opts ...grpc.CallOption) (*GetParamsResponse, error) {
	out := new(GetParamsResponse)
	err := c.cc.Invoke(ctx, "/sojanode.admin.v1.Query/GetParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	GetParams(context.Context, *GetParamsRequest) (*GetParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ListAccounts(ctx context.Context, req *ListAccountsRequest) (*ListAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (*UnimplementedQueryServer) GetParams(ctx context.Context, req *GetParamsRequest) (*GetParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParams not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sojanode.admin.v1.Query/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sojanode.admin.v1.Query/GetParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetParams(ctx, req.(*GetParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sojanode.admin.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAccounts",
			Handler:    _Query_ListAccounts_Handler,
		},
		{
			MethodName: "GetParams",
			Handler:    _Query_GetParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sojanode/admin/v1/query.proto",
}

func (m *ListAccountsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAccountsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAccountsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListAccountsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAccountsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAccountsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Keys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListAccountsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListAccountsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Keys) > 0 {
		for _, e := range m.Keys {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *GetParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListAccountsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListAccountsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAccountsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListAccountsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListAccountsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAccountsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, &AdminAccount{})
			if err := m.Keys[len(m.Keys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
