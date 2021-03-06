// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vnode_scs.proto

/*
Package moac is a generated protocol buffer package.

It is generated from these files:
	vnode_scs.proto

It has these top-level messages:
	ScbPublicCallRequest
	ScbPublicCallReply
	RemoteCallRequest
	RemoteCallReply
	StorageRequest
	ChainInfoRequest
	ChainInfoReply
	AccountInfoRequest
	AccountInfoReply
	ScsPushMsg
	BlockSyncRequest
	BlockSyncReply
	UploadBlockRequest
	UploadBlockReply
	DownloadBlockRequest
	DownloadBlockReply
*/
package moac

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ScbPublicCallRequest struct {
	Requestid    uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Sender       []byte `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Contractaddr []byte `protobuf:"bytes,3,opt,name=contractaddr,proto3" json:"contractaddr,omitempty"`
	Data         []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ScbPublicCallRequest) Reset()                    { *m = ScbPublicCallRequest{} }
func (m *ScbPublicCallRequest) String() string            { return proto.CompactTextString(m) }
func (*ScbPublicCallRequest) ProtoMessage()               {}
func (*ScbPublicCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ScbPublicCallRequest) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *ScbPublicCallRequest) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *ScbPublicCallRequest) GetContractaddr() []byte {
	if m != nil {
		return m.Contractaddr
	}
	return nil
}

func (m *ScbPublicCallRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ScbPublicCallReply struct {
	Requestid uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Replybody []byte `protobuf:"bytes,2,opt,name=replybody,proto3" json:"replybody,omitempty"`
}

func (m *ScbPublicCallReply) Reset()                    { *m = ScbPublicCallReply{} }
func (m *ScbPublicCallReply) String() string            { return proto.CompactTextString(m) }
func (*ScbPublicCallReply) ProtoMessage()               {}
func (*ScbPublicCallReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ScbPublicCallReply) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *ScbPublicCallReply) GetReplybody() []byte {
	if m != nil {
		return m.Replybody
	}
	return nil
}

type RemoteCallRequest struct {
	Requestid    uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Sender       []byte `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Contractaddr []byte `protobuf:"bytes,3,opt,name=contractaddr,proto3" json:"contractaddr,omitempty"`
	Data         []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RemoteCallRequest) Reset()                    { *m = RemoteCallRequest{} }
func (m *RemoteCallRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoteCallRequest) ProtoMessage()               {}
func (*RemoteCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RemoteCallRequest) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *RemoteCallRequest) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *RemoteCallRequest) GetContractaddr() []byte {
	if m != nil {
		return m.Contractaddr
	}
	return nil
}

func (m *RemoteCallRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RemoteCallReply struct {
	Requestid uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Replybody []byte `protobuf:"bytes,2,opt,name=replybody,proto3" json:"replybody,omitempty"`
}

func (m *RemoteCallReply) Reset()                    { *m = RemoteCallReply{} }
func (m *RemoteCallReply) String() string            { return proto.CompactTextString(m) }
func (*RemoteCallReply) ProtoMessage()               {}
func (*RemoteCallReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RemoteCallReply) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *RemoteCallReply) GetReplybody() []byte {
	if m != nil {
		return m.Replybody
	}
	return nil
}

type StorageRequest struct {
	Reqtype      uint32 `protobuf:"varint,1,opt,name=reqtype" json:"reqtype,omitempty"`
	Storagekey   []byte `protobuf:"bytes,2,opt,name=storagekey,proto3" json:"storagekey,omitempty"`
	Position     []byte `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Structformat []byte `protobuf:"bytes,4,opt,name=structformat,proto3" json:"structformat,omitempty"`
}

func (m *StorageRequest) Reset()                    { *m = StorageRequest{} }
func (m *StorageRequest) String() string            { return proto.CompactTextString(m) }
func (*StorageRequest) ProtoMessage()               {}
func (*StorageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StorageRequest) GetReqtype() uint32 {
	if m != nil {
		return m.Reqtype
	}
	return 0
}

func (m *StorageRequest) GetStoragekey() []byte {
	if m != nil {
		return m.Storagekey
	}
	return nil
}

func (m *StorageRequest) GetPosition() []byte {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *StorageRequest) GetStructformat() []byte {
	if m != nil {
		return m.Structformat
	}
	return nil
}

type ChainInfoRequest struct {
	Requestid     uint32            `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Consensusaddr []byte            `protobuf:"bytes,2,opt,name=consensusaddr,proto3" json:"consensusaddr,omitempty"`
	Request       []*StorageRequest `protobuf:"bytes,3,rep,name=request" json:"request,omitempty"`
}

func (m *ChainInfoRequest) Reset()                    { *m = ChainInfoRequest{} }
func (m *ChainInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*ChainInfoRequest) ProtoMessage()               {}
func (*ChainInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ChainInfoRequest) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *ChainInfoRequest) GetConsensusaddr() []byte {
	if m != nil {
		return m.Consensusaddr
	}
	return nil
}

func (m *ChainInfoRequest) GetRequest() []*StorageRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type ChainInfoReply struct {
	Requestid uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Replybody []byte `protobuf:"bytes,2,opt,name=replybody,proto3" json:"replybody,omitempty"`
}

func (m *ChainInfoReply) Reset()                    { *m = ChainInfoReply{} }
func (m *ChainInfoReply) String() string            { return proto.CompactTextString(m) }
func (*ChainInfoReply) ProtoMessage()               {}
func (*ChainInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ChainInfoReply) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *ChainInfoReply) GetReplybody() []byte {
	if m != nil {
		return m.Replybody
	}
	return nil
}

type AccountInfoRequest struct {
	Requestid uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Addr      []byte `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *AccountInfoRequest) Reset()                    { *m = AccountInfoRequest{} }
func (m *AccountInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountInfoRequest) ProtoMessage()               {}
func (*AccountInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AccountInfoRequest) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *AccountInfoRequest) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

type AccountInfoReply struct {
	Requestid uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Replybody []byte `protobuf:"bytes,2,opt,name=replybody,proto3" json:"replybody,omitempty"`
}

func (m *AccountInfoReply) Reset()                    { *m = AccountInfoReply{} }
func (m *AccountInfoReply) String() string            { return proto.CompactTextString(m) }
func (*AccountInfoReply) ProtoMessage()               {}
func (*AccountInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AccountInfoReply) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *AccountInfoReply) GetReplybody() []byte {
	if m != nil {
		return m.Replybody
	}
	return nil
}

type ScsPushMsg struct {
	Requestid   []byte `protobuf:"bytes,1,opt,name=requestid,proto3" json:"requestid,omitempty"`
	Timestamp   []byte `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Requestflag bool   `protobuf:"varint,3,opt,name=requestflag" json:"requestflag,omitempty"`
	Type        []byte `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Status      []byte `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Scsid       []byte `protobuf:"bytes,6,opt,name=scsid,proto3" json:"scsid,omitempty"`
	Subchainid  []byte `protobuf:"bytes,7,opt,name=subchainid,proto3" json:"subchainid,omitempty"`
	Sender      []byte `protobuf:"bytes,8,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver    []byte `protobuf:"bytes,9,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Msghash     []byte `protobuf:"bytes,10,opt,name=msghash,proto3" json:"msghash,omitempty"`
}

func (m *ScsPushMsg) Reset()                    { *m = ScsPushMsg{} }
func (m *ScsPushMsg) String() string            { return proto.CompactTextString(m) }
func (*ScsPushMsg) ProtoMessage()               {}
func (*ScsPushMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ScsPushMsg) GetRequestid() []byte {
	if m != nil {
		return m.Requestid
	}
	return nil
}

func (m *ScsPushMsg) GetTimestamp() []byte {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ScsPushMsg) GetRequestflag() bool {
	if m != nil {
		return m.Requestflag
	}
	return false
}

func (m *ScsPushMsg) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ScsPushMsg) GetStatus() []byte {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ScsPushMsg) GetScsid() []byte {
	if m != nil {
		return m.Scsid
	}
	return nil
}

func (m *ScsPushMsg) GetSubchainid() []byte {
	if m != nil {
		return m.Subchainid
	}
	return nil
}

func (m *ScsPushMsg) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *ScsPushMsg) GetReceiver() []byte {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *ScsPushMsg) GetMsghash() []byte {
	if m != nil {
		return m.Msghash
	}
	return nil
}

type BlockSyncRequest struct {
	Requestid   uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Sender      []byte `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Blocknumber uint32 `protobuf:"varint,3,opt,name=blocknumber" json:"blocknumber,omitempty"`
	Scsid       []byte `protobuf:"bytes,4,opt,name=scsid,proto3" json:"scsid,omitempty"`
}

func (m *BlockSyncRequest) Reset()                    { *m = BlockSyncRequest{} }
func (m *BlockSyncRequest) String() string            { return proto.CompactTextString(m) }
func (*BlockSyncRequest) ProtoMessage()               {}
func (*BlockSyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BlockSyncRequest) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *BlockSyncRequest) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *BlockSyncRequest) GetBlocknumber() uint32 {
	if m != nil {
		return m.Blocknumber
	}
	return 0
}

func (m *BlockSyncRequest) GetScsid() []byte {
	if m != nil {
		return m.Scsid
	}
	return nil
}

type BlockSyncReply struct {
	Requestid uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Replybody []byte `protobuf:"bytes,2,opt,name=replybody,proto3" json:"replybody,omitempty"`
}

func (m *BlockSyncReply) Reset()                    { *m = BlockSyncReply{} }
func (m *BlockSyncReply) String() string            { return proto.CompactTextString(m) }
func (*BlockSyncReply) ProtoMessage()               {}
func (*BlockSyncReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *BlockSyncReply) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *BlockSyncReply) GetReplybody() []byte {
	if m != nil {
		return m.Replybody
	}
	return nil
}

type UploadBlockRequest struct {
	Requestid   uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Sender      []byte `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Blocknumber uint64 `protobuf:"varint,3,opt,name=blocknumber" json:"blocknumber,omitempty"`
	Blockhash   []byte `protobuf:"bytes,4,opt,name=blockhash,proto3" json:"blockhash,omitempty"`
	Scsid       []byte `protobuf:"bytes,5,opt,name=scsid,proto3" json:"scsid,omitempty"`
	Subchainid  []byte `protobuf:"bytes,6,opt,name=subchainid,proto3" json:"subchainid,omitempty"`
	Blockdata   []byte `protobuf:"bytes,7,opt,name=blockdata,proto3" json:"blockdata,omitempty"`
}

func (m *UploadBlockRequest) Reset()                    { *m = UploadBlockRequest{} }
func (m *UploadBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*UploadBlockRequest) ProtoMessage()               {}
func (*UploadBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UploadBlockRequest) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *UploadBlockRequest) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *UploadBlockRequest) GetBlocknumber() uint64 {
	if m != nil {
		return m.Blocknumber
	}
	return 0
}

func (m *UploadBlockRequest) GetBlockhash() []byte {
	if m != nil {
		return m.Blockhash
	}
	return nil
}

func (m *UploadBlockRequest) GetScsid() []byte {
	if m != nil {
		return m.Scsid
	}
	return nil
}

func (m *UploadBlockRequest) GetSubchainid() []byte {
	if m != nil {
		return m.Subchainid
	}
	return nil
}

func (m *UploadBlockRequest) GetBlockdata() []byte {
	if m != nil {
		return m.Blockdata
	}
	return nil
}

type UploadBlockReply struct {
	Requestid uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Replybody []byte `protobuf:"bytes,2,opt,name=replybody,proto3" json:"replybody,omitempty"`
}

func (m *UploadBlockReply) Reset()                    { *m = UploadBlockReply{} }
func (m *UploadBlockReply) String() string            { return proto.CompactTextString(m) }
func (*UploadBlockReply) ProtoMessage()               {}
func (*UploadBlockReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UploadBlockReply) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *UploadBlockReply) GetReplybody() []byte {
	if m != nil {
		return m.Replybody
	}
	return nil
}

type DownloadBlockRequest struct {
	Requestid   uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Sender      []byte `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Blocknumber uint64 `protobuf:"varint,3,opt,name=blocknumber" json:"blocknumber,omitempty"`
	Blockhash   []byte `protobuf:"bytes,4,opt,name=blockhash,proto3" json:"blockhash,omitempty"`
	Scsid       []byte `protobuf:"bytes,5,opt,name=scsid,proto3" json:"scsid,omitempty"`
	Uploadscsid []byte `protobuf:"bytes,6,opt,name=uploadscsid,proto3" json:"uploadscsid,omitempty"`
	Subchainid  []byte `protobuf:"bytes,7,opt,name=subchainid,proto3" json:"subchainid,omitempty"`
}

func (m *DownloadBlockRequest) Reset()                    { *m = DownloadBlockRequest{} }
func (m *DownloadBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*DownloadBlockRequest) ProtoMessage()               {}
func (*DownloadBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *DownloadBlockRequest) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *DownloadBlockRequest) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *DownloadBlockRequest) GetBlocknumber() uint64 {
	if m != nil {
		return m.Blocknumber
	}
	return 0
}

func (m *DownloadBlockRequest) GetBlockhash() []byte {
	if m != nil {
		return m.Blockhash
	}
	return nil
}

func (m *DownloadBlockRequest) GetScsid() []byte {
	if m != nil {
		return m.Scsid
	}
	return nil
}

func (m *DownloadBlockRequest) GetUploadscsid() []byte {
	if m != nil {
		return m.Uploadscsid
	}
	return nil
}

func (m *DownloadBlockRequest) GetSubchainid() []byte {
	if m != nil {
		return m.Subchainid
	}
	return nil
}

type DownloadBlockReply struct {
	Requestid uint32 `protobuf:"varint,1,opt,name=requestid" json:"requestid,omitempty"`
	Replybody []byte `protobuf:"bytes,2,opt,name=replybody,proto3" json:"replybody,omitempty"`
}

func (m *DownloadBlockReply) Reset()                    { *m = DownloadBlockReply{} }
func (m *DownloadBlockReply) String() string            { return proto.CompactTextString(m) }
func (*DownloadBlockReply) ProtoMessage()               {}
func (*DownloadBlockReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *DownloadBlockReply) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *DownloadBlockReply) GetReplybody() []byte {
	if m != nil {
		return m.Replybody
	}
	return nil
}

func init() {
	proto.RegisterType((*ScbPublicCallRequest)(nil), "moac.ScbPublicCallRequest")
	proto.RegisterType((*ScbPublicCallReply)(nil), "moac.ScbPublicCallReply")
	proto.RegisterType((*RemoteCallRequest)(nil), "moac.RemoteCallRequest")
	proto.RegisterType((*RemoteCallReply)(nil), "moac.RemoteCallReply")
	proto.RegisterType((*StorageRequest)(nil), "moac.StorageRequest")
	proto.RegisterType((*ChainInfoRequest)(nil), "moac.ChainInfoRequest")
	proto.RegisterType((*ChainInfoReply)(nil), "moac.ChainInfoReply")
	proto.RegisterType((*AccountInfoRequest)(nil), "moac.AccountInfoRequest")
	proto.RegisterType((*AccountInfoReply)(nil), "moac.AccountInfoReply")
	proto.RegisterType((*ScsPushMsg)(nil), "moac.ScsPushMsg")
	proto.RegisterType((*BlockSyncRequest)(nil), "moac.BlockSyncRequest")
	proto.RegisterType((*BlockSyncReply)(nil), "moac.BlockSyncReply")
	proto.RegisterType((*UploadBlockRequest)(nil), "moac.UploadBlockRequest")
	proto.RegisterType((*UploadBlockReply)(nil), "moac.UploadBlockReply")
	proto.RegisterType((*DownloadBlockRequest)(nil), "moac.DownloadBlockRequest")
	proto.RegisterType((*DownloadBlockReply)(nil), "moac.DownloadBlockReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SCS service

type SCSClient interface {
	Sync(ctx context.Context, in *BlockSyncRequest, opts ...grpc.CallOption) (*BlockSyncReply, error)
}

type sCSClient struct {
	cc *grpc.ClientConn
}

func NewSCSClient(cc *grpc.ClientConn) SCSClient {
	return &sCSClient{cc}
}

func (c *sCSClient) Sync(ctx context.Context, in *BlockSyncRequest, opts ...grpc.CallOption) (*BlockSyncReply, error) {
	out := new(BlockSyncReply)
	err := grpc.Invoke(ctx, "/moac.SCS/Sync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SCS service

type SCSServer interface {
	Sync(context.Context, *BlockSyncRequest) (*BlockSyncReply, error)
}

func RegisterSCSServer(s *grpc.Server, srv SCSServer) {
	s.RegisterService(&_SCS_serviceDesc, srv)
}

func _SCS_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCSServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moac.SCS/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCSServer).Sync(ctx, req.(*BlockSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SCS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moac.SCS",
	HandlerType: (*SCSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _SCS_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vnode_scs.proto",
}

// Client API for Vnode service

type VnodeClient interface {
	ScsPush(ctx context.Context, opts ...grpc.CallOption) (Vnode_ScsPushClient, error)
	AccountInfo(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*AccountInfoReply, error)
	ChainInfo(ctx context.Context, in *ChainInfoRequest, opts ...grpc.CallOption) (*ChainInfoReply, error)
	RemoteCall(ctx context.Context, in *RemoteCallRequest, opts ...grpc.CallOption) (*RemoteCallReply, error)
	ScbPublicCall(ctx context.Context, in *ScbPublicCallRequest, opts ...grpc.CallOption) (*ScbPublicCallReply, error)
	UploadBlock(ctx context.Context, in *UploadBlockRequest, opts ...grpc.CallOption) (*UploadBlockReply, error)
	DownloadBlock(ctx context.Context, in *DownloadBlockRequest, opts ...grpc.CallOption) (*DownloadBlockReply, error)
}

type vnodeClient struct {
	cc *grpc.ClientConn
}

func NewVnodeClient(cc *grpc.ClientConn) VnodeClient {
	return &vnodeClient{cc}
}

func (c *vnodeClient) ScsPush(ctx context.Context, opts ...grpc.CallOption) (Vnode_ScsPushClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Vnode_serviceDesc.Streams[0], c.cc, "/moac.Vnode/ScsPush", opts...)
	if err != nil {
		return nil, err
	}
	x := &vnodeScsPushClient{stream}
	return x, nil
}

type Vnode_ScsPushClient interface {
	Send(*ScsPushMsg) error
	Recv() (*ScsPushMsg, error)
	grpc.ClientStream
}

type vnodeScsPushClient struct {
	grpc.ClientStream
}

func (x *vnodeScsPushClient) Send(m *ScsPushMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vnodeScsPushClient) Recv() (*ScsPushMsg, error) {
	m := new(ScsPushMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vnodeClient) AccountInfo(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*AccountInfoReply, error) {
	out := new(AccountInfoReply)
	err := grpc.Invoke(ctx, "/moac.Vnode/AccountInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnodeClient) ChainInfo(ctx context.Context, in *ChainInfoRequest, opts ...grpc.CallOption) (*ChainInfoReply, error) {
	out := new(ChainInfoReply)
	err := grpc.Invoke(ctx, "/moac.Vnode/ChainInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnodeClient) RemoteCall(ctx context.Context, in *RemoteCallRequest, opts ...grpc.CallOption) (*RemoteCallReply, error) {
	out := new(RemoteCallReply)
	err := grpc.Invoke(ctx, "/moac.Vnode/RemoteCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnodeClient) ScbPublicCall(ctx context.Context, in *ScbPublicCallRequest, opts ...grpc.CallOption) (*ScbPublicCallReply, error) {
	out := new(ScbPublicCallReply)
	err := grpc.Invoke(ctx, "/moac.Vnode/ScbPublicCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnodeClient) UploadBlock(ctx context.Context, in *UploadBlockRequest, opts ...grpc.CallOption) (*UploadBlockReply, error) {
	out := new(UploadBlockReply)
	err := grpc.Invoke(ctx, "/moac.Vnode/UploadBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnodeClient) DownloadBlock(ctx context.Context, in *DownloadBlockRequest, opts ...grpc.CallOption) (*DownloadBlockReply, error) {
	out := new(DownloadBlockReply)
	err := grpc.Invoke(ctx, "/moac.Vnode/DownloadBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Vnode service

type VnodeServer interface {
	ScsPush(Vnode_ScsPushServer) error
	AccountInfo(context.Context, *AccountInfoRequest) (*AccountInfoReply, error)
	ChainInfo(context.Context, *ChainInfoRequest) (*ChainInfoReply, error)
	RemoteCall(context.Context, *RemoteCallRequest) (*RemoteCallReply, error)
	ScbPublicCall(context.Context, *ScbPublicCallRequest) (*ScbPublicCallReply, error)
	UploadBlock(context.Context, *UploadBlockRequest) (*UploadBlockReply, error)
	DownloadBlock(context.Context, *DownloadBlockRequest) (*DownloadBlockReply, error)
}

func RegisterVnodeServer(s *grpc.Server, srv VnodeServer) {
	s.RegisterService(&_Vnode_serviceDesc, srv)
}

func _Vnode_ScsPush_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VnodeServer).ScsPush(&vnodeScsPushServer{stream})
}

type Vnode_ScsPushServer interface {
	Send(*ScsPushMsg) error
	Recv() (*ScsPushMsg, error)
	grpc.ServerStream
}

type vnodeScsPushServer struct {
	grpc.ServerStream
}

func (x *vnodeScsPushServer) Send(m *ScsPushMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vnodeScsPushServer) Recv() (*ScsPushMsg, error) {
	m := new(ScsPushMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Vnode_AccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnodeServer).AccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moac.Vnode/AccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnodeServer).AccountInfo(ctx, req.(*AccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnode_ChainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnodeServer).ChainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moac.Vnode/ChainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnodeServer).ChainInfo(ctx, req.(*ChainInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnode_RemoteCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnodeServer).RemoteCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moac.Vnode/RemoteCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnodeServer).RemoteCall(ctx, req.(*RemoteCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnode_ScbPublicCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScbPublicCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnodeServer).ScbPublicCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moac.Vnode/ScbPublicCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnodeServer).ScbPublicCall(ctx, req.(*ScbPublicCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnode_UploadBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnodeServer).UploadBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moac.Vnode/UploadBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnodeServer).UploadBlock(ctx, req.(*UploadBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnode_DownloadBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnodeServer).DownloadBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moac.Vnode/DownloadBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnodeServer).DownloadBlock(ctx, req.(*DownloadBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vnode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moac.Vnode",
	HandlerType: (*VnodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountInfo",
			Handler:    _Vnode_AccountInfo_Handler,
		},
		{
			MethodName: "ChainInfo",
			Handler:    _Vnode_ChainInfo_Handler,
		},
		{
			MethodName: "RemoteCall",
			Handler:    _Vnode_RemoteCall_Handler,
		},
		{
			MethodName: "ScbPublicCall",
			Handler:    _Vnode_ScbPublicCall_Handler,
		},
		{
			MethodName: "UploadBlock",
			Handler:    _Vnode_UploadBlock_Handler,
		},
		{
			MethodName: "DownloadBlock",
			Handler:    _Vnode_DownloadBlock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ScsPush",
			Handler:       _Vnode_ScsPush_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "vnode_scs.proto",
}

func init() { proto.RegisterFile("vnode_scs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcf, 0x6a, 0xdb, 0x4e,
	0x10, 0xb6, 0xe2, 0xff, 0xe3, 0x38, 0xf1, 0x6f, 0xf1, 0x2f, 0x15, 0x22, 0x14, 0x23, 0x7a, 0xf0,
	0xc9, 0x94, 0x04, 0x7a, 0x29, 0x2d, 0x24, 0x29, 0x2d, 0x85, 0x26, 0x18, 0x9b, 0xf6, 0x5a, 0x56,
	0xab, 0x8d, 0x2d, 0x22, 0x69, 0x15, 0xed, 0x2a, 0xc5, 0xb7, 0x16, 0x4a, 0x0f, 0x7d, 0x84, 0xbe,
	0x55, 0x5f, 0xa0, 0xa7, 0x3e, 0x48, 0xd1, 0x4a, 0x6b, 0xaf, 0xac, 0x94, 0x04, 0x54, 0x68, 0x6f,
	0x9a, 0x6f, 0x56, 0xc3, 0xf7, 0xcd, 0xce, 0xcc, 0x0e, 0xec, 0xdf, 0x84, 0xcc, 0xa5, 0xef, 0x39,
	0xe1, 0x93, 0x28, 0x66, 0x82, 0xa1, 0x46, 0xc0, 0x30, 0xb1, 0x3f, 0x1b, 0x30, 0x9c, 0x13, 0x67,
	0x9a, 0x38, 0xbe, 0x47, 0xce, 0xb0, 0xef, 0xcf, 0xe8, 0x75, 0x42, 0xb9, 0x40, 0x87, 0xd0, 0x8d,
	0xb3, 0x4f, 0xcf, 0x35, 0x8d, 0x91, 0x31, 0xee, 0xcf, 0x36, 0x00, 0x3a, 0x80, 0x16, 0xa7, 0xa1,
	0x4b, 0x63, 0x73, 0x67, 0x64, 0x8c, 0x77, 0x67, 0xb9, 0x85, 0x6c, 0xd8, 0x25, 0x2c, 0x14, 0x31,
	0x26, 0x02, 0xbb, 0x6e, 0x6c, 0xd6, 0xa5, 0xb7, 0x80, 0x21, 0x04, 0x0d, 0x17, 0x0b, 0x6c, 0x36,
	0xa4, 0x4f, 0x7e, 0xdb, 0x53, 0x40, 0x5b, 0x2c, 0x22, 0x7f, 0x75, 0x07, 0x07, 0xe9, 0x8d, 0xfc,
	0x95, 0xc3, 0xdc, 0x55, 0x4e, 0x63, 0x03, 0xd8, 0x9f, 0x0c, 0xf8, 0x6f, 0x46, 0x03, 0x26, 0xe8,
	0xdf, 0x53, 0x75, 0x0e, 0xfb, 0x3a, 0x85, 0xaa, 0x92, 0xbe, 0x1a, 0xb0, 0x37, 0x17, 0x2c, 0xc6,
	0x0b, 0xaa, 0xf4, 0x98, 0xd0, 0x8e, 0xe9, 0xb5, 0x58, 0x45, 0x34, 0x0f, 0xa6, 0x4c, 0xf4, 0x10,
	0x80, 0x67, 0x67, 0xaf, 0xa8, 0x8a, 0xa5, 0x21, 0xc8, 0x82, 0x4e, 0xc4, 0xb8, 0x27, 0x3c, 0x16,
	0xe6, 0x7a, 0xd6, 0x76, 0xaa, 0x97, 0x8b, 0x38, 0x21, 0xe2, 0x92, 0xc5, 0x01, 0x16, 0xb9, 0xa6,
	0x02, 0x66, 0x7f, 0x31, 0x60, 0x70, 0xb6, 0xc4, 0x5e, 0xf8, 0x3a, 0xbc, 0x64, 0xf7, 0x4b, 0xef,
	0x23, 0xe8, 0x13, 0x16, 0x72, 0x1a, 0xf2, 0x84, 0xcb, 0x3c, 0x66, 0xac, 0x8a, 0x20, 0x9a, 0x48,
	0x49, 0xe9, 0x2f, 0x66, 0x7d, 0x54, 0x1f, 0xf7, 0x8e, 0x86, 0x93, 0xb4, 0x52, 0x27, 0x45, 0xe5,
	0x33, 0x75, 0xc8, 0x7e, 0x03, 0x7b, 0x1a, 0x8f, 0xaa, 0x39, 0x7e, 0x09, 0xe8, 0x84, 0x10, 0x96,
	0x84, 0xe2, 0xfe, 0xba, 0x10, 0x34, 0x34, 0x39, 0xf2, 0xdb, 0xbe, 0x80, 0x41, 0x21, 0x4e, 0x55,
	0x5e, 0xdf, 0x76, 0x00, 0xe6, 0x84, 0x4f, 0x13, 0xbe, 0x3c, 0xe7, 0x8b, 0x72, 0xa8, 0xdd, 0xad,
	0x50, 0xc2, 0x0b, 0x28, 0x17, 0x38, 0x88, 0x54, 0xa8, 0x35, 0x80, 0x46, 0xd0, 0xcb, 0x8f, 0x5e,
	0xfa, 0x78, 0x21, 0x2f, 0xbf, 0x33, 0xd3, 0xa1, 0x54, 0x90, 0x2c, 0xa9, 0xbc, 0x96, 0x65, 0x3d,
	0xa5, 0xbd, 0x21, 0xb0, 0x48, 0xb8, 0xd9, 0xcc, 0x7b, 0x43, 0x5a, 0x68, 0x08, 0x4d, 0x4e, 0xb8,
	0xe7, 0x9a, 0x2d, 0x09, 0x67, 0x86, 0xac, 0xbe, 0xc4, 0x21, 0xe9, 0xbd, 0x78, 0xae, 0xd9, 0xce,
	0xab, 0x6f, 0x8d, 0x68, 0x9d, 0xd6, 0x29, 0x74, 0x9a, 0x05, 0x9d, 0x98, 0x12, 0xea, 0xdd, 0xd0,
	0xd8, 0xec, 0x66, 0x55, 0xa9, 0xec, 0xb4, 0xd6, 0x03, 0xbe, 0x58, 0x62, 0xbe, 0x34, 0x41, 0xba,
	0x94, 0x69, 0x7f, 0x34, 0x60, 0x70, 0xea, 0x33, 0x72, 0x35, 0x5f, 0x85, 0xa4, 0x5a, 0xab, 0x8f,
	0xa0, 0xe7, 0xa4, 0x91, 0xc2, 0x24, 0x70, 0x68, 0xd6, 0xe9, 0xfd, 0x99, 0x0e, 0x6d, 0x04, 0x37,
	0x34, 0xc1, 0x69, 0x15, 0x6a, 0x0c, 0xaa, 0xde, 0xf6, 0x0f, 0x03, 0xd0, 0xdb, 0xc8, 0x67, 0xd8,
	0x95, 0x41, 0xff, 0xb8, 0xa4, 0x46, 0x51, 0xd2, 0x21, 0x74, 0xa5, 0x29, 0x73, 0x9b, 0xc9, 0xda,
	0x00, 0x1b, 0xc1, 0xcd, 0xdf, 0xdf, 0x70, 0xab, 0x74, 0xc3, 0x2a, 0xa6, 0x1c, 0x8a, 0x6d, 0x2d,
	0xa6, 0x9c, 0x8c, 0x17, 0x30, 0x28, 0xe8, 0xab, 0x9a, 0xb0, 0x9f, 0x06, 0x0c, 0x5f, 0xb0, 0x0f,
	0xe1, 0x3f, 0x9d, 0xb2, 0x11, 0xf4, 0x12, 0x29, 0x5a, 0x6f, 0x18, 0x1d, 0xba, 0xab, 0x6d, 0xd2,
	0x67, 0x72, 0x4b, 0x65, 0xc5, 0xc4, 0x1d, 0x3d, 0x83, 0xfa, 0xfc, 0x6c, 0x8e, 0x9e, 0x40, 0x23,
	0xad, 0x5c, 0x74, 0x90, 0xcd, 0xda, 0xed, 0x66, 0xb2, 0x86, 0x25, 0x3c, 0xf2, 0x57, 0x76, 0xed,
	0xe8, 0x7b, 0x1d, 0x9a, 0xef, 0xd2, 0xc5, 0x02, 0x1d, 0x43, 0x3b, 0x9f, 0x4f, 0x68, 0x90, 0x0f,
	0xec, 0xf5, 0xb8, 0xb2, 0x4a, 0x88, 0x5d, 0x1b, 0x1b, 0x8f, 0x0d, 0x74, 0x02, 0x3d, 0x6d, 0x4a,
	0x22, 0x33, 0x3b, 0x56, 0x1e, 0xc0, 0xd6, 0xc1, 0x2d, 0x1e, 0xc9, 0x00, 0x3d, 0x85, 0xee, 0x7a,
	0xfc, 0x2b, 0xfa, 0xdb, 0xef, 0x92, 0xa2, 0x5f, 0x7c, 0x27, 0xec, 0x1a, 0x7a, 0x0e, 0xb0, 0x79,
	0xa0, 0xd1, 0x83, 0xec, 0x54, 0x69, 0x6b, 0xb0, 0xfe, 0x2f, 0x3b, 0xb2, 0xff, 0x5f, 0x41, 0xbf,
	0xb0, 0xb6, 0x20, 0x4b, 0x09, 0x2d, 0x6f, 0x54, 0x96, 0x79, 0xab, 0x2f, 0x0b, 0x74, 0x02, 0x3d,
	0xad, 0x1f, 0x54, 0x22, 0xca, 0x23, 0x40, 0x25, 0x62, 0xbb, 0x79, 0x32, 0x2e, 0x85, 0xda, 0x50,
	0x5c, 0x6e, 0x6b, 0x0b, 0xc5, 0xa5, 0x5c, 0x4c, 0x76, 0xed, 0xb4, 0x7f, 0xda, 0x3d, 0x67, 0x98,
	0x4c, 0xd3, 0x2d, 0x71, 0x6a, 0x38, 0x2d, 0xb9, 0x2e, 0x1e, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0x48, 0x1b, 0x92, 0x41, 0x0a, 0x00, 0x00,
}
