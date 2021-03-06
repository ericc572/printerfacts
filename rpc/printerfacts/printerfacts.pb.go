// Code generated by protoc-gen-go. DO NOT EDIT.
// source: printerfacts.proto

/*
Package printerfacts is a generated protocol buffer package.

It is generated from these files:
	printerfacts.proto

It has these top-level messages:
	FactParams
	Facts
*/
package printerfacts

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// FactParams stores paramaters about what printer facts a user wants.
type FactParams struct {
	// count is the number of facts to return.
	Count int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *FactParams) Reset()                    { *m = FactParams{} }
func (m *FactParams) String() string            { return proto.CompactTextString(m) }
func (*FactParams) ProtoMessage()               {}
func (*FactParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FactParams) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// Facts is the resulting list of printer facts.
type Facts struct {
	// facts is the individual string facts about printers.
	Facts []string `protobuf:"bytes,1,rep,name=facts" json:"facts,omitempty"`
}

func (m *Facts) Reset()                    { *m = Facts{} }
func (m *Facts) String() string            { return proto.CompactTextString(m) }
func (*Facts) ProtoMessage()               {}
func (*Facts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Facts) GetFacts() []string {
	if m != nil {
		return m.Facts
	}
	return nil
}

func init() {
	proto.RegisterType((*FactParams)(nil), "us.xeserv.api.FactParams")
	proto.RegisterType((*Facts)(nil), "us.xeserv.api.Facts")
}

func init() { proto.RegisterFile("printerfacts.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0xcc,
	0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x2d, 0x2d, 0xd6, 0xab, 0x48, 0x2d, 0x4e, 0x2d, 0x2a, 0xd3, 0x4b, 0x2c, 0xc8, 0x54, 0x52, 0xe2,
	0xe2, 0x72, 0x4b, 0x4c, 0x2e, 0x09, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0x16, 0x12, 0xe1, 0x62, 0x4d,
	0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x70, 0x94, 0x64, 0xb9,
	0x58, 0x41, 0x6a, 0xc0, 0xd2, 0x60, 0xa3, 0x24, 0x18, 0x15, 0x98, 0x35, 0x38, 0x83, 0x20, 0x1c,
	0x23, 0x4f, 0x2e, 0x9e, 0x00, 0x24, 0x7b, 0x84, 0x2c, 0xb9, 0x58, 0x40, 0xca, 0x85, 0x24, 0xf5,
	0x50, 0xac, 0xd2, 0x43, 0xd8, 0x23, 0x25, 0x82, 0x45, 0xaa, 0x58, 0x89, 0xc1, 0x89, 0x2f, 0x8a,
	0x07, 0xd9, 0xc9, 0x49, 0x6c, 0x60, 0x37, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x89, 0xf0,
	0x66, 0x04, 0xc9, 0x00, 0x00, 0x00,
}
