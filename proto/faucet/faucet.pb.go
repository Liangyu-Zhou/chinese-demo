// Code generated by protoc-gen-go. DO NOT EDIT.
// source: faucet/faucet.proto

package faucet_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FundingRequest struct {
	WalletAddress        string   `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	CaptchaResponse      string   `protobuf:"bytes,2,opt,name=captcha_response,json=captchaResponse,proto3" json:"captcha_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FundingRequest) Reset()         { *m = FundingRequest{} }
func (m *FundingRequest) String() string { return proto.CompactTextString(m) }
func (*FundingRequest) ProtoMessage()    {}
func (*FundingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb075c8fce600193, []int{0}
}

func (m *FundingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingRequest.Unmarshal(m, b)
}
func (m *FundingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingRequest.Marshal(b, m, deterministic)
}
func (m *FundingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingRequest.Merge(m, src)
}
func (m *FundingRequest) XXX_Size() int {
	return xxx_messageInfo_FundingRequest.Size(m)
}
func (m *FundingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FundingRequest proto.InternalMessageInfo

func (m *FundingRequest) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *FundingRequest) GetCaptchaResponse() string {
	if m != nil {
		return m.CaptchaResponse
	}
	return ""
}

type FundingResponse struct {
	Amount               string   `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	TransactionHash      string   `protobuf:"bytes,2,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FundingResponse) Reset()         { *m = FundingResponse{} }
func (m *FundingResponse) String() string { return proto.CompactTextString(m) }
func (*FundingResponse) ProtoMessage()    {}
func (*FundingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb075c8fce600193, []int{1}
}

func (m *FundingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingResponse.Unmarshal(m, b)
}
func (m *FundingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingResponse.Marshal(b, m, deterministic)
}
func (m *FundingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingResponse.Merge(m, src)
}
func (m *FundingResponse) XXX_Size() int {
	return xxx_messageInfo_FundingResponse.Size(m)
}
func (m *FundingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FundingResponse proto.InternalMessageInfo

func (m *FundingResponse) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *FundingResponse) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func init() {
	proto.RegisterType((*FundingRequest)(nil), "faucet.FundingRequest")
	proto.RegisterType((*FundingResponse)(nil), "faucet.FundingResponse")
}

func init() { proto.RegisterFile("faucet/faucet.proto", fileDescriptor_fb075c8fce600193) }

var fileDescriptor_fb075c8fce600193 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xc9, 0x15, 0x01, 0x17, 0xbd, 0x93, 0x15, 0xe3, 0x71, 0x5a, 0x68, 0x40, 0x50, 0x8b,
	0x0b, 0x6a, 0x67, 0xa7, 0x45, 0xb0, 0x0e, 0x56, 0x36, 0x61, 0x92, 0xac, 0x49, 0x20, 0xee, 0xc4,
	0x9d, 0x89, 0xf6, 0xbe, 0x82, 0x8f, 0xe6, 0x2b, 0xf8, 0x20, 0x92, 0xec, 0x28, 0x8a, 0x56, 0xc3,
	0x7c, 0xcc, 0xfc, 0xf3, 0xcf, 0xaf, 0x76, 0x1e, 0x60, 0x28, 0x0d, 0x27, 0xbe, 0xac, 0x7b, 0x87,
	0x8c, 0x3a, 0xf4, 0xdd, 0xea, 0xa0, 0x46, 0xac, 0x3b, 0x93, 0x40, 0xdf, 0x26, 0x60, 0x2d, 0x32,
	0x70, 0x8b, 0x96, 0xfc, 0x54, 0x5c, 0xa8, 0x79, 0x3a, 0xd8, 0xaa, 0xb5, 0x75, 0x66, 0x9e, 0x06,
	0x43, 0xac, 0x8f, 0xd5, 0xfc, 0x05, 0xba, 0xce, 0x70, 0x0e, 0x55, 0xe5, 0x0c, 0xd1, 0x32, 0x38,
	0x0c, 0x4e, 0x36, 0xb2, 0x2d, 0x4f, 0xaf, 0x3d, 0xd4, 0xa7, 0x6a, 0xbb, 0x84, 0x9e, 0xcb, 0x06,
	0x72, 0x67, 0xa8, 0x47, 0x4b, 0x66, 0x39, 0x9b, 0x06, 0x17, 0xc2, 0x33, 0xc1, 0xf1, 0x9d, 0x5a,
	0x7c, 0xdf, 0xf0, 0x48, 0x47, 0x2a, 0x84, 0x47, 0x1c, 0x2c, 0x8b, 0xb8, 0x74, 0xa3, 0x2a, 0x3b,
	0xb0, 0x04, 0xe5, 0x68, 0x32, 0x6f, 0x80, 0x9a, 0x2f, 0xd5, 0x1f, 0xfc, 0x16, 0xa8, 0xb9, 0xe8,
	0x54, 0x98, 0x4e, 0x1f, 0xea, 0x42, 0x6d, 0x8a, 0xf9, 0xf1, 0x0c, 0xe9, 0x68, 0x2d, 0x41, 0xfc,
	0xfe, 0x6c, 0xb5, 0xf7, 0x87, 0x8b, 0xc1, 0xa3, 0xd7, 0xf7, 0x8f, 0xb7, 0xd9, 0x7e, 0x1c, 0x4d,
	0x21, 0x3d, 0x9f, 0x4b, 0x90, 0x89, 0xf3, 0x8b, 0x57, 0xc1, 0xd9, 0xcd, 0xee, 0xfd, 0x7f, 0x21,
	0x17, 0xe1, 0x54, 0x2e, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x82, 0xec, 0xa9, 0xb6, 0x82, 0x01,
	0x00, 0x00,
}
