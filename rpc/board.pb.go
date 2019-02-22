// Code generated by protoc-gen-go. DO NOT EDIT.
// source: board.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type BoardDetailsReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Fqbn                 string    `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardDetailsReq) Reset()         { *m = BoardDetailsReq{} }
func (m *BoardDetailsReq) String() string { return proto.CompactTextString(m) }
func (*BoardDetailsReq) ProtoMessage()    {}
func (*BoardDetailsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{0}
}

func (m *BoardDetailsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardDetailsReq.Unmarshal(m, b)
}
func (m *BoardDetailsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardDetailsReq.Marshal(b, m, deterministic)
}
func (m *BoardDetailsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardDetailsReq.Merge(m, src)
}
func (m *BoardDetailsReq) XXX_Size() int {
	return xxx_messageInfo_BoardDetailsReq.Size(m)
}
func (m *BoardDetailsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardDetailsReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardDetailsReq proto.InternalMessageInfo

func (m *BoardDetailsReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *BoardDetailsReq) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

type BoardDetailsResp struct {
	Result               *Result         `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ConfigOptions        []*ConfigOption `protobuf:"bytes,3,rep,name=config_options,json=configOptions,proto3" json:"config_options,omitempty"`
	RequiredTools        []*RequiredTool `protobuf:"bytes,4,rep,name=required_tools,json=requiredTools,proto3" json:"required_tools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BoardDetailsResp) Reset()         { *m = BoardDetailsResp{} }
func (m *BoardDetailsResp) String() string { return proto.CompactTextString(m) }
func (*BoardDetailsResp) ProtoMessage()    {}
func (*BoardDetailsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{1}
}

func (m *BoardDetailsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardDetailsResp.Unmarshal(m, b)
}
func (m *BoardDetailsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardDetailsResp.Marshal(b, m, deterministic)
}
func (m *BoardDetailsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardDetailsResp.Merge(m, src)
}
func (m *BoardDetailsResp) XXX_Size() int {
	return xxx_messageInfo_BoardDetailsResp.Size(m)
}
func (m *BoardDetailsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardDetailsResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardDetailsResp proto.InternalMessageInfo

func (m *BoardDetailsResp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *BoardDetailsResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BoardDetailsResp) GetConfigOptions() []*ConfigOption {
	if m != nil {
		return m.ConfigOptions
	}
	return nil
}

func (m *BoardDetailsResp) GetRequiredTools() []*RequiredTool {
	if m != nil {
		return m.RequiredTools
	}
	return nil
}

type ConfigOption struct {
	Option               string         `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
	OptionLabel          string         `protobuf:"bytes,2,opt,name=option_label,json=optionLabel,proto3" json:"option_label,omitempty"`
	Values               []*ConfigValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfigOption) Reset()         { *m = ConfigOption{} }
func (m *ConfigOption) String() string { return proto.CompactTextString(m) }
func (*ConfigOption) ProtoMessage()    {}
func (*ConfigOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{2}
}

func (m *ConfigOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigOption.Unmarshal(m, b)
}
func (m *ConfigOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigOption.Marshal(b, m, deterministic)
}
func (m *ConfigOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigOption.Merge(m, src)
}
func (m *ConfigOption) XXX_Size() int {
	return xxx_messageInfo_ConfigOption.Size(m)
}
func (m *ConfigOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigOption.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigOption proto.InternalMessageInfo

func (m *ConfigOption) GetOption() string {
	if m != nil {
		return m.Option
	}
	return ""
}

func (m *ConfigOption) GetOptionLabel() string {
	if m != nil {
		return m.OptionLabel
	}
	return ""
}

func (m *ConfigOption) GetValues() []*ConfigValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type ConfigValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ValueLabel           string   `protobuf:"bytes,2,opt,name=value_label,json=valueLabel,proto3" json:"value_label,omitempty"`
	Selected             bool     `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigValue) Reset()         { *m = ConfigValue{} }
func (m *ConfigValue) String() string { return proto.CompactTextString(m) }
func (*ConfigValue) ProtoMessage()    {}
func (*ConfigValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{3}
}

func (m *ConfigValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigValue.Unmarshal(m, b)
}
func (m *ConfigValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigValue.Marshal(b, m, deterministic)
}
func (m *ConfigValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigValue.Merge(m, src)
}
func (m *ConfigValue) XXX_Size() int {
	return xxx_messageInfo_ConfigValue.Size(m)
}
func (m *ConfigValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigValue.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigValue proto.InternalMessageInfo

func (m *ConfigValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ConfigValue) GetValueLabel() string {
	if m != nil {
		return m.ValueLabel
	}
	return ""
}

func (m *ConfigValue) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type RequiredTool struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Packager             string   `protobuf:"bytes,3,opt,name=packager,proto3" json:"packager,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequiredTool) Reset()         { *m = RequiredTool{} }
func (m *RequiredTool) String() string { return proto.CompactTextString(m) }
func (*RequiredTool) ProtoMessage()    {}
func (*RequiredTool) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{4}
}

func (m *RequiredTool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequiredTool.Unmarshal(m, b)
}
func (m *RequiredTool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequiredTool.Marshal(b, m, deterministic)
}
func (m *RequiredTool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequiredTool.Merge(m, src)
}
func (m *RequiredTool) XXX_Size() int {
	return xxx_messageInfo_RequiredTool.Size(m)
}
func (m *RequiredTool) XXX_DiscardUnknown() {
	xxx_messageInfo_RequiredTool.DiscardUnknown(m)
}

var xxx_messageInfo_RequiredTool proto.InternalMessageInfo

func (m *RequiredTool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequiredTool) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RequiredTool) GetPackager() string {
	if m != nil {
		return m.Packager
	}
	return ""
}

func init() {
	proto.RegisterType((*BoardDetailsReq)(nil), "arduino.BoardDetailsReq")
	proto.RegisterType((*BoardDetailsResp)(nil), "arduino.BoardDetailsResp")
	proto.RegisterType((*ConfigOption)(nil), "arduino.ConfigOption")
	proto.RegisterType((*ConfigValue)(nil), "arduino.ConfigValue")
	proto.RegisterType((*RequiredTool)(nil), "arduino.RequiredTool")
}

func init() { proto.RegisterFile("board.proto", fileDescriptor_937f74b042f92c0f) }

var fileDescriptor_937f74b042f92c0f = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x41, 0x8f, 0xd3, 0x30,
	0x10, 0x85, 0x15, 0xba, 0x64, 0xdb, 0x49, 0x96, 0x05, 0x6b, 0x41, 0xd1, 0x5e, 0x08, 0x11, 0x12,
	0x39, 0xb0, 0x59, 0xa9, 0x5c, 0x39, 0x15, 0x2e, 0x48, 0x48, 0x48, 0x56, 0x85, 0x10, 0x97, 0xe2,
	0x38, 0x6e, 0xb1, 0x70, 0xec, 0xd4, 0x76, 0xca, 0x8f, 0xe4, 0x4f, 0x21, 0x3b, 0x4e, 0x94, 0xf6,
	0x94, 0x79, 0xef, 0xd9, 0xdf, 0xcc, 0x58, 0x81, 0xa4, 0x56, 0x44, 0x37, 0x55, 0xa7, 0x95, 0x55,
	0xe8, 0x9a, 0xe8, 0xa6, 0xe7, 0x52, 0xdd, 0xa7, 0x54, 0xb5, 0xad, 0x92, 0x83, 0x5d, 0x6c, 0xe1,
	0x76, 0xe3, 0x4e, 0x7d, 0x66, 0x96, 0x70, 0x61, 0x30, 0x3b, 0xa2, 0x07, 0x58, 0x72, 0x69, 0x2c,
	0x91, 0x94, 0x65, 0x51, 0x1e, 0x95, 0xc9, 0xfa, 0x45, 0x15, 0x2e, 0x57, 0x5f, 0x42, 0x80, 0xa7,
	0x23, 0x08, 0xc1, 0xd5, 0xfe, 0x58, 0xcb, 0xec, 0x49, 0x1e, 0x95, 0x2b, 0xec, 0xeb, 0xe2, 0x5f,
	0x04, 0xcf, 0xcf, 0xb1, 0xa6, 0x43, 0xef, 0x20, 0xd6, 0xcc, 0xf4, 0xc2, 0x06, 0xea, 0xed, 0x44,
	0xc5, 0xde, 0xc6, 0x21, 0x76, 0x44, 0x49, 0x5a, 0x36, 0x12, 0x5d, 0x8d, 0x3e, 0xc2, 0x33, 0xaa,
	0xe4, 0x9e, 0x1f, 0x76, 0xaa, 0xb3, 0x5c, 0x49, 0x93, 0x2d, 0xf2, 0x45, 0x99, 0xac, 0x5f, 0x4e,
	0x90, 0x4f, 0x3e, 0xfe, 0xe6, 0x53, 0x7c, 0x43, 0x67, 0xca, 0xb8, 0xdb, 0x9a, 0x1d, 0x7b, 0xae,
	0x59, 0xb3, 0xb3, 0x4a, 0x09, 0x93, 0x5d, 0x5d, 0xdc, 0xc6, 0x21, 0xde, 0x2a, 0x25, 0xf0, 0x8d,
	0x9e, 0x29, 0x53, 0xfc, 0x85, 0x74, 0x0e, 0x47, 0xaf, 0x20, 0x1e, 0x86, 0xf0, 0x8b, 0xac, 0x70,
	0x50, 0xe8, 0x0d, 0xa4, 0x43, 0xb5, 0x13, 0xa4, 0x66, 0x22, 0xcc, 0x9f, 0x0c, 0xde, 0x57, 0x67,
	0xa1, 0xf7, 0x10, 0x9f, 0x88, 0xe8, 0xd9, 0x38, 0xfe, 0xdd, 0xc5, 0xf8, 0xdf, 0x5d, 0x88, 0xc3,
	0x99, 0xe2, 0x17, 0x24, 0x33, 0x1b, 0xdd, 0xc1, 0x53, 0x1f, 0x84, 0xb6, 0x83, 0x40, 0xaf, 0x21,
	0xf1, 0xc5, 0x59, 0x53, 0xf0, 0xd6, 0xd0, 0xf3, 0x1e, 0x96, 0x86, 0x09, 0x46, 0x2d, 0x6b, 0xb2,
	0x45, 0x1e, 0x95, 0x4b, 0x3c, 0xe9, 0xe2, 0x07, 0xa4, 0xf3, 0xcd, 0xa7, 0xa7, 0x8f, 0x66, 0x4f,
	0x9f, 0xc1, 0xf5, 0x89, 0x69, 0xe3, 0xf6, 0x1d, 0xe0, 0xa3, 0x74, 0xe4, 0x8e, 0xd0, 0x3f, 0xe4,
	0xc0, 0xb4, 0x27, 0xaf, 0xf0, 0xa4, 0x37, 0x6f, 0x7f, 0x16, 0x07, 0x6e, 0x7f, 0xf7, 0x75, 0x45,
	0x55, 0xfb, 0x18, 0xb6, 0x1c, 0xbf, 0x0f, 0x54, 0xf0, 0x47, 0xdd, 0xd1, 0x3a, 0xf6, 0x7f, 0xe1,
	0x87, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x83, 0xcb, 0x9d, 0xab, 0x02, 0x00, 0x00,
}
