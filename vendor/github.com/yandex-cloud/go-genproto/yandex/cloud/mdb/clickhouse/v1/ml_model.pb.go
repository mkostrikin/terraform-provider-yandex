// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/clickhouse/v1/ml_model.proto

package clickhouse

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

type MlModelType int32

const (
	MlModelType_ML_MODEL_TYPE_UNSPECIFIED MlModelType = 0
	MlModelType_ML_MODEL_TYPE_CATBOOST    MlModelType = 1
)

var MlModelType_name = map[int32]string{
	0: "ML_MODEL_TYPE_UNSPECIFIED",
	1: "ML_MODEL_TYPE_CATBOOST",
}

var MlModelType_value = map[string]int32{
	"ML_MODEL_TYPE_UNSPECIFIED": 0,
	"ML_MODEL_TYPE_CATBOOST":    1,
}

func (x MlModelType) String() string {
	return proto.EnumName(MlModelType_name, int32(x))
}

func (MlModelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94ab55beb9dff409, []int{0}
}

type MlModel struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClusterId            string      `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Type                 MlModelType `protobuf:"varint,3,opt,name=type,proto3,enum=yandex.cloud.mdb.clickhouse.v1.MlModelType" json:"type,omitempty"`
	Uri                  string      `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MlModel) Reset()         { *m = MlModel{} }
func (m *MlModel) String() string { return proto.CompactTextString(m) }
func (*MlModel) ProtoMessage()    {}
func (*MlModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ab55beb9dff409, []int{0}
}

func (m *MlModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MlModel.Unmarshal(m, b)
}
func (m *MlModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MlModel.Marshal(b, m, deterministic)
}
func (m *MlModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MlModel.Merge(m, src)
}
func (m *MlModel) XXX_Size() int {
	return xxx_messageInfo_MlModel.Size(m)
}
func (m *MlModel) XXX_DiscardUnknown() {
	xxx_messageInfo_MlModel.DiscardUnknown(m)
}

var xxx_messageInfo_MlModel proto.InternalMessageInfo

func (m *MlModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MlModel) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *MlModel) GetType() MlModelType {
	if m != nil {
		return m.Type
	}
	return MlModelType_ML_MODEL_TYPE_UNSPECIFIED
}

func (m *MlModel) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func init() {
	proto.RegisterEnum("yandex.cloud.mdb.clickhouse.v1.MlModelType", MlModelType_name, MlModelType_value)
	proto.RegisterType((*MlModel)(nil), "yandex.cloud.mdb.clickhouse.v1.MlModel")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/clickhouse/v1/ml_model.proto", fileDescriptor_94ab55beb9dff409)
}

var fileDescriptor_94ab55beb9dff409 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xad, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0xcf, 0x4d, 0x49, 0xd2, 0x4f, 0xce, 0xc9,
	0x4c, 0xce, 0xce, 0xc8, 0x2f, 0x2d, 0x4e, 0xd5, 0x2f, 0x33, 0xd4, 0xcf, 0xcd, 0x89, 0xcf, 0xcd,
	0x4f, 0x49, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x83, 0x28, 0xd7, 0x03, 0x2b,
	0xd7, 0xcb, 0x4d, 0x49, 0xd2, 0x43, 0x28, 0xd7, 0x2b, 0x33, 0x54, 0xea, 0x67, 0xe4, 0x62, 0xf7,
	0xcd, 0xf1, 0x05, 0xe9, 0x10, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x64, 0xb9, 0xb8, 0x92, 0x73, 0x4a, 0x8b, 0x4b, 0x52, 0x8b,
	0xe2, 0x33, 0x53, 0x24, 0x98, 0xc0, 0x32, 0x9c, 0x50, 0x11, 0xcf, 0x14, 0x21, 0x7b, 0x2e, 0x96,
	0x92, 0xca, 0x82, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x6d, 0x3d, 0xfc, 0xb6, 0xe9,
	0x41, 0x6d, 0x0a, 0xa9, 0x2c, 0x48, 0x0d, 0x02, 0x6b, 0x14, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca,
	0x94, 0x60, 0x01, 0x1b, 0x0c, 0x62, 0x6a, 0x79, 0x70, 0x71, 0x23, 0x29, 0x13, 0x92, 0xe5, 0x92,
	0xf4, 0xf5, 0x89, 0xf7, 0xf5, 0x77, 0x71, 0xf5, 0x89, 0x0f, 0x89, 0x0c, 0x70, 0x8d, 0x0f, 0xf5,
	0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x10, 0x92, 0xe2, 0x12, 0x43,
	0x95, 0x76, 0x76, 0x0c, 0x71, 0xf2, 0xf7, 0x0f, 0x0e, 0x11, 0x60, 0x74, 0x2a, 0xe6, 0x52, 0x42,
	0x71, 0x4f, 0x62, 0x41, 0x26, 0xa6, 0x9b, 0xa2, 0x7c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0x21, 0xca, 0x75, 0x21, 0x61, 0x9b, 0x9e, 0xaf, 0x9b, 0x9e, 0x9a, 0x07,
	0x0e, 0x46, 0x7d, 0xfc, 0x81, 0x6e, 0x8d, 0xe0, 0x25, 0xb1, 0x81, 0x35, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xce, 0xb6, 0x78, 0x8c, 0xa8, 0x01, 0x00, 0x00,
}