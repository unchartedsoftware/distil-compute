// Code generated by protoc-gen-go. DO NOT EDIT.
// source: value.proto

package pipeline

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

type NullValue int32

const (
	NullValue_NULL_VALUE NullValue = 0
)

var NullValue_name = map[int32]string{
	0: "NULL_VALUE",
}

var NullValue_value = map[string]int32{
	"NULL_VALUE": 0,
}

func (x NullValue) String() string {
	return proto.EnumName(NullValue_name, int32(x))
}

func (NullValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6d8b663a521ecf69, []int{0}
}

type ValueError struct {
	// A error message useful for debugging or logging. Not meant to be very end-user friendly.
	// If a list of supported/allowed value types could not support a given value, then message
	// should say so. On the other hand, if there was really an error using a value type which
	// would otherwise support a given value, then the error message should communicate this error.
	// If there was such an error but some later value type allowed for recovery, then there
	// should be no error.
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueError) Reset()         { *m = ValueError{} }
func (m *ValueError) String() string { return proto.CompactTextString(m) }
func (*ValueError) ProtoMessage()    {}
func (*ValueError) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d8b663a521ecf69, []int{0}
}

func (m *ValueError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueError.Unmarshal(m, b)
}
func (m *ValueError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueError.Marshal(b, m, deterministic)
}
func (m *ValueError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueError.Merge(m, src)
}
func (m *ValueError) XXX_Size() int {
	return xxx_messageInfo_ValueError.Size(m)
}
func (m *ValueError) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueError.DiscardUnknown(m)
}

var xxx_messageInfo_ValueError proto.InternalMessageInfo

func (m *ValueError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ValueList struct {
	Items                []*ValueRaw `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ValueList) Reset()         { *m = ValueList{} }
func (m *ValueList) String() string { return proto.CompactTextString(m) }
func (*ValueList) ProtoMessage()    {}
func (*ValueList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d8b663a521ecf69, []int{1}
}

func (m *ValueList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueList.Unmarshal(m, b)
}
func (m *ValueList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueList.Marshal(b, m, deterministic)
}
func (m *ValueList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueList.Merge(m, src)
}
func (m *ValueList) XXX_Size() int {
	return xxx_messageInfo_ValueList.Size(m)
}
func (m *ValueList) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueList.DiscardUnknown(m)
}

var xxx_messageInfo_ValueList proto.InternalMessageInfo

func (m *ValueList) GetItems() []*ValueRaw {
	if m != nil {
		return m.Items
	}
	return nil
}

type ValueDict struct {
	Items                map[string]*ValueRaw `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ValueDict) Reset()         { *m = ValueDict{} }
func (m *ValueDict) String() string { return proto.CompactTextString(m) }
func (*ValueDict) ProtoMessage()    {}
func (*ValueDict) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d8b663a521ecf69, []int{2}
}

func (m *ValueDict) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueDict.Unmarshal(m, b)
}
func (m *ValueDict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueDict.Marshal(b, m, deterministic)
}
func (m *ValueDict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueDict.Merge(m, src)
}
func (m *ValueDict) XXX_Size() int {
	return xxx_messageInfo_ValueDict.Size(m)
}
func (m *ValueDict) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueDict.DiscardUnknown(m)
}

var xxx_messageInfo_ValueDict proto.InternalMessageInfo

func (m *ValueDict) GetItems() map[string]*ValueRaw {
	if m != nil {
		return m.Items
	}
	return nil
}

type ValueRaw struct {
	// Types that are valid to be assigned to Raw:
	//	*ValueRaw_Null
	//	*ValueRaw_Double
	//	*ValueRaw_Int64
	//	*ValueRaw_Bool
	//	*ValueRaw_String_
	//	*ValueRaw_Bytes
	//	*ValueRaw_List
	//	*ValueRaw_Dict
	Raw                  isValueRaw_Raw `protobuf_oneof:"raw"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ValueRaw) Reset()         { *m = ValueRaw{} }
func (m *ValueRaw) String() string { return proto.CompactTextString(m) }
func (*ValueRaw) ProtoMessage()    {}
func (*ValueRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d8b663a521ecf69, []int{3}
}

func (m *ValueRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueRaw.Unmarshal(m, b)
}
func (m *ValueRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueRaw.Marshal(b, m, deterministic)
}
func (m *ValueRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueRaw.Merge(m, src)
}
func (m *ValueRaw) XXX_Size() int {
	return xxx_messageInfo_ValueRaw.Size(m)
}
func (m *ValueRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueRaw.DiscardUnknown(m)
}

var xxx_messageInfo_ValueRaw proto.InternalMessageInfo

type isValueRaw_Raw interface {
	isValueRaw_Raw()
}

type ValueRaw_Null struct {
	Null NullValue `protobuf:"varint,1,opt,name=null,proto3,enum=NullValue,oneof"`
}

type ValueRaw_Double struct {
	Double float64 `protobuf:"fixed64,2,opt,name=double,proto3,oneof"`
}

type ValueRaw_Int64 struct {
	Int64 int64 `protobuf:"varint,3,opt,name=int64,proto3,oneof"`
}

type ValueRaw_Bool struct {
	Bool bool `protobuf:"varint,4,opt,name=bool,proto3,oneof"`
}

type ValueRaw_String_ struct {
	String_ string `protobuf:"bytes,5,opt,name=string,proto3,oneof"`
}

type ValueRaw_Bytes struct {
	Bytes []byte `protobuf:"bytes,6,opt,name=bytes,proto3,oneof"`
}

type ValueRaw_List struct {
	List *ValueList `protobuf:"bytes,7,opt,name=list,proto3,oneof"`
}

type ValueRaw_Dict struct {
	Dict *ValueDict `protobuf:"bytes,8,opt,name=dict,proto3,oneof"`
}

func (*ValueRaw_Null) isValueRaw_Raw() {}

func (*ValueRaw_Double) isValueRaw_Raw() {}

func (*ValueRaw_Int64) isValueRaw_Raw() {}

func (*ValueRaw_Bool) isValueRaw_Raw() {}

func (*ValueRaw_String_) isValueRaw_Raw() {}

func (*ValueRaw_Bytes) isValueRaw_Raw() {}

func (*ValueRaw_List) isValueRaw_Raw() {}

func (*ValueRaw_Dict) isValueRaw_Raw() {}

func (m *ValueRaw) GetRaw() isValueRaw_Raw {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *ValueRaw) GetNull() NullValue {
	if x, ok := m.GetRaw().(*ValueRaw_Null); ok {
		return x.Null
	}
	return NullValue_NULL_VALUE
}

func (m *ValueRaw) GetDouble() float64 {
	if x, ok := m.GetRaw().(*ValueRaw_Double); ok {
		return x.Double
	}
	return 0
}

func (m *ValueRaw) GetInt64() int64 {
	if x, ok := m.GetRaw().(*ValueRaw_Int64); ok {
		return x.Int64
	}
	return 0
}

func (m *ValueRaw) GetBool() bool {
	if x, ok := m.GetRaw().(*ValueRaw_Bool); ok {
		return x.Bool
	}
	return false
}

func (m *ValueRaw) GetString_() string {
	if x, ok := m.GetRaw().(*ValueRaw_String_); ok {
		return x.String_
	}
	return ""
}

func (m *ValueRaw) GetBytes() []byte {
	if x, ok := m.GetRaw().(*ValueRaw_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (m *ValueRaw) GetList() *ValueList {
	if x, ok := m.GetRaw().(*ValueRaw_List); ok {
		return x.List
	}
	return nil
}

func (m *ValueRaw) GetDict() *ValueDict {
	if x, ok := m.GetRaw().(*ValueRaw_Dict); ok {
		return x.Dict
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ValueRaw) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ValueRaw_Null)(nil),
		(*ValueRaw_Double)(nil),
		(*ValueRaw_Int64)(nil),
		(*ValueRaw_Bool)(nil),
		(*ValueRaw_String_)(nil),
		(*ValueRaw_Bytes)(nil),
		(*ValueRaw_List)(nil),
		(*ValueRaw_Dict)(nil),
	}
}

type Value struct {
	// Types that are valid to be assigned to Value:
	//	*Value_Error
	//	*Value_Raw
	//	*Value_DatasetUri
	//	*Value_CsvUri
	//	*Value_PickleUri
	//	*Value_PickleBlob
	//	*Value_PlasmaId
	//	*Value_ParquetUri
	Value                isValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d8b663a521ecf69, []int{4}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Value interface {
	isValue_Value()
}

type Value_Error struct {
	Error *ValueError `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type Value_Raw struct {
	Raw *ValueRaw `protobuf:"bytes,2,opt,name=raw,proto3,oneof"`
}

type Value_DatasetUri struct {
	DatasetUri string `protobuf:"bytes,3,opt,name=dataset_uri,json=datasetUri,proto3,oneof"`
}

type Value_CsvUri struct {
	CsvUri string `protobuf:"bytes,4,opt,name=csv_uri,json=csvUri,proto3,oneof"`
}

type Value_PickleUri struct {
	PickleUri string `protobuf:"bytes,5,opt,name=pickle_uri,json=pickleUri,proto3,oneof"`
}

type Value_PickleBlob struct {
	PickleBlob []byte `protobuf:"bytes,6,opt,name=pickle_blob,json=pickleBlob,proto3,oneof"`
}

type Value_PlasmaId struct {
	PlasmaId []byte `protobuf:"bytes,7,opt,name=plasma_id,json=plasmaId,proto3,oneof"`
}

type Value_ParquetUri struct {
	ParquetUri string `protobuf:"bytes,8,opt,name=parquet_uri,json=parquetUri,proto3,oneof"`
}

func (*Value_Error) isValue_Value() {}

func (*Value_Raw) isValue_Value() {}

func (*Value_DatasetUri) isValue_Value() {}

func (*Value_CsvUri) isValue_Value() {}

func (*Value_PickleUri) isValue_Value() {}

func (*Value_PickleBlob) isValue_Value() {}

func (*Value_PlasmaId) isValue_Value() {}

func (*Value_ParquetUri) isValue_Value() {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetError() *ValueError {
	if x, ok := m.GetValue().(*Value_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Value) GetRaw() *ValueRaw {
	if x, ok := m.GetValue().(*Value_Raw); ok {
		return x.Raw
	}
	return nil
}

func (m *Value) GetDatasetUri() string {
	if x, ok := m.GetValue().(*Value_DatasetUri); ok {
		return x.DatasetUri
	}
	return ""
}

func (m *Value) GetCsvUri() string {
	if x, ok := m.GetValue().(*Value_CsvUri); ok {
		return x.CsvUri
	}
	return ""
}

func (m *Value) GetPickleUri() string {
	if x, ok := m.GetValue().(*Value_PickleUri); ok {
		return x.PickleUri
	}
	return ""
}

func (m *Value) GetPickleBlob() []byte {
	if x, ok := m.GetValue().(*Value_PickleBlob); ok {
		return x.PickleBlob
	}
	return nil
}

func (m *Value) GetPlasmaId() []byte {
	if x, ok := m.GetValue().(*Value_PlasmaId); ok {
		return x.PlasmaId
	}
	return nil
}

func (m *Value) GetParquetUri() string {
	if x, ok := m.GetValue().(*Value_ParquetUri); ok {
		return x.ParquetUri
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_Error)(nil),
		(*Value_Raw)(nil),
		(*Value_DatasetUri)(nil),
		(*Value_CsvUri)(nil),
		(*Value_PickleUri)(nil),
		(*Value_PickleBlob)(nil),
		(*Value_PlasmaId)(nil),
		(*Value_ParquetUri)(nil),
	}
}

func init() {
	proto.RegisterEnum("NullValue", NullValue_name, NullValue_value)
	proto.RegisterType((*ValueError)(nil), "ValueError")
	proto.RegisterType((*ValueList)(nil), "ValueList")
	proto.RegisterType((*ValueDict)(nil), "ValueDict")
	proto.RegisterMapType((map[string]*ValueRaw)(nil), "ValueDict.ItemsEntry")
	proto.RegisterType((*ValueRaw)(nil), "ValueRaw")
	proto.RegisterType((*Value)(nil), "Value")
}

func init() { proto.RegisterFile("value.proto", fileDescriptor_6d8b663a521ecf69) }

var fileDescriptor_6d8b663a521ecf69 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xee, 0x6c, 0x9b, 0x36, 0x79, 0x23, 0xcb, 0x32, 0xa8, 0x8c, 0xca, 0xb2, 0xd9, 0x08, 0x12,
	0x54, 0x72, 0xa8, 0x22, 0xe2, 0xcd, 0x6a, 0xa1, 0x0b, 0x65, 0x0f, 0x81, 0xee, 0xc1, 0x4b, 0xc9,
	0xc7, 0xb0, 0x0c, 0x3b, 0x4d, 0xe2, 0xcc, 0xa4, 0x4b, 0x0f, 0xfe, 0x0f, 0xff, 0xa6, 0xff, 0x40,
	0xde, 0xc9, 0xa4, 0xeb, 0xc7, 0xad, 0xcf, 0xc7, 0x3c, 0xf3, 0xbe, 0xcf, 0x34, 0x10, 0xee, 0x73,
	0xd9, 0xf1, 0xb4, 0x55, 0x8d, 0x69, 0xe2, 0x57, 0x00, 0x37, 0x08, 0x97, 0x4a, 0x35, 0x8a, 0x32,
	0x98, 0xed, 0xb8, 0xd6, 0xf9, 0x2d, 0x67, 0x24, 0x22, 0x49, 0x90, 0x0d, 0x30, 0x7e, 0x0b, 0x81,
	0xf5, 0xad, 0x85, 0x36, 0xf4, 0x02, 0x3c, 0x61, 0xf8, 0x4e, 0x33, 0x12, 0x8d, 0x93, 0x70, 0x1e,
	0xa4, 0x56, 0xca, 0xf2, 0xfb, 0xac, 0xe7, 0xe3, 0x1f, 0xce, 0xfd, 0x55, 0x94, 0x86, 0xbe, 0xf9,
	0xdb, 0xfd, 0x24, 0x3d, 0x4a, 0xe9, 0x15, 0xf2, 0xcb, 0xda, 0xa8, 0x83, 0x3b, 0xf9, 0xfc, 0x0b,
	0xc0, 0x03, 0x49, 0xcf, 0x60, 0x7c, 0xc7, 0x0f, 0x6e, 0x16, 0xfc, 0x89, 0x57, 0xdb, 0xf1, 0xd9,
	0x49, 0x44, 0xfe, 0xb9, 0xda, 0xf2, 0x9f, 0x4e, 0x3e, 0x92, 0xf8, 0x17, 0x01, 0x7f, 0xe0, 0x69,
	0x04, 0x93, 0xba, 0x93, 0xd2, 0x86, 0x9c, 0xce, 0x21, 0xbd, 0xee, 0xa4, 0xb4, 0xe2, 0x6a, 0x94,
	0x59, 0x85, 0x32, 0x98, 0x56, 0x4d, 0x57, 0xc8, 0x3e, 0x94, 0xac, 0x46, 0x99, 0xc3, 0xf4, 0x29,
	0x78, 0xa2, 0x36, 0x1f, 0xde, 0xb3, 0x71, 0x44, 0x92, 0xf1, 0x6a, 0x94, 0xf5, 0x90, 0x3e, 0x86,
	0x49, 0xd1, 0x34, 0x92, 0x4d, 0x22, 0x92, 0xf8, 0x98, 0x83, 0x08, 0x73, 0xb4, 0x51, 0xa2, 0xbe,
	0x65, 0x1e, 0x0e, 0x8c, 0x39, 0x3d, 0xc6, 0x9c, 0xe2, 0x60, 0xb8, 0x66, 0xd3, 0x88, 0x24, 0x8f,
	0x30, 0xc7, 0x42, 0x9c, 0x4d, 0x0a, 0x6d, 0xd8, 0xcc, 0x2e, 0x03, 0xe9, 0xb1, 0x62, 0xcc, 0x44,
	0x05, 0x1d, 0x95, 0x28, 0x0d, 0xf3, 0xff, 0x74, 0x60, 0x77, 0xe8, 0x40, 0x65, 0xe1, 0xc1, 0x58,
	0xe5, 0xf7, 0xf1, 0xcf, 0x13, 0xf0, 0xac, 0x48, 0x5f, 0x82, 0xc7, 0xf1, 0x35, 0xed, 0xc6, 0xe1,
	0x3c, 0x4c, 0x1f, 0x1e, 0x18, 0x6f, 0xb6, 0x1a, 0x3d, 0xb7, 0xa7, 0xfe, 0x6b, 0x71, 0x35, 0xca,
	0x90, 0xa7, 0x97, 0x10, 0x56, 0xb9, 0xc9, 0x35, 0x37, 0xdb, 0x4e, 0x09, 0xbb, 0x3e, 0xee, 0x03,
	0x8e, 0xdc, 0x28, 0x41, 0x9f, 0xc1, 0xac, 0xd4, 0x7b, 0x2b, 0x4f, 0x86, 0x75, 0x4b, 0xbd, 0x47,
	0xe9, 0x02, 0xa0, 0x15, 0xe5, 0x9d, 0xe4, 0x56, 0x1d, 0xca, 0x08, 0x7a, 0x0e, 0x0d, 0x97, 0x10,
	0x3a, 0x43, 0x21, 0x9b, 0xe2, 0xd8, 0x8a, 0x3b, 0xb5, 0x90, 0x4d, 0x41, 0xcf, 0x21, 0x68, 0x65,
	0xae, 0x77, 0xf9, 0x56, 0x54, 0xb6, 0x1f, 0x34, 0xf8, 0x3d, 0x75, 0x55, 0xd9, 0x84, 0x5c, 0x7d,
	0xef, 0xdc, 0x80, 0xfe, 0x30, 0xa0, 0x23, 0x37, 0x4a, 0x2c, 0x66, 0xee, 0xaf, 0xf2, 0xfa, 0x05,
	0x04, 0xc7, 0x47, 0xa7, 0xa7, 0x00, 0xd7, 0x9b, 0xf5, 0x7a, 0x7b, 0xf3, 0x79, 0xbd, 0x59, 0x9e,
	0x8d, 0x16, 0xf0, 0xcd, 0x6f, 0x45, 0xcb, 0xa5, 0xa8, 0x79, 0x31, 0xb5, 0xdf, 0xc4, 0xbb, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xea, 0x10, 0x97, 0x47, 0x22, 0x03, 0x00, 0x00,
}
