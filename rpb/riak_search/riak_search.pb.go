// Code generated by protoc-gen-go.
// source: riak_search.proto
// DO NOT EDIT!

/*
Package riak_search is a generated protocol buffer package.

It is generated from these files:
	riak_search.proto

It has these top-level messages:
	RpbSearchDoc
	RpbSearchQueryReq
	RpbSearchQueryResp
*/
package riak_search

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import riak "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type RpbSearchDoc struct {
	Fields           []*riak.RpbPair `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbSearchDoc) Reset()                    { *m = RpbSearchDoc{} }
func (m *RpbSearchDoc) String() string            { return proto.CompactTextString(m) }
func (*RpbSearchDoc) ProtoMessage()               {}
func (*RpbSearchDoc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RpbSearchDoc) GetFields() []*riak.RpbPair {
	if m != nil {
		return m.Fields
	}
	return nil
}

type RpbSearchQueryReq struct {
	Q                []byte   `protobuf:"bytes,1,req,name=q" json:"q,omitempty"`
	Index            []byte   `protobuf:"bytes,2,req,name=index" json:"index,omitempty"`
	Rows             *uint32  `protobuf:"varint,3,opt,name=rows" json:"rows,omitempty"`
	Start            *uint32  `protobuf:"varint,4,opt,name=start" json:"start,omitempty"`
	Sort             []byte   `protobuf:"bytes,5,opt,name=sort" json:"sort,omitempty"`
	Filter           []byte   `protobuf:"bytes,6,opt,name=filter" json:"filter,omitempty"`
	Df               []byte   `protobuf:"bytes,7,opt,name=df" json:"df,omitempty"`
	Op               []byte   `protobuf:"bytes,8,opt,name=op" json:"op,omitempty"`
	Fl               [][]byte `protobuf:"bytes,9,rep,name=fl" json:"fl,omitempty"`
	Presort          []byte   `protobuf:"bytes,10,opt,name=presort" json:"presort,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RpbSearchQueryReq) Reset()                    { *m = RpbSearchQueryReq{} }
func (m *RpbSearchQueryReq) String() string            { return proto.CompactTextString(m) }
func (*RpbSearchQueryReq) ProtoMessage()               {}
func (*RpbSearchQueryReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RpbSearchQueryReq) GetQ() []byte {
	if m != nil {
		return m.Q
	}
	return nil
}

func (m *RpbSearchQueryReq) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *RpbSearchQueryReq) GetRows() uint32 {
	if m != nil && m.Rows != nil {
		return *m.Rows
	}
	return 0
}

func (m *RpbSearchQueryReq) GetStart() uint32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *RpbSearchQueryReq) GetSort() []byte {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *RpbSearchQueryReq) GetFilter() []byte {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *RpbSearchQueryReq) GetDf() []byte {
	if m != nil {
		return m.Df
	}
	return nil
}

func (m *RpbSearchQueryReq) GetOp() []byte {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *RpbSearchQueryReq) GetFl() [][]byte {
	if m != nil {
		return m.Fl
	}
	return nil
}

func (m *RpbSearchQueryReq) GetPresort() []byte {
	if m != nil {
		return m.Presort
	}
	return nil
}

type RpbSearchQueryResp struct {
	Docs             []*RpbSearchDoc `protobuf:"bytes,1,rep,name=docs" json:"docs,omitempty"`
	MaxScore         *float32        `protobuf:"fixed32,2,opt,name=max_score" json:"max_score,omitempty"`
	NumFound         *uint32         `protobuf:"varint,3,opt,name=num_found" json:"num_found,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbSearchQueryResp) Reset()                    { *m = RpbSearchQueryResp{} }
func (m *RpbSearchQueryResp) String() string            { return proto.CompactTextString(m) }
func (*RpbSearchQueryResp) ProtoMessage()               {}
func (*RpbSearchQueryResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RpbSearchQueryResp) GetDocs() []*RpbSearchDoc {
	if m != nil {
		return m.Docs
	}
	return nil
}

func (m *RpbSearchQueryResp) GetMaxScore() float32 {
	if m != nil && m.MaxScore != nil {
		return *m.MaxScore
	}
	return 0
}

func (m *RpbSearchQueryResp) GetNumFound() uint32 {
	if m != nil && m.NumFound != nil {
		return *m.NumFound
	}
	return 0
}

func init() {
	proto.RegisterType((*RpbSearchDoc)(nil), "RpbSearchDoc")
	proto.RegisterType((*RpbSearchQueryReq)(nil), "RpbSearchQueryReq")
	proto.RegisterType((*RpbSearchQueryResp)(nil), "RpbSearchQueryResp")
}

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4e, 0xbc, 0x30,
	0x10, 0xc7, 0x03, 0xcb, 0xfe, 0x61, 0x7e, 0xe5, 0x67, 0xe8, 0xc5, 0x46, 0x2f, 0x1b, 0x2e, 0x72,
	0xe2, 0xe0, 0x23, 0x10, 0x1f, 0x60, 0xc5, 0x83, 0x47, 0x52, 0xa0, 0x64, 0x89, 0xc0, 0xb0, 0x53,
	0x88, 0xeb, 0xc3, 0xf8, 0xae, 0x96, 0x12, 0x34, 0x7a, 0x9b, 0xcf, 0x7c, 0x26, 0xd3, 0xef, 0x14,
	0x42, 0x6a, 0xe4, 0x5b, 0xae, 0x95, 0xa4, 0xf2, 0x9c, 0x0c, 0x84, 0x23, 0xde, 0xc1, 0xdc, 0x5a,
	0xea, 0x28, 0x06, 0x96, 0x0d, 0xc5, 0x8b, 0xd5, 0x4f, 0x58, 0x72, 0x01, 0xbb, 0xba, 0x51, 0x6d,
	0xa5, 0x85, 0x73, 0xdc, 0xc4, 0xff, 0x1e, 0x0f, 0x89, 0xd1, 0x27, 0xd9, 0x50, 0xf4, 0xe9, 0x40,
	0xf8, 0x3d, 0xfa, 0x3c, 0x29, 0xfa, 0xc8, 0xd4, 0x85, 0xfb, 0xe0, 0x5c, 0xcc, 0xa8, 0x1b, 0x33,
	0x1e, 0xc0, 0xb6, 0xe9, 0x2b, 0x75, 0x15, 0xae, 0x45, 0x06, 0x1e, 0xe1, 0xbb, 0x16, 0x9b, 0xa3,
	0x13, 0x07, 0xb3, 0xd4, 0xa3, 0xa4, 0x51, 0x78, 0x16, 0x8d, 0xd4, 0x68, 0x68, 0x6b, 0x88, 0xf1,
	0xff, 0xf3, 0xa3, 0xed, 0xa8, 0x48, 0xec, 0x2c, 0x03, 0xb8, 0x55, 0x2d, 0xf6, 0x6b, 0x8d, 0x83,
	0x38, 0xac, 0x75, 0xdd, 0x0a, 0xdf, 0x04, 0x63, 0xfc, 0x06, 0xf6, 0x03, 0x29, 0xbb, 0x04, 0x66,
	0x19, 0xbd, 0x02, 0xff, 0x1b, 0x4f, 0x0f, 0xfc, 0x1e, 0xbc, 0x0a, 0xcb, 0xf5, 0x9a, 0x20, 0xf9,
	0x75, 0x6c, 0x08, 0x7e, 0x27, 0xaf, 0xb9, 0x2e, 0x91, 0x94, 0x49, 0xed, 0xc4, 0xee, 0xdc, 0xea,
	0xa7, 0x2e, 0xaf, 0x71, 0xea, 0xab, 0x25, 0x7a, 0xfa, 0x00, 0xb7, 0x25, 0x76, 0x49, 0x21, 0xf5,
	0x19, 0x93, 0x9f, 0xaf, 0x2b, 0xa6, 0x3a, 0x65, 0x99, 0xc1, 0x65, 0xdf, 0x29, 0xfd, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0xf4, 0xd4, 0x3b, 0x37, 0x6b, 0x01, 0x00, 0x00,
}
