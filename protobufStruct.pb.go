// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: protobufStruct.proto

package go_qq_protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeviceInfoBytesStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bootloader   *string `protobuf:"bytes,1,opt,name=Bootloader,proto3,oneof" json:"Bootloader,omitempty"`
	ProcVersion  *string `protobuf:"bytes,2,opt,name=ProcVersion,proto3,oneof" json:"ProcVersion,omitempty"`
	CodeName     *string `protobuf:"bytes,3,opt,name=CodeName,proto3,oneof" json:"CodeName,omitempty"`
	Incremental  *string `protobuf:"bytes,4,opt,name=Incremental,proto3,oneof" json:"Incremental,omitempty"`
	FingerPrint  *string `protobuf:"bytes,5,opt,name=FingerPrint,proto3,oneof" json:"FingerPrint,omitempty"`
	BootId       *string `protobuf:"bytes,6,opt,name=BootId,proto3,oneof" json:"BootId,omitempty"`
	AndroidId    *string `protobuf:"bytes,7,opt,name=AndroidId,proto3,oneof" json:"AndroidId,omitempty"`
	BaseBand     *string `protobuf:"bytes,8,opt,name=BaseBand,proto3,oneof" json:"BaseBand,omitempty"`
	InnerVersion *string `protobuf:"bytes,9,opt,name=InnerVersion,proto3,oneof" json:"InnerVersion,omitempty"`
}

func (x *DeviceInfoBytesStruct) Reset() {
	*x = DeviceInfoBytesStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufStruct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfoBytesStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfoBytesStruct) ProtoMessage() {}

func (x *DeviceInfoBytesStruct) ProtoReflect() protoreflect.Message {
	mi := &file_protobufStruct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfoBytesStruct.ProtoReflect.Descriptor instead.
func (*DeviceInfoBytesStruct) Descriptor() ([]byte, []int) {
	return file_protobufStruct_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceInfoBytesStruct) GetBootloader() string {
	if x != nil && x.Bootloader != nil {
		return *x.Bootloader
	}
	return ""
}

func (x *DeviceInfoBytesStruct) GetProcVersion() string {
	if x != nil && x.ProcVersion != nil {
		return *x.ProcVersion
	}
	return ""
}

func (x *DeviceInfoBytesStruct) GetCodeName() string {
	if x != nil && x.CodeName != nil {
		return *x.CodeName
	}
	return ""
}

func (x *DeviceInfoBytesStruct) GetIncremental() string {
	if x != nil && x.Incremental != nil {
		return *x.Incremental
	}
	return ""
}

func (x *DeviceInfoBytesStruct) GetFingerPrint() string {
	if x != nil && x.FingerPrint != nil {
		return *x.FingerPrint
	}
	return ""
}

func (x *DeviceInfoBytesStruct) GetBootId() string {
	if x != nil && x.BootId != nil {
		return *x.BootId
	}
	return ""
}

func (x *DeviceInfoBytesStruct) GetAndroidId() string {
	if x != nil && x.AndroidId != nil {
		return *x.AndroidId
	}
	return ""
}

func (x *DeviceInfoBytesStruct) GetBaseBand() string {
	if x != nil && x.BaseBand != nil {
		return *x.BaseBand
	}
	return ""
}

func (x *DeviceInfoBytesStruct) GetInnerVersion() string {
	if x != nil && x.InnerVersion != nil {
		return *x.InnerVersion
	}
	return ""
}

type D50RequestStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId                       *uint64  `protobuf:"varint,1,opt,name=AppId,proto3,oneof" json:"AppId,omitempty"`
	MaxPackageSize              *uint32  `protobuf:"varint,2,opt,name=MaxPackageSize,proto3,oneof" json:"MaxPackageSize,omitempty"`
	StartTime                   *uint32  `protobuf:"varint,3,opt,name=StartTime,proto3,oneof" json:"StartTime,omitempty"`
	StartIndex                  *uint32  `protobuf:"varint,4,opt,name=StartIndex,proto3,oneof" json:"StartIndex,omitempty"`
	RequestNum                  *uint32  `protobuf:"varint,5,opt,name=RequestNum,proto3,oneof" json:"RequestNum,omitempty"`
	UinList                     []uint64 `protobuf:"varint,6,rep,packed,name=UinList,proto3" json:"UinList,omitempty"`
	RequestMusicSwitch          *uint32  `protobuf:"varint,91001,opt,name=RequestMusicSwitch,proto3,oneof" json:"RequestMusicSwitch,omitempty"`
	RequestMutualmarkAlienation *uint32  `protobuf:"varint,101001,opt,name=RequestMutualmarkAlienation,proto3,oneof" json:"RequestMutualmarkAlienation,omitempty"`
	RequestMutualmarkScore      *uint32  `protobuf:"varint,141001,opt,name=RequestMutualmarkScore,proto3,oneof" json:"RequestMutualmarkScore,omitempty"`
	RequestKsingSwitch          *uint32  `protobuf:"varint,151001,opt,name=RequestKsingSwitch,proto3,oneof" json:"RequestKsingSwitch,omitempty"`
	RequestMutalmarkLbsShare    *uint32  `protobuf:"varint,181001,opt,name=RequestMutalmarkLbsShare,proto3,oneof" json:"RequestMutalmarkLbsShare,omitempty"`
}

func (x *D50RequestStruct) Reset() {
	*x = D50RequestStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufStruct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D50RequestStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D50RequestStruct) ProtoMessage() {}

func (x *D50RequestStruct) ProtoReflect() protoreflect.Message {
	mi := &file_protobufStruct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D50RequestStruct.ProtoReflect.Descriptor instead.
func (*D50RequestStruct) Descriptor() ([]byte, []int) {
	return file_protobufStruct_proto_rawDescGZIP(), []int{1}
}

func (x *D50RequestStruct) GetAppId() uint64 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *D50RequestStruct) GetMaxPackageSize() uint32 {
	if x != nil && x.MaxPackageSize != nil {
		return *x.MaxPackageSize
	}
	return 0
}

func (x *D50RequestStruct) GetStartTime() uint32 {
	if x != nil && x.StartTime != nil {
		return *x.StartTime
	}
	return 0
}

func (x *D50RequestStruct) GetStartIndex() uint32 {
	if x != nil && x.StartIndex != nil {
		return *x.StartIndex
	}
	return 0
}

func (x *D50RequestStruct) GetRequestNum() uint32 {
	if x != nil && x.RequestNum != nil {
		return *x.RequestNum
	}
	return 0
}

func (x *D50RequestStruct) GetUinList() []uint64 {
	if x != nil {
		return x.UinList
	}
	return nil
}

func (x *D50RequestStruct) GetRequestMusicSwitch() uint32 {
	if x != nil && x.RequestMusicSwitch != nil {
		return *x.RequestMusicSwitch
	}
	return 0
}

func (x *D50RequestStruct) GetRequestMutualmarkAlienation() uint32 {
	if x != nil && x.RequestMutualmarkAlienation != nil {
		return *x.RequestMutualmarkAlienation
	}
	return 0
}

func (x *D50RequestStruct) GetRequestMutualmarkScore() uint32 {
	if x != nil && x.RequestMutualmarkScore != nil {
		return *x.RequestMutualmarkScore
	}
	return 0
}

func (x *D50RequestStruct) GetRequestKsingSwitch() uint32 {
	if x != nil && x.RequestKsingSwitch != nil {
		return *x.RequestKsingSwitch
	}
	return 0
}

func (x *D50RequestStruct) GetRequestMutalmarkLbsShare() uint32 {
	if x != nil && x.RequestMutalmarkLbsShare != nil {
		return *x.RequestMutalmarkLbsShare
	}
	return 0
}

var File_protobufStruct_proto protoreflect.FileDescriptor

var file_protobufStruct_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x03, 0x0a, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x12, 0x23, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x74, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x42, 0x6f, 0x6f, 0x74, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x50, 0x72,
	0x6f, 0x63, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x43, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x08, 0x43, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a,
	0x0b, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0b, 0x46, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x42,
	0x6f, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x42,
	0x6f, 0x6f, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x41, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x09, 0x41,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x42,
	0x61, 0x73, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52,
	0x08, 0x42, 0x61, 0x73, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x08, 0x52, 0x0c, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x42, 0x6f, 0x6f, 0x74, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x50, 0x72, 0x6f, 0x63, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x43, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x42, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x42,
	0x61, 0x73, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe9, 0x05, 0x0a, 0x10, 0x44, 0x35, 0x30,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x19, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x01, 0x52, 0x0e, 0x4d, 0x61, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03, 0x52, 0x0a,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a,
	0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x04, 0x52, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x88,
	0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x07, 0x55, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x12,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x18, 0xf9, 0xc6, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x05, 0x52, 0x12, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x88, 0x01, 0x01, 0x12, 0x47, 0x0a, 0x1b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x75,
	0x74, 0x75, 0x61, 0x6c, 0x6d, 0x61, 0x72, 0x6b, 0x41, 0x6c, 0x69, 0x65, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x89, 0x95, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x06, 0x52, 0x1b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x6d, 0x61, 0x72, 0x6b, 0x41,
	0x6c, 0x69, 0x65, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x16,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x6d, 0x61, 0x72,
	0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0xc9, 0xcd, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x07,
	0x52, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x6d,
	0x61, 0x72, 0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x12, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x18, 0xd9, 0x9b, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x08, 0x52, 0x12, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4b, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x88,
	0x01, 0x01, 0x12, 0x41, 0x0a, 0x18, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x74,
	0x61, 0x6c, 0x6d, 0x61, 0x72, 0x6b, 0x4c, 0x62, 0x73, 0x53, 0x68, 0x61, 0x72, 0x65, 0x18, 0x89,
	0x86, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x09, 0x52, 0x18, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x75, 0x74, 0x61, 0x6c, 0x6d, 0x61, 0x72, 0x6b, 0x4c, 0x62, 0x73, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x64, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x4d, 0x61, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x42, 0x15,
	0x0a, 0x13, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x53,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x42, 0x1e, 0x0a, 0x1c, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x6d, 0x61, 0x72, 0x6b, 0x41, 0x6c, 0x69, 0x65, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x6d, 0x61, 0x72, 0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x73, 0x69, 0x6e,
	0x67, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x75, 0x74, 0x61, 0x6c, 0x6d, 0x61, 0x72, 0x6b, 0x4c, 0x62, 0x73, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x66, 0x69, 0x73, 0x68, 0x31, 0x32, 0x33,
	0x34, 0x35, 0x2f, 0x67, 0x6f, 0x2d, 0x71, 0x71, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobufStruct_proto_rawDescOnce sync.Once
	file_protobufStruct_proto_rawDescData = file_protobufStruct_proto_rawDesc
)

func file_protobufStruct_proto_rawDescGZIP() []byte {
	file_protobufStruct_proto_rawDescOnce.Do(func() {
		file_protobufStruct_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobufStruct_proto_rawDescData)
	})
	return file_protobufStruct_proto_rawDescData
}

var file_protobufStruct_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobufStruct_proto_goTypes = []interface{}{
	(*DeviceInfoBytesStruct)(nil), // 0: DeviceInfoBytesStruct
	(*D50RequestStruct)(nil),      // 1: D50RequestStruct
}
var file_protobufStruct_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobufStruct_proto_init() }
func file_protobufStruct_proto_init() {
	if File_protobufStruct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobufStruct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfoBytesStruct); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobufStruct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D50RequestStruct); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_protobufStruct_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_protobufStruct_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobufStruct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobufStruct_proto_goTypes,
		DependencyIndexes: file_protobufStruct_proto_depIdxs,
		MessageInfos:      file_protobufStruct_proto_msgTypes,
	}.Build()
	File_protobufStruct_proto = out.File
	file_protobufStruct_proto_rawDesc = nil
	file_protobufStruct_proto_goTypes = nil
	file_protobufStruct_proto_depIdxs = nil
}
