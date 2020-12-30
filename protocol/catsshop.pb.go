// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: catsshop.proto

package protocol

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

type Catparams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CatName   string `protobuf:"bytes,2,opt,name=CatName,proto3" json:"CatName,omitempty"`
	CatAge    int32  `protobuf:"varint,3,opt,name=CatAge,proto3" json:"CatAge,omitempty"`
	CatGender string `protobuf:"bytes,4,opt,name=CatGender,proto3" json:"CatGender,omitempty"`
	CatBreed  string `protobuf:"bytes,5,opt,name=CatBreed,proto3" json:"CatBreed,omitempty"`
	Error     string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Catparams) Reset() {
	*x = Catparams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catsshop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Catparams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Catparams) ProtoMessage() {}

func (x *Catparams) ProtoReflect() protoreflect.Message {
	mi := &file_catsshop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Catparams.ProtoReflect.Descriptor instead.
func (*Catparams) Descriptor() ([]byte, []int) {
	return file_catsshop_proto_rawDescGZIP(), []int{0}
}

func (x *Catparams) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Catparams) GetCatName() string {
	if x != nil {
		return x.CatName
	}
	return ""
}

func (x *Catparams) GetCatAge() int32 {
	if x != nil {
		return x.CatAge
	}
	return 0
}

func (x *Catparams) GetCatGender() string {
	if x != nil {
		return x.CatGender
	}
	return ""
}

func (x *Catparams) GetCatBreed() string {
	if x != nil {
		return x.CatBreed
	}
	return ""
}

func (x *Catparams) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Errmsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Errmsg) Reset() {
	*x = Errmsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catsshop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Errmsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Errmsg) ProtoMessage() {}

func (x *Errmsg) ProtoReflect() protoreflect.Message {
	mi := &file_catsshop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Errmsg.ProtoReflect.Descriptor instead.
func (*Errmsg) Descriptor() ([]byte, []int) {
	return file_catsshop_proto_rawDescGZIP(), []int{1}
}

func (x *Errmsg) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Jwtoken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Jwtoken) Reset() {
	*x = Jwtoken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catsshop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jwtoken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jwtoken) ProtoMessage() {}

func (x *Jwtoken) ProtoReflect() protoreflect.Message {
	mi := &file_catsshop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jwtoken.ProtoReflect.Descriptor instead.
func (*Jwtoken) Descriptor() ([]byte, []int) {
	return file_catsshop_proto_rawDescGZIP(), []int{2}
}

func (x *Jwtoken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Jwtoken) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Userparams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName string `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Admin    bool   `protobuf:"varint,2,opt,name=admin,proto3" json:"admin,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *Userparams) Reset() {
	*x = Userparams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catsshop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Userparams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Userparams) ProtoMessage() {}

func (x *Userparams) ProtoReflect() protoreflect.Message {
	mi := &file_catsshop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Userparams.ProtoReflect.Descriptor instead.
func (*Userparams) Descriptor() ([]byte, []int) {
	return file_catsshop_proto_rawDescGZIP(), []int{3}
}

func (x *Userparams) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *Userparams) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *Userparams) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Userparams) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Json struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json  string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Json) Reset() {
	*x = Json{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catsshop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Json) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Json) ProtoMessage() {}

func (x *Json) ProtoReflect() protoreflect.Message {
	mi := &file_catsshop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Json.ProtoReflect.Descriptor instead.
func (*Json) Descriptor() ([]byte, []int) {
	return file_catsshop_proto_rawDescGZIP(), []int{4}
}

func (x *Json) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

func (x *Json) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Operationparams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewOwnersNick string `protobuf:"bytes,1,opt,name=NewOwnersNick,proto3" json:"NewOwnersNick,omitempty"`
	CatID         int32  `protobuf:"varint,2,opt,name=CatID,proto3" json:"CatID,omitempty"`
	CatName       string `protobuf:"bytes,3,opt,name=CatName,proto3" json:"CatName,omitempty"`
	Date          string `protobuf:"bytes,4,opt,name=Date,proto3" json:"Date,omitempty"`
	Status        string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Error         string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Operationparams) Reset() {
	*x = Operationparams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catsshop_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operationparams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operationparams) ProtoMessage() {}

func (x *Operationparams) ProtoReflect() protoreflect.Message {
	mi := &file_catsshop_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operationparams.ProtoReflect.Descriptor instead.
func (*Operationparams) Descriptor() ([]byte, []int) {
	return file_catsshop_proto_rawDescGZIP(), []int{5}
}

func (x *Operationparams) GetNewOwnersNick() string {
	if x != nil {
		return x.NewOwnersNick
	}
	return ""
}

func (x *Operationparams) GetCatID() int32 {
	if x != nil {
		return x.CatID
	}
	return 0
}

func (x *Operationparams) GetCatName() string {
	if x != nil {
		return x.CatName
	}
	return ""
}

func (x *Operationparams) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Operationparams) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Operationparams) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_catsshop_proto protoreflect.FileDescriptor

var file_catsshop_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x74, 0x73, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x22, 0x9d, 0x01, 0x0a, 0x09, 0x63, 0x61, 0x74,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x74, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x43, 0x61, 0x74, 0x41, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61, 0x74, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x61, 0x74,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x42, 0x72, 0x65,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x74, 0x42, 0x72, 0x65,
	0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x1e, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d,
	0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x07, 0x6a, 0x77, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x74, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x4e,
	0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x4e, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x4e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x4e, 0x69, 0x63,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x61, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x43, 0x61, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x61, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x32, 0xf4, 0x05, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x73, 0x53, 0x68, 0x6f, 0x70,
	0x12, 0x30, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x12, 0x11, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x63, 0x61, 0x74, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x12,
	0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x63, 0x61, 0x74, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x65, 0x72, 0x72, 0x6d,
	0x73, 0x67, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x63, 0x61, 0x74, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x65, 0x72,
	0x72, 0x6d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x63, 0x61, 0x74, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x63, 0x61, 0x74,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x63, 0x61,
	0x74, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x70, 0x62, 0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x70, 0x62, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73,
	0x67, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x6a, 0x77, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x6a, 0x73,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x70, 0x62, 0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e,
	0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catsshop_proto_rawDescOnce sync.Once
	file_catsshop_proto_rawDescData = file_catsshop_proto_rawDesc
)

func file_catsshop_proto_rawDescGZIP() []byte {
	file_catsshop_proto_rawDescOnce.Do(func() {
		file_catsshop_proto_rawDescData = protoimpl.X.CompressGZIP(file_catsshop_proto_rawDescData)
	})
	return file_catsshop_proto_rawDescData
}

var file_catsshop_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_catsshop_proto_goTypes = []interface{}{
	(*Catparams)(nil),       // 0: protocol.catparams
	(*Errmsg)(nil),          // 1: protocol.errmsg
	(*Jwtoken)(nil),         // 2: protocol.jwtoken
	(*Userparams)(nil),      // 3: protocol.userparams
	(*Json)(nil),            // 4: protocol.json
	(*Operationparams)(nil), // 5: protocol.operationparams
}
var file_catsshop_proto_depIdxs = []int32{
	0,  // 0: protocol.CatsShop.CreateCat:input_type -> protocol.catparams
	0,  // 1: protocol.CatsShop.DeleteCat:input_type -> protocol.catparams
	0,  // 2: protocol.CatsShop.UpdateCat:input_type -> protocol.catparams
	0,  // 3: protocol.CatsShop.GetCat:input_type -> protocol.catparams
	0,  // 4: protocol.CatsShop.ListCats:input_type -> protocol.catparams
	3,  // 5: protocol.CatsShop.CreateUser:input_type -> protocol.userparams
	3,  // 6: protocol.CatsShop.DeleteUser:input_type -> protocol.userparams
	3,  // 7: protocol.CatsShop.UpdateUser:input_type -> protocol.userparams
	3,  // 8: protocol.CatsShop.Login:input_type -> protocol.userparams
	3,  // 9: protocol.CatsShop.ListUsers:input_type -> protocol.userparams
	5,  // 10: protocol.CatsShop.AddOperation:input_type -> protocol.operationparams
	5,  // 11: protocol.CatsShop.EditOperation:input_type -> protocol.operationparams
	5,  // 12: protocol.CatsShop.ListOperations:input_type -> protocol.operationparams
	5,  // 13: protocol.CatsShop.GetOperation:input_type -> protocol.operationparams
	1,  // 14: protocol.CatsShop.CreateCat:output_type -> protocol.errmsg
	1,  // 15: protocol.CatsShop.DeleteCat:output_type -> protocol.errmsg
	1,  // 16: protocol.CatsShop.UpdateCat:output_type -> protocol.errmsg
	0,  // 17: protocol.CatsShop.GetCat:output_type -> protocol.catparams
	4,  // 18: protocol.CatsShop.ListCats:output_type -> protocol.json
	1,  // 19: protocol.CatsShop.CreateUser:output_type -> protocol.errmsg
	1,  // 20: protocol.CatsShop.DeleteUser:output_type -> protocol.errmsg
	1,  // 21: protocol.CatsShop.UpdateUser:output_type -> protocol.errmsg
	2,  // 22: protocol.CatsShop.Login:output_type -> protocol.jwtoken
	4,  // 23: protocol.CatsShop.ListUsers:output_type -> protocol.json
	1,  // 24: protocol.CatsShop.AddOperation:output_type -> protocol.errmsg
	1,  // 25: protocol.CatsShop.EditOperation:output_type -> protocol.errmsg
	4,  // 26: protocol.CatsShop.ListOperations:output_type -> protocol.json
	5,  // 27: protocol.CatsShop.GetOperation:output_type -> protocol.operationparams
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_catsshop_proto_init() }
func file_catsshop_proto_init() {
	if File_catsshop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catsshop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Catparams); i {
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
		file_catsshop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Errmsg); i {
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
		file_catsshop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jwtoken); i {
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
		file_catsshop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Userparams); i {
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
		file_catsshop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Json); i {
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
		file_catsshop_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operationparams); i {
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
			RawDescriptor: file_catsshop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catsshop_proto_goTypes,
		DependencyIndexes: file_catsshop_proto_depIdxs,
		MessageInfos:      file_catsshop_proto_msgTypes,
	}.Build()
	File_catsshop_proto = out.File
	file_catsshop_proto_rawDesc = nil
	file_catsshop_proto_goTypes = nil
	file_catsshop_proto_depIdxs = nil
}