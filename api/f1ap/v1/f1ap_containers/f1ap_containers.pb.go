//
//SPDX-FileCopyrightText: 2022-present Intel Corporation
//
//SPDX-License-Identifier: Apache-2.0

////////////////////// f1ap-containers.proto //////////////////////
// Protobuf generated from /f1ap_v1.asn1 by asn1c-0.9.29
// F1AP-Containers { itu-t(0) identified-organization(4) etsi(0) mobileDomain(0) ngran-access(22) modules(3) f1ap(3) version1(1) f1ap-Containers(5) }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.4
// source: api/f1ap/v1/f1ap_containers.proto

package f1apcontainersv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_commondatatypes"
	_ "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_constants"
	_ "github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
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

// sequence from f1ap_v1.asn1:10878
// Param F1AP-PROTOCOL-IES:IEsSetParam
// {ProtocolIE-Container001}
type ProtocolIeContainer001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:""
	Value []int32 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeContainer001) Reset() {
	*x = ProtocolIeContainer001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeContainer001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeContainer001) ProtoMessage() {}

func (x *ProtocolIeContainer001) ProtoReflect() protoreflect.Message {
	mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeContainer001.ProtoReflect.Descriptor instead.
func (*ProtocolIeContainer001) Descriptor() ([]byte, []int) {
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP(), []int{0}
}

func (x *ProtocolIeContainer001) GetValue() []int32 {
	if x != nil {
		return x.Value
	}
	return nil
}

// reference from f1ap_v1.asn1:10881
// Param F1AP-PROTOCOL-IES:IEsSetParam
// {ProtocolIE-SingleContainer001}
type ProtocolIeSingleContainer001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *ProtocolIeField001 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeSingleContainer001) Reset() {
	*x = ProtocolIeSingleContainer001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeSingleContainer001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeSingleContainer001) ProtoMessage() {}

func (x *ProtocolIeSingleContainer001) ProtoReflect() protoreflect.Message {
	mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeSingleContainer001.ProtoReflect.Descriptor instead.
func (*ProtocolIeSingleContainer001) Descriptor() ([]byte, []int) {
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP(), []int{1}
}

func (x *ProtocolIeSingleContainer001) GetValue() *ProtocolIeField001 {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from f1ap_v1.asn1:10884
// Param F1AP-PROTOCOL-IES:IEsSetParam
// {ProtocolIE-Field001}
type ProtocolIeField001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Criticality int32 `protobuf:"varint,2,opt,name=criticality,proto3" json:"criticality,omitempty"`
	Value       int32 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeField001) Reset() {
	*x = ProtocolIeField001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeField001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeField001) ProtoMessage() {}

func (x *ProtocolIeField001) ProtoReflect() protoreflect.Message {
	mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeField001.ProtoReflect.Descriptor instead.
func (*ProtocolIeField001) Descriptor() ([]byte, []int) {
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP(), []int{2}
}

func (x *ProtocolIeField001) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProtocolIeField001) GetCriticality() int32 {
	if x != nil {
		return x.Criticality
	}
	return 0
}

func (x *ProtocolIeField001) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// sequence from f1ap_v1.asn1:10897
// Param F1AP-PROTOCOL-IES-PAIR:IEsSetParam
// {ProtocolIE-ContainerPair}
type ProtocolIeContainerPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:""
	Value []int32 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeContainerPair) Reset() {
	*x = ProtocolIeContainerPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeContainerPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeContainerPair) ProtoMessage() {}

func (x *ProtocolIeContainerPair) ProtoReflect() protoreflect.Message {
	mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeContainerPair.ProtoReflect.Descriptor instead.
func (*ProtocolIeContainerPair) Descriptor() ([]byte, []int) {
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP(), []int{3}
}

func (x *ProtocolIeContainerPair) GetValue() []int32 {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from f1ap_v1.asn1:10900
// Param F1AP-PROTOCOL-IES-PAIR:IEsSetParam
// {ProtocolIE-FieldPair}
type ProtocolIeFieldPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProtocolIeFieldPair) Reset() {
	*x = ProtocolIeFieldPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeFieldPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeFieldPair) ProtoMessage() {}

func (x *ProtocolIeFieldPair) ProtoReflect() protoreflect.Message {
	mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeFieldPair.ProtoReflect.Descriptor instead.
func (*ProtocolIeFieldPair) Descriptor() ([]byte, []int) {
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP(), []int{4}
}

func (x *ProtocolIeFieldPair) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// sequence from f1ap_v1.asn1:10915
// Param F1AP-PROTOCOL-EXTENSION:ExtensionSetParam
// {ProtocolExtensionContainer001}
type ProtocolExtensionContainer001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:""
	Value []int32 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolExtensionContainer001) Reset() {
	*x = ProtocolExtensionContainer001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolExtensionContainer001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolExtensionContainer001) ProtoMessage() {}

func (x *ProtocolExtensionContainer001) ProtoReflect() protoreflect.Message {
	mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolExtensionContainer001.ProtoReflect.Descriptor instead.
func (*ProtocolExtensionContainer001) Descriptor() ([]byte, []int) {
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP(), []int{5}
}

func (x *ProtocolExtensionContainer001) GetValue() []int32 {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from f1ap_v1.asn1:10918
// Param F1AP-PROTOCOL-EXTENSION:ExtensionSetParam
// {ProtocolExtensionField001}
type ProtocolExtensionField001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Criticality    int32 `protobuf:"varint,2,opt,name=criticality,proto3" json:"criticality,omitempty"`
	ExtensionValue int32 `protobuf:"varint,3,opt,name=extension_value,json=extensionValue,proto3" json:"extension_value,omitempty"`
}

func (x *ProtocolExtensionField001) Reset() {
	*x = ProtocolExtensionField001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolExtensionField001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolExtensionField001) ProtoMessage() {}

func (x *ProtocolExtensionField001) ProtoReflect() protoreflect.Message {
	mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolExtensionField001.ProtoReflect.Descriptor instead.
func (*ProtocolExtensionField001) Descriptor() ([]byte, []int) {
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP(), []int{6}
}

func (x *ProtocolExtensionField001) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProtocolExtensionField001) GetCriticality() int32 {
	if x != nil {
		return x.Criticality
	}
	return 0
}

func (x *ProtocolExtensionField001) GetExtensionValue() int32 {
	if x != nil {
		return x.ExtensionValue
	}
	return 0
}

// sequence from f1ap_v1.asn1:10931
// Param F1AP-PRIVATE-IES:IEsSetParam
// {PrivateIE-Container001}
type PrivateIeContainer001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:""
	Value []int32 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *PrivateIeContainer001) Reset() {
	*x = PrivateIeContainer001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateIeContainer001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateIeContainer001) ProtoMessage() {}

func (x *PrivateIeContainer001) ProtoReflect() protoreflect.Message {
	mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateIeContainer001.ProtoReflect.Descriptor instead.
func (*PrivateIeContainer001) Descriptor() ([]byte, []int) {
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP(), []int{7}
}

func (x *PrivateIeContainer001) GetValue() []int32 {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from f1ap_v1.asn1:10934
// Param F1AP-PRIVATE-IES:IEsSetParam
// {PrivateIE-Field001}
type PrivateIeField001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Criticality int32 `protobuf:"varint,2,opt,name=criticality,proto3" json:"criticality,omitempty"`
	Value       int32 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PrivateIeField001) Reset() {
	*x = PrivateIeField001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateIeField001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateIeField001) ProtoMessage() {}

func (x *PrivateIeField001) ProtoReflect() protoreflect.Message {
	mi := &file_api_f1ap_v1_f1ap_containers_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateIeField001.ProtoReflect.Descriptor instead.
func (*PrivateIeField001) Descriptor() ([]byte, []int) {
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP(), []int{8}
}

func (x *PrivateIeField001) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PrivateIeField001) GetCriticality() int32 {
	if x != nil {
		return x.Criticality
	}
	return 0
}

func (x *PrivateIeField001) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_api_f1ap_v1_f1ap_containers_proto protoreflect.FileDescriptor

var file_api_f1ap_v1_f1ap_containers_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x31, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x31,
	0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x31, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x66, 0x31,
	0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x31, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x66, 0x31, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x31, 0x61, 0x70, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x73, 0x6e, 0x31,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e,
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x30, 0x30, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x51,
	0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x30, 0x30, 0x31, 0x12, 0x31,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x66, 0x31, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x49, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x30, 0x30, 0x31, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x5c, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x30, 0x30, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x72,
	0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x2f, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x25, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x50, 0x61, 0x69, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x1d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x30, 0x30, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x76,
	0x0a, 0x19, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x30, 0x30, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a,
	0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2d, 0x0a, 0x15, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x49, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x30, 0x30, 0x31, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5b, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x49, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x30, 0x30, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72,
	0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6f, 0x6e, 0x6f,
	0x73, 0x2d, 0x65, 0x32, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x31, 0x61, 0x70, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x31, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x3b, 0x66, 0x31, 0x61, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_f1ap_v1_f1ap_containers_proto_rawDescOnce sync.Once
	file_api_f1ap_v1_f1ap_containers_proto_rawDescData = file_api_f1ap_v1_f1ap_containers_proto_rawDesc
)

func file_api_f1ap_v1_f1ap_containers_proto_rawDescGZIP() []byte {
	file_api_f1ap_v1_f1ap_containers_proto_rawDescOnce.Do(func() {
		file_api_f1ap_v1_f1ap_containers_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_f1ap_v1_f1ap_containers_proto_rawDescData)
	})
	return file_api_f1ap_v1_f1ap_containers_proto_rawDescData
}

var file_api_f1ap_v1_f1ap_containers_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_f1ap_v1_f1ap_containers_proto_goTypes = []interface{}{
	(*ProtocolIeContainer001)(nil),        // 0: f1ap.v1.ProtocolIeContainer001
	(*ProtocolIeSingleContainer001)(nil),  // 1: f1ap.v1.ProtocolIeSingleContainer001
	(*ProtocolIeField001)(nil),            // 2: f1ap.v1.ProtocolIeField001
	(*ProtocolIeContainerPair)(nil),       // 3: f1ap.v1.ProtocolIeContainerPair
	(*ProtocolIeFieldPair)(nil),           // 4: f1ap.v1.ProtocolIeFieldPair
	(*ProtocolExtensionContainer001)(nil), // 5: f1ap.v1.ProtocolExtensionContainer001
	(*ProtocolExtensionField001)(nil),     // 6: f1ap.v1.ProtocolExtensionField001
	(*PrivateIeContainer001)(nil),         // 7: f1ap.v1.PrivateIeContainer001
	(*PrivateIeField001)(nil),             // 8: f1ap.v1.PrivateIeField001
}
var file_api_f1ap_v1_f1ap_containers_proto_depIdxs = []int32{
	2, // 0: f1ap.v1.ProtocolIeSingleContainer001.value:type_name -> f1ap.v1.ProtocolIeField001
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_f1ap_v1_f1ap_containers_proto_init() }
func file_api_f1ap_v1_f1ap_containers_proto_init() {
	if File_api_f1ap_v1_f1ap_containers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_f1ap_v1_f1ap_containers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeContainer001); i {
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
		file_api_f1ap_v1_f1ap_containers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeSingleContainer001); i {
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
		file_api_f1ap_v1_f1ap_containers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeField001); i {
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
		file_api_f1ap_v1_f1ap_containers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeContainerPair); i {
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
		file_api_f1ap_v1_f1ap_containers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeFieldPair); i {
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
		file_api_f1ap_v1_f1ap_containers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolExtensionContainer001); i {
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
		file_api_f1ap_v1_f1ap_containers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolExtensionField001); i {
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
		file_api_f1ap_v1_f1ap_containers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateIeContainer001); i {
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
		file_api_f1ap_v1_f1ap_containers_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateIeField001); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_f1ap_v1_f1ap_containers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_f1ap_v1_f1ap_containers_proto_goTypes,
		DependencyIndexes: file_api_f1ap_v1_f1ap_containers_proto_depIdxs,
		MessageInfos:      file_api_f1ap_v1_f1ap_containers_proto_msgTypes,
	}.Build()
	File_api_f1ap_v1_f1ap_containers_proto = out.File
	file_api_f1ap_v1_f1ap_containers_proto_rawDesc = nil
	file_api_f1ap_v1_f1ap_containers_proto_goTypes = nil
	file_api_f1ap_v1_f1ap_containers_proto_depIdxs = nil
}
