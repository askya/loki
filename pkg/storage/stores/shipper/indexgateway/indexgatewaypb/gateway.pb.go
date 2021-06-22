// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/storage/stores/shipper/indexgateway/indexgatewaypb/gateway.proto

package indexgatewaypb

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type QueryIndexResponse struct {
	QueryKey string `protobuf:"bytes,1,opt,name=QueryKey,proto3" json:"QueryKey,omitempty"`
	Rows     []*Row `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (m *QueryIndexResponse) Reset()      { *m = QueryIndexResponse{} }
func (*QueryIndexResponse) ProtoMessage() {}
func (*QueryIndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a7bd4603d312b2, []int{0}
}
func (m *QueryIndexResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIndexResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIndexResponse.Merge(m, src)
}
func (m *QueryIndexResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIndexResponse proto.InternalMessageInfo

func (m *QueryIndexResponse) GetQueryKey() string {
	if m != nil {
		return m.QueryKey
	}
	return ""
}

func (m *QueryIndexResponse) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type Row struct {
	RangeValue []byte `protobuf:"bytes,1,opt,name=rangeValue,proto3" json:"rangeValue,omitempty"`
	Value      []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Row) Reset()      { *m = Row{} }
func (*Row) ProtoMessage() {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a7bd4603d312b2, []int{1}
}
func (m *Row) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Row.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return m.Size()
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetRangeValue() []byte {
	if m != nil {
		return m.RangeValue
	}
	return nil
}

func (m *Row) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type QueryIndexRequest struct {
	Queries []*IndexQuery `protobuf:"bytes,1,rep,name=Queries,proto3" json:"Queries,omitempty"`
}

func (m *QueryIndexRequest) Reset()      { *m = QueryIndexRequest{} }
func (*QueryIndexRequest) ProtoMessage() {}
func (*QueryIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a7bd4603d312b2, []int{2}
}
func (m *QueryIndexRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIndexRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIndexRequest.Merge(m, src)
}
func (m *QueryIndexRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIndexRequest proto.InternalMessageInfo

func (m *QueryIndexRequest) GetQueries() []*IndexQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

type IndexQuery struct {
	TableName        string `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	HashValue        string `protobuf:"bytes,2,opt,name=hashValue,proto3" json:"hashValue,omitempty"`
	RangeValuePrefix []byte `protobuf:"bytes,3,opt,name=rangeValuePrefix,proto3" json:"rangeValuePrefix,omitempty"`
	RangeValueStart  []byte `protobuf:"bytes,4,opt,name=rangeValueStart,proto3" json:"rangeValueStart,omitempty"`
	ValueEqual       []byte `protobuf:"bytes,5,opt,name=valueEqual,proto3" json:"valueEqual,omitempty"`
}

func (m *IndexQuery) Reset()      { *m = IndexQuery{} }
func (*IndexQuery) ProtoMessage() {}
func (*IndexQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a7bd4603d312b2, []int{3}
}
func (m *IndexQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexQuery.Merge(m, src)
}
func (m *IndexQuery) XXX_Size() int {
	return m.Size()
}
func (m *IndexQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexQuery.DiscardUnknown(m)
}

var xxx_messageInfo_IndexQuery proto.InternalMessageInfo

func (m *IndexQuery) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *IndexQuery) GetHashValue() string {
	if m != nil {
		return m.HashValue
	}
	return ""
}

func (m *IndexQuery) GetRangeValuePrefix() []byte {
	if m != nil {
		return m.RangeValuePrefix
	}
	return nil
}

func (m *IndexQuery) GetRangeValueStart() []byte {
	if m != nil {
		return m.RangeValueStart
	}
	return nil
}

func (m *IndexQuery) GetValueEqual() []byte {
	if m != nil {
		return m.ValueEqual
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryIndexResponse)(nil), "indexgatewaypb.QueryIndexResponse")
	proto.RegisterType((*Row)(nil), "indexgatewaypb.Row")
	proto.RegisterType((*QueryIndexRequest)(nil), "indexgatewaypb.QueryIndexRequest")
	proto.RegisterType((*IndexQuery)(nil), "indexgatewaypb.IndexQuery")
}

func init() {
	proto.RegisterFile("pkg/storage/stores/shipper/indexgateway/indexgatewaypb/gateway.proto", fileDescriptor_33a7bd4603d312b2)
}

var fileDescriptor_33a7bd4603d312b2 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbb, 0x4e, 0xe3, 0x40,
	0x14, 0x86, 0x3d, 0xb9, 0xec, 0x6e, 0xce, 0x46, 0x7b, 0x99, 0xdd, 0xc2, 0x8a, 0x56, 0xa3, 0xac,
	0x9b, 0x8d, 0xb6, 0x88, 0x51, 0x48, 0x47, 0x87, 0x40, 0x28, 0x42, 0x42, 0x30, 0x88, 0x48, 0x94,
	0x13, 0x71, 0x70, 0x2c, 0x42, 0xec, 0xcc, 0xd8, 0x38, 0xe9, 0x78, 0x04, 0x1e, 0x83, 0xa7, 0xa0,
	0xa6, 0x4c, 0x99, 0x92, 0x38, 0x0d, 0x65, 0x1e, 0x01, 0x79, 0x1c, 0xe2, 0x5c, 0x24, 0x2a, 0xfb,
	0x7c, 0xe7, 0xf7, 0x9c, 0xff, 0xfc, 0x1e, 0x38, 0xf0, 0x6f, 0x1c, 0x5b, 0x05, 0x9e, 0x14, 0x0e,
	0xea, 0x27, 0x2a, 0x5b, 0x75, 0x5d, 0xdf, 0x47, 0x69, 0xbb, 0xfd, 0x2b, 0x1c, 0x3a, 0x22, 0xc0,
	0x48, 0x8c, 0xd6, 0x0a, 0xbf, 0x63, 0x2f, 0xde, 0xea, 0xbe, 0xf4, 0x02, 0x8f, 0x7e, 0x5b, 0xef,
	0x5a, 0x97, 0x40, 0xcf, 0x42, 0x94, 0xa3, 0x56, 0x82, 0x39, 0x2a, 0xdf, 0xeb, 0x2b, 0xa4, 0x15,
	0xf8, 0xa2, 0xe9, 0x31, 0x8e, 0x4c, 0x52, 0x25, 0xb5, 0x12, 0x5f, 0xd6, 0xf4, 0x1f, 0x14, 0xa4,
	0x17, 0x29, 0x33, 0x57, 0xcd, 0xd7, 0xbe, 0x36, 0x7e, 0xd5, 0xd7, 0x0f, 0xac, 0x73, 0x2f, 0xe2,
	0x5a, 0x60, 0xed, 0x41, 0x9e, 0x7b, 0x11, 0x65, 0x00, 0x52, 0xf4, 0x1d, 0x6c, 0x8b, 0x5e, 0x88,
	0xfa, 0xb4, 0x32, 0x5f, 0x21, 0xf4, 0x37, 0x14, 0xef, 0x74, 0x2b, 0xa7, 0x5b, 0x69, 0x61, 0xb5,
	0xe0, 0xe7, 0xaa, 0xaf, 0x41, 0x88, 0x2a, 0xa0, 0x4d, 0xf8, 0x9c, 0x40, 0x17, 0x95, 0x49, 0xf4,
	0xf4, 0xca, 0xe6, 0x74, 0x2d, 0xd7, 0x1f, 0xf2, 0x77, 0xa9, 0xf5, 0x44, 0x00, 0x32, 0x4e, 0xff,
	0x40, 0x29, 0x10, 0x9d, 0x1e, 0x9e, 0x88, 0x5b, 0x5c, 0x2c, 0x97, 0x81, 0xa4, 0xdb, 0x15, 0xaa,
	0xdb, 0x5e, 0x3a, 0x2a, 0xf1, 0x0c, 0xd0, 0xff, 0xf0, 0x23, 0x73, 0x7e, 0x2a, 0xf1, 0xda, 0x1d,
	0x9a, 0x79, 0x6d, 0x7b, 0x8b, 0xd3, 0x1a, 0x7c, 0xcf, 0xd8, 0x79, 0x20, 0x64, 0x60, 0x16, 0xb4,
	0x74, 0x13, 0x27, 0x09, 0xe9, 0xa5, 0x0f, 0x07, 0xa1, 0xe8, 0x99, 0xc5, 0x34, 0xa1, 0x8c, 0x34,
	0x10, 0xca, 0xda, 0xff, 0x51, 0xba, 0x26, 0xbd, 0x00, 0xc8, 0xb2, 0xa1, 0x7f, 0x37, 0x33, 0xd8,
	0xca, 0xad, 0x62, 0x7d, 0x24, 0x49, 0x7f, 0xf9, 0x0e, 0xd9, 0x6f, 0x8e, 0xa7, 0xcc, 0x98, 0x4c,
	0x99, 0x31, 0x9f, 0x32, 0x72, 0x1f, 0x33, 0xf2, 0x18, 0x33, 0xf2, 0x1c, 0x33, 0x32, 0x8e, 0x19,
	0x79, 0x89, 0x19, 0x79, 0x8d, 0x99, 0x31, 0x8f, 0x19, 0x79, 0x98, 0x31, 0x63, 0x3c, 0x63, 0xc6,
	0x64, 0xc6, 0x8c, 0xce, 0x27, 0x7d, 0xaf, 0x76, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x95, 0x3c,
	0x4e, 0x5e, 0x9f, 0x02, 0x00, 0x00,
}

func (this *QueryIndexResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryIndexResponse)
	if !ok {
		that2, ok := that.(QueryIndexResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.QueryKey != that1.QueryKey {
		return false
	}
	if len(this.Rows) != len(that1.Rows) {
		return false
	}
	for i := range this.Rows {
		if !this.Rows[i].Equal(that1.Rows[i]) {
			return false
		}
	}
	return true
}
func (this *Row) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Row)
	if !ok {
		that2, ok := that.(Row)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.RangeValue, that1.RangeValue) {
		return false
	}
	if !bytes.Equal(this.Value, that1.Value) {
		return false
	}
	return true
}
func (this *QueryIndexRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryIndexRequest)
	if !ok {
		that2, ok := that.(QueryIndexRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Queries) != len(that1.Queries) {
		return false
	}
	for i := range this.Queries {
		if !this.Queries[i].Equal(that1.Queries[i]) {
			return false
		}
	}
	return true
}
func (this *IndexQuery) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IndexQuery)
	if !ok {
		that2, ok := that.(IndexQuery)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TableName != that1.TableName {
		return false
	}
	if this.HashValue != that1.HashValue {
		return false
	}
	if !bytes.Equal(this.RangeValuePrefix, that1.RangeValuePrefix) {
		return false
	}
	if !bytes.Equal(this.RangeValueStart, that1.RangeValueStart) {
		return false
	}
	if !bytes.Equal(this.ValueEqual, that1.ValueEqual) {
		return false
	}
	return true
}
func (this *QueryIndexResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&indexgatewaypb.QueryIndexResponse{")
	s = append(s, "QueryKey: "+fmt.Sprintf("%#v", this.QueryKey)+",\n")
	if this.Rows != nil {
		s = append(s, "Rows: "+fmt.Sprintf("%#v", this.Rows)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Row) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&indexgatewaypb.Row{")
	s = append(s, "RangeValue: "+fmt.Sprintf("%#v", this.RangeValue)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryIndexRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&indexgatewaypb.QueryIndexRequest{")
	if this.Queries != nil {
		s = append(s, "Queries: "+fmt.Sprintf("%#v", this.Queries)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *IndexQuery) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&indexgatewaypb.IndexQuery{")
	s = append(s, "TableName: "+fmt.Sprintf("%#v", this.TableName)+",\n")
	s = append(s, "HashValue: "+fmt.Sprintf("%#v", this.HashValue)+",\n")
	s = append(s, "RangeValuePrefix: "+fmt.Sprintf("%#v", this.RangeValuePrefix)+",\n")
	s = append(s, "RangeValueStart: "+fmt.Sprintf("%#v", this.RangeValueStart)+",\n")
	s = append(s, "ValueEqual: "+fmt.Sprintf("%#v", this.ValueEqual)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGateway(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IndexGatewayClient is the client API for IndexGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexGatewayClient interface {
	/// QueryIndex reads the indexes required for given query & sends back the batch of rows
	/// in rpc streams
	QueryIndex(ctx context.Context, in *QueryIndexRequest, opts ...grpc.CallOption) (IndexGateway_QueryIndexClient, error)
}

type indexGatewayClient struct {
	cc *grpc.ClientConn
}

func NewIndexGatewayClient(cc *grpc.ClientConn) IndexGatewayClient {
	return &indexGatewayClient{cc}
}

func (c *indexGatewayClient) QueryIndex(ctx context.Context, in *QueryIndexRequest, opts ...grpc.CallOption) (IndexGateway_QueryIndexClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IndexGateway_serviceDesc.Streams[0], "/indexgatewaypb.IndexGateway/QueryIndex", opts...)
	if err != nil {
		return nil, err
	}
	x := &indexGatewayQueryIndexClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IndexGateway_QueryIndexClient interface {
	Recv() (*QueryIndexResponse, error)
	grpc.ClientStream
}

type indexGatewayQueryIndexClient struct {
	grpc.ClientStream
}

func (x *indexGatewayQueryIndexClient) Recv() (*QueryIndexResponse, error) {
	m := new(QueryIndexResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IndexGatewayServer is the server API for IndexGateway service.
type IndexGatewayServer interface {
	/// QueryIndex reads the indexes required for given query & sends back the batch of rows
	/// in rpc streams
	QueryIndex(*QueryIndexRequest, IndexGateway_QueryIndexServer) error
}

func RegisterIndexGatewayServer(s *grpc.Server, srv IndexGatewayServer) {
	s.RegisterService(&_IndexGateway_serviceDesc, srv)
}

func _IndexGateway_QueryIndex_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryIndexRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IndexGatewayServer).QueryIndex(m, &indexGatewayQueryIndexServer{stream})
}

type IndexGateway_QueryIndexServer interface {
	Send(*QueryIndexResponse) error
	grpc.ServerStream
}

type indexGatewayQueryIndexServer struct {
	grpc.ServerStream
}

func (x *indexGatewayQueryIndexServer) Send(m *QueryIndexResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _IndexGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "indexgatewaypb.IndexGateway",
	HandlerType: (*IndexGatewayServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryIndex",
			Handler:       _IndexGateway_QueryIndex_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/storage/stores/shipper/indexgateway/indexgatewaypb/gateway.proto",
}

func (m *QueryIndexResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIndexResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.QueryKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGateway(dAtA, i, uint64(len(m.QueryKey)))
		i += copy(dAtA[i:], m.QueryKey)
	}
	if len(m.Rows) > 0 {
		for _, msg := range m.Rows {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGateway(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Row) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Row) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RangeValue) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGateway(dAtA, i, uint64(len(m.RangeValue)))
		i += copy(dAtA[i:], m.RangeValue)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGateway(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *QueryIndexRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIndexRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, msg := range m.Queries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGateway(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *IndexQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TableName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGateway(dAtA, i, uint64(len(m.TableName)))
		i += copy(dAtA[i:], m.TableName)
	}
	if len(m.HashValue) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGateway(dAtA, i, uint64(len(m.HashValue)))
		i += copy(dAtA[i:], m.HashValue)
	}
	if len(m.RangeValuePrefix) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGateway(dAtA, i, uint64(len(m.RangeValuePrefix)))
		i += copy(dAtA[i:], m.RangeValuePrefix)
	}
	if len(m.RangeValueStart) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGateway(dAtA, i, uint64(len(m.RangeValueStart)))
		i += copy(dAtA[i:], m.RangeValueStart)
	}
	if len(m.ValueEqual) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGateway(dAtA, i, uint64(len(m.ValueEqual)))
		i += copy(dAtA[i:], m.ValueEqual)
	}
	return i, nil
}

func encodeVarintGateway(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *QueryIndexResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.QueryKey)
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	if len(m.Rows) > 0 {
		for _, e := range m.Rows {
			l = e.Size()
			n += 1 + l + sovGateway(uint64(l))
		}
	}
	return n
}

func (m *Row) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RangeValue)
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	return n
}

func (m *QueryIndexRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovGateway(uint64(l))
		}
	}
	return n
}

func (m *IndexQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TableName)
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	l = len(m.HashValue)
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	l = len(m.RangeValuePrefix)
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	l = len(m.RangeValueStart)
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	l = len(m.ValueEqual)
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	return n
}

func sovGateway(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGateway(x uint64) (n int) {
	return sovGateway(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QueryIndexResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryIndexResponse{`,
		`QueryKey:` + fmt.Sprintf("%v", this.QueryKey) + `,`,
		`Rows:` + strings.Replace(fmt.Sprintf("%v", this.Rows), "Row", "Row", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Row) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Row{`,
		`RangeValue:` + fmt.Sprintf("%v", this.RangeValue) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueryIndexRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryIndexRequest{`,
		`Queries:` + strings.Replace(fmt.Sprintf("%v", this.Queries), "IndexQuery", "IndexQuery", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IndexQuery) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IndexQuery{`,
		`TableName:` + fmt.Sprintf("%v", this.TableName) + `,`,
		`HashValue:` + fmt.Sprintf("%v", this.HashValue) + `,`,
		`RangeValuePrefix:` + fmt.Sprintf("%v", this.RangeValuePrefix) + `,`,
		`RangeValueStart:` + fmt.Sprintf("%v", this.RangeValueStart) + `,`,
		`ValueEqual:` + fmt.Sprintf("%v", this.ValueEqual) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGateway(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QueryIndexResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGateway
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
			return fmt.Errorf("proto: QueryIndexResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIndexResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
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
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rows = append(m.Rows, &Row{})
			if err := m.Rows[len(m.Rows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGateway
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGateway
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
func (m *Row) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGateway
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
			return fmt.Errorf("proto: Row: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Row: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeValue = append(m.RangeValue[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeValue == nil {
				m.RangeValue = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGateway
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGateway
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
func (m *QueryIndexRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGateway
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
			return fmt.Errorf("proto: QueryIndexRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIndexRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
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
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queries = append(m.Queries, &IndexQuery{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGateway
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGateway
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
func (m *IndexQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGateway
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
			return fmt.Errorf("proto: IndexQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TableName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeValuePrefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeValuePrefix = append(m.RangeValuePrefix[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeValuePrefix == nil {
				m.RangeValuePrefix = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeValueStart", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeValueStart = append(m.RangeValueStart[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeValueStart == nil {
				m.RangeValueStart = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueEqual", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueEqual = append(m.ValueEqual[:0], dAtA[iNdEx:postIndex]...)
			if m.ValueEqual == nil {
				m.ValueEqual = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGateway
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGateway
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
func skipGateway(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGateway
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
					return 0, ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGateway
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
				return 0, ErrInvalidLengthGateway
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGateway
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGateway
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGateway(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGateway
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGateway = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGateway   = fmt.Errorf("proto: integer overflow")
)
