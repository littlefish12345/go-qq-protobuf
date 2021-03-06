// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobufStruct.proto

package go_qq_protobuf

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type DeviceInfoBytesStruct struct {
	Bootloader   string `protobuf:"bytes,1,opt,name=Bootloader,proto3" json:"Bootloader,omitempty"`
	ProcVersion  string `protobuf:"bytes,2,opt,name=ProcVersion,proto3" json:"ProcVersion,omitempty"`
	CodeName     string `protobuf:"bytes,3,opt,name=CodeName,proto3" json:"CodeName,omitempty"`
	Incremental  string `protobuf:"bytes,4,opt,name=Incremental,proto3" json:"Incremental,omitempty"`
	FingerPrint  string `protobuf:"bytes,5,opt,name=FingerPrint,proto3" json:"FingerPrint,omitempty"`
	BootId       string `protobuf:"bytes,6,opt,name=BootId,proto3" json:"BootId,omitempty"`
	AndroidId    string `protobuf:"bytes,7,opt,name=AndroidId,proto3" json:"AndroidId,omitempty"`
	BaseBand     string `protobuf:"bytes,8,opt,name=BaseBand,proto3" json:"BaseBand,omitempty"`
	InnerVersion string `protobuf:"bytes,9,opt,name=InnerVersion,proto3" json:"InnerVersion,omitempty"`
}

func (m *DeviceInfoBytesStruct) Reset()         { *m = DeviceInfoBytesStruct{} }
func (m *DeviceInfoBytesStruct) String() string { return proto.CompactTextString(m) }
func (*DeviceInfoBytesStruct) ProtoMessage()    {}
func (*DeviceInfoBytesStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_434ddb4d57c89a2a, []int{0}
}
func (m *DeviceInfoBytesStruct) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceInfoBytesStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceInfoBytesStruct.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceInfoBytesStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceInfoBytesStruct.Merge(m, src)
}
func (m *DeviceInfoBytesStruct) XXX_Size() int {
	return m.Size()
}
func (m *DeviceInfoBytesStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceInfoBytesStruct.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceInfoBytesStruct proto.InternalMessageInfo

func (m *DeviceInfoBytesStruct) GetBootloader() string {
	if m != nil {
		return m.Bootloader
	}
	return ""
}

func (m *DeviceInfoBytesStruct) GetProcVersion() string {
	if m != nil {
		return m.ProcVersion
	}
	return ""
}

func (m *DeviceInfoBytesStruct) GetCodeName() string {
	if m != nil {
		return m.CodeName
	}
	return ""
}

func (m *DeviceInfoBytesStruct) GetIncremental() string {
	if m != nil {
		return m.Incremental
	}
	return ""
}

func (m *DeviceInfoBytesStruct) GetFingerPrint() string {
	if m != nil {
		return m.FingerPrint
	}
	return ""
}

func (m *DeviceInfoBytesStruct) GetBootId() string {
	if m != nil {
		return m.BootId
	}
	return ""
}

func (m *DeviceInfoBytesStruct) GetAndroidId() string {
	if m != nil {
		return m.AndroidId
	}
	return ""
}

func (m *DeviceInfoBytesStruct) GetBaseBand() string {
	if m != nil {
		return m.BaseBand
	}
	return ""
}

func (m *DeviceInfoBytesStruct) GetInnerVersion() string {
	if m != nil {
		return m.InnerVersion
	}
	return ""
}

type D50RequestStruct struct {
	AppId                       uint64   `protobuf:"varint,1,opt,name=AppId,proto3" json:"AppId,omitempty"`
	MaxPackageSize              uint32   `protobuf:"varint,2,opt,name=MaxPackageSize,proto3" json:"MaxPackageSize,omitempty"`
	StartTime                   uint32   `protobuf:"varint,3,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	StartIndex                  uint32   `protobuf:"varint,4,opt,name=StartIndex,proto3" json:"StartIndex,omitempty"`
	RequestNum                  uint32   `protobuf:"varint,5,opt,name=RequestNum,proto3" json:"RequestNum,omitempty"`
	UinList                     []uint64 `protobuf:"varint,6,rep,packed,name=UinList,proto3" json:"UinList,omitempty"`
	RequestMusicSwitch          uint32   `protobuf:"varint,91001,opt,name=RequestMusicSwitch,proto3" json:"RequestMusicSwitch,omitempty"`
	RequestMutualmarkAlienation uint32   `protobuf:"varint,101001,opt,name=RequestMutualmarkAlienation,proto3" json:"RequestMutualmarkAlienation,omitempty"`
	RequestMutualmarkScore      uint32   `protobuf:"varint,141001,opt,name=RequestMutualmarkScore,proto3" json:"RequestMutualmarkScore,omitempty"`
	RequestKsingSwitch          uint32   `protobuf:"varint,151001,opt,name=RequestKsingSwitch,proto3" json:"RequestKsingSwitch,omitempty"`
	RequestMutalmarkLbsShare    uint32   `protobuf:"varint,181001,opt,name=RequestMutalmarkLbsShare,proto3" json:"RequestMutalmarkLbsShare,omitempty"`
}

func (m *D50RequestStruct) Reset()         { *m = D50RequestStruct{} }
func (m *D50RequestStruct) String() string { return proto.CompactTextString(m) }
func (*D50RequestStruct) ProtoMessage()    {}
func (*D50RequestStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_434ddb4d57c89a2a, []int{1}
}
func (m *D50RequestStruct) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *D50RequestStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_D50RequestStruct.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *D50RequestStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_D50RequestStruct.Merge(m, src)
}
func (m *D50RequestStruct) XXX_Size() int {
	return m.Size()
}
func (m *D50RequestStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_D50RequestStruct.DiscardUnknown(m)
}

var xxx_messageInfo_D50RequestStruct proto.InternalMessageInfo

func (m *D50RequestStruct) GetAppId() uint64 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *D50RequestStruct) GetMaxPackageSize() uint32 {
	if m != nil {
		return m.MaxPackageSize
	}
	return 0
}

func (m *D50RequestStruct) GetStartTime() uint32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *D50RequestStruct) GetStartIndex() uint32 {
	if m != nil {
		return m.StartIndex
	}
	return 0
}

func (m *D50RequestStruct) GetRequestNum() uint32 {
	if m != nil {
		return m.RequestNum
	}
	return 0
}

func (m *D50RequestStruct) GetUinList() []uint64 {
	if m != nil {
		return m.UinList
	}
	return nil
}

func (m *D50RequestStruct) GetRequestMusicSwitch() uint32 {
	if m != nil {
		return m.RequestMusicSwitch
	}
	return 0
}

func (m *D50RequestStruct) GetRequestMutualmarkAlienation() uint32 {
	if m != nil {
		return m.RequestMutualmarkAlienation
	}
	return 0
}

func (m *D50RequestStruct) GetRequestMutualmarkScore() uint32 {
	if m != nil {
		return m.RequestMutualmarkScore
	}
	return 0
}

func (m *D50RequestStruct) GetRequestKsingSwitch() uint32 {
	if m != nil {
		return m.RequestKsingSwitch
	}
	return 0
}

func (m *D50RequestStruct) GetRequestMutalmarkLbsShare() uint32 {
	if m != nil {
		return m.RequestMutalmarkLbsShare
	}
	return 0
}

func init() {
	proto.RegisterType((*DeviceInfoBytesStruct)(nil), "DeviceInfoBytesStruct")
	proto.RegisterType((*D50RequestStruct)(nil), "D50RequestStruct")
}

func init() { proto.RegisterFile("protobufStruct.proto", fileDescriptor_434ddb4d57c89a2a) }

var fileDescriptor_434ddb4d57c89a2a = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xe3, 0x26, 0x75, 0x93, 0xe9, 0x97, 0x4f, 0x68, 0x54, 0xaa, 0x11, 0x20, 0x2b, 0xca,
	0x02, 0x95, 0x45, 0xff, 0x40, 0xa9, 0xc4, 0x82, 0x4d, 0xdc, 0x0a, 0xc9, 0xa2, 0xad, 0x22, 0x1b,
	0x58, 0xb0, 0x9b, 0xd8, 0x37, 0xce, 0xa8, 0xf6, 0x4c, 0x32, 0x33, 0x86, 0xc2, 0x03, 0x20, 0x75,
	0xc7, 0x86, 0x15, 0x0f, 0xc1, 0x5b, 0x20, 0x58, 0x20, 0x75, 0x09, 0x3b, 0x94, 0x6c, 0x78, 0x05,
	0x76, 0xc8, 0x63, 0xa7, 0x71, 0x81, 0xb2, 0xf3, 0xf9, 0x9d, 0x7b, 0xaf, 0x8e, 0x8f, 0x65, 0xb4,
	0x36, 0x96, 0x42, 0x8b, 0x41, 0x36, 0x0c, 0xb4, 0xcc, 0x42, 0xbd, 0x65, 0x64, 0xf7, 0xc3, 0x12,
	0xba, 0x7e, 0x00, 0x2f, 0x58, 0x08, 0x1e, 0x1f, 0x0a, 0xf7, 0x95, 0x06, 0x55, 0xf8, 0xd8, 0x41,
	0xc8, 0x15, 0x42, 0x27, 0x82, 0x46, 0x20, 0x89, 0xd5, 0xb1, 0x36, 0x5a, 0x7e, 0x85, 0xe0, 0x0e,
	0x5a, 0xed, 0x4b, 0x11, 0x3e, 0x03, 0xa9, 0x98, 0xe0, 0x64, 0xc9, 0x0c, 0x54, 0x11, 0xbe, 0x81,
	0x9a, 0xfb, 0x22, 0x82, 0x63, 0x9a, 0x02, 0xa9, 0x1b, 0xfb, 0x42, 0xe7, 0xdb, 0x1e, 0x0f, 0x25,
	0xa4, 0xc0, 0x35, 0x4d, 0x48, 0xa3, 0xd8, 0xae, 0xa0, 0x7c, 0xe2, 0x11, 0xe3, 0x31, 0xc8, 0xbe,
	0x64, 0x5c, 0x93, 0xe5, 0x62, 0xa2, 0x82, 0xf0, 0x3a, 0xb2, 0xf3, 0x3c, 0x5e, 0x44, 0x6c, 0x63,
	0x96, 0x0a, 0xdf, 0x42, 0xad, 0x1e, 0x8f, 0xa4, 0x60, 0x91, 0x17, 0x91, 0x15, 0x63, 0x2d, 0x40,
	0x9e, 0xca, 0xa5, 0x0a, 0x5c, 0xca, 0x23, 0xd2, 0x2c, 0x52, 0xcd, 0x35, 0xee, 0xa2, 0xff, 0x3c,
	0xce, 0x41, 0xce, 0x5f, 0xaa, 0x65, 0xfc, 0x4b, 0xac, 0xfb, 0xa3, 0x8e, 0xae, 0x1d, 0xec, 0xed,
	0xf8, 0x30, 0xc9, 0x40, 0xe9, 0xb2, 0xac, 0x35, 0xb4, 0xdc, 0x1b, 0x8f, 0xbd, 0xc8, 0xf4, 0xd4,
	0xf0, 0x0b, 0x81, 0x6f, 0xa3, 0xff, 0x8f, 0xe8, 0x69, 0x9f, 0x86, 0x27, 0x34, 0x86, 0x80, 0xbd,
	0x06, 0xd3, 0x52, 0xdb, 0xff, 0x8d, 0xe6, 0x81, 0x03, 0x4d, 0xa5, 0x7e, 0xc2, 0xca, 0xa6, 0xda,
	0xfe, 0x02, 0xe4, 0x1f, 0xc2, 0x08, 0x8f, 0x47, 0x70, 0x6a, 0x9a, 0x6a, 0xfb, 0x15, 0x92, 0xfb,
	0x65, 0x98, 0xe3, 0x2c, 0x35, 0x3d, 0xb5, 0xfd, 0x0a, 0xc1, 0x04, 0xad, 0x3c, 0x65, 0xfc, 0x90,
	0x29, 0x4d, 0xec, 0x4e, 0x7d, 0xa3, 0xe1, 0xcf, 0x25, 0xde, 0x41, 0xb8, 0x9c, 0x3b, 0xca, 0x14,
	0x0b, 0x83, 0x97, 0x4c, 0x87, 0x23, 0xf2, 0xf3, 0x63, 0x71, 0xe3, 0x2f, 0x1e, 0x76, 0xd1, 0xcd,
	0x0b, 0xaa, 0x33, 0x9a, 0xa4, 0x54, 0x9e, 0xf4, 0x12, 0x06, 0x9c, 0xea, 0xbc, 0xaf, 0xb3, 0x77,
	0xb6, 0x59, 0xfd, 0xd7, 0x10, 0x7e, 0x80, 0xd6, 0xff, 0xb0, 0x83, 0x50, 0x48, 0x20, 0x9f, 0xbf,
	0x34, 0xcd, 0xfa, 0x15, 0x7e, 0x25, 0xef, 0x63, 0xc5, 0x78, 0x5c, 0xe6, 0xfd, 0xf6, 0xbe, 0x75,
	0x29, 0x6f, 0xc5, 0xc3, 0x0f, 0x11, 0x59, 0xdc, 0x2a, 0x4e, 0x1d, 0x0e, 0x54, 0x30, 0xa2, 0x12,
	0xc8, 0xd9, 0x9b, 0x55, 0xb3, 0x77, 0xe5, 0x84, 0xbb, 0xff, 0x69, 0xea, 0x58, 0xe7, 0x53, 0xc7,
	0xfa, 0x3e, 0x75, 0xac, 0xb7, 0x33, 0xa7, 0x76, 0x3e, 0x73, 0x6a, 0x5f, 0x67, 0x4e, 0xed, 0xf9,
	0x9d, 0x98, 0xe9, 0x51, 0x36, 0xd8, 0x0a, 0x45, 0xba, 0x9d, 0x30, 0xad, 0x13, 0x18, 0x32, 0x35,
	0xba, 0x7b, 0x6f, 0xf7, 0xfe, 0xde, 0x76, 0x2c, 0x36, 0x27, 0x93, 0xcd, 0xf9, 0xdf, 0x36, 0xb0,
	0xcd, 0xd3, 0xee, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x2c, 0x95, 0xee, 0x80, 0x03, 0x00,
	0x00,
}

func (m *DeviceInfoBytesStruct) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceInfoBytesStruct) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceInfoBytesStruct) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InnerVersion) > 0 {
		i -= len(m.InnerVersion)
		copy(dAtA[i:], m.InnerVersion)
		i = encodeVarintProtobufStruct(dAtA, i, uint64(len(m.InnerVersion)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.BaseBand) > 0 {
		i -= len(m.BaseBand)
		copy(dAtA[i:], m.BaseBand)
		i = encodeVarintProtobufStruct(dAtA, i, uint64(len(m.BaseBand)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.AndroidId) > 0 {
		i -= len(m.AndroidId)
		copy(dAtA[i:], m.AndroidId)
		i = encodeVarintProtobufStruct(dAtA, i, uint64(len(m.AndroidId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.BootId) > 0 {
		i -= len(m.BootId)
		copy(dAtA[i:], m.BootId)
		i = encodeVarintProtobufStruct(dAtA, i, uint64(len(m.BootId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.FingerPrint) > 0 {
		i -= len(m.FingerPrint)
		copy(dAtA[i:], m.FingerPrint)
		i = encodeVarintProtobufStruct(dAtA, i, uint64(len(m.FingerPrint)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Incremental) > 0 {
		i -= len(m.Incremental)
		copy(dAtA[i:], m.Incremental)
		i = encodeVarintProtobufStruct(dAtA, i, uint64(len(m.Incremental)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CodeName) > 0 {
		i -= len(m.CodeName)
		copy(dAtA[i:], m.CodeName)
		i = encodeVarintProtobufStruct(dAtA, i, uint64(len(m.CodeName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ProcVersion) > 0 {
		i -= len(m.ProcVersion)
		copy(dAtA[i:], m.ProcVersion)
		i = encodeVarintProtobufStruct(dAtA, i, uint64(len(m.ProcVersion)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Bootloader) > 0 {
		i -= len(m.Bootloader)
		copy(dAtA[i:], m.Bootloader)
		i = encodeVarintProtobufStruct(dAtA, i, uint64(len(m.Bootloader)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *D50RequestStruct) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *D50RequestStruct) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *D50RequestStruct) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestMutalmarkLbsShare != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.RequestMutalmarkLbsShare))
		i--
		dAtA[i] = 0x58
		i--
		dAtA[i] = 0xb0
		i--
		dAtA[i] = 0xc8
	}
	if m.RequestKsingSwitch != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.RequestKsingSwitch))
		i--
		dAtA[i] = 0x49
		i--
		dAtA[i] = 0xdd
		i--
		dAtA[i] = 0xc8
	}
	if m.RequestMutualmarkScore != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.RequestMutualmarkScore))
		i--
		dAtA[i] = 0x44
		i--
		dAtA[i] = 0xec
		i--
		dAtA[i] = 0xc8
	}
	if m.RequestMutualmarkAlienation != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.RequestMutualmarkAlienation))
		i--
		dAtA[i] = 0x31
		i--
		dAtA[i] = 0xa8
		i--
		dAtA[i] = 0xc8
	}
	if m.RequestMusicSwitch != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.RequestMusicSwitch))
		i--
		dAtA[i] = 0x2c
		i--
		dAtA[i] = 0xb7
		i--
		dAtA[i] = 0xc8
	}
	if len(m.UinList) > 0 {
		dAtA2 := make([]byte, len(m.UinList)*10)
		var j1 int
		for _, num := range m.UinList {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintProtobufStruct(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x32
	}
	if m.RequestNum != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.RequestNum))
		i--
		dAtA[i] = 0x28
	}
	if m.StartIndex != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.StartIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.StartTime != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxPackageSize != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.MaxPackageSize))
		i--
		dAtA[i] = 0x10
	}
	if m.AppId != 0 {
		i = encodeVarintProtobufStruct(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtobufStruct(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtobufStruct(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceInfoBytesStruct) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bootloader)
	if l > 0 {
		n += 1 + l + sovProtobufStruct(uint64(l))
	}
	l = len(m.ProcVersion)
	if l > 0 {
		n += 1 + l + sovProtobufStruct(uint64(l))
	}
	l = len(m.CodeName)
	if l > 0 {
		n += 1 + l + sovProtobufStruct(uint64(l))
	}
	l = len(m.Incremental)
	if l > 0 {
		n += 1 + l + sovProtobufStruct(uint64(l))
	}
	l = len(m.FingerPrint)
	if l > 0 {
		n += 1 + l + sovProtobufStruct(uint64(l))
	}
	l = len(m.BootId)
	if l > 0 {
		n += 1 + l + sovProtobufStruct(uint64(l))
	}
	l = len(m.AndroidId)
	if l > 0 {
		n += 1 + l + sovProtobufStruct(uint64(l))
	}
	l = len(m.BaseBand)
	if l > 0 {
		n += 1 + l + sovProtobufStruct(uint64(l))
	}
	l = len(m.InnerVersion)
	if l > 0 {
		n += 1 + l + sovProtobufStruct(uint64(l))
	}
	return n
}

func (m *D50RequestStruct) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppId != 0 {
		n += 1 + sovProtobufStruct(uint64(m.AppId))
	}
	if m.MaxPackageSize != 0 {
		n += 1 + sovProtobufStruct(uint64(m.MaxPackageSize))
	}
	if m.StartTime != 0 {
		n += 1 + sovProtobufStruct(uint64(m.StartTime))
	}
	if m.StartIndex != 0 {
		n += 1 + sovProtobufStruct(uint64(m.StartIndex))
	}
	if m.RequestNum != 0 {
		n += 1 + sovProtobufStruct(uint64(m.RequestNum))
	}
	if len(m.UinList) > 0 {
		l = 0
		for _, e := range m.UinList {
			l += sovProtobufStruct(uint64(e))
		}
		n += 1 + sovProtobufStruct(uint64(l)) + l
	}
	if m.RequestMusicSwitch != 0 {
		n += 3 + sovProtobufStruct(uint64(m.RequestMusicSwitch))
	}
	if m.RequestMutualmarkAlienation != 0 {
		n += 3 + sovProtobufStruct(uint64(m.RequestMutualmarkAlienation))
	}
	if m.RequestMutualmarkScore != 0 {
		n += 3 + sovProtobufStruct(uint64(m.RequestMutualmarkScore))
	}
	if m.RequestKsingSwitch != 0 {
		n += 3 + sovProtobufStruct(uint64(m.RequestKsingSwitch))
	}
	if m.RequestMutalmarkLbsShare != 0 {
		n += 3 + sovProtobufStruct(uint64(m.RequestMutalmarkLbsShare))
	}
	return n
}

func sovProtobufStruct(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtobufStruct(x uint64) (n int) {
	return sovProtobufStruct(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceInfoBytesStruct) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtobufStruct
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
			return fmt.Errorf("proto: DeviceInfoBytesStruct: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceInfoBytesStruct: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bootloader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
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
				return ErrInvalidLengthProtobufStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtobufStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bootloader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
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
				return ErrInvalidLengthProtobufStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtobufStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProcVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
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
				return ErrInvalidLengthProtobufStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtobufStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Incremental", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
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
				return ErrInvalidLengthProtobufStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtobufStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Incremental = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FingerPrint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
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
				return ErrInvalidLengthProtobufStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtobufStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FingerPrint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BootId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
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
				return ErrInvalidLengthProtobufStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtobufStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BootId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AndroidId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
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
				return ErrInvalidLengthProtobufStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtobufStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AndroidId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseBand", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
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
				return ErrInvalidLengthProtobufStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtobufStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseBand = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InnerVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
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
				return ErrInvalidLengthProtobufStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtobufStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InnerVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtobufStruct(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtobufStruct
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
func (m *D50RequestStruct) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtobufStruct
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
			return fmt.Errorf("proto: D50RequestStruct: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: D50RequestStruct: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPackageSize", wireType)
			}
			m.MaxPackageSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPackageSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartIndex", wireType)
			}
			m.StartIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestNum", wireType)
			}
			m.RequestNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestNum |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProtobufStruct
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.UinList = append(m.UinList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProtobufStruct
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProtobufStruct
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProtobufStruct
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.UinList) == 0 {
					m.UinList = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProtobufStruct
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.UinList = append(m.UinList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field UinList", wireType)
			}
		case 91001:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMusicSwitch", wireType)
			}
			m.RequestMusicSwitch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestMusicSwitch |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 101001:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMutualmarkAlienation", wireType)
			}
			m.RequestMutualmarkAlienation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestMutualmarkAlienation |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 141001:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMutualmarkScore", wireType)
			}
			m.RequestMutualmarkScore = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestMutualmarkScore |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 151001:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestKsingSwitch", wireType)
			}
			m.RequestKsingSwitch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestKsingSwitch |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 181001:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMutalmarkLbsShare", wireType)
			}
			m.RequestMutalmarkLbsShare = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestMutalmarkLbsShare |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtobufStruct(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtobufStruct
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
func skipProtobufStruct(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtobufStruct
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
					return 0, ErrIntOverflowProtobufStruct
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
					return 0, ErrIntOverflowProtobufStruct
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
				return 0, ErrInvalidLengthProtobufStruct
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProtobufStruct
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProtobufStruct
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProtobufStruct        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtobufStruct          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProtobufStruct = fmt.Errorf("proto: unexpected end of group")
)
