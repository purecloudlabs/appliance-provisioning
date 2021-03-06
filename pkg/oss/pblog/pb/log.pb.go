// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: BSD-3-Clause
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EvtType int32

const (
	EvtType_LOG        EvtType = 0
	EvtType_MSG        EvtType = 1
	EvtType_SUCCESS    EvtType = 2
	EvtType_ERROR      EvtType = 3
	EvtType_SESS_START EvtType = 4
	//following are used in server app; likely not useful on wire.
	EvtType_PRINT     EvtType = 5
	EvtType_PRINT_ERR EvtType = 6
	EvtType_CODENAME  EvtType = 7
	EvtType_STATE     EvtType = 8
	EvtType_SETPW     EvtType = 9
)

var EvtType_name = map[int32]string{
	0: "LOG",
	1: "MSG",
	2: "SUCCESS",
	3: "ERROR",
	4: "SESS_START",
	5: "PRINT",
	6: "PRINT_ERR",
	7: "CODENAME",
	8: "STATE",
	9: "SETPW",
}

var EvtType_value = map[string]int32{
	"LOG":        0,
	"MSG":        1,
	"SUCCESS":    2,
	"ERROR":      3,
	"SESS_START": 4,
	"PRINT":      5,
	"PRINT_ERR":  6,
	"CODENAME":   7,
	"STATE":      8,
	"SETPW":      9,
}

func (x EvtType) String() string {
	return proto.EnumName(EvtType_name, int32(x))
}

func (EvtType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{0}
}

type ProcessState int32

const (
	ProcessState_UNKNOWN        ProcessState = 0
	ProcessState_MFG_STARTED    ProcessState = 1
	ProcessState_MFG_SUCCEEDED  ProcessState = 2
	ProcessState_MFG_FAILED     ProcessState = 3
	ProcessState_FR_STARTED     ProcessState = 4
	ProcessState_FR_SUCCEEDED   ProcessState = 5
	ProcessState_FR_FAILED      ProcessState = 6
	ProcessState_INIT_STARTED   ProcessState = 7
	ProcessState_INIT_SUCCEEDED ProcessState = 8
	ProcessState_INIT_FAILED    ProcessState = 9
)

var ProcessState_name = map[int32]string{
	0: "UNKNOWN",
	1: "MFG_STARTED",
	2: "MFG_SUCCEEDED",
	3: "MFG_FAILED",
	4: "FR_STARTED",
	5: "FR_SUCCEEDED",
	6: "FR_FAILED",
	7: "INIT_STARTED",
	8: "INIT_SUCCEEDED",
	9: "INIT_FAILED",
}

var ProcessState_value = map[string]int32{
	"UNKNOWN":        0,
	"MFG_STARTED":    1,
	"MFG_SUCCEEDED":  2,
	"MFG_FAILED":     3,
	"FR_STARTED":     4,
	"FR_SUCCEEDED":   5,
	"FR_FAILED":      6,
	"INIT_STARTED":   7,
	"INIT_SUCCEEDED": 8,
	"INIT_FAILED":    9,
}

func (x ProcessState) String() string {
	return proto.EnumName(ProcessState_name, int32(x))
}

func (ProcessState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{1}
}

type LogEvent struct {
	SN                   string     `protobuf:"bytes,1,opt,name=SN,proto3" json:"SN,omitempty"`
	EventType            EvtType    `protobuf:"varint,2,opt,name=EventType,proto3,enum=pb.EvtType" json:"EventType,omitempty"`
	Time                 *Timestamp `protobuf:"bytes,3,opt,name=Time,proto3" json:"Time,omitempty"`
	Payload              string     `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LogEvent) Reset()         { *m = LogEvent{} }
func (m *LogEvent) String() string { return proto.CompactTextString(m) }
func (*LogEvent) ProtoMessage()    {}
func (*LogEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{0}
}

func (m *LogEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEvent.Unmarshal(m, b)
}
func (m *LogEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEvent.Marshal(b, m, deterministic)
}
func (m *LogEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEvent.Merge(m, src)
}
func (m *LogEvent) XXX_Size() int {
	return xxx_messageInfo_LogEvent.Size(m)
}
func (m *LogEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LogEvent proto.InternalMessageInfo

func (m *LogEvent) GetSN() string {
	if m != nil {
		return m.SN
	}
	return ""
}

func (m *LogEvent) GetEventType() EvtType {
	if m != nil {
		return m.EventType
	}
	return EvtType_LOG
}

func (m *LogEvent) GetTime() *Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *LogEvent) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

//used for storage of multiple le's
type LogEvents struct {
	Evt                  []*LogEvent `protobuf:"bytes,1,rep,name=Evt,proto3" json:"Evt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LogEvents) Reset()         { *m = LogEvents{} }
func (m *LogEvents) String() string { return proto.CompactTextString(m) }
func (*LogEvents) ProtoMessage()    {}
func (*LogEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{1}
}

func (m *LogEvents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEvents.Unmarshal(m, b)
}
func (m *LogEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEvents.Marshal(b, m, deterministic)
}
func (m *LogEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEvents.Merge(m, src)
}
func (m *LogEvents) XXX_Size() int {
	return xxx_messageInfo_LogEvents.Size(m)
}
func (m *LogEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEvents.DiscardUnknown(m)
}

var xxx_messageInfo_LogEvents proto.InternalMessageInfo

func (m *LogEvents) GetEvt() []*LogEvent {
	if m != nil {
		return m.Evt
	}
	return nil
}

type Timestamp struct {
	TS                   int64    `protobuf:"varint,1,opt,name=TS,proto3" json:"TS,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{2}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetTS() int64 {
	if m != nil {
		return m.TS
	}
	return 0
}

type GenericResponse struct {
	EventType            EvtType  `protobuf:"varint,1,opt,name=EventType,proto3,enum=pb.EvtType" json:"EventType,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResponse) Reset()         { *m = GenericResponse{} }
func (m *GenericResponse) String() string { return proto.CompactTextString(m) }
func (*GenericResponse) ProtoMessage()    {}
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{3}
}

func (m *GenericResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponse.Unmarshal(m, b)
}
func (m *GenericResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponse.Marshal(b, m, deterministic)
}
func (m *GenericResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponse.Merge(m, src)
}
func (m *GenericResponse) XXX_Size() int {
	return xxx_messageInfo_GenericResponse.Size(m)
}
func (m *GenericResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponse proto.InternalMessageInfo

func (m *GenericResponse) GetEventType() EvtType {
	if m != nil {
		return m.EventType
	}
	return EvtType_LOG
}

func (m *GenericResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type MACs struct {
	SN                   string   `protobuf:"bytes,1,opt,name=SN,proto3" json:"SN,omitempty"`
	MAC                  []string `protobuf:"bytes,2,rep,name=MAC,proto3" json:"MAC,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MACs) Reset()         { *m = MACs{} }
func (m *MACs) String() string { return proto.CompactTextString(m) }
func (*MACs) ProtoMessage()    {}
func (*MACs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{4}
}

func (m *MACs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MACs.Unmarshal(m, b)
}
func (m *MACs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MACs.Marshal(b, m, deterministic)
}
func (m *MACs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MACs.Merge(m, src)
}
func (m *MACs) XXX_Size() int {
	return xxx_messageInfo_MACs.Size(m)
}
func (m *MACs) XXX_DiscardUnknown() {
	xxx_messageInfo_MACs.DiscardUnknown(m)
}

var xxx_messageInfo_MACs proto.InternalMessageInfo

func (m *MACs) GetSN() string {
	if m != nil {
		return m.SN
	}
	return ""
}

func (m *MACs) GetMAC() []string {
	if m != nil {
		return m.MAC
	}
	return nil
}

type ProcessStage struct {
	SN                   string       `protobuf:"bytes,1,opt,name=SN,proto3" json:"SN,omitempty"`
	State                ProcessState `protobuf:"varint,2,opt,name=State,proto3,enum=pb.ProcessState" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProcessStage) Reset()         { *m = ProcessStage{} }
func (m *ProcessStage) String() string { return proto.CompactTextString(m) }
func (*ProcessStage) ProtoMessage()    {}
func (*ProcessStage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{5}
}

func (m *ProcessStage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessStage.Unmarshal(m, b)
}
func (m *ProcessStage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessStage.Marshal(b, m, deterministic)
}
func (m *ProcessStage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessStage.Merge(m, src)
}
func (m *ProcessStage) XXX_Size() int {
	return xxx_messageInfo_ProcessStage.Size(m)
}
func (m *ProcessStage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessStage.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessStage proto.InternalMessageInfo

func (m *ProcessStage) GetSN() string {
	if m != nil {
		return m.SN
	}
	return ""
}

func (m *ProcessStage) GetState() ProcessState {
	if m != nil {
		return m.State
	}
	return ProcessState_UNKNOWN
}

type Codename struct {
	SN                   string   `protobuf:"bytes,1,opt,name=SN,proto3" json:"SN,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Codename) Reset()         { *m = Codename{} }
func (m *Codename) String() string { return proto.CompactTextString(m) }
func (*Codename) ProtoMessage()    {}
func (*Codename) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{6}
}

func (m *Codename) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Codename.Unmarshal(m, b)
}
func (m *Codename) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Codename.Marshal(b, m, deterministic)
}
func (m *Codename) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Codename.Merge(m, src)
}
func (m *Codename) XXX_Size() int {
	return xxx_messageInfo_Codename.Size(m)
}
func (m *Codename) XXX_DiscardUnknown() {
	xxx_messageInfo_Codename.DiscardUnknown(m)
}

var xxx_messageInfo_Codename proto.InternalMessageInfo

func (m *Codename) GetSN() string {
	if m != nil {
		return m.SN
	}
	return ""
}

func (m *Codename) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Document struct {
	SN                   string   `protobuf:"bytes,1,opt,name=SN,proto3" json:"SN,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Doctype              string   `protobuf:"bytes,3,opt,name=Doctype,proto3" json:"Doctype,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{7}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetSN() string {
	if m != nil {
		return m.SN
	}
	return ""
}

func (m *Document) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Document) GetDoctype() string {
	if m != nil {
		return m.Doctype
	}
	return ""
}

func (m *Document) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Identifier struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{8}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Credentials struct {
	OS                   string   `protobuf:"bytes,1,opt,name=OS,proto3" json:"OS,omitempty"`
	BIOS                 string   `protobuf:"bytes,2,opt,name=BIOS,proto3" json:"BIOS,omitempty"`
	IPMI                 string   `protobuf:"bytes,3,opt,name=IPMI,proto3" json:"IPMI,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{9}
}

func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (m *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(m, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetOS() string {
	if m != nil {
		return m.OS
	}
	return ""
}

func (m *Credentials) GetBIOS() string {
	if m != nil {
		return m.BIOS
	}
	return ""
}

func (m *Credentials) GetIPMI() string {
	if m != nil {
		return m.IPMI
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.EvtType", EvtType_name, EvtType_value)
	proto.RegisterEnum("pb.ProcessState", ProcessState_name, ProcessState_value)
	proto.RegisterType((*LogEvent)(nil), "pb.LogEvent")
	proto.RegisterType((*LogEvents)(nil), "pb.LogEvents")
	proto.RegisterType((*Timestamp)(nil), "pb.Timestamp")
	proto.RegisterType((*GenericResponse)(nil), "pb.GenericResponse")
	proto.RegisterType((*MACs)(nil), "pb.MACs")
	proto.RegisterType((*ProcessStage)(nil), "pb.ProcessStage")
	proto.RegisterType((*Codename)(nil), "pb.Codename")
	proto.RegisterType((*Document)(nil), "pb.Document")
	proto.RegisterType((*Identifier)(nil), "pb.Identifier")
	proto.RegisterType((*Credentials)(nil), "pb.Credentials")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor_a153da538f858886) }

var fileDescriptor_a153da538f858886 = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xf2, 0x46,
	0x10, 0x7e, 0x6d, 0x13, 0x8c, 0x87, 0x8f, 0x6c, 0xb7, 0x52, 0x84, 0x48, 0x55, 0x51, 0x1f, 0x2a,
	0x9a, 0x4a, 0x4e, 0x4b, 0xaa, 0x9e, 0xaa, 0x4a, 0x14, 0x16, 0x64, 0x05, 0x0c, 0xda, 0x75, 0x94,
	0x4b, 0xa5, 0xc8, 0xc0, 0xc6, 0x42, 0x05, 0xd6, 0xb2, 0x37, 0x48, 0x1c, 0xdb, 0x5b, 0x7f, 0x51,
	0xff, 0x5e, 0xb5, 0x6b, 0x0c, 0x25, 0x41, 0xd5, 0x7b, 0x9b, 0xaf, 0x67, 0xf6, 0x99, 0x99, 0x9d,
	0x01, 0x67, 0x2d, 0x62, 0x2f, 0x49, 0x85, 0x14, 0xd8, 0x4c, 0xe6, 0xad, 0xdb, 0x58, 0x88, 0x78,
	0xcd, 0xef, 0xb5, 0x65, 0xfe, 0xf6, 0x7a, 0xcf, 0x37, 0x89, 0xdc, 0xe7, 0x01, 0xee, 0x5f, 0x06,
	0x54, 0xc6, 0x22, 0x26, 0x3b, 0xbe, 0x95, 0xb8, 0x01, 0x26, 0x0b, 0x9a, 0x46, 0xdb, 0xe8, 0x38,
	0xd4, 0x64, 0x01, 0xfe, 0x0e, 0x1c, 0xed, 0x08, 0xf7, 0x09, 0x6f, 0x9a, 0x6d, 0xa3, 0xd3, 0xe8,
	0x56, 0xbd, 0x64, 0xee, 0x91, 0x9d, 0x36, 0xd1, 0x93, 0x17, 0x7f, 0x03, 0xa5, 0x70, 0xb5, 0xe1,
	0x4d, 0xab, 0x6d, 0x74, 0xaa, 0xdd, 0xba, 0x8a, 0x52, 0x7a, 0x26, 0xa3, 0x4d, 0x42, 0xb5, 0x0b,
	0x37, 0xc1, 0x9e, 0x45, 0xfb, 0xb5, 0x88, 0x96, 0xcd, 0x92, 0x7e, 0xa2, 0x50, 0xdd, 0xef, 0xc1,
	0x29, 0x38, 0x64, 0xf8, 0x6b, 0xb0, 0xc8, 0x4e, 0x36, 0x8d, 0xb6, 0xd5, 0xa9, 0x76, 0x6b, 0x2a,
	0x51, 0xe1, 0xa3, 0xca, 0xe1, 0xde, 0x82, 0x73, 0xcc, 0xac, 0x18, 0x87, 0x4c, 0x33, 0xb6, 0xa8,
	0x19, 0x32, 0x37, 0x84, 0xeb, 0x11, 0xdf, 0xf2, 0x74, 0xb5, 0xa0, 0x3c, 0x4b, 0xc4, 0x36, 0xe3,
	0xe7, 0x45, 0x18, 0xff, 0x5b, 0xc4, 0x0d, 0x94, 0x49, 0x9a, 0x4e, 0xb2, 0x58, 0x17, 0xeb, 0xd0,
	0x83, 0xe6, 0x76, 0xa0, 0x34, 0xe9, 0xf5, 0xb3, 0x0f, 0xfd, 0x41, 0x60, 0x4d, 0x7a, 0xfd, 0xa6,
	0xd9, 0xb6, 0x3a, 0x0e, 0x55, 0xa2, 0x3b, 0x84, 0xda, 0x2c, 0x15, 0x0b, 0x9e, 0x65, 0x4c, 0x46,
	0x31, 0xff, 0x80, 0xf8, 0x16, 0xae, 0x98, 0x8c, 0x64, 0xd1, 0x4d, 0xa4, 0x88, 0x9c, 0x00, 0x92,
	0xd3, 0xdc, 0xed, 0x7a, 0x50, 0xe9, 0x8b, 0x25, 0xdf, 0x46, 0x9b, 0x8f, 0x39, 0x30, 0x94, 0x82,
	0x68, 0xc3, 0x0f, 0x1c, 0xb5, 0xec, 0xfe, 0x0e, 0x95, 0x81, 0x58, 0xbc, 0x6d, 0x2e, 0x4d, 0xf1,
	0x42, 0xbc, 0x9a, 0xc5, 0x40, 0x2c, 0xa4, 0x6a, 0x89, 0x95, 0xcf, 0xe2, 0xa0, 0xaa, 0xe8, 0xdf,
	0xc4, 0x72, 0xaf, 0x47, 0x54, 0xa3, 0x5a, 0x76, 0xbf, 0x02, 0xf0, 0x97, 0x7c, 0x2b, 0x57, 0xaf,
	0x2b, 0x9e, 0xaa, 0xfc, 0xfe, 0xb2, 0xc8, 0xef, 0x2f, 0x5d, 0x02, 0xd5, 0x7e, 0xca, 0xb5, 0x3f,
	0x5a, 0xeb, 0x26, 0x4d, 0x59, 0xe1, 0x9e, 0x32, 0x9d, 0xd0, 0x9f, 0xb2, 0xe2, 0x79, 0x25, 0x2b,
	0x9b, 0x3f, 0x9b, 0xf8, 0x87, 0xb7, 0xb5, 0x7c, 0xf7, 0xa7, 0x01, 0xf6, 0x61, 0x26, 0xd8, 0x06,
	0x6b, 0x3c, 0x1d, 0xa1, 0x4f, 0x4a, 0x98, 0xb0, 0x11, 0x32, 0x70, 0x15, 0x6c, 0xf6, 0xd4, 0xef,
	0x13, 0xc6, 0x90, 0x89, 0x1d, 0xb8, 0x22, 0x94, 0x4e, 0x29, 0xb2, 0x70, 0x03, 0x80, 0x11, 0xc6,
	0x5e, 0x58, 0xd8, 0xa3, 0x21, 0x2a, 0x29, 0xd7, 0x8c, 0xfa, 0x41, 0x88, 0xae, 0x70, 0x1d, 0x1c,
	0x2d, 0xbe, 0x10, 0x4a, 0x51, 0x19, 0xd7, 0xa0, 0xd2, 0x9f, 0x0e, 0x48, 0xd0, 0x9b, 0x10, 0x64,
	0xab, 0x38, 0x16, 0xf6, 0x42, 0x82, 0x2a, 0x5a, 0x24, 0xe1, 0xec, 0x19, 0x39, 0x77, 0xff, 0x18,
	0xff, 0x9d, 0x9f, 0xe4, 0xea, 0xd9, 0xa7, 0xe0, 0x31, 0x98, 0x3e, 0x07, 0xe8, 0x13, 0xbe, 0x86,
	0xea, 0x64, 0x38, 0xca, 0x9f, 0x22, 0x03, 0x64, 0xe0, 0x2f, 0xa0, 0xae, 0x0d, 0x8a, 0x18, 0x19,
	0x90, 0x01, 0x32, 0x15, 0x1f, 0x65, 0x1a, 0xf6, 0xfc, 0x31, 0x19, 0xe4, 0xfc, 0x86, 0xf4, 0x08,
	0x29, 0x61, 0x04, 0x35, 0xa5, 0x1f, 0x11, 0x9a, 0xe6, 0x90, 0x16, 0x80, 0xb2, 0x0a, 0xf0, 0x03,
	0x3f, 0x3c, 0x42, 0x6c, 0x8c, 0xa1, 0x91, 0x5b, 0x8e, 0xa0, 0x8a, 0xa2, 0xa2, 0x6d, 0x07, 0x98,
	0xd3, 0xfd, 0x19, 0x60, 0x2c, 0x62, 0xc6, 0xd3, 0xdd, 0x6a, 0xc1, 0x71, 0x07, 0xac, 0xb1, 0x88,
	0xf1, 0xd9, 0xf6, 0xb4, 0xbe, 0x54, 0xda, 0xbb, 0xed, 0xe8, 0xfe, 0x6d, 0x42, 0x8d, 0xf2, 0x85,
	0x48, 0x97, 0x8f, 0x9c, 0x27, 0x3c, 0xc5, 0x77, 0xe0, 0x30, 0x29, 0x52, 0xae, 0x3f, 0x7c, 0x45,
	0x41, 0x94, 0x74, 0x11, 0x8c, 0x3d, 0xa8, 0xeb, 0x58, 0x35, 0xbf, 0xcf, 0x89, 0xff, 0x09, 0xaa,
	0x94, 0x27, 0x22, 0x95, 0x79, 0x73, 0xdf, 0xfd, 0xfe, 0x98, 0x5f, 0x46, 0x3d, 0x40, 0x23, 0x47,
	0x1d, 0x37, 0x42, 0xd7, 0x55, 0x68, 0x97, 0x41, 0xdd, 0x03, 0xb5, 0xe3, 0x56, 0x68, 0x4c, 0xa1,
	0x5d, 0xee, 0xc5, 0xaf, 0x00, 0xea, 0xb2, 0xfc, 0x91, 0x37, 0xe2, 0x07, 0xb0, 0x47, 0x5c, 0xea,
	0xcb, 0x75, 0xe3, 0xe5, 0x27, 0xd4, 0x2b, 0x4e, 0xa8, 0x47, 0xd4, 0x09, 0x6d, 0x9d, 0x9f, 0xb9,
	0xee, 0x2f, 0x60, 0x33, 0xbe, 0x48, 0xb9, 0xcc, 0xf0, 0x8f, 0xd0, 0x18, 0x71, 0x79, 0xb6, 0x16,
	0x2a, 0xf6, 0xb4, 0x45, 0xad, 0x6b, 0x5d, 0xc3, 0x29, 0x60, 0x5e, 0xd6, 0xc9, 0x1f, 0xfe, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xbf, 0x22, 0x93, 0x07, 0xbe, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogServiceClient interface {
	//stream or no?
	Log(ctx context.Context, in *LogEvent, opts ...grpc.CallOption) (*GenericResponse, error)
}

type logServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogServiceClient(cc *grpc.ClientConn) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) Log(ctx context.Context, in *LogEvent, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pb.LogService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
type LogServiceServer interface {
	//stream or no?
	Log(context.Context, *LogEvent) (*GenericResponse, error)
}

func RegisterLogServiceServer(s *grpc.Server, srv LogServiceServer) {
	s.RegisterService(&_LogService_serviceDesc, srv)
}

func _LogService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Log(ctx, req.(*LogEvent))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _LogService_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}

// RecordKeeperClient is the client API for RecordKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecordKeeperClient interface {
	StoreMACs(ctx context.Context, in *MACs, opts ...grpc.CallOption) (*GenericResponse, error)
	StoreIPMIMACs(ctx context.Context, in *MACs, opts ...grpc.CallOption) (*GenericResponse, error)
	ReportState(ctx context.Context, in *ProcessStage, opts ...grpc.CallOption) (*GenericResponse, error)
	ReportCodename(ctx context.Context, in *Codename, opts ...grpc.CallOption) (*GenericResponse, error)
	StoreDocument(ctx context.Context, in *Document, opts ...grpc.CallOption) (*GenericResponse, error)
}

type recordKeeperClient struct {
	cc *grpc.ClientConn
}

func NewRecordKeeperClient(cc *grpc.ClientConn) RecordKeeperClient {
	return &recordKeeperClient{cc}
}

func (c *recordKeeperClient) StoreMACs(ctx context.Context, in *MACs, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pb.RecordKeeper/StoreMACs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordKeeperClient) StoreIPMIMACs(ctx context.Context, in *MACs, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pb.RecordKeeper/StoreIPMIMACs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordKeeperClient) ReportState(ctx context.Context, in *ProcessStage, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pb.RecordKeeper/ReportState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordKeeperClient) ReportCodename(ctx context.Context, in *Codename, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pb.RecordKeeper/ReportCodename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordKeeperClient) StoreDocument(ctx context.Context, in *Document, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pb.RecordKeeper/StoreDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordKeeperServer is the server API for RecordKeeper service.
type RecordKeeperServer interface {
	StoreMACs(context.Context, *MACs) (*GenericResponse, error)
	StoreIPMIMACs(context.Context, *MACs) (*GenericResponse, error)
	ReportState(context.Context, *ProcessStage) (*GenericResponse, error)
	ReportCodename(context.Context, *Codename) (*GenericResponse, error)
	StoreDocument(context.Context, *Document) (*GenericResponse, error)
}

func RegisterRecordKeeperServer(s *grpc.Server, srv RecordKeeperServer) {
	s.RegisterService(&_RecordKeeper_serviceDesc, srv)
}

func _RecordKeeper_StoreMACs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MACs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordKeeperServer).StoreMACs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecordKeeper/StoreMACs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordKeeperServer).StoreMACs(ctx, req.(*MACs))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordKeeper_StoreIPMIMACs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MACs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordKeeperServer).StoreIPMIMACs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecordKeeper/StoreIPMIMACs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordKeeperServer).StoreIPMIMACs(ctx, req.(*MACs))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordKeeper_ReportState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessStage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordKeeperServer).ReportState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecordKeeper/ReportState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordKeeperServer).ReportState(ctx, req.(*ProcessStage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordKeeper_ReportCodename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Codename)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordKeeperServer).ReportCodename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecordKeeper/ReportCodename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordKeeperServer).ReportCodename(ctx, req.(*Codename))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordKeeper_StoreDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Document)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordKeeperServer).StoreDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecordKeeper/StoreDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordKeeperServer).StoreDocument(ctx, req.(*Document))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecordKeeper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RecordKeeper",
	HandlerType: (*RecordKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreMACs",
			Handler:    _RecordKeeper_StoreMACs_Handler,
		},
		{
			MethodName: "StoreIPMIMACs",
			Handler:    _RecordKeeper_StoreIPMIMACs_Handler,
		},
		{
			MethodName: "ReportState",
			Handler:    _RecordKeeper_ReportState_Handler,
		},
		{
			MethodName: "ReportCodename",
			Handler:    _RecordKeeper_ReportCodename_Handler,
		},
		{
			MethodName: "StoreDocument",
			Handler:    _RecordKeeper_StoreDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}

// TimekeeperClient is the client API for Timekeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimekeeperClient interface {
	GetTime(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Timestamp, error)
}

type timekeeperClient struct {
	cc *grpc.ClientConn
}

func NewTimekeeperClient(cc *grpc.ClientConn) TimekeeperClient {
	return &timekeeperClient{cc}
}

func (c *timekeeperClient) GetTime(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Timestamp, error) {
	out := new(Timestamp)
	err := c.cc.Invoke(ctx, "/pb.Timekeeper/GetTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimekeeperServer is the server API for Timekeeper service.
type TimekeeperServer interface {
	GetTime(context.Context, *empty.Empty) (*Timestamp, error)
}

func RegisterTimekeeperServer(s *grpc.Server, srv TimekeeperServer) {
	s.RegisterService(&_Timekeeper_serviceDesc, srv)
}

func _Timekeeper_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimekeeperServer).GetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Timekeeper/GetTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimekeeperServer).GetTime(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Timekeeper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Timekeeper",
	HandlerType: (*TimekeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTime",
			Handler:    _Timekeeper_GetTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}

// SecretsClient is the client API for Secrets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretsClient interface {
	GetCredentials(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Credentials, error)
}

type secretsClient struct {
	cc *grpc.ClientConn
}

func NewSecretsClient(cc *grpc.ClientConn) SecretsClient {
	return &secretsClient{cc}
}

func (c *secretsClient) GetCredentials(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Credentials, error) {
	out := new(Credentials)
	err := c.cc.Invoke(ctx, "/pb.Secrets/GetCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretsServer is the server API for Secrets service.
type SecretsServer interface {
	GetCredentials(context.Context, *Identifier) (*Credentials, error)
}

func RegisterSecretsServer(s *grpc.Server, srv SecretsServer) {
	s.RegisterService(&_Secrets_serviceDesc, srv)
}

func _Secrets_GetCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServer).GetCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Secrets/GetCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServer).GetCredentials(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _Secrets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Secrets",
	HandlerType: (*SecretsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCredentials",
			Handler:    _Secrets_GetCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}
